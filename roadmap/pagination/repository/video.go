package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log" // For logging errors or warnings

	"github.com/vldcreation/go-ressources/roadmap/pagination/model"
)

// VideoRepository handles database operations for videos.
type VideoRepository struct {
	DB *sql.DB
}

// NewVideoRepository creates a new VideoRepository.
func NewVideoRepository(db *sql.DB) *VideoRepository {
	return &VideoRepository{DB: db}
}

// GetRandomVideos fetches videos in a seeded random order with pagination.
// It uses a transaction to ensure `SET LOCAL SEED` applies correctly.
func (r *VideoRepository) GetRandomVideos(ctx context.Context, dbSeed float64, perPage, offset int) ([]model.Video, error) {
	videos := []model.Video{}

	// Start a transaction. SET LOCAL SEED is transaction-scoped.
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	// Defer a rollback in case anything goes wrong.
	// The rollback will be ignored if Commit() is called successfully.
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Re-panic after rollback
		} else if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("Error rolling back transaction: %v (original error: %v)", rbErr, err)
			}
		}
	}()

	// Set the seed for the random number generator for the current transaction.
	// `SET LOCAL SEED` is preferred as it's cleaner than `SELECT setseed()`.
	_, err = tx.ExecContext(ctx, "SELECT SETSEED($1)", dbSeed)
	if err != nil {
		// err is already set, defer will handle rollback
		return nil, fmt.Errorf("failed to set database seed: %w", err)
	}

	// Query to fetch videos.
	// Ensure 'status' value (e.g., 'active') matches your data.
	// The 'deleted_at IS NULL' ensures soft-deleted videos are excluded.
	query := `
        SELECT id, "name", url, thumbnail_url, merchant_id, status, caption,
               created_at, updated_at, deleted_at, created_by, updated_by, deleted_by
        FROM video
        WHERE deleted_at IS NULL AND status = 'ACTIVE'
        ORDER BY random()
        LIMIT $1 OFFSET $2;
    `
	rows, err := tx.QueryContext(ctx, query, perPage, offset)
	if err != nil {
		// err is already set, defer will handle rollback
		return nil, fmt.Errorf("failed to query videos: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var v model.Video
		// Scan into the Video struct fields.
		// Pointers in the struct (*time.Time, *int64) will correctly handle NULLs
		// by remaining nil if the DB value is NULL.
		if scanErr := rows.Scan(
			&v.ID, &v.Name, &v.URL, &v.ThumbnailURL, &v.MerchantID,
			&v.Status, &v.Caption, &v.CreatedAt, &v.UpdatedAt,
			&v.DeletedAt, // Scans to *time.Time
			&v.CreatedBy, &v.UpdatedBy,
			&v.DeletedBy, // Scans to *int64
		); scanErr != nil {
			err = fmt.Errorf("failed to scan video row: %w", scanErr) // Set err for defer
			return nil, err
		}
		videos = append(videos, v)
	}
	// Check for errors encountered during iteration.
	if err = rows.Err(); err != nil {
		// err is already set, defer will handle rollback
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	// If all operations were successful, commit the transaction.
	if err = tx.Commit(); err != nil {
		// err is already set, defer will handle rollback (though commit failed, so it already did)
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return videos, nil // err is nil here
}

// GetTotalActiveVideos counts all active, non-deleted videos.
// This is used to inform the client about the total number of items in the current random sequence.
func (r *VideoRepository) GetTotalActiveVideos(ctx context.Context) (int, error) {
	var count int
	// Ensure 'status' value matches your data.
	query := `SELECT COUNT(*) FROM public.video WHERE deleted_at IS NULL AND status = 'active'`
	err := r.DB.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count active videos: %w", err)
	}
	return count, nil
}
