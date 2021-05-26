// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// MaintenanceRunSummary Details of a maintenance run.
type MaintenanceRunSummary struct {

	// The OCID of the maintenance run.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the maintenance run.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the maintenance run. For Autonomous Database on shared Exadata infrastructure, valid states are IN_PROGRESS, SUCCEEDED and FAILED.
	LifecycleState MaintenanceRunSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the maintenance run is scheduled to occur.
	TimeScheduled *common.SDKTime `mandatory:"true" json:"timeScheduled"`

	// Description of the maintenance run.
	Description *string `mandatory:"false" json:"description"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the maintenance run starts.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the maintenance run was completed.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// The type of the target resource on which the maintenance run occurs.
	TargetResourceType MaintenanceRunSummaryTargetResourceTypeEnum `mandatory:"false" json:"targetResourceType,omitempty"`

	// The ID of the target resource on which the maintenance run occurs.
	TargetResourceId *string `mandatory:"false" json:"targetResourceId"`

	// Maintenance type.
	MaintenanceType MaintenanceRunSummaryMaintenanceTypeEnum `mandatory:"false" json:"maintenanceType,omitempty"`

	// The unique identifier of the patch. The identifier string includes the patch type, the Oracle Database version, and the patch creation date (using the format YYMMDD). For example, the identifier `ru_patch_19.9.0.0_201030` is used for an RU patch for Oracle Database 19.9.0.0 that was released October 30, 2020.
	PatchId *string `mandatory:"false" json:"patchId"`

	// Maintenance sub-type.
	MaintenanceSubtype MaintenanceRunSummaryMaintenanceSubtypeEnum `mandatory:"false" json:"maintenanceSubtype,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance run for the Autonomous Data Guard association's peer container database.
	PeerMaintenanceRunId *string `mandatory:"false" json:"peerMaintenanceRunId"`

	// Maintenance method, it will be either "ROLLING" or "NONROLLING". Default value is ROLLING.
	PatchingMode MaintenanceRunSummaryPatchingModeEnum `mandatory:"false" json:"patchingMode,omitempty"`

	// Contain the patch failure count.
	PatchFailureCount *int `mandatory:"false" json:"patchFailureCount"`
}

func (m MaintenanceRunSummary) String() string {
	return common.PointerString(m)
}

// MaintenanceRunSummaryLifecycleStateEnum Enum with underlying type: string
type MaintenanceRunSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for MaintenanceRunSummaryLifecycleStateEnum
const (
	MaintenanceRunSummaryLifecycleStateScheduled  MaintenanceRunSummaryLifecycleStateEnum = "SCHEDULED"
	MaintenanceRunSummaryLifecycleStateInProgress MaintenanceRunSummaryLifecycleStateEnum = "IN_PROGRESS"
	MaintenanceRunSummaryLifecycleStateSucceeded  MaintenanceRunSummaryLifecycleStateEnum = "SUCCEEDED"
	MaintenanceRunSummaryLifecycleStateSkipped    MaintenanceRunSummaryLifecycleStateEnum = "SKIPPED"
	MaintenanceRunSummaryLifecycleStateFailed     MaintenanceRunSummaryLifecycleStateEnum = "FAILED"
	MaintenanceRunSummaryLifecycleStateUpdating   MaintenanceRunSummaryLifecycleStateEnum = "UPDATING"
	MaintenanceRunSummaryLifecycleStateDeleting   MaintenanceRunSummaryLifecycleStateEnum = "DELETING"
	MaintenanceRunSummaryLifecycleStateDeleted    MaintenanceRunSummaryLifecycleStateEnum = "DELETED"
	MaintenanceRunSummaryLifecycleStateCanceled   MaintenanceRunSummaryLifecycleStateEnum = "CANCELED"
)

var mappingMaintenanceRunSummaryLifecycleState = map[string]MaintenanceRunSummaryLifecycleStateEnum{
	"SCHEDULED":   MaintenanceRunSummaryLifecycleStateScheduled,
	"IN_PROGRESS": MaintenanceRunSummaryLifecycleStateInProgress,
	"SUCCEEDED":   MaintenanceRunSummaryLifecycleStateSucceeded,
	"SKIPPED":     MaintenanceRunSummaryLifecycleStateSkipped,
	"FAILED":      MaintenanceRunSummaryLifecycleStateFailed,
	"UPDATING":    MaintenanceRunSummaryLifecycleStateUpdating,
	"DELETING":    MaintenanceRunSummaryLifecycleStateDeleting,
	"DELETED":     MaintenanceRunSummaryLifecycleStateDeleted,
	"CANCELED":    MaintenanceRunSummaryLifecycleStateCanceled,
}

// GetMaintenanceRunSummaryLifecycleStateEnumValues Enumerates the set of values for MaintenanceRunSummaryLifecycleStateEnum
func GetMaintenanceRunSummaryLifecycleStateEnumValues() []MaintenanceRunSummaryLifecycleStateEnum {
	values := make([]MaintenanceRunSummaryLifecycleStateEnum, 0)
	for _, v := range mappingMaintenanceRunSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}

// MaintenanceRunSummaryTargetResourceTypeEnum Enum with underlying type: string
type MaintenanceRunSummaryTargetResourceTypeEnum string

// Set of constants representing the allowable values for MaintenanceRunSummaryTargetResourceTypeEnum
const (
	MaintenanceRunSummaryTargetResourceTypeAutonomousExadataInfrastructure MaintenanceRunSummaryTargetResourceTypeEnum = "AUTONOMOUS_EXADATA_INFRASTRUCTURE"
	MaintenanceRunSummaryTargetResourceTypeAutonomousContainerDatabase     MaintenanceRunSummaryTargetResourceTypeEnum = "AUTONOMOUS_CONTAINER_DATABASE"
	MaintenanceRunSummaryTargetResourceTypeExadataDbSystem                 MaintenanceRunSummaryTargetResourceTypeEnum = "EXADATA_DB_SYSTEM"
	MaintenanceRunSummaryTargetResourceTypeCloudExadataInfrastructure      MaintenanceRunSummaryTargetResourceTypeEnum = "CLOUD_EXADATA_INFRASTRUCTURE"
	MaintenanceRunSummaryTargetResourceTypeExaccInfrastructure             MaintenanceRunSummaryTargetResourceTypeEnum = "EXACC_INFRASTRUCTURE"
	MaintenanceRunSummaryTargetResourceTypeAutonomousDatabase              MaintenanceRunSummaryTargetResourceTypeEnum = "AUTONOMOUS_DATABASE"
)

var mappingMaintenanceRunSummaryTargetResourceType = map[string]MaintenanceRunSummaryTargetResourceTypeEnum{
	"AUTONOMOUS_EXADATA_INFRASTRUCTURE": MaintenanceRunSummaryTargetResourceTypeAutonomousExadataInfrastructure,
	"AUTONOMOUS_CONTAINER_DATABASE":     MaintenanceRunSummaryTargetResourceTypeAutonomousContainerDatabase,
	"EXADATA_DB_SYSTEM":                 MaintenanceRunSummaryTargetResourceTypeExadataDbSystem,
	"CLOUD_EXADATA_INFRASTRUCTURE":      MaintenanceRunSummaryTargetResourceTypeCloudExadataInfrastructure,
	"EXACC_INFRASTRUCTURE":              MaintenanceRunSummaryTargetResourceTypeExaccInfrastructure,
	"AUTONOMOUS_DATABASE":               MaintenanceRunSummaryTargetResourceTypeAutonomousDatabase,
}

// GetMaintenanceRunSummaryTargetResourceTypeEnumValues Enumerates the set of values for MaintenanceRunSummaryTargetResourceTypeEnum
func GetMaintenanceRunSummaryTargetResourceTypeEnumValues() []MaintenanceRunSummaryTargetResourceTypeEnum {
	values := make([]MaintenanceRunSummaryTargetResourceTypeEnum, 0)
	for _, v := range mappingMaintenanceRunSummaryTargetResourceType {
		values = append(values, v)
	}
	return values
}

// MaintenanceRunSummaryMaintenanceTypeEnum Enum with underlying type: string
type MaintenanceRunSummaryMaintenanceTypeEnum string

// Set of constants representing the allowable values for MaintenanceRunSummaryMaintenanceTypeEnum
const (
	MaintenanceRunSummaryMaintenanceTypePlanned   MaintenanceRunSummaryMaintenanceTypeEnum = "PLANNED"
	MaintenanceRunSummaryMaintenanceTypeUnplanned MaintenanceRunSummaryMaintenanceTypeEnum = "UNPLANNED"
)

var mappingMaintenanceRunSummaryMaintenanceType = map[string]MaintenanceRunSummaryMaintenanceTypeEnum{
	"PLANNED":   MaintenanceRunSummaryMaintenanceTypePlanned,
	"UNPLANNED": MaintenanceRunSummaryMaintenanceTypeUnplanned,
}

// GetMaintenanceRunSummaryMaintenanceTypeEnumValues Enumerates the set of values for MaintenanceRunSummaryMaintenanceTypeEnum
func GetMaintenanceRunSummaryMaintenanceTypeEnumValues() []MaintenanceRunSummaryMaintenanceTypeEnum {
	values := make([]MaintenanceRunSummaryMaintenanceTypeEnum, 0)
	for _, v := range mappingMaintenanceRunSummaryMaintenanceType {
		values = append(values, v)
	}
	return values
}

// MaintenanceRunSummaryMaintenanceSubtypeEnum Enum with underlying type: string
type MaintenanceRunSummaryMaintenanceSubtypeEnum string

// Set of constants representing the allowable values for MaintenanceRunSummaryMaintenanceSubtypeEnum
const (
	MaintenanceRunSummaryMaintenanceSubtypeQuarterly      MaintenanceRunSummaryMaintenanceSubtypeEnum = "QUARTERLY"
	MaintenanceRunSummaryMaintenanceSubtypeHardware       MaintenanceRunSummaryMaintenanceSubtypeEnum = "HARDWARE"
	MaintenanceRunSummaryMaintenanceSubtypeCritical       MaintenanceRunSummaryMaintenanceSubtypeEnum = "CRITICAL"
	MaintenanceRunSummaryMaintenanceSubtypeInfrastructure MaintenanceRunSummaryMaintenanceSubtypeEnum = "INFRASTRUCTURE"
	MaintenanceRunSummaryMaintenanceSubtypeDatabase       MaintenanceRunSummaryMaintenanceSubtypeEnum = "DATABASE"
	MaintenanceRunSummaryMaintenanceSubtypeOneoff         MaintenanceRunSummaryMaintenanceSubtypeEnum = "ONEOFF"
)

var mappingMaintenanceRunSummaryMaintenanceSubtype = map[string]MaintenanceRunSummaryMaintenanceSubtypeEnum{
	"QUARTERLY":      MaintenanceRunSummaryMaintenanceSubtypeQuarterly,
	"HARDWARE":       MaintenanceRunSummaryMaintenanceSubtypeHardware,
	"CRITICAL":       MaintenanceRunSummaryMaintenanceSubtypeCritical,
	"INFRASTRUCTURE": MaintenanceRunSummaryMaintenanceSubtypeInfrastructure,
	"DATABASE":       MaintenanceRunSummaryMaintenanceSubtypeDatabase,
	"ONEOFF":         MaintenanceRunSummaryMaintenanceSubtypeOneoff,
}

// GetMaintenanceRunSummaryMaintenanceSubtypeEnumValues Enumerates the set of values for MaintenanceRunSummaryMaintenanceSubtypeEnum
func GetMaintenanceRunSummaryMaintenanceSubtypeEnumValues() []MaintenanceRunSummaryMaintenanceSubtypeEnum {
	values := make([]MaintenanceRunSummaryMaintenanceSubtypeEnum, 0)
	for _, v := range mappingMaintenanceRunSummaryMaintenanceSubtype {
		values = append(values, v)
	}
	return values
}

// MaintenanceRunSummaryPatchingModeEnum Enum with underlying type: string
type MaintenanceRunSummaryPatchingModeEnum string

// Set of constants representing the allowable values for MaintenanceRunSummaryPatchingModeEnum
const (
	MaintenanceRunSummaryPatchingModeRolling    MaintenanceRunSummaryPatchingModeEnum = "ROLLING"
	MaintenanceRunSummaryPatchingModeNonrolling MaintenanceRunSummaryPatchingModeEnum = "NONROLLING"
)

var mappingMaintenanceRunSummaryPatchingMode = map[string]MaintenanceRunSummaryPatchingModeEnum{
	"ROLLING":    MaintenanceRunSummaryPatchingModeRolling,
	"NONROLLING": MaintenanceRunSummaryPatchingModeNonrolling,
}

// GetMaintenanceRunSummaryPatchingModeEnumValues Enumerates the set of values for MaintenanceRunSummaryPatchingModeEnum
func GetMaintenanceRunSummaryPatchingModeEnumValues() []MaintenanceRunSummaryPatchingModeEnum {
	values := make([]MaintenanceRunSummaryPatchingModeEnum, 0)
	for _, v := range mappingMaintenanceRunSummaryPatchingMode {
		values = append(values, v)
	}
	return values
}
