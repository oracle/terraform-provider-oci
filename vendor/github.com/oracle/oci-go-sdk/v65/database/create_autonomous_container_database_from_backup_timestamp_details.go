// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// CreateAutonomousContainerDatabaseFromBackupTimestampDetails Details to create an Autonomous Container Database (ACD) by cloning backup available, of an existing ACD, at the requested time stamp.
type CreateAutonomousContainerDatabaseFromBackupTimestampDetails struct {

	// The display name for the Autonomous Container Database.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source ACD that you will clone to create a new ACD.
	SourceAutonomousContainerDatabaseId *string `mandatory:"true" json:"sourceAutonomousContainerDatabaseId"`

	NfsStorageDetails *NfsStorageDetails `mandatory:"false" json:"nfsStorageDetails"`

	// Customer Contacts. Setting this to an empty list removes all customer contacts.
	CustomerContacts []CustomerContact `mandatory:"false" json:"customerContacts"`

	// The OKV End Point Group name for the Autonomous Container Database.
	OkvEndPointGroupName *string `mandatory:"false" json:"okvEndPointGroupName"`

	// **Deprecated.** The `DB_UNIQUE_NAME` value is set by Oracle Cloud Infrastructure.  Do not specify a value for this parameter. Specifying a value for this field will cause Terraform operations to fail.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// The Database name for the Autonomous Container Database. The name must be unique within the Cloud Autonomous VM Cluster, starting with an alphabetic character, followed by 1 to 7 alphanumeric characters.
	DbName *string `mandatory:"false" json:"dbName"`

	// **No longer used.** This parameter is no longer used for Autonomous AI Database on dedicated Exadata infrasture. Specify a `cloudAutonomousVmClusterId` instead. Using this parameter will cause the operation to fail.
	AutonomousExadataInfrastructureId *string `mandatory:"false" json:"autonomousExadataInfrastructureId"`

	// The base version for the Autonomous Container Database.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	// The Autonomous AI Database Software Image OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	DatabaseSoftwareImageId *string `mandatory:"false" json:"databaseSoftwareImageId"`

	// *No longer used.* This parameter is no longer used for Autonomous AI Database on dedicated Exadata infrasture. Specify a `peerCloudAutonomousVmClusterId` instead. Using this parameter will cause the operation to fail.
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

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous AI Database Serverless does not use key versions, hence is not applicable for Autonomous AI Database Serverless instances.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
	KeyStoreId *string `mandatory:"false" json:"keyStoreId"`

	EncryptionKeyLocationDetails EncryptionKeyLocationDetails `mandatory:"false" json:"encryptionKeyLocationDetails"`

	// Indicates if FIPS-140 compliant cryptography is enabled for the Autonomous Container Database. The default
	// value is `TRUE` for regions in Oracle Cloud's Government, ONSR realms (https://docs.oracle.com/iaas/Content/General/Concepts/regions.htm),
	// and `FALSE` for regions in Oracle Cloud's commercial realm.
	IsFipsEnabled *bool `mandatory:"false" json:"isFipsEnabled"`

	// The CPU value beyond which an Autonomous AI Database will be opened across multiple nodes. The default value of this attribute is 16 for OCPUs and 64 for ECPUs.
	DbSplitThreshold *int `mandatory:"false" json:"dbSplitThreshold"`

	// The percentage of CPUs reserved across nodes to support node failover. Allowed values are 0%, 25%, 50%, 75%, and 100%, with 50% being the default option.
	VmFailoverReservation *int `mandatory:"false" json:"vmFailoverReservation"`

	// The time stamp representing the point in time to which the Autonomous Container Database should be cloned from backup. And the requested timeStamp should be in the past.
	TimeStampToUseForCloning *common.SDKTime `mandatory:"false" json:"timeStampToUseForCloning"`

	// If set to true, OCI shall attempt to create a point in time to the latest available backup of the source Autonomous Container Database.
	ShouldUseLatestAvailableBackupTimeStamp *bool `mandatory:"false" json:"shouldUseLatestAvailableBackupTimeStamp"`

	// A list of Autonomous Databases ( display name of the ADB in specific ) to be cloned from backup of the source Autonomous Container Database.
	AutonomousDatabasesToClone []string `mandatory:"false" json:"autonomousDatabasesToClone"`

	// The Autonomous AI Database clone type.
	CloneType CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnum `mandatory:"true" json:"cloneType"`

	// The speed at which the Autonomous Container Database Clone from backup operation to be performed by OCI.
	CloneBandWidth CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnum `mandatory:"false" json:"cloneBandWidth,omitempty"`

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

	// Determines whether an Autonomous AI Database must be opened across a minimum or maximum of nodes. By default, Minimum nodes is selected.
	DistributionAffinity CreateAutonomousContainerDatabaseBaseDistributionAffinityEnum `mandatory:"false" json:"distributionAffinity,omitempty"`

	// Enabling SHARED server architecture enables a database server to allow many client processes to share very few server processes, thereby increasing the number of supported users.
	NetServicesArchitecture CreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum `mandatory:"false" json:"netServicesArchitecture,omitempty"`
}

// GetNfsStorageDetails returns NfsStorageDetails
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetNfsStorageDetails() *NfsStorageDetails {
	return m.NfsStorageDetails
}

// GetCustomerContacts returns CustomerContacts
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetCustomerContacts() []CustomerContact {
	return m.CustomerContacts
}

// GetOkvEndPointGroupName returns OkvEndPointGroupName
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetOkvEndPointGroupName() *string {
	return m.OkvEndPointGroupName
}

// GetDisplayName returns DisplayName
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDbUniqueName returns DbUniqueName
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetDbUniqueName() *string {
	return m.DbUniqueName
}

// GetDbName returns DbName
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetDbName() *string {
	return m.DbName
}

// GetServiceLevelAgreementType returns ServiceLevelAgreementType
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetServiceLevelAgreementType() CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum {
	return m.ServiceLevelAgreementType
}

// GetAutonomousExadataInfrastructureId returns AutonomousExadataInfrastructureId
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetAutonomousExadataInfrastructureId() *string {
	return m.AutonomousExadataInfrastructureId
}

// GetDbVersion returns DbVersion
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetDbVersion() *string {
	return m.DbVersion
}

// GetDatabaseSoftwareImageId returns DatabaseSoftwareImageId
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetDatabaseSoftwareImageId() *string {
	return m.DatabaseSoftwareImageId
}

// GetPeerAutonomousExadataInfrastructureId returns PeerAutonomousExadataInfrastructureId
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetPeerAutonomousExadataInfrastructureId() *string {
	return m.PeerAutonomousExadataInfrastructureId
}

// GetPeerAutonomousContainerDatabaseDisplayName returns PeerAutonomousContainerDatabaseDisplayName
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetPeerAutonomousContainerDatabaseDisplayName() *string {
	return m.PeerAutonomousContainerDatabaseDisplayName
}

// GetProtectionMode returns ProtectionMode
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetProtectionMode() CreateAutonomousContainerDatabaseBaseProtectionModeEnum {
	return m.ProtectionMode
}

// GetFastStartFailOverLagLimitInSeconds returns FastStartFailOverLagLimitInSeconds
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetFastStartFailOverLagLimitInSeconds() *int {
	return m.FastStartFailOverLagLimitInSeconds
}

// GetIsAutomaticFailoverEnabled returns IsAutomaticFailoverEnabled
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetIsAutomaticFailoverEnabled() *bool {
	return m.IsAutomaticFailoverEnabled
}

// GetPeerCloudAutonomousVmClusterId returns PeerCloudAutonomousVmClusterId
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetPeerCloudAutonomousVmClusterId() *string {
	return m.PeerCloudAutonomousVmClusterId
}

// GetPeerAutonomousVmClusterId returns PeerAutonomousVmClusterId
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetPeerAutonomousVmClusterId() *string {
	return m.PeerAutonomousVmClusterId
}

// GetPeerAutonomousContainerDatabaseCompartmentId returns PeerAutonomousContainerDatabaseCompartmentId
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetPeerAutonomousContainerDatabaseCompartmentId() *string {
	return m.PeerAutonomousContainerDatabaseCompartmentId
}

// GetPeerAutonomousContainerDatabaseBackupConfig returns PeerAutonomousContainerDatabaseBackupConfig
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetPeerAutonomousContainerDatabaseBackupConfig() *PeerAutonomousContainerDatabaseBackupConfig {
	return m.PeerAutonomousContainerDatabaseBackupConfig
}

// GetPeerDbUniqueName returns PeerDbUniqueName
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetPeerDbUniqueName() *string {
	return m.PeerDbUniqueName
}

// GetAutonomousVmClusterId returns AutonomousVmClusterId
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetAutonomousVmClusterId() *string {
	return m.AutonomousVmClusterId
}

// GetCloudAutonomousVmClusterId returns CloudAutonomousVmClusterId
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetCloudAutonomousVmClusterId() *string {
	return m.CloudAutonomousVmClusterId
}

// GetCompartmentId returns CompartmentId
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetPatchModel returns PatchModel
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetPatchModel() CreateAutonomousContainerDatabaseBasePatchModelEnum {
	return m.PatchModel
}

// GetMaintenanceWindowDetails returns MaintenanceWindowDetails
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetMaintenanceWindowDetails() *MaintenanceWindow {
	return m.MaintenanceWindowDetails
}

// GetStandbyMaintenanceBufferInDays returns StandbyMaintenanceBufferInDays
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetStandbyMaintenanceBufferInDays() *int {
	return m.StandbyMaintenanceBufferInDays
}

// GetVersionPreference returns VersionPreference
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetVersionPreference() CreateAutonomousContainerDatabaseBaseVersionPreferenceEnum {
	return m.VersionPreference
}

// GetIsDstFileUpdateEnabled returns IsDstFileUpdateEnabled
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetIsDstFileUpdateEnabled() *bool {
	return m.IsDstFileUpdateEnabled
}

// GetFreeformTags returns FreeformTags
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetBackupConfig returns BackupConfig
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetBackupConfig() *AutonomousContainerDatabaseBackupConfig {
	return m.BackupConfig
}

// GetKmsKeyId returns KmsKeyId
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetKmsKeyId() *string {
	return m.KmsKeyId
}

// GetKmsKeyVersionId returns KmsKeyVersionId
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetKmsKeyVersionId() *string {
	return m.KmsKeyVersionId
}

// GetVaultId returns VaultId
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetVaultId() *string {
	return m.VaultId
}

// GetKeyStoreId returns KeyStoreId
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetKeyStoreId() *string {
	return m.KeyStoreId
}

// GetEncryptionKeyLocationDetails returns EncryptionKeyLocationDetails
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetEncryptionKeyLocationDetails() EncryptionKeyLocationDetails {
	return m.EncryptionKeyLocationDetails
}

// GetIsFipsEnabled returns IsFipsEnabled
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetIsFipsEnabled() *bool {
	return m.IsFipsEnabled
}

// GetDbSplitThreshold returns DbSplitThreshold
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetDbSplitThreshold() *int {
	return m.DbSplitThreshold
}

// GetVmFailoverReservation returns VmFailoverReservation
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetVmFailoverReservation() *int {
	return m.VmFailoverReservation
}

// GetDistributionAffinity returns DistributionAffinity
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetDistributionAffinity() CreateAutonomousContainerDatabaseBaseDistributionAffinityEnum {
	return m.DistributionAffinity
}

// GetNetServicesArchitecture returns NetServicesArchitecture
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) GetNetServicesArchitecture() CreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum {
	return m.NetServicesArchitecture
}

func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnum(string(m.CloneType)); !ok && m.CloneType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CloneType: %s. Supported values are: %s.", m.CloneType, strings.Join(GetCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnum(string(m.CloneBandWidth)); !ok && m.CloneBandWidth != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CloneBandWidth: %s. Supported values are: %s.", m.CloneBandWidth, strings.Join(GetCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnumStringValues(), ",")))
	}

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
func (m CreateAutonomousContainerDatabaseFromBackupTimestampDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateAutonomousContainerDatabaseFromBackupTimestampDetails CreateAutonomousContainerDatabaseFromBackupTimestampDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateAutonomousContainerDatabaseFromBackupTimestampDetails
	}{
		"BACKUP_FROM_TIMESTAMP",
		(MarshalTypeCreateAutonomousContainerDatabaseFromBackupTimestampDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateAutonomousContainerDatabaseFromBackupTimestampDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		NfsStorageDetails                            *NfsStorageDetails                                                            `json:"nfsStorageDetails"`
		CustomerContacts                             []CustomerContact                                                             `json:"customerContacts"`
		OkvEndPointGroupName                         *string                                                                       `json:"okvEndPointGroupName"`
		DbUniqueName                                 *string                                                                       `json:"dbUniqueName"`
		DbName                                       *string                                                                       `json:"dbName"`
		ServiceLevelAgreementType                    CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum            `json:"serviceLevelAgreementType"`
		AutonomousExadataInfrastructureId            *string                                                                       `json:"autonomousExadataInfrastructureId"`
		DbVersion                                    *string                                                                       `json:"dbVersion"`
		DatabaseSoftwareImageId                      *string                                                                       `json:"databaseSoftwareImageId"`
		PeerAutonomousExadataInfrastructureId        *string                                                                       `json:"peerAutonomousExadataInfrastructureId"`
		PeerAutonomousContainerDatabaseDisplayName   *string                                                                       `json:"peerAutonomousContainerDatabaseDisplayName"`
		ProtectionMode                               CreateAutonomousContainerDatabaseBaseProtectionModeEnum                       `json:"protectionMode"`
		FastStartFailOverLagLimitInSeconds           *int                                                                          `json:"fastStartFailOverLagLimitInSeconds"`
		IsAutomaticFailoverEnabled                   *bool                                                                         `json:"isAutomaticFailoverEnabled"`
		PeerCloudAutonomousVmClusterId               *string                                                                       `json:"peerCloudAutonomousVmClusterId"`
		PeerAutonomousVmClusterId                    *string                                                                       `json:"peerAutonomousVmClusterId"`
		PeerAutonomousContainerDatabaseCompartmentId *string                                                                       `json:"peerAutonomousContainerDatabaseCompartmentId"`
		PeerAutonomousContainerDatabaseBackupConfig  *PeerAutonomousContainerDatabaseBackupConfig                                  `json:"peerAutonomousContainerDatabaseBackupConfig"`
		PeerDbUniqueName                             *string                                                                       `json:"peerDbUniqueName"`
		AutonomousVmClusterId                        *string                                                                       `json:"autonomousVmClusterId"`
		CloudAutonomousVmClusterId                   *string                                                                       `json:"cloudAutonomousVmClusterId"`
		CompartmentId                                *string                                                                       `json:"compartmentId"`
		MaintenanceWindowDetails                     *MaintenanceWindow                                                            `json:"maintenanceWindowDetails"`
		StandbyMaintenanceBufferInDays               *int                                                                          `json:"standbyMaintenanceBufferInDays"`
		VersionPreference                            CreateAutonomousContainerDatabaseBaseVersionPreferenceEnum                    `json:"versionPreference"`
		IsDstFileUpdateEnabled                       *bool                                                                         `json:"isDstFileUpdateEnabled"`
		FreeformTags                                 map[string]string                                                             `json:"freeformTags"`
		DefinedTags                                  map[string]map[string]interface{}                                             `json:"definedTags"`
		BackupConfig                                 *AutonomousContainerDatabaseBackupConfig                                      `json:"backupConfig"`
		KmsKeyId                                     *string                                                                       `json:"kmsKeyId"`
		KmsKeyVersionId                              *string                                                                       `json:"kmsKeyVersionId"`
		VaultId                                      *string                                                                       `json:"vaultId"`
		KeyStoreId                                   *string                                                                       `json:"keyStoreId"`
		EncryptionKeyLocationDetails                 encryptionkeylocationdetails                                                  `json:"encryptionKeyLocationDetails"`
		IsFipsEnabled                                *bool                                                                         `json:"isFipsEnabled"`
		DbSplitThreshold                             *int                                                                          `json:"dbSplitThreshold"`
		VmFailoverReservation                        *int                                                                          `json:"vmFailoverReservation"`
		DistributionAffinity                         CreateAutonomousContainerDatabaseBaseDistributionAffinityEnum                 `json:"distributionAffinity"`
		NetServicesArchitecture                      CreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum              `json:"netServicesArchitecture"`
		TimeStampToUseForCloning                     *common.SDKTime                                                               `json:"timeStampToUseForCloning"`
		ShouldUseLatestAvailableBackupTimeStamp      *bool                                                                         `json:"shouldUseLatestAvailableBackupTimeStamp"`
		AutonomousDatabasesToClone                   []string                                                                      `json:"autonomousDatabasesToClone"`
		CloneBandWidth                               CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnum `json:"cloneBandWidth"`
		DisplayName                                  *string                                                                       `json:"displayName"`
		PatchModel                                   CreateAutonomousContainerDatabaseBasePatchModelEnum                           `json:"patchModel"`
		SourceAutonomousContainerDatabaseId          *string                                                                       `json:"sourceAutonomousContainerDatabaseId"`
		CloneType                                    CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnum      `json:"cloneType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.NfsStorageDetails = model.NfsStorageDetails

	m.CustomerContacts = make([]CustomerContact, len(model.CustomerContacts))
	copy(m.CustomerContacts, model.CustomerContacts)
	m.OkvEndPointGroupName = model.OkvEndPointGroupName

	m.DbUniqueName = model.DbUniqueName

	m.DbName = model.DbName

	m.ServiceLevelAgreementType = model.ServiceLevelAgreementType

	m.AutonomousExadataInfrastructureId = model.AutonomousExadataInfrastructureId

	m.DbVersion = model.DbVersion

	m.DatabaseSoftwareImageId = model.DatabaseSoftwareImageId

	m.PeerAutonomousExadataInfrastructureId = model.PeerAutonomousExadataInfrastructureId

	m.PeerAutonomousContainerDatabaseDisplayName = model.PeerAutonomousContainerDatabaseDisplayName

	m.ProtectionMode = model.ProtectionMode

	m.FastStartFailOverLagLimitInSeconds = model.FastStartFailOverLagLimitInSeconds

	m.IsAutomaticFailoverEnabled = model.IsAutomaticFailoverEnabled

	m.PeerCloudAutonomousVmClusterId = model.PeerCloudAutonomousVmClusterId

	m.PeerAutonomousVmClusterId = model.PeerAutonomousVmClusterId

	m.PeerAutonomousContainerDatabaseCompartmentId = model.PeerAutonomousContainerDatabaseCompartmentId

	m.PeerAutonomousContainerDatabaseBackupConfig = model.PeerAutonomousContainerDatabaseBackupConfig

	m.PeerDbUniqueName = model.PeerDbUniqueName

	m.AutonomousVmClusterId = model.AutonomousVmClusterId

	m.CloudAutonomousVmClusterId = model.CloudAutonomousVmClusterId

	m.CompartmentId = model.CompartmentId

	m.MaintenanceWindowDetails = model.MaintenanceWindowDetails

	m.StandbyMaintenanceBufferInDays = model.StandbyMaintenanceBufferInDays

	m.VersionPreference = model.VersionPreference

	m.IsDstFileUpdateEnabled = model.IsDstFileUpdateEnabled

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.BackupConfig = model.BackupConfig

	m.KmsKeyId = model.KmsKeyId

	m.KmsKeyVersionId = model.KmsKeyVersionId

	m.VaultId = model.VaultId

	m.KeyStoreId = model.KeyStoreId

	nn, e = model.EncryptionKeyLocationDetails.UnmarshalPolymorphicJSON(model.EncryptionKeyLocationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.EncryptionKeyLocationDetails = nn.(EncryptionKeyLocationDetails)
	} else {
		m.EncryptionKeyLocationDetails = nil
	}

	m.IsFipsEnabled = model.IsFipsEnabled

	m.DbSplitThreshold = model.DbSplitThreshold

	m.VmFailoverReservation = model.VmFailoverReservation

	m.DistributionAffinity = model.DistributionAffinity

	m.NetServicesArchitecture = model.NetServicesArchitecture

	m.TimeStampToUseForCloning = model.TimeStampToUseForCloning

	m.ShouldUseLatestAvailableBackupTimeStamp = model.ShouldUseLatestAvailableBackupTimeStamp

	m.AutonomousDatabasesToClone = make([]string, len(model.AutonomousDatabasesToClone))
	copy(m.AutonomousDatabasesToClone, model.AutonomousDatabasesToClone)
	m.CloneBandWidth = model.CloneBandWidth

	m.DisplayName = model.DisplayName

	m.PatchModel = model.PatchModel

	m.SourceAutonomousContainerDatabaseId = model.SourceAutonomousContainerDatabaseId

	m.CloneType = model.CloneType

	return
}

// CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnum Enum with underlying type: string
type CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnum string

// Set of constants representing the allowable values for CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnum
const (
	CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeFull     CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnum = "FULL"
	CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeMetadata CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnum = "METADATA"
	CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypePartial  CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnum = "PARTIAL"
)

var mappingCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnum = map[string]CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnum{
	"FULL":     CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeFull,
	"METADATA": CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeMetadata,
	"PARTIAL":  CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypePartial,
}

var mappingCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnumLowerCase = map[string]CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnum{
	"full":     CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeFull,
	"metadata": CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeMetadata,
	"partial":  CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypePartial,
}

// GetCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnumValues Enumerates the set of values for CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnum
func GetCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnumValues() []CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnum {
	values := make([]CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnum, 0)
	for _, v := range mappingCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnumStringValues Enumerates the set of values in String for CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnum
func GetCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnumStringValues() []string {
	return []string{
		"FULL",
		"METADATA",
		"PARTIAL",
	}
}

// GetMappingCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnum(val string) (CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnum, bool) {
	enum, ok := mappingCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnum Enum with underlying type: string
type CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnum string

// Set of constants representing the allowable values for CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnum
const (
	CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthSlow   CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnum = "SLOW"
	CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthMedium CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnum = "MEDIUM"
	CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthFast   CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnum = "FAST"
)

var mappingCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnum = map[string]CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnum{
	"SLOW":   CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthSlow,
	"MEDIUM": CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthMedium,
	"FAST":   CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthFast,
}

var mappingCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnumLowerCase = map[string]CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnum{
	"slow":   CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthSlow,
	"medium": CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthMedium,
	"fast":   CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthFast,
}

// GetCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnumValues Enumerates the set of values for CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnum
func GetCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnumValues() []CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnum {
	values := make([]CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnum, 0)
	for _, v := range mappingCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnumStringValues Enumerates the set of values in String for CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnum
func GetCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnumStringValues() []string {
	return []string{
		"SLOW",
		"MEDIUM",
		"FAST",
	}
}

// GetMappingCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnum(val string) (CreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnum, bool) {
	enum, ok := mappingCreateAutonomousContainerDatabaseFromBackupTimestampDetailsCloneBandWidthEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
