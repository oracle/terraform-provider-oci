// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DistributedAutonomousDbBackupDestination Backup destination details
type DistributedAutonomousDbBackupDestination struct {

	// Type of the database backup destination.
	Type DistributedAutonomousDbBackupDestinationTypeEnum `mandatory:"true" json:"type"`

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

	// Indicates whether the backup destination is cross-region or local region.
	IsRemote *bool `mandatory:"false" json:"isRemote"`

	// The name of the remote region where the remote automatic incremental backups will be stored.
	// For information about valid region names, see
	// Regions and Availability Domains (https://docs.oracle.com/iaas/Content/General/Concepts/regions.htm).
	RemoteRegion *string `mandatory:"false" json:"remoteRegion"`
}

func (m DistributedAutonomousDbBackupDestination) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedAutonomousDbBackupDestination) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedAutonomousDbBackupDestinationTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDistributedAutonomousDbBackupDestinationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DistributedAutonomousDbBackupDestinationTypeEnum Enum with underlying type: string
type DistributedAutonomousDbBackupDestinationTypeEnum string

// Set of constants representing the allowable values for DistributedAutonomousDbBackupDestinationTypeEnum
const (
	DistributedAutonomousDbBackupDestinationTypeNfs               DistributedAutonomousDbBackupDestinationTypeEnum = "NFS"
	DistributedAutonomousDbBackupDestinationTypeRecoveryAppliance DistributedAutonomousDbBackupDestinationTypeEnum = "RECOVERY_APPLIANCE"
	DistributedAutonomousDbBackupDestinationTypeObjectStore       DistributedAutonomousDbBackupDestinationTypeEnum = "OBJECT_STORE"
	DistributedAutonomousDbBackupDestinationTypeLocal             DistributedAutonomousDbBackupDestinationTypeEnum = "LOCAL"
	DistributedAutonomousDbBackupDestinationTypeDbrs              DistributedAutonomousDbBackupDestinationTypeEnum = "DBRS"
)

var mappingDistributedAutonomousDbBackupDestinationTypeEnum = map[string]DistributedAutonomousDbBackupDestinationTypeEnum{
	"NFS":                DistributedAutonomousDbBackupDestinationTypeNfs,
	"RECOVERY_APPLIANCE": DistributedAutonomousDbBackupDestinationTypeRecoveryAppliance,
	"OBJECT_STORE":       DistributedAutonomousDbBackupDestinationTypeObjectStore,
	"LOCAL":              DistributedAutonomousDbBackupDestinationTypeLocal,
	"DBRS":               DistributedAutonomousDbBackupDestinationTypeDbrs,
}

var mappingDistributedAutonomousDbBackupDestinationTypeEnumLowerCase = map[string]DistributedAutonomousDbBackupDestinationTypeEnum{
	"nfs":                DistributedAutonomousDbBackupDestinationTypeNfs,
	"recovery_appliance": DistributedAutonomousDbBackupDestinationTypeRecoveryAppliance,
	"object_store":       DistributedAutonomousDbBackupDestinationTypeObjectStore,
	"local":              DistributedAutonomousDbBackupDestinationTypeLocal,
	"dbrs":               DistributedAutonomousDbBackupDestinationTypeDbrs,
}

// GetDistributedAutonomousDbBackupDestinationTypeEnumValues Enumerates the set of values for DistributedAutonomousDbBackupDestinationTypeEnum
func GetDistributedAutonomousDbBackupDestinationTypeEnumValues() []DistributedAutonomousDbBackupDestinationTypeEnum {
	values := make([]DistributedAutonomousDbBackupDestinationTypeEnum, 0)
	for _, v := range mappingDistributedAutonomousDbBackupDestinationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedAutonomousDbBackupDestinationTypeEnumStringValues Enumerates the set of values in String for DistributedAutonomousDbBackupDestinationTypeEnum
func GetDistributedAutonomousDbBackupDestinationTypeEnumStringValues() []string {
	return []string{
		"NFS",
		"RECOVERY_APPLIANCE",
		"OBJECT_STORE",
		"LOCAL",
		"DBRS",
	}
}

// GetMappingDistributedAutonomousDbBackupDestinationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedAutonomousDbBackupDestinationTypeEnum(val string) (DistributedAutonomousDbBackupDestinationTypeEnum, bool) {
	enum, ok := mappingDistributedAutonomousDbBackupDestinationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
