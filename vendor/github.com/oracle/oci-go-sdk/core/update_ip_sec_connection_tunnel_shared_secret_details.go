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

// UpdateIpSecConnectionTunnelSharedSecretDetails The representation of UpdateIpSecConnectionTunnelSharedSecretDetails
type UpdateIpSecConnectionTunnelSharedSecretDetails struct {

	// The shared secret of the IPSec tunnel.
	// Example: `vFG2IF6TWq4UToUiLSRDoJEUs6j1c.p8G.dVQxiMfMO0yXMLi.lZTbYIWhGu4V8o`
	SharedSecret *string `mandatory:"false" json:"sharedSecret"`
}

func (m UpdateIpSecConnectionTunnelSharedSecretDetails) String() string {
	return common.PointerString(m)
}
