package main

import "github.com/MustWin/baremtlclient"

type BareMetalClient interface {
	CreateUser(name, description string, options ...baremtlsdk.Options) (*baremtlsdk.Resource, error)
	GetUser(userID string) (*baremtlsdk.Resource, error)
	UpdateUser(userID, userDescription string, opts ...baremtlsdk.Options) (*baremtlsdk.Resource, error)
}
