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

// Storage Host storage information
type Storage struct {

	// A unique identifier.
	Id *string `mandatory:"false" json:"id"`

	// Address of storage domain.
	Address *string `mandatory:"false" json:"address"`

	// Free text containing comments about this object.
	Comment *string `mandatory:"false" json:"comment"`

	// A human-readable description in plain text.
	StorageDescription *string `mandatory:"false" json:"storageDescription"`

	// The options to be passed when creating a storage domain using a cinder driver.
	DriverOptions []OlvmProperty `mandatory:"false" json:"driverOptions"`

	// The options containing sensitive information to be passed when creating a storage domain using a cinder driver.
	DriverSensitiveOptions []OlvmProperty `mandatory:"false" json:"driverSensitiveOptions"`

	// Logical Units of the host storage
	LogicalUnits []LogicalUnit `mandatory:"false" json:"logicalUnits"`

	// Mount options
	MountOptions *string `mandatory:"false" json:"mountOptions"`

	// The number of times to retry a request before attempting further recovery actions.
	NfsRetrans *int `mandatory:"false" json:"nfsRetrans"`

	// The time in tenths of a second to wait for a response before retrying NFS requests.
	NfsTimeo *int `mandatory:"false" json:"nfsTimeo"`

	// Version of NFS used.
	NfsVersion StorageNfsVersionEnum `mandatory:"false" json:"nfsVersion,omitempty"`

	// Whether to override LUNs
	IsOverrideLuns *bool `mandatory:"false" json:"isOverrideLuns"`

	// Password of the host storage.
	Password *string `mandatory:"false" json:"password"`

	// Paths of the host storage.
	Paths *int `mandatory:"false" json:"paths"`

	// Port of the host storage.
	Port *int `mandatory:"false" json:"port"`

	// Portal of the host storage.
	Portal *string `mandatory:"false" json:"portal"`

	// Target of the host storage.
	Target *string `mandatory:"false" json:"target"`

	// Username of the host storage.
	Username *string `mandatory:"false" json:"username"`

	// Type representing a storage domain type.
	Type StorageTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Type of VFS
	VfsType *string `mandatory:"false" json:"vfsType"`

	VolumeGroup *VolumeGroup `mandatory:"false" json:"volumeGroup"`
}

func (m Storage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Storage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingStorageNfsVersionEnum(string(m.NfsVersion)); !ok && m.NfsVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NfsVersion: %s. Supported values are: %s.", m.NfsVersion, strings.Join(GetStorageNfsVersionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingStorageTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetStorageTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StorageNfsVersionEnum Enum with underlying type: string
type StorageNfsVersionEnum string

// Set of constants representing the allowable values for StorageNfsVersionEnum
const (
	StorageNfsVersionAuto StorageNfsVersionEnum = "AUTO"
	StorageNfsVersionV3   StorageNfsVersionEnum = "V3"
	StorageNfsVersionV4   StorageNfsVersionEnum = "V4"
	StorageNfsVersionV40  StorageNfsVersionEnum = "V4_0"
	StorageNfsVersionV41  StorageNfsVersionEnum = "V4_1"
	StorageNfsVersionV42  StorageNfsVersionEnum = "V4_2"
)

var mappingStorageNfsVersionEnum = map[string]StorageNfsVersionEnum{
	"AUTO": StorageNfsVersionAuto,
	"V3":   StorageNfsVersionV3,
	"V4":   StorageNfsVersionV4,
	"V4_0": StorageNfsVersionV40,
	"V4_1": StorageNfsVersionV41,
	"V4_2": StorageNfsVersionV42,
}

var mappingStorageNfsVersionEnumLowerCase = map[string]StorageNfsVersionEnum{
	"auto": StorageNfsVersionAuto,
	"v3":   StorageNfsVersionV3,
	"v4":   StorageNfsVersionV4,
	"v4_0": StorageNfsVersionV40,
	"v4_1": StorageNfsVersionV41,
	"v4_2": StorageNfsVersionV42,
}

// GetStorageNfsVersionEnumValues Enumerates the set of values for StorageNfsVersionEnum
func GetStorageNfsVersionEnumValues() []StorageNfsVersionEnum {
	values := make([]StorageNfsVersionEnum, 0)
	for _, v := range mappingStorageNfsVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetStorageNfsVersionEnumStringValues Enumerates the set of values in String for StorageNfsVersionEnum
func GetStorageNfsVersionEnumStringValues() []string {
	return []string{
		"AUTO",
		"V3",
		"V4",
		"V4_0",
		"V4_1",
		"V4_2",
	}
}

// GetMappingStorageNfsVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStorageNfsVersionEnum(val string) (StorageNfsVersionEnum, bool) {
	enum, ok := mappingStorageNfsVersionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// StorageTypeEnum Enum with underlying type: string
type StorageTypeEnum string

// Set of constants representing the allowable values for StorageTypeEnum
const (
	StorageTypeCinder              StorageTypeEnum = "CINDER"
	StorageTypeFcp                 StorageTypeEnum = "FCP"
	StorageTypeGlance              StorageTypeEnum = "GLANCE"
	StorageTypeGlusterfs           StorageTypeEnum = "GLUSTERFS"
	StorageTypeIscsi               StorageTypeEnum = "ISCSI"
	StorageTypeLocalfs             StorageTypeEnum = "LOCALFS"
	StorageTypeManagedBlockStorage StorageTypeEnum = "MANAGED_BLOCK_STORAGE"
	StorageTypeNfs                 StorageTypeEnum = "NFS"
	StorageTypePosixfs             StorageTypeEnum = "POSIXFS"
)

var mappingStorageTypeEnum = map[string]StorageTypeEnum{
	"CINDER":                StorageTypeCinder,
	"FCP":                   StorageTypeFcp,
	"GLANCE":                StorageTypeGlance,
	"GLUSTERFS":             StorageTypeGlusterfs,
	"ISCSI":                 StorageTypeIscsi,
	"LOCALFS":               StorageTypeLocalfs,
	"MANAGED_BLOCK_STORAGE": StorageTypeManagedBlockStorage,
	"NFS":                   StorageTypeNfs,
	"POSIXFS":               StorageTypePosixfs,
}

var mappingStorageTypeEnumLowerCase = map[string]StorageTypeEnum{
	"cinder":                StorageTypeCinder,
	"fcp":                   StorageTypeFcp,
	"glance":                StorageTypeGlance,
	"glusterfs":             StorageTypeGlusterfs,
	"iscsi":                 StorageTypeIscsi,
	"localfs":               StorageTypeLocalfs,
	"managed_block_storage": StorageTypeManagedBlockStorage,
	"nfs":                   StorageTypeNfs,
	"posixfs":               StorageTypePosixfs,
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
		"CINDER",
		"FCP",
		"GLANCE",
		"GLUSTERFS",
		"ISCSI",
		"LOCALFS",
		"MANAGED_BLOCK_STORAGE",
		"NFS",
		"POSIXFS",
	}
}

// GetMappingStorageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStorageTypeEnum(val string) (StorageTypeEnum, bool) {
	enum, ok := mappingStorageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
