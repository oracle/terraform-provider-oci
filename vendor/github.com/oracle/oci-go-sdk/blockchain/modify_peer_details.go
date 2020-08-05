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

// ModifyPeerDetails peer to modify ocpu allocation
type ModifyPeerDetails struct {

	// peer identifier
	PeerName *string `mandatory:"true" json:"peerName"`

	OcpuAllocationParam *OcpuAllocationNumberParam `mandatory:"true" json:"ocpuAllocationParam"`
}

func (m ModifyPeerDetails) String() string {
	return common.PointerString(m)
}
