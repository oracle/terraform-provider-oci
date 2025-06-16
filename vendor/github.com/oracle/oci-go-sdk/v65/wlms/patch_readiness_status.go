// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// WebLogic Management Service API
//
// WebLogic Management Service is an OCI service that enables a unified view and management of WebLogic domains
// in Oracle Cloud Infrastructure. Features include on-demand patching of WebLogic domains, rollback of the
// last applied patch, discovery and management of WebLogic instances on a compute host.
//

package wlms

import (
	"strings"
)

// PatchReadinessStatusEnum Enum with underlying type: string
type PatchReadinessStatusEnum string

// Set of constants representing the allowable values for PatchReadinessStatusEnum
const (
	PatchReadinessStatusOk      PatchReadinessStatusEnum = "OK"
	PatchReadinessStatusWarning PatchReadinessStatusEnum = "WARNING"
	PatchReadinessStatusError   PatchReadinessStatusEnum = "ERROR"
	PatchReadinessStatusUnknown PatchReadinessStatusEnum = "UNKNOWN"
)

var mappingPatchReadinessStatusEnum = map[string]PatchReadinessStatusEnum{
	"OK":      PatchReadinessStatusOk,
	"WARNING": PatchReadinessStatusWarning,
	"ERROR":   PatchReadinessStatusError,
	"UNKNOWN": PatchReadinessStatusUnknown,
}

var mappingPatchReadinessStatusEnumLowerCase = map[string]PatchReadinessStatusEnum{
	"ok":      PatchReadinessStatusOk,
	"warning": PatchReadinessStatusWarning,
	"error":   PatchReadinessStatusError,
	"unknown": PatchReadinessStatusUnknown,
}

// GetPatchReadinessStatusEnumValues Enumerates the set of values for PatchReadinessStatusEnum
func GetPatchReadinessStatusEnumValues() []PatchReadinessStatusEnum {
	values := make([]PatchReadinessStatusEnum, 0)
	for _, v := range mappingPatchReadinessStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchReadinessStatusEnumStringValues Enumerates the set of values in String for PatchReadinessStatusEnum
func GetPatchReadinessStatusEnumStringValues() []string {
	return []string{
		"OK",
		"WARNING",
		"ERROR",
		"UNKNOWN",
	}
}

// GetMappingPatchReadinessStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchReadinessStatusEnum(val string) (PatchReadinessStatusEnum, bool) {
	enum, ok := mappingPatchReadinessStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
