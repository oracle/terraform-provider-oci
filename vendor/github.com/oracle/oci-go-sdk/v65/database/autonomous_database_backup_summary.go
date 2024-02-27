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

// AutonomousDatabaseBackupSummary An Autonomous Database backup.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type AutonomousDatabaseBackupSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous Database backup.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
	AutonomousDatabaseId *string `mandatory:"true" json:"autonomousDatabaseId"`

	// The user-friendly name for the backup. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of backup.
	Type AutonomousDatabaseBackupSummaryTypeEnum `mandatory:"true" json:"type"`

	// Indicates whether the backup is user-initiated or automatic.
	IsAutomatic *bool `mandatory:"true" json:"isAutomatic"`

	// The current state of the backup.
	LifecycleState AutonomousDatabaseBackupSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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

func (m AutonomousDatabaseBackupSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDatabaseBackupSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousDatabaseBackupSummaryTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAutonomousDatabaseBackupSummaryTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseBackupSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousDatabaseBackupSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousDatabaseBackupSummaryTypeEnum Enum with underlying type: string
type AutonomousDatabaseBackupSummaryTypeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseBackupSummaryTypeEnum
const (
	AutonomousDatabaseBackupSummaryTypeIncremental AutonomousDatabaseBackupSummaryTypeEnum = "INCREMENTAL"
	AutonomousDatabaseBackupSummaryTypeFull        AutonomousDatabaseBackupSummaryTypeEnum = "FULL"
	AutonomousDatabaseBackupSummaryTypeLongterm    AutonomousDatabaseBackupSummaryTypeEnum = "LONGTERM"
)

var mappingAutonomousDatabaseBackupSummaryTypeEnum = map[string]AutonomousDatabaseBackupSummaryTypeEnum{
	"INCREMENTAL": AutonomousDatabaseBackupSummaryTypeIncremental,
	"FULL":        AutonomousDatabaseBackupSummaryTypeFull,
	"LONGTERM":    AutonomousDatabaseBackupSummaryTypeLongterm,
}

var mappingAutonomousDatabaseBackupSummaryTypeEnumLowerCase = map[string]AutonomousDatabaseBackupSummaryTypeEnum{
	"incremental": AutonomousDatabaseBackupSummaryTypeIncremental,
	"full":        AutonomousDatabaseBackupSummaryTypeFull,
	"longterm":    AutonomousDatabaseBackupSummaryTypeLongterm,
}

// GetAutonomousDatabaseBackupSummaryTypeEnumValues Enumerates the set of values for AutonomousDatabaseBackupSummaryTypeEnum
func GetAutonomousDatabaseBackupSummaryTypeEnumValues() []AutonomousDatabaseBackupSummaryTypeEnum {
	values := make([]AutonomousDatabaseBackupSummaryTypeEnum, 0)
	for _, v := range mappingAutonomousDatabaseBackupSummaryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseBackupSummaryTypeEnumStringValues Enumerates the set of values in String for AutonomousDatabaseBackupSummaryTypeEnum
func GetAutonomousDatabaseBackupSummaryTypeEnumStringValues() []string {
	return []string{
		"INCREMENTAL",
		"FULL",
		"LONGTERM",
	}
}

// GetMappingAutonomousDatabaseBackupSummaryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseBackupSummaryTypeEnum(val string) (AutonomousDatabaseBackupSummaryTypeEnum, bool) {
	enum, ok := mappingAutonomousDatabaseBackupSummaryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseBackupSummaryLifecycleStateEnum Enum with underlying type: string
type AutonomousDatabaseBackupSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseBackupSummaryLifecycleStateEnum
const (
	AutonomousDatabaseBackupSummaryLifecycleStateCreating AutonomousDatabaseBackupSummaryLifecycleStateEnum = "CREATING"
	AutonomousDatabaseBackupSummaryLifecycleStateActive   AutonomousDatabaseBackupSummaryLifecycleStateEnum = "ACTIVE"
	AutonomousDatabaseBackupSummaryLifecycleStateDeleting AutonomousDatabaseBackupSummaryLifecycleStateEnum = "DELETING"
	AutonomousDatabaseBackupSummaryLifecycleStateDeleted  AutonomousDatabaseBackupSummaryLifecycleStateEnum = "DELETED"
	AutonomousDatabaseBackupSummaryLifecycleStateFailed   AutonomousDatabaseBackupSummaryLifecycleStateEnum = "FAILED"
	AutonomousDatabaseBackupSummaryLifecycleStateUpdating AutonomousDatabaseBackupSummaryLifecycleStateEnum = "UPDATING"
)

var mappingAutonomousDatabaseBackupSummaryLifecycleStateEnum = map[string]AutonomousDatabaseBackupSummaryLifecycleStateEnum{
	"CREATING": AutonomousDatabaseBackupSummaryLifecycleStateCreating,
	"ACTIVE":   AutonomousDatabaseBackupSummaryLifecycleStateActive,
	"DELETING": AutonomousDatabaseBackupSummaryLifecycleStateDeleting,
	"DELETED":  AutonomousDatabaseBackupSummaryLifecycleStateDeleted,
	"FAILED":   AutonomousDatabaseBackupSummaryLifecycleStateFailed,
	"UPDATING": AutonomousDatabaseBackupSummaryLifecycleStateUpdating,
}

var mappingAutonomousDatabaseBackupSummaryLifecycleStateEnumLowerCase = map[string]AutonomousDatabaseBackupSummaryLifecycleStateEnum{
	"creating": AutonomousDatabaseBackupSummaryLifecycleStateCreating,
	"active":   AutonomousDatabaseBackupSummaryLifecycleStateActive,
	"deleting": AutonomousDatabaseBackupSummaryLifecycleStateDeleting,
	"deleted":  AutonomousDatabaseBackupSummaryLifecycleStateDeleted,
	"failed":   AutonomousDatabaseBackupSummaryLifecycleStateFailed,
	"updating": AutonomousDatabaseBackupSummaryLifecycleStateUpdating,
}

// GetAutonomousDatabaseBackupSummaryLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseBackupSummaryLifecycleStateEnum
func GetAutonomousDatabaseBackupSummaryLifecycleStateEnumValues() []AutonomousDatabaseBackupSummaryLifecycleStateEnum {
	values := make([]AutonomousDatabaseBackupSummaryLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseBackupSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseBackupSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousDatabaseBackupSummaryLifecycleStateEnum
func GetAutonomousDatabaseBackupSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingAutonomousDatabaseBackupSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseBackupSummaryLifecycleStateEnum(val string) (AutonomousDatabaseBackupSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousDatabaseBackupSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
