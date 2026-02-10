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

// BulkCreatePrivateIpItem Secondary private IPv4 address object to create as part of bulk creation.
type BulkCreatePrivateIpItem struct {

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

	// The hostname for the private IP. Used for DNS. The value
	// is the hostname portion of the private IP's fully qualified domain name (FQDN)
	// (for example, `bminstance1` in FQDN `bminstance1.subnet123.vcn1.oraclevcn.com`).
	// Must be unique across all VNICs in the subnet and comply with
	// RFC 952 (https://tools.ietf.org/html/rfc952) and
	// RFC 1123 (https://tools.ietf.org/html/rfc1123).
	// For more information, see
	// DNS in Your Virtual Cloud Network (https://docs.oracle.com/iaas/Content/Network/Concepts/dns.htm).
	// Example: `bminstance1`
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// A private IP address of your choice. Must be an available IP address within
	// the subnet's CIDR. If you don't specify a value, Oracle automatically
	// assigns a private IP address from the subnet.
	// Example: `10.0.3.3`
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// Lifetime of the IP address.
	// There are two types of IPs:
	//  - Ephemeral
	//  - Reserved
	Lifetime BulkCreatePrivateIpItemLifetimeEnum `mandatory:"false" json:"lifetime,omitempty"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the IP address or VNIC will use. For more information, see
	// Per-resource Routing (https://docs.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm#Overview_of_Routing_for_Your_VCN__source_routing).
	RouteTableId *string `mandatory:"false" json:"routeTableId"`

	// An optional field that when combined with the ipAddress field, will be used to allocate secondary IPv4 CIDRs.
	// The CIDR range created by this combination must be within the subnet's CIDR
	// and the CIDR range should not collide with any existing IPv4 address allocation.
	// The VNIC ID specified in the request object should not already been assigned more than the max IPv4 addresses.
	// If you don't specify a value, this option will be ignored.
	// Example: 18
	CidrPrefixLength *int `mandatory:"false" json:"cidrPrefixLength"`

	// Any one of the IPv4 CIDRs allocated to the subnet.
	Ipv4SubnetCidrAtCreation *string `mandatory:"false" json:"ipv4SubnetCidrAtCreation"`
}

func (m BulkCreatePrivateIpItem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkCreatePrivateIpItem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBulkCreatePrivateIpItemLifetimeEnum(string(m.Lifetime)); !ok && m.Lifetime != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Lifetime: %s. Supported values are: %s.", m.Lifetime, strings.Join(GetBulkCreatePrivateIpItemLifetimeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BulkCreatePrivateIpItemLifetimeEnum Enum with underlying type: string
type BulkCreatePrivateIpItemLifetimeEnum string

// Set of constants representing the allowable values for BulkCreatePrivateIpItemLifetimeEnum
const (
	BulkCreatePrivateIpItemLifetimeEphemeral BulkCreatePrivateIpItemLifetimeEnum = "EPHEMERAL"
	BulkCreatePrivateIpItemLifetimeReserved  BulkCreatePrivateIpItemLifetimeEnum = "RESERVED"
)

var mappingBulkCreatePrivateIpItemLifetimeEnum = map[string]BulkCreatePrivateIpItemLifetimeEnum{
	"EPHEMERAL": BulkCreatePrivateIpItemLifetimeEphemeral,
	"RESERVED":  BulkCreatePrivateIpItemLifetimeReserved,
}

var mappingBulkCreatePrivateIpItemLifetimeEnumLowerCase = map[string]BulkCreatePrivateIpItemLifetimeEnum{
	"ephemeral": BulkCreatePrivateIpItemLifetimeEphemeral,
	"reserved":  BulkCreatePrivateIpItemLifetimeReserved,
}

// GetBulkCreatePrivateIpItemLifetimeEnumValues Enumerates the set of values for BulkCreatePrivateIpItemLifetimeEnum
func GetBulkCreatePrivateIpItemLifetimeEnumValues() []BulkCreatePrivateIpItemLifetimeEnum {
	values := make([]BulkCreatePrivateIpItemLifetimeEnum, 0)
	for _, v := range mappingBulkCreatePrivateIpItemLifetimeEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkCreatePrivateIpItemLifetimeEnumStringValues Enumerates the set of values in String for BulkCreatePrivateIpItemLifetimeEnum
func GetBulkCreatePrivateIpItemLifetimeEnumStringValues() []string {
	return []string{
		"EPHEMERAL",
		"RESERVED",
	}
}

// GetMappingBulkCreatePrivateIpItemLifetimeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkCreatePrivateIpItemLifetimeEnum(val string) (BulkCreatePrivateIpItemLifetimeEnum, bool) {
	enum, ok := mappingBulkCreatePrivateIpItemLifetimeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
