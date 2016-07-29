package main

import "github.com/MustWin/baremtlclient"

type BareMetalClient interface {
	CreateUser(name, description string, options ...baremtlclient.Options) (*baremtlclient.Resource, error)
}
