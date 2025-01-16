package entities

type User struct {
	Resource
}

type Project struct {
	Resource
}

type ResourceType string

const (
	ResourceTypeUser    = "user"
	ResourceTypeProject = "project"
)

type Resource struct {
	GID          string       `json:"gid"`
	Name         string       `json:"name"`
	ResourceType ResourceType `json:"resource_type"`
}
