// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// CreateAutonomousContainerDatabaseDetails Describes the required parameters for the creation of an Autonomous Container Database.
type CreateAutonomousContainerDatabaseDetails struct {

	// The display name for the Autonomous Container Database.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Database Patch model preference.
	PatchModel CreateAutonomousContainerDatabaseDetailsPatchModelEnum `mandatory:"true" json:"patchModel"`

	// The `DB_UNIQUE_NAME` of the Oracle Database being backed up.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// The service level agreement type of the Autonomous Container Database. The default is STANDARD. For an autonomous dataguard Autonomous Container Database, the specified Autonomous Exadata Infrastructure must be associated with a remote Autonomous Exadata Infrastructure.
	ServiceLevelAgreementType CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum `mandatory:"false" json:"serviceLevelAgreementType,omitempty"`

	// The OCID of the Autonomous Exadata Infrastructure.
	AutonomousExadataInfrastructureId *string `mandatory:"false" json:"autonomousExadataInfrastructureId"`

	// The OCID of the peer Autonomous Exadata Infrastructure for Autonomous Data Guard.
	PeerAutonomousExadataInfrastructureId *string `mandatory:"false" json:"peerAutonomousExadataInfrastructureId"`

	// The display name for the peer Autonomous Container Database.
	PeerAutonomousContainerDatabaseDisplayName *string `mandatory:"false" json:"peerAutonomousContainerDatabaseDisplayName"`

	// The protection mode of this Autonomous Data Guard association. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode CreateAutonomousContainerDatabaseDetailsProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the peer Autonomous VM cluster for Autonomous Data Guard. Required to enable Data Guard.
	PeerAutonomousVmClusterId *string `mandatory:"false" json:"peerAutonomousVmClusterId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment where the standby Autonomous Container Database
	// will be created.
	PeerAutonomousContainerDatabaseCompartmentId *string `mandatory:"false" json:"peerAutonomousContainerDatabaseCompartmentId"`

	PeerAutonomousContainerDatabaseBackupConfig *PeerAutonomousContainerDatabaseBackupConfig `mandatory:"false" json:"peerAutonomousContainerDatabaseBackupConfig"`

	// The `DB_UNIQUE_NAME` of the peer Autonomous Container Database in a Data Guard association.
	PeerDbUniqueName *string `mandatory:"false" json:"peerDbUniqueName"`

	// The OCID of the Autonomous VM Cluster.
	AutonomousVmClusterId *string `mandatory:"false" json:"autonomousVmClusterId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment containing the Autonomous Container Database.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	MaintenanceWindowDetails *MaintenanceWindow `mandatory:"false" json:"maintenanceWindowDetails"`

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

	BackupConfig *AutonomousContainerDatabaseBackupConfig `mandatory:"false" json:"backupConfig"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.cloud.oracle.com/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the key store.
	KeyStoreId *string `mandatory:"false" json:"keyStoreId"`
}

func (m CreateAutonomousContainerDatabaseDetails) String() string {
	return common.PointerString(m)
}

// CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum Enum with underlying type: string
type CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum string

// Set of constants representing the allowable values for CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum
const (
	CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeStandard            CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum = "STANDARD"
	CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeAutonomousDataguard CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum = "AUTONOMOUS_DATAGUARD"
)

var mappingCreateAutonomousContainerDatabaseDetailsServiceLevelAgreementType = map[string]CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum{
	"STANDARD":             CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeStandard,
	"AUTONOMOUS_DATAGUARD": CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeAutonomousDataguard,
}

// GetCreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnumValues Enumerates the set of values for CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum
func GetCreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnumValues() []CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum {
	values := make([]CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum, 0)
	for _, v := range mappingCreateAutonomousContainerDatabaseDetailsServiceLevelAgreementType {
		values = append(values, v)
	}
	return values
}

// CreateAutonomousContainerDatabaseDetailsProtectionModeEnum Enum with underlying type: string
type CreateAutonomousContainerDatabaseDetailsProtectionModeEnum string

// Set of constants representing the allowable values for CreateAutonomousContainerDatabaseDetailsProtectionModeEnum
const (
	CreateAutonomousContainerDatabaseDetailsProtectionModeAvailability CreateAutonomousContainerDatabaseDetailsProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	CreateAutonomousContainerDatabaseDetailsProtectionModePerformance  CreateAutonomousContainerDatabaseDetailsProtectionModeEnum = "MAXIMUM_PERFORMANCE"
)

var mappingCreateAutonomousContainerDatabaseDetailsProtectionMode = map[string]CreateAutonomousContainerDatabaseDetailsProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": CreateAutonomousContainerDatabaseDetailsProtectionModeAvailability,
	"MAXIMUM_PERFORMANCE":  CreateAutonomousContainerDatabaseDetailsProtectionModePerformance,
}

// GetCreateAutonomousContainerDatabaseDetailsProtectionModeEnumValues Enumerates the set of values for CreateAutonomousContainerDatabaseDetailsProtectionModeEnum
func GetCreateAutonomousContainerDatabaseDetailsProtectionModeEnumValues() []CreateAutonomousContainerDatabaseDetailsProtectionModeEnum {
	values := make([]CreateAutonomousContainerDatabaseDetailsProtectionModeEnum, 0)
	for _, v := range mappingCreateAutonomousContainerDatabaseDetailsProtectionMode {
		values = append(values, v)
	}
	return values
}

// CreateAutonomousContainerDatabaseDetailsPatchModelEnum Enum with underlying type: string
type CreateAutonomousContainerDatabaseDetailsPatchModelEnum string

// Set of constants representing the allowable values for CreateAutonomousContainerDatabaseDetailsPatchModelEnum
const (
	CreateAutonomousContainerDatabaseDetailsPatchModelUpdates         CreateAutonomousContainerDatabaseDetailsPatchModelEnum = "RELEASE_UPDATES"
	CreateAutonomousContainerDatabaseDetailsPatchModelUpdateRevisions CreateAutonomousContainerDatabaseDetailsPatchModelEnum = "RELEASE_UPDATE_REVISIONS"
)

var mappingCreateAutonomousContainerDatabaseDetailsPatchModel = map[string]CreateAutonomousContainerDatabaseDetailsPatchModelEnum{
	"RELEASE_UPDATES":          CreateAutonomousContainerDatabaseDetailsPatchModelUpdates,
	"RELEASE_UPDATE_REVISIONS": CreateAutonomousContainerDatabaseDetailsPatchModelUpdateRevisions,
}

// GetCreateAutonomousContainerDatabaseDetailsPatchModelEnumValues Enumerates the set of values for CreateAutonomousContainerDatabaseDetailsPatchModelEnum
func GetCreateAutonomousContainerDatabaseDetailsPatchModelEnumValues() []CreateAutonomousContainerDatabaseDetailsPatchModelEnum {
	values := make([]CreateAutonomousContainerDatabaseDetailsPatchModelEnum, 0)
	for _, v := range mappingCreateAutonomousContainerDatabaseDetailsPatchModel {
		values = append(values, v)
	}
	return values
}
