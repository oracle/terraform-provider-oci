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

// DbNode The representation of DbNode
type DbNode struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database node.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DB system.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VNIC.
	VnicId *string `mandatory:"true" json:"vnicId"`

	// The current state of the database node.
	LifecycleState DbNodeLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time that the database node was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the backup VNIC.
	BackupVnicId *string `mandatory:"false" json:"backupVnicId"`

	// The host name for the database node.
	Hostname *string `mandatory:"false" json:"hostname"`

	// The name of the Fault Domain the instance is contained in.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// The size (in GB) of the block storage volume allocation for the DB system. This attribute applies only for virtual machine DB systems.
	SoftwareStorageSizeInGB *int `mandatory:"false" json:"softwareStorageSizeInGB"`

	// The type of database node maintenance.
	MaintenanceType DbNodeMaintenanceTypeEnum `mandatory:"false" json:"maintenanceType,omitempty"`

	// Start date and time of maintenance window.
	TimeMaintenanceWindowStart *common.SDKTime `mandatory:"false" json:"timeMaintenanceWindowStart"`

	// End date and time of maintenance window.
	TimeMaintenanceWindowEnd *common.SDKTime `mandatory:"false" json:"timeMaintenanceWindowEnd"`

	// Additional information about the planned maintenance.
	AdditionalDetails *string `mandatory:"false" json:"additionalDetails"`
}

func (m DbNode) String() string {
	return common.PointerString(m)
}

// DbNodeLifecycleStateEnum Enum with underlying type: string
type DbNodeLifecycleStateEnum string

// Set of constants representing the allowable values for DbNodeLifecycleStateEnum
const (
	DbNodeLifecycleStateProvisioning DbNodeLifecycleStateEnum = "PROVISIONING"
	DbNodeLifecycleStateAvailable    DbNodeLifecycleStateEnum = "AVAILABLE"
	DbNodeLifecycleStateUpdating     DbNodeLifecycleStateEnum = "UPDATING"
	DbNodeLifecycleStateStopping     DbNodeLifecycleStateEnum = "STOPPING"
	DbNodeLifecycleStateStopped      DbNodeLifecycleStateEnum = "STOPPED"
	DbNodeLifecycleStateStarting     DbNodeLifecycleStateEnum = "STARTING"
	DbNodeLifecycleStateTerminating  DbNodeLifecycleStateEnum = "TERMINATING"
	DbNodeLifecycleStateTerminated   DbNodeLifecycleStateEnum = "TERMINATED"
	DbNodeLifecycleStateFailed       DbNodeLifecycleStateEnum = "FAILED"
)

var mappingDbNodeLifecycleState = map[string]DbNodeLifecycleStateEnum{
	"PROVISIONING": DbNodeLifecycleStateProvisioning,
	"AVAILABLE":    DbNodeLifecycleStateAvailable,
	"UPDATING":     DbNodeLifecycleStateUpdating,
	"STOPPING":     DbNodeLifecycleStateStopping,
	"STOPPED":      DbNodeLifecycleStateStopped,
	"STARTING":     DbNodeLifecycleStateStarting,
	"TERMINATING":  DbNodeLifecycleStateTerminating,
	"TERMINATED":   DbNodeLifecycleStateTerminated,
	"FAILED":       DbNodeLifecycleStateFailed,
}

// GetDbNodeLifecycleStateEnumValues Enumerates the set of values for DbNodeLifecycleStateEnum
func GetDbNodeLifecycleStateEnumValues() []DbNodeLifecycleStateEnum {
	values := make([]DbNodeLifecycleStateEnum, 0)
	for _, v := range mappingDbNodeLifecycleState {
		values = append(values, v)
	}
	return values
}

// DbNodeMaintenanceTypeEnum Enum with underlying type: string
type DbNodeMaintenanceTypeEnum string

// Set of constants representing the allowable values for DbNodeMaintenanceTypeEnum
const (
	DbNodeMaintenanceTypeVmdbRebootMigration DbNodeMaintenanceTypeEnum = "VMDB_REBOOT_MIGRATION"
)

var mappingDbNodeMaintenanceType = map[string]DbNodeMaintenanceTypeEnum{
	"VMDB_REBOOT_MIGRATION": DbNodeMaintenanceTypeVmdbRebootMigration,
}

// GetDbNodeMaintenanceTypeEnumValues Enumerates the set of values for DbNodeMaintenanceTypeEnum
func GetDbNodeMaintenanceTypeEnumValues() []DbNodeMaintenanceTypeEnum {
	values := make([]DbNodeMaintenanceTypeEnum, 0)
	for _, v := range mappingDbNodeMaintenanceType {
		values = append(values, v)
	}
	return values
}
