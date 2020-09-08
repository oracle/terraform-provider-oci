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

// AutonomousContainerDatabase The representation of AutonomousContainerDatabase
type AutonomousContainerDatabase struct {

	// The OCID of the Autonomous Container Database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-provided name for the Autonomous Container Database.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The service level agreement type of the container database. The default is STANDARD.
	ServiceLevelAgreementType AutonomousContainerDatabaseServiceLevelAgreementTypeEnum `mandatory:"true" json:"serviceLevelAgreementType"`

	// The current state of the Autonomous Container Database.
	LifecycleState AutonomousContainerDatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Database patch model preference.
	PatchModel AutonomousContainerDatabasePatchModelEnum `mandatory:"true" json:"patchModel"`

	// The `DB_UNIQUE_NAME` of the Oracle Database being backed up.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// The OCID of the Autonomous Exadata Infrastructure.
	AutonomousExadataInfrastructureId *string `mandatory:"false" json:"autonomousExadataInfrastructureId"`

	// The OCID of the Autonomous VM Cluster.
	AutonomousVmClusterId *string `mandatory:"false" json:"autonomousVmClusterId"`

	// The infrastructure type this resource belongs to.
	InfrastructureType AutonomousContainerDatabaseInfrastructureTypeEnum `mandatory:"false" json:"infrastructureType,omitempty"`

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

func (m AutonomousContainerDatabase) String() string {
	return common.PointerString(m)
}

// AutonomousContainerDatabaseServiceLevelAgreementTypeEnum Enum with underlying type: string
type AutonomousContainerDatabaseServiceLevelAgreementTypeEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseServiceLevelAgreementTypeEnum
const (
	AutonomousContainerDatabaseServiceLevelAgreementTypeStandard        AutonomousContainerDatabaseServiceLevelAgreementTypeEnum = "STANDARD"
	AutonomousContainerDatabaseServiceLevelAgreementTypeMissionCritical AutonomousContainerDatabaseServiceLevelAgreementTypeEnum = "MISSION_CRITICAL"
)

var mappingAutonomousContainerDatabaseServiceLevelAgreementType = map[string]AutonomousContainerDatabaseServiceLevelAgreementTypeEnum{
	"STANDARD":         AutonomousContainerDatabaseServiceLevelAgreementTypeStandard,
	"MISSION_CRITICAL": AutonomousContainerDatabaseServiceLevelAgreementTypeMissionCritical,
}

// GetAutonomousContainerDatabaseServiceLevelAgreementTypeEnumValues Enumerates the set of values for AutonomousContainerDatabaseServiceLevelAgreementTypeEnum
func GetAutonomousContainerDatabaseServiceLevelAgreementTypeEnumValues() []AutonomousContainerDatabaseServiceLevelAgreementTypeEnum {
	values := make([]AutonomousContainerDatabaseServiceLevelAgreementTypeEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseServiceLevelAgreementType {
		values = append(values, v)
	}
	return values
}

// AutonomousContainerDatabaseInfrastructureTypeEnum Enum with underlying type: string
type AutonomousContainerDatabaseInfrastructureTypeEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseInfrastructureTypeEnum
const (
	AutonomousContainerDatabaseInfrastructureTypeCloud           AutonomousContainerDatabaseInfrastructureTypeEnum = "CLOUD"
	AutonomousContainerDatabaseInfrastructureTypeCloudAtCustomer AutonomousContainerDatabaseInfrastructureTypeEnum = "CLOUD_AT_CUSTOMER"
)

var mappingAutonomousContainerDatabaseInfrastructureType = map[string]AutonomousContainerDatabaseInfrastructureTypeEnum{
	"CLOUD":             AutonomousContainerDatabaseInfrastructureTypeCloud,
	"CLOUD_AT_CUSTOMER": AutonomousContainerDatabaseInfrastructureTypeCloudAtCustomer,
}

// GetAutonomousContainerDatabaseInfrastructureTypeEnumValues Enumerates the set of values for AutonomousContainerDatabaseInfrastructureTypeEnum
func GetAutonomousContainerDatabaseInfrastructureTypeEnumValues() []AutonomousContainerDatabaseInfrastructureTypeEnum {
	values := make([]AutonomousContainerDatabaseInfrastructureTypeEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseInfrastructureType {
		values = append(values, v)
	}
	return values
}

// AutonomousContainerDatabaseLifecycleStateEnum Enum with underlying type: string
type AutonomousContainerDatabaseLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseLifecycleStateEnum
const (
	AutonomousContainerDatabaseLifecycleStateProvisioning          AutonomousContainerDatabaseLifecycleStateEnum = "PROVISIONING"
	AutonomousContainerDatabaseLifecycleStateAvailable             AutonomousContainerDatabaseLifecycleStateEnum = "AVAILABLE"
	AutonomousContainerDatabaseLifecycleStateUpdating              AutonomousContainerDatabaseLifecycleStateEnum = "UPDATING"
	AutonomousContainerDatabaseLifecycleStateTerminating           AutonomousContainerDatabaseLifecycleStateEnum = "TERMINATING"
	AutonomousContainerDatabaseLifecycleStateTerminated            AutonomousContainerDatabaseLifecycleStateEnum = "TERMINATED"
	AutonomousContainerDatabaseLifecycleStateFailed                AutonomousContainerDatabaseLifecycleStateEnum = "FAILED"
	AutonomousContainerDatabaseLifecycleStateBackupInProgress      AutonomousContainerDatabaseLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	AutonomousContainerDatabaseLifecycleStateRestoring             AutonomousContainerDatabaseLifecycleStateEnum = "RESTORING"
	AutonomousContainerDatabaseLifecycleStateRestoreFailed         AutonomousContainerDatabaseLifecycleStateEnum = "RESTORE_FAILED"
	AutonomousContainerDatabaseLifecycleStateRestarting            AutonomousContainerDatabaseLifecycleStateEnum = "RESTARTING"
	AutonomousContainerDatabaseLifecycleStateMaintenanceInProgress AutonomousContainerDatabaseLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
)

var mappingAutonomousContainerDatabaseLifecycleState = map[string]AutonomousContainerDatabaseLifecycleStateEnum{
	"PROVISIONING":            AutonomousContainerDatabaseLifecycleStateProvisioning,
	"AVAILABLE":               AutonomousContainerDatabaseLifecycleStateAvailable,
	"UPDATING":                AutonomousContainerDatabaseLifecycleStateUpdating,
	"TERMINATING":             AutonomousContainerDatabaseLifecycleStateTerminating,
	"TERMINATED":              AutonomousContainerDatabaseLifecycleStateTerminated,
	"FAILED":                  AutonomousContainerDatabaseLifecycleStateFailed,
	"BACKUP_IN_PROGRESS":      AutonomousContainerDatabaseLifecycleStateBackupInProgress,
	"RESTORING":               AutonomousContainerDatabaseLifecycleStateRestoring,
	"RESTORE_FAILED":          AutonomousContainerDatabaseLifecycleStateRestoreFailed,
	"RESTARTING":              AutonomousContainerDatabaseLifecycleStateRestarting,
	"MAINTENANCE_IN_PROGRESS": AutonomousContainerDatabaseLifecycleStateMaintenanceInProgress,
}

// GetAutonomousContainerDatabaseLifecycleStateEnumValues Enumerates the set of values for AutonomousContainerDatabaseLifecycleStateEnum
func GetAutonomousContainerDatabaseLifecycleStateEnumValues() []AutonomousContainerDatabaseLifecycleStateEnum {
	values := make([]AutonomousContainerDatabaseLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseLifecycleState {
		values = append(values, v)
	}
	return values
}

// AutonomousContainerDatabasePatchModelEnum Enum with underlying type: string
type AutonomousContainerDatabasePatchModelEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabasePatchModelEnum
const (
	AutonomousContainerDatabasePatchModelUpdates         AutonomousContainerDatabasePatchModelEnum = "RELEASE_UPDATES"
	AutonomousContainerDatabasePatchModelUpdateRevisions AutonomousContainerDatabasePatchModelEnum = "RELEASE_UPDATE_REVISIONS"
)

var mappingAutonomousContainerDatabasePatchModel = map[string]AutonomousContainerDatabasePatchModelEnum{
	"RELEASE_UPDATES":          AutonomousContainerDatabasePatchModelUpdates,
	"RELEASE_UPDATE_REVISIONS": AutonomousContainerDatabasePatchModelUpdateRevisions,
}

// GetAutonomousContainerDatabasePatchModelEnumValues Enumerates the set of values for AutonomousContainerDatabasePatchModelEnum
func GetAutonomousContainerDatabasePatchModelEnumValues() []AutonomousContainerDatabasePatchModelEnum {
	values := make([]AutonomousContainerDatabasePatchModelEnum, 0)
	for _, v := range mappingAutonomousContainerDatabasePatchModel {
		values = append(values, v)
	}
	return values
}
