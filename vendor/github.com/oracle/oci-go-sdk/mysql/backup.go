// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Backup A full or incremental copy of a DB System which can be used to create a
// new DB System or recover a DB System.
// To use any of the API operations, you must be authorized in an IAM
// policy. If you're not authorized, talk to an administrator. If you're an
// administrator who needs to write policies to give users access, see
// Getting Started with
// Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type Backup struct {

	// OCID of the backup itself
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the backup record was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time at which the backup was updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The state of the backup.
	LifecycleState BackupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// The type of backup.
	BackupType BackupBackupTypeEnum `mandatory:"true" json:"backupType"`

	// If the backup was created automatically, or by a manual request.
	CreationType BackupCreationTypeEnum `mandatory:"true" json:"creationType"`

	// The OCID of the DB System the backup is associated with.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// A user-supplied display name for the backup.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A user-supplied description for the backup.
	Description *string `mandatory:"false" json:"description"`

	// The size of the backup in base-2 (IEC) gibibytes. (GiB).
	BackupSizeInGBs *int `mandatory:"false" json:"backupSizeInGBs"`

	// Number of days to retain this backup.
	RetentionInDays *int `mandatory:"false" json:"retentionInDays"`

	// Initial size of the data volume in GiBs.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	// The MySQL server version of the DB System used for backup.
	MysqlVersion *string `mandatory:"false" json:"mysqlVersion"`

	// The shape of the DB System used for backup.
	ShapeName *string `mandatory:"false" json:"shapeName"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Backup) String() string {
	return common.PointerString(m)
}

// BackupLifecycleStateEnum Enum with underlying type: string
type BackupLifecycleStateEnum string

// Set of constants representing the allowable values for BackupLifecycleStateEnum
const (
	BackupLifecycleStateCreating BackupLifecycleStateEnum = "CREATING"
	BackupLifecycleStateActive   BackupLifecycleStateEnum = "ACTIVE"
	BackupLifecycleStateInactive BackupLifecycleStateEnum = "INACTIVE"
	BackupLifecycleStateUpdating BackupLifecycleStateEnum = "UPDATING"
	BackupLifecycleStateDeleting BackupLifecycleStateEnum = "DELETING"
	BackupLifecycleStateDeleted  BackupLifecycleStateEnum = "DELETED"
	BackupLifecycleStateFailed   BackupLifecycleStateEnum = "FAILED"
)

var mappingBackupLifecycleState = map[string]BackupLifecycleStateEnum{
	"CREATING": BackupLifecycleStateCreating,
	"ACTIVE":   BackupLifecycleStateActive,
	"INACTIVE": BackupLifecycleStateInactive,
	"UPDATING": BackupLifecycleStateUpdating,
	"DELETING": BackupLifecycleStateDeleting,
	"DELETED":  BackupLifecycleStateDeleted,
	"FAILED":   BackupLifecycleStateFailed,
}

// GetBackupLifecycleStateEnumValues Enumerates the set of values for BackupLifecycleStateEnum
func GetBackupLifecycleStateEnumValues() []BackupLifecycleStateEnum {
	values := make([]BackupLifecycleStateEnum, 0)
	for _, v := range mappingBackupLifecycleState {
		values = append(values, v)
	}
	return values
}

// BackupBackupTypeEnum Enum with underlying type: string
type BackupBackupTypeEnum string

// Set of constants representing the allowable values for BackupBackupTypeEnum
const (
	BackupBackupTypeFull        BackupBackupTypeEnum = "FULL"
	BackupBackupTypeIncremental BackupBackupTypeEnum = "INCREMENTAL"
)

var mappingBackupBackupType = map[string]BackupBackupTypeEnum{
	"FULL":        BackupBackupTypeFull,
	"INCREMENTAL": BackupBackupTypeIncremental,
}

// GetBackupBackupTypeEnumValues Enumerates the set of values for BackupBackupTypeEnum
func GetBackupBackupTypeEnumValues() []BackupBackupTypeEnum {
	values := make([]BackupBackupTypeEnum, 0)
	for _, v := range mappingBackupBackupType {
		values = append(values, v)
	}
	return values
}

// BackupCreationTypeEnum Enum with underlying type: string
type BackupCreationTypeEnum string

// Set of constants representing the allowable values for BackupCreationTypeEnum
const (
	BackupCreationTypeManual    BackupCreationTypeEnum = "MANUAL"
	BackupCreationTypeAutomatic BackupCreationTypeEnum = "AUTOMATIC"
)

var mappingBackupCreationType = map[string]BackupCreationTypeEnum{
	"MANUAL":    BackupCreationTypeManual,
	"AUTOMATIC": BackupCreationTypeAutomatic,
}

// GetBackupCreationTypeEnumValues Enumerates the set of values for BackupCreationTypeEnum
func GetBackupCreationTypeEnumValues() []BackupCreationTypeEnum {
	values := make([]BackupCreationTypeEnum, 0)
	for _, v := range mappingBackupCreationType {
		values = append(values, v)
	}
	return values
}
