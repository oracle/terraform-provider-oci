// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MoveDistributedAutonomousDatabaseReplicationUnitDetails The details for moving replication units from source shard to destination shard for the Globally distributed autonomous database.
type MoveDistributedAutonomousDatabaseReplicationUnitDetails struct {

	// The name of the source shard from which to move the chunks out to other shards.
	SourceShardName *string `mandatory:"true" json:"sourceShardName"`

	// The name of the destination shard to which the chunks moved out from source shard should be relocate to.
	DestinationShardName *string `mandatory:"false" json:"destinationShardName"`

	// For RAFT databases please provide replication unit numbers to be moved from source shard to destination shard.
	ReplicationUnits []int `mandatory:"false" json:"replicationUnits"`
}

func (m MoveDistributedAutonomousDatabaseReplicationUnitDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MoveDistributedAutonomousDatabaseReplicationUnitDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
