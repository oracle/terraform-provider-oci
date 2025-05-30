// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrivateIp A *private IP* is a conceptual term that refers to an IPv4 private IP address and related properties.
// The `privateIp` object is the API representation of a private IP.
// **Note:** For information about IPv6 addresses, see Ipv6.
// Each instance has a *primary private IP* that is automatically created and
// assigned to the primary VNIC during instance launch. If you add a secondary
// VNIC to the instance, it also automatically gets a primary private IP. You
// can't remove a primary private IP from its VNIC. The primary private IP is
// automatically deleted when the VNIC is terminated.
// You can add *secondary private IPs* to a VNIC after it's created. For more
// information, see the `privateIp` operations and also
// IP Addresses (https://docs.oracle.com/iaas/Content/Network/Tasks/managingIPaddresses.htm).
// **Note:** Only
// ListPrivateIps and
// GetPrivateIp work with
// *primary* private IPs. To create and update primary private IPs, you instead
// work with instance and VNIC operations. For example, a primary private IP's
// properties come from the values you specify in
// CreateVnicDetails when calling either
// LaunchInstance or
// AttachVnic. To update the hostname
// for a primary private IP, you use UpdateVnic.
// `PrivateIp` objects that are created for use with the Oracle Cloud VMware Solution are
// assigned to a VLAN and not a VNIC in a subnet. See the
// descriptions of the relevant attributes in the `PrivateIp` object. Also see
// Vlan.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
type PrivateIp struct {

	// The private IP's availability domain. This attribute will be null if this is a *secondary*
	// private IP assigned to a VNIC that is in a *regional* subnet.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the private IP.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The hostname for the private IP. Used for DNS. The value is the hostname
	// portion of the private IP's fully qualified domain name (FQDN)
	// (for example, `bminstance1` in FQDN `bminstance1.subnet123.vcn1.oraclevcn.com`).
	// Must be unique across all VNICs in the subnet and comply with
	// RFC 952 (https://tools.ietf.org/html/rfc952) and
	// RFC 1123 (https://tools.ietf.org/html/rfc1123).
	// For more information, see
	// DNS in Your Virtual Cloud Network (https://docs.oracle.com/iaas/Content/Network/Concepts/dns.htm).
	// Example: `bminstance1`
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// The private IP's Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
	Id *string `mandatory:"false" json:"id"`

	// The private IP address of the `privateIp` object. The address is within the CIDR
	// of the VNIC's subnet.
	// However, if the `PrivateIp` object is being used with a VLAN as part of
	// the Oracle Cloud VMware Solution, the address is from the range specified by the
	// `cidrBlock` attribute for the VLAN. See Vlan.
	// Example: `10.0.3.3`
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// Whether this private IP is the primary one on the VNIC. Primary private IPs
	// are unassigned and deleted automatically when the VNIC is terminated.
	// Example: `true`
	IsPrimary *bool `mandatory:"false" json:"isPrimary"`

	// Applicable only if the `PrivateIp` object is being used with a VLAN as part of
	// the Oracle Cloud VMware Solution. The `vlanId` is the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN. See
	// Vlan.
	VlanId *string `mandatory:"false" json:"vlanId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the VNIC is in.
	// However, if the `PrivateIp` object is being used with a VLAN as part of
	// the Oracle Cloud VMware Solution, the `subnetId` is null.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The date and time the private IP was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC the private IP is assigned to. The VNIC and private IP
	// must be in the same subnet.
	// However, if the `PrivateIp` object is being used with a VLAN as part of
	// the Oracle Cloud VMware Solution, the `vnicId` is null.
	VnicId *string `mandatory:"false" json:"vnicId"`

	// State of the IP address. If an IP address is assigned to a VNIC it is ASSIGNED, otherwise it is AVAILABLE.
	IpState PrivateIpIpStateEnum `mandatory:"false" json:"ipState,omitempty"`

	// Lifetime of the IP address.
	// There are two types of IPs:
	//  - Ephemeral
	//  - Reserved
	Lifetime PrivateIpLifetimeEnum `mandatory:"false" json:"lifetime,omitempty"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the IP address or VNIC will use. For more information, see
	// Per-resource Routing (https://docs.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm#Overview_of_Routing_for_Your_VCN__source_routing).
	RouteTableId *string `mandatory:"false" json:"routeTableId"`
}

func (m PrivateIp) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrivateIp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPrivateIpIpStateEnum(string(m.IpState)); !ok && m.IpState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IpState: %s. Supported values are: %s.", m.IpState, strings.Join(GetPrivateIpIpStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPrivateIpLifetimeEnum(string(m.Lifetime)); !ok && m.Lifetime != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Lifetime: %s. Supported values are: %s.", m.Lifetime, strings.Join(GetPrivateIpLifetimeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PrivateIpIpStateEnum Enum with underlying type: string
type PrivateIpIpStateEnum string

// Set of constants representing the allowable values for PrivateIpIpStateEnum
const (
	PrivateIpIpStateAssigned  PrivateIpIpStateEnum = "ASSIGNED"
	PrivateIpIpStateAvailable PrivateIpIpStateEnum = "AVAILABLE"
)

var mappingPrivateIpIpStateEnum = map[string]PrivateIpIpStateEnum{
	"ASSIGNED":  PrivateIpIpStateAssigned,
	"AVAILABLE": PrivateIpIpStateAvailable,
}

var mappingPrivateIpIpStateEnumLowerCase = map[string]PrivateIpIpStateEnum{
	"assigned":  PrivateIpIpStateAssigned,
	"available": PrivateIpIpStateAvailable,
}

// GetPrivateIpIpStateEnumValues Enumerates the set of values for PrivateIpIpStateEnum
func GetPrivateIpIpStateEnumValues() []PrivateIpIpStateEnum {
	values := make([]PrivateIpIpStateEnum, 0)
	for _, v := range mappingPrivateIpIpStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPrivateIpIpStateEnumStringValues Enumerates the set of values in String for PrivateIpIpStateEnum
func GetPrivateIpIpStateEnumStringValues() []string {
	return []string{
		"ASSIGNED",
		"AVAILABLE",
	}
}

// GetMappingPrivateIpIpStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPrivateIpIpStateEnum(val string) (PrivateIpIpStateEnum, bool) {
	enum, ok := mappingPrivateIpIpStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PrivateIpLifetimeEnum Enum with underlying type: string
type PrivateIpLifetimeEnum string

// Set of constants representing the allowable values for PrivateIpLifetimeEnum
const (
	PrivateIpLifetimeEphemeral PrivateIpLifetimeEnum = "EPHEMERAL"
	PrivateIpLifetimeReserved  PrivateIpLifetimeEnum = "RESERVED"
)

var mappingPrivateIpLifetimeEnum = map[string]PrivateIpLifetimeEnum{
	"EPHEMERAL": PrivateIpLifetimeEphemeral,
	"RESERVED":  PrivateIpLifetimeReserved,
}

var mappingPrivateIpLifetimeEnumLowerCase = map[string]PrivateIpLifetimeEnum{
	"ephemeral": PrivateIpLifetimeEphemeral,
	"reserved":  PrivateIpLifetimeReserved,
}

// GetPrivateIpLifetimeEnumValues Enumerates the set of values for PrivateIpLifetimeEnum
func GetPrivateIpLifetimeEnumValues() []PrivateIpLifetimeEnum {
	values := make([]PrivateIpLifetimeEnum, 0)
	for _, v := range mappingPrivateIpLifetimeEnum {
		values = append(values, v)
	}
	return values
}

// GetPrivateIpLifetimeEnumStringValues Enumerates the set of values in String for PrivateIpLifetimeEnum
func GetPrivateIpLifetimeEnumStringValues() []string {
	return []string{
		"EPHEMERAL",
		"RESERVED",
	}
}

// GetMappingPrivateIpLifetimeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPrivateIpLifetimeEnum(val string) (PrivateIpLifetimeEnum, bool) {
	enum, ok := mappingPrivateIpLifetimeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
