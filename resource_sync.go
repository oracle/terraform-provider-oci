package main

import "github.com/MustWin/baremetal-sdk-go"

// Reads BareMetal entity
type ResourceReader interface {
	Get() error
	SetData()
}

// ResourceSync synchronizes a ResourceData instance and a BareMetal entity.
type ResourceSync interface {
	ResourceReader
	Id() string
	Create() error
	Update() error
	Delete() error
}

type StatefulResourceSync interface {
	ResourceSync
	CreatedPending() []string
	CreatedTarget() []string
	State() string
}

type IdentitySync struct{}

func (s *IdentitySync) CreatedPending() []string {
	return []string{baremetal.ResourceCreating}
}

func (s *IdentitySync) CreatedTarget() []string {
	return []string{baremetal.ResourceCreated}
}
