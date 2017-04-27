// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package crud

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
)

// Gets the current BareMetal Resource
type ResourceFetcher interface {
	// Get should update the s.Resource, and is used by crud.ReadResource() to populate s.D
	// Get() may expect s.D.Id() to be set, but not s.Resource, or anything else in s.D
	Get() error
}

// ResourceDataWriter populates ResourceData based on current BareMetal Resource
type ResourceDataWriter interface {
	ResourceVoider
	SetData()
}

// ResourceCreator creates a BareMetal resource based on ResourceData
type ResourceCreator interface {
	ResourceDataWriter
	// ID identifies the resource, or a work request to create the resource.
	ID() string
	Create() error
}

type ResourceVoider interface {
	VoidState() // Call this when the resource is gone
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
	Delete() error
}

// Some resources in the oracle API are removed asynchronously, so even
// after they claim to be gone, other dependencies haven't been notified
// of that fact. This facility allows us to add an artificial delay for
// resources that need a little time to let the oracle API backend catch
// up with reality.
type ExtraWaitPostDelete interface {
	ExtraWaitPostDelete() time.Duration
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
