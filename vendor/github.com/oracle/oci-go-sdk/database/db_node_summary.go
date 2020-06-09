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

// DbNodeSummary A server where Oracle Database software is running.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type DbNodeSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database node.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DB system.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VNIC.
	VnicId *string `mandatory:"true" json:"vnicId"`

	// The current state of the database node.
	LifecycleState DbNodeSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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
	MaintenanceType DbNodeSummaryMaintenanceTypeEnum `mandatory:"false" json:"maintenanceType,omitempty"`

	// Start date and time of maintenance window.
	TimeMaintenanceWindowStart *common.SDKTime `mandatory:"false" json:"timeMaintenanceWindowStart"`

	// End date and time of maintenance window.
	TimeMaintenanceWindowEnd *common.SDKTime `mandatory:"false" json:"timeMaintenanceWindowEnd"`

	// Additional information about the planned maintenance.
	AdditionalDetails *string `mandatory:"false" json:"additionalDetails"`
}

func (m DbNodeSummary) String() string {
	return common.PointerString(m)
}

// DbNodeSummaryLifecycleStateEnum Enum with underlying type: string
type DbNodeSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for DbNodeSummaryLifecycleStateEnum
const (
	DbNodeSummaryLifecycleStateProvisioning DbNodeSummaryLifecycleStateEnum = "PROVISIONING"
	DbNodeSummaryLifecycleStateAvailable    DbNodeSummaryLifecycleStateEnum = "AVAILABLE"
	DbNodeSummaryLifecycleStateUpdating     DbNodeSummaryLifecycleStateEnum = "UPDATING"
	DbNodeSummaryLifecycleStateStopping     DbNodeSummaryLifecycleStateEnum = "STOPPING"
	DbNodeSummaryLifecycleStateStopped      DbNodeSummaryLifecycleStateEnum = "STOPPED"
	DbNodeSummaryLifecycleStateStarting     DbNodeSummaryLifecycleStateEnum = "STARTING"
	DbNodeSummaryLifecycleStateTerminating  DbNodeSummaryLifecycleStateEnum = "TERMINATING"
	DbNodeSummaryLifecycleStateTerminated   DbNodeSummaryLifecycleStateEnum = "TERMINATED"
	DbNodeSummaryLifecycleStateFailed       DbNodeSummaryLifecycleStateEnum = "FAILED"
)

var mappingDbNodeSummaryLifecycleState = map[string]DbNodeSummaryLifecycleStateEnum{
	"PROVISIONING": DbNodeSummaryLifecycleStateProvisioning,
	"AVAILABLE":    DbNodeSummaryLifecycleStateAvailable,
	"UPDATING":     DbNodeSummaryLifecycleStateUpdating,
	"STOPPING":     DbNodeSummaryLifecycleStateStopping,
	"STOPPED":      DbNodeSummaryLifecycleStateStopped,
	"STARTING":     DbNodeSummaryLifecycleStateStarting,
	"TERMINATING":  DbNodeSummaryLifecycleStateTerminating,
	"TERMINATED":   DbNodeSummaryLifecycleStateTerminated,
	"FAILED":       DbNodeSummaryLifecycleStateFailed,
}

// GetDbNodeSummaryLifecycleStateEnumValues Enumerates the set of values for DbNodeSummaryLifecycleStateEnum
func GetDbNodeSummaryLifecycleStateEnumValues() []DbNodeSummaryLifecycleStateEnum {
	values := make([]DbNodeSummaryLifecycleStateEnum, 0)
	for _, v := range mappingDbNodeSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}

// DbNodeSummaryMaintenanceTypeEnum Enum with underlying type: string
type DbNodeSummaryMaintenanceTypeEnum string

// Set of constants representing the allowable values for DbNodeSummaryMaintenanceTypeEnum
const (
	DbNodeSummaryMaintenanceTypeVmdbRebootMigration DbNodeSummaryMaintenanceTypeEnum = "VMDB_REBOOT_MIGRATION"
)

var mappingDbNodeSummaryMaintenanceType = map[string]DbNodeSummaryMaintenanceTypeEnum{
	"VMDB_REBOOT_MIGRATION": DbNodeSummaryMaintenanceTypeVmdbRebootMigration,
}

// GetDbNodeSummaryMaintenanceTypeEnumValues Enumerates the set of values for DbNodeSummaryMaintenanceTypeEnum
func GetDbNodeSummaryMaintenanceTypeEnumValues() []DbNodeSummaryMaintenanceTypeEnum {
	values := make([]DbNodeSummaryMaintenanceTypeEnum, 0)
	for _, v := range mappingDbNodeSummaryMaintenanceType {
		values = append(values, v)
	}
	return values
}
