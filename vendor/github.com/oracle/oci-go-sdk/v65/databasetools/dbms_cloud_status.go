// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"strings"
)

// DbmsCloudStatusEnum Enum with underlying type: string
type DbmsCloudStatusEnum string

// Set of constants representing the allowable values for DbmsCloudStatusEnum
const (
	DbmsCloudStatusAvailable   DbmsCloudStatusEnum = "AVAILABLE"
	DbmsCloudStatusUnavailable DbmsCloudStatusEnum = "UNAVAILABLE"
)

var mappingDbmsCloudStatusEnum = map[string]DbmsCloudStatusEnum{
	"AVAILABLE":   DbmsCloudStatusAvailable,
	"UNAVAILABLE": DbmsCloudStatusUnavailable,
}

var mappingDbmsCloudStatusEnumLowerCase = map[string]DbmsCloudStatusEnum{
	"available":   DbmsCloudStatusAvailable,
	"unavailable": DbmsCloudStatusUnavailable,
}

// GetDbmsCloudStatusEnumValues Enumerates the set of values for DbmsCloudStatusEnum
func GetDbmsCloudStatusEnumValues() []DbmsCloudStatusEnum {
	values := make([]DbmsCloudStatusEnum, 0)
	for _, v := range mappingDbmsCloudStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDbmsCloudStatusEnumStringValues Enumerates the set of values in String for DbmsCloudStatusEnum
func GetDbmsCloudStatusEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"UNAVAILABLE",
	}
}

// GetMappingDbmsCloudStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbmsCloudStatusEnum(val string) (DbmsCloudStatusEnum, bool) {
	enum, ok := mappingDbmsCloudStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
