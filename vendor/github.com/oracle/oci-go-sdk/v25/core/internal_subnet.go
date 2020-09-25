// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// InternalSubnet A logical subdivision of a VCN. Each subnet
// consists of a contiguous range of IP addresses that do not overlap with
// other subnets in the VCN. Example: 172.16.1.0/24. For more information, see
// Overview of the Networking Service (https://docs.cloud.oracle.com/Content/Network/Concepts/overview.htm) and
// VCNs and Subnets (https://docs.cloud.oracle.com/Content/Network/Tasks/managingVCNs.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you
// supply string values using the API.
type InternalSubnet struct {

	// The subnet's CIDR block.
	// Example: `10.0.1.0/24`
	CidrBlock *string `mandatory:"true" json:"cidrBlock"`

	// The OCID of the compartment containing the subnet.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The subnet's Oracle ID (OCID).
	Id *string `mandatory:"true" json:"id"`

	// The subnet's current state.
	LifecycleState InternalSubnetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the route table that the subnet uses.
	RouteTableId *string `mandatory:"true" json:"routeTableId"`

	// The OCID of the VCN the subnet is in.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The IP address of the virtual router.
	// Example: `10.0.14.1`
	VirtualRouterIp *string `mandatory:"true" json:"virtualRouterIp"`

	// The MAC address of the virtual router.
	// Example: `00:00:00:00:00:01`
	VirtualRouterMac *string `mandatory:"true" json:"virtualRouterMac"`

	// The subnet's availability domain. This attribute will be null if this is a regional subnet
	// instead of an AD-specific subnet. Oracle recommends creating regional subnets.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID of the set of DHCP options that the subnet uses.
	DhcpOptionsId *string `mandatory:"false" json:"dhcpOptionsId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A DNS label for the subnet, used in conjunction with the VNIC's hostname and
	// VCN's DNS label to form a fully qualified domain name (FQDN) for each VNIC
	// within this subnet (for example, `bminstance-1.subnet123.vcn1.oraclevcn.com`).
	// Must be an alphanumeric string that begins with a letter and is unique within the VCN.
	// The value cannot be changed.
	// The absence of this parameter means the Internet and VCN Resolver
	// will not resolve hostnames of instances in this subnet.
	// For more information, see
	// DNS in Your Virtual Cloud Network (https://docs.cloud.oracle.com/Content/Network/Concepts/dns.htm).
	// Example: `subnet123`
	DnsLabel *string `mandatory:"false" json:"dnsLabel"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The OCID of the Internal Hosted Zone the DnsRecord belongs to.
	InternalHostedZoneId *string `mandatory:"false" json:"internalHostedZoneId"`

	// For an IPv6-enabled subnet, this is the IPv6 CIDR block for the subnet's private IP address
	// space. The subnet size is always /64. Note that IPv6 addressing is currently supported only
	// in certain regions. See IPv6 Addresses (https://docs.cloud.oracle.com/Content/Network/Concepts/ipv6.htm).
	// Example: `2001:0db8:0123:1111::/64`
	Ipv6CidrBlock *string `mandatory:"false" json:"ipv6CidrBlock"`

	// For an IPv6-enabled subnet, this is the IPv6 CIDR block for the subnet's public IP address
	// space. The subnet size is always /64. The left 48 bits are inherited from the
	// `ipv6PublicCidrBlock` of the Vcn,
	// and the remaining 16 bits are from the subnet's `ipv6CidrBlock`.
	// Example: `2001:0db8:0123:1111::/64`
	Ipv6PublicCidrBlock *string `mandatory:"false" json:"ipv6PublicCidrBlock"`

	// For an IPv6-enabled subnet, this is the IPv6 address of the virtual router.
	// Example: `2001:0db8:0123:1111:89ab:cdef:1234:5678`
	Ipv6VirtualRouterIp *string `mandatory:"false" json:"ipv6VirtualRouterIp"`

	// Whether learning mode is enabled for this subnet. The default is `false`.
	// **Note:** When a subnet has learning mode enabled, only certain types
	// of resources can be launched in the subnet.
	// Example: `true`
	IsLearningEnabled *bool `mandatory:"false" json:"isLearningEnabled"`

	// The VLAN tag assigned to VNIC Attachments within this Subnet if the Subnet has learning enabled.
	// **Note:** When a subnet does not have learning enabled, this field will be null.
	// Example: `100`
	VlanTag *int `mandatory:"false" json:"vlanTag"`

	// Whether VNICs within this subnet can have public IP addresses.
	// Defaults to false, which means VNICs created in this subnet will
	// automatically be assigned public IP addresses unless specified
	// otherwise during instance launch or VNIC creation (with the
	// `assignPublicIp` flag in
	// CreateVnicDetails).
	// If `prohibitPublicIpOnVnic` is set to true, VNICs created in this
	// subnet cannot have public IP addresses (that is, it's a private
	// subnet).
	// Example: `true`
	ProhibitPublicIpOnVnic *bool `mandatory:"false" json:"prohibitPublicIpOnVnic"`

	// The OCIDs of the security list or lists that the subnet uses. Remember
	// that security lists are associated *with the subnet*, but the
	// rules are applied to the individual VNICs in the subnet.
	SecurityListIds []string `mandatory:"false" json:"securityListIds"`

	// The subnet's domain name, which consists of the subnet's DNS label,
	// the VCN's DNS label, and the `oraclevcn.com` domain.
	// For more information, see
	// DNS in Your Virtual Cloud Network (https://docs.cloud.oracle.com/Content/Network/Concepts/dns.htm).
	// Example: `subnet123.vcn1.oraclevcn.com`
	SubnetDomainName *string `mandatory:"false" json:"subnetDomainName"`

	// The date and time the subnet was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m InternalSubnet) String() string {
	return common.PointerString(m)
}

// InternalSubnetLifecycleStateEnum Enum with underlying type: string
type InternalSubnetLifecycleStateEnum string

// Set of constants representing the allowable values for InternalSubnetLifecycleStateEnum
const (
	InternalSubnetLifecycleStateProvisioning InternalSubnetLifecycleStateEnum = "PROVISIONING"
	InternalSubnetLifecycleStateAvailable    InternalSubnetLifecycleStateEnum = "AVAILABLE"
	InternalSubnetLifecycleStateTerminating  InternalSubnetLifecycleStateEnum = "TERMINATING"
	InternalSubnetLifecycleStateTerminated   InternalSubnetLifecycleStateEnum = "TERMINATED"
	InternalSubnetLifecycleStateUpdating     InternalSubnetLifecycleStateEnum = "UPDATING"
)

var mappingInternalSubnetLifecycleState = map[string]InternalSubnetLifecycleStateEnum{
	"PROVISIONING": InternalSubnetLifecycleStateProvisioning,
	"AVAILABLE":    InternalSubnetLifecycleStateAvailable,
	"TERMINATING":  InternalSubnetLifecycleStateTerminating,
	"TERMINATED":   InternalSubnetLifecycleStateTerminated,
	"UPDATING":     InternalSubnetLifecycleStateUpdating,
}

// GetInternalSubnetLifecycleStateEnumValues Enumerates the set of values for InternalSubnetLifecycleStateEnum
func GetInternalSubnetLifecycleStateEnumValues() []InternalSubnetLifecycleStateEnum {
	values := make([]InternalSubnetLifecycleStateEnum, 0)
	for _, v := range mappingInternalSubnetLifecycleState {
		values = append(values, v)
	}
	return values
}
