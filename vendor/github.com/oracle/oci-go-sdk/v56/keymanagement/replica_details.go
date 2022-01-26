// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Service Key Management API
//
// API for managing and performing operations with keys and vaults. (For the API for managing secrets, see the Vault Service
// Secret Management API. For the API for retrieving secrets, see the Vault Service Secret Retrieval API.)
//

package keymanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ReplicaDetails Details of replication status
type ReplicaDetails struct {

	// The replica region
	Region *string `mandatory:"false" json:"region"`

	// Replication status associated with a replicationId
	Status ReplicaDetailsStatusEnum `mandatory:"false" json:"status,omitempty"`
}

func (m ReplicaDetails) String() string {
	return common.PointerString(m)
}

// ReplicaDetailsStatusEnum Enum with underlying type: string
type ReplicaDetailsStatusEnum string

// Set of constants representing the allowable values for ReplicaDetailsStatusEnum
const (
	ReplicaDetailsStatusReplicating ReplicaDetailsStatusEnum = "REPLICATING"
	ReplicaDetailsStatusReplicated  ReplicaDetailsStatusEnum = "REPLICATED"
)

var mappingReplicaDetailsStatus = map[string]ReplicaDetailsStatusEnum{
	"REPLICATING": ReplicaDetailsStatusReplicating,
	"REPLICATED":  ReplicaDetailsStatusReplicated,
}

// GetReplicaDetailsStatusEnumValues Enumerates the set of values for ReplicaDetailsStatusEnum
func GetReplicaDetailsStatusEnumValues() []ReplicaDetailsStatusEnum {
	values := make([]ReplicaDetailsStatusEnum, 0)
	for _, v := range mappingReplicaDetailsStatus {
		values = append(values, v)
	}
	return values
}
