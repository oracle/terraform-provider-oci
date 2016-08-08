package main

// ResourceSync synchronizes a ResourceData instance and a BareMetal entity.
type ResourceSync interface {
	Create() (BareMetalResource, error)
	Get() (BareMetalResource, error)
	Update() (BareMetalResource, error)
	SetData(res BareMetalResource)
	Delete() error
}
