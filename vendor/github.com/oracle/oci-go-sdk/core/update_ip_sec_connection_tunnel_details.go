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

// UpdateIpSecConnectionTunnelDetails Details to modify an IPSecConnection Tunnel paramaters.
type UpdateIpSecConnectionTunnelDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// the routing strategy used for this tunnel, either static route or BGP.
	Routing UpdateIpSecConnectionTunnelDetailsRoutingEnum `mandatory:"false" json:"routing,omitempty"`

	// Information needed to establish a BGP Session on an interface.
	BgpSessionConfig *UpdateIpSecTunnelBgpSessionDetails `mandatory:"false" json:"bgpSessionConfig"`
}

func (m UpdateIpSecConnectionTunnelDetails) String() string {
	return common.PointerString(m)
}

// UpdateIpSecConnectionTunnelDetailsRoutingEnum Enum with underlying type: string
type UpdateIpSecConnectionTunnelDetailsRoutingEnum string

// Set of constants representing the allowable values for UpdateIpSecConnectionTunnelDetailsRoutingEnum
const (
	UpdateIpSecConnectionTunnelDetailsRoutingBgp    UpdateIpSecConnectionTunnelDetailsRoutingEnum = "BGP"
	UpdateIpSecConnectionTunnelDetailsRoutingStatic UpdateIpSecConnectionTunnelDetailsRoutingEnum = "STATIC"
)

var mappingUpdateIpSecConnectionTunnelDetailsRouting = map[string]UpdateIpSecConnectionTunnelDetailsRoutingEnum{
	"BGP":    UpdateIpSecConnectionTunnelDetailsRoutingBgp,
	"STATIC": UpdateIpSecConnectionTunnelDetailsRoutingStatic,
}

// GetUpdateIpSecConnectionTunnelDetailsRoutingEnumValues Enumerates the set of values for UpdateIpSecConnectionTunnelDetailsRoutingEnum
func GetUpdateIpSecConnectionTunnelDetailsRoutingEnumValues() []UpdateIpSecConnectionTunnelDetailsRoutingEnum {
	values := make([]UpdateIpSecConnectionTunnelDetailsRoutingEnum, 0)
	for _, v := range mappingUpdateIpSecConnectionTunnelDetailsRouting {
		values = append(values, v)
	}
	return values
}
