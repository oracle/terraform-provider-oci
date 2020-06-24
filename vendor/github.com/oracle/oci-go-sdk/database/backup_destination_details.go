// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// BackupDestinationDetails Backup destination details
type BackupDestinationDetails struct {

	// Type of the database backup destination.
	Type BackupDestinationDetailsTypeEnum `mandatory:"true" json:"type"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the backup destination.
	Id *string `mandatory:"false" json:"id"`

	// For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) user that is used to access the Recovery Appliance.
	VpcUser *string `mandatory:"false" json:"vpcUser"`

	// For a RECOVERY_APPLIANCE backup destination, the password for the VPC user that is used to access the Recovery Appliance.
	VpcPassword *string `mandatory:"false" json:"vpcPassword"`

	// Proxy URL to connect to object store.
	InternetProxy *string `mandatory:"false" json:"internetProxy"`
}

func (m BackupDestinationDetails) String() string {
	return common.PointerString(m)
}

// BackupDestinationDetailsTypeEnum Enum with underlying type: string
type BackupDestinationDetailsTypeEnum string

// Set of constants representing the allowable values for BackupDestinationDetailsTypeEnum
const (
	BackupDestinationDetailsTypeNfs               BackupDestinationDetailsTypeEnum = "NFS"
	BackupDestinationDetailsTypeRecoveryAppliance BackupDestinationDetailsTypeEnum = "RECOVERY_APPLIANCE"
	BackupDestinationDetailsTypeObjectStore       BackupDestinationDetailsTypeEnum = "OBJECT_STORE"
	BackupDestinationDetailsTypeLocal             BackupDestinationDetailsTypeEnum = "LOCAL"
)

var mappingBackupDestinationDetailsType = map[string]BackupDestinationDetailsTypeEnum{
	"NFS":                BackupDestinationDetailsTypeNfs,
	"RECOVERY_APPLIANCE": BackupDestinationDetailsTypeRecoveryAppliance,
	"OBJECT_STORE":       BackupDestinationDetailsTypeObjectStore,
	"LOCAL":              BackupDestinationDetailsTypeLocal,
}

// GetBackupDestinationDetailsTypeEnumValues Enumerates the set of values for BackupDestinationDetailsTypeEnum
func GetBackupDestinationDetailsTypeEnumValues() []BackupDestinationDetailsTypeEnum {
	values := make([]BackupDestinationDetailsTypeEnum, 0)
	for _, v := range mappingBackupDestinationDetailsType {
		values = append(values, v)
	}
	return values
}
