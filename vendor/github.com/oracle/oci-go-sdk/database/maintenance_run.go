// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// MaintenanceRun Details of a maintenance run.
type MaintenanceRun struct {

	// The OCID of the maintenance run.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the maintenance run.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the maintenance run.
	LifecycleState MaintenanceRunLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the maintenance run is scheduled to occur.
	TimeScheduled *common.SDKTime `mandatory:"true" json:"timeScheduled"`

	// Description of the maintenance run.
	Description *string `mandatory:"false" json:"description"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the maintenance run starts.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the maintenance run was completed.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// The type of the target resource on which the maintenance run occurs.
	TargetResourceType MaintenanceRunTargetResourceTypeEnum `mandatory:"false" json:"targetResourceType,omitempty"`

	// The ID of the target resource on which the maintenance run occurs.
	TargetResourceId *string `mandatory:"false" json:"targetResourceId"`

	// Maintenance type.
	MaintenanceType MaintenanceRunMaintenanceTypeEnum `mandatory:"false" json:"maintenanceType,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the patch to be applied in the maintenance run.
	PatchId *string `mandatory:"false" json:"patchId"`

	// Maintenance sub-type.
	MaintenanceSubtype MaintenanceRunMaintenanceSubtypeEnum `mandatory:"false" json:"maintenanceSubtype,omitempty"`
}

func (m MaintenanceRun) String() string {
	return common.PointerString(m)
}

// MaintenanceRunLifecycleStateEnum Enum with underlying type: string
type MaintenanceRunLifecycleStateEnum string

// Set of constants representing the allowable values for MaintenanceRunLifecycleStateEnum
const (
	MaintenanceRunLifecycleStateScheduled  MaintenanceRunLifecycleStateEnum = "SCHEDULED"
	MaintenanceRunLifecycleStateInProgress MaintenanceRunLifecycleStateEnum = "IN_PROGRESS"
	MaintenanceRunLifecycleStateSucceeded  MaintenanceRunLifecycleStateEnum = "SUCCEEDED"
	MaintenanceRunLifecycleStateSkipped    MaintenanceRunLifecycleStateEnum = "SKIPPED"
	MaintenanceRunLifecycleStateFailed     MaintenanceRunLifecycleStateEnum = "FAILED"
	MaintenanceRunLifecycleStateUpdating   MaintenanceRunLifecycleStateEnum = "UPDATING"
	MaintenanceRunLifecycleStateDeleting   MaintenanceRunLifecycleStateEnum = "DELETING"
	MaintenanceRunLifecycleStateDeleted    MaintenanceRunLifecycleStateEnum = "DELETED"
)

var mappingMaintenanceRunLifecycleState = map[string]MaintenanceRunLifecycleStateEnum{
	"SCHEDULED":   MaintenanceRunLifecycleStateScheduled,
	"IN_PROGRESS": MaintenanceRunLifecycleStateInProgress,
	"SUCCEEDED":   MaintenanceRunLifecycleStateSucceeded,
	"SKIPPED":     MaintenanceRunLifecycleStateSkipped,
	"FAILED":      MaintenanceRunLifecycleStateFailed,
	"UPDATING":    MaintenanceRunLifecycleStateUpdating,
	"DELETING":    MaintenanceRunLifecycleStateDeleting,
	"DELETED":     MaintenanceRunLifecycleStateDeleted,
}

// GetMaintenanceRunLifecycleStateEnumValues Enumerates the set of values for MaintenanceRunLifecycleStateEnum
func GetMaintenanceRunLifecycleStateEnumValues() []MaintenanceRunLifecycleStateEnum {
	values := make([]MaintenanceRunLifecycleStateEnum, 0)
	for _, v := range mappingMaintenanceRunLifecycleState {
		values = append(values, v)
	}
	return values
}

// MaintenanceRunTargetResourceTypeEnum Enum with underlying type: string
type MaintenanceRunTargetResourceTypeEnum string

// Set of constants representing the allowable values for MaintenanceRunTargetResourceTypeEnum
const (
	MaintenanceRunTargetResourceTypeAutonomousExadataInfrastructure MaintenanceRunTargetResourceTypeEnum = "AUTONOMOUS_EXADATA_INFRASTRUCTURE"
	MaintenanceRunTargetResourceTypeAutonomousContainerDatabase     MaintenanceRunTargetResourceTypeEnum = "AUTONOMOUS_CONTAINER_DATABASE"
	MaintenanceRunTargetResourceTypeExadataDbSystem                 MaintenanceRunTargetResourceTypeEnum = "EXADATA_DB_SYSTEM"
)

var mappingMaintenanceRunTargetResourceType = map[string]MaintenanceRunTargetResourceTypeEnum{
	"AUTONOMOUS_EXADATA_INFRASTRUCTURE": MaintenanceRunTargetResourceTypeAutonomousExadataInfrastructure,
	"AUTONOMOUS_CONTAINER_DATABASE":     MaintenanceRunTargetResourceTypeAutonomousContainerDatabase,
	"EXADATA_DB_SYSTEM":                 MaintenanceRunTargetResourceTypeExadataDbSystem,
}

// GetMaintenanceRunTargetResourceTypeEnumValues Enumerates the set of values for MaintenanceRunTargetResourceTypeEnum
func GetMaintenanceRunTargetResourceTypeEnumValues() []MaintenanceRunTargetResourceTypeEnum {
	values := make([]MaintenanceRunTargetResourceTypeEnum, 0)
	for _, v := range mappingMaintenanceRunTargetResourceType {
		values = append(values, v)
	}
	return values
}

// MaintenanceRunMaintenanceTypeEnum Enum with underlying type: string
type MaintenanceRunMaintenanceTypeEnum string

// Set of constants representing the allowable values for MaintenanceRunMaintenanceTypeEnum
const (
	MaintenanceRunMaintenanceTypePlanned   MaintenanceRunMaintenanceTypeEnum = "PLANNED"
	MaintenanceRunMaintenanceTypeUnplanned MaintenanceRunMaintenanceTypeEnum = "UNPLANNED"
)

var mappingMaintenanceRunMaintenanceType = map[string]MaintenanceRunMaintenanceTypeEnum{
	"PLANNED":   MaintenanceRunMaintenanceTypePlanned,
	"UNPLANNED": MaintenanceRunMaintenanceTypeUnplanned,
}

// GetMaintenanceRunMaintenanceTypeEnumValues Enumerates the set of values for MaintenanceRunMaintenanceTypeEnum
func GetMaintenanceRunMaintenanceTypeEnumValues() []MaintenanceRunMaintenanceTypeEnum {
	values := make([]MaintenanceRunMaintenanceTypeEnum, 0)
	for _, v := range mappingMaintenanceRunMaintenanceType {
		values = append(values, v)
	}
	return values
}

// MaintenanceRunMaintenanceSubtypeEnum Enum with underlying type: string
type MaintenanceRunMaintenanceSubtypeEnum string

// Set of constants representing the allowable values for MaintenanceRunMaintenanceSubtypeEnum
const (
	MaintenanceRunMaintenanceSubtypeQuarterly MaintenanceRunMaintenanceSubtypeEnum = "QUARTERLY"
	MaintenanceRunMaintenanceSubtypeHardware  MaintenanceRunMaintenanceSubtypeEnum = "HARDWARE"
	MaintenanceRunMaintenanceSubtypeCritical  MaintenanceRunMaintenanceSubtypeEnum = "CRITICAL"
)

var mappingMaintenanceRunMaintenanceSubtype = map[string]MaintenanceRunMaintenanceSubtypeEnum{
	"QUARTERLY": MaintenanceRunMaintenanceSubtypeQuarterly,
	"HARDWARE":  MaintenanceRunMaintenanceSubtypeHardware,
	"CRITICAL":  MaintenanceRunMaintenanceSubtypeCritical,
}

// GetMaintenanceRunMaintenanceSubtypeEnumValues Enumerates the set of values for MaintenanceRunMaintenanceSubtypeEnum
func GetMaintenanceRunMaintenanceSubtypeEnumValues() []MaintenanceRunMaintenanceSubtypeEnum {
	values := make([]MaintenanceRunMaintenanceSubtypeEnum, 0)
	for _, v := range mappingMaintenanceRunMaintenanceSubtype {
		values = append(values, v)
	}
	return values
}
