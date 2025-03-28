// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// DbServer Details of the Db server resource.
type DbServer struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Db server.
	Id *string `mandatory:"false" json:"id"`

	// The user-friendly name for the Db server. The name does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
	ExadataInfrastructureId *string `mandatory:"false" json:"exadataInfrastructureId"`

	// The number of CPU cores enabled on the Db server.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// The allocated memory in GBs on the Db server.
	MemorySizeInGBs *int `mandatory:"false" json:"memorySizeInGBs"`

	// The allocated local node storage in GBs on the Db server.
	DbNodeStorageSizeInGBs *int `mandatory:"false" json:"dbNodeStorageSizeInGBs"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM Clusters associated with the Db server.
	VmClusterIds []string `mandatory:"false" json:"vmClusterIds"`

	// The list of OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous VM Clusters associated with the Db server.
	AutonomousVmClusterIds []string `mandatory:"false" json:"autonomousVmClusterIds"`

	// The list of OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Virtual Machines associated with the Db server.
	AutonomousVirtualMachineIds []string `mandatory:"false" json:"autonomousVirtualMachineIds"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Db nodes associated with the Db server.
	DbNodeIds []string `mandatory:"false" json:"dbNodeIds"`

	// The shape of the Db server. The shape determines the amount of CPU, storage, and memory resources available.
	Shape *string `mandatory:"false" json:"shape"`

	// The current state of the Db server.
	LifecycleState DbServerLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The total number of CPU cores available.
	MaxCpuCount *int `mandatory:"false" json:"maxCpuCount"`

	// The total memory available in GBs.
	MaxMemoryInGBs *int `mandatory:"false" json:"maxMemoryInGBs"`

	// The total local node storage available in GBs.
	MaxDbNodeStorageInGBs *int `mandatory:"false" json:"maxDbNodeStorageInGBs"`

	// The date and time that the Db Server was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	DbServerPatchingDetails *DbServerPatchingDetails `mandatory:"false" json:"dbServerPatchingDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The compute model of the Autonomous Database. This is required if using the `computeCount` parameter. If using `cpuCoreCount` then it is an error to specify `computeModel` to a non-null value. ECPU compute model is the recommended model and OCPU compute model is legacy.
	ComputeModel DbServerComputeModelEnum `mandatory:"false" json:"computeModel,omitempty"`
}

func (m DbServer) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbServer) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDbServerLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbServerLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbServerComputeModelEnum(string(m.ComputeModel)); !ok && m.ComputeModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComputeModel: %s. Supported values are: %s.", m.ComputeModel, strings.Join(GetDbServerComputeModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbServerLifecycleStateEnum Enum with underlying type: string
type DbServerLifecycleStateEnum string

// Set of constants representing the allowable values for DbServerLifecycleStateEnum
const (
	DbServerLifecycleStateCreating              DbServerLifecycleStateEnum = "CREATING"
	DbServerLifecycleStateAvailable             DbServerLifecycleStateEnum = "AVAILABLE"
	DbServerLifecycleStateUnavailable           DbServerLifecycleStateEnum = "UNAVAILABLE"
	DbServerLifecycleStateDeleting              DbServerLifecycleStateEnum = "DELETING"
	DbServerLifecycleStateDeleted               DbServerLifecycleStateEnum = "DELETED"
	DbServerLifecycleStateMaintenanceInProgress DbServerLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
)

var mappingDbServerLifecycleStateEnum = map[string]DbServerLifecycleStateEnum{
	"CREATING":                DbServerLifecycleStateCreating,
	"AVAILABLE":               DbServerLifecycleStateAvailable,
	"UNAVAILABLE":             DbServerLifecycleStateUnavailable,
	"DELETING":                DbServerLifecycleStateDeleting,
	"DELETED":                 DbServerLifecycleStateDeleted,
	"MAINTENANCE_IN_PROGRESS": DbServerLifecycleStateMaintenanceInProgress,
}

var mappingDbServerLifecycleStateEnumLowerCase = map[string]DbServerLifecycleStateEnum{
	"creating":                DbServerLifecycleStateCreating,
	"available":               DbServerLifecycleStateAvailable,
	"unavailable":             DbServerLifecycleStateUnavailable,
	"deleting":                DbServerLifecycleStateDeleting,
	"deleted":                 DbServerLifecycleStateDeleted,
	"maintenance_in_progress": DbServerLifecycleStateMaintenanceInProgress,
}

// GetDbServerLifecycleStateEnumValues Enumerates the set of values for DbServerLifecycleStateEnum
func GetDbServerLifecycleStateEnumValues() []DbServerLifecycleStateEnum {
	values := make([]DbServerLifecycleStateEnum, 0)
	for _, v := range mappingDbServerLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDbServerLifecycleStateEnumStringValues Enumerates the set of values in String for DbServerLifecycleStateEnum
func GetDbServerLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"AVAILABLE",
		"UNAVAILABLE",
		"DELETING",
		"DELETED",
		"MAINTENANCE_IN_PROGRESS",
	}
}

// GetMappingDbServerLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbServerLifecycleStateEnum(val string) (DbServerLifecycleStateEnum, bool) {
	enum, ok := mappingDbServerLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DbServerComputeModelEnum Enum with underlying type: string
type DbServerComputeModelEnum string

// Set of constants representing the allowable values for DbServerComputeModelEnum
const (
	DbServerComputeModelEcpu DbServerComputeModelEnum = "ECPU"
	DbServerComputeModelOcpu DbServerComputeModelEnum = "OCPU"
)

var mappingDbServerComputeModelEnum = map[string]DbServerComputeModelEnum{
	"ECPU": DbServerComputeModelEcpu,
	"OCPU": DbServerComputeModelOcpu,
}

var mappingDbServerComputeModelEnumLowerCase = map[string]DbServerComputeModelEnum{
	"ecpu": DbServerComputeModelEcpu,
	"ocpu": DbServerComputeModelOcpu,
}

// GetDbServerComputeModelEnumValues Enumerates the set of values for DbServerComputeModelEnum
func GetDbServerComputeModelEnumValues() []DbServerComputeModelEnum {
	values := make([]DbServerComputeModelEnum, 0)
	for _, v := range mappingDbServerComputeModelEnum {
		values = append(values, v)
	}
	return values
}

// GetDbServerComputeModelEnumStringValues Enumerates the set of values in String for DbServerComputeModelEnum
func GetDbServerComputeModelEnumStringValues() []string {
	return []string{
		"ECPU",
		"OCPU",
	}
}

// GetMappingDbServerComputeModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbServerComputeModelEnum(val string) (DbServerComputeModelEnum, bool) {
	enum, ok := mappingDbServerComputeModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
