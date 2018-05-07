// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// LocalPeeringTokenDetails An object containing a generated peering token to be given to a peer who then accepts the token as part of the peering handshake process.
type LocalPeeringTokenDetails struct {

	// An opaque token to be shared with a peer.
	TokenForPeer *string `mandatory:"true" json:"tokenForPeer"`
}

func (m LocalPeeringTokenDetails) String() string {
	return common.PointerString(m)
}
