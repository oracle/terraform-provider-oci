// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// MaintenanceRun Details of a maintenance run.
type MaintenanceRun struct {

	// The OCID of the maintenance run.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the maintenance run.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the maintenance run. For Autonomous Database on shared Exadata infrastructure, valid states are IN_PROGRESS, SUCCEEDED and FAILED.
	LifecycleState MaintenanceRunLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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
	TargetResourceType MaintenanceRunTargetResourceTypeEnum `mandatory:"false" json:"targetResourceType,omitempty"`

	// The ID of the target resource on which the maintenance run occurs.
	TargetResourceId *string `mandatory:"false" json:"targetResourceId"`

	// Maintenance type.
	MaintenanceType MaintenanceRunMaintenanceTypeEnum `mandatory:"false" json:"maintenanceType,omitempty"`

	// The unique identifier of the patch. The identifier string includes the patch type, the Oracle Database version, and the patch creation date (using the format YYMMDD). For example, the identifier `ru_patch_19.9.0.0_201030` is used for an RU patch for Oracle Database 19.9.0.0 that was released October 30, 2020.
	PatchId *string `mandatory:"false" json:"patchId"`

	// Maintenance sub-type.
	MaintenanceSubtype MaintenanceRunMaintenanceSubtypeEnum `mandatory:"false" json:"maintenanceSubtype,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance run for the Autonomous Data Guard association's peer container database.
	PeerMaintenanceRunId *string `mandatory:"false" json:"peerMaintenanceRunId"`

	// Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.
	// *IMPORTANT*: Non-rolling infrastructure patching involves system down time. See Oracle-Managed Infrastructure Maintenance Updates (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle) for more information.
	PatchingMode MaintenanceRunPatchingModeEnum `mandatory:"false" json:"patchingMode,omitempty"`

	// Contain the patch failure count.
	PatchFailureCount *int `mandatory:"false" json:"patchFailureCount"`
}

func (m MaintenanceRun) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaintenanceRun) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMaintenanceRunLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMaintenanceRunLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingMaintenanceRunTargetResourceTypeEnum(string(m.TargetResourceType)); !ok && m.TargetResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetResourceType: %s. Supported values are: %s.", m.TargetResourceType, strings.Join(GetMaintenanceRunTargetResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMaintenanceRunMaintenanceTypeEnum(string(m.MaintenanceType)); !ok && m.MaintenanceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceType: %s. Supported values are: %s.", m.MaintenanceType, strings.Join(GetMaintenanceRunMaintenanceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMaintenanceRunMaintenanceSubtypeEnum(string(m.MaintenanceSubtype)); !ok && m.MaintenanceSubtype != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceSubtype: %s. Supported values are: %s.", m.MaintenanceSubtype, strings.Join(GetMaintenanceRunMaintenanceSubtypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMaintenanceRunPatchingModeEnum(string(m.PatchingMode)); !ok && m.PatchingMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchingMode: %s. Supported values are: %s.", m.PatchingMode, strings.Join(GetMaintenanceRunPatchingModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
	MaintenanceRunLifecycleStateCanceled   MaintenanceRunLifecycleStateEnum = "CANCELED"
)

var mappingMaintenanceRunLifecycleStateEnum = map[string]MaintenanceRunLifecycleStateEnum{
	"SCHEDULED":   MaintenanceRunLifecycleStateScheduled,
	"IN_PROGRESS": MaintenanceRunLifecycleStateInProgress,
	"SUCCEEDED":   MaintenanceRunLifecycleStateSucceeded,
	"SKIPPED":     MaintenanceRunLifecycleStateSkipped,
	"FAILED":      MaintenanceRunLifecycleStateFailed,
	"UPDATING":    MaintenanceRunLifecycleStateUpdating,
	"DELETING":    MaintenanceRunLifecycleStateDeleting,
	"DELETED":     MaintenanceRunLifecycleStateDeleted,
	"CANCELED":    MaintenanceRunLifecycleStateCanceled,
}

// GetMaintenanceRunLifecycleStateEnumValues Enumerates the set of values for MaintenanceRunLifecycleStateEnum
func GetMaintenanceRunLifecycleStateEnumValues() []MaintenanceRunLifecycleStateEnum {
	values := make([]MaintenanceRunLifecycleStateEnum, 0)
	for _, v := range mappingMaintenanceRunLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceRunLifecycleStateEnumStringValues Enumerates the set of values in String for MaintenanceRunLifecycleStateEnum
func GetMaintenanceRunLifecycleStateEnumStringValues() []string {
	return []string{
		"SCHEDULED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"SKIPPED",
		"FAILED",
		"UPDATING",
		"DELETING",
		"DELETED",
		"CANCELED",
	}
}

// GetMappingMaintenanceRunLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceRunLifecycleStateEnum(val string) (MaintenanceRunLifecycleStateEnum, bool) {
	mappingMaintenanceRunLifecycleStateEnumIgnoreCase := make(map[string]MaintenanceRunLifecycleStateEnum)
	for k, v := range mappingMaintenanceRunLifecycleStateEnum {
		mappingMaintenanceRunLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingMaintenanceRunLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// MaintenanceRunTargetResourceTypeEnum Enum with underlying type: string
type MaintenanceRunTargetResourceTypeEnum string

// Set of constants representing the allowable values for MaintenanceRunTargetResourceTypeEnum
const (
	MaintenanceRunTargetResourceTypeAutonomousExadataInfrastructure MaintenanceRunTargetResourceTypeEnum = "AUTONOMOUS_EXADATA_INFRASTRUCTURE"
	MaintenanceRunTargetResourceTypeAutonomousContainerDatabase     MaintenanceRunTargetResourceTypeEnum = "AUTONOMOUS_CONTAINER_DATABASE"
	MaintenanceRunTargetResourceTypeExadataDbSystem                 MaintenanceRunTargetResourceTypeEnum = "EXADATA_DB_SYSTEM"
	MaintenanceRunTargetResourceTypeCloudExadataInfrastructure      MaintenanceRunTargetResourceTypeEnum = "CLOUD_EXADATA_INFRASTRUCTURE"
	MaintenanceRunTargetResourceTypeExaccInfrastructure             MaintenanceRunTargetResourceTypeEnum = "EXACC_INFRASTRUCTURE"
	MaintenanceRunTargetResourceTypeAutonomousDatabase              MaintenanceRunTargetResourceTypeEnum = "AUTONOMOUS_DATABASE"
)

var mappingMaintenanceRunTargetResourceTypeEnum = map[string]MaintenanceRunTargetResourceTypeEnum{
	"AUTONOMOUS_EXADATA_INFRASTRUCTURE": MaintenanceRunTargetResourceTypeAutonomousExadataInfrastructure,
	"AUTONOMOUS_CONTAINER_DATABASE":     MaintenanceRunTargetResourceTypeAutonomousContainerDatabase,
	"EXADATA_DB_SYSTEM":                 MaintenanceRunTargetResourceTypeExadataDbSystem,
	"CLOUD_EXADATA_INFRASTRUCTURE":      MaintenanceRunTargetResourceTypeCloudExadataInfrastructure,
	"EXACC_INFRASTRUCTURE":              MaintenanceRunTargetResourceTypeExaccInfrastructure,
	"AUTONOMOUS_DATABASE":               MaintenanceRunTargetResourceTypeAutonomousDatabase,
}

// GetMaintenanceRunTargetResourceTypeEnumValues Enumerates the set of values for MaintenanceRunTargetResourceTypeEnum
func GetMaintenanceRunTargetResourceTypeEnumValues() []MaintenanceRunTargetResourceTypeEnum {
	values := make([]MaintenanceRunTargetResourceTypeEnum, 0)
	for _, v := range mappingMaintenanceRunTargetResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceRunTargetResourceTypeEnumStringValues Enumerates the set of values in String for MaintenanceRunTargetResourceTypeEnum
func GetMaintenanceRunTargetResourceTypeEnumStringValues() []string {
	return []string{
		"AUTONOMOUS_EXADATA_INFRASTRUCTURE",
		"AUTONOMOUS_CONTAINER_DATABASE",
		"EXADATA_DB_SYSTEM",
		"CLOUD_EXADATA_INFRASTRUCTURE",
		"EXACC_INFRASTRUCTURE",
		"AUTONOMOUS_DATABASE",
	}
}

// GetMappingMaintenanceRunTargetResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceRunTargetResourceTypeEnum(val string) (MaintenanceRunTargetResourceTypeEnum, bool) {
	mappingMaintenanceRunTargetResourceTypeEnumIgnoreCase := make(map[string]MaintenanceRunTargetResourceTypeEnum)
	for k, v := range mappingMaintenanceRunTargetResourceTypeEnum {
		mappingMaintenanceRunTargetResourceTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingMaintenanceRunTargetResourceTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// MaintenanceRunMaintenanceTypeEnum Enum with underlying type: string
type MaintenanceRunMaintenanceTypeEnum string

// Set of constants representing the allowable values for MaintenanceRunMaintenanceTypeEnum
const (
	MaintenanceRunMaintenanceTypePlanned   MaintenanceRunMaintenanceTypeEnum = "PLANNED"
	MaintenanceRunMaintenanceTypeUnplanned MaintenanceRunMaintenanceTypeEnum = "UNPLANNED"
)

var mappingMaintenanceRunMaintenanceTypeEnum = map[string]MaintenanceRunMaintenanceTypeEnum{
	"PLANNED":   MaintenanceRunMaintenanceTypePlanned,
	"UNPLANNED": MaintenanceRunMaintenanceTypeUnplanned,
}

// GetMaintenanceRunMaintenanceTypeEnumValues Enumerates the set of values for MaintenanceRunMaintenanceTypeEnum
func GetMaintenanceRunMaintenanceTypeEnumValues() []MaintenanceRunMaintenanceTypeEnum {
	values := make([]MaintenanceRunMaintenanceTypeEnum, 0)
	for _, v := range mappingMaintenanceRunMaintenanceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceRunMaintenanceTypeEnumStringValues Enumerates the set of values in String for MaintenanceRunMaintenanceTypeEnum
func GetMaintenanceRunMaintenanceTypeEnumStringValues() []string {
	return []string{
		"PLANNED",
		"UNPLANNED",
	}
}

// GetMappingMaintenanceRunMaintenanceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceRunMaintenanceTypeEnum(val string) (MaintenanceRunMaintenanceTypeEnum, bool) {
	mappingMaintenanceRunMaintenanceTypeEnumIgnoreCase := make(map[string]MaintenanceRunMaintenanceTypeEnum)
	for k, v := range mappingMaintenanceRunMaintenanceTypeEnum {
		mappingMaintenanceRunMaintenanceTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingMaintenanceRunMaintenanceTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// MaintenanceRunMaintenanceSubtypeEnum Enum with underlying type: string
type MaintenanceRunMaintenanceSubtypeEnum string

// Set of constants representing the allowable values for MaintenanceRunMaintenanceSubtypeEnum
const (
	MaintenanceRunMaintenanceSubtypeQuarterly      MaintenanceRunMaintenanceSubtypeEnum = "QUARTERLY"
	MaintenanceRunMaintenanceSubtypeHardware       MaintenanceRunMaintenanceSubtypeEnum = "HARDWARE"
	MaintenanceRunMaintenanceSubtypeCritical       MaintenanceRunMaintenanceSubtypeEnum = "CRITICAL"
	MaintenanceRunMaintenanceSubtypeInfrastructure MaintenanceRunMaintenanceSubtypeEnum = "INFRASTRUCTURE"
	MaintenanceRunMaintenanceSubtypeDatabase       MaintenanceRunMaintenanceSubtypeEnum = "DATABASE"
	MaintenanceRunMaintenanceSubtypeOneoff         MaintenanceRunMaintenanceSubtypeEnum = "ONEOFF"
)

var mappingMaintenanceRunMaintenanceSubtypeEnum = map[string]MaintenanceRunMaintenanceSubtypeEnum{
	"QUARTERLY":      MaintenanceRunMaintenanceSubtypeQuarterly,
	"HARDWARE":       MaintenanceRunMaintenanceSubtypeHardware,
	"CRITICAL":       MaintenanceRunMaintenanceSubtypeCritical,
	"INFRASTRUCTURE": MaintenanceRunMaintenanceSubtypeInfrastructure,
	"DATABASE":       MaintenanceRunMaintenanceSubtypeDatabase,
	"ONEOFF":         MaintenanceRunMaintenanceSubtypeOneoff,
}

// GetMaintenanceRunMaintenanceSubtypeEnumValues Enumerates the set of values for MaintenanceRunMaintenanceSubtypeEnum
func GetMaintenanceRunMaintenanceSubtypeEnumValues() []MaintenanceRunMaintenanceSubtypeEnum {
	values := make([]MaintenanceRunMaintenanceSubtypeEnum, 0)
	for _, v := range mappingMaintenanceRunMaintenanceSubtypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceRunMaintenanceSubtypeEnumStringValues Enumerates the set of values in String for MaintenanceRunMaintenanceSubtypeEnum
func GetMaintenanceRunMaintenanceSubtypeEnumStringValues() []string {
	return []string{
		"QUARTERLY",
		"HARDWARE",
		"CRITICAL",
		"INFRASTRUCTURE",
		"DATABASE",
		"ONEOFF",
	}
}

// GetMappingMaintenanceRunMaintenanceSubtypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceRunMaintenanceSubtypeEnum(val string) (MaintenanceRunMaintenanceSubtypeEnum, bool) {
	mappingMaintenanceRunMaintenanceSubtypeEnumIgnoreCase := make(map[string]MaintenanceRunMaintenanceSubtypeEnum)
	for k, v := range mappingMaintenanceRunMaintenanceSubtypeEnum {
		mappingMaintenanceRunMaintenanceSubtypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingMaintenanceRunMaintenanceSubtypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// MaintenanceRunPatchingModeEnum Enum with underlying type: string
type MaintenanceRunPatchingModeEnum string

// Set of constants representing the allowable values for MaintenanceRunPatchingModeEnum
const (
	MaintenanceRunPatchingModeRolling    MaintenanceRunPatchingModeEnum = "ROLLING"
	MaintenanceRunPatchingModeNonrolling MaintenanceRunPatchingModeEnum = "NONROLLING"
)

var mappingMaintenanceRunPatchingModeEnum = map[string]MaintenanceRunPatchingModeEnum{
	"ROLLING":    MaintenanceRunPatchingModeRolling,
	"NONROLLING": MaintenanceRunPatchingModeNonrolling,
}

// GetMaintenanceRunPatchingModeEnumValues Enumerates the set of values for MaintenanceRunPatchingModeEnum
func GetMaintenanceRunPatchingModeEnumValues() []MaintenanceRunPatchingModeEnum {
	values := make([]MaintenanceRunPatchingModeEnum, 0)
	for _, v := range mappingMaintenanceRunPatchingModeEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceRunPatchingModeEnumStringValues Enumerates the set of values in String for MaintenanceRunPatchingModeEnum
func GetMaintenanceRunPatchingModeEnumStringValues() []string {
	return []string{
		"ROLLING",
		"NONROLLING",
	}
}

// GetMappingMaintenanceRunPatchingModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceRunPatchingModeEnum(val string) (MaintenanceRunPatchingModeEnum, bool) {
	mappingMaintenanceRunPatchingModeEnumIgnoreCase := make(map[string]MaintenanceRunPatchingModeEnum)
	for k, v := range mappingMaintenanceRunPatchingModeEnum {
		mappingMaintenanceRunPatchingModeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingMaintenanceRunPatchingModeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
