// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
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

	// Customer Contacts. Setting this to an empty list removes all customer contacts.
	CustomerContacts []CustomerContact `mandatory:"false" json:"customerContacts"`

	// **Deprecated.** The `DB_UNIQUE_NAME` value is set by Oracle Cloud Infrastructure.  Do not specify a value for this parameter. Specifying a value for this field will cause Terraform operations to fail.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// The Database name for the Autonomous Container Database. The name must be unique within the Cloud Autonomous VM Cluster, starting with an alphabetic character, followed by 1 to 7 alphanumeric characters.
	DbName *string `mandatory:"false" json:"dbName"`

	// **No longer used.** For Autonomous Database on dedicated Exadata infrastructure, the container database is created within a specified `cloudAutonomousVmCluster`.
	AutonomousExadataInfrastructureId *string `mandatory:"false" json:"autonomousExadataInfrastructureId"`

	// The OCID of the Autonomous VM Cluster.
	AutonomousVmClusterId *string `mandatory:"false" json:"autonomousVmClusterId"`

	// The infrastructure type this resource belongs to.
	InfrastructureType AutonomousContainerDatabaseInfrastructureTypeEnum `mandatory:"false" json:"infrastructureType,omitempty"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
	CloudAutonomousVmClusterId *string `mandatory:"false" json:"cloudAutonomousVmClusterId"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// Key History Entry.
	KeyHistoryEntry []AutonomousDatabaseKeyHistoryEntry `mandatory:"false" json:"keyHistoryEntry"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the Autonomous Container Database was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the Autonomous Container Database will be reverted to Standby from Snapshot Standby.
	TimeSnapshotStandbyRevert *common.SDKTime `mandatory:"false" json:"timeSnapshotStandbyRevert"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last patch applied on the system.
	PatchId *string `mandatory:"false" json:"patchId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance run.
	LastMaintenanceRunId *string `mandatory:"false" json:"lastMaintenanceRunId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next maintenance run.
	NextMaintenanceRunId *string `mandatory:"false" json:"nextMaintenanceRunId"`

	MaintenanceWindow *MaintenanceWindow `mandatory:"false" json:"maintenanceWindow"`

	// The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database.
	// This value represents the number of days before scheduled maintenance of the primary database.
	StandbyMaintenanceBufferInDays *int `mandatory:"false" json:"standbyMaintenanceBufferInDays"`

	// The next maintenance version preference.
	VersionPreference AutonomousContainerDatabaseVersionPreferenceEnum `mandatory:"false" json:"versionPreference,omitempty"`

	// Indicates if an automatic DST Time Zone file update is enabled for the Autonomous Container Database. If enabled along with Release Update, patching will be done in a Non-Rolling manner.
	IsDstFileUpdateEnabled *bool `mandatory:"false" json:"isDstFileUpdateEnabled"`

	// DST Time-Zone File version of the Autonomous Container Database.
	DstFileVersion *string `mandatory:"false" json:"dstFileVersion"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The Data Guard role of the Autonomous Container Database or Autonomous Database, if Autonomous Data Guard is enabled.
	Role AutonomousContainerDatabaseRoleEnum `mandatory:"false" json:"role,omitempty"`

	// The availability domain of the Autonomous Container Database.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Oracle Database version of the Autonomous Container Database.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	BackupConfig *AutonomousContainerDatabaseBackupConfig `mandatory:"false" json:"backupConfig"`

	// This list describes the backup destination properties associated with the Autonomous Container Database (ACD) 's preferred backup destination. The object at a given index is associated with the destination present at the same index in the backup destination details list of the ACD Backup Configuration.
	BackupDestinationPropertiesList []BackupDestinationProperties `mandatory:"false" json:"backupDestinationPropertiesList"`

	// A backup config object holds information about preferred backup destinations only. This object holds information about the associated backup destinations, such as secondary backup destinations created for local backups or remote replicated backups.
	AssociatedBackupConfigurationDetails []BackupDestinationConfigurationSummary `mandatory:"false" json:"associatedBackupConfigurationDetails"`

	RecoveryApplianceDetails *RecoveryApplianceDetails `mandatory:"false" json:"recoveryApplianceDetails"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
	KeyStoreId *string `mandatory:"false" json:"keyStoreId"`

	// The wallet name for Oracle Key Vault.
	KeyStoreWalletName *string `mandatory:"false" json:"keyStoreWalletName"`

	// The amount of memory (in GBs) enabled per ECPU or OCPU in the Autonomous VM Cluster.
	MemoryPerOracleComputeUnitInGBs *int `mandatory:"false" json:"memoryPerOracleComputeUnitInGBs"`

	// Sum of CPUs available on the Autonomous VM Cluster + Sum of reclaimable CPUs available in the Autonomous Container Database.
	AvailableCpus *float32 `mandatory:"false" json:"availableCpus"`

	// The number of CPUs allocated to the Autonomous VM cluster.
	TotalCpus *int `mandatory:"false" json:"totalCpus"`

	// CPUs that continue to be included in the count of CPUs available to the Autonomous Container Database even after one of its Autonomous Database is terminated or scaled down. You can release them to the available CPUs at its parent Autonomous VM Cluster level by restarting the Autonomous Container Database.
	ReclaimableCpus *float32 `mandatory:"false" json:"reclaimableCpus"`

	// An array of CPU values that can be used to successfully provision a single Autonomous Database.
	ProvisionableCpus []float32 `mandatory:"false" json:"provisionableCpus"`

	// List of One-Off patches that has been successfully applied to Autonomous Container Database
	ListOneOffPatches []string `mandatory:"false" json:"listOneOffPatches"`

	// The compute model of the Autonomous Container Database. For Autonomous Database on Dedicated Exadata Infrastructure, the CPU type (ECPUs or OCPUs) is determined by the parent Autonomous Exadata VM Cluster's compute model. ECPU compute model is the recommended model and OCPU compute model is legacy. See Compute Models in Autonomous Database on Dedicated Exadata Infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/adbak) for more details.
	ComputeModel AutonomousContainerDatabaseComputeModelEnum `mandatory:"false" json:"computeModel,omitempty"`

	// The number of CPUs provisioned in an Autonomous Container Database.
	ProvisionedCpus *float32 `mandatory:"false" json:"provisionedCpus"`

	// The number of CPUs reserved in an Autonomous Container Database.
	ReservedCpus *float32 `mandatory:"false" json:"reservedCpus"`

	// The largest Autonomous Database (CPU) that can be created in a new Autonomous Container Database.
	LargestProvisionableAutonomousDatabaseInCpus *float32 `mandatory:"false" json:"largestProvisionableAutonomousDatabaseInCpus"`

	// The timestamp of last successful backup. Here NULL value represents either there are no successful backups or backups are not configured for this Autonomous Container Database.
	TimeOfLastBackup *common.SDKTime `mandatory:"false" json:"timeOfLastBackup"`

	// The CPU value beyond which an Autonomous Database will be opened across multiple nodes. The default value of this attribute is 16 for OCPUs and 64 for ECPUs.
	DbSplitThreshold *int `mandatory:"false" json:"dbSplitThreshold"`

	// The percentage of CPUs reserved across nodes to support node failover. Allowed values are 0%, 25%, and 50%, with 50% being the default option.
	VmFailoverReservation *int `mandatory:"false" json:"vmFailoverReservation"`

	// Determines whether an Autonomous Database must be opened across the maximum number of nodes or the least number of nodes. By default, Minimum nodes is selected.
	DistributionAffinity AutonomousContainerDatabaseDistributionAffinityEnum `mandatory:"false" json:"distributionAffinity,omitempty"`

	// Enabling SHARED server architecture enables a database server to allow many client processes to share very few server processes, thereby increasing the number of supported users.
	NetServicesArchitecture AutonomousContainerDatabaseNetServicesArchitectureEnum `mandatory:"false" json:"netServicesArchitecture,omitempty"`

	// Whether it is multiple standby Autonomous Dataguard
	IsMultipleStandby *bool `mandatory:"false" json:"isMultipleStandby"`

	// **Deprecated.** Indicates whether the Autonomous Database has local (in-region) Data Guard enabled. Not applicable to cross-region Autonomous Data Guard associations, or to Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure.
	IsDataGuardEnabled *bool `mandatory:"false" json:"isDataGuardEnabled"`

	Dataguard *AutonomousContainerDatabaseDataguard `mandatory:"false" json:"dataguard"`

	// Array of Dg associations.
	DataguardGroupMembers []AutonomousContainerDatabaseDataguard `mandatory:"false" json:"dataguardGroupMembers"`
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
	if _, ok := GetMappingAutonomousContainerDatabaseVersionPreferenceEnum(string(m.VersionPreference)); !ok && m.VersionPreference != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VersionPreference: %s. Supported values are: %s.", m.VersionPreference, strings.Join(GetAutonomousContainerDatabaseVersionPreferenceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousContainerDatabaseRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetAutonomousContainerDatabaseRoleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousContainerDatabaseComputeModelEnum(string(m.ComputeModel)); !ok && m.ComputeModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComputeModel: %s. Supported values are: %s.", m.ComputeModel, strings.Join(GetAutonomousContainerDatabaseComputeModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousContainerDatabaseDistributionAffinityEnum(string(m.DistributionAffinity)); !ok && m.DistributionAffinity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DistributionAffinity: %s. Supported values are: %s.", m.DistributionAffinity, strings.Join(GetAutonomousContainerDatabaseDistributionAffinityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousContainerDatabaseNetServicesArchitectureEnum(string(m.NetServicesArchitecture)); !ok && m.NetServicesArchitecture != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetServicesArchitecture: %s. Supported values are: %s.", m.NetServicesArchitecture, strings.Join(GetAutonomousContainerDatabaseNetServicesArchitectureEnumStringValues(), ",")))
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
	AutonomousContainerDatabaseLifecycleStateProvisioning                AutonomousContainerDatabaseLifecycleStateEnum = "PROVISIONING"
	AutonomousContainerDatabaseLifecycleStateAvailable                   AutonomousContainerDatabaseLifecycleStateEnum = "AVAILABLE"
	AutonomousContainerDatabaseLifecycleStateUpdating                    AutonomousContainerDatabaseLifecycleStateEnum = "UPDATING"
	AutonomousContainerDatabaseLifecycleStateTerminating                 AutonomousContainerDatabaseLifecycleStateEnum = "TERMINATING"
	AutonomousContainerDatabaseLifecycleStateTerminated                  AutonomousContainerDatabaseLifecycleStateEnum = "TERMINATED"
	AutonomousContainerDatabaseLifecycleStateFailed                      AutonomousContainerDatabaseLifecycleStateEnum = "FAILED"
	AutonomousContainerDatabaseLifecycleStateBackupInProgress            AutonomousContainerDatabaseLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	AutonomousContainerDatabaseLifecycleStateRestoring                   AutonomousContainerDatabaseLifecycleStateEnum = "RESTORING"
	AutonomousContainerDatabaseLifecycleStateRestoreFailed               AutonomousContainerDatabaseLifecycleStateEnum = "RESTORE_FAILED"
	AutonomousContainerDatabaseLifecycleStateRestarting                  AutonomousContainerDatabaseLifecycleStateEnum = "RESTARTING"
	AutonomousContainerDatabaseLifecycleStateMaintenanceInProgress       AutonomousContainerDatabaseLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
	AutonomousContainerDatabaseLifecycleStateRoleChangeInProgress        AutonomousContainerDatabaseLifecycleStateEnum = "ROLE_CHANGE_IN_PROGRESS"
	AutonomousContainerDatabaseLifecycleStateEnablingAutonomousDataGuard AutonomousContainerDatabaseLifecycleStateEnum = "ENABLING_AUTONOMOUS_DATA_GUARD"
	AutonomousContainerDatabaseLifecycleStateUnavailable                 AutonomousContainerDatabaseLifecycleStateEnum = "UNAVAILABLE"
)

var mappingAutonomousContainerDatabaseLifecycleStateEnum = map[string]AutonomousContainerDatabaseLifecycleStateEnum{
	"PROVISIONING":                   AutonomousContainerDatabaseLifecycleStateProvisioning,
	"AVAILABLE":                      AutonomousContainerDatabaseLifecycleStateAvailable,
	"UPDATING":                       AutonomousContainerDatabaseLifecycleStateUpdating,
	"TERMINATING":                    AutonomousContainerDatabaseLifecycleStateTerminating,
	"TERMINATED":                     AutonomousContainerDatabaseLifecycleStateTerminated,
	"FAILED":                         AutonomousContainerDatabaseLifecycleStateFailed,
	"BACKUP_IN_PROGRESS":             AutonomousContainerDatabaseLifecycleStateBackupInProgress,
	"RESTORING":                      AutonomousContainerDatabaseLifecycleStateRestoring,
	"RESTORE_FAILED":                 AutonomousContainerDatabaseLifecycleStateRestoreFailed,
	"RESTARTING":                     AutonomousContainerDatabaseLifecycleStateRestarting,
	"MAINTENANCE_IN_PROGRESS":        AutonomousContainerDatabaseLifecycleStateMaintenanceInProgress,
	"ROLE_CHANGE_IN_PROGRESS":        AutonomousContainerDatabaseLifecycleStateRoleChangeInProgress,
	"ENABLING_AUTONOMOUS_DATA_GUARD": AutonomousContainerDatabaseLifecycleStateEnablingAutonomousDataGuard,
	"UNAVAILABLE":                    AutonomousContainerDatabaseLifecycleStateUnavailable,
}

var mappingAutonomousContainerDatabaseLifecycleStateEnumLowerCase = map[string]AutonomousContainerDatabaseLifecycleStateEnum{
	"provisioning":                   AutonomousContainerDatabaseLifecycleStateProvisioning,
	"available":                      AutonomousContainerDatabaseLifecycleStateAvailable,
	"updating":                       AutonomousContainerDatabaseLifecycleStateUpdating,
	"terminating":                    AutonomousContainerDatabaseLifecycleStateTerminating,
	"terminated":                     AutonomousContainerDatabaseLifecycleStateTerminated,
	"failed":                         AutonomousContainerDatabaseLifecycleStateFailed,
	"backup_in_progress":             AutonomousContainerDatabaseLifecycleStateBackupInProgress,
	"restoring":                      AutonomousContainerDatabaseLifecycleStateRestoring,
	"restore_failed":                 AutonomousContainerDatabaseLifecycleStateRestoreFailed,
	"restarting":                     AutonomousContainerDatabaseLifecycleStateRestarting,
	"maintenance_in_progress":        AutonomousContainerDatabaseLifecycleStateMaintenanceInProgress,
	"role_change_in_progress":        AutonomousContainerDatabaseLifecycleStateRoleChangeInProgress,
	"enabling_autonomous_data_guard": AutonomousContainerDatabaseLifecycleStateEnablingAutonomousDataGuard,
	"unavailable":                    AutonomousContainerDatabaseLifecycleStateUnavailable,
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
		"ENABLING_AUTONOMOUS_DATA_GUARD",
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

// AutonomousContainerDatabaseVersionPreferenceEnum Enum with underlying type: string
type AutonomousContainerDatabaseVersionPreferenceEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseVersionPreferenceEnum
const (
	AutonomousContainerDatabaseVersionPreferenceNextReleaseUpdate   AutonomousContainerDatabaseVersionPreferenceEnum = "NEXT_RELEASE_UPDATE"
	AutonomousContainerDatabaseVersionPreferenceLatestReleaseUpdate AutonomousContainerDatabaseVersionPreferenceEnum = "LATEST_RELEASE_UPDATE"
)

var mappingAutonomousContainerDatabaseVersionPreferenceEnum = map[string]AutonomousContainerDatabaseVersionPreferenceEnum{
	"NEXT_RELEASE_UPDATE":   AutonomousContainerDatabaseVersionPreferenceNextReleaseUpdate,
	"LATEST_RELEASE_UPDATE": AutonomousContainerDatabaseVersionPreferenceLatestReleaseUpdate,
}

var mappingAutonomousContainerDatabaseVersionPreferenceEnumLowerCase = map[string]AutonomousContainerDatabaseVersionPreferenceEnum{
	"next_release_update":   AutonomousContainerDatabaseVersionPreferenceNextReleaseUpdate,
	"latest_release_update": AutonomousContainerDatabaseVersionPreferenceLatestReleaseUpdate,
}

// GetAutonomousContainerDatabaseVersionPreferenceEnumValues Enumerates the set of values for AutonomousContainerDatabaseVersionPreferenceEnum
func GetAutonomousContainerDatabaseVersionPreferenceEnumValues() []AutonomousContainerDatabaseVersionPreferenceEnum {
	values := make([]AutonomousContainerDatabaseVersionPreferenceEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseVersionPreferenceEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseVersionPreferenceEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseVersionPreferenceEnum
func GetAutonomousContainerDatabaseVersionPreferenceEnumStringValues() []string {
	return []string{
		"NEXT_RELEASE_UPDATE",
		"LATEST_RELEASE_UPDATE",
	}
}

// GetMappingAutonomousContainerDatabaseVersionPreferenceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseVersionPreferenceEnum(val string) (AutonomousContainerDatabaseVersionPreferenceEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseVersionPreferenceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousContainerDatabaseRoleEnum Enum with underlying type: string
type AutonomousContainerDatabaseRoleEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseRoleEnum
const (
	AutonomousContainerDatabaseRolePrimary         AutonomousContainerDatabaseRoleEnum = "PRIMARY"
	AutonomousContainerDatabaseRoleStandby         AutonomousContainerDatabaseRoleEnum = "STANDBY"
	AutonomousContainerDatabaseRoleDisabledStandby AutonomousContainerDatabaseRoleEnum = "DISABLED_STANDBY"
	AutonomousContainerDatabaseRoleBackupCopy      AutonomousContainerDatabaseRoleEnum = "BACKUP_COPY"
	AutonomousContainerDatabaseRoleSnapshotStandby AutonomousContainerDatabaseRoleEnum = "SNAPSHOT_STANDBY"
)

var mappingAutonomousContainerDatabaseRoleEnum = map[string]AutonomousContainerDatabaseRoleEnum{
	"PRIMARY":          AutonomousContainerDatabaseRolePrimary,
	"STANDBY":          AutonomousContainerDatabaseRoleStandby,
	"DISABLED_STANDBY": AutonomousContainerDatabaseRoleDisabledStandby,
	"BACKUP_COPY":      AutonomousContainerDatabaseRoleBackupCopy,
	"SNAPSHOT_STANDBY": AutonomousContainerDatabaseRoleSnapshotStandby,
}

var mappingAutonomousContainerDatabaseRoleEnumLowerCase = map[string]AutonomousContainerDatabaseRoleEnum{
	"primary":          AutonomousContainerDatabaseRolePrimary,
	"standby":          AutonomousContainerDatabaseRoleStandby,
	"disabled_standby": AutonomousContainerDatabaseRoleDisabledStandby,
	"backup_copy":      AutonomousContainerDatabaseRoleBackupCopy,
	"snapshot_standby": AutonomousContainerDatabaseRoleSnapshotStandby,
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
		"BACKUP_COPY",
		"SNAPSHOT_STANDBY",
	}
}

// GetMappingAutonomousContainerDatabaseRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseRoleEnum(val string) (AutonomousContainerDatabaseRoleEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousContainerDatabaseComputeModelEnum Enum with underlying type: string
type AutonomousContainerDatabaseComputeModelEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseComputeModelEnum
const (
	AutonomousContainerDatabaseComputeModelEcpu AutonomousContainerDatabaseComputeModelEnum = "ECPU"
	AutonomousContainerDatabaseComputeModelOcpu AutonomousContainerDatabaseComputeModelEnum = "OCPU"
)

var mappingAutonomousContainerDatabaseComputeModelEnum = map[string]AutonomousContainerDatabaseComputeModelEnum{
	"ECPU": AutonomousContainerDatabaseComputeModelEcpu,
	"OCPU": AutonomousContainerDatabaseComputeModelOcpu,
}

var mappingAutonomousContainerDatabaseComputeModelEnumLowerCase = map[string]AutonomousContainerDatabaseComputeModelEnum{
	"ecpu": AutonomousContainerDatabaseComputeModelEcpu,
	"ocpu": AutonomousContainerDatabaseComputeModelOcpu,
}

// GetAutonomousContainerDatabaseComputeModelEnumValues Enumerates the set of values for AutonomousContainerDatabaseComputeModelEnum
func GetAutonomousContainerDatabaseComputeModelEnumValues() []AutonomousContainerDatabaseComputeModelEnum {
	values := make([]AutonomousContainerDatabaseComputeModelEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseComputeModelEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseComputeModelEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseComputeModelEnum
func GetAutonomousContainerDatabaseComputeModelEnumStringValues() []string {
	return []string{
		"ECPU",
		"OCPU",
	}
}

// GetMappingAutonomousContainerDatabaseComputeModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseComputeModelEnum(val string) (AutonomousContainerDatabaseComputeModelEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseComputeModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousContainerDatabaseDistributionAffinityEnum Enum with underlying type: string
type AutonomousContainerDatabaseDistributionAffinityEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseDistributionAffinityEnum
const (
	AutonomousContainerDatabaseDistributionAffinityMinimumDistribution AutonomousContainerDatabaseDistributionAffinityEnum = "MINIMUM_DISTRIBUTION"
	AutonomousContainerDatabaseDistributionAffinityMaximumDistribution AutonomousContainerDatabaseDistributionAffinityEnum = "MAXIMUM_DISTRIBUTION"
)

var mappingAutonomousContainerDatabaseDistributionAffinityEnum = map[string]AutonomousContainerDatabaseDistributionAffinityEnum{
	"MINIMUM_DISTRIBUTION": AutonomousContainerDatabaseDistributionAffinityMinimumDistribution,
	"MAXIMUM_DISTRIBUTION": AutonomousContainerDatabaseDistributionAffinityMaximumDistribution,
}

var mappingAutonomousContainerDatabaseDistributionAffinityEnumLowerCase = map[string]AutonomousContainerDatabaseDistributionAffinityEnum{
	"minimum_distribution": AutonomousContainerDatabaseDistributionAffinityMinimumDistribution,
	"maximum_distribution": AutonomousContainerDatabaseDistributionAffinityMaximumDistribution,
}

// GetAutonomousContainerDatabaseDistributionAffinityEnumValues Enumerates the set of values for AutonomousContainerDatabaseDistributionAffinityEnum
func GetAutonomousContainerDatabaseDistributionAffinityEnumValues() []AutonomousContainerDatabaseDistributionAffinityEnum {
	values := make([]AutonomousContainerDatabaseDistributionAffinityEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseDistributionAffinityEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseDistributionAffinityEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseDistributionAffinityEnum
func GetAutonomousContainerDatabaseDistributionAffinityEnumStringValues() []string {
	return []string{
		"MINIMUM_DISTRIBUTION",
		"MAXIMUM_DISTRIBUTION",
	}
}

// GetMappingAutonomousContainerDatabaseDistributionAffinityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseDistributionAffinityEnum(val string) (AutonomousContainerDatabaseDistributionAffinityEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseDistributionAffinityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousContainerDatabaseNetServicesArchitectureEnum Enum with underlying type: string
type AutonomousContainerDatabaseNetServicesArchitectureEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseNetServicesArchitectureEnum
const (
	AutonomousContainerDatabaseNetServicesArchitectureDedicated AutonomousContainerDatabaseNetServicesArchitectureEnum = "DEDICATED"
	AutonomousContainerDatabaseNetServicesArchitectureShared    AutonomousContainerDatabaseNetServicesArchitectureEnum = "SHARED"
)

var mappingAutonomousContainerDatabaseNetServicesArchitectureEnum = map[string]AutonomousContainerDatabaseNetServicesArchitectureEnum{
	"DEDICATED": AutonomousContainerDatabaseNetServicesArchitectureDedicated,
	"SHARED":    AutonomousContainerDatabaseNetServicesArchitectureShared,
}

var mappingAutonomousContainerDatabaseNetServicesArchitectureEnumLowerCase = map[string]AutonomousContainerDatabaseNetServicesArchitectureEnum{
	"dedicated": AutonomousContainerDatabaseNetServicesArchitectureDedicated,
	"shared":    AutonomousContainerDatabaseNetServicesArchitectureShared,
}

// GetAutonomousContainerDatabaseNetServicesArchitectureEnumValues Enumerates the set of values for AutonomousContainerDatabaseNetServicesArchitectureEnum
func GetAutonomousContainerDatabaseNetServicesArchitectureEnumValues() []AutonomousContainerDatabaseNetServicesArchitectureEnum {
	values := make([]AutonomousContainerDatabaseNetServicesArchitectureEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseNetServicesArchitectureEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseNetServicesArchitectureEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseNetServicesArchitectureEnum
func GetAutonomousContainerDatabaseNetServicesArchitectureEnumStringValues() []string {
	return []string{
		"DEDICATED",
		"SHARED",
	}
}

// GetMappingAutonomousContainerDatabaseNetServicesArchitectureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseNetServicesArchitectureEnum(val string) (AutonomousContainerDatabaseNetServicesArchitectureEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseNetServicesArchitectureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
