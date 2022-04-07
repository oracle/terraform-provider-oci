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

// BackupDestination Backup destination details.
type BackupDestination struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the backup destination.
	Id *string `mandatory:"false" json:"id"`

	// The user-provided name of the backup destination.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Type of the backup destination.
	Type BackupDestinationTypeEnum `mandatory:"false" json:"type,omitempty"`

	// List of databases associated with the backup destination.
	AssociatedDatabases []AssociatedDatabaseDetails `mandatory:"false" json:"associatedDatabases"`

	// For a RECOVERY_APPLIANCE backup destination, the connection string for connecting to the Recovery Appliance.
	ConnectionString *string `mandatory:"false" json:"connectionString"`

	// For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) users that are used to access the Recovery Appliance.
	VpcUsers []string `mandatory:"false" json:"vpcUsers"`

	// The local directory path on each VM cluster node where the NFS server location is mounted. The local directory path and the NFS server location must each be the same across all of the VM cluster nodes. Ensure that the NFS mount is maintained continuously on all of the VM cluster nodes.
	LocalMountPointPath *string `mandatory:"false" json:"localMountPointPath"`

	// NFS Mount type for backup destination.
	NfsMountType BackupDestinationNfsMountTypeEnum `mandatory:"false" json:"nfsMountType,omitempty"`

	// Host names or IP addresses for NFS Auto mount.
	NfsServer []string `mandatory:"false" json:"nfsServer"`

	// Specifies the directory on which to mount the file system
	NfsServerExport *string `mandatory:"false" json:"nfsServerExport"`

	// The current lifecycle state of the backup destination.
	LifecycleState BackupDestinationLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the backup destination was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// A descriptive text associated with the lifecycleState.
	// Typically contains additional displayable text
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m BackupDestination) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackupDestination) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBackupDestinationTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetBackupDestinationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBackupDestinationNfsMountTypeEnum(string(m.NfsMountType)); !ok && m.NfsMountType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NfsMountType: %s. Supported values are: %s.", m.NfsMountType, strings.Join(GetBackupDestinationNfsMountTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBackupDestinationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBackupDestinationLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BackupDestinationTypeEnum Enum with underlying type: string
type BackupDestinationTypeEnum string

// Set of constants representing the allowable values for BackupDestinationTypeEnum
const (
	BackupDestinationTypeNfs               BackupDestinationTypeEnum = "NFS"
	BackupDestinationTypeRecoveryAppliance BackupDestinationTypeEnum = "RECOVERY_APPLIANCE"
)

var mappingBackupDestinationTypeEnum = map[string]BackupDestinationTypeEnum{
	"NFS":                BackupDestinationTypeNfs,
	"RECOVERY_APPLIANCE": BackupDestinationTypeRecoveryAppliance,
}

var mappingBackupDestinationTypeEnumLowerCase = map[string]BackupDestinationTypeEnum{
	"nfs":                BackupDestinationTypeNfs,
	"recovery_appliance": BackupDestinationTypeRecoveryAppliance,
}

// GetBackupDestinationTypeEnumValues Enumerates the set of values for BackupDestinationTypeEnum
func GetBackupDestinationTypeEnumValues() []BackupDestinationTypeEnum {
	values := make([]BackupDestinationTypeEnum, 0)
	for _, v := range mappingBackupDestinationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupDestinationTypeEnumStringValues Enumerates the set of values in String for BackupDestinationTypeEnum
func GetBackupDestinationTypeEnumStringValues() []string {
	return []string{
		"NFS",
		"RECOVERY_APPLIANCE",
	}
}

// GetMappingBackupDestinationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupDestinationTypeEnum(val string) (BackupDestinationTypeEnum, bool) {
	enum, ok := mappingBackupDestinationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BackupDestinationNfsMountTypeEnum Enum with underlying type: string
type BackupDestinationNfsMountTypeEnum string

// Set of constants representing the allowable values for BackupDestinationNfsMountTypeEnum
const (
	BackupDestinationNfsMountTypeSelfMount      BackupDestinationNfsMountTypeEnum = "SELF_MOUNT"
	BackupDestinationNfsMountTypeAutomatedMount BackupDestinationNfsMountTypeEnum = "AUTOMATED_MOUNT"
)

var mappingBackupDestinationNfsMountTypeEnum = map[string]BackupDestinationNfsMountTypeEnum{
	"SELF_MOUNT":      BackupDestinationNfsMountTypeSelfMount,
	"AUTOMATED_MOUNT": BackupDestinationNfsMountTypeAutomatedMount,
}

var mappingBackupDestinationNfsMountTypeEnumLowerCase = map[string]BackupDestinationNfsMountTypeEnum{
	"self_mount":      BackupDestinationNfsMountTypeSelfMount,
	"automated_mount": BackupDestinationNfsMountTypeAutomatedMount,
}

// GetBackupDestinationNfsMountTypeEnumValues Enumerates the set of values for BackupDestinationNfsMountTypeEnum
func GetBackupDestinationNfsMountTypeEnumValues() []BackupDestinationNfsMountTypeEnum {
	values := make([]BackupDestinationNfsMountTypeEnum, 0)
	for _, v := range mappingBackupDestinationNfsMountTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupDestinationNfsMountTypeEnumStringValues Enumerates the set of values in String for BackupDestinationNfsMountTypeEnum
func GetBackupDestinationNfsMountTypeEnumStringValues() []string {
	return []string{
		"SELF_MOUNT",
		"AUTOMATED_MOUNT",
	}
}

// GetMappingBackupDestinationNfsMountTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupDestinationNfsMountTypeEnum(val string) (BackupDestinationNfsMountTypeEnum, bool) {
	enum, ok := mappingBackupDestinationNfsMountTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BackupDestinationLifecycleStateEnum Enum with underlying type: string
type BackupDestinationLifecycleStateEnum string

// Set of constants representing the allowable values for BackupDestinationLifecycleStateEnum
const (
	BackupDestinationLifecycleStateActive  BackupDestinationLifecycleStateEnum = "ACTIVE"
	BackupDestinationLifecycleStateFailed  BackupDestinationLifecycleStateEnum = "FAILED"
	BackupDestinationLifecycleStateDeleted BackupDestinationLifecycleStateEnum = "DELETED"
)

var mappingBackupDestinationLifecycleStateEnum = map[string]BackupDestinationLifecycleStateEnum{
	"ACTIVE":  BackupDestinationLifecycleStateActive,
	"FAILED":  BackupDestinationLifecycleStateFailed,
	"DELETED": BackupDestinationLifecycleStateDeleted,
}

var mappingBackupDestinationLifecycleStateEnumLowerCase = map[string]BackupDestinationLifecycleStateEnum{
	"active":  BackupDestinationLifecycleStateActive,
	"failed":  BackupDestinationLifecycleStateFailed,
	"deleted": BackupDestinationLifecycleStateDeleted,
}

// GetBackupDestinationLifecycleStateEnumValues Enumerates the set of values for BackupDestinationLifecycleStateEnum
func GetBackupDestinationLifecycleStateEnumValues() []BackupDestinationLifecycleStateEnum {
	values := make([]BackupDestinationLifecycleStateEnum, 0)
	for _, v := range mappingBackupDestinationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupDestinationLifecycleStateEnumStringValues Enumerates the set of values in String for BackupDestinationLifecycleStateEnum
func GetBackupDestinationLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"FAILED",
		"DELETED",
	}
}

// GetMappingBackupDestinationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupDestinationLifecycleStateEnum(val string) (BackupDestinationLifecycleStateEnum, bool) {
	enum, ok := mappingBackupDestinationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
