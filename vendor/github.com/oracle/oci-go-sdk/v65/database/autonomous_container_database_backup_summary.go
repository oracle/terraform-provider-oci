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

// AutonomousContainerDatabaseBackupSummary An Autonomous Container Database backup.
// To use any API operation, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type AutonomousContainerDatabaseBackupSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database backup.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Container Database.
	AutonomousContainerDatabaseId *string `mandatory:"true" json:"autonomousContainerDatabaseId"`

	// A user-friendly name for the backup. This name need not be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of backup.
	Type AutonomousContainerDatabaseBackupSummaryTypeEnum `mandatory:"true" json:"type"`

	// Indicates whether the backup is user-initiated or automatic.
	IsAutomatic *bool `mandatory:"true" json:"isAutomatic"`

	// The current state of the backup.
	LifecycleState AutonomousContainerDatabaseBackupSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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
	InfrastructureType AutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnum `mandatory:"false" json:"infrastructureType,omitempty"`

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

func (m AutonomousContainerDatabaseBackupSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousContainerDatabaseBackupSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousContainerDatabaseBackupSummaryTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAutonomousContainerDatabaseBackupSummaryTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousContainerDatabaseBackupSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousContainerDatabaseBackupSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnum(string(m.InfrastructureType)); !ok && m.InfrastructureType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InfrastructureType: %s. Supported values are: %s.", m.InfrastructureType, strings.Join(GetAutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousContainerDatabaseBackupSummaryTypeEnum Enum with underlying type: string
type AutonomousContainerDatabaseBackupSummaryTypeEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseBackupSummaryTypeEnum
const (
	AutonomousContainerDatabaseBackupSummaryTypeIncremental AutonomousContainerDatabaseBackupSummaryTypeEnum = "INCREMENTAL"
	AutonomousContainerDatabaseBackupSummaryTypeFull        AutonomousContainerDatabaseBackupSummaryTypeEnum = "FULL"
)

var mappingAutonomousContainerDatabaseBackupSummaryTypeEnum = map[string]AutonomousContainerDatabaseBackupSummaryTypeEnum{
	"INCREMENTAL": AutonomousContainerDatabaseBackupSummaryTypeIncremental,
	"FULL":        AutonomousContainerDatabaseBackupSummaryTypeFull,
}

var mappingAutonomousContainerDatabaseBackupSummaryTypeEnumLowerCase = map[string]AutonomousContainerDatabaseBackupSummaryTypeEnum{
	"incremental": AutonomousContainerDatabaseBackupSummaryTypeIncremental,
	"full":        AutonomousContainerDatabaseBackupSummaryTypeFull,
}

// GetAutonomousContainerDatabaseBackupSummaryTypeEnumValues Enumerates the set of values for AutonomousContainerDatabaseBackupSummaryTypeEnum
func GetAutonomousContainerDatabaseBackupSummaryTypeEnumValues() []AutonomousContainerDatabaseBackupSummaryTypeEnum {
	values := make([]AutonomousContainerDatabaseBackupSummaryTypeEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseBackupSummaryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseBackupSummaryTypeEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseBackupSummaryTypeEnum
func GetAutonomousContainerDatabaseBackupSummaryTypeEnumStringValues() []string {
	return []string{
		"INCREMENTAL",
		"FULL",
	}
}

// GetMappingAutonomousContainerDatabaseBackupSummaryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseBackupSummaryTypeEnum(val string) (AutonomousContainerDatabaseBackupSummaryTypeEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseBackupSummaryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousContainerDatabaseBackupSummaryLifecycleStateEnum Enum with underlying type: string
type AutonomousContainerDatabaseBackupSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseBackupSummaryLifecycleStateEnum
const (
	AutonomousContainerDatabaseBackupSummaryLifecycleStateCreating AutonomousContainerDatabaseBackupSummaryLifecycleStateEnum = "CREATING"
	AutonomousContainerDatabaseBackupSummaryLifecycleStateActive   AutonomousContainerDatabaseBackupSummaryLifecycleStateEnum = "ACTIVE"
	AutonomousContainerDatabaseBackupSummaryLifecycleStateDeleting AutonomousContainerDatabaseBackupSummaryLifecycleStateEnum = "DELETING"
	AutonomousContainerDatabaseBackupSummaryLifecycleStateDeleted  AutonomousContainerDatabaseBackupSummaryLifecycleStateEnum = "DELETED"
	AutonomousContainerDatabaseBackupSummaryLifecycleStateFailed   AutonomousContainerDatabaseBackupSummaryLifecycleStateEnum = "FAILED"
)

var mappingAutonomousContainerDatabaseBackupSummaryLifecycleStateEnum = map[string]AutonomousContainerDatabaseBackupSummaryLifecycleStateEnum{
	"CREATING": AutonomousContainerDatabaseBackupSummaryLifecycleStateCreating,
	"ACTIVE":   AutonomousContainerDatabaseBackupSummaryLifecycleStateActive,
	"DELETING": AutonomousContainerDatabaseBackupSummaryLifecycleStateDeleting,
	"DELETED":  AutonomousContainerDatabaseBackupSummaryLifecycleStateDeleted,
	"FAILED":   AutonomousContainerDatabaseBackupSummaryLifecycleStateFailed,
}

var mappingAutonomousContainerDatabaseBackupSummaryLifecycleStateEnumLowerCase = map[string]AutonomousContainerDatabaseBackupSummaryLifecycleStateEnum{
	"creating": AutonomousContainerDatabaseBackupSummaryLifecycleStateCreating,
	"active":   AutonomousContainerDatabaseBackupSummaryLifecycleStateActive,
	"deleting": AutonomousContainerDatabaseBackupSummaryLifecycleStateDeleting,
	"deleted":  AutonomousContainerDatabaseBackupSummaryLifecycleStateDeleted,
	"failed":   AutonomousContainerDatabaseBackupSummaryLifecycleStateFailed,
}

// GetAutonomousContainerDatabaseBackupSummaryLifecycleStateEnumValues Enumerates the set of values for AutonomousContainerDatabaseBackupSummaryLifecycleStateEnum
func GetAutonomousContainerDatabaseBackupSummaryLifecycleStateEnumValues() []AutonomousContainerDatabaseBackupSummaryLifecycleStateEnum {
	values := make([]AutonomousContainerDatabaseBackupSummaryLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseBackupSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseBackupSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseBackupSummaryLifecycleStateEnum
func GetAutonomousContainerDatabaseBackupSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingAutonomousContainerDatabaseBackupSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseBackupSummaryLifecycleStateEnum(val string) (AutonomousContainerDatabaseBackupSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseBackupSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnum Enum with underlying type: string
type AutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnum
const (
	AutonomousContainerDatabaseBackupSummaryInfrastructureTypeCloud           AutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnum = "CLOUD"
	AutonomousContainerDatabaseBackupSummaryInfrastructureTypeCloudAtCustomer AutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnum = "CLOUD_AT_CUSTOMER"
)

var mappingAutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnum = map[string]AutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnum{
	"CLOUD":             AutonomousContainerDatabaseBackupSummaryInfrastructureTypeCloud,
	"CLOUD_AT_CUSTOMER": AutonomousContainerDatabaseBackupSummaryInfrastructureTypeCloudAtCustomer,
}

var mappingAutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnumLowerCase = map[string]AutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnum{
	"cloud":             AutonomousContainerDatabaseBackupSummaryInfrastructureTypeCloud,
	"cloud_at_customer": AutonomousContainerDatabaseBackupSummaryInfrastructureTypeCloudAtCustomer,
}

// GetAutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnumValues Enumerates the set of values for AutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnum
func GetAutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnumValues() []AutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnum {
	values := make([]AutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnum
func GetAutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnumStringValues() []string {
	return []string{
		"CLOUD",
		"CLOUD_AT_CUSTOMER",
	}
}

// GetMappingAutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnum(val string) (AutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
