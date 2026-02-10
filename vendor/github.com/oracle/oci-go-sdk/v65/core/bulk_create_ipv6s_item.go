// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// BulkCreateIpv6sItem Secondary IPv6 object to creation as part of bulk creation .
type BulkCreateIpv6sItem struct {

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

	// An IPv6 address of your choice. Must be an available IP address within
	// the subnet's CIDR. If you don't specify a value, Oracle automatically
	// assigns an IPv6 address from the subnet. The subnet is the one that
	// contains the VNIC you specify in `vnicId`.
	// Example: `2001:DB8::`
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// Lifetime of the IP address.
	// There are two types of IPs:
	//  - Ephemeral
	//  - Reserved
	Lifetime BulkCreateIpv6sItemLifetimeEnum `mandatory:"false" json:"lifetime,omitempty"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the IP address or VNIC will use. For more information, see
	// Per-resource Routing (https://docs.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm#Overview_of_Routing_for_Your_VCN__source_routing).
	RouteTableId *string `mandatory:"false" json:"routeTableId"`

	// The IPv6 prefix allocated to the subnet. This is required if more than one IPv6 prefix exists on the subnet.
	Ipv6SubnetCidr *string `mandatory:"false" json:"ipv6SubnetCidr"`

	// Length of cidr range. Optional field to specify flexible cidr.
	CidrPrefixLength *int `mandatory:"false" json:"cidrPrefixLength"`
}

func (m BulkCreateIpv6sItem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkCreateIpv6sItem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBulkCreateIpv6sItemLifetimeEnum(string(m.Lifetime)); !ok && m.Lifetime != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Lifetime: %s. Supported values are: %s.", m.Lifetime, strings.Join(GetBulkCreateIpv6sItemLifetimeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BulkCreateIpv6sItemLifetimeEnum Enum with underlying type: string
type BulkCreateIpv6sItemLifetimeEnum string

// Set of constants representing the allowable values for BulkCreateIpv6sItemLifetimeEnum
const (
	BulkCreateIpv6sItemLifetimeEphemeral BulkCreateIpv6sItemLifetimeEnum = "EPHEMERAL"
	BulkCreateIpv6sItemLifetimeReserved  BulkCreateIpv6sItemLifetimeEnum = "RESERVED"
)

var mappingBulkCreateIpv6sItemLifetimeEnum = map[string]BulkCreateIpv6sItemLifetimeEnum{
	"EPHEMERAL": BulkCreateIpv6sItemLifetimeEphemeral,
	"RESERVED":  BulkCreateIpv6sItemLifetimeReserved,
}

var mappingBulkCreateIpv6sItemLifetimeEnumLowerCase = map[string]BulkCreateIpv6sItemLifetimeEnum{
	"ephemeral": BulkCreateIpv6sItemLifetimeEphemeral,
	"reserved":  BulkCreateIpv6sItemLifetimeReserved,
}

// GetBulkCreateIpv6sItemLifetimeEnumValues Enumerates the set of values for BulkCreateIpv6sItemLifetimeEnum
func GetBulkCreateIpv6sItemLifetimeEnumValues() []BulkCreateIpv6sItemLifetimeEnum {
	values := make([]BulkCreateIpv6sItemLifetimeEnum, 0)
	for _, v := range mappingBulkCreateIpv6sItemLifetimeEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkCreateIpv6sItemLifetimeEnumStringValues Enumerates the set of values in String for BulkCreateIpv6sItemLifetimeEnum
func GetBulkCreateIpv6sItemLifetimeEnumStringValues() []string {
	return []string{
		"EPHEMERAL",
		"RESERVED",
	}
}

// GetMappingBulkCreateIpv6sItemLifetimeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkCreateIpv6sItemLifetimeEnum(val string) (BulkCreateIpv6sItemLifetimeEnum, bool) {
	enum, ok := mappingBulkCreateIpv6sItemLifetimeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
