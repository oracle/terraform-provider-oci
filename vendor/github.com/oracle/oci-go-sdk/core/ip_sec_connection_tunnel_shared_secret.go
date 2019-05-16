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

// IpSecConnectionTunnelSharedSecret The shared secret of a IPSec connection's specified tunnel.
type IpSecConnectionTunnelSharedSecret struct {

	// The shared secret of the IPSec tunnel.
	// Example: `vFG2IF6TWq4UToUiLSRDoJEUs6j1c.p8G.dVQxiMfMO0yXMLi.lZTbYIWhGu4V8o`
	SharedSecret *string `mandatory:"true" json:"sharedSecret"`
}

func (m IpSecConnectionTunnelSharedSecret) String() string {
	return common.PointerString(m)
}
