// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DbServerSummary Details of the Exadata Cloud@Customer Db server.
type DbServerSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exacc Db server.
	Id *string `mandatory:"false" json:"id"`

	// The user-friendly name for the Db server. The name does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
	ExadataInfrastructureId *string `mandatory:"false" json:"exadataInfrastructureId"`

	// The number of CPU cores enabled on the Db server.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// The allocated memory in GBs on the Db server.
	MemorySizeInGBs *int `mandatory:"false" json:"memorySizeInGBs"`

	// The allocated local node storage in GBs on the Db server.
	DbNodeStorageSizeInGBs *int `mandatory:"false" json:"dbNodeStorageSizeInGBs"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VM Clusters associated with the Db server.
	VmClusterIds []string `mandatory:"false" json:"vmClusterIds"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Db nodes associated with the Db server.
	DbNodeIds []string `mandatory:"false" json:"dbNodeIds"`

	// The current state of the Db server.
	LifecycleState DbServerSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

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

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m DbServerSummary) String() string {
	return common.PointerString(m)
}

// DbServerSummaryLifecycleStateEnum Enum with underlying type: string
type DbServerSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for DbServerSummaryLifecycleStateEnum
const (
	DbServerSummaryLifecycleStateCreating              DbServerSummaryLifecycleStateEnum = "CREATING"
	DbServerSummaryLifecycleStateAvailable             DbServerSummaryLifecycleStateEnum = "AVAILABLE"
	DbServerSummaryLifecycleStateUnavailable           DbServerSummaryLifecycleStateEnum = "UNAVAILABLE"
	DbServerSummaryLifecycleStateDeleting              DbServerSummaryLifecycleStateEnum = "DELETING"
	DbServerSummaryLifecycleStateDeleted               DbServerSummaryLifecycleStateEnum = "DELETED"
	DbServerSummaryLifecycleStateMaintenanceInProgress DbServerSummaryLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
)

var mappingDbServerSummaryLifecycleState = map[string]DbServerSummaryLifecycleStateEnum{
	"CREATING":                DbServerSummaryLifecycleStateCreating,
	"AVAILABLE":               DbServerSummaryLifecycleStateAvailable,
	"UNAVAILABLE":             DbServerSummaryLifecycleStateUnavailable,
	"DELETING":                DbServerSummaryLifecycleStateDeleting,
	"DELETED":                 DbServerSummaryLifecycleStateDeleted,
	"MAINTENANCE_IN_PROGRESS": DbServerSummaryLifecycleStateMaintenanceInProgress,
}

// GetDbServerSummaryLifecycleStateEnumValues Enumerates the set of values for DbServerSummaryLifecycleStateEnum
func GetDbServerSummaryLifecycleStateEnumValues() []DbServerSummaryLifecycleStateEnum {
	values := make([]DbServerSummaryLifecycleStateEnum, 0)
	for _, v := range mappingDbServerSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
