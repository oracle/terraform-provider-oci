// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutonomousContainerDatabaseBackup An Autonomous Database backup.
type AutonomousContainerDatabaseBackup struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database backup.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Container Database.
	AutonomousContainerDatabaseId *string `mandatory:"true" json:"autonomousContainerDatabaseId"`

	// A user-friendly name for the backup. This name need not be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of backup.
	Type AutonomousContainerDatabaseBackupTypeEnum `mandatory:"true" json:"type"`

	// Indicates whether the backup is user-initiated or automatic.
	IsAutomatic *bool `mandatory:"true" json:"isAutomatic"`

	// The current state of the backup.
	LifecycleState AutonomousContainerDatabaseBackupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The user-friendly name for the Autonomous Container Database when the Backup was initiated. This name need not be unique. This field captures the name at the time of backup creation, accounting for possible later updates to the display name.
	AcdDisplayName *string `mandatory:"false" json:"acdDisplayName"`

	// The date and time the backup started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the backup completed.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Retention period, in days, for long-term backups
	RetentionPeriodInDays *int `mandatory:"false" json:"retentionPeriodInDays"`

	// Whether backup is for remote-region or local region
	IsRemoteBackup *bool `mandatory:"false" json:"isRemoteBackup"`

	// The infrastructure type this resource belongs to.
	InfrastructureType AutonomousContainerDatabaseBackupInfrastructureTypeEnum `mandatory:"false" json:"infrastructureType,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// List of Autonomous Databases that is part of this Autonomous Container Database Backup
	AutonomousDatabases []AutonomousDatabaseInBackup `mandatory:"false" json:"autonomousDatabases"`
}

func (m AutonomousContainerDatabaseBackup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousContainerDatabaseBackup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousContainerDatabaseBackupTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAutonomousContainerDatabaseBackupTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousContainerDatabaseBackupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousContainerDatabaseBackupLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAutonomousContainerDatabaseBackupInfrastructureTypeEnum(string(m.InfrastructureType)); !ok && m.InfrastructureType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InfrastructureType: %s. Supported values are: %s.", m.InfrastructureType, strings.Join(GetAutonomousContainerDatabaseBackupInfrastructureTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousContainerDatabaseBackupTypeEnum Enum with underlying type: string
type AutonomousContainerDatabaseBackupTypeEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseBackupTypeEnum
const (
	AutonomousContainerDatabaseBackupTypeIncremental AutonomousContainerDatabaseBackupTypeEnum = "INCREMENTAL"
	AutonomousContainerDatabaseBackupTypeFull        AutonomousContainerDatabaseBackupTypeEnum = "FULL"
)

var mappingAutonomousContainerDatabaseBackupTypeEnum = map[string]AutonomousContainerDatabaseBackupTypeEnum{
	"INCREMENTAL": AutonomousContainerDatabaseBackupTypeIncremental,
	"FULL":        AutonomousContainerDatabaseBackupTypeFull,
}

var mappingAutonomousContainerDatabaseBackupTypeEnumLowerCase = map[string]AutonomousContainerDatabaseBackupTypeEnum{
	"incremental": AutonomousContainerDatabaseBackupTypeIncremental,
	"full":        AutonomousContainerDatabaseBackupTypeFull,
}

// GetAutonomousContainerDatabaseBackupTypeEnumValues Enumerates the set of values for AutonomousContainerDatabaseBackupTypeEnum
func GetAutonomousContainerDatabaseBackupTypeEnumValues() []AutonomousContainerDatabaseBackupTypeEnum {
	values := make([]AutonomousContainerDatabaseBackupTypeEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseBackupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseBackupTypeEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseBackupTypeEnum
func GetAutonomousContainerDatabaseBackupTypeEnumStringValues() []string {
	return []string{
		"INCREMENTAL",
		"FULL",
	}
}

// GetMappingAutonomousContainerDatabaseBackupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseBackupTypeEnum(val string) (AutonomousContainerDatabaseBackupTypeEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseBackupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousContainerDatabaseBackupLifecycleStateEnum Enum with underlying type: string
type AutonomousContainerDatabaseBackupLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseBackupLifecycleStateEnum
const (
	AutonomousContainerDatabaseBackupLifecycleStateCreating AutonomousContainerDatabaseBackupLifecycleStateEnum = "CREATING"
	AutonomousContainerDatabaseBackupLifecycleStateActive   AutonomousContainerDatabaseBackupLifecycleStateEnum = "ACTIVE"
	AutonomousContainerDatabaseBackupLifecycleStateDeleting AutonomousContainerDatabaseBackupLifecycleStateEnum = "DELETING"
	AutonomousContainerDatabaseBackupLifecycleStateDeleted  AutonomousContainerDatabaseBackupLifecycleStateEnum = "DELETED"
	AutonomousContainerDatabaseBackupLifecycleStateFailed   AutonomousContainerDatabaseBackupLifecycleStateEnum = "FAILED"
)

var mappingAutonomousContainerDatabaseBackupLifecycleStateEnum = map[string]AutonomousContainerDatabaseBackupLifecycleStateEnum{
	"CREATING": AutonomousContainerDatabaseBackupLifecycleStateCreating,
	"ACTIVE":   AutonomousContainerDatabaseBackupLifecycleStateActive,
	"DELETING": AutonomousContainerDatabaseBackupLifecycleStateDeleting,
	"DELETED":  AutonomousContainerDatabaseBackupLifecycleStateDeleted,
	"FAILED":   AutonomousContainerDatabaseBackupLifecycleStateFailed,
}

var mappingAutonomousContainerDatabaseBackupLifecycleStateEnumLowerCase = map[string]AutonomousContainerDatabaseBackupLifecycleStateEnum{
	"creating": AutonomousContainerDatabaseBackupLifecycleStateCreating,
	"active":   AutonomousContainerDatabaseBackupLifecycleStateActive,
	"deleting": AutonomousContainerDatabaseBackupLifecycleStateDeleting,
	"deleted":  AutonomousContainerDatabaseBackupLifecycleStateDeleted,
	"failed":   AutonomousContainerDatabaseBackupLifecycleStateFailed,
}

// GetAutonomousContainerDatabaseBackupLifecycleStateEnumValues Enumerates the set of values for AutonomousContainerDatabaseBackupLifecycleStateEnum
func GetAutonomousContainerDatabaseBackupLifecycleStateEnumValues() []AutonomousContainerDatabaseBackupLifecycleStateEnum {
	values := make([]AutonomousContainerDatabaseBackupLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseBackupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseBackupLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseBackupLifecycleStateEnum
func GetAutonomousContainerDatabaseBackupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingAutonomousContainerDatabaseBackupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseBackupLifecycleStateEnum(val string) (AutonomousContainerDatabaseBackupLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseBackupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousContainerDatabaseBackupInfrastructureTypeEnum Enum with underlying type: string
type AutonomousContainerDatabaseBackupInfrastructureTypeEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseBackupInfrastructureTypeEnum
const (
	AutonomousContainerDatabaseBackupInfrastructureTypeCloud           AutonomousContainerDatabaseBackupInfrastructureTypeEnum = "CLOUD"
	AutonomousContainerDatabaseBackupInfrastructureTypeCloudAtCustomer AutonomousContainerDatabaseBackupInfrastructureTypeEnum = "CLOUD_AT_CUSTOMER"
)

var mappingAutonomousContainerDatabaseBackupInfrastructureTypeEnum = map[string]AutonomousContainerDatabaseBackupInfrastructureTypeEnum{
	"CLOUD":             AutonomousContainerDatabaseBackupInfrastructureTypeCloud,
	"CLOUD_AT_CUSTOMER": AutonomousContainerDatabaseBackupInfrastructureTypeCloudAtCustomer,
}

var mappingAutonomousContainerDatabaseBackupInfrastructureTypeEnumLowerCase = map[string]AutonomousContainerDatabaseBackupInfrastructureTypeEnum{
	"cloud":             AutonomousContainerDatabaseBackupInfrastructureTypeCloud,
	"cloud_at_customer": AutonomousContainerDatabaseBackupInfrastructureTypeCloudAtCustomer,
}

// GetAutonomousContainerDatabaseBackupInfrastructureTypeEnumValues Enumerates the set of values for AutonomousContainerDatabaseBackupInfrastructureTypeEnum
func GetAutonomousContainerDatabaseBackupInfrastructureTypeEnumValues() []AutonomousContainerDatabaseBackupInfrastructureTypeEnum {
	values := make([]AutonomousContainerDatabaseBackupInfrastructureTypeEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseBackupInfrastructureTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseBackupInfrastructureTypeEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseBackupInfrastructureTypeEnum
func GetAutonomousContainerDatabaseBackupInfrastructureTypeEnumStringValues() []string {
	return []string{
		"CLOUD",
		"CLOUD_AT_CUSTOMER",
	}
}

// GetMappingAutonomousContainerDatabaseBackupInfrastructureTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseBackupInfrastructureTypeEnum(val string) (AutonomousContainerDatabaseBackupInfrastructureTypeEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseBackupInfrastructureTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
