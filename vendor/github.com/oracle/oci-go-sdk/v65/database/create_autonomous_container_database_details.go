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

// CreateAutonomousContainerDatabaseDetails Describes the required parameters for the creation of an Autonomous Container Database.
type CreateAutonomousContainerDatabaseDetails struct {

	// The display name for the Autonomous Container Database.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Database Patch model preference.
	PatchModel CreateAutonomousContainerDatabaseDetailsPatchModelEnum `mandatory:"true" json:"patchModel"`

	// **Deprecated.** The `DB_UNIQUE_NAME` value is set by Oracle Cloud Infrastructure.  Do not specify a value for this parameter. Specifying a value for this field will cause Terraform operations to fail.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// The service level agreement type of the Autonomous Container Database. The default is STANDARD. For an autonomous dataguard Autonomous Container Database, the specified Autonomous Exadata Infrastructure must be associated with a remote Autonomous Exadata Infrastructure.
	ServiceLevelAgreementType CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum `mandatory:"false" json:"serviceLevelAgreementType,omitempty"`

	// **No longer used.** This parameter is no longer used for Autonomous Database on dedicated Exadata infrasture. Specify a `cloudAutonomousVmClusterId` instead. Using this parameter will cause the operation to fail.
	AutonomousExadataInfrastructureId *string `mandatory:"false" json:"autonomousExadataInfrastructureId"`

	// *No longer used.* This parameter is no longer used for Autonomous Database on dedicated Exadata infrasture. Specify a `peerCloudAutonomousVmClusterId` instead. Using this parameter will cause the operation to fail.
	PeerAutonomousExadataInfrastructureId *string `mandatory:"false" json:"peerAutonomousExadataInfrastructureId"`

	// The display name for the peer Autonomous Container Database.
	PeerAutonomousContainerDatabaseDisplayName *string `mandatory:"false" json:"peerAutonomousContainerDatabaseDisplayName"`

	// The protection mode of this Autonomous Data Guard association. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode CreateAutonomousContainerDatabaseDetailsProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// Indicates whether Automatic Failover is enabled for Autonomous Container Database Dataguard Association
	IsAutomaticFailoverEnabled *bool `mandatory:"false" json:"isAutomaticFailoverEnabled"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the peer cloud Autonomous Exadata VM Cluster.
	PeerCloudAutonomousVmClusterId *string `mandatory:"false" json:"peerCloudAutonomousVmClusterId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the peer Autonomous VM cluster for Autonomous Data Guard. Required to enable Data Guard.
	PeerAutonomousVmClusterId *string `mandatory:"false" json:"peerAutonomousVmClusterId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment where the standby Autonomous Container Database
	// will be created.
	PeerAutonomousContainerDatabaseCompartmentId *string `mandatory:"false" json:"peerAutonomousContainerDatabaseCompartmentId"`

	PeerAutonomousContainerDatabaseBackupConfig *PeerAutonomousContainerDatabaseBackupConfig `mandatory:"false" json:"peerAutonomousContainerDatabaseBackupConfig"`

	// **Deprecated.** The `DB_UNIQUE_NAME` of the peer Autonomous Container Database in a Data Guard association is set by Oracle Cloud Infrastructure.  Do not specify a value for this parameter. Specifying a value for this field will cause Terraform operations to fail.
	PeerDbUniqueName *string `mandatory:"false" json:"peerDbUniqueName"`

	// The OCID of the Autonomous VM Cluster.
	AutonomousVmClusterId *string `mandatory:"false" json:"autonomousVmClusterId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
	CloudAutonomousVmClusterId *string `mandatory:"false" json:"cloudAutonomousVmClusterId"`

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
