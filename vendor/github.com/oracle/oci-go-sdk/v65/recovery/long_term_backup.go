// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database Autonomous Recovery Service API
//
// Use Oracle Database Autonomous Recovery Service API to manage Protected Databases.
//

package recovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LongTermBackup The details of a Long Term Backup.
type LongTermBackup struct {

	// The OCID of the long-term backup.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the long-term backup.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the protected database associated with the long-term backup.
	ProtectedDatabaseId *string `mandatory:"true" json:"protectedDatabaseId"`

	// The maximum number of DAYS or YEARS to store the long-term backup. You can retain a long-term backup for a period ranging from 90-3650 days or 1-10 years.
	RetentionPeriod []RetentionPeriodValue `mandatory:"true" json:"retentionPeriod"`

	// Indicates that Recovery Service must retain the backup for the specified long-term retention period.
	RetentionUntilDateTime *common.SDKTime `mandatory:"true" json:"retentionUntilDateTime"`

	// The current state of the long term backup.
	LifecycleState LongTermBackupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The user-provided name for the long-term backup. You can change the displayName. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The unique system change number (SCN) or the target point in the database until which the long-term backup is consistent.
	RetentionScn *int `mandatory:"false" json:"retentionScn"`

	// An RFC3339 formatted datetime string that indicates the target point in time until which the long-term backup is consistent. For example, '2020-05-22T21:10:29.600Z'.
	RetentionPointInTime *common.SDKTime `mandatory:"false" json:"retentionPointInTime"`

	// An RFC3339 formatted datetime string that indicates the time when the long-term backup was created. For example, '2020-05-22T21:10:29.600Z'.
	TimeBackupInitiated *common.SDKTime `mandatory:"false" json:"timeBackupInitiated"`

	// An RFC3339 formatted datetime string that indicates the time when the long-term backup completed. For example, '2020-05-22T21:10:29.600Z'.
	TimeBackupCompleted *common.SDKTime `mandatory:"false" json:"timeBackupCompleted"`

	// The Oracle Database ID, which identifies an Oracle Database located outside of Oracle Cloud.
	DatabaseIdentifier *string `mandatory:"false" json:"databaseIdentifier"`

	// The size of the database, in gigabytes.
	DatabaseSizeInGBs *int `mandatory:"false" json:"databaseSizeInGBs"`

	// An RFC3339 formatted datetime string that indicates the time when the long-term backup was created. For example: '2020-05-22T21:10:29.600Z'.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// An RFC3339 formatted datetime string that indicates the time when the long term backup was last updated. For example: '2020-05-22T21:10:29.600Z'.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// More details on the state of the backup when it is in Creating lifecycleState.
	LifecycleSubstate LongTermBackupLifecycleSubstateEnum `mandatory:"false" json:"lifecycleSubstate,omitempty"`

	// A detailed message about the current lifecycle state of the long-term backup. For example, it can be used to provide actionable information for a resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Recovery Manager (RMAN) assigned unique identifier for the long-term backup.
	RmanTag *string `mandatory:"false" json:"rmanTag"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`. For more information, see Resource Tags (https://docs.oracle.com/en-us/iaas/Content/General/Concepts/resourcetags.htm)
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`. For more information, see Resource Tags (https://docs.oracle.com/en-us/iaas/Content/General/Concepts/resourcetags.htm)
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m LongTermBackup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LongTermBackup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLongTermBackupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLongTermBackupLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLongTermBackupLifecycleSubstateEnum(string(m.LifecycleSubstate)); !ok && m.LifecycleSubstate != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleSubstate: %s. Supported values are: %s.", m.LifecycleSubstate, strings.Join(GetLongTermBackupLifecycleSubstateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LongTermBackupLifecycleStateEnum Enum with underlying type: string
type LongTermBackupLifecycleStateEnum string

// Set of constants representing the allowable values for LongTermBackupLifecycleStateEnum
const (
	LongTermBackupLifecycleStateCreating  LongTermBackupLifecycleStateEnum = "CREATING"
	LongTermBackupLifecycleStateUpdating  LongTermBackupLifecycleStateEnum = "UPDATING"
	LongTermBackupLifecycleStateActive    LongTermBackupLifecycleStateEnum = "ACTIVE"
	LongTermBackupLifecycleStateDeleting  LongTermBackupLifecycleStateEnum = "DELETING"
	LongTermBackupLifecycleStateDeleted   LongTermBackupLifecycleStateEnum = "DELETED"
	LongTermBackupLifecycleStateFailed    LongTermBackupLifecycleStateEnum = "FAILED"
	LongTermBackupLifecycleStateCanceling LongTermBackupLifecycleStateEnum = "CANCELING"
	LongTermBackupLifecycleStateCanceled  LongTermBackupLifecycleStateEnum = "CANCELED"
)

var mappingLongTermBackupLifecycleStateEnum = map[string]LongTermBackupLifecycleStateEnum{
	"CREATING":  LongTermBackupLifecycleStateCreating,
	"UPDATING":  LongTermBackupLifecycleStateUpdating,
	"ACTIVE":    LongTermBackupLifecycleStateActive,
	"DELETING":  LongTermBackupLifecycleStateDeleting,
	"DELETED":   LongTermBackupLifecycleStateDeleted,
	"FAILED":    LongTermBackupLifecycleStateFailed,
	"CANCELING": LongTermBackupLifecycleStateCanceling,
	"CANCELED":  LongTermBackupLifecycleStateCanceled,
}

var mappingLongTermBackupLifecycleStateEnumLowerCase = map[string]LongTermBackupLifecycleStateEnum{
	"creating":  LongTermBackupLifecycleStateCreating,
	"updating":  LongTermBackupLifecycleStateUpdating,
	"active":    LongTermBackupLifecycleStateActive,
	"deleting":  LongTermBackupLifecycleStateDeleting,
	"deleted":   LongTermBackupLifecycleStateDeleted,
	"failed":    LongTermBackupLifecycleStateFailed,
	"canceling": LongTermBackupLifecycleStateCanceling,
	"canceled":  LongTermBackupLifecycleStateCanceled,
}

// GetLongTermBackupLifecycleStateEnumValues Enumerates the set of values for LongTermBackupLifecycleStateEnum
func GetLongTermBackupLifecycleStateEnumValues() []LongTermBackupLifecycleStateEnum {
	values := make([]LongTermBackupLifecycleStateEnum, 0)
	for _, v := range mappingLongTermBackupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLongTermBackupLifecycleStateEnumStringValues Enumerates the set of values in String for LongTermBackupLifecycleStateEnum
func GetLongTermBackupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingLongTermBackupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLongTermBackupLifecycleStateEnum(val string) (LongTermBackupLifecycleStateEnum, bool) {
	enum, ok := mappingLongTermBackupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// LongTermBackupLifecycleSubstateEnum Enum with underlying type: string
type LongTermBackupLifecycleSubstateEnum string

// Set of constants representing the allowable values for LongTermBackupLifecycleSubstateEnum
const (
	LongTermBackupLifecycleSubstateWaitingForBackupFromDb LongTermBackupLifecycleSubstateEnum = "WAITING_FOR_BACKUP_FROM_DB"
	LongTermBackupLifecycleSubstateScheduledForArchival   LongTermBackupLifecycleSubstateEnum = "SCHEDULED_FOR_ARCHIVAL"
	LongTermBackupLifecycleSubstateArchivalInProgress     LongTermBackupLifecycleSubstateEnum = "ARCHIVAL_IN_PROGRESS"
)

var mappingLongTermBackupLifecycleSubstateEnum = map[string]LongTermBackupLifecycleSubstateEnum{
	"WAITING_FOR_BACKUP_FROM_DB": LongTermBackupLifecycleSubstateWaitingForBackupFromDb,
	"SCHEDULED_FOR_ARCHIVAL":     LongTermBackupLifecycleSubstateScheduledForArchival,
	"ARCHIVAL_IN_PROGRESS":       LongTermBackupLifecycleSubstateArchivalInProgress,
}

var mappingLongTermBackupLifecycleSubstateEnumLowerCase = map[string]LongTermBackupLifecycleSubstateEnum{
	"waiting_for_backup_from_db": LongTermBackupLifecycleSubstateWaitingForBackupFromDb,
	"scheduled_for_archival":     LongTermBackupLifecycleSubstateScheduledForArchival,
	"archival_in_progress":       LongTermBackupLifecycleSubstateArchivalInProgress,
}

// GetLongTermBackupLifecycleSubstateEnumValues Enumerates the set of values for LongTermBackupLifecycleSubstateEnum
func GetLongTermBackupLifecycleSubstateEnumValues() []LongTermBackupLifecycleSubstateEnum {
	values := make([]LongTermBackupLifecycleSubstateEnum, 0)
	for _, v := range mappingLongTermBackupLifecycleSubstateEnum {
		values = append(values, v)
	}
	return values
}

// GetLongTermBackupLifecycleSubstateEnumStringValues Enumerates the set of values in String for LongTermBackupLifecycleSubstateEnum
func GetLongTermBackupLifecycleSubstateEnumStringValues() []string {
	return []string{
		"WAITING_FOR_BACKUP_FROM_DB",
		"SCHEDULED_FOR_ARCHIVAL",
		"ARCHIVAL_IN_PROGRESS",
	}
}

// GetMappingLongTermBackupLifecycleSubstateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLongTermBackupLifecycleSubstateEnum(val string) (LongTermBackupLifecycleSubstateEnum, bool) {
	enum, ok := mappingLongTermBackupLifecycleSubstateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
