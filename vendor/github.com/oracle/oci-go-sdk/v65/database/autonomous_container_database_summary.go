// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

	// **Deprecated.** The `DB_UNIQUE_NAME` value is set by Oracle Cloud Infrastructure.  Do not specify a value for this parameter. Specifying a value for this field will cause Terraform operations to fail.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// **No longer used.** For Autonomous Database on dedicated Exadata infrastructure, the container database is created within a specified `cloudAutonomousVmCluster`.
	AutonomousExadataInfrastructureId *string `mandatory:"false" json:"autonomousExadataInfrastructureId"`

	// The OCID of the Autonomous VM Cluster.
	AutonomousVmClusterId *string `mandatory:"false" json:"autonomousVmClusterId"`

	// The infrastructure type this resource belongs to.
	InfrastructureType AutonomousContainerDatabaseSummaryInfrastructureTypeEnum `mandatory:"false" json:"infrastructureType,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
	CloudAutonomousVmClusterId *string `mandatory:"false" json:"cloudAutonomousVmClusterId"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.cloud.oracle.com/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// Key History Entry.
	KeyHistoryEntry []AutonomousDatabaseKeyHistoryEntry `mandatory:"false" json:"keyHistoryEntry"`

	// Additional information about the current lifecycle state.
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

	// The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database.
	// This value represents the number of days before scheduled maintenance of the primary database.
	StandbyMaintenanceBufferInDays *int `mandatory:"false" json:"standbyMaintenanceBufferInDays"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The Data Guard role of the Autonomous Container Database or Autonomous Database, if Autonomous Data Guard is enabled.
	Role AutonomousContainerDatabaseSummaryRoleEnum `mandatory:"false" json:"role,omitempty"`

	// The availability domain of the Autonomous Container Database.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Oracle Database version of the Autonomous Container Database.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	BackupConfig *AutonomousContainerDatabaseBackupConfig `mandatory:"false" json:"backupConfig"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the key store.
	KeyStoreId *string `mandatory:"false" json:"keyStoreId"`

	// The wallet name for Oracle Key Vault.
	KeyStoreWalletName *string `mandatory:"false" json:"keyStoreWalletName"`

	// The amount of memory (in GBs) enabled per each OCPU core in Autonomous VM Cluster.
	MemoryPerOracleComputeUnitInGBs *int `mandatory:"false" json:"memoryPerOracleComputeUnitInGBs"`

	// Sum of OCPUs available on the Autonomous VM Cluster + Sum of reclaimable OCPUs available in the Autonomous Container Database.
	AvailableCpus *float32 `mandatory:"false" json:"availableCpus"`

	// The number of CPU cores allocated to the Autonomous VM cluster.
	TotalCpus *int `mandatory:"false" json:"totalCpus"`

	// CPU cores that continue to be included in the count of OCPUs available to the Autonomous Container Database even after one of its Autonomous Database is terminated or scaled down. You can release them to the available OCPUs at its parent AVMC level by restarting the Autonomous Container Database.
	ReclaimableCpus *float32 `mandatory:"false" json:"reclaimableCpus"`

	// An array of CPU values that can be used to successfully provision a single Autonomous Database.
	ProvisionableCpus []float32 `mandatory:"false" json:"provisionableCpus"`
}

func (m AutonomousContainerDatabaseSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousContainerDatabaseSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum(string(m.ServiceLevelAgreementType)); !ok && m.ServiceLevelAgreementType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceLevelAgreementType: %s. Supported values are: %s.", m.ServiceLevelAgreementType, strings.Join(GetAutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousContainerDatabaseSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousContainerDatabaseSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousContainerDatabaseSummaryPatchModelEnum(string(m.PatchModel)); !ok && m.PatchModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchModel: %s. Supported values are: %s.", m.PatchModel, strings.Join(GetAutonomousContainerDatabaseSummaryPatchModelEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAutonomousContainerDatabaseSummaryInfrastructureTypeEnum(string(m.InfrastructureType)); !ok && m.InfrastructureType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InfrastructureType: %s. Supported values are: %s.", m.InfrastructureType, strings.Join(GetAutonomousContainerDatabaseSummaryInfrastructureTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousContainerDatabaseSummaryRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetAutonomousContainerDatabaseSummaryRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum Enum with underlying type: string
type AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum
const (
	AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeStandard            AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum = "STANDARD"
	AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeMissionCritical     AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum = "MISSION_CRITICAL"
	AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeAutonomousDataguard AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum = "AUTONOMOUS_DATAGUARD"
)

var mappingAutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum = map[string]AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum{
	"STANDARD":             AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeStandard,
	"MISSION_CRITICAL":     AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeMissionCritical,
	"AUTONOMOUS_DATAGUARD": AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeAutonomousDataguard,
}

var mappingAutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnumLowerCase = map[string]AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum{
	"standard":             AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeStandard,
	"mission_critical":     AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeMissionCritical,
	"autonomous_dataguard": AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeAutonomousDataguard,
}

// GetAutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnumValues Enumerates the set of values for AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum
func GetAutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnumValues() []AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum {
	values := make([]AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum
func GetAutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnumStringValues() []string {
	return []string{
		"STANDARD",
		"MISSION_CRITICAL",
		"AUTONOMOUS_DATAGUARD",
	}
}

// GetMappingAutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum(val string) (AutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseSummaryServiceLevelAgreementTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousContainerDatabaseSummaryInfrastructureTypeEnum Enum with underlying type: string
type AutonomousContainerDatabaseSummaryInfrastructureTypeEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseSummaryInfrastructureTypeEnum
const (
	AutonomousContainerDatabaseSummaryInfrastructureTypeCloud           AutonomousContainerDatabaseSummaryInfrastructureTypeEnum = "CLOUD"
	AutonomousContainerDatabaseSummaryInfrastructureTypeCloudAtCustomer AutonomousContainerDatabaseSummaryInfrastructureTypeEnum = "CLOUD_AT_CUSTOMER"
)

var mappingAutonomousContainerDatabaseSummaryInfrastructureTypeEnum = map[string]AutonomousContainerDatabaseSummaryInfrastructureTypeEnum{
	"CLOUD":             AutonomousContainerDatabaseSummaryInfrastructureTypeCloud,
	"CLOUD_AT_CUSTOMER": AutonomousContainerDatabaseSummaryInfrastructureTypeCloudAtCustomer,
}

var mappingAutonomousContainerDatabaseSummaryInfrastructureTypeEnumLowerCase = map[string]AutonomousContainerDatabaseSummaryInfrastructureTypeEnum{
	"cloud":             AutonomousContainerDatabaseSummaryInfrastructureTypeCloud,
	"cloud_at_customer": AutonomousContainerDatabaseSummaryInfrastructureTypeCloudAtCustomer,
}

// GetAutonomousContainerDatabaseSummaryInfrastructureTypeEnumValues Enumerates the set of values for AutonomousContainerDatabaseSummaryInfrastructureTypeEnum
func GetAutonomousContainerDatabaseSummaryInfrastructureTypeEnumValues() []AutonomousContainerDatabaseSummaryInfrastructureTypeEnum {
	values := make([]AutonomousContainerDatabaseSummaryInfrastructureTypeEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseSummaryInfrastructureTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseSummaryInfrastructureTypeEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseSummaryInfrastructureTypeEnum
func GetAutonomousContainerDatabaseSummaryInfrastructureTypeEnumStringValues() []string {
	return []string{
		"CLOUD",
		"CLOUD_AT_CUSTOMER",
	}
}

// GetMappingAutonomousContainerDatabaseSummaryInfrastructureTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseSummaryInfrastructureTypeEnum(val string) (AutonomousContainerDatabaseSummaryInfrastructureTypeEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseSummaryInfrastructureTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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
	AutonomousContainerDatabaseSummaryLifecycleStateRoleChangeInProgress  AutonomousContainerDatabaseSummaryLifecycleStateEnum = "ROLE_CHANGE_IN_PROGRESS"
	AutonomousContainerDatabaseSummaryLifecycleStateUnavailable           AutonomousContainerDatabaseSummaryLifecycleStateEnum = "UNAVAILABLE"
)

var mappingAutonomousContainerDatabaseSummaryLifecycleStateEnum = map[string]AutonomousContainerDatabaseSummaryLifecycleStateEnum{
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
	"ROLE_CHANGE_IN_PROGRESS": AutonomousContainerDatabaseSummaryLifecycleStateRoleChangeInProgress,
	"UNAVAILABLE":             AutonomousContainerDatabaseSummaryLifecycleStateUnavailable,
}

var mappingAutonomousContainerDatabaseSummaryLifecycleStateEnumLowerCase = map[string]AutonomousContainerDatabaseSummaryLifecycleStateEnum{
	"provisioning":            AutonomousContainerDatabaseSummaryLifecycleStateProvisioning,
	"available":               AutonomousContainerDatabaseSummaryLifecycleStateAvailable,
	"updating":                AutonomousContainerDatabaseSummaryLifecycleStateUpdating,
	"terminating":             AutonomousContainerDatabaseSummaryLifecycleStateTerminating,
	"terminated":              AutonomousContainerDatabaseSummaryLifecycleStateTerminated,
	"failed":                  AutonomousContainerDatabaseSummaryLifecycleStateFailed,
	"backup_in_progress":      AutonomousContainerDatabaseSummaryLifecycleStateBackupInProgress,
	"restoring":               AutonomousContainerDatabaseSummaryLifecycleStateRestoring,
	"restore_failed":          AutonomousContainerDatabaseSummaryLifecycleStateRestoreFailed,
	"restarting":              AutonomousContainerDatabaseSummaryLifecycleStateRestarting,
	"maintenance_in_progress": AutonomousContainerDatabaseSummaryLifecycleStateMaintenanceInProgress,
	"role_change_in_progress": AutonomousContainerDatabaseSummaryLifecycleStateRoleChangeInProgress,
	"unavailable":             AutonomousContainerDatabaseSummaryLifecycleStateUnavailable,
}

// GetAutonomousContainerDatabaseSummaryLifecycleStateEnumValues Enumerates the set of values for AutonomousContainerDatabaseSummaryLifecycleStateEnum
func GetAutonomousContainerDatabaseSummaryLifecycleStateEnumValues() []AutonomousContainerDatabaseSummaryLifecycleStateEnum {
	values := make([]AutonomousContainerDatabaseSummaryLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseSummaryLifecycleStateEnum
func GetAutonomousContainerDatabaseSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"UPDATING",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
		"BACKUP_IN_PROGRESS",
		"RESTORING",
		"RESTORE_FAILED",
		"RESTARTING",
		"MAINTENANCE_IN_PROGRESS",
		"ROLE_CHANGE_IN_PROGRESS",
		"UNAVAILABLE",
	}
}

// GetMappingAutonomousContainerDatabaseSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseSummaryLifecycleStateEnum(val string) (AutonomousContainerDatabaseSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousContainerDatabaseSummaryPatchModelEnum Enum with underlying type: string
type AutonomousContainerDatabaseSummaryPatchModelEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseSummaryPatchModelEnum
const (
	AutonomousContainerDatabaseSummaryPatchModelUpdates         AutonomousContainerDatabaseSummaryPatchModelEnum = "RELEASE_UPDATES"
	AutonomousContainerDatabaseSummaryPatchModelUpdateRevisions AutonomousContainerDatabaseSummaryPatchModelEnum = "RELEASE_UPDATE_REVISIONS"
)

var mappingAutonomousContainerDatabaseSummaryPatchModelEnum = map[string]AutonomousContainerDatabaseSummaryPatchModelEnum{
	"RELEASE_UPDATES":          AutonomousContainerDatabaseSummaryPatchModelUpdates,
	"RELEASE_UPDATE_REVISIONS": AutonomousContainerDatabaseSummaryPatchModelUpdateRevisions,
}

var mappingAutonomousContainerDatabaseSummaryPatchModelEnumLowerCase = map[string]AutonomousContainerDatabaseSummaryPatchModelEnum{
	"release_updates":          AutonomousContainerDatabaseSummaryPatchModelUpdates,
	"release_update_revisions": AutonomousContainerDatabaseSummaryPatchModelUpdateRevisions,
}

// GetAutonomousContainerDatabaseSummaryPatchModelEnumValues Enumerates the set of values for AutonomousContainerDatabaseSummaryPatchModelEnum
func GetAutonomousContainerDatabaseSummaryPatchModelEnumValues() []AutonomousContainerDatabaseSummaryPatchModelEnum {
	values := make([]AutonomousContainerDatabaseSummaryPatchModelEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseSummaryPatchModelEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseSummaryPatchModelEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseSummaryPatchModelEnum
func GetAutonomousContainerDatabaseSummaryPatchModelEnumStringValues() []string {
	return []string{
		"RELEASE_UPDATES",
		"RELEASE_UPDATE_REVISIONS",
	}
}

// GetMappingAutonomousContainerDatabaseSummaryPatchModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseSummaryPatchModelEnum(val string) (AutonomousContainerDatabaseSummaryPatchModelEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseSummaryPatchModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousContainerDatabaseSummaryRoleEnum Enum with underlying type: string
type AutonomousContainerDatabaseSummaryRoleEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseSummaryRoleEnum
const (
	AutonomousContainerDatabaseSummaryRolePrimary         AutonomousContainerDatabaseSummaryRoleEnum = "PRIMARY"
	AutonomousContainerDatabaseSummaryRoleStandby         AutonomousContainerDatabaseSummaryRoleEnum = "STANDBY"
	AutonomousContainerDatabaseSummaryRoleDisabledStandby AutonomousContainerDatabaseSummaryRoleEnum = "DISABLED_STANDBY"
)

var mappingAutonomousContainerDatabaseSummaryRoleEnum = map[string]AutonomousContainerDatabaseSummaryRoleEnum{
	"PRIMARY":          AutonomousContainerDatabaseSummaryRolePrimary,
	"STANDBY":          AutonomousContainerDatabaseSummaryRoleStandby,
	"DISABLED_STANDBY": AutonomousContainerDatabaseSummaryRoleDisabledStandby,
}

var mappingAutonomousContainerDatabaseSummaryRoleEnumLowerCase = map[string]AutonomousContainerDatabaseSummaryRoleEnum{
	"primary":          AutonomousContainerDatabaseSummaryRolePrimary,
	"standby":          AutonomousContainerDatabaseSummaryRoleStandby,
	"disabled_standby": AutonomousContainerDatabaseSummaryRoleDisabledStandby,
}

// GetAutonomousContainerDatabaseSummaryRoleEnumValues Enumerates the set of values for AutonomousContainerDatabaseSummaryRoleEnum
func GetAutonomousContainerDatabaseSummaryRoleEnumValues() []AutonomousContainerDatabaseSummaryRoleEnum {
	values := make([]AutonomousContainerDatabaseSummaryRoleEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseSummaryRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseSummaryRoleEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseSummaryRoleEnum
func GetAutonomousContainerDatabaseSummaryRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"STANDBY",
		"DISABLED_STANDBY",
	}
}

// GetMappingAutonomousContainerDatabaseSummaryRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseSummaryRoleEnum(val string) (AutonomousContainerDatabaseSummaryRoleEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseSummaryRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
