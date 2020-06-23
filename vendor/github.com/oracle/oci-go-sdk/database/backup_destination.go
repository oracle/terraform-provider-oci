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

// BackupDestinationTypeEnum Enum with underlying type: string
type BackupDestinationTypeEnum string

// Set of constants representing the allowable values for BackupDestinationTypeEnum
const (
	BackupDestinationTypeNfs               BackupDestinationTypeEnum = "NFS"
	BackupDestinationTypeRecoveryAppliance BackupDestinationTypeEnum = "RECOVERY_APPLIANCE"
)

var mappingBackupDestinationType = map[string]BackupDestinationTypeEnum{
	"NFS":                BackupDestinationTypeNfs,
	"RECOVERY_APPLIANCE": BackupDestinationTypeRecoveryAppliance,
}

// GetBackupDestinationTypeEnumValues Enumerates the set of values for BackupDestinationTypeEnum
func GetBackupDestinationTypeEnumValues() []BackupDestinationTypeEnum {
	values := make([]BackupDestinationTypeEnum, 0)
	for _, v := range mappingBackupDestinationType {
		values = append(values, v)
	}
	return values
}

// BackupDestinationNfsMountTypeEnum Enum with underlying type: string
type BackupDestinationNfsMountTypeEnum string

// Set of constants representing the allowable values for BackupDestinationNfsMountTypeEnum
const (
	BackupDestinationNfsMountTypeSelfMount      BackupDestinationNfsMountTypeEnum = "SELF_MOUNT"
	BackupDestinationNfsMountTypeAutomatedMount BackupDestinationNfsMountTypeEnum = "AUTOMATED_MOUNT"
)

var mappingBackupDestinationNfsMountType = map[string]BackupDestinationNfsMountTypeEnum{
	"SELF_MOUNT":      BackupDestinationNfsMountTypeSelfMount,
	"AUTOMATED_MOUNT": BackupDestinationNfsMountTypeAutomatedMount,
}

// GetBackupDestinationNfsMountTypeEnumValues Enumerates the set of values for BackupDestinationNfsMountTypeEnum
func GetBackupDestinationNfsMountTypeEnumValues() []BackupDestinationNfsMountTypeEnum {
	values := make([]BackupDestinationNfsMountTypeEnum, 0)
	for _, v := range mappingBackupDestinationNfsMountType {
		values = append(values, v)
	}
	return values
}

// BackupDestinationLifecycleStateEnum Enum with underlying type: string
type BackupDestinationLifecycleStateEnum string

// Set of constants representing the allowable values for BackupDestinationLifecycleStateEnum
const (
	BackupDestinationLifecycleStateActive  BackupDestinationLifecycleStateEnum = "ACTIVE"
	BackupDestinationLifecycleStateFailed  BackupDestinationLifecycleStateEnum = "FAILED"
	BackupDestinationLifecycleStateDeleted BackupDestinationLifecycleStateEnum = "DELETED"
)

var mappingBackupDestinationLifecycleState = map[string]BackupDestinationLifecycleStateEnum{
	"ACTIVE":  BackupDestinationLifecycleStateActive,
	"FAILED":  BackupDestinationLifecycleStateFailed,
	"DELETED": BackupDestinationLifecycleStateDeleted,
}

// GetBackupDestinationLifecycleStateEnumValues Enumerates the set of values for BackupDestinationLifecycleStateEnum
func GetBackupDestinationLifecycleStateEnumValues() []BackupDestinationLifecycleStateEnum {
	values := make([]BackupDestinationLifecycleStateEnum, 0)
	for _, v := range mappingBackupDestinationLifecycleState {
		values = append(values, v)
	}
	return values
}
