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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

	// **Deprecated.** The `DB_UNIQUE_NAME` value is set by Oracle Cloud Infrastructure.  Do not specify a value for this parameter. Specifying a value for this field will cause Terraform operations to fail.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// **No longer used.** For Autonomous Database on dedicated Exadata infrastructure, the container database is created within a specified `cloudAutonomousVmCluster`.
	AutonomousExadataInfrastructureId *string `mandatory:"false" json:"autonomousExadataInfrastructureId"`

	// The OCID of the Autonomous VM Cluster.
	AutonomousVmClusterId *string `mandatory:"false" json:"autonomousVmClusterId"`

	// The infrastructure type this resource belongs to.
	InfrastructureType AutonomousContainerDatabaseInfrastructureTypeEnum `mandatory:"false" json:"infrastructureType,omitempty"`

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

	// The role of the Autonomous Data Guard-enabled Autonomous Container Database.
	Role AutonomousContainerDatabaseRoleEnum `mandatory:"false" json:"role,omitempty"`

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
}

func (m AutonomousContainerDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousContainerDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousContainerDatabaseServiceLevelAgreementTypeEnum(string(m.ServiceLevelAgreementType)); !ok && m.ServiceLevelAgreementType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceLevelAgreementType: %s. Supported values are: %s.", m.ServiceLevelAgreementType, strings.Join(GetAutonomousContainerDatabaseServiceLevelAgreementTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousContainerDatabaseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousContainerDatabaseLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousContainerDatabasePatchModelEnum(string(m.PatchModel)); !ok && m.PatchModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchModel: %s. Supported values are: %s.", m.PatchModel, strings.Join(GetAutonomousContainerDatabasePatchModelEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAutonomousContainerDatabaseInfrastructureTypeEnum(string(m.InfrastructureType)); !ok && m.InfrastructureType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InfrastructureType: %s. Supported values are: %s.", m.InfrastructureType, strings.Join(GetAutonomousContainerDatabaseInfrastructureTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousContainerDatabaseRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetAutonomousContainerDatabaseRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousContainerDatabaseServiceLevelAgreementTypeEnum Enum with underlying type: string
type AutonomousContainerDatabaseServiceLevelAgreementTypeEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseServiceLevelAgreementTypeEnum
const (
	AutonomousContainerDatabaseServiceLevelAgreementTypeStandard            AutonomousContainerDatabaseServiceLevelAgreementTypeEnum = "STANDARD"
	AutonomousContainerDatabaseServiceLevelAgreementTypeMissionCritical     AutonomousContainerDatabaseServiceLevelAgreementTypeEnum = "MISSION_CRITICAL"
	AutonomousContainerDatabaseServiceLevelAgreementTypeAutonomousDataguard AutonomousContainerDatabaseServiceLevelAgreementTypeEnum = "AUTONOMOUS_DATAGUARD"
)

var mappingAutonomousContainerDatabaseServiceLevelAgreementTypeEnum = map[string]AutonomousContainerDatabaseServiceLevelAgreementTypeEnum{
	"STANDARD":             AutonomousContainerDatabaseServiceLevelAgreementTypeStandard,
	"MISSION_CRITICAL":     AutonomousContainerDatabaseServiceLevelAgreementTypeMissionCritical,
	"AUTONOMOUS_DATAGUARD": AutonomousContainerDatabaseServiceLevelAgreementTypeAutonomousDataguard,
}

var mappingAutonomousContainerDatabaseServiceLevelAgreementTypeEnumLowerCase = map[string]AutonomousContainerDatabaseServiceLevelAgreementTypeEnum{
	"standard":             AutonomousContainerDatabaseServiceLevelAgreementTypeStandard,
	"mission_critical":     AutonomousContainerDatabaseServiceLevelAgreementTypeMissionCritical,
	"autonomous_dataguard": AutonomousContainerDatabaseServiceLevelAgreementTypeAutonomousDataguard,
}

// GetAutonomousContainerDatabaseServiceLevelAgreementTypeEnumValues Enumerates the set of values for AutonomousContainerDatabaseServiceLevelAgreementTypeEnum
func GetAutonomousContainerDatabaseServiceLevelAgreementTypeEnumValues() []AutonomousContainerDatabaseServiceLevelAgreementTypeEnum {
	values := make([]AutonomousContainerDatabaseServiceLevelAgreementTypeEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseServiceLevelAgreementTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseServiceLevelAgreementTypeEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseServiceLevelAgreementTypeEnum
func GetAutonomousContainerDatabaseServiceLevelAgreementTypeEnumStringValues() []string {
	return []string{
		"STANDARD",
		"MISSION_CRITICAL",
		"AUTONOMOUS_DATAGUARD",
	}
}

// GetMappingAutonomousContainerDatabaseServiceLevelAgreementTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseServiceLevelAgreementTypeEnum(val string) (AutonomousContainerDatabaseServiceLevelAgreementTypeEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseServiceLevelAgreementTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousContainerDatabaseInfrastructureTypeEnum Enum with underlying type: string
type AutonomousContainerDatabaseInfrastructureTypeEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseInfrastructureTypeEnum
const (
	AutonomousContainerDatabaseInfrastructureTypeCloud           AutonomousContainerDatabaseInfrastructureTypeEnum = "CLOUD"
	AutonomousContainerDatabaseInfrastructureTypeCloudAtCustomer AutonomousContainerDatabaseInfrastructureTypeEnum = "CLOUD_AT_CUSTOMER"
)

var mappingAutonomousContainerDatabaseInfrastructureTypeEnum = map[string]AutonomousContainerDatabaseInfrastructureTypeEnum{
	"CLOUD":             AutonomousContainerDatabaseInfrastructureTypeCloud,
	"CLOUD_AT_CUSTOMER": AutonomousContainerDatabaseInfrastructureTypeCloudAtCustomer,
}

var mappingAutonomousContainerDatabaseInfrastructureTypeEnumLowerCase = map[string]AutonomousContainerDatabaseInfrastructureTypeEnum{
	"cloud":             AutonomousContainerDatabaseInfrastructureTypeCloud,
	"cloud_at_customer": AutonomousContainerDatabaseInfrastructureTypeCloudAtCustomer,
}

// GetAutonomousContainerDatabaseInfrastructureTypeEnumValues Enumerates the set of values for AutonomousContainerDatabaseInfrastructureTypeEnum
func GetAutonomousContainerDatabaseInfrastructureTypeEnumValues() []AutonomousContainerDatabaseInfrastructureTypeEnum {
	values := make([]AutonomousContainerDatabaseInfrastructureTypeEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseInfrastructureTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseInfrastructureTypeEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseInfrastructureTypeEnum
func GetAutonomousContainerDatabaseInfrastructureTypeEnumStringValues() []string {
	return []string{
		"CLOUD",
		"CLOUD_AT_CUSTOMER",
	}
}

// GetMappingAutonomousContainerDatabaseInfrastructureTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseInfrastructureTypeEnum(val string) (AutonomousContainerDatabaseInfrastructureTypeEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseInfrastructureTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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
	AutonomousContainerDatabaseLifecycleStateRoleChangeInProgress  AutonomousContainerDatabaseLifecycleStateEnum = "ROLE_CHANGE_IN_PROGRESS"
	AutonomousContainerDatabaseLifecycleStateUnavailable           AutonomousContainerDatabaseLifecycleStateEnum = "UNAVAILABLE"
)

var mappingAutonomousContainerDatabaseLifecycleStateEnum = map[string]AutonomousContainerDatabaseLifecycleStateEnum{
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
	"ROLE_CHANGE_IN_PROGRESS": AutonomousContainerDatabaseLifecycleStateRoleChangeInProgress,
	"UNAVAILABLE":             AutonomousContainerDatabaseLifecycleStateUnavailable,
}

var mappingAutonomousContainerDatabaseLifecycleStateEnumLowerCase = map[string]AutonomousContainerDatabaseLifecycleStateEnum{
	"provisioning":            AutonomousContainerDatabaseLifecycleStateProvisioning,
	"available":               AutonomousContainerDatabaseLifecycleStateAvailable,
	"updating":                AutonomousContainerDatabaseLifecycleStateUpdating,
	"terminating":             AutonomousContainerDatabaseLifecycleStateTerminating,
	"terminated":              AutonomousContainerDatabaseLifecycleStateTerminated,
	"failed":                  AutonomousContainerDatabaseLifecycleStateFailed,
	"backup_in_progress":      AutonomousContainerDatabaseLifecycleStateBackupInProgress,
	"restoring":               AutonomousContainerDatabaseLifecycleStateRestoring,
	"restore_failed":          AutonomousContainerDatabaseLifecycleStateRestoreFailed,
	"restarting":              AutonomousContainerDatabaseLifecycleStateRestarting,
	"maintenance_in_progress": AutonomousContainerDatabaseLifecycleStateMaintenanceInProgress,
	"role_change_in_progress": AutonomousContainerDatabaseLifecycleStateRoleChangeInProgress,
	"unavailable":             AutonomousContainerDatabaseLifecycleStateUnavailable,
}

// GetAutonomousContainerDatabaseLifecycleStateEnumValues Enumerates the set of values for AutonomousContainerDatabaseLifecycleStateEnum
func GetAutonomousContainerDatabaseLifecycleStateEnumValues() []AutonomousContainerDatabaseLifecycleStateEnum {
	values := make([]AutonomousContainerDatabaseLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseLifecycleStateEnum
func GetAutonomousContainerDatabaseLifecycleStateEnumStringValues() []string {
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

// GetMappingAutonomousContainerDatabaseLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseLifecycleStateEnum(val string) (AutonomousContainerDatabaseLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousContainerDatabasePatchModelEnum Enum with underlying type: string
type AutonomousContainerDatabasePatchModelEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabasePatchModelEnum
const (
	AutonomousContainerDatabasePatchModelUpdates         AutonomousContainerDatabasePatchModelEnum = "RELEASE_UPDATES"
	AutonomousContainerDatabasePatchModelUpdateRevisions AutonomousContainerDatabasePatchModelEnum = "RELEASE_UPDATE_REVISIONS"
)

var mappingAutonomousContainerDatabasePatchModelEnum = map[string]AutonomousContainerDatabasePatchModelEnum{
	"RELEASE_UPDATES":          AutonomousContainerDatabasePatchModelUpdates,
	"RELEASE_UPDATE_REVISIONS": AutonomousContainerDatabasePatchModelUpdateRevisions,
}

var mappingAutonomousContainerDatabasePatchModelEnumLowerCase = map[string]AutonomousContainerDatabasePatchModelEnum{
	"release_updates":          AutonomousContainerDatabasePatchModelUpdates,
	"release_update_revisions": AutonomousContainerDatabasePatchModelUpdateRevisions,
}

// GetAutonomousContainerDatabasePatchModelEnumValues Enumerates the set of values for AutonomousContainerDatabasePatchModelEnum
func GetAutonomousContainerDatabasePatchModelEnumValues() []AutonomousContainerDatabasePatchModelEnum {
	values := make([]AutonomousContainerDatabasePatchModelEnum, 0)
	for _, v := range mappingAutonomousContainerDatabasePatchModelEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabasePatchModelEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabasePatchModelEnum
func GetAutonomousContainerDatabasePatchModelEnumStringValues() []string {
	return []string{
		"RELEASE_UPDATES",
		"RELEASE_UPDATE_REVISIONS",
	}
}

// GetMappingAutonomousContainerDatabasePatchModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabasePatchModelEnum(val string) (AutonomousContainerDatabasePatchModelEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabasePatchModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousContainerDatabaseRoleEnum Enum with underlying type: string
type AutonomousContainerDatabaseRoleEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseRoleEnum
const (
	AutonomousContainerDatabaseRolePrimary         AutonomousContainerDatabaseRoleEnum = "PRIMARY"
	AutonomousContainerDatabaseRoleStandby         AutonomousContainerDatabaseRoleEnum = "STANDBY"
	AutonomousContainerDatabaseRoleDisabledStandby AutonomousContainerDatabaseRoleEnum = "DISABLED_STANDBY"
)

var mappingAutonomousContainerDatabaseRoleEnum = map[string]AutonomousContainerDatabaseRoleEnum{
	"PRIMARY":          AutonomousContainerDatabaseRolePrimary,
	"STANDBY":          AutonomousContainerDatabaseRoleStandby,
	"DISABLED_STANDBY": AutonomousContainerDatabaseRoleDisabledStandby,
}

var mappingAutonomousContainerDatabaseRoleEnumLowerCase = map[string]AutonomousContainerDatabaseRoleEnum{
	"primary":          AutonomousContainerDatabaseRolePrimary,
	"standby":          AutonomousContainerDatabaseRoleStandby,
	"disabled_standby": AutonomousContainerDatabaseRoleDisabledStandby,
}

// GetAutonomousContainerDatabaseRoleEnumValues Enumerates the set of values for AutonomousContainerDatabaseRoleEnum
func GetAutonomousContainerDatabaseRoleEnumValues() []AutonomousContainerDatabaseRoleEnum {
	values := make([]AutonomousContainerDatabaseRoleEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseRoleEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseRoleEnum
func GetAutonomousContainerDatabaseRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"STANDBY",
		"DISABLED_STANDBY",
	}
}

// GetMappingAutonomousContainerDatabaseRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseRoleEnum(val string) (AutonomousContainerDatabaseRoleEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
