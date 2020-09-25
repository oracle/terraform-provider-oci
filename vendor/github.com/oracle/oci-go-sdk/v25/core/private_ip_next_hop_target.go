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

// PrivateIpNextHopTarget Details of a private IP nextHop target.
type PrivateIpNextHopTarget struct {

	// NextHop target's MPLS label.
	MplsLabel *int `mandatory:"false" json:"mplsLabel"`

	PortRange *PortRange `mandatory:"false" json:"portRange"`

	// NextHop target's substrate IP.
	SubstrateIp *string `mandatory:"false" json:"substrateIp"`

	// The OCID of the nextHop target entity.
	TargetId *string `mandatory:"false" json:"targetId"`

	// Type of nextHop target.
	TargetType PrivateIpNextHopTargetTargetTypeEnum `mandatory:"false" json:"targetType,omitempty"`
}

func (m PrivateIpNextHopTarget) String() string {
	return common.PointerString(m)
}

// PrivateIpNextHopTargetTargetTypeEnum Enum with underlying type: string
type PrivateIpNextHopTargetTargetTypeEnum string

// Set of constants representing the allowable values for PrivateIpNextHopTargetTargetTypeEnum
const (
	PrivateIpNextHopTargetTargetTypePadp       PrivateIpNextHopTargetTargetTypeEnum = "PADP"
	PrivateIpNextHopTargetTargetTypeVnicWorker PrivateIpNextHopTargetTargetTypeEnum = "VNIC_WORKER"
)

var mappingPrivateIpNextHopTargetTargetType = map[string]PrivateIpNextHopTargetTargetTypeEnum{
	"PADP":        PrivateIpNextHopTargetTargetTypePadp,
	"VNIC_WORKER": PrivateIpNextHopTargetTargetTypeVnicWorker,
}

// GetPrivateIpNextHopTargetTargetTypeEnumValues Enumerates the set of values for PrivateIpNextHopTargetTargetTypeEnum
func GetPrivateIpNextHopTargetTargetTypeEnumValues() []PrivateIpNextHopTargetTargetTypeEnum {
	values := make([]PrivateIpNextHopTargetTargetTypeEnum, 0)
	for _, v := range mappingPrivateIpNextHopTargetTargetType {
		values = append(values, v)
	}
	return values
}
