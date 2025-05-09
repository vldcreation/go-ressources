package util

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math" // Required for math.Mod and Abs

	"github.com/google/uuid"
)

// GenerateNewSessionSeed creates a new unique string seed (UUID).
func GenerateNewSessionSeed() string {
	return uuid.NewString()
}

// DeriveDBSeedFromString converts a string seed (like a UUID) to a float64
// in the range [-1.0, 1.0] for PostgreSQL's setseed().
// This derivation must be deterministic: same input string yields same float.
func DeriveDBSeedFromString(sessionSeed string) (float64, error) {
	if sessionSeed == "" {
		return 0, fmt.Errorf("session seed cannot be empty for derivation")
	}

	// Create a SHA256 hash of the session seed.
	h := sha256.New()
	h.Write([]byte(sessionSeed))
	hashBytes := h.Sum(nil)

	// Use the first 8 bytes of the hash to form a uint64.
	// This provides a large number based on the seed.
	if len(hashBytes) < 8 {
		// This should not happen with SHA256, which produces 32 bytes.
		return 0, fmt.Errorf("hash too short to derive seed value, got %d bytes", len(hashBytes))
	}
	seedAsUint64 := binary.BigEndian.Uint64(hashBytes[:8])

	// Convert uint64 to int64 to get a signed value.
	seedAsInt64 := int64(seedAsUint64)

	// Normalize the int64 value to a float64 between -1.0 and 1.0.
	// One way: use modulo arithmetic on a large prime, then scale.
	// Or, more simply, scale based on int64 range.
	// MaxInt64 is approx 9.22e18.
	// A simple normalization:
	// Take the int64, divide by MaxInt64. This gives a value in [-1, 1].
	// Using math.Mod can also help distribute values if int64 is too large or small.
	// For setseed(), the distribution quality is important.
	// Let's use a simpler approach: treat the int64 as a point on a large number line
	// and scale it to the -1 to 1 range.
	// Example: (value % N) / (N/2.0) - 1.0 might not be ideal due to modulo bias.
	// A direct scaling is often preferred:
	// dbSeed := float64(seedAsInt64) / float64(math.MaxInt64)

	// Alternative approach for better distribution for setseed:
	// Convert the int64 to a float64, then use math.Mod to bring it into a manageable range,
	// then scale to [-1.0, 1.0].
	// For example, map to [0, 2*PI), then use sin() or cos().
	// Or simpler:
	// Take the absolute value, mod by a large number (e.g., 2,000,000,000),
	// then scale this to [0, 1] and then map to [-1, 1].
	// The sign of original int64 can determine if it's positive or negative part.

	// Let's use a direct scaling of the int64.
	// If seedAsInt64 is 0, dbSeed will be 0.
	// If seedAsInt64 is math.MaxInt64, dbSeed will be 1.0.
	// If seedAsInt64 is math.MinInt64, dbSeed will be -1.0.
	var dbSeed float64
	if seedAsInt64 >= 0 {
		dbSeed = float64(seedAsInt64) / float64(math.MaxInt64)
	} else {
		// For negative numbers, MinInt64 is -(MaxInt64 + 1).
		// So, dividing by -MinInt64 (which is MaxInt64+1) can map it to [-1, 0).
		dbSeed = float64(seedAsInt64) / float64(math.MaxInt64+1) // approximately -math.Abs(float64(math.MinInt64))
	}

	// Ensure it's strictly within [-1.0, 1.0] due to potential float inaccuracies.
	if dbSeed > 1.0 {
		dbSeed = 1.0
	} else if dbSeed < -1.0 {
		dbSeed = -1.0
	}

	return dbSeed, nil
}
