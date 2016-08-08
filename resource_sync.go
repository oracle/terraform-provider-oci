package main

// ResourceSync synchronizes a ResourceData instance and a BareMetal entity.
type ResourceSync interface {
	Id() string
	State() string
	Create() error
	Get() error
	Update() error
	SetData()
	Delete() error
}
