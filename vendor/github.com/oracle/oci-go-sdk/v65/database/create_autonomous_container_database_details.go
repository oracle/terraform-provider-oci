// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateAutonomousContainerDatabaseDetails Details to create an Autonomous Container Database (ACD).
type CreateAutonomousContainerDatabaseDetails struct {

	// The display name for the Autonomous Container Database.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Customer Contacts. Setting this to an empty list removes all customer contacts.
	CustomerContacts []CustomerContact `mandatory:"false" json:"customerContacts"`

	// The OKV End Point Group name for the Autonomous Container Database.
	OkvEndPointGroupName *string `mandatory:"false" json:"okvEndPointGroupName"`

	// **Deprecated.** The `DB_UNIQUE_NAME` value is set by Oracle Cloud Infrastructure.  Do not specify a value for this parameter. Specifying a value for this field will cause Terraform operations to fail.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// The Database name for the Autonomous Container Database. The name must be unique within the Cloud Autonomous VM Cluster, starting with an alphabetic character, followed by 1 to 7 alphanumeric characters.
	DbName *string `mandatory:"false" json:"dbName"`

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

	// The service level agreement type of the Autonomous Container Database. The default is STANDARD. For an autonomous dataguard Autonomous Container Database, the specified Autonomous Exadata Infrastructure must be associated with a remote Autonomous Exadata Infrastructure.
	ServiceLevelAgreementType CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum `mandatory:"false" json:"serviceLevelAgreementType,omitempty"`

	// The protection mode of this Autonomous Data Guard association. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode CreateAutonomousContainerDatabaseBaseProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// Database Patch model preference.
	PatchModel CreateAutonomousContainerDatabaseBasePatchModelEnum `mandatory:"true" json:"patchModel"`

	// The next maintenance version preference.
	VersionPreference CreateAutonomousContainerDatabaseBaseVersionPreferenceEnum `mandatory:"false" json:"versionPreference,omitempty"`

	// Determines whether an Autonomous Database must be opened across a minimum or maximum of nodes. By default, Minimum nodes is selected.
	DistributionAffinity CreateAutonomousContainerDatabaseBaseDistributionAffinityEnum `mandatory:"false" json:"distributionAffinity,omitempty"`

	// Enabling SHARED server architecture enables a database server to allow many client processes to share very few server processes, thereby increasing the number of supported users.
	NetServicesArchitecture CreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum `mandatory:"false" json:"netServicesArchitecture,omitempty"`
}

// GetCustomerContacts returns CustomerContacts
func (m CreateAutonomousContainerDatabaseDetails) GetCustomerContacts() []CustomerContact {
	return m.CustomerContacts
}

// GetOkvEndPointGroupName returns OkvEndPointGroupName
func (m CreateAutonomousContainerDatabaseDetails) GetOkvEndPointGroupName() *string {
	return m.OkvEndPointGroupName
}

// GetDisplayName returns DisplayName
func (m CreateAutonomousContainerDatabaseDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDbUniqueName returns DbUniqueName
func (m CreateAutonomousContainerDatabaseDetails) GetDbUniqueName() *string {
	return m.DbUniqueName
}

// GetDbName returns DbName
func (m CreateAutonomousContainerDatabaseDetails) GetDbName() *string {
	return m.DbName
}

// GetServiceLevelAgreementType returns ServiceLevelAgreementType
func (m CreateAutonomousContainerDatabaseDetails) GetServiceLevelAgreementType() CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum {
	return m.ServiceLevelAgreementType
}

// GetAutonomousExadataInfrastructureId returns AutonomousExadataInfrastructureId
func (m CreateAutonomousContainerDatabaseDetails) GetAutonomousExadataInfrastructureId() *string {
	return m.AutonomousExadataInfrastructureId
}

// GetDbVersion returns DbVersion
func (m CreateAutonomousContainerDatabaseDetails) GetDbVersion() *string {
	return m.DbVersion
}

// GetDatabaseSoftwareImageId returns DatabaseSoftwareImageId
func (m CreateAutonomousContainerDatabaseDetails) GetDatabaseSoftwareImageId() *string {
	return m.DatabaseSoftwareImageId
}

// GetPeerAutonomousExadataInfrastructureId returns PeerAutonomousExadataInfrastructureId
func (m CreateAutonomousContainerDatabaseDetails) GetPeerAutonomousExadataInfrastructureId() *string {
	return m.PeerAutonomousExadataInfrastructureId
}

// GetPeerAutonomousContainerDatabaseDisplayName returns PeerAutonomousContainerDatabaseDisplayName
func (m CreateAutonomousContainerDatabaseDetails) GetPeerAutonomousContainerDatabaseDisplayName() *string {
	return m.PeerAutonomousContainerDatabaseDisplayName
}

// GetProtectionMode returns ProtectionMode
func (m CreateAutonomousContainerDatabaseDetails) GetProtectionMode() CreateAutonomousContainerDatabaseBaseProtectionModeEnum {
	return m.ProtectionMode
}

// GetFastStartFailOverLagLimitInSeconds returns FastStartFailOverLagLimitInSeconds
func (m CreateAutonomousContainerDatabaseDetails) GetFastStartFailOverLagLimitInSeconds() *int {
	return m.FastStartFailOverLagLimitInSeconds
}

// GetIsAutomaticFailoverEnabled returns IsAutomaticFailoverEnabled
func (m CreateAutonomousContainerDatabaseDetails) GetIsAutomaticFailoverEnabled() *bool {
	return m.IsAutomaticFailoverEnabled
}

// GetPeerCloudAutonomousVmClusterId returns PeerCloudAutonomousVmClusterId
func (m CreateAutonomousContainerDatabaseDetails) GetPeerCloudAutonomousVmClusterId() *string {
	return m.PeerCloudAutonomousVmClusterId
}

// GetPeerAutonomousVmClusterId returns PeerAutonomousVmClusterId
func (m CreateAutonomousContainerDatabaseDetails) GetPeerAutonomousVmClusterId() *string {
	return m.PeerAutonomousVmClusterId
}

// GetPeerAutonomousContainerDatabaseCompartmentId returns PeerAutonomousContainerDatabaseCompartmentId
func (m CreateAutonomousContainerDatabaseDetails) GetPeerAutonomousContainerDatabaseCompartmentId() *string {
	return m.PeerAutonomousContainerDatabaseCompartmentId
}

// GetPeerAutonomousContainerDatabaseBackupConfig returns PeerAutonomousContainerDatabaseBackupConfig
func (m CreateAutonomousContainerDatabaseDetails) GetPeerAutonomousContainerDatabaseBackupConfig() *PeerAutonomousContainerDatabaseBackupConfig {
	return m.PeerAutonomousContainerDatabaseBackupConfig
}

// GetPeerDbUniqueName returns PeerDbUniqueName
func (m CreateAutonomousContainerDatabaseDetails) GetPeerDbUniqueName() *string {
	return m.PeerDbUniqueName
}

// GetAutonomousVmClusterId returns AutonomousVmClusterId
func (m CreateAutonomousContainerDatabaseDetails) GetAutonomousVmClusterId() *string {
	return m.AutonomousVmClusterId
}

// GetCloudAutonomousVmClusterId returns CloudAutonomousVmClusterId
func (m CreateAutonomousContainerDatabaseDetails) GetCloudAutonomousVmClusterId() *string {
	return m.CloudAutonomousVmClusterId
}

// GetCompartmentId returns CompartmentId
func (m CreateAutonomousContainerDatabaseDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetPatchModel returns PatchModel
func (m CreateAutonomousContainerDatabaseDetails) GetPatchModel() CreateAutonomousContainerDatabaseBasePatchModelEnum {
	return m.PatchModel
}

// GetMaintenanceWindowDetails returns MaintenanceWindowDetails
func (m CreateAutonomousContainerDatabaseDetails) GetMaintenanceWindowDetails() *MaintenanceWindow {
	return m.MaintenanceWindowDetails
}

// GetStandbyMaintenanceBufferInDays returns StandbyMaintenanceBufferInDays
func (m CreateAutonomousContainerDatabaseDetails) GetStandbyMaintenanceBufferInDays() *int {
	return m.StandbyMaintenanceBufferInDays
}

// GetVersionPreference returns VersionPreference
func (m CreateAutonomousContainerDatabaseDetails) GetVersionPreference() CreateAutonomousContainerDatabaseBaseVersionPreferenceEnum {
	return m.VersionPreference
}

// GetIsDstFileUpdateEnabled returns IsDstFileUpdateEnabled
func (m CreateAutonomousContainerDatabaseDetails) GetIsDstFileUpdateEnabled() *bool {
	return m.IsDstFileUpdateEnabled
}

// GetFreeformTags returns FreeformTags
func (m CreateAutonomousContainerDatabaseDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateAutonomousContainerDatabaseDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetBackupConfig returns BackupConfig
func (m CreateAutonomousContainerDatabaseDetails) GetBackupConfig() *AutonomousContainerDatabaseBackupConfig {
	return m.BackupConfig
}

// GetKmsKeyId returns KmsKeyId
func (m CreateAutonomousContainerDatabaseDetails) GetKmsKeyId() *string {
	return m.KmsKeyId
}

// GetKmsKeyVersionId returns KmsKeyVersionId
func (m CreateAutonomousContainerDatabaseDetails) GetKmsKeyVersionId() *string {
	return m.KmsKeyVersionId
}

// GetVaultId returns VaultId
func (m CreateAutonomousContainerDatabaseDetails) GetVaultId() *string {
	return m.VaultId
}

// GetKeyStoreId returns KeyStoreId
func (m CreateAutonomousContainerDatabaseDetails) GetKeyStoreId() *string {
	return m.KeyStoreId
}

// GetDbSplitThreshold returns DbSplitThreshold
func (m CreateAutonomousContainerDatabaseDetails) GetDbSplitThreshold() *int {
	return m.DbSplitThreshold
}

// GetVmFailoverReservation returns VmFailoverReservation
func (m CreateAutonomousContainerDatabaseDetails) GetVmFailoverReservation() *int {
	return m.VmFailoverReservation
}

// GetDistributionAffinity returns DistributionAffinity
func (m CreateAutonomousContainerDatabaseDetails) GetDistributionAffinity() CreateAutonomousContainerDatabaseBaseDistributionAffinityEnum {
	return m.DistributionAffinity
}

// GetNetServicesArchitecture returns NetServicesArchitecture
func (m CreateAutonomousContainerDatabaseDetails) GetNetServicesArchitecture() CreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum {
	return m.NetServicesArchitecture
}

func (m CreateAutonomousContainerDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAutonomousContainerDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum(string(m.ServiceLevelAgreementType)); !ok && m.ServiceLevelAgreementType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceLevelAgreementType: %s. Supported values are: %s.", m.ServiceLevelAgreementType, strings.Join(GetCreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateAutonomousContainerDatabaseBaseProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetCreateAutonomousContainerDatabaseBaseProtectionModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateAutonomousContainerDatabaseBasePatchModelEnum(string(m.PatchModel)); !ok && m.PatchModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchModel: %s. Supported values are: %s.", m.PatchModel, strings.Join(GetCreateAutonomousContainerDatabaseBasePatchModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateAutonomousContainerDatabaseBaseVersionPreferenceEnum(string(m.VersionPreference)); !ok && m.VersionPreference != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VersionPreference: %s. Supported values are: %s.", m.VersionPreference, strings.Join(GetCreateAutonomousContainerDatabaseBaseVersionPreferenceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateAutonomousContainerDatabaseBaseDistributionAffinityEnum(string(m.DistributionAffinity)); !ok && m.DistributionAffinity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DistributionAffinity: %s. Supported values are: %s.", m.DistributionAffinity, strings.Join(GetCreateAutonomousContainerDatabaseBaseDistributionAffinityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum(string(m.NetServicesArchitecture)); !ok && m.NetServicesArchitecture != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetServicesArchitecture: %s. Supported values are: %s.", m.NetServicesArchitecture, strings.Join(GetCreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateAutonomousContainerDatabaseDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateAutonomousContainerDatabaseDetails CreateAutonomousContainerDatabaseDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateAutonomousContainerDatabaseDetails
	}{
		"NONE",
		(MarshalTypeCreateAutonomousContainerDatabaseDetails)(m),
	}

	return json.Marshal(&s)
}
