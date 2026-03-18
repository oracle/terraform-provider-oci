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

// CreateDistributedDatabaseShardWithExadbXsDetails Globally distributed database shard based on exadbxs.
type CreateDistributedDatabaseShardWithExadbXsDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VmCluster.
	VmClusterId *string `mandatory:"true" json:"vmClusterId"`

	// The admin password for the shard associated with Globally distributed database.
	AdminPassword *string `mandatory:"true" json:"adminPassword"`

	// This field is deprecated. This should not be used while creation of new distributed database. To set the peers
	// on new shards of distributed database please use peerDetails.
	PeerVmClusterIds []string `mandatory:"false" json:"peerVmClusterIds"`

	// The details required for creation of the peer for the ExadbXs infrastructure based shard.
	PeerDetails []CreateShardPeerWithExadbXsDetails `mandatory:"false" json:"peerDetails"`

	// The shard space name for the Globally distributed database. Shard space for existing shard cannot be changed, once shard is created.
	// Shard space name shall be used while creation of new shards.
	ShardSpace *string `mandatory:"false" json:"shardSpace"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`
}

func (m CreateDistributedDatabaseShardWithExadbXsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDistributedDatabaseShardWithExadbXsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDistributedDatabaseShardWithExadbXsDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDistributedDatabaseShardWithExadbXsDetails CreateDistributedDatabaseShardWithExadbXsDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateDistributedDatabaseShardWithExadbXsDetails
	}{
		"EXADB_XS",
		(MarshalTypeCreateDistributedDatabaseShardWithExadbXsDetails)(m),
	}

	return json.Marshal(&s)
}
