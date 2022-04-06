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

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the host IP address associated with the database node. Use this OCID with either the
	// GetPrivateIp or the GetPublicIpByPrivateIpId API to get the IP address
	// needed to make a database connection.
	// **Note:** Applies only to Exadata Cloud Service.
	HostIpId *string `mandatory:"false" json:"hostIpId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the backup IP address associated with the database node. Use this OCID with either the
	// GetPrivateIp or the GetPublicIpByPrivateIpId API to get the IP address
	// needed to make a database connection.
	// **Note:** Applies only to Exadata Cloud Service.
	BackupIpId *string `mandatory:"false" json:"backupIpId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the second VNIC.
	// **Note:** Applies only to Exadata Cloud Service.
	Vnic2Id *string `mandatory:"false" json:"vnic2Id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the second backup VNIC.
	// **Note:** Applies only to Exadata Cloud Service.
	BackupVnic2Id *string `mandatory:"false" json:"backupVnic2Id"`

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

	// The number of CPU cores enabled on the Db node.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// The allocated memory in GBs on the Db node.
	MemorySizeInGBs *int `mandatory:"false" json:"memorySizeInGBs"`

	// The allocated local node storage in GBs on the Db node.
	DbNodeStorageSizeInGBs *int `mandatory:"false" json:"dbNodeStorageSizeInGBs"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exacc Db server associated with the database node.
	DbServerId *string `mandatory:"false" json:"dbServerId"`
}

func (m DbNodeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbNodeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbNodeSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbNodeSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDbNodeSummaryMaintenanceTypeEnum(string(m.MaintenanceType)); !ok && m.MaintenanceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceType: %s. Supported values are: %s.", m.MaintenanceType, strings.Join(GetDbNodeSummaryMaintenanceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingDbNodeSummaryLifecycleStateEnum = map[string]DbNodeSummaryLifecycleStateEnum{
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

var mappingDbNodeSummaryLifecycleStateEnumLowerCase = map[string]DbNodeSummaryLifecycleStateEnum{
	"provisioning": DbNodeSummaryLifecycleStateProvisioning,
	"available":    DbNodeSummaryLifecycleStateAvailable,
	"updating":     DbNodeSummaryLifecycleStateUpdating,
	"stopping":     DbNodeSummaryLifecycleStateStopping,
	"stopped":      DbNodeSummaryLifecycleStateStopped,
	"starting":     DbNodeSummaryLifecycleStateStarting,
	"terminating":  DbNodeSummaryLifecycleStateTerminating,
	"terminated":   DbNodeSummaryLifecycleStateTerminated,
	"failed":       DbNodeSummaryLifecycleStateFailed,
}

// GetDbNodeSummaryLifecycleStateEnumValues Enumerates the set of values for DbNodeSummaryLifecycleStateEnum
func GetDbNodeSummaryLifecycleStateEnumValues() []DbNodeSummaryLifecycleStateEnum {
	values := make([]DbNodeSummaryLifecycleStateEnum, 0)
	for _, v := range mappingDbNodeSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDbNodeSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for DbNodeSummaryLifecycleStateEnum
func GetDbNodeSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"UPDATING",
		"STOPPING",
		"STOPPED",
		"STARTING",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
	}
}

// GetMappingDbNodeSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbNodeSummaryLifecycleStateEnum(val string) (DbNodeSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingDbNodeSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DbNodeSummaryMaintenanceTypeEnum Enum with underlying type: string
type DbNodeSummaryMaintenanceTypeEnum string

// Set of constants representing the allowable values for DbNodeSummaryMaintenanceTypeEnum
const (
	DbNodeSummaryMaintenanceTypeVmdbRebootMigration DbNodeSummaryMaintenanceTypeEnum = "VMDB_REBOOT_MIGRATION"
)

var mappingDbNodeSummaryMaintenanceTypeEnum = map[string]DbNodeSummaryMaintenanceTypeEnum{
	"VMDB_REBOOT_MIGRATION": DbNodeSummaryMaintenanceTypeVmdbRebootMigration,
}

var mappingDbNodeSummaryMaintenanceTypeEnumLowerCase = map[string]DbNodeSummaryMaintenanceTypeEnum{
	"vmdb_reboot_migration": DbNodeSummaryMaintenanceTypeVmdbRebootMigration,
}

// GetDbNodeSummaryMaintenanceTypeEnumValues Enumerates the set of values for DbNodeSummaryMaintenanceTypeEnum
func GetDbNodeSummaryMaintenanceTypeEnumValues() []DbNodeSummaryMaintenanceTypeEnum {
	values := make([]DbNodeSummaryMaintenanceTypeEnum, 0)
	for _, v := range mappingDbNodeSummaryMaintenanceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDbNodeSummaryMaintenanceTypeEnumStringValues Enumerates the set of values in String for DbNodeSummaryMaintenanceTypeEnum
func GetDbNodeSummaryMaintenanceTypeEnumStringValues() []string {
	return []string{
		"VMDB_REBOOT_MIGRATION",
	}
}

// GetMappingDbNodeSummaryMaintenanceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbNodeSummaryMaintenanceTypeEnum(val string) (DbNodeSummaryMaintenanceTypeEnum, bool) {
	enum, ok := mappingDbNodeSummaryMaintenanceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
