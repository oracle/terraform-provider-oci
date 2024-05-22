// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NodeBackupConfiguration The information about the NodeBackupConfiguration.
type NodeBackupConfiguration struct {

	// The unique identifier for the NodeBackupConfiguration.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the bdsInstance which is the parent resource id.
	BdsInstanceId *string `mandatory:"true" json:"bdsInstanceId"`

	// A user-friendly name. Only ASCII alphanumeric characters with no spaces allowed. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	LevelTypeDetails LevelTypeDetails `mandatory:"true" json:"levelTypeDetails"`

	// The state of the NodeBackupConfiguration.
	LifecycleState NodeBackupConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the NodeBackupConfiguration was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the NodeBackupConfiguration was updated, shown as an RFC 3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The time zone of the execution schedule, in IANA time zone database name format
	Timezone *string `mandatory:"true" json:"timezone"`

	// Day/time recurrence (specified following RFC 5545) at which to trigger the backup process. Currently only DAILY, WEEKLY and MONTHLY frequency is supported. Days of the week are specified using BYDAY field. Time of the day is specified using BYHOUR. Other fields are not supported.
	Schedule *string `mandatory:"true" json:"schedule"`

	// Number of backup copies to retain.
	NumberOfBackupsToRetain *int `mandatory:"true" json:"numberOfBackupsToRetain"`

	// Incremental backup type includes only the changes since the last backup. Full backup type includes all changes since the volume was created.
	BackupType NodeBackupBackupTypeEnum `mandatory:"false" json:"backupType,omitempty"`
}

func (m NodeBackupConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NodeBackupConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNodeBackupConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNodeBackupConfigurationLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingNodeBackupBackupTypeEnum(string(m.BackupType)); !ok && m.BackupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupType: %s. Supported values are: %s.", m.BackupType, strings.Join(GetNodeBackupBackupTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *NodeBackupConfiguration) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		BackupType              NodeBackupBackupTypeEnum                  `json:"backupType"`
		Id                      *string                                   `json:"id"`
		BdsInstanceId           *string                                   `json:"bdsInstanceId"`
		DisplayName             *string                                   `json:"displayName"`
		LevelTypeDetails        leveltypedetails                          `json:"levelTypeDetails"`
		LifecycleState          NodeBackupConfigurationLifecycleStateEnum `json:"lifecycleState"`
		TimeCreated             *common.SDKTime                           `json:"timeCreated"`
		TimeUpdated             *common.SDKTime                           `json:"timeUpdated"`
		Timezone                *string                                   `json:"timezone"`
		Schedule                *string                                   `json:"schedule"`
		NumberOfBackupsToRetain *int                                      `json:"numberOfBackupsToRetain"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.BackupType = model.BackupType

	m.Id = model.Id

	m.BdsInstanceId = model.BdsInstanceId

	m.DisplayName = model.DisplayName

	nn, e = model.LevelTypeDetails.UnmarshalPolymorphicJSON(model.LevelTypeDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.LevelTypeDetails = nn.(LevelTypeDetails)
	} else {
		m.LevelTypeDetails = nil
	}

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.Timezone = model.Timezone

	m.Schedule = model.Schedule

	m.NumberOfBackupsToRetain = model.NumberOfBackupsToRetain

	return
}

// NodeBackupConfigurationLifecycleStateEnum Enum with underlying type: string
type NodeBackupConfigurationLifecycleStateEnum string

// Set of constants representing the allowable values for NodeBackupConfigurationLifecycleStateEnum
const (
	NodeBackupConfigurationLifecycleStateCreating NodeBackupConfigurationLifecycleStateEnum = "CREATING"
	NodeBackupConfigurationLifecycleStateActive   NodeBackupConfigurationLifecycleStateEnum = "ACTIVE"
	NodeBackupConfigurationLifecycleStateUpdating NodeBackupConfigurationLifecycleStateEnum = "UPDATING"
	NodeBackupConfigurationLifecycleStateDeleting NodeBackupConfigurationLifecycleStateEnum = "DELETING"
	NodeBackupConfigurationLifecycleStateDeleted  NodeBackupConfigurationLifecycleStateEnum = "DELETED"
	NodeBackupConfigurationLifecycleStateFailed   NodeBackupConfigurationLifecycleStateEnum = "FAILED"
)

var mappingNodeBackupConfigurationLifecycleStateEnum = map[string]NodeBackupConfigurationLifecycleStateEnum{
	"CREATING": NodeBackupConfigurationLifecycleStateCreating,
	"ACTIVE":   NodeBackupConfigurationLifecycleStateActive,
	"UPDATING": NodeBackupConfigurationLifecycleStateUpdating,
	"DELETING": NodeBackupConfigurationLifecycleStateDeleting,
	"DELETED":  NodeBackupConfigurationLifecycleStateDeleted,
	"FAILED":   NodeBackupConfigurationLifecycleStateFailed,
}

var mappingNodeBackupConfigurationLifecycleStateEnumLowerCase = map[string]NodeBackupConfigurationLifecycleStateEnum{
	"creating": NodeBackupConfigurationLifecycleStateCreating,
	"active":   NodeBackupConfigurationLifecycleStateActive,
	"updating": NodeBackupConfigurationLifecycleStateUpdating,
	"deleting": NodeBackupConfigurationLifecycleStateDeleting,
	"deleted":  NodeBackupConfigurationLifecycleStateDeleted,
	"failed":   NodeBackupConfigurationLifecycleStateFailed,
}

// GetNodeBackupConfigurationLifecycleStateEnumValues Enumerates the set of values for NodeBackupConfigurationLifecycleStateEnum
func GetNodeBackupConfigurationLifecycleStateEnumValues() []NodeBackupConfigurationLifecycleStateEnum {
	values := make([]NodeBackupConfigurationLifecycleStateEnum, 0)
	for _, v := range mappingNodeBackupConfigurationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetNodeBackupConfigurationLifecycleStateEnumStringValues Enumerates the set of values in String for NodeBackupConfigurationLifecycleStateEnum
func GetNodeBackupConfigurationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingNodeBackupConfigurationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNodeBackupConfigurationLifecycleStateEnum(val string) (NodeBackupConfigurationLifecycleStateEnum, bool) {
	enum, ok := mappingNodeBackupConfigurationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
