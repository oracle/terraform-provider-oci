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

// CreateAutonomousContainerDatabaseBase Describes the required parameters for the creation of an Autonomous Container Database.
type CreateAutonomousContainerDatabaseBase interface {

	// The display name for the Autonomous Container Database.
	GetDisplayName() *string

	// Database Patch model preference.
	GetPatchModel() CreateAutonomousContainerDatabaseBasePatchModelEnum

	// Customer Contacts. Setting this to an empty list removes all customer contacts.
	GetCustomerContacts() []CustomerContact

	// The OKV End Point Group name for the Autonomous Container Database.
	GetOkvEndPointGroupName() *string

	// **Deprecated.** The `DB_UNIQUE_NAME` value is set by Oracle Cloud Infrastructure.  Do not specify a value for this parameter. Specifying a value for this field will cause Terraform operations to fail.
	GetDbUniqueName() *string

	// The Database name for the Autonomous Container Database. The name must be unique within the Cloud Autonomous VM Cluster, starting with an alphabetic character, followed by 1 to 7 alphanumeric characters.
	GetDbName() *string

	// The service level agreement type of the Autonomous Container Database. The default is STANDARD. For an autonomous dataguard Autonomous Container Database, the specified Autonomous Exadata Infrastructure must be associated with a remote Autonomous Exadata Infrastructure.
	GetServiceLevelAgreementType() CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum

	// **No longer used.** This parameter is no longer used for Autonomous AI Database on dedicated Exadata infrasture. Specify a `cloudAutonomousVmClusterId` instead. Using this parameter will cause the operation to fail.
	GetAutonomousExadataInfrastructureId() *string

	// The base version for the Autonomous Container Database.
	GetDbVersion() *string

	// The Autonomous AI Database Software Image OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	GetDatabaseSoftwareImageId() *string

	// *No longer used.* This parameter is no longer used for Autonomous AI Database on dedicated Exadata infrasture. Specify a `peerCloudAutonomousVmClusterId` instead. Using this parameter will cause the operation to fail.
	GetPeerAutonomousExadataInfrastructureId() *string

	// The display name for the peer Autonomous Container Database.
	GetPeerAutonomousContainerDatabaseDisplayName() *string

	// The protection mode of this Autonomous Data Guard association. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	GetProtectionMode() CreateAutonomousContainerDatabaseBaseProtectionModeEnum

	// The lag time for my preference based on data loss tolerance in seconds.
	GetFastStartFailOverLagLimitInSeconds() *int

	// Indicates whether Automatic Failover is enabled for Autonomous Container Database Dataguard Association
	GetIsAutomaticFailoverEnabled() *bool

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer cloud Autonomous Exadata VM Cluster.
	GetPeerCloudAutonomousVmClusterId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer Autonomous VM cluster for Autonomous Data Guard. Required to enable Data Guard.
	GetPeerAutonomousVmClusterId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the standby Autonomous Container Database
	// will be created.
	GetPeerAutonomousContainerDatabaseCompartmentId() *string

	GetPeerAutonomousContainerDatabaseBackupConfig() *PeerAutonomousContainerDatabaseBackupConfig

	// **Deprecated.** The `DB_UNIQUE_NAME` of the peer Autonomous Container Database in a Data Guard association is set by Oracle Cloud Infrastructure.  Do not specify a value for this parameter. Specifying a value for this field will cause Terraform operations to fail.
	GetPeerDbUniqueName() *string

	// The OCID of the Autonomous VM Cluster.
	GetAutonomousVmClusterId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
	GetCloudAutonomousVmClusterId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Autonomous Container Database.
	GetCompartmentId() *string

	GetMaintenanceWindowDetails() *MaintenanceWindow

	// The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database.
	// This value represents the number of days before scheduled maintenance of the primary database.
	GetStandbyMaintenanceBufferInDays() *int

	// The next maintenance version preference.
	GetVersionPreference() CreateAutonomousContainerDatabaseBaseVersionPreferenceEnum

	// Indicates if an automatic DST Time Zone file update is enabled for the Autonomous Container Database. If enabled along with Release Update, patching will be done in a Non-Rolling manner.
	GetIsDstFileUpdateEnabled() *bool

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	GetDefinedTags() map[string]map[string]interface{}

	GetBackupConfig() *AutonomousContainerDatabaseBackupConfig

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	GetKmsKeyId() *string

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous AI Database Serverless does not use key versions, hence is not applicable for Autonomous AI Database Serverless instances.
	GetKmsKeyVersionId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
	GetVaultId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
	GetKeyStoreId() *string

	// The CPU value beyond which an Autonomous AI Database will be opened across multiple nodes. The default value of this attribute is 16 for OCPUs and 64 for ECPUs.
	GetDbSplitThreshold() *int

	// The percentage of CPUs reserved across nodes to support node failover. Allowed values are 0%, 25%, and 50%, with 50% being the default option.
	GetVmFailoverReservation() *int

	// Determines whether an Autonomous AI Database must be opened across a minimum or maximum of nodes. By default, Minimum nodes is selected.
	GetDistributionAffinity() CreateAutonomousContainerDatabaseBaseDistributionAffinityEnum

	// Enabling SHARED server architecture enables a database server to allow many client processes to share very few server processes, thereby increasing the number of supported users.
	GetNetServicesArchitecture() CreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum
}

type createautonomouscontainerdatabasebase struct {
	JsonData                                     []byte
	CustomerContacts                             []CustomerContact                                                  `mandatory:"false" json:"customerContacts"`
	OkvEndPointGroupName                         *string                                                            `mandatory:"false" json:"okvEndPointGroupName"`
	DbUniqueName                                 *string                                                            `mandatory:"false" json:"dbUniqueName"`
	DbName                                       *string                                                            `mandatory:"false" json:"dbName"`
	ServiceLevelAgreementType                    CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum `mandatory:"false" json:"serviceLevelAgreementType,omitempty"`
	AutonomousExadataInfrastructureId            *string                                                            `mandatory:"false" json:"autonomousExadataInfrastructureId"`
	DbVersion                                    *string                                                            `mandatory:"false" json:"dbVersion"`
	DatabaseSoftwareImageId                      *string                                                            `mandatory:"false" json:"databaseSoftwareImageId"`
	PeerAutonomousExadataInfrastructureId        *string                                                            `mandatory:"false" json:"peerAutonomousExadataInfrastructureId"`
	PeerAutonomousContainerDatabaseDisplayName   *string                                                            `mandatory:"false" json:"peerAutonomousContainerDatabaseDisplayName"`
	ProtectionMode                               CreateAutonomousContainerDatabaseBaseProtectionModeEnum            `mandatory:"false" json:"protectionMode,omitempty"`
	FastStartFailOverLagLimitInSeconds           *int                                                               `mandatory:"false" json:"fastStartFailOverLagLimitInSeconds"`
	IsAutomaticFailoverEnabled                   *bool                                                              `mandatory:"false" json:"isAutomaticFailoverEnabled"`
	PeerCloudAutonomousVmClusterId               *string                                                            `mandatory:"false" json:"peerCloudAutonomousVmClusterId"`
	PeerAutonomousVmClusterId                    *string                                                            `mandatory:"false" json:"peerAutonomousVmClusterId"`
	PeerAutonomousContainerDatabaseCompartmentId *string                                                            `mandatory:"false" json:"peerAutonomousContainerDatabaseCompartmentId"`
	PeerAutonomousContainerDatabaseBackupConfig  *PeerAutonomousContainerDatabaseBackupConfig                       `mandatory:"false" json:"peerAutonomousContainerDatabaseBackupConfig"`
	PeerDbUniqueName                             *string                                                            `mandatory:"false" json:"peerDbUniqueName"`
	AutonomousVmClusterId                        *string                                                            `mandatory:"false" json:"autonomousVmClusterId"`
	CloudAutonomousVmClusterId                   *string                                                            `mandatory:"false" json:"cloudAutonomousVmClusterId"`
	CompartmentId                                *string                                                            `mandatory:"false" json:"compartmentId"`
	MaintenanceWindowDetails                     *MaintenanceWindow                                                 `mandatory:"false" json:"maintenanceWindowDetails"`
	StandbyMaintenanceBufferInDays               *int                                                               `mandatory:"false" json:"standbyMaintenanceBufferInDays"`
	VersionPreference                            CreateAutonomousContainerDatabaseBaseVersionPreferenceEnum         `mandatory:"false" json:"versionPreference,omitempty"`
	IsDstFileUpdateEnabled                       *bool                                                              `mandatory:"false" json:"isDstFileUpdateEnabled"`
	FreeformTags                                 map[string]string                                                  `mandatory:"false" json:"freeformTags"`
	DefinedTags                                  map[string]map[string]interface{}                                  `mandatory:"false" json:"definedTags"`
	BackupConfig                                 *AutonomousContainerDatabaseBackupConfig                           `mandatory:"false" json:"backupConfig"`
	KmsKeyId                                     *string                                                            `mandatory:"false" json:"kmsKeyId"`
	KmsKeyVersionId                              *string                                                            `mandatory:"false" json:"kmsKeyVersionId"`
	VaultId                                      *string                                                            `mandatory:"false" json:"vaultId"`
	KeyStoreId                                   *string                                                            `mandatory:"false" json:"keyStoreId"`
	DbSplitThreshold                             *int                                                               `mandatory:"false" json:"dbSplitThreshold"`
	VmFailoverReservation                        *int                                                               `mandatory:"false" json:"vmFailoverReservation"`
	DistributionAffinity                         CreateAutonomousContainerDatabaseBaseDistributionAffinityEnum      `mandatory:"false" json:"distributionAffinity,omitempty"`
	NetServicesArchitecture                      CreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum   `mandatory:"false" json:"netServicesArchitecture,omitempty"`
	DisplayName                                  *string                                                            `mandatory:"true" json:"displayName"`
	PatchModel                                   CreateAutonomousContainerDatabaseBasePatchModelEnum                `mandatory:"true" json:"patchModel"`
	Source                                       string                                                             `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *createautonomouscontainerdatabasebase) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateautonomouscontainerdatabasebase createautonomouscontainerdatabasebase
	s := struct {
		Model Unmarshalercreateautonomouscontainerdatabasebase
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.PatchModel = s.Model.PatchModel
	m.CustomerContacts = s.Model.CustomerContacts
	m.OkvEndPointGroupName = s.Model.OkvEndPointGroupName
	m.DbUniqueName = s.Model.DbUniqueName
	m.DbName = s.Model.DbName
	m.ServiceLevelAgreementType = s.Model.ServiceLevelAgreementType
	m.AutonomousExadataInfrastructureId = s.Model.AutonomousExadataInfrastructureId
	m.DbVersion = s.Model.DbVersion
	m.DatabaseSoftwareImageId = s.Model.DatabaseSoftwareImageId
	m.PeerAutonomousExadataInfrastructureId = s.Model.PeerAutonomousExadataInfrastructureId
	m.PeerAutonomousContainerDatabaseDisplayName = s.Model.PeerAutonomousContainerDatabaseDisplayName
	m.ProtectionMode = s.Model.ProtectionMode
	m.FastStartFailOverLagLimitInSeconds = s.Model.FastStartFailOverLagLimitInSeconds
	m.IsAutomaticFailoverEnabled = s.Model.IsAutomaticFailoverEnabled
	m.PeerCloudAutonomousVmClusterId = s.Model.PeerCloudAutonomousVmClusterId
	m.PeerAutonomousVmClusterId = s.Model.PeerAutonomousVmClusterId
	m.PeerAutonomousContainerDatabaseCompartmentId = s.Model.PeerAutonomousContainerDatabaseCompartmentId
	m.PeerAutonomousContainerDatabaseBackupConfig = s.Model.PeerAutonomousContainerDatabaseBackupConfig
	m.PeerDbUniqueName = s.Model.PeerDbUniqueName
	m.AutonomousVmClusterId = s.Model.AutonomousVmClusterId
	m.CloudAutonomousVmClusterId = s.Model.CloudAutonomousVmClusterId
	m.CompartmentId = s.Model.CompartmentId
	m.MaintenanceWindowDetails = s.Model.MaintenanceWindowDetails
	m.StandbyMaintenanceBufferInDays = s.Model.StandbyMaintenanceBufferInDays
	m.VersionPreference = s.Model.VersionPreference
	m.IsDstFileUpdateEnabled = s.Model.IsDstFileUpdateEnabled
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.BackupConfig = s.Model.BackupConfig
	m.KmsKeyId = s.Model.KmsKeyId
	m.KmsKeyVersionId = s.Model.KmsKeyVersionId
	m.VaultId = s.Model.VaultId
	m.KeyStoreId = s.Model.KeyStoreId
	m.DbSplitThreshold = s.Model.DbSplitThreshold
	m.VmFailoverReservation = s.Model.VmFailoverReservation
	m.DistributionAffinity = s.Model.DistributionAffinity
	m.NetServicesArchitecture = s.Model.NetServicesArchitecture
	m.Source = s.Model.Source

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createautonomouscontainerdatabasebase) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "NONE":
		mm := CreateAutonomousContainerDatabaseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BACKUP_FROM_ID":
		mm := CreateAutonomousContainerDatabaseFromBackupDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateAutonomousContainerDatabaseBase: %s.", m.Source)
		return *m, nil
	}
}

// GetCustomerContacts returns CustomerContacts
func (m createautonomouscontainerdatabasebase) GetCustomerContacts() []CustomerContact {
	return m.CustomerContacts
}

// GetOkvEndPointGroupName returns OkvEndPointGroupName
func (m createautonomouscontainerdatabasebase) GetOkvEndPointGroupName() *string {
	return m.OkvEndPointGroupName
}

// GetDbUniqueName returns DbUniqueName
func (m createautonomouscontainerdatabasebase) GetDbUniqueName() *string {
	return m.DbUniqueName
}

// GetDbName returns DbName
func (m createautonomouscontainerdatabasebase) GetDbName() *string {
	return m.DbName
}

// GetServiceLevelAgreementType returns ServiceLevelAgreementType
func (m createautonomouscontainerdatabasebase) GetServiceLevelAgreementType() CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum {
	return m.ServiceLevelAgreementType
}

// GetAutonomousExadataInfrastructureId returns AutonomousExadataInfrastructureId
func (m createautonomouscontainerdatabasebase) GetAutonomousExadataInfrastructureId() *string {
	return m.AutonomousExadataInfrastructureId
}

// GetDbVersion returns DbVersion
func (m createautonomouscontainerdatabasebase) GetDbVersion() *string {
	return m.DbVersion
}

// GetDatabaseSoftwareImageId returns DatabaseSoftwareImageId
func (m createautonomouscontainerdatabasebase) GetDatabaseSoftwareImageId() *string {
	return m.DatabaseSoftwareImageId
}

// GetPeerAutonomousExadataInfrastructureId returns PeerAutonomousExadataInfrastructureId
func (m createautonomouscontainerdatabasebase) GetPeerAutonomousExadataInfrastructureId() *string {
	return m.PeerAutonomousExadataInfrastructureId
}

// GetPeerAutonomousContainerDatabaseDisplayName returns PeerAutonomousContainerDatabaseDisplayName
func (m createautonomouscontainerdatabasebase) GetPeerAutonomousContainerDatabaseDisplayName() *string {
	return m.PeerAutonomousContainerDatabaseDisplayName
}

// GetProtectionMode returns ProtectionMode
func (m createautonomouscontainerdatabasebase) GetProtectionMode() CreateAutonomousContainerDatabaseBaseProtectionModeEnum {
	return m.ProtectionMode
}

// GetFastStartFailOverLagLimitInSeconds returns FastStartFailOverLagLimitInSeconds
func (m createautonomouscontainerdatabasebase) GetFastStartFailOverLagLimitInSeconds() *int {
	return m.FastStartFailOverLagLimitInSeconds
}

// GetIsAutomaticFailoverEnabled returns IsAutomaticFailoverEnabled
func (m createautonomouscontainerdatabasebase) GetIsAutomaticFailoverEnabled() *bool {
	return m.IsAutomaticFailoverEnabled
}

// GetPeerCloudAutonomousVmClusterId returns PeerCloudAutonomousVmClusterId
func (m createautonomouscontainerdatabasebase) GetPeerCloudAutonomousVmClusterId() *string {
	return m.PeerCloudAutonomousVmClusterId
}

// GetPeerAutonomousVmClusterId returns PeerAutonomousVmClusterId
func (m createautonomouscontainerdatabasebase) GetPeerAutonomousVmClusterId() *string {
	return m.PeerAutonomousVmClusterId
}

// GetPeerAutonomousContainerDatabaseCompartmentId returns PeerAutonomousContainerDatabaseCompartmentId
func (m createautonomouscontainerdatabasebase) GetPeerAutonomousContainerDatabaseCompartmentId() *string {
	return m.PeerAutonomousContainerDatabaseCompartmentId
}

// GetPeerAutonomousContainerDatabaseBackupConfig returns PeerAutonomousContainerDatabaseBackupConfig
func (m createautonomouscontainerdatabasebase) GetPeerAutonomousContainerDatabaseBackupConfig() *PeerAutonomousContainerDatabaseBackupConfig {
	return m.PeerAutonomousContainerDatabaseBackupConfig
}

// GetPeerDbUniqueName returns PeerDbUniqueName
func (m createautonomouscontainerdatabasebase) GetPeerDbUniqueName() *string {
	return m.PeerDbUniqueName
}

// GetAutonomousVmClusterId returns AutonomousVmClusterId
func (m createautonomouscontainerdatabasebase) GetAutonomousVmClusterId() *string {
	return m.AutonomousVmClusterId
}

// GetCloudAutonomousVmClusterId returns CloudAutonomousVmClusterId
func (m createautonomouscontainerdatabasebase) GetCloudAutonomousVmClusterId() *string {
	return m.CloudAutonomousVmClusterId
}

// GetCompartmentId returns CompartmentId
func (m createautonomouscontainerdatabasebase) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetMaintenanceWindowDetails returns MaintenanceWindowDetails
func (m createautonomouscontainerdatabasebase) GetMaintenanceWindowDetails() *MaintenanceWindow {
	return m.MaintenanceWindowDetails
}

// GetStandbyMaintenanceBufferInDays returns StandbyMaintenanceBufferInDays
func (m createautonomouscontainerdatabasebase) GetStandbyMaintenanceBufferInDays() *int {
	return m.StandbyMaintenanceBufferInDays
}

// GetVersionPreference returns VersionPreference
func (m createautonomouscontainerdatabasebase) GetVersionPreference() CreateAutonomousContainerDatabaseBaseVersionPreferenceEnum {
	return m.VersionPreference
}

// GetIsDstFileUpdateEnabled returns IsDstFileUpdateEnabled
func (m createautonomouscontainerdatabasebase) GetIsDstFileUpdateEnabled() *bool {
	return m.IsDstFileUpdateEnabled
}

// GetFreeformTags returns FreeformTags
func (m createautonomouscontainerdatabasebase) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createautonomouscontainerdatabasebase) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetBackupConfig returns BackupConfig
func (m createautonomouscontainerdatabasebase) GetBackupConfig() *AutonomousContainerDatabaseBackupConfig {
	return m.BackupConfig
}

// GetKmsKeyId returns KmsKeyId
func (m createautonomouscontainerdatabasebase) GetKmsKeyId() *string {
	return m.KmsKeyId
}

// GetKmsKeyVersionId returns KmsKeyVersionId
func (m createautonomouscontainerdatabasebase) GetKmsKeyVersionId() *string {
	return m.KmsKeyVersionId
}

// GetVaultId returns VaultId
func (m createautonomouscontainerdatabasebase) GetVaultId() *string {
	return m.VaultId
}

// GetKeyStoreId returns KeyStoreId
func (m createautonomouscontainerdatabasebase) GetKeyStoreId() *string {
	return m.KeyStoreId
}

// GetDbSplitThreshold returns DbSplitThreshold
func (m createautonomouscontainerdatabasebase) GetDbSplitThreshold() *int {
	return m.DbSplitThreshold
}

// GetVmFailoverReservation returns VmFailoverReservation
func (m createautonomouscontainerdatabasebase) GetVmFailoverReservation() *int {
	return m.VmFailoverReservation
}

// GetDistributionAffinity returns DistributionAffinity
func (m createautonomouscontainerdatabasebase) GetDistributionAffinity() CreateAutonomousContainerDatabaseBaseDistributionAffinityEnum {
	return m.DistributionAffinity
}

// GetNetServicesArchitecture returns NetServicesArchitecture
func (m createautonomouscontainerdatabasebase) GetNetServicesArchitecture() CreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum {
	return m.NetServicesArchitecture
}

// GetDisplayName returns DisplayName
func (m createautonomouscontainerdatabasebase) GetDisplayName() *string {
	return m.DisplayName
}

// GetPatchModel returns PatchModel
func (m createautonomouscontainerdatabasebase) GetPatchModel() CreateAutonomousContainerDatabaseBasePatchModelEnum {
	return m.PatchModel
}

func (m createautonomouscontainerdatabasebase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createautonomouscontainerdatabasebase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateAutonomousContainerDatabaseBasePatchModelEnum(string(m.PatchModel)); !ok && m.PatchModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchModel: %s. Supported values are: %s.", m.PatchModel, strings.Join(GetCreateAutonomousContainerDatabaseBasePatchModelEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum(string(m.ServiceLevelAgreementType)); !ok && m.ServiceLevelAgreementType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceLevelAgreementType: %s. Supported values are: %s.", m.ServiceLevelAgreementType, strings.Join(GetCreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateAutonomousContainerDatabaseBaseProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetCreateAutonomousContainerDatabaseBaseProtectionModeEnumStringValues(), ",")))
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

// CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum Enum with underlying type: string
type CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum string

// Set of constants representing the allowable values for CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum
const (
	CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeStandard            CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum = "STANDARD"
	CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeAutonomousDataguard CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum = "AUTONOMOUS_DATAGUARD"
)

var mappingCreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum = map[string]CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum{
	"STANDARD":             CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeStandard,
	"AUTONOMOUS_DATAGUARD": CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeAutonomousDataguard,
}

var mappingCreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnumLowerCase = map[string]CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum{
	"standard":             CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeStandard,
	"autonomous_dataguard": CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeAutonomousDataguard,
}

// GetCreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnumValues Enumerates the set of values for CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum
func GetCreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnumValues() []CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum {
	values := make([]CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum, 0)
	for _, v := range mappingCreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnumStringValues Enumerates the set of values in String for CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum
func GetCreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnumStringValues() []string {
	return []string{
		"STANDARD",
		"AUTONOMOUS_DATAGUARD",
	}
}

// GetMappingCreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum(val string) (CreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnum, bool) {
	enum, ok := mappingCreateAutonomousContainerDatabaseBaseServiceLevelAgreementTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateAutonomousContainerDatabaseBaseProtectionModeEnum Enum with underlying type: string
type CreateAutonomousContainerDatabaseBaseProtectionModeEnum string

// Set of constants representing the allowable values for CreateAutonomousContainerDatabaseBaseProtectionModeEnum
const (
	CreateAutonomousContainerDatabaseBaseProtectionModeAvailability CreateAutonomousContainerDatabaseBaseProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	CreateAutonomousContainerDatabaseBaseProtectionModePerformance  CreateAutonomousContainerDatabaseBaseProtectionModeEnum = "MAXIMUM_PERFORMANCE"
)

var mappingCreateAutonomousContainerDatabaseBaseProtectionModeEnum = map[string]CreateAutonomousContainerDatabaseBaseProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": CreateAutonomousContainerDatabaseBaseProtectionModeAvailability,
	"MAXIMUM_PERFORMANCE":  CreateAutonomousContainerDatabaseBaseProtectionModePerformance,
}

var mappingCreateAutonomousContainerDatabaseBaseProtectionModeEnumLowerCase = map[string]CreateAutonomousContainerDatabaseBaseProtectionModeEnum{
	"maximum_availability": CreateAutonomousContainerDatabaseBaseProtectionModeAvailability,
	"maximum_performance":  CreateAutonomousContainerDatabaseBaseProtectionModePerformance,
}

// GetCreateAutonomousContainerDatabaseBaseProtectionModeEnumValues Enumerates the set of values for CreateAutonomousContainerDatabaseBaseProtectionModeEnum
func GetCreateAutonomousContainerDatabaseBaseProtectionModeEnumValues() []CreateAutonomousContainerDatabaseBaseProtectionModeEnum {
	values := make([]CreateAutonomousContainerDatabaseBaseProtectionModeEnum, 0)
	for _, v := range mappingCreateAutonomousContainerDatabaseBaseProtectionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousContainerDatabaseBaseProtectionModeEnumStringValues Enumerates the set of values in String for CreateAutonomousContainerDatabaseBaseProtectionModeEnum
func GetCreateAutonomousContainerDatabaseBaseProtectionModeEnumStringValues() []string {
	return []string{
		"MAXIMUM_AVAILABILITY",
		"MAXIMUM_PERFORMANCE",
	}
}

// GetMappingCreateAutonomousContainerDatabaseBaseProtectionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousContainerDatabaseBaseProtectionModeEnum(val string) (CreateAutonomousContainerDatabaseBaseProtectionModeEnum, bool) {
	enum, ok := mappingCreateAutonomousContainerDatabaseBaseProtectionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateAutonomousContainerDatabaseBasePatchModelEnum Enum with underlying type: string
type CreateAutonomousContainerDatabaseBasePatchModelEnum string

// Set of constants representing the allowable values for CreateAutonomousContainerDatabaseBasePatchModelEnum
const (
	CreateAutonomousContainerDatabaseBasePatchModelUpdates         CreateAutonomousContainerDatabaseBasePatchModelEnum = "RELEASE_UPDATES"
	CreateAutonomousContainerDatabaseBasePatchModelUpdateRevisions CreateAutonomousContainerDatabaseBasePatchModelEnum = "RELEASE_UPDATE_REVISIONS"
)

var mappingCreateAutonomousContainerDatabaseBasePatchModelEnum = map[string]CreateAutonomousContainerDatabaseBasePatchModelEnum{
	"RELEASE_UPDATES":          CreateAutonomousContainerDatabaseBasePatchModelUpdates,
	"RELEASE_UPDATE_REVISIONS": CreateAutonomousContainerDatabaseBasePatchModelUpdateRevisions,
}

var mappingCreateAutonomousContainerDatabaseBasePatchModelEnumLowerCase = map[string]CreateAutonomousContainerDatabaseBasePatchModelEnum{
	"release_updates":          CreateAutonomousContainerDatabaseBasePatchModelUpdates,
	"release_update_revisions": CreateAutonomousContainerDatabaseBasePatchModelUpdateRevisions,
}

// GetCreateAutonomousContainerDatabaseBasePatchModelEnumValues Enumerates the set of values for CreateAutonomousContainerDatabaseBasePatchModelEnum
func GetCreateAutonomousContainerDatabaseBasePatchModelEnumValues() []CreateAutonomousContainerDatabaseBasePatchModelEnum {
	values := make([]CreateAutonomousContainerDatabaseBasePatchModelEnum, 0)
	for _, v := range mappingCreateAutonomousContainerDatabaseBasePatchModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousContainerDatabaseBasePatchModelEnumStringValues Enumerates the set of values in String for CreateAutonomousContainerDatabaseBasePatchModelEnum
func GetCreateAutonomousContainerDatabaseBasePatchModelEnumStringValues() []string {
	return []string{
		"RELEASE_UPDATES",
		"RELEASE_UPDATE_REVISIONS",
	}
}

// GetMappingCreateAutonomousContainerDatabaseBasePatchModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousContainerDatabaseBasePatchModelEnum(val string) (CreateAutonomousContainerDatabaseBasePatchModelEnum, bool) {
	enum, ok := mappingCreateAutonomousContainerDatabaseBasePatchModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateAutonomousContainerDatabaseBaseVersionPreferenceEnum Enum with underlying type: string
type CreateAutonomousContainerDatabaseBaseVersionPreferenceEnum string

// Set of constants representing the allowable values for CreateAutonomousContainerDatabaseBaseVersionPreferenceEnum
const (
	CreateAutonomousContainerDatabaseBaseVersionPreferenceNextReleaseUpdate   CreateAutonomousContainerDatabaseBaseVersionPreferenceEnum = "NEXT_RELEASE_UPDATE"
	CreateAutonomousContainerDatabaseBaseVersionPreferenceLatestReleaseUpdate CreateAutonomousContainerDatabaseBaseVersionPreferenceEnum = "LATEST_RELEASE_UPDATE"
)

var mappingCreateAutonomousContainerDatabaseBaseVersionPreferenceEnum = map[string]CreateAutonomousContainerDatabaseBaseVersionPreferenceEnum{
	"NEXT_RELEASE_UPDATE":   CreateAutonomousContainerDatabaseBaseVersionPreferenceNextReleaseUpdate,
	"LATEST_RELEASE_UPDATE": CreateAutonomousContainerDatabaseBaseVersionPreferenceLatestReleaseUpdate,
}

var mappingCreateAutonomousContainerDatabaseBaseVersionPreferenceEnumLowerCase = map[string]CreateAutonomousContainerDatabaseBaseVersionPreferenceEnum{
	"next_release_update":   CreateAutonomousContainerDatabaseBaseVersionPreferenceNextReleaseUpdate,
	"latest_release_update": CreateAutonomousContainerDatabaseBaseVersionPreferenceLatestReleaseUpdate,
}

// GetCreateAutonomousContainerDatabaseBaseVersionPreferenceEnumValues Enumerates the set of values for CreateAutonomousContainerDatabaseBaseVersionPreferenceEnum
func GetCreateAutonomousContainerDatabaseBaseVersionPreferenceEnumValues() []CreateAutonomousContainerDatabaseBaseVersionPreferenceEnum {
	values := make([]CreateAutonomousContainerDatabaseBaseVersionPreferenceEnum, 0)
	for _, v := range mappingCreateAutonomousContainerDatabaseBaseVersionPreferenceEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousContainerDatabaseBaseVersionPreferenceEnumStringValues Enumerates the set of values in String for CreateAutonomousContainerDatabaseBaseVersionPreferenceEnum
func GetCreateAutonomousContainerDatabaseBaseVersionPreferenceEnumStringValues() []string {
	return []string{
		"NEXT_RELEASE_UPDATE",
		"LATEST_RELEASE_UPDATE",
	}
}

// GetMappingCreateAutonomousContainerDatabaseBaseVersionPreferenceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousContainerDatabaseBaseVersionPreferenceEnum(val string) (CreateAutonomousContainerDatabaseBaseVersionPreferenceEnum, bool) {
	enum, ok := mappingCreateAutonomousContainerDatabaseBaseVersionPreferenceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateAutonomousContainerDatabaseBaseDistributionAffinityEnum Enum with underlying type: string
type CreateAutonomousContainerDatabaseBaseDistributionAffinityEnum string

// Set of constants representing the allowable values for CreateAutonomousContainerDatabaseBaseDistributionAffinityEnum
const (
	CreateAutonomousContainerDatabaseBaseDistributionAffinityMinimumDistribution CreateAutonomousContainerDatabaseBaseDistributionAffinityEnum = "MINIMUM_DISTRIBUTION"
	CreateAutonomousContainerDatabaseBaseDistributionAffinityMaximumDistribution CreateAutonomousContainerDatabaseBaseDistributionAffinityEnum = "MAXIMUM_DISTRIBUTION"
)

var mappingCreateAutonomousContainerDatabaseBaseDistributionAffinityEnum = map[string]CreateAutonomousContainerDatabaseBaseDistributionAffinityEnum{
	"MINIMUM_DISTRIBUTION": CreateAutonomousContainerDatabaseBaseDistributionAffinityMinimumDistribution,
	"MAXIMUM_DISTRIBUTION": CreateAutonomousContainerDatabaseBaseDistributionAffinityMaximumDistribution,
}

var mappingCreateAutonomousContainerDatabaseBaseDistributionAffinityEnumLowerCase = map[string]CreateAutonomousContainerDatabaseBaseDistributionAffinityEnum{
	"minimum_distribution": CreateAutonomousContainerDatabaseBaseDistributionAffinityMinimumDistribution,
	"maximum_distribution": CreateAutonomousContainerDatabaseBaseDistributionAffinityMaximumDistribution,
}

// GetCreateAutonomousContainerDatabaseBaseDistributionAffinityEnumValues Enumerates the set of values for CreateAutonomousContainerDatabaseBaseDistributionAffinityEnum
func GetCreateAutonomousContainerDatabaseBaseDistributionAffinityEnumValues() []CreateAutonomousContainerDatabaseBaseDistributionAffinityEnum {
	values := make([]CreateAutonomousContainerDatabaseBaseDistributionAffinityEnum, 0)
	for _, v := range mappingCreateAutonomousContainerDatabaseBaseDistributionAffinityEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousContainerDatabaseBaseDistributionAffinityEnumStringValues Enumerates the set of values in String for CreateAutonomousContainerDatabaseBaseDistributionAffinityEnum
func GetCreateAutonomousContainerDatabaseBaseDistributionAffinityEnumStringValues() []string {
	return []string{
		"MINIMUM_DISTRIBUTION",
		"MAXIMUM_DISTRIBUTION",
	}
}

// GetMappingCreateAutonomousContainerDatabaseBaseDistributionAffinityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousContainerDatabaseBaseDistributionAffinityEnum(val string) (CreateAutonomousContainerDatabaseBaseDistributionAffinityEnum, bool) {
	enum, ok := mappingCreateAutonomousContainerDatabaseBaseDistributionAffinityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum Enum with underlying type: string
type CreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum string

// Set of constants representing the allowable values for CreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum
const (
	CreateAutonomousContainerDatabaseBaseNetServicesArchitectureDedicated CreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum = "DEDICATED"
	CreateAutonomousContainerDatabaseBaseNetServicesArchitectureShared    CreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum = "SHARED"
)

var mappingCreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum = map[string]CreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum{
	"DEDICATED": CreateAutonomousContainerDatabaseBaseNetServicesArchitectureDedicated,
	"SHARED":    CreateAutonomousContainerDatabaseBaseNetServicesArchitectureShared,
}

var mappingCreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnumLowerCase = map[string]CreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum{
	"dedicated": CreateAutonomousContainerDatabaseBaseNetServicesArchitectureDedicated,
	"shared":    CreateAutonomousContainerDatabaseBaseNetServicesArchitectureShared,
}

// GetCreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnumValues Enumerates the set of values for CreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum
func GetCreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnumValues() []CreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum {
	values := make([]CreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum, 0)
	for _, v := range mappingCreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnumStringValues Enumerates the set of values in String for CreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum
func GetCreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnumStringValues() []string {
	return []string{
		"DEDICATED",
		"SHARED",
	}
}

// GetMappingCreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum(val string) (CreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnum, bool) {
	enum, ok := mappingCreateAutonomousContainerDatabaseBaseNetServicesArchitectureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateAutonomousContainerDatabaseBaseSourceEnum Enum with underlying type: string
type CreateAutonomousContainerDatabaseBaseSourceEnum string

// Set of constants representing the allowable values for CreateAutonomousContainerDatabaseBaseSourceEnum
const (
	CreateAutonomousContainerDatabaseBaseSourceNone         CreateAutonomousContainerDatabaseBaseSourceEnum = "NONE"
	CreateAutonomousContainerDatabaseBaseSourceBackupFromId CreateAutonomousContainerDatabaseBaseSourceEnum = "BACKUP_FROM_ID"
)

var mappingCreateAutonomousContainerDatabaseBaseSourceEnum = map[string]CreateAutonomousContainerDatabaseBaseSourceEnum{
	"NONE":           CreateAutonomousContainerDatabaseBaseSourceNone,
	"BACKUP_FROM_ID": CreateAutonomousContainerDatabaseBaseSourceBackupFromId,
}

var mappingCreateAutonomousContainerDatabaseBaseSourceEnumLowerCase = map[string]CreateAutonomousContainerDatabaseBaseSourceEnum{
	"none":           CreateAutonomousContainerDatabaseBaseSourceNone,
	"backup_from_id": CreateAutonomousContainerDatabaseBaseSourceBackupFromId,
}

// GetCreateAutonomousContainerDatabaseBaseSourceEnumValues Enumerates the set of values for CreateAutonomousContainerDatabaseBaseSourceEnum
func GetCreateAutonomousContainerDatabaseBaseSourceEnumValues() []CreateAutonomousContainerDatabaseBaseSourceEnum {
	values := make([]CreateAutonomousContainerDatabaseBaseSourceEnum, 0)
	for _, v := range mappingCreateAutonomousContainerDatabaseBaseSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousContainerDatabaseBaseSourceEnumStringValues Enumerates the set of values in String for CreateAutonomousContainerDatabaseBaseSourceEnum
func GetCreateAutonomousContainerDatabaseBaseSourceEnumStringValues() []string {
	return []string{
		"NONE",
		"BACKUP_FROM_ID",
	}
}

// GetMappingCreateAutonomousContainerDatabaseBaseSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousContainerDatabaseBaseSourceEnum(val string) (CreateAutonomousContainerDatabaseBaseSourceEnum, bool) {
	enum, ok := mappingCreateAutonomousContainerDatabaseBaseSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
