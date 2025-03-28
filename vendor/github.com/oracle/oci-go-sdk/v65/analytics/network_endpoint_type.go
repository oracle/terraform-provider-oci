// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"strings"
)

// NetworkEndpointTypeEnum Enum with underlying type: string
type NetworkEndpointTypeEnum string

// Set of constants representing the allowable values for NetworkEndpointTypeEnum
const (
	NetworkEndpointTypePublic  NetworkEndpointTypeEnum = "PUBLIC"
	NetworkEndpointTypePrivate NetworkEndpointTypeEnum = "PRIVATE"
)

var mappingNetworkEndpointTypeEnum = map[string]NetworkEndpointTypeEnum{
	"PUBLIC":  NetworkEndpointTypePublic,
	"PRIVATE": NetworkEndpointTypePrivate,
}

var mappingNetworkEndpointTypeEnumLowerCase = map[string]NetworkEndpointTypeEnum{
	"public":  NetworkEndpointTypePublic,
	"private": NetworkEndpointTypePrivate,
}

// GetNetworkEndpointTypeEnumValues Enumerates the set of values for NetworkEndpointTypeEnum
func GetNetworkEndpointTypeEnumValues() []NetworkEndpointTypeEnum {
	values := make([]NetworkEndpointTypeEnum, 0)
	for _, v := range mappingNetworkEndpointTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkEndpointTypeEnumStringValues Enumerates the set of values in String for NetworkEndpointTypeEnum
func GetNetworkEndpointTypeEnumStringValues() []string {
	return []string{
		"PUBLIC",
		"PRIVATE",
	}
}

// GetMappingNetworkEndpointTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkEndpointTypeEnum(val string) (NetworkEndpointTypeEnum, bool) {
	enum, ok := mappingNetworkEndpointTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
