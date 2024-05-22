// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NodeBackup The information about the Node's backup.
type NodeBackup struct {

	// The id of the node backup.
	Id *string `mandatory:"true" json:"id"`

	// BDS generated name for the backup. Format is nodeHostName_timeCreated
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The instance OCID of the node, which is the resource from which the node backup was acquired.
	NodeInstanceId *string `mandatory:"true" json:"nodeInstanceId"`

	// Host name of the node to which this backup belongs.
	NodeHostName *string `mandatory:"true" json:"nodeHostName"`

	// type based on how backup action was initiated.
	BackupTriggerType NodeBackupBackupTriggerTypeEnum `mandatory:"true" json:"backupTriggerType"`

	// Incremental backup type includes only the changes since the last backup. Full backup type includes all changes since the volume was created.
	BackupType NodeBackupBackupTypeEnum `mandatory:"true" json:"backupType"`

	// The state of the NodeBackup.
	LifecycleState NodeBackupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the cluster was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The ID of the nodeBackupConfiguration if the NodeBackup is automatically created by applying the configuration.
	NodeBackupConfigId *string `mandatory:"false" json:"nodeBackupConfigId"`
}

func (m NodeBackup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NodeBackup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNodeBackupBackupTriggerTypeEnum(string(m.BackupTriggerType)); !ok && m.BackupTriggerType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupTriggerType: %s. Supported values are: %s.", m.BackupTriggerType, strings.Join(GetNodeBackupBackupTriggerTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNodeBackupBackupTypeEnum(string(m.BackupType)); !ok && m.BackupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupType: %s. Supported values are: %s.", m.BackupType, strings.Join(GetNodeBackupBackupTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNodeBackupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNodeBackupLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NodeBackupBackupTriggerTypeEnum Enum with underlying type: string
type NodeBackupBackupTriggerTypeEnum string

// Set of constants representing the allowable values for NodeBackupBackupTriggerTypeEnum
const (
	NodeBackupBackupTriggerTypeManual    NodeBackupBackupTriggerTypeEnum = "MANUAL"
	NodeBackupBackupTriggerTypeScheduled NodeBackupBackupTriggerTypeEnum = "SCHEDULED"
)

var mappingNodeBackupBackupTriggerTypeEnum = map[string]NodeBackupBackupTriggerTypeEnum{
	"MANUAL":    NodeBackupBackupTriggerTypeManual,
	"SCHEDULED": NodeBackupBackupTriggerTypeScheduled,
}

var mappingNodeBackupBackupTriggerTypeEnumLowerCase = map[string]NodeBackupBackupTriggerTypeEnum{
	"manual":    NodeBackupBackupTriggerTypeManual,
	"scheduled": NodeBackupBackupTriggerTypeScheduled,
}

// GetNodeBackupBackupTriggerTypeEnumValues Enumerates the set of values for NodeBackupBackupTriggerTypeEnum
func GetNodeBackupBackupTriggerTypeEnumValues() []NodeBackupBackupTriggerTypeEnum {
	values := make([]NodeBackupBackupTriggerTypeEnum, 0)
	for _, v := range mappingNodeBackupBackupTriggerTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNodeBackupBackupTriggerTypeEnumStringValues Enumerates the set of values in String for NodeBackupBackupTriggerTypeEnum
func GetNodeBackupBackupTriggerTypeEnumStringValues() []string {
	return []string{
		"MANUAL",
		"SCHEDULED",
	}
}

// GetMappingNodeBackupBackupTriggerTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNodeBackupBackupTriggerTypeEnum(val string) (NodeBackupBackupTriggerTypeEnum, bool) {
	enum, ok := mappingNodeBackupBackupTriggerTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// NodeBackupBackupTypeEnum Enum with underlying type: string
type NodeBackupBackupTypeEnum string

// Set of constants representing the allowable values for NodeBackupBackupTypeEnum
const (
	NodeBackupBackupTypeFull        NodeBackupBackupTypeEnum = "FULL"
	NodeBackupBackupTypeIncremental NodeBackupBackupTypeEnum = "INCREMENTAL"
)

var mappingNodeBackupBackupTypeEnum = map[string]NodeBackupBackupTypeEnum{
	"FULL":        NodeBackupBackupTypeFull,
	"INCREMENTAL": NodeBackupBackupTypeIncremental,
}

var mappingNodeBackupBackupTypeEnumLowerCase = map[string]NodeBackupBackupTypeEnum{
	"full":        NodeBackupBackupTypeFull,
	"incremental": NodeBackupBackupTypeIncremental,
}

// GetNodeBackupBackupTypeEnumValues Enumerates the set of values for NodeBackupBackupTypeEnum
func GetNodeBackupBackupTypeEnumValues() []NodeBackupBackupTypeEnum {
	values := make([]NodeBackupBackupTypeEnum, 0)
	for _, v := range mappingNodeBackupBackupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNodeBackupBackupTypeEnumStringValues Enumerates the set of values in String for NodeBackupBackupTypeEnum
func GetNodeBackupBackupTypeEnumStringValues() []string {
	return []string{
		"FULL",
		"INCREMENTAL",
	}
}

// GetMappingNodeBackupBackupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNodeBackupBackupTypeEnum(val string) (NodeBackupBackupTypeEnum, bool) {
	enum, ok := mappingNodeBackupBackupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// NodeBackupLifecycleStateEnum Enum with underlying type: string
type NodeBackupLifecycleStateEnum string

// Set of constants representing the allowable values for NodeBackupLifecycleStateEnum
const (
	NodeBackupLifecycleStateCreating NodeBackupLifecycleStateEnum = "CREATING"
	NodeBackupLifecycleStateActive   NodeBackupLifecycleStateEnum = "ACTIVE"
	NodeBackupLifecycleStateUpdating NodeBackupLifecycleStateEnum = "UPDATING"
	NodeBackupLifecycleStateDeleting NodeBackupLifecycleStateEnum = "DELETING"
	NodeBackupLifecycleStateDeleted  NodeBackupLifecycleStateEnum = "DELETED"
	NodeBackupLifecycleStateFailed   NodeBackupLifecycleStateEnum = "FAILED"
	NodeBackupLifecycleStatePartial  NodeBackupLifecycleStateEnum = "PARTIAL"
)

var mappingNodeBackupLifecycleStateEnum = map[string]NodeBackupLifecycleStateEnum{
	"CREATING": NodeBackupLifecycleStateCreating,
	"ACTIVE":   NodeBackupLifecycleStateActive,
	"UPDATING": NodeBackupLifecycleStateUpdating,
	"DELETING": NodeBackupLifecycleStateDeleting,
	"DELETED":  NodeBackupLifecycleStateDeleted,
	"FAILED":   NodeBackupLifecycleStateFailed,
	"PARTIAL":  NodeBackupLifecycleStatePartial,
}

var mappingNodeBackupLifecycleStateEnumLowerCase = map[string]NodeBackupLifecycleStateEnum{
	"creating": NodeBackupLifecycleStateCreating,
	"active":   NodeBackupLifecycleStateActive,
	"updating": NodeBackupLifecycleStateUpdating,
	"deleting": NodeBackupLifecycleStateDeleting,
	"deleted":  NodeBackupLifecycleStateDeleted,
	"failed":   NodeBackupLifecycleStateFailed,
	"partial":  NodeBackupLifecycleStatePartial,
}

// GetNodeBackupLifecycleStateEnumValues Enumerates the set of values for NodeBackupLifecycleStateEnum
func GetNodeBackupLifecycleStateEnumValues() []NodeBackupLifecycleStateEnum {
	values := make([]NodeBackupLifecycleStateEnum, 0)
	for _, v := range mappingNodeBackupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetNodeBackupLifecycleStateEnumStringValues Enumerates the set of values in String for NodeBackupLifecycleStateEnum
func GetNodeBackupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"PARTIAL",
	}
}

// GetMappingNodeBackupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNodeBackupLifecycleStateEnum(val string) (NodeBackupLifecycleStateEnum, bool) {
	enum, ok := mappingNodeBackupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
