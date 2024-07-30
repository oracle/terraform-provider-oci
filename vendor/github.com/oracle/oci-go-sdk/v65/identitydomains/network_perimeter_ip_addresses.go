// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
// Use this pattern to construct endpoints for identity domains: `https://<domainURL>/admin/v1/`. See Finding an Identity Domain URL (https://docs.oracle.com/en-us/iaas/Content/Identity/api-getstarted/locate-identity-domain-url.htm) to locate the domain URL you need.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NetworkPerimeterIpAddresses IPAddresses or Ranges assigned to the NetworkPerimeter
type NetworkPerimeterIpAddresses struct {

	// Value of exact ipaddress or the range in CIDR or the range with start and end ip addresses
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`

	// type of the ip address value
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type NetworkPerimeterIpAddressesTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Indicates the type of Ip Address example, IPV4 or IPV6
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Version NetworkPerimeterIpAddressesVersionEnum `mandatory:"false" json:"version,omitempty"`
}

func (m NetworkPerimeterIpAddresses) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NetworkPerimeterIpAddresses) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingNetworkPerimeterIpAddressesTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetNetworkPerimeterIpAddressesTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNetworkPerimeterIpAddressesVersionEnum(string(m.Version)); !ok && m.Version != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Version: %s. Supported values are: %s.", m.Version, strings.Join(GetNetworkPerimeterIpAddressesVersionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NetworkPerimeterIpAddressesTypeEnum Enum with underlying type: string
type NetworkPerimeterIpAddressesTypeEnum string

// Set of constants representing the allowable values for NetworkPerimeterIpAddressesTypeEnum
const (
	NetworkPerimeterIpAddressesTypeCidr  NetworkPerimeterIpAddressesTypeEnum = "CIDR"
	NetworkPerimeterIpAddressesTypeRange NetworkPerimeterIpAddressesTypeEnum = "RANGE"
	NetworkPerimeterIpAddressesTypeExact NetworkPerimeterIpAddressesTypeEnum = "EXACT"
)

var mappingNetworkPerimeterIpAddressesTypeEnum = map[string]NetworkPerimeterIpAddressesTypeEnum{
	"CIDR":  NetworkPerimeterIpAddressesTypeCidr,
	"RANGE": NetworkPerimeterIpAddressesTypeRange,
	"EXACT": NetworkPerimeterIpAddressesTypeExact,
}

var mappingNetworkPerimeterIpAddressesTypeEnumLowerCase = map[string]NetworkPerimeterIpAddressesTypeEnum{
	"cidr":  NetworkPerimeterIpAddressesTypeCidr,
	"range": NetworkPerimeterIpAddressesTypeRange,
	"exact": NetworkPerimeterIpAddressesTypeExact,
}

// GetNetworkPerimeterIpAddressesTypeEnumValues Enumerates the set of values for NetworkPerimeterIpAddressesTypeEnum
func GetNetworkPerimeterIpAddressesTypeEnumValues() []NetworkPerimeterIpAddressesTypeEnum {
	values := make([]NetworkPerimeterIpAddressesTypeEnum, 0)
	for _, v := range mappingNetworkPerimeterIpAddressesTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkPerimeterIpAddressesTypeEnumStringValues Enumerates the set of values in String for NetworkPerimeterIpAddressesTypeEnum
func GetNetworkPerimeterIpAddressesTypeEnumStringValues() []string {
	return []string{
		"CIDR",
		"RANGE",
		"EXACT",
	}
}

// GetMappingNetworkPerimeterIpAddressesTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkPerimeterIpAddressesTypeEnum(val string) (NetworkPerimeterIpAddressesTypeEnum, bool) {
	enum, ok := mappingNetworkPerimeterIpAddressesTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// NetworkPerimeterIpAddressesVersionEnum Enum with underlying type: string
type NetworkPerimeterIpAddressesVersionEnum string

// Set of constants representing the allowable values for NetworkPerimeterIpAddressesVersionEnum
const (
	NetworkPerimeterIpAddressesVersionIpv4 NetworkPerimeterIpAddressesVersionEnum = "IPV4"
	NetworkPerimeterIpAddressesVersionIpv6 NetworkPerimeterIpAddressesVersionEnum = "IPV6"
)

var mappingNetworkPerimeterIpAddressesVersionEnum = map[string]NetworkPerimeterIpAddressesVersionEnum{
	"IPV4": NetworkPerimeterIpAddressesVersionIpv4,
	"IPV6": NetworkPerimeterIpAddressesVersionIpv6,
}

var mappingNetworkPerimeterIpAddressesVersionEnumLowerCase = map[string]NetworkPerimeterIpAddressesVersionEnum{
	"ipv4": NetworkPerimeterIpAddressesVersionIpv4,
	"ipv6": NetworkPerimeterIpAddressesVersionIpv6,
}

// GetNetworkPerimeterIpAddressesVersionEnumValues Enumerates the set of values for NetworkPerimeterIpAddressesVersionEnum
func GetNetworkPerimeterIpAddressesVersionEnumValues() []NetworkPerimeterIpAddressesVersionEnum {
	values := make([]NetworkPerimeterIpAddressesVersionEnum, 0)
	for _, v := range mappingNetworkPerimeterIpAddressesVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkPerimeterIpAddressesVersionEnumStringValues Enumerates the set of values in String for NetworkPerimeterIpAddressesVersionEnum
func GetNetworkPerimeterIpAddressesVersionEnumStringValues() []string {
	return []string{
		"IPV4",
		"IPV6",
	}
}

// GetMappingNetworkPerimeterIpAddressesVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkPerimeterIpAddressesVersionEnum(val string) (NetworkPerimeterIpAddressesVersionEnum, bool) {
	enum, ok := mappingNetworkPerimeterIpAddressesVersionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
