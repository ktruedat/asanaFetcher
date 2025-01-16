package service

import (
	"fmt"

	"github.com/ktruedat/recoAssignment/internal/domain/entities"
	"github.com/pkg/errors"
)

var ErrInvalidResourceType = errors.New("invalid resource type: the only valid options are 'user', 'project'")

const (
	// projectsAPIURLFormattedString - API URL extension for fetching the projects.
	projectsAPIURLFormattedString = "/projects"
	// userAPIURLFormattedString - API URL extension for fetching the users from a workspace. The %s format arg is a
	// placeholder for the WorkspaceGID.
	userAPIURLFormattedString = "/workspaces/%s/users"
)

func (s *service[T]) constructAPIUrlFromResource() (string, error) {
	switch s.resourceType {
	case entities.ResourceTypeUser:
		return fmt.Sprintf(userAPIURLFormattedString, s.cfg.WorkspaceGID), nil
	case entities.ResourceTypeProject:
		return projectsAPIURLFormattedString, nil
	default:
		return "", ErrInvalidResourceType
	}
}
