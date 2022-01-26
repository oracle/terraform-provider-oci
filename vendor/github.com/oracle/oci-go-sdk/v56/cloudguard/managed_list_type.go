// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

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

var mappingManagedListType = map[string]ManagedListTypeEnum{
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

// GetManagedListTypeEnumValues Enumerates the set of values for ManagedListTypeEnum
func GetManagedListTypeEnumValues() []ManagedListTypeEnum {
	values := make([]ManagedListTypeEnum, 0)
	for _, v := range mappingManagedListType {
		values = append(values, v)
	}
	return values
}
