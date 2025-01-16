package http

import (
	"github.com/ktruedat/recoAssignment/internal/domain/entities"
)

type Client interface {
	Get(url string, response any) error
	GetRaw(url string) ([]byte, error)
}

type Data[T entities.User | entities.Project] struct {
	Data []T `json:"data"`
}

type Token struct{}

type Limiter interface {
	Tries() <-chan Token
}
