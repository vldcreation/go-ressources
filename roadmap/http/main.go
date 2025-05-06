package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

// --- Configuration & Options ---

// RetryOptions holds configuration for the retry mechanism.
type RetryOptions struct {
	MaxRetries     int
	InitialBackoff time.Duration
	MaxBackoff     time.Duration
	Multiplier     float64
	JitterFactor   float64 // e.g., 0.1 for 10% jitter. (delay ± (delay * JitterFactor / 2))
}

// DefaultRetryOptions provides some sensible defaults.
var DefaultRetryOptions = RetryOptions{
	MaxRetries:     3,
	InitialBackoff: 200 * time.Millisecond,
	MaxBackoff:     5 * time.Second,
	Multiplier:     2.0,
	JitterFactor:   0.2,
}

// --- HTTP Client Creation ---

func createCustomHTTPClient() *http.Client {
	// Configure the transport
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment, // Respect proxy settings
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second, // Connection timeout
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,              // Max total idle connections
		MaxIdleConnsPerHost:   10,               // Max idle connections per host
		IdleConnTimeout:       90 * time.Second, // Timeout for idle connections
		TLSHandshakeTimeout:   10 * time.Second, // TLS handshake timeout
		ExpectContinueTimeout: 1 * time.Second,  // Timeout for waiting for 100 Continue

		// Basic TLS Configuration
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			// In a real scenario, you might load custom CAs or set other TLS properties:
			// RootCAs: caCertPool,
			// Certificates: []tls.Certificate{clientCert},
		},
	}

	// Create the client with the custom transport and an overall request timeout
	client := &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second, // Overall timeout for a single Do call (including redirects)
	}
	return client
}

// --- Retryable Logic Helpers ---

// isRetryableError checks if an error suggests a retry is warranted.
// This includes network errors like timeouts, temporary DNS issues, etc.
func isRetryableError(err error) bool {
	if err == nil {
		return false
	}
	var netErr net.Error
	if errors.As(err, &netErr) && (netErr.Timeout() || netErr.Temporary()) {
		return true
	}

	// Specific errors that might be considered retryable
	// (This can be highly dependent on the services you're calling)
	if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
		// Sometimes EOF or UnexpectedEOF can be transient network issues
		return true
	}
	// syscall.ECONNREFUSED, syscall.ECONNRESET could also be candidates under certain conditions

	return false
}

// isRetryableStatusCode checks if an HTTP status code indicates a transient server error.
func isRetryableStatusCode(statusCode int) bool {
	switch statusCode {
	case http.StatusServiceUnavailable, // 503
		http.StatusBadGateway,      // 502
		http.StatusGatewayTimeout,  // 504
		http.StatusTooManyRequests: // 429 (often needs special handling for Retry-After header)
		return true
	default:
		return false
	}
}

// --- Core Request with Retry Function ---

// DoRequestWithRetries performs an HTTP request with a retry mechanism.
func DoRequestWithRetries(
	ctx context.Context,
	client *http.Client,
	req *http.Request, // The original request. It will be cloned for each attempt.
	opts RetryOptions,
) (*http.Response, error) {
	var lastErr error
	var resp *http.Response
	_ = resp

	// Ensure the request uses the provided context
	// req = req.WithContext(ctx) // req.Clone below will also use this context

	currentBackoff := opts.InitialBackoff

	for attempt := 0; attempt <= opts.MaxRetries; attempt++ {
		// Check for context cancellation *before* this attempt.
		select {
		case <-ctx.Done():
			log.Printf("[Attempt %d] Context cancelled before starting: %v", attempt, ctx.Err())
			if lastErr != nil {
				return nil, fmt.Errorf("context cancelled during retries (attempt %d): %w (last error: %v)", attempt, ctx.Err(), lastErr)
			}
			return nil, ctx.Err()
		default:
		}

		if attempt > 0 {
			// Calculate jitter: delay ± (delay * JitterFactor / 2)
			jitter := time.Duration(rand.Float64()*float64(currentBackoff)*opts.JitterFactor) - (time.Duration(float64(currentBackoff)*opts.JitterFactor) / 2)
			actualDelay := currentBackoff + jitter
			if actualDelay < 0 {
				actualDelay = 0 // Ensure delay is not negative
			}

			log.Printf("[Attempt %d] Previous attempt failed. Retrying in %v...", attempt, actualDelay)

			select {
			case <-time.After(actualDelay):
				// Waited successfully
			case <-ctx.Done():
				log.Printf("[Attempt %d] Context cancelled during backoff: %v", attempt, ctx.Err())
				// Return context error, possibly wrapping the last known operational error
				return nil, fmt.Errorf("context cancelled during retry backoff (attempt %d): %w (last error: %v)", attempt, ctx.Err(), lastErr)
			}

			// Increase backoff for the next potential retry
			currentBackoff = time.Duration(float64(currentBackoff) * opts.Multiplier)
			if currentBackoff > opts.MaxBackoff {
				currentBackoff = opts.MaxBackoff
			}
		}

		// Clone the request for each attempt. This is crucial because:
		// 1. The request body (if any) might have been consumed in a previous attempt.
		//    req.Clone() handles body duplication correctly if req.GetBody is set
		//    or if the body is one of a few known replayable types (bytes.Reader, strings.Reader).
		//    For production, ensure req.GetBody is properly set for complex bodies.
		// 2. It associates the most current context with this specific attempt's clone.
		clonedReq := req.Clone(ctx)

		// For POST/PUT with bodies not handled by default by req.Clone (e.g. io.Pipe),
		// you would need to ensure req.GetBody is set on the *original* request, or manually reset the body.
		// Example: if req.Body was from `io.NopCloser(bytes.NewReader(someBytes))`,
		// `GetBody` would be set by `http.NewRequest`. If it was a custom stream, more care is needed.
		if req.Body != nil && req.GetBody == nil {
			// This is a simplification. In real-world scenarios, if GetBody isn't set for a non-nil body,
			// you need a strategy to "reset" or "re-create" the body for the clonedReq.
			// For common types like bytes.Reader, strings.Reader, http.NewRequest sets GetBody.
			// If using a streaming body that can only be read once, this will fail on retries
			// unless GetBody is correctly implemented.
			// log.Println("[Warn] Request body is present but GetBody is nil. Retries might fail if body is not replayable.")
		}

		resp, err := client.Do(clonedReq)
		lastErr = err // Store the error from this attempt

		// Check for context cancellation *immediately* after the call returns.
		// The client.Do might return an error *because* of cancellation, or it might
		// return success just as cancellation occurs. Prioritize context error.
		select {
		case <-ctx.Done():
			log.Printf("[Attempt %d] Context cancelled immediately after Do: %v", attempt, ctx.Err())
			if resp != nil {
				io.Copy(io.Discard, resp.Body) // Drain and close body
				resp.Body.Close()
			}
			return nil, ctx.Err() // Context error is primary
		default:
			// Context not cancelled, proceed with normal error handling
		}

		if lastErr != nil {
			log.Printf("[Attempt %d] Request failed with client error: %v", attempt, lastErr)
			if !isRetryableError(lastErr) {
				log.Printf("[Attempt %d] Non-retryable client error: %v. Not retrying.", attempt, lastErr)
				// Close response body if a response was somehow received despite the error
				if resp != nil {
					io.Copy(io.Discard, resp.Body)
					resp.Body.Close()
				}
				return nil, lastErr // Non-retryable error, return it
			}
			// It's a retryable client error, continue to the next attempt (backoff will apply)
			if resp != nil { // Should close/drain body even if there was an error
				io.Copy(io.Discard, resp.Body)
				resp.Body.Close()
			}
			continue // Go to next retry iteration
		}

		// No client error from client.Do(), check response status code
		log.Printf("[Attempt %d] Request successful, Status: %s", attempt, resp.Status)
		if !isRetryableStatusCode(resp.StatusCode) {
			// Success or a non-retryable status code (e.g., 2xx, 4xx that isn't 429)
			return resp, nil
		}

		// It's a retryable status code (e.g., 503, 429)
		log.Printf("[Attempt %d] Received retryable status code: %d. Will retry.", attempt, resp.StatusCode)
		lastErr = fmt.Errorf("server returned a retryable status code: %d", resp.StatusCode)

		// IMPORTANT: Consume and close the response body before retrying,
		// otherwise, the connection might not be released back to the pool.
		io.Copy(io.Discard, resp.Body) // Read to EOF to allow connection reuse.
		resp.Body.Close()
		// Continue to next retry iteration
	}

	// If loop finishes, all retries were exhausted
	log.Printf("Max retries (%d) reached. Last error: %v", opts.MaxRetries, lastErr)
	// Return the last error encountered
	return nil, fmt.Errorf("max retries (%d) reached: %w", opts.MaxRetries, lastErr)
}

// --- Mock HTTP Server for Demonstration ---

// requestCounter for the mock server to change behavior
var requestCountSuccess int32
var requestCountRetryable int32
var requestCountTimeout int32

func mockServerHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[MockServer] Received request for %s", r.URL.Path)
	switch r.URL.Path {
	case "/success":
		atomic.AddInt32(&requestCountSuccess, 1)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Hello, client! This is a successful response.")
	case "/retry-once": // Fails once with 503, then succeeds
		count := atomic.AddInt32(&requestCountRetryable, 1)
		if count <= 1 {
			log.Printf("[MockServer] /retry-once: Sending 503 (attempt %d)", count)
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, "Service temporarily unavailable. Please try again.")
		} else {
			log.Printf("[MockServer] /retry-once: Sending 200 (attempt %d)", count)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "Service is back online!")
		}
	case "/fail-always": // Always fails with 500
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error. This will always fail.")
	case "/client-error": // Non-retryable client error
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Bad Request. Your input is invalid.")
	case "/slow-response": // Simulates a slow response that might trigger client timeout
		count := atomic.AddInt32(&requestCountTimeout, 1)
		log.Printf("[MockServer] /slow-response: attempt %d. Sleeping for 3s...", count)
		time.Sleep(3 * time.Second) // Simulate work
		// Check if client context was cancelled during our sleep
		select {
		case <-r.Context().Done():
			log.Printf("[MockServer] /slow-response: Client cancelled request while server was processing.")
			// Don't write anything if client is gone, or http.ErrHandlerTimeout if using http.TimeoutHandler
			return
		default:
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Finally! Here is your slow response.")
	case "/post-echo":
		if r.Method != http.MethodPost {
			http.Error(w, "POST only", http.StatusMethodNotAllowed)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read body", http.StatusInternalServerError)
			return
		}
		log.Printf("[MockServer] /post-echo: Received body: %s", string(body))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"received": "%s"}`, string(body))

	default:
		http.NotFound(w, r)
	}
}

// --- Main Demonstration ---

func main() {
	// Seed random number generator for jitter
	rand.Seed(time.Now().UnixNano())
	log.SetOutput(os.Stdout) // Ensure logs go to stdout

	// Start the mock server
	mockServer := httptest.NewServer(http.HandlerFunc(mockServerHandler))
	defer mockServer.Close()
	log.Printf("Mock server started at: %s", mockServer.URL)

	// Create our custom HTTP client
	httpClient := createCustomHTTPClient()

	// --- Demo 1: Successful request ---
	log.Println("\n--- Demo 1: Successful Request ---")
	reqSuccess, _ := http.NewRequest("GET", mockServer.URL+"/success", nil)
	ctxSuccess, cancelSuccess := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelSuccess()
	resp, err := DoRequestWithRetries(ctxSuccess, httpClient, reqSuccess, DefaultRetryOptions)
	if err != nil {
		log.Printf("Error on /success: %v", err)
	} else {
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("/success Response Status: %s, Body: %s", resp.Status, string(bodyBytes))
		resp.Body.Close()
	}

	// --- Demo 2: Request that retries once and then succeeds ---
	log.Println("\n--- Demo 2: Request Retries Once then Succeeds ---")
	atomic.StoreInt32(&requestCountRetryable, 0) // Reset counter for this demo
	reqRetryOnce, _ := http.NewRequest("GET", mockServer.URL+"/retry-once", nil)
	ctxRetryOnce, cancelRetryOnce := context.WithTimeout(context.Background(), 10*time.Second) // Longer timeout for retries
	defer cancelRetryOnce()
	resp, err = DoRequestWithRetries(ctxRetryOnce, httpClient, reqRetryOnce, DefaultRetryOptions)
	if err != nil {
		log.Printf("Error on /retry-once: %v", err)
	} else {
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("/retry-once Response Status: %s, Body: %s", resp.Status, string(bodyBytes))
		resp.Body.Close()
	}

	// --- Demo 3: Request that always fails (exhausts retries) ---
	log.Println("\n--- Demo 3: Request Always Fails (Exhausts Retries) ---")
	reqFailAlways, _ := http.NewRequest("GET", mockServer.URL+"/fail-always", nil)
	ctxFailAlways, cancelFailAlways := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFailAlways()
	resp, err = DoRequestWithRetries(ctxFailAlways, httpClient, reqFailAlways, DefaultRetryOptions)
	if err != nil {
		log.Printf("Expected error on /fail-always: %v", err)
		if resp != nil { // Should be nil, but good practice
			resp.Body.Close()
		}
	} else {
		log.Printf("/fail-always unexpected success: %s", resp.Status)
		resp.Body.Close()
	}

	// --- Demo 4: Non-retryable client error ---
	log.Println("\n--- Demo 4: Non-Retryable Client Error ---")
	reqClientError, _ := http.NewRequest("GET", mockServer.URL+"/client-error", nil)
	ctxClientError, cancelClientError := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelClientError()
	resp, err = DoRequestWithRetries(ctxClientError, httpClient, reqClientError, DefaultRetryOptions)
	if err != nil {
		log.Printf("Expected error on /client-error: %v", err)
		if resp != nil {
			resp.Body.Close()
		}
	} else {
		log.Printf("/client-error unexpected success: %s", resp.Status)
		resp.Body.Close()
	}

	// --- Demo 5: Context Cancellation during backoff ---
	log.Println("\n--- Demo 5: Context Cancellation During Backoff ---")
	atomic.StoreInt32(&requestCountRetryable, 0)                              // Reset counter
	reqCancel, _ := http.NewRequest("GET", mockServer.URL+"/retry-once", nil) // Will try to retry
	// Short-lived context that will be cancelled quickly
	ctxCancel, cancelCtxFunc := context.WithTimeout(context.Background(), DefaultRetryOptions.InitialBackoff/2) // Cancel before first retry sleep finishes
	defer cancelCtxFunc()
	// For this specific test, make the initial backoff slightly longer so cancellation is more likely to hit during sleep
	customOptsCancel := DefaultRetryOptions
	customOptsCancel.InitialBackoff = 300 * time.Millisecond
	customOptsCancel.MaxRetries = 1

	log.Printf("Starting request that will be cancelled. Context timeout: %v, Initial Backoff for retry: %v", DefaultRetryOptions.InitialBackoff/2, customOptsCancel.InitialBackoff)

	resp, err = DoRequestWithRetries(ctxCancel, httpClient, reqCancel, customOptsCancel)
	if err != nil {
		log.Printf("Expected error on /retry-once due to cancellation: %v", err)
		if !errors.Is(err, context.DeadlineExceeded) && !errors.Is(err, context.Canceled) {
			log.Printf("  WARNING: Error was not context.DeadlineExceeded or context.Canceled: %T", err)
		}
		if resp != nil {
			resp.Body.Close()
		}
	} else {
		log.Printf("/retry-once with cancellation, unexpected success: %s", resp.Status)
		resp.Body.Close()
	}
	// Wait a moment to ensure cancellation message from server (if any) is printed
	time.Sleep(100 * time.Millisecond)

	// --- Demo 6: Client-side timeout with slow server ---
	log.Println("\n--- Demo 6: Client-Side Timeout with Slow Server ---")
	// Custom client with a very short timeout for this specific test
	shortTimeoutClient := &http.Client{
		Transport: httpClient.Transport, // Reuse transport settings
		Timeout:   1 * time.Second,      // Very short overall timeout for client.Do()
	}
	reqSlow, _ := http.NewRequest("GET", mockServer.URL+"/slow-response", nil)
	// Context for the retry wrapper, can be longer than client.Timeout
	ctxSlowOp, cancelSlowOp := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelSlowOp()
	resp, err = DoRequestWithRetries(ctxSlowOp, shortTimeoutClient, reqSlow, DefaultRetryOptions)
	if err != nil {
		log.Printf("Expected error on /slow-response due to client timeout: %v", err)
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			log.Println("  Verified: Error is a net.Error Timeout.")
		} else {
			log.Printf("  WARNING: Error was not a net.Error Timeout: %T: %v", err, err)
		}
		if resp != nil {
			resp.Body.Close()
		}
	} else {
		log.Printf("/slow-response with client timeout, unexpected success: %s", resp.Status)
		resp.Body.Close()
	}
	// Give server time to log its cancellation if client timed out
	time.Sleep(3 * time.Second)

	// --- Demo 7: Concurrent Requests with Semaphore ---
	log.Println("\n--- Demo 7: Concurrent Requests with Semaphore ---")
	var wg sync.WaitGroup
	numConcurrentRequests := 5
	maxConcurrentLimit := 2 // Semaphore limit
	semaphore := make(chan struct{}, maxConcurrentLimit)

	urlsToFetch := []string{
		mockServer.URL + "/success?id=1",
		mockServer.URL + "/retry-once?id=2", // this will cause one retry
		mockServer.URL + "/success?id=3",
		mockServer.URL + "/slow-response?id=4", // this might timeout depending on context for this sub-task
		mockServer.URL + "/success?id=5",
	}
	atomic.StoreInt32(&requestCountRetryable, 0) // Reset retry counter
	atomic.StoreInt32(&requestCountTimeout, 0)   // Reset timeout counter

	opCtx, opCancel := context.WithTimeout(context.Background(), 20*time.Second) // Overall context for all concurrent ops
	defer opCancel()

	for i, urlStr := range urlsToFetch {
		wg.Add(1)
		go func(id int, u string, opNum int) {
			defer wg.Done()

			log.Printf("[Concurrent-%d] Waiting to acquire semaphore for %s", opNum, u)
			semaphore <- struct{}{} // Acquire slot
			log.Printf("[Concurrent-%d] Acquired semaphore for %s. Making request...", opNum, u)
			defer func() {
				log.Printf("[Concurrent-%d] Releasing semaphore for %s", opNum, u)
				<-semaphore // Release slot
			}()

			// Create a new context for this specific goroutine's operation, derived from the parent opCtx.
			// This allows individual requests to have shorter effective timeouts if needed, or just to ensure cancellation propagates.
			// For /slow-response, if its individual timeout is shorter than its processing time, it will timeout.
			var reqCtx context.Context
			var reqCancel context.CancelFunc
			if u == mockServer.URL+"/slow-response?id=4" {
				// Give slow response a shorter timeout to see it hit.
				reqCtx, reqCancel = context.WithTimeout(opCtx, 2*time.Second)
			} else {
				reqCtx, reqCancel = context.WithTimeout(opCtx, 10*time.Second) // Default for others
			}
			defer reqCancel()

			// Check parent context before even starting the request with retries
			select {
			case <-opCtx.Done():
				log.Printf("[Concurrent-%d] Overall operation cancelled before starting request for %s: %v", opNum, u, opCtx.Err())
				return
			default:
			}

			postData := []byte(fmt.Sprintf(`{"message": "hello from concurrent %d"}`, opNum))
			var req *http.Request
			var errReq error

			// Alternate between GET and POST for some requests to test body handling
			if opNum%2 == 0 { // Make even numbered operations POST
				req, errReq = http.NewRequestWithContext(reqCtx, "POST", mockServer.URL+"/post-echo", bytes.NewReader(postData))
				if errReq == nil {
					req.Header.Set("Content-Type", "application/json")
				}
			} else { // Odd ones are GET
				req, errReq = http.NewRequestWithContext(reqCtx, "GET", u, nil)
			}

			if errReq != nil {
				log.Printf("[Concurrent-%d] Failed to create request for %s: %v", opNum, u, errReq)
				return
			}

			log.Printf("[Concurrent-%d] Fetching %s", opNum, req.URL.String())
			resp, err := DoRequestWithRetries(reqCtx, httpClient, req, DefaultRetryOptions) // Use the derived reqCtx
			if err != nil {
				log.Printf("[Concurrent-%d] Error fetching %s: %v", opNum, req.URL.String(), err)
				if resp != nil { // Should be nil on error, but defensive
					io.Copy(io.Discard, resp.Body)
					resp.Body.Close()
				}
				return
			}
			defer resp.Body.Close()

			bodyBytes, _ := io.ReadAll(resp.Body)
			log.Printf("[Concurrent-%d] Successfully fetched %s, Status: %s, Body: %s", opNum, req.URL.String(), resp.Status, string(bodyBytes))

		}(i, urlStr, i+1) // Pass i as opNum for logging
		// Stagger goroutine creation slightly to make semaphore acquisition more obvious in logs
		if i < numConcurrentRequests-1 { // Don't sleep after the last one
			time.Sleep(50 * time.Millisecond)
		}
	}

	wg.Wait()
	log.Println("\n--- All concurrent requests completed ---")

	// --- Demo 8: POST request with retryable body ---
	log.Println("\n--- Demo 8: POST Request with Retryable Body ---")
	atomic.StoreInt32(&requestCountRetryable, 0) // Reset counter for this demo
	postBody := []byte(`{"key": "value", "attempt": 1}`)
	// bytes.NewReader is seekable, so http.NewRequest can set GetBody, making it retryable.
	reqPost, _ := http.NewRequest("POST", mockServer.URL+"/retry-once", bytes.NewReader(postBody)) // using /retry-once to force a retry
	reqPost.Header.Set("Content-Type", "application/json")

	ctxPost, cancelPost := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelPost()
	resp, err = DoRequestWithRetries(ctxPost, httpClient, reqPost, DefaultRetryOptions)
	if err != nil {
		log.Printf("Error on POST /retry-once: %v", err)
	} else {
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("POST /retry-once Response Status: %s, Body: %s", resp.Status, string(bodyBytes))
		resp.Body.Close()
	}

	log.Println("\n--- All demos finished ---")
}
