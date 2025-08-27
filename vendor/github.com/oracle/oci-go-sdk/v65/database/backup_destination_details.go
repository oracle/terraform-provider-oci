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
