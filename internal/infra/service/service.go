package service

import (
	"github.com/ktruedat/recoAssignment/config"
	"github.com/ktruedat/recoAssignment/internal/domain/entities"
	"github.com/ktruedat/recoAssignment/internal/infra/http"
	"github.com/pkg/errors"
)

// ResourceGetService is an interface for fetching resources from the API.
type ResourceGetService[T ResourceConstraint] interface {
	// Get fetches the resources from the API.
	Get() ([]T, error)
}

// ResourceConstraint is an interface for constraining the resource type.
type ResourceConstraint interface {
	entities.User | entities.Project
}

type service[T ResourceConstraint] struct {
	cl           http.Client
	resourceType entities.ResourceType
	cfg          *config.Config
}

// NewUserGetService creates a new user get service.
func NewUserGetService(client http.Client, cfg *config.Config) ResourceGetService[entities.User] {
	return &service[entities.User]{
		cl:           client,
		resourceType: entities.ResourceTypeUser,
		cfg:          cfg,
	}
}

// NewProjectGetService creates a new project get service.
func NewProjectGetService(client http.Client, cfg *config.Config) ResourceGetService[entities.Project] {
	return &service[entities.Project]{
		cl:           client,
		resourceType: entities.ResourceTypeProject,
		cfg:          cfg,
	}
}

// Get fetches the resources from the API.
func (s *service[T]) Get() ([]T, error) {
	url, err := s.constructAPIUrlFromResource()
	if err != nil {
		return nil, errors.Wrap(err, "failed to determine url based on resource type")
	}

	var ents http.Data[T]
	if err := s.cl.Get(url, &ents); err != nil {
		return nil, errors.Wrap(err, "failed to fetch the resources")
	}

	return ents.Data, nil
}
