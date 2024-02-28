// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// Backup The representation of Backup
type Backup struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the backup.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	// The user-friendly name for the backup. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The type of backup.
	Type BackupTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The date and time the backup started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the backup was completed.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The name of the availability domain where the database backup is stored.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The current state of the backup.
	LifecycleState BackupLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The Oracle Database edition of the DB system from which the database backup was taken.
	DatabaseEdition BackupDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`

	// The size of the database in gigabytes at the time the backup was taken.
	DatabaseSizeInGBs *float64 `mandatory:"false" json:"databaseSizeInGBs"`

	// Shape of the backup's source database.
	Shape *string `mandatory:"false" json:"shape"`

	// Version of the backup's source database
	Version *string `mandatory:"false" json:"version"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.cloud.oracle.com/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
	KeyStoreId *string `mandatory:"false" json:"keyStoreId"`

	// The wallet name for Oracle Key Vault.
	KeyStoreWalletName *string `mandatory:"false" json:"keyStoreWalletName"`
}

func (m Backup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Backup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBackupTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetBackupTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBackupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBackupLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBackupDatabaseEditionEnum(string(m.DatabaseEdition)); !ok && m.DatabaseEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseEdition: %s. Supported values are: %s.", m.DatabaseEdition, strings.Join(GetBackupDatabaseEditionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BackupTypeEnum Enum with underlying type: string
type BackupTypeEnum string

// Set of constants representing the allowable values for BackupTypeEnum
const (
	BackupTypeIncremental BackupTypeEnum = "INCREMENTAL"
	BackupTypeFull        BackupTypeEnum = "FULL"
	BackupTypeVirtualFull BackupTypeEnum = "VIRTUAL_FULL"
)

var mappingBackupTypeEnum = map[string]BackupTypeEnum{
	"INCREMENTAL":  BackupTypeIncremental,
	"FULL":         BackupTypeFull,
	"VIRTUAL_FULL": BackupTypeVirtualFull,
}

var mappingBackupTypeEnumLowerCase = map[string]BackupTypeEnum{
	"incremental":  BackupTypeIncremental,
	"full":         BackupTypeFull,
	"virtual_full": BackupTypeVirtualFull,
}

// GetBackupTypeEnumValues Enumerates the set of values for BackupTypeEnum
func GetBackupTypeEnumValues() []BackupTypeEnum {
	values := make([]BackupTypeEnum, 0)
	for _, v := range mappingBackupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupTypeEnumStringValues Enumerates the set of values in String for BackupTypeEnum
func GetBackupTypeEnumStringValues() []string {
	return []string{
		"INCREMENTAL",
		"FULL",
		"VIRTUAL_FULL",
	}
}

// GetMappingBackupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupTypeEnum(val string) (BackupTypeEnum, bool) {
	enum, ok := mappingBackupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BackupLifecycleStateEnum Enum with underlying type: string
type BackupLifecycleStateEnum string

// Set of constants representing the allowable values for BackupLifecycleStateEnum
const (
	BackupLifecycleStateCreating  BackupLifecycleStateEnum = "CREATING"
	BackupLifecycleStateActive    BackupLifecycleStateEnum = "ACTIVE"
	BackupLifecycleStateDeleting  BackupLifecycleStateEnum = "DELETING"
	BackupLifecycleStateDeleted   BackupLifecycleStateEnum = "DELETED"
	BackupLifecycleStateFailed    BackupLifecycleStateEnum = "FAILED"
	BackupLifecycleStateRestoring BackupLifecycleStateEnum = "RESTORING"
	BackupLifecycleStateCanceling BackupLifecycleStateEnum = "CANCELING"
	BackupLifecycleStateCanceled  BackupLifecycleStateEnum = "CANCELED"
)

var mappingBackupLifecycleStateEnum = map[string]BackupLifecycleStateEnum{
	"CREATING":  BackupLifecycleStateCreating,
	"ACTIVE":    BackupLifecycleStateActive,
	"DELETING":  BackupLifecycleStateDeleting,
	"DELETED":   BackupLifecycleStateDeleted,
	"FAILED":    BackupLifecycleStateFailed,
	"RESTORING": BackupLifecycleStateRestoring,
	"CANCELING": BackupLifecycleStateCanceling,
	"CANCELED":  BackupLifecycleStateCanceled,
}

var mappingBackupLifecycleStateEnumLowerCase = map[string]BackupLifecycleStateEnum{
	"creating":  BackupLifecycleStateCreating,
	"active":    BackupLifecycleStateActive,
	"deleting":  BackupLifecycleStateDeleting,
	"deleted":   BackupLifecycleStateDeleted,
	"failed":    BackupLifecycleStateFailed,
	"restoring": BackupLifecycleStateRestoring,
	"canceling": BackupLifecycleStateCanceling,
	"canceled":  BackupLifecycleStateCanceled,
}

// GetBackupLifecycleStateEnumValues Enumerates the set of values for BackupLifecycleStateEnum
func GetBackupLifecycleStateEnumValues() []BackupLifecycleStateEnum {
	values := make([]BackupLifecycleStateEnum, 0)
	for _, v := range mappingBackupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupLifecycleStateEnumStringValues Enumerates the set of values in String for BackupLifecycleStateEnum
func GetBackupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"RESTORING",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingBackupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupLifecycleStateEnum(val string) (BackupLifecycleStateEnum, bool) {
	enum, ok := mappingBackupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BackupDatabaseEditionEnum Enum with underlying type: string
type BackupDatabaseEditionEnum string

// Set of constants representing the allowable values for BackupDatabaseEditionEnum
const (
	BackupDatabaseEditionStandardEdition                     BackupDatabaseEditionEnum = "STANDARD_EDITION"
	BackupDatabaseEditionEnterpriseEdition                   BackupDatabaseEditionEnum = "ENTERPRISE_EDITION"
	BackupDatabaseEditionEnterpriseEditionHighPerformance    BackupDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	BackupDatabaseEditionEnterpriseEditionExtremePerformance BackupDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
)

var mappingBackupDatabaseEditionEnum = map[string]BackupDatabaseEditionEnum{
	"STANDARD_EDITION":                       BackupDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     BackupDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    BackupDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": BackupDatabaseEditionEnterpriseEditionExtremePerformance,
}

var mappingBackupDatabaseEditionEnumLowerCase = map[string]BackupDatabaseEditionEnum{
	"standard_edition":                       BackupDatabaseEditionStandardEdition,
	"enterprise_edition":                     BackupDatabaseEditionEnterpriseEdition,
	"enterprise_edition_high_performance":    BackupDatabaseEditionEnterpriseEditionHighPerformance,
	"enterprise_edition_extreme_performance": BackupDatabaseEditionEnterpriseEditionExtremePerformance,
}

// GetBackupDatabaseEditionEnumValues Enumerates the set of values for BackupDatabaseEditionEnum
func GetBackupDatabaseEditionEnumValues() []BackupDatabaseEditionEnum {
	values := make([]BackupDatabaseEditionEnum, 0)
	for _, v := range mappingBackupDatabaseEditionEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupDatabaseEditionEnumStringValues Enumerates the set of values in String for BackupDatabaseEditionEnum
func GetBackupDatabaseEditionEnumStringValues() []string {
	return []string{
		"STANDARD_EDITION",
		"ENTERPRISE_EDITION",
		"ENTERPRISE_EDITION_HIGH_PERFORMANCE",
		"ENTERPRISE_EDITION_EXTREME_PERFORMANCE",
	}
}

// GetMappingBackupDatabaseEditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupDatabaseEditionEnum(val string) (BackupDatabaseEditionEnum, bool) {
	enum, ok := mappingBackupDatabaseEditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
