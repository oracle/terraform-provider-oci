// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// PrivateIpNextHopProtocolPort Details containing the port number and the protocol type
type PrivateIpNextHopProtocolPort struct {

	// VNICaaS uses this to identify the port number to flow-hash traffic
	Port *int `mandatory:"false" json:"port"`

	// The type of protocol i.e. TCP, UDP or ALL accompanied by port number used for flow-hash by VNICaaS
	Protocol PrivateIpNextHopProtocolPortProtocolEnum `mandatory:"false" json:"protocol,omitempty"`
}

func (m PrivateIpNextHopProtocolPort) String() string {
	return common.PointerString(m)
}

// PrivateIpNextHopProtocolPortProtocolEnum Enum with underlying type: string
type PrivateIpNextHopProtocolPortProtocolEnum string

// Set of constants representing the allowable values for PrivateIpNextHopProtocolPortProtocolEnum
const (
	PrivateIpNextHopProtocolPortProtocolTcp PrivateIpNextHopProtocolPortProtocolEnum = "TCP"
	PrivateIpNextHopProtocolPortProtocolUdp PrivateIpNextHopProtocolPortProtocolEnum = "UDP"
	PrivateIpNextHopProtocolPortProtocolAll PrivateIpNextHopProtocolPortProtocolEnum = "ALL"
)

var mappingPrivateIpNextHopProtocolPortProtocol = map[string]PrivateIpNextHopProtocolPortProtocolEnum{
	"TCP": PrivateIpNextHopProtocolPortProtocolTcp,
	"UDP": PrivateIpNextHopProtocolPortProtocolUdp,
	"ALL": PrivateIpNextHopProtocolPortProtocolAll,
}

// GetPrivateIpNextHopProtocolPortProtocolEnumValues Enumerates the set of values for PrivateIpNextHopProtocolPortProtocolEnum
func GetPrivateIpNextHopProtocolPortProtocolEnumValues() []PrivateIpNextHopProtocolPortProtocolEnum {
	values := make([]PrivateIpNextHopProtocolPortProtocolEnum, 0)
	for _, v := range mappingPrivateIpNextHopProtocolPortProtocol {
		values = append(values, v)
	}
	return values
}
