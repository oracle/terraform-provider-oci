// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"strings"
)

// PatchingStatusEnumEnum Enum with underlying type: string
type PatchingStatusEnumEnum string

// Set of constants representing the allowable values for PatchingStatusEnumEnum
const (
	PatchingStatusEnumCreating  PatchingStatusEnumEnum = "CREATING"
	PatchingStatusEnumPatching  PatchingStatusEnumEnum = "PATCHING"
	PatchingStatusEnumWaiting   PatchingStatusEnumEnum = "WAITING"
	PatchingStatusEnumScheduled PatchingStatusEnumEnum = "SCHEDULED"
	PatchingStatusEnumComplete  PatchingStatusEnumEnum = "COMPLETE"
	PatchingStatusEnumFailed    PatchingStatusEnumEnum = "FAILED"
)

var mappingPatchingStatusEnumEnum = map[string]PatchingStatusEnumEnum{
	"CREATING":  PatchingStatusEnumCreating,
	"PATCHING":  PatchingStatusEnumPatching,
	"WAITING":   PatchingStatusEnumWaiting,
	"SCHEDULED": PatchingStatusEnumScheduled,
	"COMPLETE":  PatchingStatusEnumComplete,
	"FAILED":    PatchingStatusEnumFailed,
}

var mappingPatchingStatusEnumEnumLowerCase = map[string]PatchingStatusEnumEnum{
	"creating":  PatchingStatusEnumCreating,
	"patching":  PatchingStatusEnumPatching,
	"waiting":   PatchingStatusEnumWaiting,
	"scheduled": PatchingStatusEnumScheduled,
	"complete":  PatchingStatusEnumComplete,
	"failed":    PatchingStatusEnumFailed,
}

// GetPatchingStatusEnumEnumValues Enumerates the set of values for PatchingStatusEnumEnum
func GetPatchingStatusEnumEnumValues() []PatchingStatusEnumEnum {
	values := make([]PatchingStatusEnumEnum, 0)
	for _, v := range mappingPatchingStatusEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchingStatusEnumEnumStringValues Enumerates the set of values in String for PatchingStatusEnumEnum
func GetPatchingStatusEnumEnumStringValues() []string {
	return []string{
		"CREATING",
		"PATCHING",
		"WAITING",
		"SCHEDULED",
		"COMPLETE",
		"FAILED",
	}
}

// GetMappingPatchingStatusEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchingStatusEnumEnum(val string) (PatchingStatusEnumEnum, bool) {
	enum, ok := mappingPatchingStatusEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
