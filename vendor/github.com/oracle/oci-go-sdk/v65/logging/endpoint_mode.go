// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"strings"
)

// EndpointModeEnum Enum with underlying type: string
type EndpointModeEnum string

// Set of constants representing the allowable values for EndpointModeEnum
const (
	EndpointModePublic          EndpointModeEnum = "PUBLIC"
	EndpointModePrivateEndpoint EndpointModeEnum = "PRIVATE_ENDPOINT"
)

var mappingEndpointModeEnum = map[string]EndpointModeEnum{
	"PUBLIC":           EndpointModePublic,
	"PRIVATE_ENDPOINT": EndpointModePrivateEndpoint,
}

var mappingEndpointModeEnumLowerCase = map[string]EndpointModeEnum{
	"public":           EndpointModePublic,
	"private_endpoint": EndpointModePrivateEndpoint,
}

// GetEndpointModeEnumValues Enumerates the set of values for EndpointModeEnum
func GetEndpointModeEnumValues() []EndpointModeEnum {
	values := make([]EndpointModeEnum, 0)
	for _, v := range mappingEndpointModeEnum {
		values = append(values, v)
	}
	return values
}

// GetEndpointModeEnumStringValues Enumerates the set of values in String for EndpointModeEnum
func GetEndpointModeEnumStringValues() []string {
	return []string{
		"PUBLIC",
		"PRIVATE_ENDPOINT",
	}
}

// GetMappingEndpointModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEndpointModeEnum(val string) (EndpointModeEnum, bool) {
	enum, ok := mappingEndpointModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
