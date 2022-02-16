// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Database Tools APIs to manage Connections and Private Endpoints.
//

package databasetools

import (
	"strings"
)

// ServiceCapabilityEnum Enum with underlying type: string
type ServiceCapabilityEnum string

// Set of constants representing the allowable values for ServiceCapabilityEnum
const (
	ServiceCapabilityPrivateEndpointSupported ServiceCapabilityEnum = "PRIVATE_ENDPOINT_SUPPORTED"
)

var mappingServiceCapabilityEnum = map[string]ServiceCapabilityEnum{
	"PRIVATE_ENDPOINT_SUPPORTED": ServiceCapabilityPrivateEndpointSupported,
}

// GetServiceCapabilityEnumValues Enumerates the set of values for ServiceCapabilityEnum
func GetServiceCapabilityEnumValues() []ServiceCapabilityEnum {
	values := make([]ServiceCapabilityEnum, 0)
	for _, v := range mappingServiceCapabilityEnum {
		values = append(values, v)
	}
	return values
}

// GetServiceCapabilityEnumStringValues Enumerates the set of values in String for ServiceCapabilityEnum
func GetServiceCapabilityEnumStringValues() []string {
	return []string{
		"PRIVATE_ENDPOINT_SUPPORTED",
	}
}

// GetMappingServiceCapabilityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServiceCapabilityEnum(val string) (ServiceCapabilityEnum, bool) {
	mappingServiceCapabilityEnumIgnoreCase := make(map[string]ServiceCapabilityEnum)
	for k, v := range mappingServiceCapabilityEnum {
		mappingServiceCapabilityEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingServiceCapabilityEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
