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

// BackupDestinationDetailsSummary Information about the backup destination associated with the Autonomous Container Database.
type BackupDestinationDetailsSummary struct {

	// Type of the database backup destination.
	Type BackupDestinationDetailsSummaryTypeEnum `mandatory:"true" json:"type"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
	Id *string `mandatory:"false" json:"id"`

	// For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) user that is used to access the Recovery Appliance.
	VpcUser *string `mandatory:"false" json:"vpcUser"`

	// For a RECOVERY_APPLIANCE backup destination, the password for the VPC user that is used to access the Recovery Appliance.
	VpcPassword *string `mandatory:"false" json:"vpcPassword"`

	// Proxy URL to connect to object store.
	InternetProxy *string `mandatory:"false" json:"internetProxy"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DBRS policy used for backup.
	DbrsPolicyId *string `mandatory:"false" json:"dbrsPolicyId"`

	// Indicates if backup retention is locked for all the database backups in the Autonomous Container Database (ACD). The retention window cannot be decreased if the backup retention lock is enabled.
	// Once applied on the Autonomous Container Database, the retention lock cannot be removed, or the retention period cannot be decreased after a 14-day period.
	// If the backup is a Long Term Backup and retention lock is enabled, the backup cannot be deleted and must expire.
	// The retention lock set on the Autonomous Container Database is not applicable for cross region remote backups and backups hosted on recovery Appliance backup destination.
	IsRetentionLockEnabled *bool `mandatory:"false" json:"isRetentionLockEnabled"`

	// Defines the automatic and manual backup retention policy for the Autonomous AI Database termination.
	// The retention policy set on the Autonomous Container Database is not applicable for cross region remote backups and backups hosted on recovery Appliance backup destination.
	// Options are 'RETAIN_PER_RETENTION_WINDOW' or 'RETAIN_FOR_72_HOURS'.The default value is 'RETAIN_FOR_72_HOURS'.
	BackupRetentionPolicyOnTerminate BackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnum `mandatory:"false" json:"backupRetentionPolicyOnTerminate,omitempty"`

	// Indicates whether the backup destination is cross-region or local.
	IsRemote *bool `mandatory:"false" json:"isRemote"`

	// The name of the remote region where the remote automatic incremental backups will be stored.
	// For information about valid region names, see
	// Regions and Availability Domains (https://docs.oracle.com/iaas/Content/General/Concepts/regions.htm).
	RemoteRegion *string `mandatory:"false" json:"remoteRegion"`

	// The timestamps at which this backup destination is used as the preferred destination to host the Autonomous Container Database backups.
	BackupDestinationAttachHistory []common.SDKTime `mandatory:"false" json:"backupDestinationAttachHistory"`

	// The total space utilized (in GBs) by this Autonomous Container Database on this backup destination, rounded to the nearest integer.
	SpaceUtilizedInGBs *int `mandatory:"false" json:"spaceUtilizedInGBs"`

	// The latest timestamp when the backup destination details, such as 'spaceUtilized,' are updated.
	TimeAtWhichStorageDetailsAreUpdated *common.SDKTime `mandatory:"false" json:"timeAtWhichStorageDetailsAreUpdated"`
}

func (m BackupDestinationDetailsSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackupDestinationDetailsSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBackupDestinationDetailsSummaryTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetBackupDestinationDetailsSummaryTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnum(string(m.BackupRetentionPolicyOnTerminate)); !ok && m.BackupRetentionPolicyOnTerminate != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupRetentionPolicyOnTerminate: %s. Supported values are: %s.", m.BackupRetentionPolicyOnTerminate, strings.Join(GetBackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BackupDestinationDetailsSummaryTypeEnum Enum with underlying type: string
type BackupDestinationDetailsSummaryTypeEnum string

// Set of constants representing the allowable values for BackupDestinationDetailsSummaryTypeEnum
const (
	BackupDestinationDetailsSummaryTypeNfs               BackupDestinationDetailsSummaryTypeEnum = "NFS"
	BackupDestinationDetailsSummaryTypeRecoveryAppliance BackupDestinationDetailsSummaryTypeEnum = "RECOVERY_APPLIANCE"
	BackupDestinationDetailsSummaryTypeObjectStore       BackupDestinationDetailsSummaryTypeEnum = "OBJECT_STORE"
	BackupDestinationDetailsSummaryTypeLocal             BackupDestinationDetailsSummaryTypeEnum = "LOCAL"
	BackupDestinationDetailsSummaryTypeDbrs              BackupDestinationDetailsSummaryTypeEnum = "DBRS"
	BackupDestinationDetailsSummaryTypeAwsS3             BackupDestinationDetailsSummaryTypeEnum = "AWS_S3"
)

var mappingBackupDestinationDetailsSummaryTypeEnum = map[string]BackupDestinationDetailsSummaryTypeEnum{
	"NFS":                BackupDestinationDetailsSummaryTypeNfs,
	"RECOVERY_APPLIANCE": BackupDestinationDetailsSummaryTypeRecoveryAppliance,
	"OBJECT_STORE":       BackupDestinationDetailsSummaryTypeObjectStore,
	"LOCAL":              BackupDestinationDetailsSummaryTypeLocal,
	"DBRS":               BackupDestinationDetailsSummaryTypeDbrs,
	"AWS_S3":             BackupDestinationDetailsSummaryTypeAwsS3,
}

var mappingBackupDestinationDetailsSummaryTypeEnumLowerCase = map[string]BackupDestinationDetailsSummaryTypeEnum{
	"nfs":                BackupDestinationDetailsSummaryTypeNfs,
	"recovery_appliance": BackupDestinationDetailsSummaryTypeRecoveryAppliance,
	"object_store":       BackupDestinationDetailsSummaryTypeObjectStore,
	"local":              BackupDestinationDetailsSummaryTypeLocal,
	"dbrs":               BackupDestinationDetailsSummaryTypeDbrs,
	"aws_s3":             BackupDestinationDetailsSummaryTypeAwsS3,
}

// GetBackupDestinationDetailsSummaryTypeEnumValues Enumerates the set of values for BackupDestinationDetailsSummaryTypeEnum
func GetBackupDestinationDetailsSummaryTypeEnumValues() []BackupDestinationDetailsSummaryTypeEnum {
	values := make([]BackupDestinationDetailsSummaryTypeEnum, 0)
	for _, v := range mappingBackupDestinationDetailsSummaryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupDestinationDetailsSummaryTypeEnumStringValues Enumerates the set of values in String for BackupDestinationDetailsSummaryTypeEnum
func GetBackupDestinationDetailsSummaryTypeEnumStringValues() []string {
	return []string{
		"NFS",
		"RECOVERY_APPLIANCE",
		"OBJECT_STORE",
		"LOCAL",
		"DBRS",
		"AWS_S3",
	}
}

// GetMappingBackupDestinationDetailsSummaryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupDestinationDetailsSummaryTypeEnum(val string) (BackupDestinationDetailsSummaryTypeEnum, bool) {
	enum, ok := mappingBackupDestinationDetailsSummaryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnum Enum with underlying type: string
type BackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnum string

// Set of constants representing the allowable values for BackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnum
const (
	BackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminatePerRetentionWindow BackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnum = "RETAIN_PER_RETENTION_WINDOW"
	BackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateFor72Hours         BackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnum = "RETAIN_FOR_72_HOURS"
)

var mappingBackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnum = map[string]BackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnum{
	"RETAIN_PER_RETENTION_WINDOW": BackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminatePerRetentionWindow,
	"RETAIN_FOR_72_HOURS":         BackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateFor72Hours,
}

var mappingBackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnumLowerCase = map[string]BackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnum{
	"retain_per_retention_window": BackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminatePerRetentionWindow,
	"retain_for_72_hours":         BackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateFor72Hours,
}

// GetBackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnumValues Enumerates the set of values for BackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnum
func GetBackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnumValues() []BackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnum {
	values := make([]BackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnum, 0)
	for _, v := range mappingBackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnumStringValues Enumerates the set of values in String for BackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnum
func GetBackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnumStringValues() []string {
	return []string{
		"RETAIN_PER_RETENTION_WINDOW",
		"RETAIN_FOR_72_HOURS",
	}
}

// GetMappingBackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnum(val string) (BackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnum, bool) {
	enum, ok := mappingBackupDestinationDetailsSummaryBackupRetentionPolicyOnTerminateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
