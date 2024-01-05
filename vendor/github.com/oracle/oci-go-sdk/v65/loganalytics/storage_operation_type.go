// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	StorageOperationTypePurgeArchivalData          StorageOperationTypeEnum = "PURGE_ARCHIVAL_DATA"
	StorageOperationTypeArchiveStorageData         StorageOperationTypeEnum = "ARCHIVE_STORAGE_DATA"
	StorageOperationTypeCleanupArchivalStorageData StorageOperationTypeEnum = "CLEANUP_ARCHIVAL_STORAGE_DATA"
	StorageOperationTypeEncryptActiveData          StorageOperationTypeEnum = "ENCRYPT_ACTIVE_DATA"
	StorageOperationTypeEncryptArchivalData        StorageOperationTypeEnum = "ENCRYPT_ARCHIVAL_DATA"
)

var mappingStorageOperationTypeEnum = map[string]StorageOperationTypeEnum{
	"OFFBOARD_TENANCY":              StorageOperationTypeOffboardTenancy,
	"PURGE_STORAGE_DATA":            StorageOperationTypePurgeStorageData,
	"RECALL_ARCHIVED_STORAGE_DATA":  StorageOperationTypeRecallArchivedStorageData,
	"RELEASE_RECALLED_STORAGE_DATA": StorageOperationTypeReleaseRecalledStorageData,
	"PURGE_ARCHIVAL_DATA":           StorageOperationTypePurgeArchivalData,
	"ARCHIVE_STORAGE_DATA":          StorageOperationTypeArchiveStorageData,
	"CLEANUP_ARCHIVAL_STORAGE_DATA": StorageOperationTypeCleanupArchivalStorageData,
	"ENCRYPT_ACTIVE_DATA":           StorageOperationTypeEncryptActiveData,
	"ENCRYPT_ARCHIVAL_DATA":         StorageOperationTypeEncryptArchivalData,
}

var mappingStorageOperationTypeEnumLowerCase = map[string]StorageOperationTypeEnum{
	"offboard_tenancy":              StorageOperationTypeOffboardTenancy,
	"purge_storage_data":            StorageOperationTypePurgeStorageData,
	"recall_archived_storage_data":  StorageOperationTypeRecallArchivedStorageData,
	"release_recalled_storage_data": StorageOperationTypeReleaseRecalledStorageData,
	"purge_archival_data":           StorageOperationTypePurgeArchivalData,
	"archive_storage_data":          StorageOperationTypeArchiveStorageData,
	"cleanup_archival_storage_data": StorageOperationTypeCleanupArchivalStorageData,
	"encrypt_active_data":           StorageOperationTypeEncryptActiveData,
	"encrypt_archival_data":         StorageOperationTypeEncryptArchivalData,
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
		"PURGE_ARCHIVAL_DATA",
		"ARCHIVE_STORAGE_DATA",
		"CLEANUP_ARCHIVAL_STORAGE_DATA",
		"ENCRYPT_ACTIVE_DATA",
		"ENCRYPT_ARCHIVAL_DATA",
	}
}

// GetMappingStorageOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStorageOperationTypeEnum(val string) (StorageOperationTypeEnum, bool) {
	enum, ok := mappingStorageOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
