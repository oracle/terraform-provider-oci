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
	StatefulCreation() bool
	CreatedPending() []string
	CreatedTarget() []string
	State() string
	Create() error
	Update() error
	Delete() error
}

type IdentitySync struct{}

func (s *IdentitySync) StatefulCreation() bool {
	return true
}

func (s *IdentitySync) CreatedPending() []string {
	return []string{baremetal.ResourceCreating}
}

func (s *IdentitySync) CreatedTarget() []string {
	return []string{baremetal.ResourceCreated}
}

type CoreSync struct{}

func (s *CoreSync) State() string {
	return ""
}

func (s *CoreSync) StatefulCreation() bool {
	return false
}

func (s *CoreSync) CreatedPending() []string {
	return nil
}

func (s *CoreSync) CreatedTarget() []string {
	return nil
}
