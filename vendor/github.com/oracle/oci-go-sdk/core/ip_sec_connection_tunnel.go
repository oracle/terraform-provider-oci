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

// IpSecConnectionTunnel information about IPSecConnection tunnel.
type IpSecConnectionTunnel struct {

	// The OCID of the compartment containing the tunnel.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The tunnel's Oracle ID (OCID).
	Id *string `mandatory:"true" json:"id"`

	// The IPSec connection's tunnel's lifecycle state.
	LifecycleState IpSecConnectionTunnelLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The IP address of Oracle's VPN headend.
	// Example: `129.146.17.50`
	VpnIp *string `mandatory:"false" json:"vpnIp"`

	// The IP address of Cpe headend.
	// Example: `129.146.17.50`
	CpeIp *string `mandatory:"false" json:"cpeIp"`

	// The tunnel's current state.
	Status IpSecConnectionTunnelStatusEnum `mandatory:"false" json:"status,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Information needed to establish a BGP Session on an interface.
	BgpSessionInfo *BgpSessionInfo `mandatory:"false" json:"bgpSessionInfo"`

	// the routing strategy used for this tunnel, either static route or BGP dynamic routing
	Routing IpSecConnectionTunnelRoutingEnum `mandatory:"false" json:"routing,omitempty"`

	// The date and time the IPSec connection tunnel was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// When the status of the tunnel last changed, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeStatusUpdated *common.SDKTime `mandatory:"false" json:"timeStatusUpdated"`
}

func (m IpSecConnectionTunnel) String() string {
	return common.PointerString(m)
}

// IpSecConnectionTunnelStatusEnum Enum with underlying type: string
type IpSecConnectionTunnelStatusEnum string

// Set of constants representing the allowable values for IpSecConnectionTunnelStatusEnum
const (
	IpSecConnectionTunnelStatusUp                 IpSecConnectionTunnelStatusEnum = "UP"
	IpSecConnectionTunnelStatusDown               IpSecConnectionTunnelStatusEnum = "DOWN"
	IpSecConnectionTunnelStatusDownForMaintenance IpSecConnectionTunnelStatusEnum = "DOWN_FOR_MAINTENANCE"
)

var mappingIpSecConnectionTunnelStatus = map[string]IpSecConnectionTunnelStatusEnum{
	"UP":                   IpSecConnectionTunnelStatusUp,
	"DOWN":                 IpSecConnectionTunnelStatusDown,
	"DOWN_FOR_MAINTENANCE": IpSecConnectionTunnelStatusDownForMaintenance,
}

// GetIpSecConnectionTunnelStatusEnumValues Enumerates the set of values for IpSecConnectionTunnelStatusEnum
func GetIpSecConnectionTunnelStatusEnumValues() []IpSecConnectionTunnelStatusEnum {
	values := make([]IpSecConnectionTunnelStatusEnum, 0)
	for _, v := range mappingIpSecConnectionTunnelStatus {
		values = append(values, v)
	}
	return values
}

// IpSecConnectionTunnelLifecycleStateEnum Enum with underlying type: string
type IpSecConnectionTunnelLifecycleStateEnum string

// Set of constants representing the allowable values for IpSecConnectionTunnelLifecycleStateEnum
const (
	IpSecConnectionTunnelLifecycleStateProvisioning IpSecConnectionTunnelLifecycleStateEnum = "PROVISIONING"
	IpSecConnectionTunnelLifecycleStateAvailable    IpSecConnectionTunnelLifecycleStateEnum = "AVAILABLE"
	IpSecConnectionTunnelLifecycleStateTerminating  IpSecConnectionTunnelLifecycleStateEnum = "TERMINATING"
	IpSecConnectionTunnelLifecycleStateTerminated   IpSecConnectionTunnelLifecycleStateEnum = "TERMINATED"
)

var mappingIpSecConnectionTunnelLifecycleState = map[string]IpSecConnectionTunnelLifecycleStateEnum{
	"PROVISIONING": IpSecConnectionTunnelLifecycleStateProvisioning,
	"AVAILABLE":    IpSecConnectionTunnelLifecycleStateAvailable,
	"TERMINATING":  IpSecConnectionTunnelLifecycleStateTerminating,
	"TERMINATED":   IpSecConnectionTunnelLifecycleStateTerminated,
}

// GetIpSecConnectionTunnelLifecycleStateEnumValues Enumerates the set of values for IpSecConnectionTunnelLifecycleStateEnum
func GetIpSecConnectionTunnelLifecycleStateEnumValues() []IpSecConnectionTunnelLifecycleStateEnum {
	values := make([]IpSecConnectionTunnelLifecycleStateEnum, 0)
	for _, v := range mappingIpSecConnectionTunnelLifecycleState {
		values = append(values, v)
	}
	return values
}

// IpSecConnectionTunnelRoutingEnum Enum with underlying type: string
type IpSecConnectionTunnelRoutingEnum string

// Set of constants representing the allowable values for IpSecConnectionTunnelRoutingEnum
const (
	IpSecConnectionTunnelRoutingBgp    IpSecConnectionTunnelRoutingEnum = "BGP"
	IpSecConnectionTunnelRoutingStatic IpSecConnectionTunnelRoutingEnum = "STATIC"
)

var mappingIpSecConnectionTunnelRouting = map[string]IpSecConnectionTunnelRoutingEnum{
	"BGP":    IpSecConnectionTunnelRoutingBgp,
	"STATIC": IpSecConnectionTunnelRoutingStatic,
}

// GetIpSecConnectionTunnelRoutingEnumValues Enumerates the set of values for IpSecConnectionTunnelRoutingEnum
func GetIpSecConnectionTunnelRoutingEnumValues() []IpSecConnectionTunnelRoutingEnum {
	values := make([]IpSecConnectionTunnelRoutingEnum, 0)
	for _, v := range mappingIpSecConnectionTunnelRouting {
		values = append(values, v)
	}
	return values
}
