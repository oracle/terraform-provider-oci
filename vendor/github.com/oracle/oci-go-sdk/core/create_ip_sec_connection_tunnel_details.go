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

// CreateIpSecConnectionTunnelDetails details need to create an IPSecConnection tunnel.
type CreateIpSecConnectionTunnelDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// the routing strategy used for this tunnel, either static route or BGP.
	Routing CreateIpSecConnectionTunnelDetailsRoutingEnum `mandatory:"false" json:"routing,omitempty"`

	// The shared secret of the IPSec tunnel.
	// Example: `vFG2IF6TWq4UToUiLSRDoJEUs6j1c.p8G.dVQxiMfMO0yXMLi.lZTbYIWhGu4V8o`
	SharedSecret *string `mandatory:"false" json:"sharedSecret"`

	// Information needed to establish a BGP Session on an interface.
	BgpSessionConfig *CreateIpSecTunnelBgpSessionDetails `mandatory:"false" json:"bgpSessionConfig"`
}

func (m CreateIpSecConnectionTunnelDetails) String() string {
	return common.PointerString(m)
}

// CreateIpSecConnectionTunnelDetailsRoutingEnum Enum with underlying type: string
type CreateIpSecConnectionTunnelDetailsRoutingEnum string

// Set of constants representing the allowable values for CreateIpSecConnectionTunnelDetailsRoutingEnum
const (
	CreateIpSecConnectionTunnelDetailsRoutingBgp    CreateIpSecConnectionTunnelDetailsRoutingEnum = "BGP"
	CreateIpSecConnectionTunnelDetailsRoutingStatic CreateIpSecConnectionTunnelDetailsRoutingEnum = "STATIC"
)

var mappingCreateIpSecConnectionTunnelDetailsRouting = map[string]CreateIpSecConnectionTunnelDetailsRoutingEnum{
	"BGP":    CreateIpSecConnectionTunnelDetailsRoutingBgp,
	"STATIC": CreateIpSecConnectionTunnelDetailsRoutingStatic,
}

// GetCreateIpSecConnectionTunnelDetailsRoutingEnumValues Enumerates the set of values for CreateIpSecConnectionTunnelDetailsRoutingEnum
func GetCreateIpSecConnectionTunnelDetailsRoutingEnumValues() []CreateIpSecConnectionTunnelDetailsRoutingEnum {
	values := make([]CreateIpSecConnectionTunnelDetailsRoutingEnum, 0)
	for _, v := range mappingCreateIpSecConnectionTunnelDetailsRouting {
		values = append(values, v)
	}
	return values
}
