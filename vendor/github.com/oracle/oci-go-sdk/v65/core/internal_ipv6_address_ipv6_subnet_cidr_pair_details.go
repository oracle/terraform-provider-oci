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

// InternalIpv6AddressIpv6SubnetCidrPairDetails Details to provide a pair of IPv6 Subnet Prefix and IPv6 address to assign IPv6 address on VNIC creation.
type InternalIpv6AddressIpv6SubnetCidrPairDetails struct {

	// The IPv6 prefix allocated to the subnet.
	Ipv6SubnetCidr *string `mandatory:"false" json:"ipv6SubnetCidr"`

	// An IPv6 address of your choice. Must be available IPv6 address within the subnet's CIDR.
	// If IPv6 address is not provided
	// - Oracle will automatically assign an IPv6 address from the subnet's IPv6 prefix if and only if there is only one IPv6 prefix on the subnet.
	// - Oracle will automatically assign an IPv6 address from the subnet's IPv6 Oracle GUA prefix if it exists on the subnet.
	Ipv6Address *string `mandatory:"false" json:"ipv6Address"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private IP to assign the vnic.
	Ipv6Id *string `mandatory:"false" json:"ipv6Id"`

	// Lifetime of the IP address.
	// There are two types of IPv6 IPs:
	//  - Ephemeral
	//  - Reserved
	Lifetime InternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnum `mandatory:"false" json:"lifetime,omitempty"`

	// The hostname associated with the IPv6 address. Only the hostname label, not the FQDN.
	Hostname *string `mandatory:"false" json:"hostname"`
}

func (m InternalIpv6AddressIpv6SubnetCidrPairDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalIpv6AddressIpv6SubnetCidrPairDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnum(string(m.Lifetime)); !ok && m.Lifetime != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Lifetime: %s. Supported values are: %s.", m.Lifetime, strings.Join(GetInternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnum Enum with underlying type: string
type InternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnum string

// Set of constants representing the allowable values for InternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnum
const (
	InternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEphemeral InternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnum = "EPHEMERAL"
	InternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeReserved  InternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnum = "RESERVED"
)

var mappingInternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnum = map[string]InternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnum{
	"EPHEMERAL": InternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEphemeral,
	"RESERVED":  InternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeReserved,
}

var mappingInternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnumLowerCase = map[string]InternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnum{
	"ephemeral": InternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEphemeral,
	"reserved":  InternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeReserved,
}

// GetInternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnumValues Enumerates the set of values for InternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnum
func GetInternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnumValues() []InternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnum {
	values := make([]InternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnum, 0)
	for _, v := range mappingInternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnumStringValues Enumerates the set of values in String for InternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnum
func GetInternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnumStringValues() []string {
	return []string{
		"EPHEMERAL",
		"RESERVED",
	}
}

// GetMappingInternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnum(val string) (InternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnum, bool) {
	enum, ok := mappingInternalIpv6AddressIpv6SubnetCidrPairDetailsLifetimeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
