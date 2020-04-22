// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
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
