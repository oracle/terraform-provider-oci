// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Common set of Object Storage and Archive Storage APIs for managing buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.cloud.oracle.com/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"strings"
)

// ArchivalStateEnum Enum with underlying type: string
type ArchivalStateEnum string

// Set of constants representing the allowable values for ArchivalStateEnum
const (
	ArchivalStateArchived  ArchivalStateEnum = "Archived"
	ArchivalStateRestoring ArchivalStateEnum = "Restoring"
	ArchivalStateRestored  ArchivalStateEnum = "Restored"
)

var mappingArchivalStateEnum = map[string]ArchivalStateEnum{
	"Archived":  ArchivalStateArchived,
	"Restoring": ArchivalStateRestoring,
	"Restored":  ArchivalStateRestored,
}

// GetArchivalStateEnumValues Enumerates the set of values for ArchivalStateEnum
func GetArchivalStateEnumValues() []ArchivalStateEnum {
	values := make([]ArchivalStateEnum, 0)
	for _, v := range mappingArchivalStateEnum {
		values = append(values, v)
	}
	return values
}

// GetArchivalStateEnumStringValues Enumerates the set of values in String for ArchivalStateEnum
func GetArchivalStateEnumStringValues() []string {
	return []string{
		"Archived",
		"Restoring",
		"Restored",
	}
}

// GetMappingArchivalStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingArchivalStateEnum(val string) (ArchivalStateEnum, bool) {
	mappingArchivalStateEnumIgnoreCase := make(map[string]ArchivalStateEnum)
	for k, v := range mappingArchivalStateEnum {
		mappingArchivalStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingArchivalStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
