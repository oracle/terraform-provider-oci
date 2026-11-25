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

// DistributedDbBackupDestination Backup destination details
type DistributedDbBackupDestination struct {

	// Type of the database backup destination.
	Type DistributedDbBackupDestinationTypeEnum `mandatory:"true" json:"type"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
	Id *string `mandatory:"false" json:"id"`

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

func (m DistributedDbBackupDestination) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedDbBackupDestination) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedDbBackupDestinationTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDistributedDbBackupDestinationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DistributedDbBackupDestinationTypeEnum Enum with underlying type: string
type DistributedDbBackupDestinationTypeEnum string

// Set of constants representing the allowable values for DistributedDbBackupDestinationTypeEnum
const (
	DistributedDbBackupDestinationTypeNfs         DistributedDbBackupDestinationTypeEnum = "NFS"
	DistributedDbBackupDestinationTypeObjectStore DistributedDbBackupDestinationTypeEnum = "OBJECT_STORE"
	DistributedDbBackupDestinationTypeLocal       DistributedDbBackupDestinationTypeEnum = "LOCAL"
	DistributedDbBackupDestinationTypeDbrs        DistributedDbBackupDestinationTypeEnum = "DBRS"
)

var mappingDistributedDbBackupDestinationTypeEnum = map[string]DistributedDbBackupDestinationTypeEnum{
	"NFS":          DistributedDbBackupDestinationTypeNfs,
	"OBJECT_STORE": DistributedDbBackupDestinationTypeObjectStore,
	"LOCAL":        DistributedDbBackupDestinationTypeLocal,
	"DBRS":         DistributedDbBackupDestinationTypeDbrs,
}

var mappingDistributedDbBackupDestinationTypeEnumLowerCase = map[string]DistributedDbBackupDestinationTypeEnum{
	"nfs":          DistributedDbBackupDestinationTypeNfs,
	"object_store": DistributedDbBackupDestinationTypeObjectStore,
	"local":        DistributedDbBackupDestinationTypeLocal,
	"dbrs":         DistributedDbBackupDestinationTypeDbrs,
}

// GetDistributedDbBackupDestinationTypeEnumValues Enumerates the set of values for DistributedDbBackupDestinationTypeEnum
func GetDistributedDbBackupDestinationTypeEnumValues() []DistributedDbBackupDestinationTypeEnum {
	values := make([]DistributedDbBackupDestinationTypeEnum, 0)
	for _, v := range mappingDistributedDbBackupDestinationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDbBackupDestinationTypeEnumStringValues Enumerates the set of values in String for DistributedDbBackupDestinationTypeEnum
func GetDistributedDbBackupDestinationTypeEnumStringValues() []string {
	return []string{
		"NFS",
		"OBJECT_STORE",
		"LOCAL",
		"DBRS",
	}
}

// GetMappingDistributedDbBackupDestinationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDbBackupDestinationTypeEnum(val string) (DistributedDbBackupDestinationTypeEnum, bool) {
	enum, ok := mappingDistributedDbBackupDestinationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
