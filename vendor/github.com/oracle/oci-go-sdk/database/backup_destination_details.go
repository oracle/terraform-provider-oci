// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
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

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
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
