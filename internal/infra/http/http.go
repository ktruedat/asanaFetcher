package http

import (
	"github.com/ktruedat/recoAssignment/internal/domain/entities"
)

// Client is an interface for making HTTP requests.
type Client interface {
	// Get makes a GET request to the given URL and unmarshalls the response into the given response object.
	Get(url string, response any) error
	// GetRaw makes a GET request to the given URL and returns the raw response body.
	GetRaw(url string) ([]byte, error)
}

// Data represents a generic data structures used for unmarshalling JSON responses.
type Data[T entities.User | entities.Project] struct {
	Data []T `json:"data"`
}

// Token represents a token for rate limiting.
type Token struct{}

// Limiter represents a rate limiter.
type Limiter interface {
	Tries() <-chan Token
}
