package main

import "github.com/MustWin/baremtlclient"

type BareMetalResource interface {
	GetId() string
	GetState() string
}

type BareMetalResourceAdapter struct {
	*baremtlsdk.Resource
}

func (a *BareMetalResourceAdapter) GetId() string    { return a.ID }
func (a *BareMetalResourceAdapter) GetState() string { return a.State }

type BareMetalPolicyAdapter struct {
	*baremtlsdk.Policy
}

func (a *BareMetalPolicyAdapter) GetId() string    { return a.ID }
func (a *BareMetalPolicyAdapter) GetState() string { return a.State }
