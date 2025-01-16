package entities

// User represents a user entity.
type User struct {
	Resource
}

// Project represents a project entity.
type Project struct {
	Resource
}

// ResourceType represents the type of resource. Either a user or a project.
type ResourceType string

const (
	ResourceTypeUser    = "user"
	ResourceTypeProject = "project"
)

// Resource represents the general structure of a resource, shared between users and projects.
type Resource struct {
	GID          string       `json:"gid"`
	Name         string       `json:"name"`
	ResourceType ResourceType `json:"resource_type"`
}
