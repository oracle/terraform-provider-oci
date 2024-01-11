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

// AutonomousDatabaseBackup An Autonomous Database backup.
type AutonomousDatabaseBackup struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous Database backup.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
	AutonomousDatabaseId *string `mandatory:"true" json:"autonomousDatabaseId"`

	// The user-friendly name for the backup. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of backup.
	Type AutonomousDatabaseBackupTypeEnum `mandatory:"true" json:"type"`

	// Indicates whether the backup is user-initiated or automatic.
	IsAutomatic *bool `mandatory:"true" json:"isAutomatic"`

	// The current state of the backup.
	LifecycleState AutonomousDatabaseBackupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the backup started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the backup completed.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The size of the database in terabytes at the time the backup was taken.
	DatabaseSizeInTBs *float32 `mandatory:"false" json:"databaseSizeInTBs"`

	// Indicates whether the backup can be used to restore the associated Autonomous Database.
	IsRestorable *bool `mandatory:"false" json:"isRestorable"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
	KeyStoreId *string `mandatory:"false" json:"keyStoreId"`

	// The wallet name for Oracle Key Vault.
	KeyStoreWalletName *string `mandatory:"false" json:"keyStoreWalletName"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.cloud.oracle.com/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// Retention period, in days, for long-term backups
	RetentionPeriodInDays *int `mandatory:"false" json:"retentionPeriodInDays"`

	// Timestamp until when the backup will be available
	TimeAvailableTill *common.SDKTime `mandatory:"false" json:"timeAvailableTill"`

	// A valid Oracle Database version for Autonomous Database.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	// The backup size in terrabytes (TB).
	SizeInTBs *float64 `mandatory:"false" json:"sizeInTBs"`

	BackupDestinationDetails *BackupDestinationDetails `mandatory:"false" json:"backupDestinationDetails"`
}

func (m AutonomousDatabaseBackup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDatabaseBackup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousDatabaseBackupTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAutonomousDatabaseBackupTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseBackupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousDatabaseBackupLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousDatabaseBackupTypeEnum Enum with underlying type: string
type AutonomousDatabaseBackupTypeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseBackupTypeEnum
const (
	AutonomousDatabaseBackupTypeIncremental AutonomousDatabaseBackupTypeEnum = "INCREMENTAL"
	AutonomousDatabaseBackupTypeFull        AutonomousDatabaseBackupTypeEnum = "FULL"
	AutonomousDatabaseBackupTypeLongterm    AutonomousDatabaseBackupTypeEnum = "LONGTERM"
)

var mappingAutonomousDatabaseBackupTypeEnum = map[string]AutonomousDatabaseBackupTypeEnum{
	"INCREMENTAL": AutonomousDatabaseBackupTypeIncremental,
	"FULL":        AutonomousDatabaseBackupTypeFull,
	"LONGTERM":    AutonomousDatabaseBackupTypeLongterm,
}

var mappingAutonomousDatabaseBackupTypeEnumLowerCase = map[string]AutonomousDatabaseBackupTypeEnum{
	"incremental": AutonomousDatabaseBackupTypeIncremental,
	"full":        AutonomousDatabaseBackupTypeFull,
	"longterm":    AutonomousDatabaseBackupTypeLongterm,
}

// GetAutonomousDatabaseBackupTypeEnumValues Enumerates the set of values for AutonomousDatabaseBackupTypeEnum
func GetAutonomousDatabaseBackupTypeEnumValues() []AutonomousDatabaseBackupTypeEnum {
	values := make([]AutonomousDatabaseBackupTypeEnum, 0)
	for _, v := range mappingAutonomousDatabaseBackupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseBackupTypeEnumStringValues Enumerates the set of values in String for AutonomousDatabaseBackupTypeEnum
func GetAutonomousDatabaseBackupTypeEnumStringValues() []string {
	return []string{
		"INCREMENTAL",
		"FULL",
		"LONGTERM",
	}
}

// GetMappingAutonomousDatabaseBackupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseBackupTypeEnum(val string) (AutonomousDatabaseBackupTypeEnum, bool) {
	enum, ok := mappingAutonomousDatabaseBackupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseBackupLifecycleStateEnum Enum with underlying type: string
type AutonomousDatabaseBackupLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseBackupLifecycleStateEnum
const (
	AutonomousDatabaseBackupLifecycleStateCreating AutonomousDatabaseBackupLifecycleStateEnum = "CREATING"
	AutonomousDatabaseBackupLifecycleStateActive   AutonomousDatabaseBackupLifecycleStateEnum = "ACTIVE"
	AutonomousDatabaseBackupLifecycleStateDeleting AutonomousDatabaseBackupLifecycleStateEnum = "DELETING"
	AutonomousDatabaseBackupLifecycleStateDeleted  AutonomousDatabaseBackupLifecycleStateEnum = "DELETED"
	AutonomousDatabaseBackupLifecycleStateFailed   AutonomousDatabaseBackupLifecycleStateEnum = "FAILED"
	AutonomousDatabaseBackupLifecycleStateUpdating AutonomousDatabaseBackupLifecycleStateEnum = "UPDATING"
)

var mappingAutonomousDatabaseBackupLifecycleStateEnum = map[string]AutonomousDatabaseBackupLifecycleStateEnum{
	"CREATING": AutonomousDatabaseBackupLifecycleStateCreating,
	"ACTIVE":   AutonomousDatabaseBackupLifecycleStateActive,
	"DELETING": AutonomousDatabaseBackupLifecycleStateDeleting,
	"DELETED":  AutonomousDatabaseBackupLifecycleStateDeleted,
	"FAILED":   AutonomousDatabaseBackupLifecycleStateFailed,
	"UPDATING": AutonomousDatabaseBackupLifecycleStateUpdating,
}

var mappingAutonomousDatabaseBackupLifecycleStateEnumLowerCase = map[string]AutonomousDatabaseBackupLifecycleStateEnum{
	"creating": AutonomousDatabaseBackupLifecycleStateCreating,
	"active":   AutonomousDatabaseBackupLifecycleStateActive,
	"deleting": AutonomousDatabaseBackupLifecycleStateDeleting,
	"deleted":  AutonomousDatabaseBackupLifecycleStateDeleted,
	"failed":   AutonomousDatabaseBackupLifecycleStateFailed,
	"updating": AutonomousDatabaseBackupLifecycleStateUpdating,
}

// GetAutonomousDatabaseBackupLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseBackupLifecycleStateEnum
func GetAutonomousDatabaseBackupLifecycleStateEnumValues() []AutonomousDatabaseBackupLifecycleStateEnum {
	values := make([]AutonomousDatabaseBackupLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseBackupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseBackupLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousDatabaseBackupLifecycleStateEnum
func GetAutonomousDatabaseBackupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingAutonomousDatabaseBackupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseBackupLifecycleStateEnum(val string) (AutonomousDatabaseBackupLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousDatabaseBackupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
