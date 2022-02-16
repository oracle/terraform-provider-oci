// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"strings"
)

// StorageOperationTypeEnum Enum with underlying type: string
type StorageOperationTypeEnum string

// Set of constants representing the allowable values for StorageOperationTypeEnum
const (
	StorageOperationTypeOffboardTenancy            StorageOperationTypeEnum = "OFFBOARD_TENANCY"
	StorageOperationTypePurgeStorageData           StorageOperationTypeEnum = "PURGE_STORAGE_DATA"
	StorageOperationTypeRecallArchivedStorageData  StorageOperationTypeEnum = "RECALL_ARCHIVED_STORAGE_DATA"
	StorageOperationTypeReleaseRecalledStorageData StorageOperationTypeEnum = "RELEASE_RECALLED_STORAGE_DATA"
	StorageOperationTypeArchiveStorageData         StorageOperationTypeEnum = "ARCHIVE_STORAGE_DATA"
	StorageOperationTypeCleanupArchivalStorageData StorageOperationTypeEnum = "CLEANUP_ARCHIVAL_STORAGE_DATA"
)

var mappingStorageOperationTypeEnum = map[string]StorageOperationTypeEnum{
	"OFFBOARD_TENANCY":              StorageOperationTypeOffboardTenancy,
	"PURGE_STORAGE_DATA":            StorageOperationTypePurgeStorageData,
	"RECALL_ARCHIVED_STORAGE_DATA":  StorageOperationTypeRecallArchivedStorageData,
	"RELEASE_RECALLED_STORAGE_DATA": StorageOperationTypeReleaseRecalledStorageData,
	"ARCHIVE_STORAGE_DATA":          StorageOperationTypeArchiveStorageData,
	"CLEANUP_ARCHIVAL_STORAGE_DATA": StorageOperationTypeCleanupArchivalStorageData,
}

// GetStorageOperationTypeEnumValues Enumerates the set of values for StorageOperationTypeEnum
func GetStorageOperationTypeEnumValues() []StorageOperationTypeEnum {
	values := make([]StorageOperationTypeEnum, 0)
	for _, v := range mappingStorageOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetStorageOperationTypeEnumStringValues Enumerates the set of values in String for StorageOperationTypeEnum
func GetStorageOperationTypeEnumStringValues() []string {
	return []string{
		"OFFBOARD_TENANCY",
		"PURGE_STORAGE_DATA",
		"RECALL_ARCHIVED_STORAGE_DATA",
		"RELEASE_RECALLED_STORAGE_DATA",
		"ARCHIVE_STORAGE_DATA",
		"CLEANUP_ARCHIVAL_STORAGE_DATA",
	}
}

// GetMappingStorageOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStorageOperationTypeEnum(val string) (StorageOperationTypeEnum, bool) {
	mappingStorageOperationTypeEnumIgnoreCase := make(map[string]StorageOperationTypeEnum)
	for k, v := range mappingStorageOperationTypeEnum {
		mappingStorageOperationTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingStorageOperationTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
