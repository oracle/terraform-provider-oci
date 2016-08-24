package crud

import "github.com/MustWin/baremetal-sdk-go"

// Gets the current BareMetal Resource
type ResourceFetcher interface {
	Get() error
}

// ResourceDataWriter populates ResourceData based on current BareMetal Resource
type ResourceDataWriter interface {
	SetData()
}

// ResourceCreator creates a BareMetal resource based on ResourceData
type ResourceCreator interface {
	ResourceDataWriter
	ID() string
	Create() error
}

// ResourceReader get BareMetal Resource and updated ResourceData
type ResourceReader interface {
	ResourceFetcher
	ResourceDataWriter
}

// Updates a BareMetal entity to match ResourceData
type ResourceUpdater interface {
	ResourceDataWriter
	Update() error
}

// Deletes a BareMetal entity
type ResourceDeleter interface {
	Delete() error
}

type StatefulResource interface {
	ResourceReader
	State() string
}

type StatefullyCreatedResource interface {
	StatefulResource
	CreatedPending() []string
	CreatedTarget() []string
}

type StatefullyDeletedResource interface {
	StatefulResource
	DeletedPending() []string
	DeletedTarget() []string
}

type IdentitySync struct{}

func (s *IdentitySync) CreatedPending() []string {
	return []string{baremetal.ResourceCreating}
}

func (s *IdentitySync) CreatedTarget() []string {
	return []string{baremetal.ResourceCreated}
}
