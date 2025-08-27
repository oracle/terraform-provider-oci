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

// BackupDestinationConfigurationSummary Information about the Autonomous Container Database's secondary backup destination(s).
type BackupDestinationConfigurationSummary struct {

	// Type of the database backup destination.
	Type BackupDestinationConfigurationSummaryTypeEnum `mandatory:"true" json:"type"`

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

	// Number of days between the current and earliest point of recoverability covered by automatic backups and manual backups, but not long term backups.
	RecoveryWindowInDays *int `mandatory:"false" json:"recoveryWindowInDays"`
}

func (m BackupDestinationConfigurationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackupDestinationConfigurationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBackupDestinationConfigurationSummaryTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetBackupDestinationConfigurationSummaryTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BackupDestinationConfigurationSummaryTypeEnum Enum with underlying type: string
type BackupDestinationConfigurationSummaryTypeEnum string

// Set of constants representing the allowable values for BackupDestinationConfigurationSummaryTypeEnum
const (
	BackupDestinationConfigurationSummaryTypeNfs               BackupDestinationConfigurationSummaryTypeEnum = "NFS"
	BackupDestinationConfigurationSummaryTypeRecoveryAppliance BackupDestinationConfigurationSummaryTypeEnum = "RECOVERY_APPLIANCE"
	BackupDestinationConfigurationSummaryTypeObjectStore       BackupDestinationConfigurationSummaryTypeEnum = "OBJECT_STORE"
	BackupDestinationConfigurationSummaryTypeLocal             BackupDestinationConfigurationSummaryTypeEnum = "LOCAL"
	BackupDestinationConfigurationSummaryTypeDbrs              BackupDestinationConfigurationSummaryTypeEnum = "DBRS"
	BackupDestinationConfigurationSummaryTypeAwsS3             BackupDestinationConfigurationSummaryTypeEnum = "AWS_S3"
)

var mappingBackupDestinationConfigurationSummaryTypeEnum = map[string]BackupDestinationConfigurationSummaryTypeEnum{
	"NFS":                BackupDestinationConfigurationSummaryTypeNfs,
	"RECOVERY_APPLIANCE": BackupDestinationConfigurationSummaryTypeRecoveryAppliance,
	"OBJECT_STORE":       BackupDestinationConfigurationSummaryTypeObjectStore,
	"LOCAL":              BackupDestinationConfigurationSummaryTypeLocal,
	"DBRS":               BackupDestinationConfigurationSummaryTypeDbrs,
	"AWS_S3":             BackupDestinationConfigurationSummaryTypeAwsS3,
}

var mappingBackupDestinationConfigurationSummaryTypeEnumLowerCase = map[string]BackupDestinationConfigurationSummaryTypeEnum{
	"nfs":                BackupDestinationConfigurationSummaryTypeNfs,
	"recovery_appliance": BackupDestinationConfigurationSummaryTypeRecoveryAppliance,
	"object_store":       BackupDestinationConfigurationSummaryTypeObjectStore,
	"local":              BackupDestinationConfigurationSummaryTypeLocal,
	"dbrs":               BackupDestinationConfigurationSummaryTypeDbrs,
	"aws_s3":             BackupDestinationConfigurationSummaryTypeAwsS3,
}

// GetBackupDestinationConfigurationSummaryTypeEnumValues Enumerates the set of values for BackupDestinationConfigurationSummaryTypeEnum
func GetBackupDestinationConfigurationSummaryTypeEnumValues() []BackupDestinationConfigurationSummaryTypeEnum {
	values := make([]BackupDestinationConfigurationSummaryTypeEnum, 0)
	for _, v := range mappingBackupDestinationConfigurationSummaryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupDestinationConfigurationSummaryTypeEnumStringValues Enumerates the set of values in String for BackupDestinationConfigurationSummaryTypeEnum
func GetBackupDestinationConfigurationSummaryTypeEnumStringValues() []string {
	return []string{
		"NFS",
		"RECOVERY_APPLIANCE",
		"OBJECT_STORE",
		"LOCAL",
		"DBRS",
		"AWS_S3",
	}
}

// GetMappingBackupDestinationConfigurationSummaryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupDestinationConfigurationSummaryTypeEnum(val string) (BackupDestinationConfigurationSummaryTypeEnum, bool) {
	enum, ok := mappingBackupDestinationConfigurationSummaryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
