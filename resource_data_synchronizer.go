package main

import "github.com/MustWin/baremtlclient"

// ResourceSync synchronizes a ResourceData instance and a BareMetal entity.
type ResourceSync interface {
	Create() (*baremtlsdk.Resource, error)
	Get() (*baremtlsdk.Resource, error)
	Update() (*baremtlsdk.Resource, error)
	SetData(res *baremtlsdk.Resource)
	Delete() error
}
