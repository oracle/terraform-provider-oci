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

// CreateDedicatedShardDetail Details required for creation of ATP-D based shard.
type CreateDedicatedShardDetail struct {

	// Admin password for shard database.
	AdminPassword *string `mandatory:"true" json:"adminPassword"`

	// The compute count for the shard database. It has to be in multiples of 2.
	ComputeCount *float32 `mandatory:"true" json:"computeCount"`

	// The data disk group size to be allocated in GBs for the shard database.
	DataStorageSizeInGbs *float64 `mandatory:"true" json:"dataStorageSizeInGbs"`

	// Determines the auto-scaling mode for the shard database.
	IsAutoScalingEnabled *bool `mandatory:"true" json:"isAutoScalingEnabled"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
	CloudAutonomousVmClusterId *string `mandatory:"true" json:"cloudAutonomousVmClusterId"`

	EncryptionKeyDetails *DedicatedShardOrCatalogEncryptionKeyDetails `mandatory:"false" json:"encryptionKeyDetails"`

	// The shard space name for the shard database. Shard space for existing shard cannot be changed, once shard is created.
	// Shard space name shall be used while creation of new shards. For User defined sharding, every shard must have a unique
	// shard space name. For system defined sharding, shard space name is not required.
	ShardSpace *string `mandatory:"false" json:"shardSpace"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the peer cloud Autonomous Exadata VM Cluster.
	PeerCloudAutonomousVmClusterId *string `mandatory:"false" json:"peerCloudAutonomousVmClusterId"`
}

func (m CreateDedicatedShardDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDedicatedShardDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
