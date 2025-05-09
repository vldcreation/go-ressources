package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/vldcreation/go-ressources/roadmap/pagination/service"
)

// VideoHandler handles HTTP requests related to videos.
type VideoHandler struct {
	videoService *service.VideoService
}

// NewVideoHandler creates a new VideoHandler.
func NewVideoHandler(vs *service.VideoService) *VideoHandler {
	return &VideoHandler{videoService: vs}
}

// GetVideos handles requests for the video feed.
// GET /api/videos?page=1&per_page=10[&session_seed=xxxx-xxxx-xxxx-xxxx]
func (h *VideoHandler) GetVideos(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters.
	pageStr := r.URL.Query().Get("page")
	perPageStr := r.URL.Query().Get("per_page")
	sessionSeed := r.URL.Query().Get("session_seed") // Client sends this for subsequent pages.

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1 // Default to page 1 if not specified or invalid.
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		perPage = 10 // Default to 10 items per page.
	}
	// Cap perPage to a reasonable maximum.
	if perPage > 50 { // Example cap
		perPage = 50
	}

	// Call the service layer to get the videos.
	paginatedResponse, serviceErr := h.videoService.GetVideos(r.Context(), page, perPage, sessionSeed)
	if serviceErr != nil {
		// Log the internal error for debugging.
		log.Printf("ERROR [VideoHandler]: Failed to get videos: %v", serviceErr)
		// Provide a generic error message to the client.
		http.Error(w, "An error occurred while fetching videos.", http.StatusInternalServerError)
		return
	}

	// Set content type and encode the response as JSON.
	w.Header().Set("Content-Type", "application/json")
	if encodeErr := json.NewEncoder(w).Encode(paginatedResponse); encodeErr != nil {
		log.Printf("ERROR [VideoHandler]: Failed to encode response: %v", encodeErr)
		// If encoding fails, it's likely too late to send a different HTTP error code,
		// but we log it. The client might receive a partial or malformed response.
		http.Error(w, "Failed to encode response.", http.StatusInternalServerError) // Attempt to send error
	}
}

// SetupRoutes configures the routes for video-related endpoints.
// router is an HTTP router, e.g., http.ServeMux, chi router, gin engine, etc.
func (h *VideoHandler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/videos", h.GetVideos) // GET /api/videos
}
