// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"strings"
)

// ManagedListTypeEnum Enum with underlying type: string
type ManagedListTypeEnum string

// Set of constants representing the allowable values for ManagedListTypeEnum
const (
	ManagedListTypeCidrBlock    ManagedListTypeEnum = "CIDR_BLOCK"
	ManagedListTypeUsers        ManagedListTypeEnum = "USERS"
	ManagedListTypeGroups       ManagedListTypeEnum = "GROUPS"
	ManagedListTypeIpv4Address  ManagedListTypeEnum = "IPV4ADDRESS"
	ManagedListTypeIpv6Address  ManagedListTypeEnum = "IPV6ADDRESS"
	ManagedListTypeResourceOcid ManagedListTypeEnum = "RESOURCE_OCID"
	ManagedListTypeRegion       ManagedListTypeEnum = "REGION"
	ManagedListTypeCountry      ManagedListTypeEnum = "COUNTRY"
	ManagedListTypeState        ManagedListTypeEnum = "STATE"
	ManagedListTypeCity         ManagedListTypeEnum = "CITY"
	ManagedListTypeTags         ManagedListTypeEnum = "TAGS"
	ManagedListTypeGeneric      ManagedListTypeEnum = "GENERIC"
)

var mappingManagedListTypeEnum = map[string]ManagedListTypeEnum{
	"CIDR_BLOCK":    ManagedListTypeCidrBlock,
	"USERS":         ManagedListTypeUsers,
	"GROUPS":        ManagedListTypeGroups,
	"IPV4ADDRESS":   ManagedListTypeIpv4Address,
	"IPV6ADDRESS":   ManagedListTypeIpv6Address,
	"RESOURCE_OCID": ManagedListTypeResourceOcid,
	"REGION":        ManagedListTypeRegion,
	"COUNTRY":       ManagedListTypeCountry,
	"STATE":         ManagedListTypeState,
	"CITY":          ManagedListTypeCity,
	"TAGS":          ManagedListTypeTags,
	"GENERIC":       ManagedListTypeGeneric,
}

var mappingManagedListTypeEnumLowerCase = map[string]ManagedListTypeEnum{
	"cidr_block":    ManagedListTypeCidrBlock,
	"users":         ManagedListTypeUsers,
	"groups":        ManagedListTypeGroups,
	"ipv4address":   ManagedListTypeIpv4Address,
	"ipv6address":   ManagedListTypeIpv6Address,
	"resource_ocid": ManagedListTypeResourceOcid,
	"region":        ManagedListTypeRegion,
	"country":       ManagedListTypeCountry,
	"state":         ManagedListTypeState,
	"city":          ManagedListTypeCity,
	"tags":          ManagedListTypeTags,
	"generic":       ManagedListTypeGeneric,
}

// GetManagedListTypeEnumValues Enumerates the set of values for ManagedListTypeEnum
func GetManagedListTypeEnumValues() []ManagedListTypeEnum {
	values := make([]ManagedListTypeEnum, 0)
	for _, v := range mappingManagedListTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedListTypeEnumStringValues Enumerates the set of values in String for ManagedListTypeEnum
func GetManagedListTypeEnumStringValues() []string {
	return []string{
		"CIDR_BLOCK",
		"USERS",
		"GROUPS",
		"IPV4ADDRESS",
		"IPV6ADDRESS",
		"RESOURCE_OCID",
		"REGION",
		"COUNTRY",
		"STATE",
		"CITY",
		"TAGS",
		"GENERIC",
	}
}

// GetMappingManagedListTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedListTypeEnum(val string) (ManagedListTypeEnum, bool) {
	enum, ok := mappingManagedListTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
