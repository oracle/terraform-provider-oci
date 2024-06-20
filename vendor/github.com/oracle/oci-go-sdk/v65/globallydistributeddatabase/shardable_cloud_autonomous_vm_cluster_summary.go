// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage distributed databases.
//

package globallydistributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ShardableCloudAutonomousVmClusterSummary Shardable cloud autonomous vm cluster summary.
type ShardableCloudAutonomousVmClusterSummary struct {

	// Cloud autonomous vmcluster identifier
	Id *string `mandatory:"true" json:"id"`

	// Cloud autonomous vmcluster compartment id
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Lifecycle states for shardable Cloud autonomous vm cluster.
	LifecycleState ShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Cloud autonomous vmcluster displayName
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Detailed message for the lifecycle state.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// The compute model of the Cloud Autonomous VM Cluster.
	ComputeModel *string `mandatory:"false" json:"computeModel"`

	// The number of Autonomous Container Databases that can be created with the currently available local storage.
	AvailableContainerDatabases *int `mandatory:"false" json:"availableContainerDatabases"`

	// CPU cores available for allocation to Autonomous Databases.
	AvailableCpus *float32 `mandatory:"false" json:"availableCpus"`

	// The name of the availability domain that the cloud Autonomous VM cluster is located in.
	// The format of the availability domain is the same as returned by Cloud Autonomous VM Cluster API.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The data disk group size allocated for Autonomous Databases, in TBs.
	AutonomousDataStorageSizeInTBs *float32 `mandatory:"false" json:"autonomousDataStorageSizeInTBs"`

	// The data disk group size available for Autonomous Databases, in TBs.
	AvailableAutonomousDataStorageSizeInTBs *float32 `mandatory:"false" json:"availableAutonomousDataStorageSizeInTBs"`

	// Cloud Exadata Infrastructure Identifier.
	CloudExadataInfrastructureId *string `mandatory:"false" json:"cloudExadataInfrastructureId"`

	// The time zone of the Cloud Autonomous VM Cluster.
	ClusterTimeZone *string `mandatory:"false" json:"clusterTimeZone"`

	// The total number of Autonomous Container Databases that can be created with the allocated local storage.
	TotalContainerDatabases *int `mandatory:"false" json:"totalContainerDatabases"`

	// Cloud autonomous vmcluster subnet id
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// Cloud autonomous vmcluster shape
	Shape *string `mandatory:"false" json:"shape"`

	// Cloud autonomous vmcluster node count
	NodeCount *int `mandatory:"false" json:"nodeCount"`

	// The Oracle license model that applies to the Oracle Autonomous Database.
	LicenseModel *string `mandatory:"false" json:"licenseModel"`

	// The memory allocated in GBs.
	MemorySizeInGBs *int `mandatory:"false" json:"memorySizeInGBs"`

	// The amount of memory (in GBs) enabled per OCPU or ECPU.
	MemoryPerOracleComputeUnitInGBs *int `mandatory:"false" json:"memoryPerOracleComputeUnitInGBs"`

	// The number of CPU cores on the cloud Autonomous VM cluster.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// The number of CPU cores enabled per VM cluster node.
	CpuCoreCountPerNode *int `mandatory:"false" json:"cpuCoreCountPerNode"`

	// The number of CPU cores on the cloud Autonomous VM cluster.
	OcpuCount *float32 `mandatory:"false" json:"ocpuCount"`

	// The CPUs that continue to be included in the count of CPUs available to the Autonomous Container Database even after one of its Autonomous Database is terminated or scaled down. You can release them to the available CPUs at its parent Autonomous VM Cluster level by restarting the Autonomous Container Database.
	ReclaimableCpus *float32 `mandatory:"false" json:"reclaimableCpus"`

	// Number of Autonomous Container Databases that can be created in the Autonomous VM Cluster
	ProvisionableAutonomousContainerDatabases *int `mandatory:"false" json:"provisionableAutonomousContainerDatabases"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ShardableCloudAutonomousVmClusterSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ShardableCloudAutonomousVmClusterSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetShardableCloudAutonomousVmClusterSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum Enum with underlying type: string
type ShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum
const (
	ShardableCloudAutonomousVmClusterSummaryLifecycleStateActive         ShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum = "ACTIVE"
	ShardableCloudAutonomousVmClusterSummaryLifecycleStateFailed         ShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum = "FAILED"
	ShardableCloudAutonomousVmClusterSummaryLifecycleStateNeedsAttention ShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum = "NEEDS_ATTENTION"
	ShardableCloudAutonomousVmClusterSummaryLifecycleStateInactive       ShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum = "INACTIVE"
	ShardableCloudAutonomousVmClusterSummaryLifecycleStateDeleting       ShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum = "DELETING"
	ShardableCloudAutonomousVmClusterSummaryLifecycleStateDeleted        ShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum = "DELETED"
	ShardableCloudAutonomousVmClusterSummaryLifecycleStateUpdating       ShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum = "UPDATING"
	ShardableCloudAutonomousVmClusterSummaryLifecycleStateCreating       ShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum = "CREATING"
	ShardableCloudAutonomousVmClusterSummaryLifecycleStateUnavailable    ShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum = "UNAVAILABLE"
)

var mappingShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum = map[string]ShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum{
	"ACTIVE":          ShardableCloudAutonomousVmClusterSummaryLifecycleStateActive,
	"FAILED":          ShardableCloudAutonomousVmClusterSummaryLifecycleStateFailed,
	"NEEDS_ATTENTION": ShardableCloudAutonomousVmClusterSummaryLifecycleStateNeedsAttention,
	"INACTIVE":        ShardableCloudAutonomousVmClusterSummaryLifecycleStateInactive,
	"DELETING":        ShardableCloudAutonomousVmClusterSummaryLifecycleStateDeleting,
	"DELETED":         ShardableCloudAutonomousVmClusterSummaryLifecycleStateDeleted,
	"UPDATING":        ShardableCloudAutonomousVmClusterSummaryLifecycleStateUpdating,
	"CREATING":        ShardableCloudAutonomousVmClusterSummaryLifecycleStateCreating,
	"UNAVAILABLE":     ShardableCloudAutonomousVmClusterSummaryLifecycleStateUnavailable,
}

var mappingShardableCloudAutonomousVmClusterSummaryLifecycleStateEnumLowerCase = map[string]ShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum{
	"active":          ShardableCloudAutonomousVmClusterSummaryLifecycleStateActive,
	"failed":          ShardableCloudAutonomousVmClusterSummaryLifecycleStateFailed,
	"needs_attention": ShardableCloudAutonomousVmClusterSummaryLifecycleStateNeedsAttention,
	"inactive":        ShardableCloudAutonomousVmClusterSummaryLifecycleStateInactive,
	"deleting":        ShardableCloudAutonomousVmClusterSummaryLifecycleStateDeleting,
	"deleted":         ShardableCloudAutonomousVmClusterSummaryLifecycleStateDeleted,
	"updating":        ShardableCloudAutonomousVmClusterSummaryLifecycleStateUpdating,
	"creating":        ShardableCloudAutonomousVmClusterSummaryLifecycleStateCreating,
	"unavailable":     ShardableCloudAutonomousVmClusterSummaryLifecycleStateUnavailable,
}

// GetShardableCloudAutonomousVmClusterSummaryLifecycleStateEnumValues Enumerates the set of values for ShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum
func GetShardableCloudAutonomousVmClusterSummaryLifecycleStateEnumValues() []ShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum {
	values := make([]ShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum, 0)
	for _, v := range mappingShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetShardableCloudAutonomousVmClusterSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum
func GetShardableCloudAutonomousVmClusterSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"FAILED",
		"NEEDS_ATTENTION",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"UPDATING",
		"CREATING",
		"UNAVAILABLE",
	}
}

// GetMappingShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum(val string) (ShardableCloudAutonomousVmClusterSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingShardableCloudAutonomousVmClusterSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
