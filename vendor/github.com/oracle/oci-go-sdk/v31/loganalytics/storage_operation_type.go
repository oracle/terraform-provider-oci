// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

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

var mappingStorageOperationType = map[string]StorageOperationTypeEnum{
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
	for _, v := range mappingStorageOperationType {
		values = append(values, v)
	}
	return values
}
