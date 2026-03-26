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

// AutonomousDatabaseBackup An Autonomous AI Database backup.
type AutonomousDatabaseBackup struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous AI Database backup.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous AI Database.
	AutonomousDatabaseId *string `mandatory:"true" json:"autonomousDatabaseId"`

	// The user-friendly name for the backup. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of backup.
	Type AutonomousDatabaseBackupTypeEnum `mandatory:"true" json:"type"`

	// Indicates whether the backup is user-initiated or automatic.
	IsAutomatic *bool `mandatory:"true" json:"isAutomatic"`

	// The current state of the backup.
	LifecycleState AutonomousDatabaseBackupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The destination cloud provider where Autonomous AI Database backups are stored.
	BackupDestination AutonomousDatabaseBackupBackupDestinationEnum `mandatory:"false" json:"backupDestination,omitempty"`

	// The date and time the backup started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the backup completed.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The size of the database in terabytes at the time the backup was taken.
	DatabaseSizeInTBs *float32 `mandatory:"false" json:"databaseSizeInTBs"`

	// Indicates whether the backup can be used to restore the associated Autonomous AI Database.
	IsRestorable *bool `mandatory:"false" json:"isRestorable"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
	KeyStoreId *string `mandatory:"false" json:"keyStoreId"`

	// The wallet name for Oracle Key Vault.
	KeyStoreWalletName *string `mandatory:"false" json:"keyStoreWalletName"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous AI Database Serverless does not use key versions, hence is not applicable for Autonomous AI Database Serverless instances.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	EncryptionKey AutonomousDatabaseEncryptionKeyDetails `mandatory:"false" json:"encryptionKey"`

	// Retention period, in days, for backups
	RetentionPeriodInDays *int `mandatory:"false" json:"retentionPeriodInDays"`

	// Timestamp until when the backup will be available
	TimeAvailableTill *common.SDKTime `mandatory:"false" json:"timeAvailableTill"`

	// A valid Oracle AI Database version for Autonomous AI Database.
	// When you specify 23ai for dbversion, the system will provision a 23ai database, but the UI will display it as 26ai.
	// When you specify 26ai for dbversion, the system will provision and display a 26ai database as expected.
	// For new databases, it is recommended to use either 19c or 26ai.
	// **Note** Starting December 2026, 23ai will not be supported as a valid value for this parameter.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	// The backup size in terrabytes (TB).
	SizeInTBs *float64 `mandatory:"false" json:"sizeInTBs"`

	BackupDestinationDetails *BackupDestinationDetails `mandatory:"false" json:"backupDestinationDetails"`

	// The infrastructure type this resource belongs to.
	InfrastructureType AutonomousDatabaseBackupInfrastructureTypeEnum `mandatory:"false" json:"infrastructureType,omitempty"`

	// Name of the region in which backup is taken in.
	Region *string `mandatory:"false" json:"region"`

	SourceDatabaseDetails *SourceDatabaseDetails `mandatory:"false" json:"sourceDatabaseDetails"`
}

func (m AutonomousDatabaseBackup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDatabaseBackup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousDatabaseBackupTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAutonomousDatabaseBackupTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseBackupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousDatabaseBackupLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAutonomousDatabaseBackupBackupDestinationEnum(string(m.BackupDestination)); !ok && m.BackupDestination != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupDestination: %s. Supported values are: %s.", m.BackupDestination, strings.Join(GetAutonomousDatabaseBackupBackupDestinationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseBackupInfrastructureTypeEnum(string(m.InfrastructureType)); !ok && m.InfrastructureType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InfrastructureType: %s. Supported values are: %s.", m.InfrastructureType, strings.Join(GetAutonomousDatabaseBackupInfrastructureTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *AutonomousDatabaseBackup) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		BackupDestination        AutonomousDatabaseBackupBackupDestinationEnum  `json:"backupDestination"`
		TimeStarted              *common.SDKTime                                `json:"timeStarted"`
		TimeEnded                *common.SDKTime                                `json:"timeEnded"`
		LifecycleDetails         *string                                        `json:"lifecycleDetails"`
		DatabaseSizeInTBs        *float32                                       `json:"databaseSizeInTBs"`
		IsRestorable             *bool                                          `json:"isRestorable"`
		KeyStoreId               *string                                        `json:"keyStoreId"`
		KeyStoreWalletName       *string                                        `json:"keyStoreWalletName"`
		KmsKeyId                 *string                                        `json:"kmsKeyId"`
		VaultId                  *string                                        `json:"vaultId"`
		KmsKeyVersionId          *string                                        `json:"kmsKeyVersionId"`
		EncryptionKey            autonomousdatabaseencryptionkeydetails         `json:"encryptionKey"`
		RetentionPeriodInDays    *int                                           `json:"retentionPeriodInDays"`
		TimeAvailableTill        *common.SDKTime                                `json:"timeAvailableTill"`
		DbVersion                *string                                        `json:"dbVersion"`
		SizeInTBs                *float64                                       `json:"sizeInTBs"`
		BackupDestinationDetails *BackupDestinationDetails                      `json:"backupDestinationDetails"`
		InfrastructureType       AutonomousDatabaseBackupInfrastructureTypeEnum `json:"infrastructureType"`
		Region                   *string                                        `json:"region"`
		SourceDatabaseDetails    *SourceDatabaseDetails                         `json:"sourceDatabaseDetails"`
		Id                       *string                                        `json:"id"`
		CompartmentId            *string                                        `json:"compartmentId"`
		AutonomousDatabaseId     *string                                        `json:"autonomousDatabaseId"`
		DisplayName              *string                                        `json:"displayName"`
		Type                     AutonomousDatabaseBackupTypeEnum               `json:"type"`
		IsAutomatic              *bool                                          `json:"isAutomatic"`
		LifecycleState           AutonomousDatabaseBackupLifecycleStateEnum     `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.BackupDestination = model.BackupDestination

	m.TimeStarted = model.TimeStarted

	m.TimeEnded = model.TimeEnded

	m.LifecycleDetails = model.LifecycleDetails

	m.DatabaseSizeInTBs = model.DatabaseSizeInTBs

	m.IsRestorable = model.IsRestorable

	m.KeyStoreId = model.KeyStoreId

	m.KeyStoreWalletName = model.KeyStoreWalletName

	m.KmsKeyId = model.KmsKeyId

	m.VaultId = model.VaultId

	m.KmsKeyVersionId = model.KmsKeyVersionId

	nn, e = model.EncryptionKey.UnmarshalPolymorphicJSON(model.EncryptionKey.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.EncryptionKey = nn.(AutonomousDatabaseEncryptionKeyDetails)
	} else {
		m.EncryptionKey = nil
	}

	m.RetentionPeriodInDays = model.RetentionPeriodInDays

	m.TimeAvailableTill = model.TimeAvailableTill

	m.DbVersion = model.DbVersion

	m.SizeInTBs = model.SizeInTBs

	m.BackupDestinationDetails = model.BackupDestinationDetails

	m.InfrastructureType = model.InfrastructureType

	m.Region = model.Region

	m.SourceDatabaseDetails = model.SourceDatabaseDetails

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.AutonomousDatabaseId = model.AutonomousDatabaseId

	m.DisplayName = model.DisplayName

	m.Type = model.Type

	m.IsAutomatic = model.IsAutomatic

	m.LifecycleState = model.LifecycleState

	return
}

// AutonomousDatabaseBackupBackupDestinationEnum Enum with underlying type: string
type AutonomousDatabaseBackupBackupDestinationEnum string

// Set of constants representing the allowable values for AutonomousDatabaseBackupBackupDestinationEnum
const (
	AutonomousDatabaseBackupBackupDestinationOci   AutonomousDatabaseBackupBackupDestinationEnum = "OCI"
	AutonomousDatabaseBackupBackupDestinationGcp   AutonomousDatabaseBackupBackupDestinationEnum = "GCP"
	AutonomousDatabaseBackupBackupDestinationAzure AutonomousDatabaseBackupBackupDestinationEnum = "AZURE"
	AutonomousDatabaseBackupBackupDestinationAws   AutonomousDatabaseBackupBackupDestinationEnum = "AWS"
)

var mappingAutonomousDatabaseBackupBackupDestinationEnum = map[string]AutonomousDatabaseBackupBackupDestinationEnum{
	"OCI":   AutonomousDatabaseBackupBackupDestinationOci,
	"GCP":   AutonomousDatabaseBackupBackupDestinationGcp,
	"AZURE": AutonomousDatabaseBackupBackupDestinationAzure,
	"AWS":   AutonomousDatabaseBackupBackupDestinationAws,
}

var mappingAutonomousDatabaseBackupBackupDestinationEnumLowerCase = map[string]AutonomousDatabaseBackupBackupDestinationEnum{
	"oci":   AutonomousDatabaseBackupBackupDestinationOci,
	"gcp":   AutonomousDatabaseBackupBackupDestinationGcp,
	"azure": AutonomousDatabaseBackupBackupDestinationAzure,
	"aws":   AutonomousDatabaseBackupBackupDestinationAws,
}

// GetAutonomousDatabaseBackupBackupDestinationEnumValues Enumerates the set of values for AutonomousDatabaseBackupBackupDestinationEnum
func GetAutonomousDatabaseBackupBackupDestinationEnumValues() []AutonomousDatabaseBackupBackupDestinationEnum {
	values := make([]AutonomousDatabaseBackupBackupDestinationEnum, 0)
	for _, v := range mappingAutonomousDatabaseBackupBackupDestinationEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseBackupBackupDestinationEnumStringValues Enumerates the set of values in String for AutonomousDatabaseBackupBackupDestinationEnum
func GetAutonomousDatabaseBackupBackupDestinationEnumStringValues() []string {
	return []string{
		"OCI",
		"GCP",
		"AZURE",
		"AWS",
	}
}

// GetMappingAutonomousDatabaseBackupBackupDestinationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseBackupBackupDestinationEnum(val string) (AutonomousDatabaseBackupBackupDestinationEnum, bool) {
	enum, ok := mappingAutonomousDatabaseBackupBackupDestinationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseBackupTypeEnum Enum with underlying type: string
type AutonomousDatabaseBackupTypeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseBackupTypeEnum
const (
	AutonomousDatabaseBackupTypeIncremental           AutonomousDatabaseBackupTypeEnum = "INCREMENTAL"
	AutonomousDatabaseBackupTypeFull                  AutonomousDatabaseBackupTypeEnum = "FULL"
	AutonomousDatabaseBackupTypeLongterm              AutonomousDatabaseBackupTypeEnum = "LONGTERM"
	AutonomousDatabaseBackupTypeVirtualFull           AutonomousDatabaseBackupTypeEnum = "VIRTUAL_FULL"
	AutonomousDatabaseBackupTypeCumulativeIncremental AutonomousDatabaseBackupTypeEnum = "CUMULATIVE_INCREMENTAL"
	AutonomousDatabaseBackupTypeRollForwardImageCopy  AutonomousDatabaseBackupTypeEnum = "ROLL_FORWARD_IMAGE_COPY"
)

var mappingAutonomousDatabaseBackupTypeEnum = map[string]AutonomousDatabaseBackupTypeEnum{
	"INCREMENTAL":             AutonomousDatabaseBackupTypeIncremental,
	"FULL":                    AutonomousDatabaseBackupTypeFull,
	"LONGTERM":                AutonomousDatabaseBackupTypeLongterm,
	"VIRTUAL_FULL":            AutonomousDatabaseBackupTypeVirtualFull,
	"CUMULATIVE_INCREMENTAL":  AutonomousDatabaseBackupTypeCumulativeIncremental,
	"ROLL_FORWARD_IMAGE_COPY": AutonomousDatabaseBackupTypeRollForwardImageCopy,
}

var mappingAutonomousDatabaseBackupTypeEnumLowerCase = map[string]AutonomousDatabaseBackupTypeEnum{
	"incremental":             AutonomousDatabaseBackupTypeIncremental,
	"full":                    AutonomousDatabaseBackupTypeFull,
	"longterm":                AutonomousDatabaseBackupTypeLongterm,
	"virtual_full":            AutonomousDatabaseBackupTypeVirtualFull,
	"cumulative_incremental":  AutonomousDatabaseBackupTypeCumulativeIncremental,
	"roll_forward_image_copy": AutonomousDatabaseBackupTypeRollForwardImageCopy,
}

// GetAutonomousDatabaseBackupTypeEnumValues Enumerates the set of values for AutonomousDatabaseBackupTypeEnum
func GetAutonomousDatabaseBackupTypeEnumValues() []AutonomousDatabaseBackupTypeEnum {
	values := make([]AutonomousDatabaseBackupTypeEnum, 0)
	for _, v := range mappingAutonomousDatabaseBackupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseBackupTypeEnumStringValues Enumerates the set of values in String for AutonomousDatabaseBackupTypeEnum
func GetAutonomousDatabaseBackupTypeEnumStringValues() []string {
	return []string{
		"INCREMENTAL",
		"FULL",
		"LONGTERM",
		"VIRTUAL_FULL",
		"CUMULATIVE_INCREMENTAL",
		"ROLL_FORWARD_IMAGE_COPY",
	}
}

// GetMappingAutonomousDatabaseBackupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseBackupTypeEnum(val string) (AutonomousDatabaseBackupTypeEnum, bool) {
	enum, ok := mappingAutonomousDatabaseBackupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseBackupLifecycleStateEnum Enum with underlying type: string
type AutonomousDatabaseBackupLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseBackupLifecycleStateEnum
const (
	AutonomousDatabaseBackupLifecycleStateCreating AutonomousDatabaseBackupLifecycleStateEnum = "CREATING"
	AutonomousDatabaseBackupLifecycleStateActive   AutonomousDatabaseBackupLifecycleStateEnum = "ACTIVE"
	AutonomousDatabaseBackupLifecycleStateDeleting AutonomousDatabaseBackupLifecycleStateEnum = "DELETING"
	AutonomousDatabaseBackupLifecycleStateDeleted  AutonomousDatabaseBackupLifecycleStateEnum = "DELETED"
	AutonomousDatabaseBackupLifecycleStateFailed   AutonomousDatabaseBackupLifecycleStateEnum = "FAILED"
	AutonomousDatabaseBackupLifecycleStateUpdating AutonomousDatabaseBackupLifecycleStateEnum = "UPDATING"
)

var mappingAutonomousDatabaseBackupLifecycleStateEnum = map[string]AutonomousDatabaseBackupLifecycleStateEnum{
	"CREATING": AutonomousDatabaseBackupLifecycleStateCreating,
	"ACTIVE":   AutonomousDatabaseBackupLifecycleStateActive,
	"DELETING": AutonomousDatabaseBackupLifecycleStateDeleting,
	"DELETED":  AutonomousDatabaseBackupLifecycleStateDeleted,
	"FAILED":   AutonomousDatabaseBackupLifecycleStateFailed,
	"UPDATING": AutonomousDatabaseBackupLifecycleStateUpdating,
}

var mappingAutonomousDatabaseBackupLifecycleStateEnumLowerCase = map[string]AutonomousDatabaseBackupLifecycleStateEnum{
	"creating": AutonomousDatabaseBackupLifecycleStateCreating,
	"active":   AutonomousDatabaseBackupLifecycleStateActive,
	"deleting": AutonomousDatabaseBackupLifecycleStateDeleting,
	"deleted":  AutonomousDatabaseBackupLifecycleStateDeleted,
	"failed":   AutonomousDatabaseBackupLifecycleStateFailed,
	"updating": AutonomousDatabaseBackupLifecycleStateUpdating,
}

// GetAutonomousDatabaseBackupLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseBackupLifecycleStateEnum
func GetAutonomousDatabaseBackupLifecycleStateEnumValues() []AutonomousDatabaseBackupLifecycleStateEnum {
	values := make([]AutonomousDatabaseBackupLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseBackupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseBackupLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousDatabaseBackupLifecycleStateEnum
func GetAutonomousDatabaseBackupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingAutonomousDatabaseBackupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseBackupLifecycleStateEnum(val string) (AutonomousDatabaseBackupLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousDatabaseBackupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseBackupInfrastructureTypeEnum Enum with underlying type: string
type AutonomousDatabaseBackupInfrastructureTypeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseBackupInfrastructureTypeEnum
const (
	AutonomousDatabaseBackupInfrastructureTypeCloud           AutonomousDatabaseBackupInfrastructureTypeEnum = "CLOUD"
	AutonomousDatabaseBackupInfrastructureTypeCloudAtCustomer AutonomousDatabaseBackupInfrastructureTypeEnum = "CLOUD_AT_CUSTOMER"
)

var mappingAutonomousDatabaseBackupInfrastructureTypeEnum = map[string]AutonomousDatabaseBackupInfrastructureTypeEnum{
	"CLOUD":             AutonomousDatabaseBackupInfrastructureTypeCloud,
	"CLOUD_AT_CUSTOMER": AutonomousDatabaseBackupInfrastructureTypeCloudAtCustomer,
}

var mappingAutonomousDatabaseBackupInfrastructureTypeEnumLowerCase = map[string]AutonomousDatabaseBackupInfrastructureTypeEnum{
	"cloud":             AutonomousDatabaseBackupInfrastructureTypeCloud,
	"cloud_at_customer": AutonomousDatabaseBackupInfrastructureTypeCloudAtCustomer,
}

// GetAutonomousDatabaseBackupInfrastructureTypeEnumValues Enumerates the set of values for AutonomousDatabaseBackupInfrastructureTypeEnum
func GetAutonomousDatabaseBackupInfrastructureTypeEnumValues() []AutonomousDatabaseBackupInfrastructureTypeEnum {
	values := make([]AutonomousDatabaseBackupInfrastructureTypeEnum, 0)
	for _, v := range mappingAutonomousDatabaseBackupInfrastructureTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseBackupInfrastructureTypeEnumStringValues Enumerates the set of values in String for AutonomousDatabaseBackupInfrastructureTypeEnum
func GetAutonomousDatabaseBackupInfrastructureTypeEnumStringValues() []string {
	return []string{
		"CLOUD",
		"CLOUD_AT_CUSTOMER",
	}
}

// GetMappingAutonomousDatabaseBackupInfrastructureTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseBackupInfrastructureTypeEnum(val string) (AutonomousDatabaseBackupInfrastructureTypeEnum, bool) {
	enum, ok := mappingAutonomousDatabaseBackupInfrastructureTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
