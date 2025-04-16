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

// CreateAutonomousContainerDatabaseDetails Describes the required parameters for the creation of an Autonomous Container Database.
type CreateAutonomousContainerDatabaseDetails struct {

	// The display name for the Autonomous Container Database.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Database Patch model preference.
	PatchModel CreateAutonomousContainerDatabaseDetailsPatchModelEnum `mandatory:"true" json:"patchModel"`

	// Customer Contacts. Setting this to an empty list removes all customer contacts.
	CustomerContacts []CustomerContact `mandatory:"false" json:"customerContacts"`

	// **Deprecated.** The `DB_UNIQUE_NAME` value is set by Oracle Cloud Infrastructure.  Do not specify a value for this parameter. Specifying a value for this field will cause Terraform operations to fail.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// The Database name for the Autonomous Container Database. The name must be unique within the Cloud Autonomous VM Cluster, starting with an alphabetic character, followed by 1 to 7 alphanumeric characters.
	DbName *string `mandatory:"false" json:"dbName"`

	// The service level agreement type of the Autonomous Container Database. The default is STANDARD. For an autonomous dataguard Autonomous Container Database, the specified Autonomous Exadata Infrastructure must be associated with a remote Autonomous Exadata Infrastructure.
	ServiceLevelAgreementType CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum `mandatory:"false" json:"serviceLevelAgreementType,omitempty"`

	// **No longer used.** This parameter is no longer used for Autonomous Database on dedicated Exadata infrasture. Specify a `cloudAutonomousVmClusterId` instead. Using this parameter will cause the operation to fail.
	AutonomousExadataInfrastructureId *string `mandatory:"false" json:"autonomousExadataInfrastructureId"`

	// The base version for the Autonomous Container Database.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	// The Autonomous Database Software Image OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	DatabaseSoftwareImageId *string `mandatory:"false" json:"databaseSoftwareImageId"`

	// *No longer used.* This parameter is no longer used for Autonomous Database on dedicated Exadata infrasture. Specify a `peerCloudAutonomousVmClusterId` instead. Using this parameter will cause the operation to fail.
	PeerAutonomousExadataInfrastructureId *string `mandatory:"false" json:"peerAutonomousExadataInfrastructureId"`

	// The display name for the peer Autonomous Container Database.
	PeerAutonomousContainerDatabaseDisplayName *string `mandatory:"false" json:"peerAutonomousContainerDatabaseDisplayName"`

	// The protection mode of this Autonomous Data Guard association. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode CreateAutonomousContainerDatabaseDetailsProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The lag time for my preference based on data loss tolerance in seconds.
	FastStartFailOverLagLimitInSeconds *int `mandatory:"false" json:"fastStartFailOverLagLimitInSeconds"`

	// Indicates whether Automatic Failover is enabled for Autonomous Container Database Dataguard Association
	IsAutomaticFailoverEnabled *bool `mandatory:"false" json:"isAutomaticFailoverEnabled"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer cloud Autonomous Exadata VM Cluster.
	PeerCloudAutonomousVmClusterId *string `mandatory:"false" json:"peerCloudAutonomousVmClusterId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer Autonomous VM cluster for Autonomous Data Guard. Required to enable Data Guard.
	PeerAutonomousVmClusterId *string `mandatory:"false" json:"peerAutonomousVmClusterId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the standby Autonomous Container Database
	// will be created.
	PeerAutonomousContainerDatabaseCompartmentId *string `mandatory:"false" json:"peerAutonomousContainerDatabaseCompartmentId"`

	PeerAutonomousContainerDatabaseBackupConfig *PeerAutonomousContainerDatabaseBackupConfig `mandatory:"false" json:"peerAutonomousContainerDatabaseBackupConfig"`

	// **Deprecated.** The `DB_UNIQUE_NAME` of the peer Autonomous Container Database in a Data Guard association is set by Oracle Cloud Infrastructure.  Do not specify a value for this parameter. Specifying a value for this field will cause Terraform operations to fail.
	PeerDbUniqueName *string `mandatory:"false" json:"peerDbUniqueName"`

	// The OCID of the Autonomous VM Cluster.
	AutonomousVmClusterId *string `mandatory:"false" json:"autonomousVmClusterId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
	CloudAutonomousVmClusterId *string `mandatory:"false" json:"cloudAutonomousVmClusterId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Autonomous Container Database.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	MaintenanceWindowDetails *MaintenanceWindow `mandatory:"false" json:"maintenanceWindowDetails"`

	// The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database.
	// This value represents the number of days before scheduled maintenance of the primary database.
	StandbyMaintenanceBufferInDays *int `mandatory:"false" json:"standbyMaintenanceBufferInDays"`

	// The next maintenance version preference.
	VersionPreference CreateAutonomousContainerDatabaseDetailsVersionPreferenceEnum `mandatory:"false" json:"versionPreference,omitempty"`

	// Indicates if an automatic DST Time Zone file update is enabled for the Autonomous Container Database. If enabled along with Release Update, patching will be done in a Non-Rolling manner.
	IsDstFileUpdateEnabled *bool `mandatory:"false" json:"isDstFileUpdateEnabled"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	BackupConfig *AutonomousContainerDatabaseBackupConfig `mandatory:"false" json:"backupConfig"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
	KeyStoreId *string `mandatory:"false" json:"keyStoreId"`

	// The CPU value beyond which an Autonomous Database will be opened across multiple nodes. The default value of this attribute is 16 for OCPUs and 64 for ECPUs.
	DbSplitThreshold *int `mandatory:"false" json:"dbSplitThreshold"`

	// The percentage of CPUs reserved across nodes to support node failover. Allowed values are 0%, 25%, and 50%, with 50% being the default option.
	VmFailoverReservation *int `mandatory:"false" json:"vmFailoverReservation"`

	// Determines whether an Autonomous Database must be opened across a minimum or maximum of nodes. By default, Minimum nodes is selected.
	DistributionAffinity CreateAutonomousContainerDatabaseDetailsDistributionAffinityEnum `mandatory:"false" json:"distributionAffinity,omitempty"`

	// Enabling SHARED server architecture enables a database server to allow many client processes to share very few server processes, thereby increasing the number of supported users.
	NetServicesArchitecture CreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum `mandatory:"false" json:"netServicesArchitecture,omitempty"`
}

func (m CreateAutonomousContainerDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAutonomousContainerDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateAutonomousContainerDatabaseDetailsPatchModelEnum(string(m.PatchModel)); !ok && m.PatchModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchModel: %s. Supported values are: %s.", m.PatchModel, strings.Join(GetCreateAutonomousContainerDatabaseDetailsPatchModelEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum(string(m.ServiceLevelAgreementType)); !ok && m.ServiceLevelAgreementType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceLevelAgreementType: %s. Supported values are: %s.", m.ServiceLevelAgreementType, strings.Join(GetCreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateAutonomousContainerDatabaseDetailsProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetCreateAutonomousContainerDatabaseDetailsProtectionModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateAutonomousContainerDatabaseDetailsVersionPreferenceEnum(string(m.VersionPreference)); !ok && m.VersionPreference != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VersionPreference: %s. Supported values are: %s.", m.VersionPreference, strings.Join(GetCreateAutonomousContainerDatabaseDetailsVersionPreferenceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateAutonomousContainerDatabaseDetailsDistributionAffinityEnum(string(m.DistributionAffinity)); !ok && m.DistributionAffinity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DistributionAffinity: %s. Supported values are: %s.", m.DistributionAffinity, strings.Join(GetCreateAutonomousContainerDatabaseDetailsDistributionAffinityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum(string(m.NetServicesArchitecture)); !ok && m.NetServicesArchitecture != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetServicesArchitecture: %s. Supported values are: %s.", m.NetServicesArchitecture, strings.Join(GetCreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum Enum with underlying type: string
type CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum string

// Set of constants representing the allowable values for CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum
const (
	CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeStandard            CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum = "STANDARD"
	CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeAutonomousDataguard CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum = "AUTONOMOUS_DATAGUARD"
)

var mappingCreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum = map[string]CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum{
	"STANDARD":             CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeStandard,
	"AUTONOMOUS_DATAGUARD": CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeAutonomousDataguard,
}

var mappingCreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnumLowerCase = map[string]CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum{
	"standard":             CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeStandard,
	"autonomous_dataguard": CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeAutonomousDataguard,
}

// GetCreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnumValues Enumerates the set of values for CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum
func GetCreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnumValues() []CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum {
	values := make([]CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum, 0)
	for _, v := range mappingCreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnumStringValues Enumerates the set of values in String for CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum
func GetCreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnumStringValues() []string {
	return []string{
		"STANDARD",
		"AUTONOMOUS_DATAGUARD",
	}
}

// GetMappingCreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum(val string) (CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum, bool) {
	enum, ok := mappingCreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateAutonomousContainerDatabaseDetailsProtectionModeEnum Enum with underlying type: string
type CreateAutonomousContainerDatabaseDetailsProtectionModeEnum string

// Set of constants representing the allowable values for CreateAutonomousContainerDatabaseDetailsProtectionModeEnum
const (
	CreateAutonomousContainerDatabaseDetailsProtectionModeAvailability CreateAutonomousContainerDatabaseDetailsProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	CreateAutonomousContainerDatabaseDetailsProtectionModePerformance  CreateAutonomousContainerDatabaseDetailsProtectionModeEnum = "MAXIMUM_PERFORMANCE"
)

var mappingCreateAutonomousContainerDatabaseDetailsProtectionModeEnum = map[string]CreateAutonomousContainerDatabaseDetailsProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": CreateAutonomousContainerDatabaseDetailsProtectionModeAvailability,
	"MAXIMUM_PERFORMANCE":  CreateAutonomousContainerDatabaseDetailsProtectionModePerformance,
}

var mappingCreateAutonomousContainerDatabaseDetailsProtectionModeEnumLowerCase = map[string]CreateAutonomousContainerDatabaseDetailsProtectionModeEnum{
	"maximum_availability": CreateAutonomousContainerDatabaseDetailsProtectionModeAvailability,
	"maximum_performance":  CreateAutonomousContainerDatabaseDetailsProtectionModePerformance,
}

// GetCreateAutonomousContainerDatabaseDetailsProtectionModeEnumValues Enumerates the set of values for CreateAutonomousContainerDatabaseDetailsProtectionModeEnum
func GetCreateAutonomousContainerDatabaseDetailsProtectionModeEnumValues() []CreateAutonomousContainerDatabaseDetailsProtectionModeEnum {
	values := make([]CreateAutonomousContainerDatabaseDetailsProtectionModeEnum, 0)
	for _, v := range mappingCreateAutonomousContainerDatabaseDetailsProtectionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousContainerDatabaseDetailsProtectionModeEnumStringValues Enumerates the set of values in String for CreateAutonomousContainerDatabaseDetailsProtectionModeEnum
func GetCreateAutonomousContainerDatabaseDetailsProtectionModeEnumStringValues() []string {
	return []string{
		"MAXIMUM_AVAILABILITY",
		"MAXIMUM_PERFORMANCE",
	}
}

// GetMappingCreateAutonomousContainerDatabaseDetailsProtectionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousContainerDatabaseDetailsProtectionModeEnum(val string) (CreateAutonomousContainerDatabaseDetailsProtectionModeEnum, bool) {
	enum, ok := mappingCreateAutonomousContainerDatabaseDetailsProtectionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateAutonomousContainerDatabaseDetailsPatchModelEnum Enum with underlying type: string
type CreateAutonomousContainerDatabaseDetailsPatchModelEnum string

// Set of constants representing the allowable values for CreateAutonomousContainerDatabaseDetailsPatchModelEnum
const (
	CreateAutonomousContainerDatabaseDetailsPatchModelUpdates         CreateAutonomousContainerDatabaseDetailsPatchModelEnum = "RELEASE_UPDATES"
	CreateAutonomousContainerDatabaseDetailsPatchModelUpdateRevisions CreateAutonomousContainerDatabaseDetailsPatchModelEnum = "RELEASE_UPDATE_REVISIONS"
)

var mappingCreateAutonomousContainerDatabaseDetailsPatchModelEnum = map[string]CreateAutonomousContainerDatabaseDetailsPatchModelEnum{
	"RELEASE_UPDATES":          CreateAutonomousContainerDatabaseDetailsPatchModelUpdates,
	"RELEASE_UPDATE_REVISIONS": CreateAutonomousContainerDatabaseDetailsPatchModelUpdateRevisions,
}

var mappingCreateAutonomousContainerDatabaseDetailsPatchModelEnumLowerCase = map[string]CreateAutonomousContainerDatabaseDetailsPatchModelEnum{
	"release_updates":          CreateAutonomousContainerDatabaseDetailsPatchModelUpdates,
	"release_update_revisions": CreateAutonomousContainerDatabaseDetailsPatchModelUpdateRevisions,
}

// GetCreateAutonomousContainerDatabaseDetailsPatchModelEnumValues Enumerates the set of values for CreateAutonomousContainerDatabaseDetailsPatchModelEnum
func GetCreateAutonomousContainerDatabaseDetailsPatchModelEnumValues() []CreateAutonomousContainerDatabaseDetailsPatchModelEnum {
	values := make([]CreateAutonomousContainerDatabaseDetailsPatchModelEnum, 0)
	for _, v := range mappingCreateAutonomousContainerDatabaseDetailsPatchModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousContainerDatabaseDetailsPatchModelEnumStringValues Enumerates the set of values in String for CreateAutonomousContainerDatabaseDetailsPatchModelEnum
func GetCreateAutonomousContainerDatabaseDetailsPatchModelEnumStringValues() []string {
	return []string{
		"RELEASE_UPDATES",
		"RELEASE_UPDATE_REVISIONS",
	}
}

// GetMappingCreateAutonomousContainerDatabaseDetailsPatchModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousContainerDatabaseDetailsPatchModelEnum(val string) (CreateAutonomousContainerDatabaseDetailsPatchModelEnum, bool) {
	enum, ok := mappingCreateAutonomousContainerDatabaseDetailsPatchModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateAutonomousContainerDatabaseDetailsVersionPreferenceEnum Enum with underlying type: string
type CreateAutonomousContainerDatabaseDetailsVersionPreferenceEnum string

// Set of constants representing the allowable values for CreateAutonomousContainerDatabaseDetailsVersionPreferenceEnum
const (
	CreateAutonomousContainerDatabaseDetailsVersionPreferenceNextReleaseUpdate   CreateAutonomousContainerDatabaseDetailsVersionPreferenceEnum = "NEXT_RELEASE_UPDATE"
	CreateAutonomousContainerDatabaseDetailsVersionPreferenceLatestReleaseUpdate CreateAutonomousContainerDatabaseDetailsVersionPreferenceEnum = "LATEST_RELEASE_UPDATE"
)

var mappingCreateAutonomousContainerDatabaseDetailsVersionPreferenceEnum = map[string]CreateAutonomousContainerDatabaseDetailsVersionPreferenceEnum{
	"NEXT_RELEASE_UPDATE":   CreateAutonomousContainerDatabaseDetailsVersionPreferenceNextReleaseUpdate,
	"LATEST_RELEASE_UPDATE": CreateAutonomousContainerDatabaseDetailsVersionPreferenceLatestReleaseUpdate,
}

var mappingCreateAutonomousContainerDatabaseDetailsVersionPreferenceEnumLowerCase = map[string]CreateAutonomousContainerDatabaseDetailsVersionPreferenceEnum{
	"next_release_update":   CreateAutonomousContainerDatabaseDetailsVersionPreferenceNextReleaseUpdate,
	"latest_release_update": CreateAutonomousContainerDatabaseDetailsVersionPreferenceLatestReleaseUpdate,
}

// GetCreateAutonomousContainerDatabaseDetailsVersionPreferenceEnumValues Enumerates the set of values for CreateAutonomousContainerDatabaseDetailsVersionPreferenceEnum
func GetCreateAutonomousContainerDatabaseDetailsVersionPreferenceEnumValues() []CreateAutonomousContainerDatabaseDetailsVersionPreferenceEnum {
	values := make([]CreateAutonomousContainerDatabaseDetailsVersionPreferenceEnum, 0)
	for _, v := range mappingCreateAutonomousContainerDatabaseDetailsVersionPreferenceEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousContainerDatabaseDetailsVersionPreferenceEnumStringValues Enumerates the set of values in String for CreateAutonomousContainerDatabaseDetailsVersionPreferenceEnum
func GetCreateAutonomousContainerDatabaseDetailsVersionPreferenceEnumStringValues() []string {
	return []string{
		"NEXT_RELEASE_UPDATE",
		"LATEST_RELEASE_UPDATE",
	}
}

// GetMappingCreateAutonomousContainerDatabaseDetailsVersionPreferenceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousContainerDatabaseDetailsVersionPreferenceEnum(val string) (CreateAutonomousContainerDatabaseDetailsVersionPreferenceEnum, bool) {
	enum, ok := mappingCreateAutonomousContainerDatabaseDetailsVersionPreferenceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateAutonomousContainerDatabaseDetailsDistributionAffinityEnum Enum with underlying type: string
type CreateAutonomousContainerDatabaseDetailsDistributionAffinityEnum string

// Set of constants representing the allowable values for CreateAutonomousContainerDatabaseDetailsDistributionAffinityEnum
const (
	CreateAutonomousContainerDatabaseDetailsDistributionAffinityMinimumDistribution CreateAutonomousContainerDatabaseDetailsDistributionAffinityEnum = "MINIMUM_DISTRIBUTION"
	CreateAutonomousContainerDatabaseDetailsDistributionAffinityMaximumDistribution CreateAutonomousContainerDatabaseDetailsDistributionAffinityEnum = "MAXIMUM_DISTRIBUTION"
)

var mappingCreateAutonomousContainerDatabaseDetailsDistributionAffinityEnum = map[string]CreateAutonomousContainerDatabaseDetailsDistributionAffinityEnum{
	"MINIMUM_DISTRIBUTION": CreateAutonomousContainerDatabaseDetailsDistributionAffinityMinimumDistribution,
	"MAXIMUM_DISTRIBUTION": CreateAutonomousContainerDatabaseDetailsDistributionAffinityMaximumDistribution,
}

var mappingCreateAutonomousContainerDatabaseDetailsDistributionAffinityEnumLowerCase = map[string]CreateAutonomousContainerDatabaseDetailsDistributionAffinityEnum{
	"minimum_distribution": CreateAutonomousContainerDatabaseDetailsDistributionAffinityMinimumDistribution,
	"maximum_distribution": CreateAutonomousContainerDatabaseDetailsDistributionAffinityMaximumDistribution,
}

// GetCreateAutonomousContainerDatabaseDetailsDistributionAffinityEnumValues Enumerates the set of values for CreateAutonomousContainerDatabaseDetailsDistributionAffinityEnum
func GetCreateAutonomousContainerDatabaseDetailsDistributionAffinityEnumValues() []CreateAutonomousContainerDatabaseDetailsDistributionAffinityEnum {
	values := make([]CreateAutonomousContainerDatabaseDetailsDistributionAffinityEnum, 0)
	for _, v := range mappingCreateAutonomousContainerDatabaseDetailsDistributionAffinityEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousContainerDatabaseDetailsDistributionAffinityEnumStringValues Enumerates the set of values in String for CreateAutonomousContainerDatabaseDetailsDistributionAffinityEnum
func GetCreateAutonomousContainerDatabaseDetailsDistributionAffinityEnumStringValues() []string {
	return []string{
		"MINIMUM_DISTRIBUTION",
		"MAXIMUM_DISTRIBUTION",
	}
}

// GetMappingCreateAutonomousContainerDatabaseDetailsDistributionAffinityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousContainerDatabaseDetailsDistributionAffinityEnum(val string) (CreateAutonomousContainerDatabaseDetailsDistributionAffinityEnum, bool) {
	enum, ok := mappingCreateAutonomousContainerDatabaseDetailsDistributionAffinityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum Enum with underlying type: string
type CreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum string

// Set of constants representing the allowable values for CreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum
const (
	CreateAutonomousContainerDatabaseDetailsNetServicesArchitectureDedicated CreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum = "DEDICATED"
	CreateAutonomousContainerDatabaseDetailsNetServicesArchitectureShared    CreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum = "SHARED"
)

var mappingCreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum = map[string]CreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum{
	"DEDICATED": CreateAutonomousContainerDatabaseDetailsNetServicesArchitectureDedicated,
	"SHARED":    CreateAutonomousContainerDatabaseDetailsNetServicesArchitectureShared,
}

var mappingCreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnumLowerCase = map[string]CreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum{
	"dedicated": CreateAutonomousContainerDatabaseDetailsNetServicesArchitectureDedicated,
	"shared":    CreateAutonomousContainerDatabaseDetailsNetServicesArchitectureShared,
}

// GetCreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnumValues Enumerates the set of values for CreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum
func GetCreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnumValues() []CreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum {
	values := make([]CreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum, 0)
	for _, v := range mappingCreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnumStringValues Enumerates the set of values in String for CreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum
func GetCreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnumStringValues() []string {
	return []string{
		"DEDICATED",
		"SHARED",
	}
}

// GetMappingCreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum(val string) (CreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnum, bool) {
	enum, ok := mappingCreateAutonomousContainerDatabaseDetailsNetServicesArchitectureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
