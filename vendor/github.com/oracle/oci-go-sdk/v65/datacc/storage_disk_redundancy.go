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

// StorageDiskRedundancyEnum Enum with underlying type: string
type StorageDiskRedundancyEnum string

// Set of constants representing the allowable values for StorageDiskRedundancyEnum
const (
	StorageDiskRedundancyHigh   StorageDiskRedundancyEnum = "HIGH"
	StorageDiskRedundancyMirror StorageDiskRedundancyEnum = "MIRROR"
	StorageDiskRedundancyFlex   StorageDiskRedundancyEnum = "FLEX"
)

var mappingStorageDiskRedundancyEnum = map[string]StorageDiskRedundancyEnum{
	"HIGH":   StorageDiskRedundancyHigh,
	"MIRROR": StorageDiskRedundancyMirror,
	"FLEX":   StorageDiskRedundancyFlex,
}

var mappingStorageDiskRedundancyEnumLowerCase = map[string]StorageDiskRedundancyEnum{
	"high":   StorageDiskRedundancyHigh,
	"mirror": StorageDiskRedundancyMirror,
	"flex":   StorageDiskRedundancyFlex,
}

// GetStorageDiskRedundancyEnumValues Enumerates the set of values for StorageDiskRedundancyEnum
func GetStorageDiskRedundancyEnumValues() []StorageDiskRedundancyEnum {
	values := make([]StorageDiskRedundancyEnum, 0)
	for _, v := range mappingStorageDiskRedundancyEnum {
		values = append(values, v)
	}
	return values
}

// GetStorageDiskRedundancyEnumStringValues Enumerates the set of values in String for StorageDiskRedundancyEnum
func GetStorageDiskRedundancyEnumStringValues() []string {
	return []string{
		"HIGH",
		"MIRROR",
		"FLEX",
	}
}

// GetMappingStorageDiskRedundancyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStorageDiskRedundancyEnum(val string) (StorageDiskRedundancyEnum, bool) {
	enum, ok := mappingStorageDiskRedundancyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
