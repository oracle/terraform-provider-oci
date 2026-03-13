// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails Globally distributed autonomous database shard with dedicated autonomous infrastructure.
type CreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails struct {

	// Admin password for shard database.
	AdminPassword *string `mandatory:"true" json:"adminPassword"`

	// The compute count for the shard database. It has to be in multiples of 2.
	ComputeCount *float32 `mandatory:"true" json:"computeCount"`

	// The data disk group size to be allocated in GBs for the shard database.
	DataStorageSizeInGbs *float64 `mandatory:"true" json:"dataStorageSizeInGbs"`

	// Determines the auto-scaling mode for the shard database.
	IsAutoScalingEnabled *bool `mandatory:"true" json:"isAutoScalingEnabled"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
	CloudAutonomousVmClusterId *string `mandatory:"true" json:"cloudAutonomousVmClusterId"`

	// The shard space name for the shard database. Shard space for existing shard cannot be changed, once shard is created.
	// Shard space name shall be used while creation of new shards. For User defined sharding, every shard must have a unique
	// shard space name. For system defined sharding, shard space name is not required.
	ShardSpace *string `mandatory:"false" json:"shardSpace"`

	// This field is deprecated. This should not be used while creation of new distributed autonomous database. To set the peers
	// on new shards of distributed autonomous database please use peerDetails.
	PeerCloudAutonomousVmClusterIds []string `mandatory:"false" json:"peerCloudAutonomousVmClusterIds"`

	// The details required for creation of the peer for the autonomous dedicated infrastructure based shard.
	PeerDetails []CreateShardPeerWithDedicatedInfraDetails `mandatory:"false" json:"peerDetails"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the shard.
	OkvKeyStoreId *string `mandatory:"false" json:"okvKeyStoreId"`

	// The OKV endpoint name.
	OkvEndPointGroup *string `mandatory:"false" json:"okvEndPointGroup"`
}

func (m CreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails CreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails
	}{
		"ADB_D",
		(MarshalTypeCreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails)(m),
	}

	return json.Marshal(&s)
}
