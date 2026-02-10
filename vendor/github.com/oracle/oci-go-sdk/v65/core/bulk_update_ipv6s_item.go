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

// BulkUpdateIpv6sItem Secondary IPv6 object to update as part of bulk update.
type BulkUpdateIpv6sItem struct {

	// The OCID of the IPv6.
	Ipv6Id *string `mandatory:"true" json:"ipv6Id"`

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

	// Lifetime of the IP address.
	// There are two types of IPs:
	//  - Ephemeral
	//  - Reserved
	Lifetime BulkUpdateIpv6sItemLifetimeEnum `mandatory:"false" json:"lifetime,omitempty"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the IP address or VNIC will use. For more information, see
	// Per-resource Routing (https://docs.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm#Overview_of_Routing_for_Your_VCN__source_routing).
	RouteTableId *string `mandatory:"false" json:"routeTableId"`
}

func (m BulkUpdateIpv6sItem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkUpdateIpv6sItem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBulkUpdateIpv6sItemLifetimeEnum(string(m.Lifetime)); !ok && m.Lifetime != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Lifetime: %s. Supported values are: %s.", m.Lifetime, strings.Join(GetBulkUpdateIpv6sItemLifetimeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BulkUpdateIpv6sItemLifetimeEnum Enum with underlying type: string
type BulkUpdateIpv6sItemLifetimeEnum string

// Set of constants representing the allowable values for BulkUpdateIpv6sItemLifetimeEnum
const (
	BulkUpdateIpv6sItemLifetimeEphemeral BulkUpdateIpv6sItemLifetimeEnum = "EPHEMERAL"
	BulkUpdateIpv6sItemLifetimeReserved  BulkUpdateIpv6sItemLifetimeEnum = "RESERVED"
)

var mappingBulkUpdateIpv6sItemLifetimeEnum = map[string]BulkUpdateIpv6sItemLifetimeEnum{
	"EPHEMERAL": BulkUpdateIpv6sItemLifetimeEphemeral,
	"RESERVED":  BulkUpdateIpv6sItemLifetimeReserved,
}

var mappingBulkUpdateIpv6sItemLifetimeEnumLowerCase = map[string]BulkUpdateIpv6sItemLifetimeEnum{
	"ephemeral": BulkUpdateIpv6sItemLifetimeEphemeral,
	"reserved":  BulkUpdateIpv6sItemLifetimeReserved,
}

// GetBulkUpdateIpv6sItemLifetimeEnumValues Enumerates the set of values for BulkUpdateIpv6sItemLifetimeEnum
func GetBulkUpdateIpv6sItemLifetimeEnumValues() []BulkUpdateIpv6sItemLifetimeEnum {
	values := make([]BulkUpdateIpv6sItemLifetimeEnum, 0)
	for _, v := range mappingBulkUpdateIpv6sItemLifetimeEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkUpdateIpv6sItemLifetimeEnumStringValues Enumerates the set of values in String for BulkUpdateIpv6sItemLifetimeEnum
func GetBulkUpdateIpv6sItemLifetimeEnumStringValues() []string {
	return []string{
		"EPHEMERAL",
		"RESERVED",
	}
}

// GetMappingBulkUpdateIpv6sItemLifetimeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkUpdateIpv6sItemLifetimeEnum(val string) (BulkUpdateIpv6sItemLifetimeEnum, bool) {
	enum, ok := mappingBulkUpdateIpv6sItemLifetimeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
