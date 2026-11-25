// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"strings"
)

// WarehouseStorageTypeEnum Enum with underlying type: string
type WarehouseStorageTypeEnum string

// Set of constants representing the allowable values for WarehouseStorageTypeEnum
const (
	WarehouseStorageTypeOciObjectStorageAccessKey WarehouseStorageTypeEnum = "OCI_OBJECT_STORAGE_ACCESS_KEY"
	WarehouseStorageTypeZfs                       WarehouseStorageTypeEnum = "ZFS"
	WarehouseStorageTypeAzureStorage              WarehouseStorageTypeEnum = "AZURE_STORAGE"
)

var mappingWarehouseStorageTypeEnum = map[string]WarehouseStorageTypeEnum{
	"OCI_OBJECT_STORAGE_ACCESS_KEY": WarehouseStorageTypeOciObjectStorageAccessKey,
	"ZFS":                           WarehouseStorageTypeZfs,
	"AZURE_STORAGE":                 WarehouseStorageTypeAzureStorage,
}

var mappingWarehouseStorageTypeEnumLowerCase = map[string]WarehouseStorageTypeEnum{
	"oci_object_storage_access_key": WarehouseStorageTypeOciObjectStorageAccessKey,
	"zfs":                           WarehouseStorageTypeZfs,
	"azure_storage":                 WarehouseStorageTypeAzureStorage,
}

// GetWarehouseStorageTypeEnumValues Enumerates the set of values for WarehouseStorageTypeEnum
func GetWarehouseStorageTypeEnumValues() []WarehouseStorageTypeEnum {
	values := make([]WarehouseStorageTypeEnum, 0)
	for _, v := range mappingWarehouseStorageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWarehouseStorageTypeEnumStringValues Enumerates the set of values in String for WarehouseStorageTypeEnum
func GetWarehouseStorageTypeEnumStringValues() []string {
	return []string{
		"OCI_OBJECT_STORAGE_ACCESS_KEY",
		"ZFS",
		"AZURE_STORAGE",
	}
}

// GetMappingWarehouseStorageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWarehouseStorageTypeEnum(val string) (WarehouseStorageTypeEnum, bool) {
	enum, ok := mappingWarehouseStorageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
