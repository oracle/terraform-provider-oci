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

// ArchivalStateEnum Enum with underlying type: string
type ArchivalStateEnum string

// Set of constants representing the allowable values for ArchivalStateEnum
const (
	ArchivalStateArchived  ArchivalStateEnum = "Archived"
	ArchivalStateRestoring ArchivalStateEnum = "Restoring"
	ArchivalStateRestored  ArchivalStateEnum = "Restored"
)

var mappingArchivalState = map[string]ArchivalStateEnum{
	"Archived":  ArchivalStateArchived,
	"Restoring": ArchivalStateRestoring,
	"Restored":  ArchivalStateRestored,
}

// GetArchivalStateEnumValues Enumerates the set of values for ArchivalStateEnum
func GetArchivalStateEnumValues() []ArchivalStateEnum {
	values := make([]ArchivalStateEnum, 0)
	for _, v := range mappingArchivalState {
		values = append(values, v)
	}
	return values
}
