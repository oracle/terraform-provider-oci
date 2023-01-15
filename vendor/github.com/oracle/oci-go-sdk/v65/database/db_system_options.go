// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbSystemOptions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDbSystemOptionsStorageManagementEnum(string(m.StorageManagement)); !ok && m.StorageManagement != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StorageManagement: %s. Supported values are: %s.", m.StorageManagement, strings.Join(GetDbSystemOptionsStorageManagementEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbSystemOptionsStorageManagementEnum Enum with underlying type: string
type DbSystemOptionsStorageManagementEnum string

// Set of constants representing the allowable values for DbSystemOptionsStorageManagementEnum
const (
	DbSystemOptionsStorageManagementAsm DbSystemOptionsStorageManagementEnum = "ASM"
	DbSystemOptionsStorageManagementLvm DbSystemOptionsStorageManagementEnum = "LVM"
)

var mappingDbSystemOptionsStorageManagementEnum = map[string]DbSystemOptionsStorageManagementEnum{
	"ASM": DbSystemOptionsStorageManagementAsm,
	"LVM": DbSystemOptionsStorageManagementLvm,
}

var mappingDbSystemOptionsStorageManagementEnumLowerCase = map[string]DbSystemOptionsStorageManagementEnum{
	"asm": DbSystemOptionsStorageManagementAsm,
	"lvm": DbSystemOptionsStorageManagementLvm,
}

// GetDbSystemOptionsStorageManagementEnumValues Enumerates the set of values for DbSystemOptionsStorageManagementEnum
func GetDbSystemOptionsStorageManagementEnumValues() []DbSystemOptionsStorageManagementEnum {
	values := make([]DbSystemOptionsStorageManagementEnum, 0)
	for _, v := range mappingDbSystemOptionsStorageManagementEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemOptionsStorageManagementEnumStringValues Enumerates the set of values in String for DbSystemOptionsStorageManagementEnum
func GetDbSystemOptionsStorageManagementEnumStringValues() []string {
	return []string{
		"ASM",
		"LVM",
	}
}

// GetMappingDbSystemOptionsStorageManagementEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbSystemOptionsStorageManagementEnum(val string) (DbSystemOptionsStorageManagementEnum, bool) {
	enum, ok := mappingDbSystemOptionsStorageManagementEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
