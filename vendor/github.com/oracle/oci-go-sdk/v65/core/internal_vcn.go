// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InternalVcn A virtual cloud network (VCN). For more information, see
// Overview of the Networking Service (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
type InternalVcn struct {

	// Deprecated. The first CIDR IP address from cidrBlocks.
	// Example: `172.16.0.0/16`
	CidrBlock *string `mandatory:"true" json:"cidrBlock"`

	// The list of IPv4 CIDR blocks the VCN will use.
	CidrBlocks []string `mandatory:"true" json:"cidrBlocks"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the VCN.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The VCN's Oracle ID (OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
	Id *string `mandatory:"true" json:"id"`

	// The VCN's current state.
	LifecycleState InternalVcnLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The list of BYOIPv6 prefixes required to create a VCN that uses BYOIPv6 ranges.
	Byoipv6CidrBlocks []string `mandatory:"false" json:"byoipv6CidrBlocks"`

	// For an IPv6-enabled VCN, this is the list of Private IPv6 prefixes for the VCN's IP address space.
	Ipv6PrivateCidrBlocks []string `mandatory:"false" json:"ipv6PrivateCidrBlocks"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the VCN's default set of DHCP options.
	DefaultDhcpOptionsId *string `mandatory:"false" json:"defaultDhcpOptionsId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the VCN's default route table.
	DefaultRouteTableId *string `mandatory:"false" json:"defaultRouteTableId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the VCN's default security list.
	DefaultSecurityListId *string `mandatory:"false" json:"defaultSecurityListId"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A DNS label for the VCN, used in conjunction with the VNIC's hostname and
	// subnet's DNS label to form a fully qualified domain name (FQDN) for each VNIC
	// within this subnet (for example, `bminstance1.subnet123.vcn1.oraclevcn.com`).
	// Must be an alphanumeric string that begins with a letter.
	// The value cannot be changed.
	// The absence of this parameter means the Internet and VCN Resolver will
	// not work for this VCN.
	// For more information, see
	// DNS in Your Virtual Cloud Network (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).
	// Example: `vcn1`
	DnsLabel *string `mandatory:"false" json:"dnsLabel"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Security Attributes for this resource. This is unique to ZPR, and helps identify which resources are allowed to be accessed by what permission controls.
	// Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// For an IPv6-enabled VCN, this is the list of IPv6 prefixes for the VCN's IP address space.
	// The prefixes are provided by Oracle and the sizes are always /56.
	Ipv6CidrBlocks []string `mandatory:"false" json:"ipv6CidrBlocks"`

	// For an IPv6-enabled VCN, this is the IPv6 prefix for the VCN's private IP address space.
	// The VCN size is always /56. Oracle
	// provides the IPv6 prefix to use as the *same* CIDR for the `ipv6PublicCidrBlock`.
	// When creating a subnet, specify the last 8 bits, 00 to FF.
	// See IPv6 Addresses (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/ipv6.htm).
	// Example: `2001:0db8:0123::/56`
	Ipv6CidrBlock *string `mandatory:"false" json:"ipv6CidrBlock"`

	// For an IPv6-enabled VCN, this is the IPv6 prefix for the VCN's public IP address space.
	// The VCN size is always /56. This prefix is always provided by Oracle. If you don't provide a
	// custom prefix for the `ipv6CidrBlock` when creating the VCN, Oracle assigns that value and also
	// uses it for `ipv6PublicCidrBlock`. Oracle uses addresses from this block for the `publicIpAddress`
	// attribute of an Ipv6 that has internet access allowed.
	// Example: `2001:0db8:0123::/48`
	Ipv6PublicCidrBlock *string `mandatory:"false" json:"ipv6PublicCidrBlock"`

	// The date and time the VCN was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The VCN's domain name, which consists of the VCN's DNS label, and the
	// `oraclevcn.com` domain.
	// For more information, see
	// DNS in Your Virtual Cloud Network (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).
	// Example: `vcn1.oraclevcn.com`
	VcnDomainName *string `mandatory:"false" json:"vcnDomainName"`

	// Indicates whether traffic within the VCN is encrypted.
	// For more information, see VN Encryption (https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/Overview_of_VCNs_and_Subnets.htm#encryption).
	IsEncrypted *bool `mandatory:"false" json:"isEncrypted"`

	// Indicates if a VCN is Substrate Service VCN. False by default.
	IsSubstrateService *bool `mandatory:"false" json:"isSubstrateService"`

	// Indicates if a VCN is Substrate Service VCN for a Ring 0 service (Compute, Block Storage, VCN). False by default.
	IsRingZeroSubstrateService *bool `mandatory:"false" json:"isRingZeroSubstrateService"`
}

func (m InternalVcn) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalVcn) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInternalVcnLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetInternalVcnLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InternalVcnLifecycleStateEnum Enum with underlying type: string
type InternalVcnLifecycleStateEnum string

// Set of constants representing the allowable values for InternalVcnLifecycleStateEnum
const (
	InternalVcnLifecycleStateProvisioning InternalVcnLifecycleStateEnum = "PROVISIONING"
	InternalVcnLifecycleStateAvailable    InternalVcnLifecycleStateEnum = "AVAILABLE"
	InternalVcnLifecycleStateTerminating  InternalVcnLifecycleStateEnum = "TERMINATING"
	InternalVcnLifecycleStateTerminated   InternalVcnLifecycleStateEnum = "TERMINATED"
	InternalVcnLifecycleStateUpdating     InternalVcnLifecycleStateEnum = "UPDATING"
)

var mappingInternalVcnLifecycleStateEnum = map[string]InternalVcnLifecycleStateEnum{
	"PROVISIONING": InternalVcnLifecycleStateProvisioning,
	"AVAILABLE":    InternalVcnLifecycleStateAvailable,
	"TERMINATING":  InternalVcnLifecycleStateTerminating,
	"TERMINATED":   InternalVcnLifecycleStateTerminated,
	"UPDATING":     InternalVcnLifecycleStateUpdating,
}

var mappingInternalVcnLifecycleStateEnumLowerCase = map[string]InternalVcnLifecycleStateEnum{
	"provisioning": InternalVcnLifecycleStateProvisioning,
	"available":    InternalVcnLifecycleStateAvailable,
	"terminating":  InternalVcnLifecycleStateTerminating,
	"terminated":   InternalVcnLifecycleStateTerminated,
	"updating":     InternalVcnLifecycleStateUpdating,
}

// GetInternalVcnLifecycleStateEnumValues Enumerates the set of values for InternalVcnLifecycleStateEnum
func GetInternalVcnLifecycleStateEnumValues() []InternalVcnLifecycleStateEnum {
	values := make([]InternalVcnLifecycleStateEnum, 0)
	for _, v := range mappingInternalVcnLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalVcnLifecycleStateEnumStringValues Enumerates the set of values in String for InternalVcnLifecycleStateEnum
func GetInternalVcnLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
		"UPDATING",
	}
}

// GetMappingInternalVcnLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalVcnLifecycleStateEnum(val string) (InternalVcnLifecycleStateEnum, bool) {
	enum, ok := mappingInternalVcnLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
