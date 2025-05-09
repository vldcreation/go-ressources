package model

import (
	"time"
)

// Video struct matches the database table schema.
// Using pointers for nullable fields (e.g., *time.Time, *int64) allows them
// to be nil if the database value is NULL, and `omitempty` works well with JSON.
type Video struct {
	ID           int64      `json:"id"`
	Name         string     `json:"name"`
	URL          string     `json:"url"`
	ThumbnailURL string     `json:"thumbnail_url"`
	MerchantID   int64      `json:"merchant_id"`
	Status       string     `json:"status"` // e.g., "active", "pending"
	Caption      string     `json:"caption"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"` // Pointer for NULLable timestamp
	CreatedBy    int64      `json:"created_by"`
	UpdatedBy    int64      `json:"updated_by"`
	DeletedBy    *int64     `json:"deleted_by,omitempty"` // Pointer for NULLable int
}

// PaginatedVideoResponse is the structure for the API response.
type PaginatedVideoResponse struct {
	Videos     []Video `json:"videos"`
	Pagination struct {
		CurrentPage int    `json:"current_page"`
		PerPage     int    `json:"per_page"`
		TotalVideos *int   `json:"total_videos,omitempty"` // Total available videos for this seed, sent on page 1
		SessionSeed string `json:"session_seed"`           // Seed for the current random sequence
	} `json:"pagination"`
}

// Nullable types from database/sql can also be used if you prefer explicit handling
// e.g. DeletedAt sql.NullTime
//      DeletedBy sql.NullInt64
// This requires custom MarshalJSON/UnmarshalJSON or using a library that handles them.
// Pointers are often simpler for basic JSON marshalling.
