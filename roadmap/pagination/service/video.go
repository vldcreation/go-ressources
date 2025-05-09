package service

import (
	"context"
	"fmt"
	"log" // For logging warnings

	"github.com/vldcreation/go-ressources/roadmap/pagination/model"
	"github.com/vldcreation/go-ressources/roadmap/pagination/repository"
	"github.com/vldcreation/go-ressources/roadmap/pagination/util"
)

// VideoService encapsulates business logic for videos.
type VideoService struct {
	repo *repository.VideoRepository
}

// NewVideoService creates a new VideoService.
func NewVideoService(repo *repository.VideoRepository) *VideoService {
	return &VideoService{repo: repo}
}

// GetVideos retrieves a paginated list of videos, ensuring consistent randomness per session.
func (s *VideoService) GetVideos(ctx context.Context, page, perPage int, clientSessionSeed string) (*model.PaginatedVideoResponse, error) {
	// Validate and set default pagination parameters.
	if page <= 0 {
		page = 1
	}
	if perPage <= 0 {
		perPage = 10 // Default items per page
	}
	// Optional: Set a maximum limit for perPage to prevent abuse.
	if perPage > 100 {
		perPage = 100
	}

	var currentSessionSeedForClient string
	var dbSeedForQuery float64
	var err error

	// Determine the session seed to use.
	// If client doesn't provide a seed, or it's page 1 and behavior is "new list on refresh",
	// generate a new one.
	if clientSessionSeed == "" {
		currentSessionSeedForClient = util.GenerateNewSessionSeed()
	} else {
		currentSessionSeedForClient = clientSessionSeed
	}

	// Derive the numeric database seed from the string session seed.
	dbSeedForQuery, err = util.DeriveDBSeedFromString(currentSessionSeedForClient)
	if err != nil {
		return nil, fmt.Errorf("failed to derive database seed from session seed '%s': %w", currentSessionSeedForClient, err)
	}

	// Calculate offset for pagination.
	offset := (page - 1) * perPage

	// Fetch videos from the repository.
	videos, err := s.repo.GetRandomVideos(ctx, dbSeedForQuery, perPage, offset)
	if err != nil {
		// Detailed error is already logged by repository or should be here.
		return nil, fmt.Errorf("service failed to get random videos: %w", err)
	}

	// Prepare the response.
	response := &model.PaginatedVideoResponse{}
	response.Videos = videos
	response.Pagination.CurrentPage = page
	response.Pagination.PerPage = perPage
	response.Pagination.SessionSeed = currentSessionSeedForClient // Send the seed back to the client.

	// If it's the first page of a session (identified by client not sending a seed,
	// or explicitly page 1), fetch the total number of available videos for this seed sequence.
	// This helps the client know the maximum number of pages.
	if page == 1 { // Or more robustly: if clientSessionSeed == ""
		total, countErr := s.repo.GetTotalActiveVideos(ctx)
		if countErr != nil {
			// Log the error but don't fail the whole request.
			// The client can still paginate but won't know the total upfront.
			log.Printf("Warning: failed to get total active videos for session %s: %v", currentSessionSeedForClient, countErr)
			// response.Pagination.TotalVideos could be nil or explicitly set to indicate an issue.
		} else {
			response.Pagination.TotalVideos = &total
		}
	}

	return response, nil
}
