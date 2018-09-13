// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AutonomousDatabaseBackup An Autonomous Database backup.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type AutonomousDatabaseBackup struct {

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
	AutonomousDatabaseId *string `mandatory:"true" json:"autonomousDatabaseId"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the backup. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Autonomous Database backup.
	Id *string `mandatory:"true" json:"id"`

	// Indicates whether the backup is user-initiated or automatic.
	IsAutomatic *bool `mandatory:"true" json:"isAutomatic"`

	// The current state of the backup.
	LifecycleState AutonomousDatabaseBackupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The type of backup.
	Type AutonomousDatabaseBackupTypeEnum `mandatory:"true" json:"type"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the backup completed.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// The date and time the backup started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`
}

func (m AutonomousDatabaseBackup) String() string {
	return common.PointerString(m)
}

// AutonomousDatabaseBackupLifecycleStateEnum Enum with underlying type: string
type AutonomousDatabaseBackupLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseBackupLifecycleState
const (
	AutonomousDatabaseBackupLifecycleStateCreating AutonomousDatabaseBackupLifecycleStateEnum = "CREATING"
	AutonomousDatabaseBackupLifecycleStateActive   AutonomousDatabaseBackupLifecycleStateEnum = "ACTIVE"
	AutonomousDatabaseBackupLifecycleStateDeleting AutonomousDatabaseBackupLifecycleStateEnum = "DELETING"
	AutonomousDatabaseBackupLifecycleStateDeleted  AutonomousDatabaseBackupLifecycleStateEnum = "DELETED"
	AutonomousDatabaseBackupLifecycleStateFailed   AutonomousDatabaseBackupLifecycleStateEnum = "FAILED"
)

var mappingAutonomousDatabaseBackupLifecycleState = map[string]AutonomousDatabaseBackupLifecycleStateEnum{
	"CREATING": AutonomousDatabaseBackupLifecycleStateCreating,
	"ACTIVE":   AutonomousDatabaseBackupLifecycleStateActive,
	"DELETING": AutonomousDatabaseBackupLifecycleStateDeleting,
	"DELETED":  AutonomousDatabaseBackupLifecycleStateDeleted,
	"FAILED":   AutonomousDatabaseBackupLifecycleStateFailed,
}

// GetAutonomousDatabaseBackupLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseBackupLifecycleState
func GetAutonomousDatabaseBackupLifecycleStateEnumValues() []AutonomousDatabaseBackupLifecycleStateEnum {
	values := make([]AutonomousDatabaseBackupLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseBackupLifecycleState {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseBackupTypeEnum Enum with underlying type: string
type AutonomousDatabaseBackupTypeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseBackupType
const (
	AutonomousDatabaseBackupTypeIncremental AutonomousDatabaseBackupTypeEnum = "INCREMENTAL"
	AutonomousDatabaseBackupTypeFull        AutonomousDatabaseBackupTypeEnum = "FULL"
)

var mappingAutonomousDatabaseBackupType = map[string]AutonomousDatabaseBackupTypeEnum{
	"INCREMENTAL": AutonomousDatabaseBackupTypeIncremental,
	"FULL":        AutonomousDatabaseBackupTypeFull,
}

// GetAutonomousDatabaseBackupTypeEnumValues Enumerates the set of values for AutonomousDatabaseBackupType
func GetAutonomousDatabaseBackupTypeEnumValues() []AutonomousDatabaseBackupTypeEnum {
	values := make([]AutonomousDatabaseBackupTypeEnum, 0)
	for _, v := range mappingAutonomousDatabaseBackupType {
		values = append(values, v)
	}
	return values
}
