package client

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

	ListShapes(compartmentID string, opt ...baremetal.Options) (*baremetal.ListShapes, error)
	ListVnicAttachments(compartmentID string, opt ...baremetal.Options) (*baremetal.ListVnicAttachments, error)

	CreateCpe(compartmentID, displayName, IPAddress string, opts ...baremetal.Options) (cpe *baremetal.Cpe, e error)
	GetCpe(id string, opts ...baremetal.Options) (cpe *baremetal.Cpe, e error)
	DeleteCpe(id string, opts ...baremetal.Options) (e error)
	ListCpes(compartmentID string, opts ...baremetal.Options) (cpes *baremetal.ListCpes, e error)

	CreateVolume(availabiltyDomain, compartmentID string, opts ...baremetal.Options) (vol *baremetal.Volume, e error)
	GetVolume(id string, opts ...baremetal.Options) (vol *baremetal.Volume, e error)
	UpdateVolume(id string, opts ...baremetal.Options) (vol *baremetal.Volume, e error)
	DeleteVolume(id string, opts ...baremetal.Options) (e error)
	ListVolumes(compartmentID string, opts ...baremetal.Options) (vols *baremetal.ListVolumes, e error)

	LaunchInstance(availabilityDomain, compartmentID, image, shape, subnetID string, metadata map[string]string, opts ...baremetal.Options) (inst *baremetal.Instance, e error)
	GetInstance(instanceID string) (inst *baremetal.Instance, e error)
	UpdateInstance(instanceID string, opts ...baremetal.Options) (inst *baremetal.Instance, e error)
	TerminateInstance(instanceID string, opts ...baremetal.Options) (e error)
	ListInstances(compartmentID string, opts ...baremetal.Options) (list *baremetal.ListInstances, e error)

	AttachVolume(compartmentID, instanceID, attachmentType, volumeID string, opts ...baremetal.Options) (vol *baremetal.VolumeAttachment, e error)
	GetVolumeAttachment(id string, opts ...baremetal.Options) (vol *baremetal.VolumeAttachment, e error)
	DetachVolume(id string, opts ...baremetal.Options) (e error)
	ListVolumeAttachments(compartmentID string, opts ...baremetal.Options) (res *baremetal.VolumeAttachmentList, e error)

	CreateSubnet(availabilityDomain, cidrBlock, compartmentID, routeTableID, vcnID string, securityListIDs []string, opts ...baremetal.Options) (*baremetal.Subnet, error)
	GetSubnet(subnetID string) (sn *baremetal.Subnet, e error)
	ListSubnets(compartmentID, vcnID string, opts ...baremetal.Options) (*baremetal.ListSubnets, error)
	DeleteSubnet(subnetID string, opts ...baremetal.Options) error

	CreateVirtualNetwork(cidrBlock, compartmentID string, opts ...baremetal.Options) (*baremetal.VirtualNetwork, error)
	GetVirtualNetwork(id string, opts ...baremetal.Options) (vcn *baremetal.VirtualNetwork, e error)
	DeleteVirtualNetwork(id string, opts ...baremetal.Options) error
	ListVirtualNetworks(compartmentID string, opts ...baremetal.Options) (*baremetal.ListVirtualNetworks, error)

	CreateIPSecConnection(compartmentID, cpeID, drgID string, staticRoutes []string, opts ...baremetal.Options) (conn *baremetal.IPSecConnection, e error)
	ListIPSecConnections(compartmentID string, opts ...baremetal.Options) (conns *baremetal.ListIPSecConnections, e error)
	GetIPSecConnection(id string) (conn *baremetal.IPSecConnection, e error)
	DeleteIPSecConnection(id string, opts ...baremetal.Options) (e error)
	GetIPSecConnectionDeviceStatus(id string) (status *baremetal.IPSecConnectionDeviceStatus, e error)
	GetIPSecConnectionDeviceConfig(id string) (status *baremetal.IPSecConnectionDeviceConfig, e error)

	CreateDrg(compartmentID string, opts ...baremetal.Options) (*baremetal.Drg, error)
	GetDrg(id string, opts ...baremetal.Options) (*baremetal.Drg, error)
	DeleteDrg(id string, opts ...baremetal.Options) error
	ListDrgs(compartmentID string, opts ...baremetal.Options) (*baremetal.ListDrgs, error)

	CreateDrgAttachment(compartmentID, drgID, vcnID string, opts ...baremetal.Options) (vol *baremetal.DrgAttachment, e error)
	GetDrgAttachment(id string, opts ...baremetal.Options) (vol *baremetal.DrgAttachment, e error)
	DeleteDrgAttachment(id string, opts ...baremetal.Options) (e error)
	ListDrgAttachments(compartmentID string, opts ...baremetal.Options) (res *baremetal.ListDrgAttachments, e error)
}
