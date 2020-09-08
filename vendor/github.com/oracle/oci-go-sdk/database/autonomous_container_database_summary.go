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

// AutonomousContainerDatabaseSummary An Autonomous Container Database is a container database service that enables the customer to host one or more databases within the container database. A basic container database runs on a single Autonomous Exadata Infrastructure from an availability domain without the Extreme Availability features enabled.
type AutonomousContainerDatabaseSummary struct {

	// The OCID of the Autonomous Container Database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-provided name for the Autonomous Container Database.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The service level agreement type of the container database. The default is STANDARD.
	ServiceLevelAgreementType AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum `mandatory:"true" json:"serviceLevelAgreementType"`

	// The current state of the Autonomous Container Database.
	LifecycleState AutonomousContainerDatabaseSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Database patch model preference.
	PatchModel AutonomousContainerDatabaseSummaryPatchModelEnum `mandatory:"true" json:"patchModel"`

	// The `DB_UNIQUE_NAME` of the Oracle Database being backed up.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// The OCID of the Autonomous Exadata Infrastructure.
	AutonomousExadataInfrastructureId *string `mandatory:"false" json:"autonomousExadataInfrastructureId"`

	// The OCID of the Autonomous VM Cluster.
	AutonomousVmClusterId *string `mandatory:"false" json:"autonomousVmClusterId"`

	// The infrastructure type this resource belongs to.
	InfrastructureType AutonomousContainerDatabaseSummaryInfrastructureTypeEnum `mandatory:"false" json:"infrastructureType,omitempty"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the Autonomous Container Database was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the last patch applied on the system.
	PatchId *string `mandatory:"false" json:"patchId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the last maintenance run.
	LastMaintenanceRunId *string `mandatory:"false" json:"lastMaintenanceRunId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the next maintenance run.
	NextMaintenanceRunId *string `mandatory:"false" json:"nextMaintenanceRunId"`

	MaintenanceWindow *MaintenanceWindow `mandatory:"false" json:"maintenanceWindow"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The availability domain of the Autonomous Container Database.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Oracle Database version of the Autonomous Container Database
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	BackupConfig *AutonomousContainerDatabaseBackupConfig `mandatory:"false" json:"backupConfig"`
}

func (m AutonomousContainerDatabaseSummary) String() string {
	return common.PointerString(m)
}

// AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum Enum with underlying type: string
type AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum
const (
	AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeStandard        AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum = "STANDARD"
	AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeMissionCritical AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum = "MISSION_CRITICAL"
)

var mappingAutonomousContainerDatabaseSummaryServiceLevelAgreementType = map[string]AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum{
	"STANDARD":         AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeStandard,
	"MISSION_CRITICAL": AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeMissionCritical,
}

// GetAutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnumValues Enumerates the set of values for AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum
func GetAutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnumValues() []AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum {
	values := make([]AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseSummaryServiceLevelAgreementType {
		values = append(values, v)
	}
	return values
}

// AutonomousContainerDatabaseSummaryInfrastructureTypeEnum Enum with underlying type: string
type AutonomousContainerDatabaseSummaryInfrastructureTypeEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseSummaryInfrastructureTypeEnum
const (
	AutonomousContainerDatabaseSummaryInfrastructureTypeCloud           AutonomousContainerDatabaseSummaryInfrastructureTypeEnum = "CLOUD"
	AutonomousContainerDatabaseSummaryInfrastructureTypeCloudAtCustomer AutonomousContainerDatabaseSummaryInfrastructureTypeEnum = "CLOUD_AT_CUSTOMER"
)

var mappingAutonomousContainerDatabaseSummaryInfrastructureType = map[string]AutonomousContainerDatabaseSummaryInfrastructureTypeEnum{
	"CLOUD":             AutonomousContainerDatabaseSummaryInfrastructureTypeCloud,
	"CLOUD_AT_CUSTOMER": AutonomousContainerDatabaseSummaryInfrastructureTypeCloudAtCustomer,
}

// GetAutonomousContainerDatabaseSummaryInfrastructureTypeEnumValues Enumerates the set of values for AutonomousContainerDatabaseSummaryInfrastructureTypeEnum
func GetAutonomousContainerDatabaseSummaryInfrastructureTypeEnumValues() []AutonomousContainerDatabaseSummaryInfrastructureTypeEnum {
	values := make([]AutonomousContainerDatabaseSummaryInfrastructureTypeEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseSummaryInfrastructureType {
		values = append(values, v)
	}
	return values
}

// AutonomousContainerDatabaseSummaryLifecycleStateEnum Enum with underlying type: string
type AutonomousContainerDatabaseSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseSummaryLifecycleStateEnum
const (
	AutonomousContainerDatabaseSummaryLifecycleStateProvisioning          AutonomousContainerDatabaseSummaryLifecycleStateEnum = "PROVISIONING"
	AutonomousContainerDatabaseSummaryLifecycleStateAvailable             AutonomousContainerDatabaseSummaryLifecycleStateEnum = "AVAILABLE"
	AutonomousContainerDatabaseSummaryLifecycleStateUpdating              AutonomousContainerDatabaseSummaryLifecycleStateEnum = "UPDATING"
	AutonomousContainerDatabaseSummaryLifecycleStateTerminating           AutonomousContainerDatabaseSummaryLifecycleStateEnum = "TERMINATING"
	AutonomousContainerDatabaseSummaryLifecycleStateTerminated            AutonomousContainerDatabaseSummaryLifecycleStateEnum = "TERMINATED"
	AutonomousContainerDatabaseSummaryLifecycleStateFailed                AutonomousContainerDatabaseSummaryLifecycleStateEnum = "FAILED"
	AutonomousContainerDatabaseSummaryLifecycleStateBackupInProgress      AutonomousContainerDatabaseSummaryLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	AutonomousContainerDatabaseSummaryLifecycleStateRestoring             AutonomousContainerDatabaseSummaryLifecycleStateEnum = "RESTORING"
	AutonomousContainerDatabaseSummaryLifecycleStateRestoreFailed         AutonomousContainerDatabaseSummaryLifecycleStateEnum = "RESTORE_FAILED"
	AutonomousContainerDatabaseSummaryLifecycleStateRestarting            AutonomousContainerDatabaseSummaryLifecycleStateEnum = "RESTARTING"
	AutonomousContainerDatabaseSummaryLifecycleStateMaintenanceInProgress AutonomousContainerDatabaseSummaryLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
)

var mappingAutonomousContainerDatabaseSummaryLifecycleState = map[string]AutonomousContainerDatabaseSummaryLifecycleStateEnum{
	"PROVISIONING":            AutonomousContainerDatabaseSummaryLifecycleStateProvisioning,
	"AVAILABLE":               AutonomousContainerDatabaseSummaryLifecycleStateAvailable,
	"UPDATING":                AutonomousContainerDatabaseSummaryLifecycleStateUpdating,
	"TERMINATING":             AutonomousContainerDatabaseSummaryLifecycleStateTerminating,
	"TERMINATED":              AutonomousContainerDatabaseSummaryLifecycleStateTerminated,
	"FAILED":                  AutonomousContainerDatabaseSummaryLifecycleStateFailed,
	"BACKUP_IN_PROGRESS":      AutonomousContainerDatabaseSummaryLifecycleStateBackupInProgress,
	"RESTORING":               AutonomousContainerDatabaseSummaryLifecycleStateRestoring,
	"RESTORE_FAILED":          AutonomousContainerDatabaseSummaryLifecycleStateRestoreFailed,
	"RESTARTING":              AutonomousContainerDatabaseSummaryLifecycleStateRestarting,
	"MAINTENANCE_IN_PROGRESS": AutonomousContainerDatabaseSummaryLifecycleStateMaintenanceInProgress,
}

// GetAutonomousContainerDatabaseSummaryLifecycleStateEnumValues Enumerates the set of values for AutonomousContainerDatabaseSummaryLifecycleStateEnum
func GetAutonomousContainerDatabaseSummaryLifecycleStateEnumValues() []AutonomousContainerDatabaseSummaryLifecycleStateEnum {
	values := make([]AutonomousContainerDatabaseSummaryLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}

// AutonomousContainerDatabaseSummaryPatchModelEnum Enum with underlying type: string
type AutonomousContainerDatabaseSummaryPatchModelEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseSummaryPatchModelEnum
const (
	AutonomousContainerDatabaseSummaryPatchModelUpdates         AutonomousContainerDatabaseSummaryPatchModelEnum = "RELEASE_UPDATES"
	AutonomousContainerDatabaseSummaryPatchModelUpdateRevisions AutonomousContainerDatabaseSummaryPatchModelEnum = "RELEASE_UPDATE_REVISIONS"
)

var mappingAutonomousContainerDatabaseSummaryPatchModel = map[string]AutonomousContainerDatabaseSummaryPatchModelEnum{
	"RELEASE_UPDATES":          AutonomousContainerDatabaseSummaryPatchModelUpdates,
	"RELEASE_UPDATE_REVISIONS": AutonomousContainerDatabaseSummaryPatchModelUpdateRevisions,
}

// GetAutonomousContainerDatabaseSummaryPatchModelEnumValues Enumerates the set of values for AutonomousContainerDatabaseSummaryPatchModelEnum
func GetAutonomousContainerDatabaseSummaryPatchModelEnumValues() []AutonomousContainerDatabaseSummaryPatchModelEnum {
	values := make([]AutonomousContainerDatabaseSummaryPatchModelEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseSummaryPatchModel {
		values = append(values, v)
	}
	return values
}
