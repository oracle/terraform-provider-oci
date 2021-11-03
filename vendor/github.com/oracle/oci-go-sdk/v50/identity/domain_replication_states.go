// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// DomainReplicationStates Domain replication replication log for all domains for a given region
type DomainReplicationStates struct {

	// The OCID of the domain
	DomainId *string `mandatory:"true" json:"domainId"`

	// The IDCS replicated region state
	State ReplicatedRegionDetailsStateEnum `mandatory:"true" json:"state"`

	// The replica region for domain.
	ReplicaRegion *string `mandatory:"true" json:"replicaRegion"`
}

func (m DomainReplicationStates) String() string {
	return common.PointerString(m)
}
