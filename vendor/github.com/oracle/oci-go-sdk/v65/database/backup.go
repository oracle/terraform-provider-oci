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

// Backup The representation of Backup
type Backup struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	// The user-friendly name for the backup. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The type of backup.
	Type BackupTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The date and time the backup started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the backup was completed.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The name of the availability domain where the database backup is stored.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The current state of the backup.
	LifecycleState BackupLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The Oracle Database Edition that applies to all the databases on the DB system.
	// Exadata DB systems and 2-node RAC DB systems require ENTERPRISE_EDITION_EXTREME_PERFORMANCE.
	DatabaseEdition BackupDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`

	// The size of the database in gigabytes at the time the backup was taken.
	DatabaseSizeInGBs *float64 `mandatory:"false" json:"databaseSizeInGBs"`

	// Shape of the backup's source database.
	Shape *string `mandatory:"false" json:"shape"`

	// Version of the backup's source database
	Version *string `mandatory:"false" json:"version"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
	KeyStoreId *string `mandatory:"false" json:"keyStoreId"`

	// The wallet name for Oracle Key Vault.
	KeyStoreWalletName *string `mandatory:"false" json:"keyStoreWalletName"`

	// List of OCIDs of the key containers used as the secondary encryption key in database transparent data encryption (TDE) operations.
	SecondaryKmsKeyIds []string `mandatory:"false" json:"secondaryKmsKeyIds"`

	// The retention period of the long term backup in days.
	RetentionPeriodInDays *int `mandatory:"false" json:"retentionPeriodInDays"`

	// The retention period of the long term backup in years.
	RetentionPeriodInYears *int `mandatory:"false" json:"retentionPeriodInYears"`

	// Expiration time of the long term database backup.
	TimeExpiryScheduled *common.SDKTime `mandatory:"false" json:"timeExpiryScheduled"`

	// True if Oracle Managed Keys is required for restore of the backup.
	IsUsingOracleManagedKeys *bool `mandatory:"false" json:"isUsingOracleManagedKeys"`

	// Type of the backup destination.
	BackupDestinationType BackupBackupDestinationTypeEnum `mandatory:"false" json:"backupDestinationType,omitempty"`

	EncryptionKeyLocationDetails EncryptionKeyLocationDetails `mandatory:"false" json:"encryptionKeyLocationDetails"`
}

func (m Backup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Backup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBackupTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetBackupTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBackupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBackupLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBackupDatabaseEditionEnum(string(m.DatabaseEdition)); !ok && m.DatabaseEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseEdition: %s. Supported values are: %s.", m.DatabaseEdition, strings.Join(GetBackupDatabaseEditionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBackupBackupDestinationTypeEnum(string(m.BackupDestinationType)); !ok && m.BackupDestinationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupDestinationType: %s. Supported values are: %s.", m.BackupDestinationType, strings.Join(GetBackupBackupDestinationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Backup) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Id                           *string                         `json:"id"`
		CompartmentId                *string                         `json:"compartmentId"`
		DatabaseId                   *string                         `json:"databaseId"`
		DisplayName                  *string                         `json:"displayName"`
		Type                         BackupTypeEnum                  `json:"type"`
		TimeStarted                  *common.SDKTime                 `json:"timeStarted"`
		TimeEnded                    *common.SDKTime                 `json:"timeEnded"`
		LifecycleDetails             *string                         `json:"lifecycleDetails"`
		AvailabilityDomain           *string                         `json:"availabilityDomain"`
		LifecycleState               BackupLifecycleStateEnum        `json:"lifecycleState"`
		DatabaseEdition              BackupDatabaseEditionEnum       `json:"databaseEdition"`
		DatabaseSizeInGBs            *float64                        `json:"databaseSizeInGBs"`
		Shape                        *string                         `json:"shape"`
		Version                      *string                         `json:"version"`
		KmsKeyId                     *string                         `json:"kmsKeyId"`
		KmsKeyVersionId              *string                         `json:"kmsKeyVersionId"`
		VaultId                      *string                         `json:"vaultId"`
		KeyStoreId                   *string                         `json:"keyStoreId"`
		KeyStoreWalletName           *string                         `json:"keyStoreWalletName"`
		SecondaryKmsKeyIds           []string                        `json:"secondaryKmsKeyIds"`
		RetentionPeriodInDays        *int                            `json:"retentionPeriodInDays"`
		RetentionPeriodInYears       *int                            `json:"retentionPeriodInYears"`
		TimeExpiryScheduled          *common.SDKTime                 `json:"timeExpiryScheduled"`
		IsUsingOracleManagedKeys     *bool                           `json:"isUsingOracleManagedKeys"`
		BackupDestinationType        BackupBackupDestinationTypeEnum `json:"backupDestinationType"`
		EncryptionKeyLocationDetails encryptionkeylocationdetails    `json:"encryptionKeyLocationDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DatabaseId = model.DatabaseId

	m.DisplayName = model.DisplayName

	m.Type = model.Type

	m.TimeStarted = model.TimeStarted

	m.TimeEnded = model.TimeEnded

	m.LifecycleDetails = model.LifecycleDetails

	m.AvailabilityDomain = model.AvailabilityDomain

	m.LifecycleState = model.LifecycleState

	m.DatabaseEdition = model.DatabaseEdition

	m.DatabaseSizeInGBs = model.DatabaseSizeInGBs

	m.Shape = model.Shape

	m.Version = model.Version

	m.KmsKeyId = model.KmsKeyId

	m.KmsKeyVersionId = model.KmsKeyVersionId

	m.VaultId = model.VaultId

	m.KeyStoreId = model.KeyStoreId

	m.KeyStoreWalletName = model.KeyStoreWalletName

	m.SecondaryKmsKeyIds = make([]string, len(model.SecondaryKmsKeyIds))
	copy(m.SecondaryKmsKeyIds, model.SecondaryKmsKeyIds)
	m.RetentionPeriodInDays = model.RetentionPeriodInDays

	m.RetentionPeriodInYears = model.RetentionPeriodInYears

	m.TimeExpiryScheduled = model.TimeExpiryScheduled

	m.IsUsingOracleManagedKeys = model.IsUsingOracleManagedKeys

	m.BackupDestinationType = model.BackupDestinationType

	nn, e = model.EncryptionKeyLocationDetails.UnmarshalPolymorphicJSON(model.EncryptionKeyLocationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.EncryptionKeyLocationDetails = nn.(EncryptionKeyLocationDetails)
	} else {
		m.EncryptionKeyLocationDetails = nil
	}

	return
}

// BackupTypeEnum Enum with underlying type: string
type BackupTypeEnum string

// Set of constants representing the allowable values for BackupTypeEnum
const (
	BackupTypeIncremental BackupTypeEnum = "INCREMENTAL"
	BackupTypeFull        BackupTypeEnum = "FULL"
	BackupTypeVirtualFull BackupTypeEnum = "VIRTUAL_FULL"
)

var mappingBackupTypeEnum = map[string]BackupTypeEnum{
	"INCREMENTAL":  BackupTypeIncremental,
	"FULL":         BackupTypeFull,
	"VIRTUAL_FULL": BackupTypeVirtualFull,
}

var mappingBackupTypeEnumLowerCase = map[string]BackupTypeEnum{
	"incremental":  BackupTypeIncremental,
	"full":         BackupTypeFull,
	"virtual_full": BackupTypeVirtualFull,
}

// GetBackupTypeEnumValues Enumerates the set of values for BackupTypeEnum
func GetBackupTypeEnumValues() []BackupTypeEnum {
	values := make([]BackupTypeEnum, 0)
	for _, v := range mappingBackupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupTypeEnumStringValues Enumerates the set of values in String for BackupTypeEnum
func GetBackupTypeEnumStringValues() []string {
	return []string{
		"INCREMENTAL",
		"FULL",
		"VIRTUAL_FULL",
	}
}

// GetMappingBackupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupTypeEnum(val string) (BackupTypeEnum, bool) {
	enum, ok := mappingBackupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BackupLifecycleStateEnum Enum with underlying type: string
type BackupLifecycleStateEnum string

// Set of constants representing the allowable values for BackupLifecycleStateEnum
const (
	BackupLifecycleStateCreating  BackupLifecycleStateEnum = "CREATING"
	BackupLifecycleStateActive    BackupLifecycleStateEnum = "ACTIVE"
	BackupLifecycleStateDeleting  BackupLifecycleStateEnum = "DELETING"
	BackupLifecycleStateDeleted   BackupLifecycleStateEnum = "DELETED"
	BackupLifecycleStateFailed    BackupLifecycleStateEnum = "FAILED"
	BackupLifecycleStateRestoring BackupLifecycleStateEnum = "RESTORING"
	BackupLifecycleStateUpdating  BackupLifecycleStateEnum = "UPDATING"
	BackupLifecycleStateCanceling BackupLifecycleStateEnum = "CANCELING"
	BackupLifecycleStateCanceled  BackupLifecycleStateEnum = "CANCELED"
)

var mappingBackupLifecycleStateEnum = map[string]BackupLifecycleStateEnum{
	"CREATING":  BackupLifecycleStateCreating,
	"ACTIVE":    BackupLifecycleStateActive,
	"DELETING":  BackupLifecycleStateDeleting,
	"DELETED":   BackupLifecycleStateDeleted,
	"FAILED":    BackupLifecycleStateFailed,
	"RESTORING": BackupLifecycleStateRestoring,
	"UPDATING":  BackupLifecycleStateUpdating,
	"CANCELING": BackupLifecycleStateCanceling,
	"CANCELED":  BackupLifecycleStateCanceled,
}

var mappingBackupLifecycleStateEnumLowerCase = map[string]BackupLifecycleStateEnum{
	"creating":  BackupLifecycleStateCreating,
	"active":    BackupLifecycleStateActive,
	"deleting":  BackupLifecycleStateDeleting,
	"deleted":   BackupLifecycleStateDeleted,
	"failed":    BackupLifecycleStateFailed,
	"restoring": BackupLifecycleStateRestoring,
	"updating":  BackupLifecycleStateUpdating,
	"canceling": BackupLifecycleStateCanceling,
	"canceled":  BackupLifecycleStateCanceled,
}

// GetBackupLifecycleStateEnumValues Enumerates the set of values for BackupLifecycleStateEnum
func GetBackupLifecycleStateEnumValues() []BackupLifecycleStateEnum {
	values := make([]BackupLifecycleStateEnum, 0)
	for _, v := range mappingBackupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupLifecycleStateEnumStringValues Enumerates the set of values in String for BackupLifecycleStateEnum
func GetBackupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"RESTORING",
		"UPDATING",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingBackupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupLifecycleStateEnum(val string) (BackupLifecycleStateEnum, bool) {
	enum, ok := mappingBackupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BackupDatabaseEditionEnum Enum with underlying type: string
type BackupDatabaseEditionEnum string

// Set of constants representing the allowable values for BackupDatabaseEditionEnum
const (
	BackupDatabaseEditionStandardEdition                     BackupDatabaseEditionEnum = "STANDARD_EDITION"
	BackupDatabaseEditionEnterpriseEdition                   BackupDatabaseEditionEnum = "ENTERPRISE_EDITION"
	BackupDatabaseEditionEnterpriseEditionHighPerformance    BackupDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	BackupDatabaseEditionEnterpriseEditionExtremePerformance BackupDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
	BackupDatabaseEditionEnterpriseEditionDeveloper          BackupDatabaseEditionEnum = "ENTERPRISE_EDITION_DEVELOPER"
)

var mappingBackupDatabaseEditionEnum = map[string]BackupDatabaseEditionEnum{
	"STANDARD_EDITION":                       BackupDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     BackupDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    BackupDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": BackupDatabaseEditionEnterpriseEditionExtremePerformance,
	"ENTERPRISE_EDITION_DEVELOPER":           BackupDatabaseEditionEnterpriseEditionDeveloper,
}

var mappingBackupDatabaseEditionEnumLowerCase = map[string]BackupDatabaseEditionEnum{
	"standard_edition":                       BackupDatabaseEditionStandardEdition,
	"enterprise_edition":                     BackupDatabaseEditionEnterpriseEdition,
	"enterprise_edition_high_performance":    BackupDatabaseEditionEnterpriseEditionHighPerformance,
	"enterprise_edition_extreme_performance": BackupDatabaseEditionEnterpriseEditionExtremePerformance,
	"enterprise_edition_developer":           BackupDatabaseEditionEnterpriseEditionDeveloper,
}

// GetBackupDatabaseEditionEnumValues Enumerates the set of values for BackupDatabaseEditionEnum
func GetBackupDatabaseEditionEnumValues() []BackupDatabaseEditionEnum {
	values := make([]BackupDatabaseEditionEnum, 0)
	for _, v := range mappingBackupDatabaseEditionEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupDatabaseEditionEnumStringValues Enumerates the set of values in String for BackupDatabaseEditionEnum
func GetBackupDatabaseEditionEnumStringValues() []string {
	return []string{
		"STANDARD_EDITION",
		"ENTERPRISE_EDITION",
		"ENTERPRISE_EDITION_HIGH_PERFORMANCE",
		"ENTERPRISE_EDITION_EXTREME_PERFORMANCE",
		"ENTERPRISE_EDITION_DEVELOPER",
	}
}

// GetMappingBackupDatabaseEditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupDatabaseEditionEnum(val string) (BackupDatabaseEditionEnum, bool) {
	enum, ok := mappingBackupDatabaseEditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BackupBackupDestinationTypeEnum Enum with underlying type: string
type BackupBackupDestinationTypeEnum string

// Set of constants representing the allowable values for BackupBackupDestinationTypeEnum
const (
	BackupBackupDestinationTypeObjectStore BackupBackupDestinationTypeEnum = "OBJECT_STORE"
	BackupBackupDestinationTypeDbrs        BackupBackupDestinationTypeEnum = "DBRS"
	BackupBackupDestinationTypeAwsS3       BackupBackupDestinationTypeEnum = "AWS_S3"
)

var mappingBackupBackupDestinationTypeEnum = map[string]BackupBackupDestinationTypeEnum{
	"OBJECT_STORE": BackupBackupDestinationTypeObjectStore,
	"DBRS":         BackupBackupDestinationTypeDbrs,
	"AWS_S3":       BackupBackupDestinationTypeAwsS3,
}

var mappingBackupBackupDestinationTypeEnumLowerCase = map[string]BackupBackupDestinationTypeEnum{
	"object_store": BackupBackupDestinationTypeObjectStore,
	"dbrs":         BackupBackupDestinationTypeDbrs,
	"aws_s3":       BackupBackupDestinationTypeAwsS3,
}

// GetBackupBackupDestinationTypeEnumValues Enumerates the set of values for BackupBackupDestinationTypeEnum
func GetBackupBackupDestinationTypeEnumValues() []BackupBackupDestinationTypeEnum {
	values := make([]BackupBackupDestinationTypeEnum, 0)
	for _, v := range mappingBackupBackupDestinationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupBackupDestinationTypeEnumStringValues Enumerates the set of values in String for BackupBackupDestinationTypeEnum
func GetBackupBackupDestinationTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORE",
		"DBRS",
		"AWS_S3",
	}
}

// GetMappingBackupBackupDestinationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupBackupDestinationTypeEnum(val string) (BackupBackupDestinationTypeEnum, bool) {
	enum, ok := mappingBackupBackupDestinationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
