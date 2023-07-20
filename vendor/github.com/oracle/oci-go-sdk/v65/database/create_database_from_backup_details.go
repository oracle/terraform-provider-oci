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

// CreateDatabaseFromBackupDetails The representation of CreateDatabaseFromBackupDetails
type CreateDatabaseFromBackupDetails struct {

	// The backup OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	BackupId *string `mandatory:"true" json:"backupId"`

	// A strong password for SYS, SYSTEM, PDB Admin and TDE Wallet. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
	AdminPassword *string `mandatory:"true" json:"adminPassword"`

	// The password to open the TDE wallet.
	BackupTDEPassword *string `mandatory:"false" json:"backupTDEPassword"`

	// The `DB_UNIQUE_NAME` of the Oracle Database being backed up.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// The display name of the database to be created from the backup. It must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted.
	DbName *string `mandatory:"false" json:"dbName"`

	// Specifies a prefix for the `Oracle SID` of the database to be created.
	SidPrefix *string `mandatory:"false" json:"sidPrefix"`

	// The list of pluggable databases that needs to be restored into new database.
	PluggableDatabases []string `mandatory:"false" json:"pluggableDatabases"`

	// Database Storage Type, this option is applicable when database on Exadata VM cluster on Exascale Infrastructure. High Capacity will be selected if not specified.
	VaultStorageType CreateDatabaseFromBackupDetailsVaultStorageTypeEnum `mandatory:"false" json:"vaultStorageType,omitempty"`
}

func (m CreateDatabaseFromBackupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDatabaseFromBackupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateDatabaseFromBackupDetailsVaultStorageTypeEnum(string(m.VaultStorageType)); !ok && m.VaultStorageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VaultStorageType: %s. Supported values are: %s.", m.VaultStorageType, strings.Join(GetCreateDatabaseFromBackupDetailsVaultStorageTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDatabaseFromBackupDetailsVaultStorageTypeEnum Enum with underlying type: string
type CreateDatabaseFromBackupDetailsVaultStorageTypeEnum string

// Set of constants representing the allowable values for CreateDatabaseFromBackupDetailsVaultStorageTypeEnum
const (
	CreateDatabaseFromBackupDetailsVaultStorageTypeHighCapacity  CreateDatabaseFromBackupDetailsVaultStorageTypeEnum = "HIGH_CAPACITY"
	CreateDatabaseFromBackupDetailsVaultStorageTypeExteremeFlash CreateDatabaseFromBackupDetailsVaultStorageTypeEnum = "EXTEREME_FLASH"
)

var mappingCreateDatabaseFromBackupDetailsVaultStorageTypeEnum = map[string]CreateDatabaseFromBackupDetailsVaultStorageTypeEnum{
	"HIGH_CAPACITY":  CreateDatabaseFromBackupDetailsVaultStorageTypeHighCapacity,
	"EXTEREME_FLASH": CreateDatabaseFromBackupDetailsVaultStorageTypeExteremeFlash,
}

var mappingCreateDatabaseFromBackupDetailsVaultStorageTypeEnumLowerCase = map[string]CreateDatabaseFromBackupDetailsVaultStorageTypeEnum{
	"high_capacity":  CreateDatabaseFromBackupDetailsVaultStorageTypeHighCapacity,
	"extereme_flash": CreateDatabaseFromBackupDetailsVaultStorageTypeExteremeFlash,
}

// GetCreateDatabaseFromBackupDetailsVaultStorageTypeEnumValues Enumerates the set of values for CreateDatabaseFromBackupDetailsVaultStorageTypeEnum
func GetCreateDatabaseFromBackupDetailsVaultStorageTypeEnumValues() []CreateDatabaseFromBackupDetailsVaultStorageTypeEnum {
	values := make([]CreateDatabaseFromBackupDetailsVaultStorageTypeEnum, 0)
	for _, v := range mappingCreateDatabaseFromBackupDetailsVaultStorageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDatabaseFromBackupDetailsVaultStorageTypeEnumStringValues Enumerates the set of values in String for CreateDatabaseFromBackupDetailsVaultStorageTypeEnum
func GetCreateDatabaseFromBackupDetailsVaultStorageTypeEnumStringValues() []string {
	return []string{
		"HIGH_CAPACITY",
		"EXTEREME_FLASH",
	}
}

// GetMappingCreateDatabaseFromBackupDetailsVaultStorageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDatabaseFromBackupDetailsVaultStorageTypeEnum(val string) (CreateDatabaseFromBackupDetailsVaultStorageTypeEnum, bool) {
	enum, ok := mappingCreateDatabaseFromBackupDetailsVaultStorageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
