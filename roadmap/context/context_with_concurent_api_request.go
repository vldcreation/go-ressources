package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func doConcurentApiRequest() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	urls := []string{
		"https://api.agify.io/?name=Viky",
		"https://api.agify.io/?name=Elum",
		"https://api.agify.io/?name=Foo",
	}

	results := make(chan string)

	for _, url := range urls {
		go fetchAPI(ctx, url, results)
	}

	for range urls {
		fmt.Println(<-results)
	}
}

func fetchAPI(ctx context.Context, url string, results chan<- string) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		results <- fmt.Sprintf("Error creating request for %s: %s", url, err.Error())
		return
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		results <- fmt.Sprintf("Error making request to %s: %s", url, err.Error())
		return
	}
	defer resp.Body.Close()

	out := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&out)
	if err != nil {
		results <- fmt.Sprintf("Error marshaling response from %s: %s", url, err.Error())
		return
	}

	outBytes, err := json.Marshal(out)
	if err != nil {
		results <- fmt.Sprintf("Error marshaling response from %s: %s", url, err.Error())
		return
	}
	if out, err := prettyPrint(outBytes); err != nil {
		results <- fmt.Sprintf("Error pretty printing response from %s: %s", url, err.Error())
		return
	} else {
		results <- fmt.Sprintf("Response from %s: %s", url, string(out))
		return
	}
}

func prettyPrint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}
