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

	// The storage option used in DB system. You can specify either Automatic Storage Management (ASM) (https://www.oracle.com/pls/topic/lookup?ctx=en/database/oracle/oracle-database/19&id=OSTMG-GUID-BC612D35-5399-4A35-843E-CF76E3D3CDB5) or Logical Volume Manager (LVM) (https://www.oracle.com/pls/topic/lookup?ctx=en/database/oracle/oracle-database/19&id=ADMIN-GUID-57C50259-9472-4ED0-8818-DB9ABA96EC8E).
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
