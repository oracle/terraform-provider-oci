// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrgUrtRouteRule Routing information needed by VCN DP to route DRG traffic.
type DrgUrtRouteRule struct {

	// The destination IP prefix (CIDR).
	IpPrefix *string `mandatory:"true" json:"ipPrefix"`

	// Encapsulation type.
	EncapType DrgUrtRouteRuleEncapTypeEnum `mandatory:"true" json:"encapType"`

	// The IP address of the next hop, for the destination prefix.
	NextHopIp *string `mandatory:"true" json:"nextHopIp"`

	// The label used in the encapsulated packet. For MPLSoUDP packets, this will be the
	// MPLS Label, and for VXLAN packets, this will be the VNI.
	EncapLabel *int `mandatory:"true" json:"encapLabel"`

	// Mac address of the next-hop for VXLAN encapped packets.
	NextHopMac *string `mandatory:"false" json:"nextHopMac"`

	// The max MTU size supported by the next hop.
	Mtu *int `mandatory:"false" json:"mtu"`

	// To be used for building/AD affinity.
	NextHopLocation *string `mandatory:"false" json:"nextHopLocation"`

	// The label of the next hop DRG attachment responsible for reaching the network destination.
	NextHopAttachmentLabel *int64 `mandatory:"false" json:"nextHopAttachmentLabel"`

	// Not supported now, reserved for future use.
	// The date and time the routes were last updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	LastModified *common.SDKTime `mandatory:"false" json:"lastModified"`
}

func (m DrgUrtRouteRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrgUrtRouteRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDrgUrtRouteRuleEncapTypeEnum(string(m.EncapType)); !ok && m.EncapType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EncapType: %s. Supported values are: %s.", m.EncapType, strings.Join(GetDrgUrtRouteRuleEncapTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DrgUrtRouteRuleEncapTypeEnum Enum with underlying type: string
type DrgUrtRouteRuleEncapTypeEnum string

// Set of constants representing the allowable values for DrgUrtRouteRuleEncapTypeEnum
const (
	DrgUrtRouteRuleEncapTypeVxlan    DrgUrtRouteRuleEncapTypeEnum = "VXLAN"
	DrgUrtRouteRuleEncapTypeMplsOUdp DrgUrtRouteRuleEncapTypeEnum = "MPLS_O_UDP"
)

var mappingDrgUrtRouteRuleEncapTypeEnum = map[string]DrgUrtRouteRuleEncapTypeEnum{
	"VXLAN":      DrgUrtRouteRuleEncapTypeVxlan,
	"MPLS_O_UDP": DrgUrtRouteRuleEncapTypeMplsOUdp,
}

var mappingDrgUrtRouteRuleEncapTypeEnumLowerCase = map[string]DrgUrtRouteRuleEncapTypeEnum{
	"vxlan":      DrgUrtRouteRuleEncapTypeVxlan,
	"mpls_o_udp": DrgUrtRouteRuleEncapTypeMplsOUdp,
}

// GetDrgUrtRouteRuleEncapTypeEnumValues Enumerates the set of values for DrgUrtRouteRuleEncapTypeEnum
func GetDrgUrtRouteRuleEncapTypeEnumValues() []DrgUrtRouteRuleEncapTypeEnum {
	values := make([]DrgUrtRouteRuleEncapTypeEnum, 0)
	for _, v := range mappingDrgUrtRouteRuleEncapTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrgUrtRouteRuleEncapTypeEnumStringValues Enumerates the set of values in String for DrgUrtRouteRuleEncapTypeEnum
func GetDrgUrtRouteRuleEncapTypeEnumStringValues() []string {
	return []string{
		"VXLAN",
		"MPLS_O_UDP",
	}
}

// GetMappingDrgUrtRouteRuleEncapTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrgUrtRouteRuleEncapTypeEnum(val string) (DrgUrtRouteRuleEncapTypeEnum, bool) {
	enum, ok := mappingDrgUrtRouteRuleEncapTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
