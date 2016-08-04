package main

import "github.com/MustWin/baremtlclient"

type CreateResourceFn func(string, string, ...baremtlsdk.Options) (*baremtlsdk.Resource, error)
type GetResourceFn func(string) (*baremtlsdk.Resource, error)
type UpdateResourceFn func(string, string, ...baremtlsdk.Options) (*baremtlsdk.Resource, error)
type DeleteResourceFn func(string, ...baremtlsdk.Options) error

type BareMetalClient interface {
	CreateUser(name, description string, options ...baremtlsdk.Options) (*baremtlsdk.Resource, error)
	GetUser(userID string) (*baremtlsdk.Resource, error)
	UpdateUser(userID, userDescription string, opts ...baremtlsdk.Options) (*baremtlsdk.Resource, error)
	DeleteUser(userID string, opts ...baremtlsdk.Options) error

	CreateGroup(name, description string, options ...baremtlsdk.Options) (*baremtlsdk.Resource, error)
	GetGroup(userID string) (*baremtlsdk.Resource, error)
	UpdateGroup(userID, userDescription string, opts ...baremtlsdk.Options) (*baremtlsdk.Resource, error)
	DeleteGroup(userID string, opts ...baremtlsdk.Options) error
}
