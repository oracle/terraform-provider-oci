// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OlvmStorageDomainProperties OLVM Storage Domain properties.
type OlvmStorageDomainProperties struct {

	// A human-readable name in plain text.
	StorageDomainName *string `mandatory:"true" json:"storageDomainName"`

	// A human-readable description in plain text.
	StorageDomainDescription *string `mandatory:"false" json:"storageDomainDescription"`

	// Free text containing comments about this object.
	Comment *string `mandatory:"false" json:"comment"`

	// Space available in bytes.
	AvailableSpaceInBytes *int64 `mandatory:"false" json:"availableSpaceInBytes"`

	// Space used in bytes.
	UsedSpaceInBytes *int64 `mandatory:"false" json:"usedSpaceInBytes"`

	// Space committed in bytes.
	CommittedSpaceInBytes *int64 `mandatory:"false" json:"committedSpaceInBytes"`

	// Block size in bytes.
	BlockSizeInBytes *int64 `mandatory:"false" json:"blockSizeInBytes"`

	// Whether a data storage domain is used as backup domain or not.
	IsBackup *bool `mandatory:"false" json:"isBackup"`

	// Indicates if this is the primary (master) storage domain of a data center.
	IsPrimary *bool `mandatory:"false" json:"isPrimary"`

	// Whether this storage domain is imported.
	IsImport *bool `mandatory:"false" json:"isImport"`

	// Indicates whether disks' blocks on block storage domains will be discarded right before they are deleted.
	IsDiscardAfterDelete *bool `mandatory:"false" json:"isDiscardAfterDelete"`

	// Indicates whether a block storage domain supports discard operations
	IsSupportDiscard *bool `mandatory:"false" json:"isSupportDiscard"`

	// Indicates whether a block storage domain supports the property that discard zeroes the data.
	IsSupportDiscardZeroesData *bool `mandatory:"false" json:"isSupportDiscardZeroesData"`

	// Serves as the default value of wipe_after_delete for disks on this storage domain.
	IsWipeAfterDelete *bool `mandatory:"false" json:"isWipeAfterDelete"`

	// Ensure storage domain always has at least this amount of unoccupied space in GBs.
	CriticalSpaceActionBlockerInGBs *int `mandatory:"false" json:"criticalSpaceActionBlockerInGBs"`

	// If the free space available on the storage domain is below this percentage, warning messages are displayed to the user and logged.
	WarningLowSpaceIndicatorInPercentage *int `mandatory:"false" json:"warningLowSpaceIndicatorInPercentage"`

	// Status of storage domain.
	ExternalStatus OlvmStorageDomainPropertiesExternalStatusEnum `mandatory:"false" json:"externalStatus,omitempty"`

	// Status of storage domain.
	StorageDomainStatus OlvmStorageDomainPropertiesStorageDomainStatusEnum `mandatory:"false" json:"storageDomainStatus,omitempty"`

	Storage *Storage `mandatory:"false" json:"storage"`

	// Type which represents a format of storage domain.
	StorageFormat OlvmStorageDomainPropertiesStorageFormatEnum `mandatory:"false" json:"storageFormat,omitempty"`

	// Indicates the kind of data managed by a storage domain.
	StorageDomainType OlvmStorageDomainPropertiesStorageDomainTypeEnum `mandatory:"false" json:"storageDomainType,omitempty"`

	// List of data centers where storage domain belongs
	DataCenters []OlvmDataCenter `mandatory:"false" json:"dataCenters"`
}

func (m OlvmStorageDomainProperties) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmStorageDomainProperties) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmStorageDomainPropertiesExternalStatusEnum(string(m.ExternalStatus)); !ok && m.ExternalStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExternalStatus: %s. Supported values are: %s.", m.ExternalStatus, strings.Join(GetOlvmStorageDomainPropertiesExternalStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOlvmStorageDomainPropertiesStorageDomainStatusEnum(string(m.StorageDomainStatus)); !ok && m.StorageDomainStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StorageDomainStatus: %s. Supported values are: %s.", m.StorageDomainStatus, strings.Join(GetOlvmStorageDomainPropertiesStorageDomainStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOlvmStorageDomainPropertiesStorageFormatEnum(string(m.StorageFormat)); !ok && m.StorageFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StorageFormat: %s. Supported values are: %s.", m.StorageFormat, strings.Join(GetOlvmStorageDomainPropertiesStorageFormatEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOlvmStorageDomainPropertiesStorageDomainTypeEnum(string(m.StorageDomainType)); !ok && m.StorageDomainType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StorageDomainType: %s. Supported values are: %s.", m.StorageDomainType, strings.Join(GetOlvmStorageDomainPropertiesStorageDomainTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmStorageDomainPropertiesExternalStatusEnum Enum with underlying type: string
type OlvmStorageDomainPropertiesExternalStatusEnum string

// Set of constants representing the allowable values for OlvmStorageDomainPropertiesExternalStatusEnum
const (
	OlvmStorageDomainPropertiesExternalStatusError   OlvmStorageDomainPropertiesExternalStatusEnum = "ERROR"
	OlvmStorageDomainPropertiesExternalStatusFailure OlvmStorageDomainPropertiesExternalStatusEnum = "FAILURE"
	OlvmStorageDomainPropertiesExternalStatusInfo    OlvmStorageDomainPropertiesExternalStatusEnum = "INFO"
	OlvmStorageDomainPropertiesExternalStatusOk      OlvmStorageDomainPropertiesExternalStatusEnum = "OK"
	OlvmStorageDomainPropertiesExternalStatusWarning OlvmStorageDomainPropertiesExternalStatusEnum = "WARNING"
)

var mappingOlvmStorageDomainPropertiesExternalStatusEnum = map[string]OlvmStorageDomainPropertiesExternalStatusEnum{
	"ERROR":   OlvmStorageDomainPropertiesExternalStatusError,
	"FAILURE": OlvmStorageDomainPropertiesExternalStatusFailure,
	"INFO":    OlvmStorageDomainPropertiesExternalStatusInfo,
	"OK":      OlvmStorageDomainPropertiesExternalStatusOk,
	"WARNING": OlvmStorageDomainPropertiesExternalStatusWarning,
}

var mappingOlvmStorageDomainPropertiesExternalStatusEnumLowerCase = map[string]OlvmStorageDomainPropertiesExternalStatusEnum{
	"error":   OlvmStorageDomainPropertiesExternalStatusError,
	"failure": OlvmStorageDomainPropertiesExternalStatusFailure,
	"info":    OlvmStorageDomainPropertiesExternalStatusInfo,
	"ok":      OlvmStorageDomainPropertiesExternalStatusOk,
	"warning": OlvmStorageDomainPropertiesExternalStatusWarning,
}

// GetOlvmStorageDomainPropertiesExternalStatusEnumValues Enumerates the set of values for OlvmStorageDomainPropertiesExternalStatusEnum
func GetOlvmStorageDomainPropertiesExternalStatusEnumValues() []OlvmStorageDomainPropertiesExternalStatusEnum {
	values := make([]OlvmStorageDomainPropertiesExternalStatusEnum, 0)
	for _, v := range mappingOlvmStorageDomainPropertiesExternalStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmStorageDomainPropertiesExternalStatusEnumStringValues Enumerates the set of values in String for OlvmStorageDomainPropertiesExternalStatusEnum
func GetOlvmStorageDomainPropertiesExternalStatusEnumStringValues() []string {
	return []string{
		"ERROR",
		"FAILURE",
		"INFO",
		"OK",
		"WARNING",
	}
}

// GetMappingOlvmStorageDomainPropertiesExternalStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmStorageDomainPropertiesExternalStatusEnum(val string) (OlvmStorageDomainPropertiesExternalStatusEnum, bool) {
	enum, ok := mappingOlvmStorageDomainPropertiesExternalStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OlvmStorageDomainPropertiesStorageDomainStatusEnum Enum with underlying type: string
type OlvmStorageDomainPropertiesStorageDomainStatusEnum string

// Set of constants representing the allowable values for OlvmStorageDomainPropertiesStorageDomainStatusEnum
const (
	OlvmStorageDomainPropertiesStorageDomainStatusActivating              OlvmStorageDomainPropertiesStorageDomainStatusEnum = "ACTIVATING"
	OlvmStorageDomainPropertiesStorageDomainStatusActive                  OlvmStorageDomainPropertiesStorageDomainStatusEnum = "ACTIVE"
	OlvmStorageDomainPropertiesStorageDomainStatusDetaching               OlvmStorageDomainPropertiesStorageDomainStatusEnum = "DETACHING"
	OlvmStorageDomainPropertiesStorageDomainStatusInactive                OlvmStorageDomainPropertiesStorageDomainStatusEnum = "INACTIVE"
	OlvmStorageDomainPropertiesStorageDomainStatusLocked                  OlvmStorageDomainPropertiesStorageDomainStatusEnum = "LOCKED"
	OlvmStorageDomainPropertiesStorageDomainStatusMaintenance             OlvmStorageDomainPropertiesStorageDomainStatusEnum = "MAINTENANCE"
	OlvmStorageDomainPropertiesStorageDomainStatusMixed                   OlvmStorageDomainPropertiesStorageDomainStatusEnum = "MIXED"
	OlvmStorageDomainPropertiesStorageDomainStatusPreparingForMaintenance OlvmStorageDomainPropertiesStorageDomainStatusEnum = "PREPARING_FOR_MAINTENANCE"
	OlvmStorageDomainPropertiesStorageDomainStatusUnattached              OlvmStorageDomainPropertiesStorageDomainStatusEnum = "UNATTACHED"
	OlvmStorageDomainPropertiesStorageDomainStatusUnknown                 OlvmStorageDomainPropertiesStorageDomainStatusEnum = "UNKNOWN"
)

var mappingOlvmStorageDomainPropertiesStorageDomainStatusEnum = map[string]OlvmStorageDomainPropertiesStorageDomainStatusEnum{
	"ACTIVATING":                OlvmStorageDomainPropertiesStorageDomainStatusActivating,
	"ACTIVE":                    OlvmStorageDomainPropertiesStorageDomainStatusActive,
	"DETACHING":                 OlvmStorageDomainPropertiesStorageDomainStatusDetaching,
	"INACTIVE":                  OlvmStorageDomainPropertiesStorageDomainStatusInactive,
	"LOCKED":                    OlvmStorageDomainPropertiesStorageDomainStatusLocked,
	"MAINTENANCE":               OlvmStorageDomainPropertiesStorageDomainStatusMaintenance,
	"MIXED":                     OlvmStorageDomainPropertiesStorageDomainStatusMixed,
	"PREPARING_FOR_MAINTENANCE": OlvmStorageDomainPropertiesStorageDomainStatusPreparingForMaintenance,
	"UNATTACHED":                OlvmStorageDomainPropertiesStorageDomainStatusUnattached,
	"UNKNOWN":                   OlvmStorageDomainPropertiesStorageDomainStatusUnknown,
}

var mappingOlvmStorageDomainPropertiesStorageDomainStatusEnumLowerCase = map[string]OlvmStorageDomainPropertiesStorageDomainStatusEnum{
	"activating":                OlvmStorageDomainPropertiesStorageDomainStatusActivating,
	"active":                    OlvmStorageDomainPropertiesStorageDomainStatusActive,
	"detaching":                 OlvmStorageDomainPropertiesStorageDomainStatusDetaching,
	"inactive":                  OlvmStorageDomainPropertiesStorageDomainStatusInactive,
	"locked":                    OlvmStorageDomainPropertiesStorageDomainStatusLocked,
	"maintenance":               OlvmStorageDomainPropertiesStorageDomainStatusMaintenance,
	"mixed":                     OlvmStorageDomainPropertiesStorageDomainStatusMixed,
	"preparing_for_maintenance": OlvmStorageDomainPropertiesStorageDomainStatusPreparingForMaintenance,
	"unattached":                OlvmStorageDomainPropertiesStorageDomainStatusUnattached,
	"unknown":                   OlvmStorageDomainPropertiesStorageDomainStatusUnknown,
}

// GetOlvmStorageDomainPropertiesStorageDomainStatusEnumValues Enumerates the set of values for OlvmStorageDomainPropertiesStorageDomainStatusEnum
func GetOlvmStorageDomainPropertiesStorageDomainStatusEnumValues() []OlvmStorageDomainPropertiesStorageDomainStatusEnum {
	values := make([]OlvmStorageDomainPropertiesStorageDomainStatusEnum, 0)
	for _, v := range mappingOlvmStorageDomainPropertiesStorageDomainStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmStorageDomainPropertiesStorageDomainStatusEnumStringValues Enumerates the set of values in String for OlvmStorageDomainPropertiesStorageDomainStatusEnum
func GetOlvmStorageDomainPropertiesStorageDomainStatusEnumStringValues() []string {
	return []string{
		"ACTIVATING",
		"ACTIVE",
		"DETACHING",
		"INACTIVE",
		"LOCKED",
		"MAINTENANCE",
		"MIXED",
		"PREPARING_FOR_MAINTENANCE",
		"UNATTACHED",
		"UNKNOWN",
	}
}

// GetMappingOlvmStorageDomainPropertiesStorageDomainStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmStorageDomainPropertiesStorageDomainStatusEnum(val string) (OlvmStorageDomainPropertiesStorageDomainStatusEnum, bool) {
	enum, ok := mappingOlvmStorageDomainPropertiesStorageDomainStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OlvmStorageDomainPropertiesStorageFormatEnum Enum with underlying type: string
type OlvmStorageDomainPropertiesStorageFormatEnum string

// Set of constants representing the allowable values for OlvmStorageDomainPropertiesStorageFormatEnum
const (
	OlvmStorageDomainPropertiesStorageFormatV1 OlvmStorageDomainPropertiesStorageFormatEnum = "V1"
	OlvmStorageDomainPropertiesStorageFormatV2 OlvmStorageDomainPropertiesStorageFormatEnum = "V2"
	OlvmStorageDomainPropertiesStorageFormatV3 OlvmStorageDomainPropertiesStorageFormatEnum = "V3"
	OlvmStorageDomainPropertiesStorageFormatV4 OlvmStorageDomainPropertiesStorageFormatEnum = "V4"
	OlvmStorageDomainPropertiesStorageFormatV5 OlvmStorageDomainPropertiesStorageFormatEnum = "V5"
)

var mappingOlvmStorageDomainPropertiesStorageFormatEnum = map[string]OlvmStorageDomainPropertiesStorageFormatEnum{
	"V1": OlvmStorageDomainPropertiesStorageFormatV1,
	"V2": OlvmStorageDomainPropertiesStorageFormatV2,
	"V3": OlvmStorageDomainPropertiesStorageFormatV3,
	"V4": OlvmStorageDomainPropertiesStorageFormatV4,
	"V5": OlvmStorageDomainPropertiesStorageFormatV5,
}

var mappingOlvmStorageDomainPropertiesStorageFormatEnumLowerCase = map[string]OlvmStorageDomainPropertiesStorageFormatEnum{
	"v1": OlvmStorageDomainPropertiesStorageFormatV1,
	"v2": OlvmStorageDomainPropertiesStorageFormatV2,
	"v3": OlvmStorageDomainPropertiesStorageFormatV3,
	"v4": OlvmStorageDomainPropertiesStorageFormatV4,
	"v5": OlvmStorageDomainPropertiesStorageFormatV5,
}

// GetOlvmStorageDomainPropertiesStorageFormatEnumValues Enumerates the set of values for OlvmStorageDomainPropertiesStorageFormatEnum
func GetOlvmStorageDomainPropertiesStorageFormatEnumValues() []OlvmStorageDomainPropertiesStorageFormatEnum {
	values := make([]OlvmStorageDomainPropertiesStorageFormatEnum, 0)
	for _, v := range mappingOlvmStorageDomainPropertiesStorageFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmStorageDomainPropertiesStorageFormatEnumStringValues Enumerates the set of values in String for OlvmStorageDomainPropertiesStorageFormatEnum
func GetOlvmStorageDomainPropertiesStorageFormatEnumStringValues() []string {
	return []string{
		"V1",
		"V2",
		"V3",
		"V4",
		"V5",
	}
}

// GetMappingOlvmStorageDomainPropertiesStorageFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmStorageDomainPropertiesStorageFormatEnum(val string) (OlvmStorageDomainPropertiesStorageFormatEnum, bool) {
	enum, ok := mappingOlvmStorageDomainPropertiesStorageFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OlvmStorageDomainPropertiesStorageDomainTypeEnum Enum with underlying type: string
type OlvmStorageDomainPropertiesStorageDomainTypeEnum string

// Set of constants representing the allowable values for OlvmStorageDomainPropertiesStorageDomainTypeEnum
const (
	OlvmStorageDomainPropertiesStorageDomainTypeData                OlvmStorageDomainPropertiesStorageDomainTypeEnum = "DATA"
	OlvmStorageDomainPropertiesStorageDomainTypeExport              OlvmStorageDomainPropertiesStorageDomainTypeEnum = "EXPORT"
	OlvmStorageDomainPropertiesStorageDomainTypeImage               OlvmStorageDomainPropertiesStorageDomainTypeEnum = "IMAGE"
	OlvmStorageDomainPropertiesStorageDomainTypeIso                 OlvmStorageDomainPropertiesStorageDomainTypeEnum = "ISO"
	OlvmStorageDomainPropertiesStorageDomainTypeManagedBlockStorage OlvmStorageDomainPropertiesStorageDomainTypeEnum = "MANAGED_BLOCK_STORAGE"
	OlvmStorageDomainPropertiesStorageDomainTypeVolume              OlvmStorageDomainPropertiesStorageDomainTypeEnum = "VOLUME"
)

var mappingOlvmStorageDomainPropertiesStorageDomainTypeEnum = map[string]OlvmStorageDomainPropertiesStorageDomainTypeEnum{
	"DATA":                  OlvmStorageDomainPropertiesStorageDomainTypeData,
	"EXPORT":                OlvmStorageDomainPropertiesStorageDomainTypeExport,
	"IMAGE":                 OlvmStorageDomainPropertiesStorageDomainTypeImage,
	"ISO":                   OlvmStorageDomainPropertiesStorageDomainTypeIso,
	"MANAGED_BLOCK_STORAGE": OlvmStorageDomainPropertiesStorageDomainTypeManagedBlockStorage,
	"VOLUME":                OlvmStorageDomainPropertiesStorageDomainTypeVolume,
}

var mappingOlvmStorageDomainPropertiesStorageDomainTypeEnumLowerCase = map[string]OlvmStorageDomainPropertiesStorageDomainTypeEnum{
	"data":                  OlvmStorageDomainPropertiesStorageDomainTypeData,
	"export":                OlvmStorageDomainPropertiesStorageDomainTypeExport,
	"image":                 OlvmStorageDomainPropertiesStorageDomainTypeImage,
	"iso":                   OlvmStorageDomainPropertiesStorageDomainTypeIso,
	"managed_block_storage": OlvmStorageDomainPropertiesStorageDomainTypeManagedBlockStorage,
	"volume":                OlvmStorageDomainPropertiesStorageDomainTypeVolume,
}

// GetOlvmStorageDomainPropertiesStorageDomainTypeEnumValues Enumerates the set of values for OlvmStorageDomainPropertiesStorageDomainTypeEnum
func GetOlvmStorageDomainPropertiesStorageDomainTypeEnumValues() []OlvmStorageDomainPropertiesStorageDomainTypeEnum {
	values := make([]OlvmStorageDomainPropertiesStorageDomainTypeEnum, 0)
	for _, v := range mappingOlvmStorageDomainPropertiesStorageDomainTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmStorageDomainPropertiesStorageDomainTypeEnumStringValues Enumerates the set of values in String for OlvmStorageDomainPropertiesStorageDomainTypeEnum
func GetOlvmStorageDomainPropertiesStorageDomainTypeEnumStringValues() []string {
	return []string{
		"DATA",
		"EXPORT",
		"IMAGE",
		"ISO",
		"MANAGED_BLOCK_STORAGE",
		"VOLUME",
	}
}

// GetMappingOlvmStorageDomainPropertiesStorageDomainTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmStorageDomainPropertiesStorageDomainTypeEnum(val string) (OlvmStorageDomainPropertiesStorageDomainTypeEnum, bool) {
	enum, ok := mappingOlvmStorageDomainPropertiesStorageDomainTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
