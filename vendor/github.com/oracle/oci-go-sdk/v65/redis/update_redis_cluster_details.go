// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Cache API
//
// Use the OCI Cache API to create and manage clusters. A cluster is a memory-based storage solution. For more information, see OCI Cache (https://docs.oracle.com/iaas/Content/ocicache/home.htm).
//

package redis

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateRedisClusterDetails The configuration to update for an existing cluster.
type UpdateRedisClusterDetails struct {

	// The number of shards in sharded cluster. Only applicable when clusterMode is SHARDED.
	ShardCount *int `mandatory:"false" json:"shardCount"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The number of nodes per shard in the cluster when clusterMode is SHARDED. This is the total number of nodes when clusterMode is NONSHARDED.
	NodeCount *int `mandatory:"false" json:"nodeCount"`

	// The amount of memory allocated to the cluster's nodes, in gigabytes.
	NodeMemoryInGBs *float32 `mandatory:"false" json:"nodeMemoryInGBs"`

	// The OCI Cache engine version that the cluster is running.
	SoftwareVersion RedisClusterSoftwareVersionEnum `mandatory:"false" json:"softwareVersion,omitempty"`

	// A list of Network Security Group (NSG) OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
	// associated with this cluster. For more information,
	// see Using an NSG for Clusters (https://docs.oracle.com/iaas/Content/ocicache/connecttocluster.htm#connecttocluster__networksecuritygroup).
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateRedisClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateRedisClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRedisClusterSoftwareVersionEnum(string(m.SoftwareVersion)); !ok && m.SoftwareVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SoftwareVersion: %s. Supported values are: %s.", m.SoftwareVersion, strings.Join(GetRedisClusterSoftwareVersionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
