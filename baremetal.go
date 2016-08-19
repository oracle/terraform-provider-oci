package main

import "github.com/MustWin/baremetal-sdk-go"

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

	ListShapes(compartmentID string, opt ...baremetal.Options) (*baremetal.ShapeList, error)
	ListVnicAttachments(compartmentID string, opt ...baremetal.Options) (*baremetal.VnicAttachmentList, error)

	CreateCpe(compartmentID, displayName, IPAddress string, opts ...baremetal.Options) (cpe *baremetal.Cpe, e error)
	GetCpe(id string, opts ...baremetal.Options) (cpe *baremetal.Cpe, e error)
	DeleteCpe(id string, opts ...baremetal.Options) (e error)

	CreateVolume(availabiltyDomain, compartmentID string, opts ...baremetal.Options) (vol *baremetal.Volume, e error)
	GetVolume(id string, opts ...baremetal.Options) (vol *baremetal.Volume, e error)
	UpdateVolume(id string, opts ...baremetal.Options) (vol *baremetal.Volume, e error)
	DeleteVolume(id string, opts ...baremetal.Options) (e error)

	AttachVolume(compartmentID, instanceID, attachmentType, volumeID string, opts ...baremetal.Options) (vol *baremetal.VolumeAttachment, e error)
	GetVolumeAttachment(id string, opts ...baremetal.Options) (vol *baremetal.VolumeAttachment, e error)
	DetachVolume(id string, opts ...baremetal.Options) (e error)
}
