// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// AutonomousVirtualMachine Autonomous Virtual Machine details.
type AutonomousVirtualMachine struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous Virtual Machine.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the Autonomous Virtual Machine.
	LifecycleState AutonomousVirtualMachineLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The name of the Autonomous Virtual Machine.
	VmName *string `mandatory:"false" json:"vmName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Db server associated with the Autonomous Virtual Machine.
	DbServerId *string `mandatory:"false" json:"dbServerId"`

	// The display name of the dbServer associated with the Autonomous Virtual Machine.
	DbServerDisplayName *string `mandatory:"false" json:"dbServerDisplayName"`

	// The number of CPU cores enabled on the Autonomous Virtual Machine.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// The allocated memory in GBs on the Autonomous Virtual Machine.
	MemorySizeInGBs *int `mandatory:"false" json:"memorySizeInGBs"`

	// The allocated local node storage in GBs on the Autonomous Virtual Machine.
	DbNodeStorageSizeInGBs *int `mandatory:"false" json:"dbNodeStorageSizeInGBs"`

	// Client IP Address.
	ClientIpAddress *string `mandatory:"false" json:"clientIpAddress"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous VM Cluster associated with the Autonomous Virtual Machine.
	AutonomousVmClusterId *string `mandatory:"false" json:"autonomousVmClusterId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Cloud Autonomous VM Cluster associated with the Autonomous Virtual Machine.
	CloudAutonomousVmClusterId *string `mandatory:"false" json:"cloudAutonomousVmClusterId"`
}

func (m AutonomousVirtualMachine) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousVirtualMachine) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousVirtualMachineLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousVirtualMachineLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousVirtualMachineLifecycleStateEnum Enum with underlying type: string
type AutonomousVirtualMachineLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousVirtualMachineLifecycleStateEnum
const (
	AutonomousVirtualMachineLifecycleStateProvisioning          AutonomousVirtualMachineLifecycleStateEnum = "PROVISIONING"
	AutonomousVirtualMachineLifecycleStateAvailable             AutonomousVirtualMachineLifecycleStateEnum = "AVAILABLE"
	AutonomousVirtualMachineLifecycleStateUpdating              AutonomousVirtualMachineLifecycleStateEnum = "UPDATING"
	AutonomousVirtualMachineLifecycleStateTerminating           AutonomousVirtualMachineLifecycleStateEnum = "TERMINATING"
	AutonomousVirtualMachineLifecycleStateTerminated            AutonomousVirtualMachineLifecycleStateEnum = "TERMINATED"
	AutonomousVirtualMachineLifecycleStateFailed                AutonomousVirtualMachineLifecycleStateEnum = "FAILED"
	AutonomousVirtualMachineLifecycleStateMaintenanceInProgress AutonomousVirtualMachineLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
)

var mappingAutonomousVirtualMachineLifecycleStateEnum = map[string]AutonomousVirtualMachineLifecycleStateEnum{
	"PROVISIONING":            AutonomousVirtualMachineLifecycleStateProvisioning,
	"AVAILABLE":               AutonomousVirtualMachineLifecycleStateAvailable,
	"UPDATING":                AutonomousVirtualMachineLifecycleStateUpdating,
	"TERMINATING":             AutonomousVirtualMachineLifecycleStateTerminating,
	"TERMINATED":              AutonomousVirtualMachineLifecycleStateTerminated,
	"FAILED":                  AutonomousVirtualMachineLifecycleStateFailed,
	"MAINTENANCE_IN_PROGRESS": AutonomousVirtualMachineLifecycleStateMaintenanceInProgress,
}

var mappingAutonomousVirtualMachineLifecycleStateEnumLowerCase = map[string]AutonomousVirtualMachineLifecycleStateEnum{
	"provisioning":            AutonomousVirtualMachineLifecycleStateProvisioning,
	"available":               AutonomousVirtualMachineLifecycleStateAvailable,
	"updating":                AutonomousVirtualMachineLifecycleStateUpdating,
	"terminating":             AutonomousVirtualMachineLifecycleStateTerminating,
	"terminated":              AutonomousVirtualMachineLifecycleStateTerminated,
	"failed":                  AutonomousVirtualMachineLifecycleStateFailed,
	"maintenance_in_progress": AutonomousVirtualMachineLifecycleStateMaintenanceInProgress,
}

// GetAutonomousVirtualMachineLifecycleStateEnumValues Enumerates the set of values for AutonomousVirtualMachineLifecycleStateEnum
func GetAutonomousVirtualMachineLifecycleStateEnumValues() []AutonomousVirtualMachineLifecycleStateEnum {
	values := make([]AutonomousVirtualMachineLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousVirtualMachineLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousVirtualMachineLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousVirtualMachineLifecycleStateEnum
func GetAutonomousVirtualMachineLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"UPDATING",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
		"MAINTENANCE_IN_PROGRESS",
	}
}

// GetMappingAutonomousVirtualMachineLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousVirtualMachineLifecycleStateEnum(val string) (AutonomousVirtualMachineLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousVirtualMachineLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
