// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, policies, and identity domains.
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DomainReplicationStates (For tenancies that support identity domains) The identity domain replication log for all identity domains for a given region.
type DomainReplicationStates struct {

	// The OCID of the identity domain.
	DomainId *string `mandatory:"true" json:"domainId"`

	// The IDCS-replicated region state.
	State ReplicatedRegionDetailsStateEnum `mandatory:"true" json:"state"`

	// The replica region for the identity domain.
	ReplicaRegion *string `mandatory:"true" json:"replicaRegion"`
}

func (m DomainReplicationStates) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DomainReplicationStates) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReplicatedRegionDetailsStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetReplicatedRegionDetailsStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
