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

// PeerSummary Peer summary information for returning in a list.
type PeerSummary struct {

	// Peer identifier
	PeerKey *string `mandatory:"false" json:"peerKey"`

	// The current state of the peer.
	LifecycleState PeerLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m PeerSummary) String() string {
	return common.PointerString(m)
}
