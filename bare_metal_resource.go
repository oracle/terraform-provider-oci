package main

import "github.com/MustWin/baremtlclient"

type BareMetalResource interface {
	Id() string
	Name() string
	Description() string
	CompartmentId() string
	State() string
	TimeModified() string
	TimeCreated() string
}

type BareMetalResourceAdapter struct {
	res *baremtlsdk.Resource
}

func (a *BareMetalResourceAdapter) Id() string            { return a.res.ID }
func (a *BareMetalResourceAdapter) Name() string          { return a.res.Name }
func (a *BareMetalResourceAdapter) Description() string   { return a.res.Description }
func (a *BareMetalResourceAdapter) CompartmentId() string { return a.res.CompartmentID }
func (a *BareMetalResourceAdapter) State() string         { return a.res.State }
func (a *BareMetalResourceAdapter) TimeModified() string  { return a.res.TimeModified.String() }
func (a *BareMetalResourceAdapter) TimeCreated() string   { return a.res.TimeCreated.String() }

type BareMetalPolicyAdapter struct {
	res *baremtlsdk.Policy
}

func (a *BareMetalPolicyAdapter) Id() string            { return a.res.ID }
func (a *BareMetalPolicyAdapter) Name() string          { return a.res.Name }
func (a *BareMetalPolicyAdapter) Description() string   { return a.res.Description }
func (a *BareMetalPolicyAdapter) CompartmentId() string { return a.res.CompartmentID }
func (a *BareMetalPolicyAdapter) State() string         { return a.res.State }
func (a *BareMetalPolicyAdapter) Statements() []string  { return a.res.Statements }
func (a *BareMetalPolicyAdapter) TimeModified() string  { return a.res.TimeModified.String() }
func (a *BareMetalPolicyAdapter) TimeCreated() string   { return a.res.TimeCreated.String() }
