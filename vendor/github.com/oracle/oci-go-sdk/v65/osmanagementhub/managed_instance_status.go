// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// ManagedInstanceStatusEnum Enum with underlying type: string
type ManagedInstanceStatusEnum string

// Set of constants representing the allowable values for ManagedInstanceStatusEnum
const (
	ManagedInstanceStatusNormal            ManagedInstanceStatusEnum = "NORMAL"
	ManagedInstanceStatusUnreachable       ManagedInstanceStatusEnum = "UNREACHABLE"
	ManagedInstanceStatusError             ManagedInstanceStatusEnum = "ERROR"
	ManagedInstanceStatusWarning           ManagedInstanceStatusEnum = "WARNING"
	ManagedInstanceStatusRegistrationError ManagedInstanceStatusEnum = "REGISTRATION_ERROR"
)

var mappingManagedInstanceStatusEnum = map[string]ManagedInstanceStatusEnum{
	"NORMAL":             ManagedInstanceStatusNormal,
	"UNREACHABLE":        ManagedInstanceStatusUnreachable,
	"ERROR":              ManagedInstanceStatusError,
	"WARNING":            ManagedInstanceStatusWarning,
	"REGISTRATION_ERROR": ManagedInstanceStatusRegistrationError,
}

var mappingManagedInstanceStatusEnumLowerCase = map[string]ManagedInstanceStatusEnum{
	"normal":             ManagedInstanceStatusNormal,
	"unreachable":        ManagedInstanceStatusUnreachable,
	"error":              ManagedInstanceStatusError,
	"warning":            ManagedInstanceStatusWarning,
	"registration_error": ManagedInstanceStatusRegistrationError,
}

// GetManagedInstanceStatusEnumValues Enumerates the set of values for ManagedInstanceStatusEnum
func GetManagedInstanceStatusEnumValues() []ManagedInstanceStatusEnum {
	values := make([]ManagedInstanceStatusEnum, 0)
	for _, v := range mappingManagedInstanceStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedInstanceStatusEnumStringValues Enumerates the set of values in String for ManagedInstanceStatusEnum
func GetManagedInstanceStatusEnumStringValues() []string {
	return []string{
		"NORMAL",
		"UNREACHABLE",
		"ERROR",
		"WARNING",
		"REGISTRATION_ERROR",
	}
}

// GetMappingManagedInstanceStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedInstanceStatusEnum(val string) (ManagedInstanceStatusEnum, bool) {
	enum, ok := mappingManagedInstanceStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
