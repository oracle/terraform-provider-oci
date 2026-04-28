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

// StorageTypeEnum Enum with underlying type: string
type StorageTypeEnum string

// Set of constants representing the allowable values for StorageTypeEnum
const (
	StorageTypeHighCapacity StorageTypeEnum = "HIGH_CAPACITY"
	StorageTypeExtended     StorageTypeEnum = "EXTENDED"
	StorageTypeExtremeFlash StorageTypeEnum = "EXTREME_FLASH"
)

var mappingStorageTypeEnum = map[string]StorageTypeEnum{
	"HIGH_CAPACITY": StorageTypeHighCapacity,
	"EXTENDED":      StorageTypeExtended,
	"EXTREME_FLASH": StorageTypeExtremeFlash,
}

var mappingStorageTypeEnumLowerCase = map[string]StorageTypeEnum{
	"high_capacity": StorageTypeHighCapacity,
	"extended":      StorageTypeExtended,
	"extreme_flash": StorageTypeExtremeFlash,
}

// GetStorageTypeEnumValues Enumerates the set of values for StorageTypeEnum
func GetStorageTypeEnumValues() []StorageTypeEnum {
	values := make([]StorageTypeEnum, 0)
	for _, v := range mappingStorageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetStorageTypeEnumStringValues Enumerates the set of values in String for StorageTypeEnum
func GetStorageTypeEnumStringValues() []string {
	return []string{
		"HIGH_CAPACITY",
		"EXTENDED",
		"EXTREME_FLASH",
	}
}

// GetMappingStorageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStorageTypeEnum(val string) (StorageTypeEnum, bool) {
	enum, ok := mappingStorageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
