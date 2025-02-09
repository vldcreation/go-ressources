package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
)

// Article represents the structure of an article in the API response
type Article struct {
	Title       *string `json:"title"`
	NumComments int     `json:"num_comments"`
	StoryTitle  *string `json:"story_title"`
}

// APIResponse represents the structure of the API response
type APIResponse struct {
	Page       int       `json:"page"`
	PerPage    int       `json:"per_page"`
	Total      int       `json:"total"`
	TotalPages int       `json:"total_pages"`
	Data       []Article `json:"data"`
}

// topArticles fetches the top articles sorted by the number of comments
func topArticles(limit int32) []string {
	baseURL := "https://jsonmock.hackerrank.com/api/articles?page="
	var allArticles []Article

	// Fetch data from all pages
	page := 1
	for {
		url := baseURL + strconv.Itoa(page)
		response, err := http.Get(url)
		if err != nil {
			fmt.Println("Error fetching data:", err)
			return nil
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return nil
		}
		defer response.Body.Close()

		var apiResponse APIResponse
		err = json.Unmarshal(body, &apiResponse)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return nil
		}

		// Append articles from the current page
		allArticles = append(allArticles, apiResponse.Data...)

		// Break if we've fetched all pages
		if page >= apiResponse.TotalPages {
			break
		}
		page++
	}

	// Sort articles by number of comments (descending)
	// if the articles doesn't have a title and a story tittle then we will skip it
	sort.Slice(allArticles, func(i, j int) bool {
		if allArticles[i].NumComments == allArticles[j].NumComments {
			// If both articles have the same number of comments, sort by title
			if allArticles[i].Title != nil && allArticles[j].Title != nil {
				return *allArticles[i].Title > *allArticles[j].Title
			} else if allArticles[i].StoryTitle != nil && allArticles[j].StoryTitle != nil {
				return *allArticles[i].StoryTitle > *allArticles[j].StoryTitle
			} else {
				return false
			}
		}

		return allArticles[i].NumComments > allArticles[j].NumComments
	})

	// write allArticles to a file
	f, err := os.Create("allArticles.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return nil
	}

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	err = enc.Encode(allArticles)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return nil
	}

	// Close the file
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	// Extract the top `limit` titles
	var topTitles []string
	for i := 0; i < len(allArticles) && i < int(limit); i++ {
		if allArticles[i].Title != nil {
			topTitles = append(topTitles, *allArticles[i].Title)
		} else if allArticles[i].StoryTitle != nil {
			topTitles = append(topTitles, *allArticles[i].StoryTitle)
		}
	}

	return topTitles
}

func main() {
	// Test the function
	fmt.Println(topArticles(2)) // Example: Fetch top 2 articles
}
