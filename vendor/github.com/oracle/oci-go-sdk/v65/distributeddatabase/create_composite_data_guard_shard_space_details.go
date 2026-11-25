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

// CreateCompositeDataGuardShardSpaceDetails Globally distributed database composite data guard shard space creation details.
type CreateCompositeDataGuardShardSpaceDetails struct {

	// The name of shard space.
	// It must start with a letter and contain only letters, digits, and underscores.
	// Maximum length is 40 characters.
	Name *string `mandatory:"true" json:"name"`

	// Count of chunks associated with the shard space.
	ChunkCount *int `mandatory:"true" json:"chunkCount"`

	OriginalReplica *CreateCompositeDataGuardShardSpaceReplicaDetails `mandatory:"true" json:"originalReplica"`

	// The details of data guard replica for the shard being created.
	DataGuardReplicas []CreateCompositeDataGuardShardSpaceReplicaDetails `mandatory:"false" json:"dataGuardReplicas"`
}

func (m CreateCompositeDataGuardShardSpaceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCompositeDataGuardShardSpaceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
