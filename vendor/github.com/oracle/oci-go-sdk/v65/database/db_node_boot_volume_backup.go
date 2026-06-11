// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// DbNodeBootVolumeBackup Details of database node boot volume backup.
type DbNodeBootVolumeBackup struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database node boot volume backup.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the database node.
	DbNodeId *string `mandatory:"true" json:"dbNodeId"`

	// A user-friendly name for the database node boot volume backup. The name does not need to be unique.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of a database node boot volume backup.
	LifecycleState DbNodeBootVolumeBackupLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The size used by the database node boot volume backup, in GBs.
	SizeInGBs *int `mandatory:"false" json:"sizeInGBs"`

	// Additional information about the current lifecycle state of the database node boot volume backup.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time when the database node boot volume backup was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`
}

func (m DbNodeBootVolumeBackup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbNodeBootVolumeBackup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDbNodeBootVolumeBackupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbNodeBootVolumeBackupLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbNodeBootVolumeBackupLifecycleStateEnum Enum with underlying type: string
type DbNodeBootVolumeBackupLifecycleStateEnum string

// Set of constants representing the allowable values for DbNodeBootVolumeBackupLifecycleStateEnum
const (
	DbNodeBootVolumeBackupLifecycleStateCreating DbNodeBootVolumeBackupLifecycleStateEnum = "CREATING"
	DbNodeBootVolumeBackupLifecycleStateActive   DbNodeBootVolumeBackupLifecycleStateEnum = "ACTIVE"
	DbNodeBootVolumeBackupLifecycleStateDeleting DbNodeBootVolumeBackupLifecycleStateEnum = "DELETING"
	DbNodeBootVolumeBackupLifecycleStateDeleted  DbNodeBootVolumeBackupLifecycleStateEnum = "DELETED"
	DbNodeBootVolumeBackupLifecycleStateFailed   DbNodeBootVolumeBackupLifecycleStateEnum = "FAILED"
)

var mappingDbNodeBootVolumeBackupLifecycleStateEnum = map[string]DbNodeBootVolumeBackupLifecycleStateEnum{
	"CREATING": DbNodeBootVolumeBackupLifecycleStateCreating,
	"ACTIVE":   DbNodeBootVolumeBackupLifecycleStateActive,
	"DELETING": DbNodeBootVolumeBackupLifecycleStateDeleting,
	"DELETED":  DbNodeBootVolumeBackupLifecycleStateDeleted,
	"FAILED":   DbNodeBootVolumeBackupLifecycleStateFailed,
}

var mappingDbNodeBootVolumeBackupLifecycleStateEnumLowerCase = map[string]DbNodeBootVolumeBackupLifecycleStateEnum{
	"creating": DbNodeBootVolumeBackupLifecycleStateCreating,
	"active":   DbNodeBootVolumeBackupLifecycleStateActive,
	"deleting": DbNodeBootVolumeBackupLifecycleStateDeleting,
	"deleted":  DbNodeBootVolumeBackupLifecycleStateDeleted,
	"failed":   DbNodeBootVolumeBackupLifecycleStateFailed,
}

// GetDbNodeBootVolumeBackupLifecycleStateEnumValues Enumerates the set of values for DbNodeBootVolumeBackupLifecycleStateEnum
func GetDbNodeBootVolumeBackupLifecycleStateEnumValues() []DbNodeBootVolumeBackupLifecycleStateEnum {
	values := make([]DbNodeBootVolumeBackupLifecycleStateEnum, 0)
	for _, v := range mappingDbNodeBootVolumeBackupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDbNodeBootVolumeBackupLifecycleStateEnumStringValues Enumerates the set of values in String for DbNodeBootVolumeBackupLifecycleStateEnum
func GetDbNodeBootVolumeBackupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingDbNodeBootVolumeBackupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbNodeBootVolumeBackupLifecycleStateEnum(val string) (DbNodeBootVolumeBackupLifecycleStateEnum, bool) {
	enum, ok := mappingDbNodeBootVolumeBackupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
