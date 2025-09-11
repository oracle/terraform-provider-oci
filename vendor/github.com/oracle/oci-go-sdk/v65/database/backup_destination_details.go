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

// BackupDestinationDetails Backup destination details
type BackupDestinationDetails struct {

	// Type of the database backup destination.
	Type BackupDestinationDetailsTypeEnum `mandatory:"true" json:"type"`

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

	// Defines the automatic and manual backup retention policy for the Autonomous Database termination.
	// The retention policy set on the Autonomous Container Database is not applicable for cross region remote backups and backups hosted on recovery Appliance backup destination.
	// Options are 'RETAIN_PER_RETENTION_WINDOW' or 'RETAIN_FOR_72_HOURS'.The default value is 'RETAIN_FOR_72_HOURS'.
	BackupRetentionPolicyOnTerminate BackupDestinationDetailsBackupRetentionPolicyOnTerminateEnum `mandatory:"false" json:"backupRetentionPolicyOnTerminate,omitempty"`

	// Indicates whether the backup destination is cross-region or local.
	IsRemote *bool `mandatory:"false" json:"isRemote"`

	// The name of the remote region where the remote automatic incremental backups will be stored.
	// For information about valid region names, see
	// Regions and Availability Domains (https://docs.oracle.com/iaas/Content/General/Concepts/regions.htm).
	RemoteRegion *string `mandatory:"false" json:"remoteRegion"`
}

func (m BackupDestinationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackupDestinationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBackupDestinationDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetBackupDestinationDetailsTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBackupDestinationDetailsBackupRetentionPolicyOnTerminateEnum(string(m.BackupRetentionPolicyOnTerminate)); !ok && m.BackupRetentionPolicyOnTerminate != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupRetentionPolicyOnTerminate: %s. Supported values are: %s.", m.BackupRetentionPolicyOnTerminate, strings.Join(GetBackupDestinationDetailsBackupRetentionPolicyOnTerminateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BackupDestinationDetailsTypeEnum Enum with underlying type: string
type BackupDestinationDetailsTypeEnum string

// Set of constants representing the allowable values for BackupDestinationDetailsTypeEnum
const (
	BackupDestinationDetailsTypeNfs               BackupDestinationDetailsTypeEnum = "NFS"
	BackupDestinationDetailsTypeRecoveryAppliance BackupDestinationDetailsTypeEnum = "RECOVERY_APPLIANCE"
	BackupDestinationDetailsTypeObjectStore       BackupDestinationDetailsTypeEnum = "OBJECT_STORE"
	BackupDestinationDetailsTypeLocal             BackupDestinationDetailsTypeEnum = "LOCAL"
	BackupDestinationDetailsTypeDbrs              BackupDestinationDetailsTypeEnum = "DBRS"
	BackupDestinationDetailsTypeAwsS3             BackupDestinationDetailsTypeEnum = "AWS_S3"
)

var mappingBackupDestinationDetailsTypeEnum = map[string]BackupDestinationDetailsTypeEnum{
	"NFS":                BackupDestinationDetailsTypeNfs,
	"RECOVERY_APPLIANCE": BackupDestinationDetailsTypeRecoveryAppliance,
	"OBJECT_STORE":       BackupDestinationDetailsTypeObjectStore,
	"LOCAL":              BackupDestinationDetailsTypeLocal,
	"DBRS":               BackupDestinationDetailsTypeDbrs,
	"AWS_S3":             BackupDestinationDetailsTypeAwsS3,
}

var mappingBackupDestinationDetailsTypeEnumLowerCase = map[string]BackupDestinationDetailsTypeEnum{
	"nfs":                BackupDestinationDetailsTypeNfs,
	"recovery_appliance": BackupDestinationDetailsTypeRecoveryAppliance,
	"object_store":       BackupDestinationDetailsTypeObjectStore,
	"local":              BackupDestinationDetailsTypeLocal,
	"dbrs":               BackupDestinationDetailsTypeDbrs,
	"aws_s3":             BackupDestinationDetailsTypeAwsS3,
}

// GetBackupDestinationDetailsTypeEnumValues Enumerates the set of values for BackupDestinationDetailsTypeEnum
func GetBackupDestinationDetailsTypeEnumValues() []BackupDestinationDetailsTypeEnum {
	values := make([]BackupDestinationDetailsTypeEnum, 0)
	for _, v := range mappingBackupDestinationDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupDestinationDetailsTypeEnumStringValues Enumerates the set of values in String for BackupDestinationDetailsTypeEnum
func GetBackupDestinationDetailsTypeEnumStringValues() []string {
	return []string{
		"NFS",
		"RECOVERY_APPLIANCE",
		"OBJECT_STORE",
		"LOCAL",
		"DBRS",
		"AWS_S3",
	}
}

// GetMappingBackupDestinationDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupDestinationDetailsTypeEnum(val string) (BackupDestinationDetailsTypeEnum, bool) {
	enum, ok := mappingBackupDestinationDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BackupDestinationDetailsBackupRetentionPolicyOnTerminateEnum Enum with underlying type: string
type BackupDestinationDetailsBackupRetentionPolicyOnTerminateEnum string

// Set of constants representing the allowable values for BackupDestinationDetailsBackupRetentionPolicyOnTerminateEnum
const (
	BackupDestinationDetailsBackupRetentionPolicyOnTerminatePerRetentionWindow BackupDestinationDetailsBackupRetentionPolicyOnTerminateEnum = "RETAIN_PER_RETENTION_WINDOW"
	BackupDestinationDetailsBackupRetentionPolicyOnTerminateFor72Hours         BackupDestinationDetailsBackupRetentionPolicyOnTerminateEnum = "RETAIN_FOR_72_HOURS"
)

var mappingBackupDestinationDetailsBackupRetentionPolicyOnTerminateEnum = map[string]BackupDestinationDetailsBackupRetentionPolicyOnTerminateEnum{
	"RETAIN_PER_RETENTION_WINDOW": BackupDestinationDetailsBackupRetentionPolicyOnTerminatePerRetentionWindow,
	"RETAIN_FOR_72_HOURS":         BackupDestinationDetailsBackupRetentionPolicyOnTerminateFor72Hours,
}

var mappingBackupDestinationDetailsBackupRetentionPolicyOnTerminateEnumLowerCase = map[string]BackupDestinationDetailsBackupRetentionPolicyOnTerminateEnum{
	"retain_per_retention_window": BackupDestinationDetailsBackupRetentionPolicyOnTerminatePerRetentionWindow,
	"retain_for_72_hours":         BackupDestinationDetailsBackupRetentionPolicyOnTerminateFor72Hours,
}

// GetBackupDestinationDetailsBackupRetentionPolicyOnTerminateEnumValues Enumerates the set of values for BackupDestinationDetailsBackupRetentionPolicyOnTerminateEnum
func GetBackupDestinationDetailsBackupRetentionPolicyOnTerminateEnumValues() []BackupDestinationDetailsBackupRetentionPolicyOnTerminateEnum {
	values := make([]BackupDestinationDetailsBackupRetentionPolicyOnTerminateEnum, 0)
	for _, v := range mappingBackupDestinationDetailsBackupRetentionPolicyOnTerminateEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupDestinationDetailsBackupRetentionPolicyOnTerminateEnumStringValues Enumerates the set of values in String for BackupDestinationDetailsBackupRetentionPolicyOnTerminateEnum
func GetBackupDestinationDetailsBackupRetentionPolicyOnTerminateEnumStringValues() []string {
	return []string{
		"RETAIN_PER_RETENTION_WINDOW",
		"RETAIN_FOR_72_HOURS",
	}
}

// GetMappingBackupDestinationDetailsBackupRetentionPolicyOnTerminateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupDestinationDetailsBackupRetentionPolicyOnTerminateEnum(val string) (BackupDestinationDetailsBackupRetentionPolicyOnTerminateEnum, bool) {
	enum, ok := mappingBackupDestinationDetailsBackupRetentionPolicyOnTerminateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
