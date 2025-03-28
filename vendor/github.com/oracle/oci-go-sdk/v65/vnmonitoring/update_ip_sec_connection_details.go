// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateIpSecConnectionDetails The representation of UpdateIpSecConnectionDetails
type UpdateIpSecConnectionDetails struct {

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Your identifier for your CPE device. Can be either an IP address or a hostname (specifically, the
	// fully qualified domain name (FQDN)). The type of identifier you provide here must correspond
	// to the value for `cpeLocalIdentifierType`.
	// For information about why you'd provide this value, see
	// If Your CPE Is Behind a NAT Device (https://docs.oracle.com/iaas/Content/Network/Tasks/overviewIPsec.htm#nat).
	// Example IP address: `10.0.3.3`
	// Example hostname: `cpe.example.com`
	CpeLocalIdentifier *string `mandatory:"false" json:"cpeLocalIdentifier"`

	// The type of identifier for your CPE device. The value you provide here must correspond to the value
	// for `cpeLocalIdentifier`.
	CpeLocalIdentifierType UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum `mandatory:"false" json:"cpeLocalIdentifierType,omitempty"`

	// Static routes to the CPE. If you provide this attribute, it replaces the entire current set of
	// static routes. A static route's CIDR must not be a multicast address or class E address.
	// The CIDR can be either IPv4 or IPv6.
	// IPv6 addressing is supported for all commercial and government regions.
	// See IPv6 Addresses (https://docs.oracle.com/iaas/Content/Network/Concepts/ipv6.htm).
	// Example: `10.0.1.0/24`
	// Example: `2001:db8::/32`
	StaticRoutes []string `mandatory:"false" json:"staticRoutes"`
}

func (m UpdateIpSecConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateIpSecConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum(string(m.CpeLocalIdentifierType)); !ok && m.CpeLocalIdentifierType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CpeLocalIdentifierType: %s. Supported values are: %s.", m.CpeLocalIdentifierType, strings.Join(GetUpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum Enum with underlying type: string
type UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum string

// Set of constants representing the allowable values for UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum
const (
	UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeIpAddress UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum = "IP_ADDRESS"
	UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeHostname  UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum = "HOSTNAME"
)

var mappingUpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum = map[string]UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum{
	"IP_ADDRESS": UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeIpAddress,
	"HOSTNAME":   UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeHostname,
}

var mappingUpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnumLowerCase = map[string]UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum{
	"ip_address": UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeIpAddress,
	"hostname":   UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeHostname,
}

// GetUpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnumValues Enumerates the set of values for UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum
func GetUpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnumValues() []UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum {
	values := make([]UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum, 0)
	for _, v := range mappingUpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnumStringValues Enumerates the set of values in String for UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum
func GetUpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnumStringValues() []string {
	return []string{
		"IP_ADDRESS",
		"HOSTNAME",
	}
}

// GetMappingUpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum(val string) (UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum, bool) {
	enum, ok := mappingUpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
