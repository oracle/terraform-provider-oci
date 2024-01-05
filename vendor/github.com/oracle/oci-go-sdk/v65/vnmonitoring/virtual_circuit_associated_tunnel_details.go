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

// VirtualCircuitAssociatedTunnelDetails Detailed private tunnel info associated with the virtual circuit.
type VirtualCircuitAssociatedTunnelDetails struct {

	// The type of the tunnel associated with the virtual circuit.
	TunnelType VirtualCircuitAssociatedTunnelDetailsTunnelTypeEnum `mandatory:"true" json:"tunnelType"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IPSec tunnel associated with the virtual circuit.
	TunnelId *string `mandatory:"true" json:"tunnelId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of IPSec connection associated with the virtual circuit.
	IpsecConnectionId *string `mandatory:"false" json:"ipsecConnectionId"`
}

func (m VirtualCircuitAssociatedTunnelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VirtualCircuitAssociatedTunnelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVirtualCircuitAssociatedTunnelDetailsTunnelTypeEnum(string(m.TunnelType)); !ok && m.TunnelType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TunnelType: %s. Supported values are: %s.", m.TunnelType, strings.Join(GetVirtualCircuitAssociatedTunnelDetailsTunnelTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VirtualCircuitAssociatedTunnelDetailsTunnelTypeEnum Enum with underlying type: string
type VirtualCircuitAssociatedTunnelDetailsTunnelTypeEnum string

// Set of constants representing the allowable values for VirtualCircuitAssociatedTunnelDetailsTunnelTypeEnum
const (
	VirtualCircuitAssociatedTunnelDetailsTunnelTypeIpsec VirtualCircuitAssociatedTunnelDetailsTunnelTypeEnum = "IPSEC"
)

var mappingVirtualCircuitAssociatedTunnelDetailsTunnelTypeEnum = map[string]VirtualCircuitAssociatedTunnelDetailsTunnelTypeEnum{
	"IPSEC": VirtualCircuitAssociatedTunnelDetailsTunnelTypeIpsec,
}

var mappingVirtualCircuitAssociatedTunnelDetailsTunnelTypeEnumLowerCase = map[string]VirtualCircuitAssociatedTunnelDetailsTunnelTypeEnum{
	"ipsec": VirtualCircuitAssociatedTunnelDetailsTunnelTypeIpsec,
}

// GetVirtualCircuitAssociatedTunnelDetailsTunnelTypeEnumValues Enumerates the set of values for VirtualCircuitAssociatedTunnelDetailsTunnelTypeEnum
func GetVirtualCircuitAssociatedTunnelDetailsTunnelTypeEnumValues() []VirtualCircuitAssociatedTunnelDetailsTunnelTypeEnum {
	values := make([]VirtualCircuitAssociatedTunnelDetailsTunnelTypeEnum, 0)
	for _, v := range mappingVirtualCircuitAssociatedTunnelDetailsTunnelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVirtualCircuitAssociatedTunnelDetailsTunnelTypeEnumStringValues Enumerates the set of values in String for VirtualCircuitAssociatedTunnelDetailsTunnelTypeEnum
func GetVirtualCircuitAssociatedTunnelDetailsTunnelTypeEnumStringValues() []string {
	return []string{
		"IPSEC",
	}
}

// GetMappingVirtualCircuitAssociatedTunnelDetailsTunnelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVirtualCircuitAssociatedTunnelDetailsTunnelTypeEnum(val string) (VirtualCircuitAssociatedTunnelDetailsTunnelTypeEnum, bool) {
	enum, ok := mappingVirtualCircuitAssociatedTunnelDetailsTunnelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
