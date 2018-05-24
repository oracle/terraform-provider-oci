// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package crud

import (
	"sync"
	"time"
)

// Common interfaces

// Gets the current BareMetal Resource
type ResourceFetcher interface {
	// Get should update the s.Resource, and is used by crud.ReadResource() to populate s.D
	// Get() may expect s.D.Id() to be set, but not s.Resource, or anything else in s.D
	Get() error
}

// ResourceVoider may set its ResourceData state to empty.
type ResourceVoider interface {
	// VoidState sets ResourceData ID to "". To be called when
	// the resource is gone.
	VoidState()
}

// ResourceDataWriter populates ResourceData state from the current
// BareMetal resource
type ResourceDataWriter interface {
	ResourceVoider
	// SetData populates ResourceData state from current
	// BareMetal resource
	SetData()
}

// CRUD standard interfaces

// ResourceCreator may create a BareMetal resource and populate into
// ResourceData state by using CreateResource()
type ResourceCreator interface {
	ResourceDataWriter
	// ID identifies the resource, or a work request to create the resource.
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
	ResourceVoider
	// ID identifies the resource, or a work request to create the resource.
	ID() string
	Delete() error
}

// Some resources in the oracle API are removed asynchronously, so even
// after they claim to be gone, other dependencies haven't been notified
// of that fact. This facility allows us to add an artificial delay for
// resources that need a little time to let the oracle API backend catch
// up with reality.
type ExtraWaitPostCreateDelete interface {
	ExtraWaitPostCreateDelete() time.Duration
}

type StatefulResource interface {
	ResourceReader
	State() string
	setState(StatefulResource) error
}

type StatefullyCreatedResource interface {
	StatefulResource
	CreatedPending() []string
	CreatedTarget() []string
}

type StatefullyUpdatedResource interface {
	StatefulResource
	UpdatedPending() []string
	UpdatedTarget() []string
}

type StatefullyDeletedResource interface {
	StatefulResource
	DeletedPending() []string
	DeletedTarget() []string
}

// This provides a mechanism for synchronizing CRUD operations from different resources
// that may concurrently modify the same resource. Implementing these interfaces will
// cause the Create/Update/Delete operations to wait on the lock before starting those
// operations.
type SynchronizedResource interface {
	GetMutex() *sync.Mutex
}
