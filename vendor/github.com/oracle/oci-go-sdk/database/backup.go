// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Backup A database backup
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type Backup struct {

	// The name of the Availability Domain that the backup is located in.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID of the database.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	// The user-friendly name for the backup. It does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the backup.
	Id *string `mandatory:"false" json:"id"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current state of the backup.
	LifecycleState BackupLifecycleStateEnum `mandatory:"false" json:"lifecycleState"`

	// The date and time the backup was completed.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// The date and time the backup starts.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The type of backup.
	Type BackupTypeEnum `mandatory:"false" json:"type"`
}

func (m Backup) String() string {
	return common.PointerString(m)
}

// BackupLifecycleStateEnum Enum with underlying type: string
type BackupLifecycleStateEnum string

// Set of constants representing the allowable values for BackupLifecycleState
const (
	BackupLifecycleStateCreating  BackupLifecycleStateEnum = "CREATING"
	BackupLifecycleStateActive    BackupLifecycleStateEnum = "ACTIVE"
	BackupLifecycleStateDeleting  BackupLifecycleStateEnum = "DELETING"
	BackupLifecycleStateDeleted   BackupLifecycleStateEnum = "DELETED"
	BackupLifecycleStateFailed    BackupLifecycleStateEnum = "FAILED"
	BackupLifecycleStateRestoring BackupLifecycleStateEnum = "RESTORING"
	BackupLifecycleStateUnknown   BackupLifecycleStateEnum = "UNKNOWN"
)

var mappingBackupLifecycleState = map[string]BackupLifecycleStateEnum{
	"CREATING":  BackupLifecycleStateCreating,
	"ACTIVE":    BackupLifecycleStateActive,
	"DELETING":  BackupLifecycleStateDeleting,
	"DELETED":   BackupLifecycleStateDeleted,
	"FAILED":    BackupLifecycleStateFailed,
	"RESTORING": BackupLifecycleStateRestoring,
	"UNKNOWN":   BackupLifecycleStateUnknown,
}

// GetBackupLifecycleStateEnumValues Enumerates the set of values for BackupLifecycleState
func GetBackupLifecycleStateEnumValues() []BackupLifecycleStateEnum {
	values := make([]BackupLifecycleStateEnum, 0)
	for _, v := range mappingBackupLifecycleState {
		if v != BackupLifecycleStateUnknown {
			values = append(values, v)
		}
	}
	return values
}

// BackupTypeEnum Enum with underlying type: string
type BackupTypeEnum string

// Set of constants representing the allowable values for BackupType
const (
	BackupTypeIncremental BackupTypeEnum = "INCREMENTAL"
	BackupTypeFull        BackupTypeEnum = "FULL"
	BackupTypeUnknown     BackupTypeEnum = "UNKNOWN"
)

var mappingBackupType = map[string]BackupTypeEnum{
	"INCREMENTAL": BackupTypeIncremental,
	"FULL":        BackupTypeFull,
	"UNKNOWN":     BackupTypeUnknown,
}

// GetBackupTypeEnumValues Enumerates the set of values for BackupType
func GetBackupTypeEnumValues() []BackupTypeEnum {
	values := make([]BackupTypeEnum, 0)
	for _, v := range mappingBackupType {
		if v != BackupTypeUnknown {
			values = append(values, v)
		}
	}
	return values
}
