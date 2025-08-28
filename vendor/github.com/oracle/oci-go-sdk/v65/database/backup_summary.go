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

// BackupSummary A database backup.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type BackupSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	// The user-friendly name for the backup. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The type of backup.
	Type BackupSummaryTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The date and time the backup started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the backup was completed.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The name of the availability domain where the database backup is stored.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The current state of the backup.
	LifecycleState BackupSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The Oracle Database Edition that applies to all the databases on the DB system.
	// Exadata DB systems and 2-node RAC DB systems require ENTERPRISE_EDITION_EXTREME_PERFORMANCE.
	DatabaseEdition BackupSummaryDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`

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
	BackupDestinationType BackupSummaryBackupDestinationTypeEnum `mandatory:"false" json:"backupDestinationType,omitempty"`

	EncryptionKeyLocationDetails EncryptionKeyLocationDetails `mandatory:"false" json:"encryptionKeyLocationDetails"`
}

func (m BackupSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackupSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBackupSummaryTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetBackupSummaryTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBackupSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBackupSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBackupSummaryDatabaseEditionEnum(string(m.DatabaseEdition)); !ok && m.DatabaseEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseEdition: %s. Supported values are: %s.", m.DatabaseEdition, strings.Join(GetBackupSummaryDatabaseEditionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBackupSummaryBackupDestinationTypeEnum(string(m.BackupDestinationType)); !ok && m.BackupDestinationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupDestinationType: %s. Supported values are: %s.", m.BackupDestinationType, strings.Join(GetBackupSummaryBackupDestinationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *BackupSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Id                           *string                                `json:"id"`
		CompartmentId                *string                                `json:"compartmentId"`
		DatabaseId                   *string                                `json:"databaseId"`
		DisplayName                  *string                                `json:"displayName"`
		Type                         BackupSummaryTypeEnum                  `json:"type"`
		TimeStarted                  *common.SDKTime                        `json:"timeStarted"`
		TimeEnded                    *common.SDKTime                        `json:"timeEnded"`
		LifecycleDetails             *string                                `json:"lifecycleDetails"`
		AvailabilityDomain           *string                                `json:"availabilityDomain"`
		LifecycleState               BackupSummaryLifecycleStateEnum        `json:"lifecycleState"`
		DatabaseEdition              BackupSummaryDatabaseEditionEnum       `json:"databaseEdition"`
		DatabaseSizeInGBs            *float64                               `json:"databaseSizeInGBs"`
		Shape                        *string                                `json:"shape"`
		Version                      *string                                `json:"version"`
		KmsKeyId                     *string                                `json:"kmsKeyId"`
		KmsKeyVersionId              *string                                `json:"kmsKeyVersionId"`
		VaultId                      *string                                `json:"vaultId"`
		KeyStoreId                   *string                                `json:"keyStoreId"`
		KeyStoreWalletName           *string                                `json:"keyStoreWalletName"`
		SecondaryKmsKeyIds           []string                               `json:"secondaryKmsKeyIds"`
		RetentionPeriodInDays        *int                                   `json:"retentionPeriodInDays"`
		RetentionPeriodInYears       *int                                   `json:"retentionPeriodInYears"`
		TimeExpiryScheduled          *common.SDKTime                        `json:"timeExpiryScheduled"`
		IsUsingOracleManagedKeys     *bool                                  `json:"isUsingOracleManagedKeys"`
		BackupDestinationType        BackupSummaryBackupDestinationTypeEnum `json:"backupDestinationType"`
		EncryptionKeyLocationDetails encryptionkeylocationdetails           `json:"encryptionKeyLocationDetails"`
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

// BackupSummaryTypeEnum Enum with underlying type: string
type BackupSummaryTypeEnum string

// Set of constants representing the allowable values for BackupSummaryTypeEnum
const (
	BackupSummaryTypeIncremental BackupSummaryTypeEnum = "INCREMENTAL"
	BackupSummaryTypeFull        BackupSummaryTypeEnum = "FULL"
	BackupSummaryTypeVirtualFull BackupSummaryTypeEnum = "VIRTUAL_FULL"
)

var mappingBackupSummaryTypeEnum = map[string]BackupSummaryTypeEnum{
	"INCREMENTAL":  BackupSummaryTypeIncremental,
	"FULL":         BackupSummaryTypeFull,
	"VIRTUAL_FULL": BackupSummaryTypeVirtualFull,
}

var mappingBackupSummaryTypeEnumLowerCase = map[string]BackupSummaryTypeEnum{
	"incremental":  BackupSummaryTypeIncremental,
	"full":         BackupSummaryTypeFull,
	"virtual_full": BackupSummaryTypeVirtualFull,
}

// GetBackupSummaryTypeEnumValues Enumerates the set of values for BackupSummaryTypeEnum
func GetBackupSummaryTypeEnumValues() []BackupSummaryTypeEnum {
	values := make([]BackupSummaryTypeEnum, 0)
	for _, v := range mappingBackupSummaryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupSummaryTypeEnumStringValues Enumerates the set of values in String for BackupSummaryTypeEnum
func GetBackupSummaryTypeEnumStringValues() []string {
	return []string{
		"INCREMENTAL",
		"FULL",
		"VIRTUAL_FULL",
	}
}

// GetMappingBackupSummaryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupSummaryTypeEnum(val string) (BackupSummaryTypeEnum, bool) {
	enum, ok := mappingBackupSummaryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BackupSummaryLifecycleStateEnum Enum with underlying type: string
type BackupSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for BackupSummaryLifecycleStateEnum
const (
	BackupSummaryLifecycleStateCreating  BackupSummaryLifecycleStateEnum = "CREATING"
	BackupSummaryLifecycleStateActive    BackupSummaryLifecycleStateEnum = "ACTIVE"
	BackupSummaryLifecycleStateDeleting  BackupSummaryLifecycleStateEnum = "DELETING"
	BackupSummaryLifecycleStateDeleted   BackupSummaryLifecycleStateEnum = "DELETED"
	BackupSummaryLifecycleStateFailed    BackupSummaryLifecycleStateEnum = "FAILED"
	BackupSummaryLifecycleStateRestoring BackupSummaryLifecycleStateEnum = "RESTORING"
	BackupSummaryLifecycleStateUpdating  BackupSummaryLifecycleStateEnum = "UPDATING"
	BackupSummaryLifecycleStateCanceling BackupSummaryLifecycleStateEnum = "CANCELING"
	BackupSummaryLifecycleStateCanceled  BackupSummaryLifecycleStateEnum = "CANCELED"
)

var mappingBackupSummaryLifecycleStateEnum = map[string]BackupSummaryLifecycleStateEnum{
	"CREATING":  BackupSummaryLifecycleStateCreating,
	"ACTIVE":    BackupSummaryLifecycleStateActive,
	"DELETING":  BackupSummaryLifecycleStateDeleting,
	"DELETED":   BackupSummaryLifecycleStateDeleted,
	"FAILED":    BackupSummaryLifecycleStateFailed,
	"RESTORING": BackupSummaryLifecycleStateRestoring,
	"UPDATING":  BackupSummaryLifecycleStateUpdating,
	"CANCELING": BackupSummaryLifecycleStateCanceling,
	"CANCELED":  BackupSummaryLifecycleStateCanceled,
}

var mappingBackupSummaryLifecycleStateEnumLowerCase = map[string]BackupSummaryLifecycleStateEnum{
	"creating":  BackupSummaryLifecycleStateCreating,
	"active":    BackupSummaryLifecycleStateActive,
	"deleting":  BackupSummaryLifecycleStateDeleting,
	"deleted":   BackupSummaryLifecycleStateDeleted,
	"failed":    BackupSummaryLifecycleStateFailed,
	"restoring": BackupSummaryLifecycleStateRestoring,
	"updating":  BackupSummaryLifecycleStateUpdating,
	"canceling": BackupSummaryLifecycleStateCanceling,
	"canceled":  BackupSummaryLifecycleStateCanceled,
}

// GetBackupSummaryLifecycleStateEnumValues Enumerates the set of values for BackupSummaryLifecycleStateEnum
func GetBackupSummaryLifecycleStateEnumValues() []BackupSummaryLifecycleStateEnum {
	values := make([]BackupSummaryLifecycleStateEnum, 0)
	for _, v := range mappingBackupSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for BackupSummaryLifecycleStateEnum
func GetBackupSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingBackupSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupSummaryLifecycleStateEnum(val string) (BackupSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingBackupSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BackupSummaryDatabaseEditionEnum Enum with underlying type: string
type BackupSummaryDatabaseEditionEnum string

// Set of constants representing the allowable values for BackupSummaryDatabaseEditionEnum
const (
	BackupSummaryDatabaseEditionStandardEdition                     BackupSummaryDatabaseEditionEnum = "STANDARD_EDITION"
	BackupSummaryDatabaseEditionEnterpriseEdition                   BackupSummaryDatabaseEditionEnum = "ENTERPRISE_EDITION"
	BackupSummaryDatabaseEditionEnterpriseEditionHighPerformance    BackupSummaryDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	BackupSummaryDatabaseEditionEnterpriseEditionExtremePerformance BackupSummaryDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
	BackupSummaryDatabaseEditionEnterpriseEditionDeveloper          BackupSummaryDatabaseEditionEnum = "ENTERPRISE_EDITION_DEVELOPER"
)

var mappingBackupSummaryDatabaseEditionEnum = map[string]BackupSummaryDatabaseEditionEnum{
	"STANDARD_EDITION":                       BackupSummaryDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     BackupSummaryDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    BackupSummaryDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": BackupSummaryDatabaseEditionEnterpriseEditionExtremePerformance,
	"ENTERPRISE_EDITION_DEVELOPER":           BackupSummaryDatabaseEditionEnterpriseEditionDeveloper,
}

var mappingBackupSummaryDatabaseEditionEnumLowerCase = map[string]BackupSummaryDatabaseEditionEnum{
	"standard_edition":                       BackupSummaryDatabaseEditionStandardEdition,
	"enterprise_edition":                     BackupSummaryDatabaseEditionEnterpriseEdition,
	"enterprise_edition_high_performance":    BackupSummaryDatabaseEditionEnterpriseEditionHighPerformance,
	"enterprise_edition_extreme_performance": BackupSummaryDatabaseEditionEnterpriseEditionExtremePerformance,
	"enterprise_edition_developer":           BackupSummaryDatabaseEditionEnterpriseEditionDeveloper,
}

// GetBackupSummaryDatabaseEditionEnumValues Enumerates the set of values for BackupSummaryDatabaseEditionEnum
func GetBackupSummaryDatabaseEditionEnumValues() []BackupSummaryDatabaseEditionEnum {
	values := make([]BackupSummaryDatabaseEditionEnum, 0)
	for _, v := range mappingBackupSummaryDatabaseEditionEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupSummaryDatabaseEditionEnumStringValues Enumerates the set of values in String for BackupSummaryDatabaseEditionEnum
func GetBackupSummaryDatabaseEditionEnumStringValues() []string {
	return []string{
		"STANDARD_EDITION",
		"ENTERPRISE_EDITION",
		"ENTERPRISE_EDITION_HIGH_PERFORMANCE",
		"ENTERPRISE_EDITION_EXTREME_PERFORMANCE",
		"ENTERPRISE_EDITION_DEVELOPER",
	}
}

// GetMappingBackupSummaryDatabaseEditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupSummaryDatabaseEditionEnum(val string) (BackupSummaryDatabaseEditionEnum, bool) {
	enum, ok := mappingBackupSummaryDatabaseEditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BackupSummaryBackupDestinationTypeEnum Enum with underlying type: string
type BackupSummaryBackupDestinationTypeEnum string

// Set of constants representing the allowable values for BackupSummaryBackupDestinationTypeEnum
const (
	BackupSummaryBackupDestinationTypeObjectStore BackupSummaryBackupDestinationTypeEnum = "OBJECT_STORE"
	BackupSummaryBackupDestinationTypeDbrs        BackupSummaryBackupDestinationTypeEnum = "DBRS"
	BackupSummaryBackupDestinationTypeAwsS3       BackupSummaryBackupDestinationTypeEnum = "AWS_S3"
)

var mappingBackupSummaryBackupDestinationTypeEnum = map[string]BackupSummaryBackupDestinationTypeEnum{
	"OBJECT_STORE": BackupSummaryBackupDestinationTypeObjectStore,
	"DBRS":         BackupSummaryBackupDestinationTypeDbrs,
	"AWS_S3":       BackupSummaryBackupDestinationTypeAwsS3,
}

var mappingBackupSummaryBackupDestinationTypeEnumLowerCase = map[string]BackupSummaryBackupDestinationTypeEnum{
	"object_store": BackupSummaryBackupDestinationTypeObjectStore,
	"dbrs":         BackupSummaryBackupDestinationTypeDbrs,
	"aws_s3":       BackupSummaryBackupDestinationTypeAwsS3,
}

// GetBackupSummaryBackupDestinationTypeEnumValues Enumerates the set of values for BackupSummaryBackupDestinationTypeEnum
func GetBackupSummaryBackupDestinationTypeEnumValues() []BackupSummaryBackupDestinationTypeEnum {
	values := make([]BackupSummaryBackupDestinationTypeEnum, 0)
	for _, v := range mappingBackupSummaryBackupDestinationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupSummaryBackupDestinationTypeEnumStringValues Enumerates the set of values in String for BackupSummaryBackupDestinationTypeEnum
func GetBackupSummaryBackupDestinationTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORE",
		"DBRS",
		"AWS_S3",
	}
}

// GetMappingBackupSummaryBackupDestinationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupSummaryBackupDestinationTypeEnum(val string) (BackupSummaryBackupDestinationTypeEnum, bool) {
	enum, ok := mappingBackupSummaryBackupDestinationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
