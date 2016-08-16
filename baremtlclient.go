package main

import "github.com/MustWin/baremetal-sdk-go"

type CreateResourceFn func(string, string, ...baremtlsdk.Options) (*baremtlsdk.IdentityResource, error)
type GetResourceFn func(string) (*baremtlsdk.IdentityResource, error)
type UpdateResourceFn func(string, string, ...baremtlsdk.Options) (*baremtlsdk.IdentityResource, error)
type DeleteResourceFn func(string, ...baremtlsdk.Options) error

type BareMetalClient interface {
	CreateUser(name, description string, options ...baremtlsdk.Options) (*baremtlsdk.IdentityResource, error)
	GetUser(userID string) (*baremtlsdk.IdentityResource, error)
	UpdateUser(userID, userDescription string, opts ...baremtlsdk.Options) (*baremtlsdk.IdentityResource, error)
	DeleteUser(userID string, opts ...baremtlsdk.Options) error

	CreateGroup(name, description string, options ...baremtlsdk.Options) (*baremtlsdk.IdentityResource, error)
	GetGroup(userID string) (*baremtlsdk.IdentityResource, error)
	UpdateGroup(userID, userDescription string, opts ...baremtlsdk.Options) (*baremtlsdk.IdentityResource, error)
	DeleteGroup(userID string, opts ...baremtlsdk.Options) error

	CreatePolicy(name, description string, statements []string, opts ...baremtlsdk.Options) (*baremtlsdk.Policy, error)
	GetPolicy(id string) (*baremtlsdk.Policy, error)
	UpdatePolicy(id, description string, statements []string, opts ...baremtlsdk.Options) (*baremtlsdk.Policy, error)
	DeletePolicy(id string, opts ...baremtlsdk.Options) error

	CreateCompartment(name, description string, options ...baremtlsdk.Options) (*baremtlsdk.IdentityResource, error)
	GetCompartment(userID string) (*baremtlsdk.IdentityResource, error)
	UpdateCompartment(userID, userDescription string, opts ...baremtlsdk.Options) (*baremtlsdk.IdentityResource, error)

	ListShapes(compartmentID string, opt ...baremtlsdk.CoreOptions) (*baremtlsdk.ShapeList, error)
	ListVnicAttachments(compartmentID string, opt ...baremtlsdk.CoreOptions) (*baremtlsdk.VnicAttachmentList, error)
}
