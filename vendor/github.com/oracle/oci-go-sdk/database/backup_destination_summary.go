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

// BackupDestinationSummary Backup destination details, including the list of databases using the backup destination.
type BackupDestinationSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the backup destination.
	Id *string `mandatory:"false" json:"id"`

	// The user-provided name of the backup destination.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Type of the backup destination.
	Type BackupDestinationSummaryTypeEnum `mandatory:"false" json:"type,omitempty"`

	// List of databases associated with the backup destination.
	AssociatedDatabases []AssociatedDatabaseDetails `mandatory:"false" json:"associatedDatabases"`

	// For a RECOVERY_APPLIANCE backup destination, the connection string for connecting to the Recovery Appliance.
	ConnectionString *string `mandatory:"false" json:"connectionString"`

	// For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) users that are used to access the Recovery Appliance.
	VpcUsers []string `mandatory:"false" json:"vpcUsers"`

	// The local directory path on each VM cluster node where the NFS server location is mounted. The local directory path and the NFS server location must each be the same across all of the VM cluster nodes. Ensure that the NFS mount is maintained continuously on all of the VM cluster nodes.
	LocalMountPointPath *string `mandatory:"false" json:"localMountPointPath"`

	// NFS Mount type for backup destination.
	NfsMountType BackupDestinationSummaryNfsMountTypeEnum `mandatory:"false" json:"nfsMountType,omitempty"`

	// Host names or IP addresses for NFS Auto mount.
	NfsServer []string `mandatory:"false" json:"nfsServer"`

	// Specifies the directory on which to mount the file system
	NfsServerExport *string `mandatory:"false" json:"nfsServerExport"`

	// The current lifecycle state of the backup destination.
	LifecycleState BackupDestinationSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

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

func (m BackupDestinationSummary) String() string {
	return common.PointerString(m)
}

// BackupDestinationSummaryTypeEnum Enum with underlying type: string
type BackupDestinationSummaryTypeEnum string

// Set of constants representing the allowable values for BackupDestinationSummaryTypeEnum
const (
	BackupDestinationSummaryTypeNfs               BackupDestinationSummaryTypeEnum = "NFS"
	BackupDestinationSummaryTypeRecoveryAppliance BackupDestinationSummaryTypeEnum = "RECOVERY_APPLIANCE"
)

var mappingBackupDestinationSummaryType = map[string]BackupDestinationSummaryTypeEnum{
	"NFS":                BackupDestinationSummaryTypeNfs,
	"RECOVERY_APPLIANCE": BackupDestinationSummaryTypeRecoveryAppliance,
}

// GetBackupDestinationSummaryTypeEnumValues Enumerates the set of values for BackupDestinationSummaryTypeEnum
func GetBackupDestinationSummaryTypeEnumValues() []BackupDestinationSummaryTypeEnum {
	values := make([]BackupDestinationSummaryTypeEnum, 0)
	for _, v := range mappingBackupDestinationSummaryType {
		values = append(values, v)
	}
	return values
}

// BackupDestinationSummaryNfsMountTypeEnum Enum with underlying type: string
type BackupDestinationSummaryNfsMountTypeEnum string

// Set of constants representing the allowable values for BackupDestinationSummaryNfsMountTypeEnum
const (
	BackupDestinationSummaryNfsMountTypeSelfMount      BackupDestinationSummaryNfsMountTypeEnum = "SELF_MOUNT"
	BackupDestinationSummaryNfsMountTypeAutomatedMount BackupDestinationSummaryNfsMountTypeEnum = "AUTOMATED_MOUNT"
)

var mappingBackupDestinationSummaryNfsMountType = map[string]BackupDestinationSummaryNfsMountTypeEnum{
	"SELF_MOUNT":      BackupDestinationSummaryNfsMountTypeSelfMount,
	"AUTOMATED_MOUNT": BackupDestinationSummaryNfsMountTypeAutomatedMount,
}

// GetBackupDestinationSummaryNfsMountTypeEnumValues Enumerates the set of values for BackupDestinationSummaryNfsMountTypeEnum
func GetBackupDestinationSummaryNfsMountTypeEnumValues() []BackupDestinationSummaryNfsMountTypeEnum {
	values := make([]BackupDestinationSummaryNfsMountTypeEnum, 0)
	for _, v := range mappingBackupDestinationSummaryNfsMountType {
		values = append(values, v)
	}
	return values
}

// BackupDestinationSummaryLifecycleStateEnum Enum with underlying type: string
type BackupDestinationSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for BackupDestinationSummaryLifecycleStateEnum
const (
	BackupDestinationSummaryLifecycleStateActive  BackupDestinationSummaryLifecycleStateEnum = "ACTIVE"
	BackupDestinationSummaryLifecycleStateFailed  BackupDestinationSummaryLifecycleStateEnum = "FAILED"
	BackupDestinationSummaryLifecycleStateDeleted BackupDestinationSummaryLifecycleStateEnum = "DELETED"
)

var mappingBackupDestinationSummaryLifecycleState = map[string]BackupDestinationSummaryLifecycleStateEnum{
	"ACTIVE":  BackupDestinationSummaryLifecycleStateActive,
	"FAILED":  BackupDestinationSummaryLifecycleStateFailed,
	"DELETED": BackupDestinationSummaryLifecycleStateDeleted,
}

// GetBackupDestinationSummaryLifecycleStateEnumValues Enumerates the set of values for BackupDestinationSummaryLifecycleStateEnum
func GetBackupDestinationSummaryLifecycleStateEnumValues() []BackupDestinationSummaryLifecycleStateEnum {
	values := make([]BackupDestinationSummaryLifecycleStateEnum, 0)
	for _, v := range mappingBackupDestinationSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
