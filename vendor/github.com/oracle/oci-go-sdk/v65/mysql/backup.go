// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

	// Indicates how the backup was created: manually, automatic, or by an Operator.
	CreationType BackupCreationTypeEnum `mandatory:"true" json:"creationType"`

	// The OCID of the DB System the backup is associated with.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// A user-supplied display name for the backup.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A user-supplied description for the backup.
	Description *string `mandatory:"false" json:"description"`

	DbSystemSnapshot *DbSystemSnapshot `mandatory:"false" json:"dbSystemSnapshot"`

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

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
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
	if _, ok := GetMappingBackupBackupTypeEnum(string(m.BackupType)); !ok && m.BackupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupType: %s. Supported values are: %s.", m.BackupType, strings.Join(GetBackupBackupTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBackupCreationTypeEnum(string(m.CreationType)); !ok && m.CreationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CreationType: %s. Supported values are: %s.", m.CreationType, strings.Join(GetBackupCreationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingBackupLifecycleStateEnum = map[string]BackupLifecycleStateEnum{
	"CREATING": BackupLifecycleStateCreating,
	"ACTIVE":   BackupLifecycleStateActive,
	"INACTIVE": BackupLifecycleStateInactive,
	"UPDATING": BackupLifecycleStateUpdating,
	"DELETING": BackupLifecycleStateDeleting,
	"DELETED":  BackupLifecycleStateDeleted,
	"FAILED":   BackupLifecycleStateFailed,
}

var mappingBackupLifecycleStateEnumLowerCase = map[string]BackupLifecycleStateEnum{
	"creating": BackupLifecycleStateCreating,
	"active":   BackupLifecycleStateActive,
	"inactive": BackupLifecycleStateInactive,
	"updating": BackupLifecycleStateUpdating,
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
		"INACTIVE",
		"UPDATING",
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

// BackupBackupTypeEnum Enum with underlying type: string
type BackupBackupTypeEnum string

// Set of constants representing the allowable values for BackupBackupTypeEnum
const (
	BackupBackupTypeFull        BackupBackupTypeEnum = "FULL"
	BackupBackupTypeIncremental BackupBackupTypeEnum = "INCREMENTAL"
)

var mappingBackupBackupTypeEnum = map[string]BackupBackupTypeEnum{
	"FULL":        BackupBackupTypeFull,
	"INCREMENTAL": BackupBackupTypeIncremental,
}

var mappingBackupBackupTypeEnumLowerCase = map[string]BackupBackupTypeEnum{
	"full":        BackupBackupTypeFull,
	"incremental": BackupBackupTypeIncremental,
}

// GetBackupBackupTypeEnumValues Enumerates the set of values for BackupBackupTypeEnum
func GetBackupBackupTypeEnumValues() []BackupBackupTypeEnum {
	values := make([]BackupBackupTypeEnum, 0)
	for _, v := range mappingBackupBackupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupBackupTypeEnumStringValues Enumerates the set of values in String for BackupBackupTypeEnum
func GetBackupBackupTypeEnumStringValues() []string {
	return []string{
		"FULL",
		"INCREMENTAL",
	}
}

// GetMappingBackupBackupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupBackupTypeEnum(val string) (BackupBackupTypeEnum, bool) {
	enum, ok := mappingBackupBackupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BackupCreationTypeEnum Enum with underlying type: string
type BackupCreationTypeEnum string

// Set of constants representing the allowable values for BackupCreationTypeEnum
const (
	BackupCreationTypeManual    BackupCreationTypeEnum = "MANUAL"
	BackupCreationTypeAutomatic BackupCreationTypeEnum = "AUTOMATIC"
	BackupCreationTypeOperator  BackupCreationTypeEnum = "OPERATOR"
)

var mappingBackupCreationTypeEnum = map[string]BackupCreationTypeEnum{
	"MANUAL":    BackupCreationTypeManual,
	"AUTOMATIC": BackupCreationTypeAutomatic,
	"OPERATOR":  BackupCreationTypeOperator,
}

var mappingBackupCreationTypeEnumLowerCase = map[string]BackupCreationTypeEnum{
	"manual":    BackupCreationTypeManual,
	"automatic": BackupCreationTypeAutomatic,
	"operator":  BackupCreationTypeOperator,
}

// GetBackupCreationTypeEnumValues Enumerates the set of values for BackupCreationTypeEnum
func GetBackupCreationTypeEnumValues() []BackupCreationTypeEnum {
	values := make([]BackupCreationTypeEnum, 0)
	for _, v := range mappingBackupCreationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupCreationTypeEnumStringValues Enumerates the set of values in String for BackupCreationTypeEnum
func GetBackupCreationTypeEnumStringValues() []string {
	return []string{
		"MANUAL",
		"AUTOMATIC",
		"OPERATOR",
	}
}

// GetMappingBackupCreationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupCreationTypeEnum(val string) (BackupCreationTypeEnum, bool) {
	enum, ok := mappingBackupCreationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
