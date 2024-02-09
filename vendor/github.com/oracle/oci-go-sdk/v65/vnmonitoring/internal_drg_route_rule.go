// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InternalDrgRouteRule An internal DRG route rule is a mapping between a destination IP address range and a DRG attachment. It is used
// to offload DRG functionality (primarily routing, but up-to-and-including all additional features associated with
// DRG attachments) onto the VCN Dataplane.
type InternalDrgRouteRule struct {

	// A singular IP address range in CIDR notation used for matching destination when routing traffic.
	IpPrefix *string `mandatory:"true" json:"ipPrefix"`

	// The IP address of the next hop, for the destination prefix.
	NextHopIp *string `mandatory:"true" json:"nextHopIp"`

	// An encapsulation type for the traffic.
	EncapType InternalDrgRouteRuleEncapTypeEnum `mandatory:"true" json:"encapType"`

	// The label used in the encapsulated packet. For MPLS_O_UDP packets, this will be the MPLS Label. For VXLAN
	// packets, this will be the VNI.
	EncapLabel *int `mandatory:"true" json:"encapLabel"`

	// The MAC address of the next hop. It can be integer value or string format separated by colon.
	NextHopMac *string `mandatory:"false" json:"nextHopMac"`
}

func (m InternalDrgRouteRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalDrgRouteRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInternalDrgRouteRuleEncapTypeEnum(string(m.EncapType)); !ok && m.EncapType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EncapType: %s. Supported values are: %s.", m.EncapType, strings.Join(GetInternalDrgRouteRuleEncapTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InternalDrgRouteRuleEncapTypeEnum Enum with underlying type: string
type InternalDrgRouteRuleEncapTypeEnum string

// Set of constants representing the allowable values for InternalDrgRouteRuleEncapTypeEnum
const (
	InternalDrgRouteRuleEncapTypeVxlan    InternalDrgRouteRuleEncapTypeEnum = "VXLAN"
	InternalDrgRouteRuleEncapTypeMplsOUdp InternalDrgRouteRuleEncapTypeEnum = "MPLS_O_UDP"
)

var mappingInternalDrgRouteRuleEncapTypeEnum = map[string]InternalDrgRouteRuleEncapTypeEnum{
	"VXLAN":      InternalDrgRouteRuleEncapTypeVxlan,
	"MPLS_O_UDP": InternalDrgRouteRuleEncapTypeMplsOUdp,
}

var mappingInternalDrgRouteRuleEncapTypeEnumLowerCase = map[string]InternalDrgRouteRuleEncapTypeEnum{
	"vxlan":      InternalDrgRouteRuleEncapTypeVxlan,
	"mpls_o_udp": InternalDrgRouteRuleEncapTypeMplsOUdp,
}

// GetInternalDrgRouteRuleEncapTypeEnumValues Enumerates the set of values for InternalDrgRouteRuleEncapTypeEnum
func GetInternalDrgRouteRuleEncapTypeEnumValues() []InternalDrgRouteRuleEncapTypeEnum {
	values := make([]InternalDrgRouteRuleEncapTypeEnum, 0)
	for _, v := range mappingInternalDrgRouteRuleEncapTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalDrgRouteRuleEncapTypeEnumStringValues Enumerates the set of values in String for InternalDrgRouteRuleEncapTypeEnum
func GetInternalDrgRouteRuleEncapTypeEnumStringValues() []string {
	return []string{
		"VXLAN",
		"MPLS_O_UDP",
	}
}

// GetMappingInternalDrgRouteRuleEncapTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalDrgRouteRuleEncapTypeEnum(val string) (InternalDrgRouteRuleEncapTypeEnum, bool) {
	enum, ok := mappingInternalDrgRouteRuleEncapTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
