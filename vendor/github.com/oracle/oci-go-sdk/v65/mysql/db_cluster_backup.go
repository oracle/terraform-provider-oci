// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DbClusterBackup A backup of a shared-storage DB cluster. The backup can be used to restore data into a new shared-storage DB cluster.
type DbClusterBackup struct {

	// The OCID of the shared-storage DB cluster backup.
	Id *string `mandatory:"true" json:"id"`

	// The type of backup.
	BackupType DbClusterBackupBackupTypeEnum `mandatory:"true" json:"backupType"`

	// The OCID of the compartment that contains the shared-storage DB cluster backup.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Name of the shared-storage DB cluster backup.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current lifecycle state of the backup.
	LifecycleState DbClusterBackupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	Source DbClusterBackupSource `mandatory:"true" json:"source"`

	// The date and time the backup was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the backup was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The size of the backup in GiBs.
	BackupSizeInGBs *int `mandatory:"false" json:"backupSizeInGBs"`

	DbClusterSnapshot *DbClusterSnapshot `mandatory:"false" json:"dbClusterSnapshot"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Description of the shared-storage DB cluster backup.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Additional information about the current lifecycle state of the backup.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The MySQL version of the shared-storage DB cluster at the time the backup was taken.
	MysqlVersion *string `mandatory:"false" json:"mysqlVersion"`

	// The number of days the backup will be retained.
	RetentionInDays *int `mandatory:"false" json:"retentionInDays"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DbClusterBackup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbClusterBackup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbClusterBackupBackupTypeEnum(string(m.BackupType)); !ok && m.BackupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupType: %s. Supported values are: %s.", m.BackupType, strings.Join(GetDbClusterBackupBackupTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbClusterBackupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbClusterBackupLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DbClusterBackup) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		BackupSizeInGBs   *int                              `json:"backupSizeInGBs"`
		DbClusterSnapshot *DbClusterSnapshot                `json:"dbClusterSnapshot"`
		DefinedTags       map[string]map[string]interface{} `json:"definedTags"`
		Description       *string                           `json:"description"`
		FreeformTags      map[string]string                 `json:"freeformTags"`
		LifecycleDetails  *string                           `json:"lifecycleDetails"`
		MysqlVersion      *string                           `json:"mysqlVersion"`
		RetentionInDays   *int                              `json:"retentionInDays"`
		SystemTags        map[string]map[string]interface{} `json:"systemTags"`
		Id                *string                           `json:"id"`
		BackupType        DbClusterBackupBackupTypeEnum     `json:"backupType"`
		CompartmentId     *string                           `json:"compartmentId"`
		DisplayName       *string                           `json:"displayName"`
		LifecycleState    DbClusterBackupLifecycleStateEnum `json:"lifecycleState"`
		Source            dbclusterbackupsource             `json:"source"`
		TimeCreated       *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated       *common.SDKTime                   `json:"timeUpdated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.BackupSizeInGBs = model.BackupSizeInGBs

	m.DbClusterSnapshot = model.DbClusterSnapshot

	m.DefinedTags = model.DefinedTags

	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.LifecycleDetails = model.LifecycleDetails

	m.MysqlVersion = model.MysqlVersion

	m.RetentionInDays = model.RetentionInDays

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.BackupType = model.BackupType

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.LifecycleState = model.LifecycleState

	nn, e = model.Source.UnmarshalPolymorphicJSON(model.Source.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Source = nn.(DbClusterBackupSource)
	} else {
		m.Source = nil
	}

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	return
}

// DbClusterBackupBackupTypeEnum Enum with underlying type: string
type DbClusterBackupBackupTypeEnum string

// Set of constants representing the allowable values for DbClusterBackupBackupTypeEnum
const (
	DbClusterBackupBackupTypeFull DbClusterBackupBackupTypeEnum = "FULL"
)

var mappingDbClusterBackupBackupTypeEnum = map[string]DbClusterBackupBackupTypeEnum{
	"FULL": DbClusterBackupBackupTypeFull,
}

var mappingDbClusterBackupBackupTypeEnumLowerCase = map[string]DbClusterBackupBackupTypeEnum{
	"full": DbClusterBackupBackupTypeFull,
}

// GetDbClusterBackupBackupTypeEnumValues Enumerates the set of values for DbClusterBackupBackupTypeEnum
func GetDbClusterBackupBackupTypeEnumValues() []DbClusterBackupBackupTypeEnum {
	values := make([]DbClusterBackupBackupTypeEnum, 0)
	for _, v := range mappingDbClusterBackupBackupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDbClusterBackupBackupTypeEnumStringValues Enumerates the set of values in String for DbClusterBackupBackupTypeEnum
func GetDbClusterBackupBackupTypeEnumStringValues() []string {
	return []string{
		"FULL",
	}
}

// GetMappingDbClusterBackupBackupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbClusterBackupBackupTypeEnum(val string) (DbClusterBackupBackupTypeEnum, bool) {
	enum, ok := mappingDbClusterBackupBackupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DbClusterBackupLifecycleStateEnum Enum with underlying type: string
type DbClusterBackupLifecycleStateEnum string

// Set of constants representing the allowable values for DbClusterBackupLifecycleStateEnum
const (
	DbClusterBackupLifecycleStateCreating DbClusterBackupLifecycleStateEnum = "CREATING"
	DbClusterBackupLifecycleStateActive   DbClusterBackupLifecycleStateEnum = "ACTIVE"
	DbClusterBackupLifecycleStateUpdating DbClusterBackupLifecycleStateEnum = "UPDATING"
	DbClusterBackupLifecycleStateDeleting DbClusterBackupLifecycleStateEnum = "DELETING"
	DbClusterBackupLifecycleStateDeleted  DbClusterBackupLifecycleStateEnum = "DELETED"
	DbClusterBackupLifecycleStateFailed   DbClusterBackupLifecycleStateEnum = "FAILED"
)

var mappingDbClusterBackupLifecycleStateEnum = map[string]DbClusterBackupLifecycleStateEnum{
	"CREATING": DbClusterBackupLifecycleStateCreating,
	"ACTIVE":   DbClusterBackupLifecycleStateActive,
	"UPDATING": DbClusterBackupLifecycleStateUpdating,
	"DELETING": DbClusterBackupLifecycleStateDeleting,
	"DELETED":  DbClusterBackupLifecycleStateDeleted,
	"FAILED":   DbClusterBackupLifecycleStateFailed,
}

var mappingDbClusterBackupLifecycleStateEnumLowerCase = map[string]DbClusterBackupLifecycleStateEnum{
	"creating": DbClusterBackupLifecycleStateCreating,
	"active":   DbClusterBackupLifecycleStateActive,
	"updating": DbClusterBackupLifecycleStateUpdating,
	"deleting": DbClusterBackupLifecycleStateDeleting,
	"deleted":  DbClusterBackupLifecycleStateDeleted,
	"failed":   DbClusterBackupLifecycleStateFailed,
}

// GetDbClusterBackupLifecycleStateEnumValues Enumerates the set of values for DbClusterBackupLifecycleStateEnum
func GetDbClusterBackupLifecycleStateEnumValues() []DbClusterBackupLifecycleStateEnum {
	values := make([]DbClusterBackupLifecycleStateEnum, 0)
	for _, v := range mappingDbClusterBackupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDbClusterBackupLifecycleStateEnumStringValues Enumerates the set of values in String for DbClusterBackupLifecycleStateEnum
func GetDbClusterBackupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingDbClusterBackupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbClusterBackupLifecycleStateEnum(val string) (DbClusterBackupLifecycleStateEnum, bool) {
	enum, ok := mappingDbClusterBackupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
