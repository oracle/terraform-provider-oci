// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// DbSystemOptions The DB system options.
type DbSystemOptions struct {

	// The storage option used in DB system.
	// ASM - Automatic storage management
	// LVM - Logical Volume management
	StorageManagement DbSystemOptionsStorageManagementEnum `mandatory:"false" json:"storageManagement,omitempty"`
}

func (m DbSystemOptions) String() string {
	return common.PointerString(m)
}

// DbSystemOptionsStorageManagementEnum Enum with underlying type: string
type DbSystemOptionsStorageManagementEnum string

// Set of constants representing the allowable values for DbSystemOptionsStorageManagementEnum
const (
	DbSystemOptionsStorageManagementAsm DbSystemOptionsStorageManagementEnum = "ASM"
	DbSystemOptionsStorageManagementLvm DbSystemOptionsStorageManagementEnum = "LVM"
)

var mappingDbSystemOptionsStorageManagement = map[string]DbSystemOptionsStorageManagementEnum{
	"ASM": DbSystemOptionsStorageManagementAsm,
	"LVM": DbSystemOptionsStorageManagementLvm,
}

// GetDbSystemOptionsStorageManagementEnumValues Enumerates the set of values for DbSystemOptionsStorageManagementEnum
func GetDbSystemOptionsStorageManagementEnumValues() []DbSystemOptionsStorageManagementEnum {
	values := make([]DbSystemOptionsStorageManagementEnum, 0)
	for _, v := range mappingDbSystemOptionsStorageManagement {
		values = append(values, v)
	}
	return values
}
