package main

import "github.com/MustWin/baremetal-sdk-go"

type CreateResourceFn func(string, string, ...baremetal.Options) (*baremetal.IdentityResource, error)
type GetResourceFn func(string) (*baremetal.IdentityResource, error)
type UpdateResourceFn func(string, string, ...baremetal.Options) (*baremetal.IdentityResource, error)
type DeleteResourceFn func(string, ...baremetal.Options) error

type BareMetalClient interface {
	CreateUser(name, description string, options ...baremetal.Options) (*baremetal.IdentityResource, error)
	GetUser(userID string) (*baremetal.IdentityResource, error)
	UpdateUser(userID, userDescription string, opts ...baremetal.Options) (*baremetal.IdentityResource, error)
	DeleteUser(userID string, opts ...baremetal.Options) error

	CreateGroup(name, description string, options ...baremetal.Options) (*baremetal.IdentityResource, error)
	GetGroup(userID string) (*baremetal.IdentityResource, error)
	UpdateGroup(userID, userDescription string, opts ...baremetal.Options) (*baremetal.IdentityResource, error)
	DeleteGroup(userID string, opts ...baremetal.Options) error

	CreatePolicy(name, description string, statements []string, opts ...baremetal.Options) (*baremetal.Policy, error)
	GetPolicy(id string) (*baremetal.Policy, error)
	UpdatePolicy(id, description string, statements []string, opts ...baremetal.Options) (*baremetal.Policy, error)
	DeletePolicy(id string, opts ...baremetal.Options) error

	CreateCompartment(name, description string, options ...baremetal.Options) (*baremetal.IdentityResource, error)
	GetCompartment(userID string) (*baremetal.IdentityResource, error)
	UpdateCompartment(userID, userDescription string, opts ...baremetal.Options) (*baremetal.IdentityResource, error)

	ListShapes(compartmentID string, opt ...baremetal.CoreOptions) (*baremetal.ShapeList, error)
	ListVnicAttachments(compartmentID string, opt ...baremetal.CoreOptions) (*baremetal.VnicAttachmentList, error)
}
