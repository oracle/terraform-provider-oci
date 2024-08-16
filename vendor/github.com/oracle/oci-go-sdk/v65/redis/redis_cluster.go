// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Cache API
//
// Use the OCI Cache API to create and manage clusters. A cluster is a memory-based storage solution. For more information, see OCI Cache (https://docs.cloud.oracle.com/iaas/Content/ocicache/home.htm).
//

package redis

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RedisCluster An OCI Cache cluster is a memory-based storage solution. For more information, see OCI Cache (https://docs.cloud.oracle.com/iaas/Content/ocicache/home.htm).
type RedisCluster struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the cluster.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the compartment that contains the cluster.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The number of nodes per shard in the cluster when clusterMode is SHARDED. This is the total number of nodes when clusterMode is NONSHARDED.
	NodeCount *int `mandatory:"true" json:"nodeCount"`

	// The amount of memory allocated to the cluster's nodes, in gigabytes.
	NodeMemoryInGBs *float32 `mandatory:"true" json:"nodeMemoryInGBs"`

	// The fully qualified domain name (FQDN) of the API endpoint for the cluster's primary node.
	PrimaryFqdn *string `mandatory:"true" json:"primaryFqdn"`

	// The private IP address of the API endpoint for the cluster's primary node.
	PrimaryEndpointIpAddress *string `mandatory:"true" json:"primaryEndpointIpAddress"`

	// The fully qualified domain name (FQDN) of the API endpoint for the cluster's replica nodes.
	ReplicasFqdn *string `mandatory:"true" json:"replicasFqdn"`

	// The private IP address of the API endpoint for the cluster's replica nodes.
	ReplicasEndpointIpAddress *string `mandatory:"true" json:"replicasEndpointIpAddress"`

	// The OCI Cache engine version that the cluster is running.
	SoftwareVersion RedisClusterSoftwareVersionEnum `mandatory:"true" json:"softwareVersion"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the cluster's subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	NodeCollection *NodeCollection `mandatory:"true" json:"nodeCollection"`

	// The current state of the cluster.
	LifecycleState RedisClusterLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, the message might provide actionable information for a resource in `FAILED` state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the cluster was created. An RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the cluster was updated. An RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Specifies whether the cluster is sharded or non-sharded.
	ClusterMode RedisClusterClusterModeEnum `mandatory:"false" json:"clusterMode,omitempty"`

	// The number of shards in a sharded cluster. Only applicable when clusterMode is SHARDED.
	ShardCount *int `mandatory:"false" json:"shardCount"`

	// A list of Network Security Group (NSG) OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
	// associated with this cluster. For more information,
	// see Using an NSG for Clusters (https://docs.cloud.oracle.com/iaas/Content/ocicache/connecttocluster.htm#connecttocluster__networksecuritygroup).
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m RedisCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RedisCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRedisClusterSoftwareVersionEnum(string(m.SoftwareVersion)); !ok && m.SoftwareVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SoftwareVersion: %s. Supported values are: %s.", m.SoftwareVersion, strings.Join(GetRedisClusterSoftwareVersionEnumStringValues(), ",")))
	}

	if _, ok := GetMappingRedisClusterLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetRedisClusterLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRedisClusterClusterModeEnum(string(m.ClusterMode)); !ok && m.ClusterMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClusterMode: %s. Supported values are: %s.", m.ClusterMode, strings.Join(GetRedisClusterClusterModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RedisClusterLifecycleStateEnum Enum with underlying type: string
type RedisClusterLifecycleStateEnum string

// Set of constants representing the allowable values for RedisClusterLifecycleStateEnum
const (
	RedisClusterLifecycleStateCreating RedisClusterLifecycleStateEnum = "CREATING"
	RedisClusterLifecycleStateUpdating RedisClusterLifecycleStateEnum = "UPDATING"
	RedisClusterLifecycleStateActive   RedisClusterLifecycleStateEnum = "ACTIVE"
	RedisClusterLifecycleStateDeleting RedisClusterLifecycleStateEnum = "DELETING"
	RedisClusterLifecycleStateDeleted  RedisClusterLifecycleStateEnum = "DELETED"
	RedisClusterLifecycleStateFailed   RedisClusterLifecycleStateEnum = "FAILED"
)

var mappingRedisClusterLifecycleStateEnum = map[string]RedisClusterLifecycleStateEnum{
	"CREATING": RedisClusterLifecycleStateCreating,
	"UPDATING": RedisClusterLifecycleStateUpdating,
	"ACTIVE":   RedisClusterLifecycleStateActive,
	"DELETING": RedisClusterLifecycleStateDeleting,
	"DELETED":  RedisClusterLifecycleStateDeleted,
	"FAILED":   RedisClusterLifecycleStateFailed,
}

var mappingRedisClusterLifecycleStateEnumLowerCase = map[string]RedisClusterLifecycleStateEnum{
	"creating": RedisClusterLifecycleStateCreating,
	"updating": RedisClusterLifecycleStateUpdating,
	"active":   RedisClusterLifecycleStateActive,
	"deleting": RedisClusterLifecycleStateDeleting,
	"deleted":  RedisClusterLifecycleStateDeleted,
	"failed":   RedisClusterLifecycleStateFailed,
}

// GetRedisClusterLifecycleStateEnumValues Enumerates the set of values for RedisClusterLifecycleStateEnum
func GetRedisClusterLifecycleStateEnumValues() []RedisClusterLifecycleStateEnum {
	values := make([]RedisClusterLifecycleStateEnum, 0)
	for _, v := range mappingRedisClusterLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetRedisClusterLifecycleStateEnumStringValues Enumerates the set of values in String for RedisClusterLifecycleStateEnum
func GetRedisClusterLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingRedisClusterLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRedisClusterLifecycleStateEnum(val string) (RedisClusterLifecycleStateEnum, bool) {
	enum, ok := mappingRedisClusterLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RedisClusterSoftwareVersionEnum Enum with underlying type: string
type RedisClusterSoftwareVersionEnum string

// Set of constants representing the allowable values for RedisClusterSoftwareVersionEnum
const (
	RedisClusterSoftwareVersionV705    RedisClusterSoftwareVersionEnum = "V7_0_5"
	RedisClusterSoftwareVersionRedis70 RedisClusterSoftwareVersionEnum = "REDIS_7_0"
)

var mappingRedisClusterSoftwareVersionEnum = map[string]RedisClusterSoftwareVersionEnum{
	"V7_0_5":    RedisClusterSoftwareVersionV705,
	"REDIS_7_0": RedisClusterSoftwareVersionRedis70,
}

var mappingRedisClusterSoftwareVersionEnumLowerCase = map[string]RedisClusterSoftwareVersionEnum{
	"v7_0_5":    RedisClusterSoftwareVersionV705,
	"redis_7_0": RedisClusterSoftwareVersionRedis70,
}

// GetRedisClusterSoftwareVersionEnumValues Enumerates the set of values for RedisClusterSoftwareVersionEnum
func GetRedisClusterSoftwareVersionEnumValues() []RedisClusterSoftwareVersionEnum {
	values := make([]RedisClusterSoftwareVersionEnum, 0)
	for _, v := range mappingRedisClusterSoftwareVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetRedisClusterSoftwareVersionEnumStringValues Enumerates the set of values in String for RedisClusterSoftwareVersionEnum
func GetRedisClusterSoftwareVersionEnumStringValues() []string {
	return []string{
		"V7_0_5",
		"REDIS_7_0",
	}
}

// GetMappingRedisClusterSoftwareVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRedisClusterSoftwareVersionEnum(val string) (RedisClusterSoftwareVersionEnum, bool) {
	enum, ok := mappingRedisClusterSoftwareVersionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RedisClusterClusterModeEnum Enum with underlying type: string
type RedisClusterClusterModeEnum string

// Set of constants representing the allowable values for RedisClusterClusterModeEnum
const (
	RedisClusterClusterModeSharded    RedisClusterClusterModeEnum = "SHARDED"
	RedisClusterClusterModeNonsharded RedisClusterClusterModeEnum = "NONSHARDED"
)

var mappingRedisClusterClusterModeEnum = map[string]RedisClusterClusterModeEnum{
	"SHARDED":    RedisClusterClusterModeSharded,
	"NONSHARDED": RedisClusterClusterModeNonsharded,
}

var mappingRedisClusterClusterModeEnumLowerCase = map[string]RedisClusterClusterModeEnum{
	"sharded":    RedisClusterClusterModeSharded,
	"nonsharded": RedisClusterClusterModeNonsharded,
}

// GetRedisClusterClusterModeEnumValues Enumerates the set of values for RedisClusterClusterModeEnum
func GetRedisClusterClusterModeEnumValues() []RedisClusterClusterModeEnum {
	values := make([]RedisClusterClusterModeEnum, 0)
	for _, v := range mappingRedisClusterClusterModeEnum {
		values = append(values, v)
	}
	return values
}

// GetRedisClusterClusterModeEnumStringValues Enumerates the set of values in String for RedisClusterClusterModeEnum
func GetRedisClusterClusterModeEnumStringValues() []string {
	return []string{
		"SHARDED",
		"NONSHARDED",
	}
}

// GetMappingRedisClusterClusterModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRedisClusterClusterModeEnum(val string) (RedisClusterClusterModeEnum, bool) {
	enum, ok := mappingRedisClusterClusterModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
