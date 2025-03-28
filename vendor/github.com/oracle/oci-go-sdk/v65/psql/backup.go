// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Backup Database system backup information.
type Backup struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly display name for the backup. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the backup.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the backup request was received, expressed in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the backup.
	LifecycleState BackupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The size of the backup, in gigabytes.
	BackupSize *int `mandatory:"true" json:"backupSize"`

	DbSystemDetails *DbSystemDetails `mandatory:"true" json:"dbSystemDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// A description for the backup.
	Description *string `mandatory:"false" json:"description"`

	// Specifies whether the backup was created manually, taken on schedule defined in the a backup policy, or copied from the remote location.
	SourceType BackupSourceTypeEnum `mandatory:"false" json:"sourceType,omitempty"`

	// The date and time the backup was created.
	// This is the time the actual point-in-time data snapshot was taken,
	// expressed in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreatedPrecise *common.SDKTime `mandatory:"false" json:"timeCreatedPrecise"`

	// The date and time the backup was updated, expressed in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Backup retention period in days.
	RetentionPeriod *int `mandatory:"false" json:"retentionPeriod"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup's source database system.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	SourceBackupDetails *SourceBackupDetails `mandatory:"false" json:"sourceBackupDetails"`

	// lastAcceptedRequestToken from MP.
	LastAcceptedRequestToken *string `mandatory:"false" json:"lastAcceptedRequestToken"`

	// lastCompletedRequestToken from MP.
	LastCompletedRequestToken *string `mandatory:"false" json:"lastCompletedRequestToken"`

	// List of status for Backup Copy
	CopyStatus []BackupCopyStatusDetails `mandatory:"false" json:"copyStatus"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Backup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Backup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBackupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBackupLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBackupSourceTypeEnum(string(m.SourceType)); !ok && m.SourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceType: %s. Supported values are: %s.", m.SourceType, strings.Join(GetBackupSourceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BackupSourceTypeEnum Enum with underlying type: string
type BackupSourceTypeEnum string

// Set of constants representing the allowable values for BackupSourceTypeEnum
const (
	BackupSourceTypeScheduled BackupSourceTypeEnum = "SCHEDULED"
	BackupSourceTypeManual    BackupSourceTypeEnum = "MANUAL"
	BackupSourceTypeCopied    BackupSourceTypeEnum = "COPIED"
)

var mappingBackupSourceTypeEnum = map[string]BackupSourceTypeEnum{
	"SCHEDULED": BackupSourceTypeScheduled,
	"MANUAL":    BackupSourceTypeManual,
	"COPIED":    BackupSourceTypeCopied,
}

var mappingBackupSourceTypeEnumLowerCase = map[string]BackupSourceTypeEnum{
	"scheduled": BackupSourceTypeScheduled,
	"manual":    BackupSourceTypeManual,
	"copied":    BackupSourceTypeCopied,
}

// GetBackupSourceTypeEnumValues Enumerates the set of values for BackupSourceTypeEnum
func GetBackupSourceTypeEnumValues() []BackupSourceTypeEnum {
	values := make([]BackupSourceTypeEnum, 0)
	for _, v := range mappingBackupSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupSourceTypeEnumStringValues Enumerates the set of values in String for BackupSourceTypeEnum
func GetBackupSourceTypeEnumStringValues() []string {
	return []string{
		"SCHEDULED",
		"MANUAL",
		"COPIED",
	}
}

// GetMappingBackupSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupSourceTypeEnum(val string) (BackupSourceTypeEnum, bool) {
	enum, ok := mappingBackupSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BackupLifecycleStateEnum Enum with underlying type: string
type BackupLifecycleStateEnum string

// Set of constants representing the allowable values for BackupLifecycleStateEnum
const (
	BackupLifecycleStateCreating BackupLifecycleStateEnum = "CREATING"
	BackupLifecycleStateActive   BackupLifecycleStateEnum = "ACTIVE"
	BackupLifecycleStateDeleting BackupLifecycleStateEnum = "DELETING"
	BackupLifecycleStateDeleted  BackupLifecycleStateEnum = "DELETED"
	BackupLifecycleStateFailed   BackupLifecycleStateEnum = "FAILED"
)

var mappingBackupLifecycleStateEnum = map[string]BackupLifecycleStateEnum{
	"CREATING": BackupLifecycleStateCreating,
	"ACTIVE":   BackupLifecycleStateActive,
	"DELETING": BackupLifecycleStateDeleting,
	"DELETED":  BackupLifecycleStateDeleted,
	"FAILED":   BackupLifecycleStateFailed,
}

var mappingBackupLifecycleStateEnumLowerCase = map[string]BackupLifecycleStateEnum{
	"creating": BackupLifecycleStateCreating,
	"active":   BackupLifecycleStateActive,
	"deleting": BackupLifecycleStateDeleting,
	"deleted":  BackupLifecycleStateDeleted,
	"failed":   BackupLifecycleStateFailed,
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
	}
}

// GetMappingBackupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupLifecycleStateEnum(val string) (BackupLifecycleStateEnum, bool) {
	enum, ok := mappingBackupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
