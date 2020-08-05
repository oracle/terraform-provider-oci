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

// CreatePeerDetails The Peer details to be added
type CreatePeerDetails struct {

	// Peer role
	Role PeerRoleRoleEnum `mandatory:"true" json:"role"`

	OcpuAllocationParam *OcpuAllocationNumberParam `mandatory:"true" json:"ocpuAllocationParam"`

	// Availability Domain to place new peer
	Ad AvailabilityDomainAdsEnum `mandatory:"true" json:"ad"`

	// peer alias
	Alias *string `mandatory:"false" json:"alias"`
}

func (m CreatePeerDetails) String() string {
	return common.PointerString(m)
}
