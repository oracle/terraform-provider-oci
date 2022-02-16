// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"strings"
)

// CloudGuardStatusEnum Enum with underlying type: string
type CloudGuardStatusEnum string

// Set of constants representing the allowable values for CloudGuardStatusEnum
const (
	CloudGuardStatusEnabled  CloudGuardStatusEnum = "ENABLED"
	CloudGuardStatusDisabled CloudGuardStatusEnum = "DISABLED"
)

var mappingCloudGuardStatusEnum = map[string]CloudGuardStatusEnum{
	"ENABLED":  CloudGuardStatusEnabled,
	"DISABLED": CloudGuardStatusDisabled,
}

// GetCloudGuardStatusEnumValues Enumerates the set of values for CloudGuardStatusEnum
func GetCloudGuardStatusEnumValues() []CloudGuardStatusEnum {
	values := make([]CloudGuardStatusEnum, 0)
	for _, v := range mappingCloudGuardStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudGuardStatusEnumStringValues Enumerates the set of values in String for CloudGuardStatusEnum
func GetCloudGuardStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingCloudGuardStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudGuardStatusEnum(val string) (CloudGuardStatusEnum, bool) {
	mappingCloudGuardStatusEnumIgnoreCase := make(map[string]CloudGuardStatusEnum)
	for k, v := range mappingCloudGuardStatusEnum {
		mappingCloudGuardStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingCloudGuardStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
