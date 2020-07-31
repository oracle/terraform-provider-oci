// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Control Plane API
//

package blockchain

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Peer A Peer details
type Peer struct {

	// peer identifier
	PeerKey *string `mandatory:"true" json:"peerKey"`

	// Peer role
	Role PeerRoleRoleEnum `mandatory:"true" json:"role"`

	// Host on which the Peer exists
	Host *string `mandatory:"true" json:"host"`

	// Availability Domain of peer
	Ad AvailabilityDomainAdsEnum `mandatory:"true" json:"ad"`

	// peer alias
	Alias *string `mandatory:"false" json:"alias"`

	OcpuAllocationParam *OcpuAllocationNumberParam `mandatory:"false" json:"ocpuAllocationParam"`

	// The current state of the peer.
	LifecycleState PeerLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m Peer) String() string {
	return common.PointerString(m)
}

// PeerLifecycleStateEnum Enum with underlying type: string
type PeerLifecycleStateEnum string

// Set of constants representing the allowable values for PeerLifecycleStateEnum
const (
	PeerLifecycleStateActive   PeerLifecycleStateEnum = "ACTIVE"
	PeerLifecycleStateInactive PeerLifecycleStateEnum = "INACTIVE"
	PeerLifecycleStateFailed   PeerLifecycleStateEnum = "FAILED"
)

var mappingPeerLifecycleState = map[string]PeerLifecycleStateEnum{
	"ACTIVE":   PeerLifecycleStateActive,
	"INACTIVE": PeerLifecycleStateInactive,
	"FAILED":   PeerLifecycleStateFailed,
}

// GetPeerLifecycleStateEnumValues Enumerates the set of values for PeerLifecycleStateEnum
func GetPeerLifecycleStateEnumValues() []PeerLifecycleStateEnum {
	values := make([]PeerLifecycleStateEnum, 0)
	for _, v := range mappingPeerLifecycleState {
		values = append(values, v)
	}
	return values
}
