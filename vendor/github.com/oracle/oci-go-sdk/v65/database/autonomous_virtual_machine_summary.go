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

// AutonomousVirtualMachineSummary Details of the Autonomous Virtual Machine.
type AutonomousVirtualMachineSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous Virtual Machine.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the Autonomous Virtual Machine.
	LifecycleState AutonomousVirtualMachineSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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

func (m AutonomousVirtualMachineSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousVirtualMachineSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousVirtualMachineSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousVirtualMachineSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousVirtualMachineSummaryLifecycleStateEnum Enum with underlying type: string
type AutonomousVirtualMachineSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousVirtualMachineSummaryLifecycleStateEnum
const (
	AutonomousVirtualMachineSummaryLifecycleStateProvisioning          AutonomousVirtualMachineSummaryLifecycleStateEnum = "PROVISIONING"
	AutonomousVirtualMachineSummaryLifecycleStateAvailable             AutonomousVirtualMachineSummaryLifecycleStateEnum = "AVAILABLE"
	AutonomousVirtualMachineSummaryLifecycleStateUpdating              AutonomousVirtualMachineSummaryLifecycleStateEnum = "UPDATING"
	AutonomousVirtualMachineSummaryLifecycleStateTerminating           AutonomousVirtualMachineSummaryLifecycleStateEnum = "TERMINATING"
	AutonomousVirtualMachineSummaryLifecycleStateTerminated            AutonomousVirtualMachineSummaryLifecycleStateEnum = "TERMINATED"
	AutonomousVirtualMachineSummaryLifecycleStateFailed                AutonomousVirtualMachineSummaryLifecycleStateEnum = "FAILED"
	AutonomousVirtualMachineSummaryLifecycleStateMaintenanceInProgress AutonomousVirtualMachineSummaryLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
)

var mappingAutonomousVirtualMachineSummaryLifecycleStateEnum = map[string]AutonomousVirtualMachineSummaryLifecycleStateEnum{
	"PROVISIONING":            AutonomousVirtualMachineSummaryLifecycleStateProvisioning,
	"AVAILABLE":               AutonomousVirtualMachineSummaryLifecycleStateAvailable,
	"UPDATING":                AutonomousVirtualMachineSummaryLifecycleStateUpdating,
	"TERMINATING":             AutonomousVirtualMachineSummaryLifecycleStateTerminating,
	"TERMINATED":              AutonomousVirtualMachineSummaryLifecycleStateTerminated,
	"FAILED":                  AutonomousVirtualMachineSummaryLifecycleStateFailed,
	"MAINTENANCE_IN_PROGRESS": AutonomousVirtualMachineSummaryLifecycleStateMaintenanceInProgress,
}

var mappingAutonomousVirtualMachineSummaryLifecycleStateEnumLowerCase = map[string]AutonomousVirtualMachineSummaryLifecycleStateEnum{
	"provisioning":            AutonomousVirtualMachineSummaryLifecycleStateProvisioning,
	"available":               AutonomousVirtualMachineSummaryLifecycleStateAvailable,
	"updating":                AutonomousVirtualMachineSummaryLifecycleStateUpdating,
	"terminating":             AutonomousVirtualMachineSummaryLifecycleStateTerminating,
	"terminated":              AutonomousVirtualMachineSummaryLifecycleStateTerminated,
	"failed":                  AutonomousVirtualMachineSummaryLifecycleStateFailed,
	"maintenance_in_progress": AutonomousVirtualMachineSummaryLifecycleStateMaintenanceInProgress,
}

// GetAutonomousVirtualMachineSummaryLifecycleStateEnumValues Enumerates the set of values for AutonomousVirtualMachineSummaryLifecycleStateEnum
func GetAutonomousVirtualMachineSummaryLifecycleStateEnumValues() []AutonomousVirtualMachineSummaryLifecycleStateEnum {
	values := make([]AutonomousVirtualMachineSummaryLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousVirtualMachineSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousVirtualMachineSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousVirtualMachineSummaryLifecycleStateEnum
func GetAutonomousVirtualMachineSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingAutonomousVirtualMachineSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousVirtualMachineSummaryLifecycleStateEnum(val string) (AutonomousVirtualMachineSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousVirtualMachineSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
