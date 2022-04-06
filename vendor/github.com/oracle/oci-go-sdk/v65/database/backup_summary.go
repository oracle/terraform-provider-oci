// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// BackupSummary A database backup.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type BackupSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the backup.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database.
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

	// The Oracle Database edition of the DB system from which the database backup was taken.
	DatabaseEdition BackupSummaryDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`

	// The size of the database in gigabytes at the time the backup was taken.
	DatabaseSizeInGBs *float64 `mandatory:"false" json:"databaseSizeInGBs"`

	// Shape of the backup's source database.
	Shape *string `mandatory:"false" json:"shape"`

	// Version of the backup's source database
	Version *string `mandatory:"false" json:"version"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.cloud.oracle.com/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
	VaultId *string `mandatory:"false" json:"vaultId"`
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
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
)

var mappingBackupSummaryLifecycleStateEnum = map[string]BackupSummaryLifecycleStateEnum{
	"CREATING":  BackupSummaryLifecycleStateCreating,
	"ACTIVE":    BackupSummaryLifecycleStateActive,
	"DELETING":  BackupSummaryLifecycleStateDeleting,
	"DELETED":   BackupSummaryLifecycleStateDeleted,
	"FAILED":    BackupSummaryLifecycleStateFailed,
	"RESTORING": BackupSummaryLifecycleStateRestoring,
}

var mappingBackupSummaryLifecycleStateEnumLowerCase = map[string]BackupSummaryLifecycleStateEnum{
	"creating":  BackupSummaryLifecycleStateCreating,
	"active":    BackupSummaryLifecycleStateActive,
	"deleting":  BackupSummaryLifecycleStateDeleting,
	"deleted":   BackupSummaryLifecycleStateDeleted,
	"failed":    BackupSummaryLifecycleStateFailed,
	"restoring": BackupSummaryLifecycleStateRestoring,
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
)

var mappingBackupSummaryDatabaseEditionEnum = map[string]BackupSummaryDatabaseEditionEnum{
	"STANDARD_EDITION":                       BackupSummaryDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     BackupSummaryDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    BackupSummaryDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": BackupSummaryDatabaseEditionEnterpriseEditionExtremePerformance,
}

var mappingBackupSummaryDatabaseEditionEnumLowerCase = map[string]BackupSummaryDatabaseEditionEnum{
	"standard_edition":                       BackupSummaryDatabaseEditionStandardEdition,
	"enterprise_edition":                     BackupSummaryDatabaseEditionEnterpriseEdition,
	"enterprise_edition_high_performance":    BackupSummaryDatabaseEditionEnterpriseEditionHighPerformance,
	"enterprise_edition_extreme_performance": BackupSummaryDatabaseEditionEnterpriseEditionExtremePerformance,
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
	}
}

// GetMappingBackupSummaryDatabaseEditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupSummaryDatabaseEditionEnum(val string) (BackupSummaryDatabaseEditionEnum, bool) {
	enum, ok := mappingBackupSummaryDatabaseEditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
