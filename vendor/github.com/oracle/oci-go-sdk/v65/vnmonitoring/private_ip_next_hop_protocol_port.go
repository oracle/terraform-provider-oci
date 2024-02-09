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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrivateIpNextHopProtocolPort) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPrivateIpNextHopProtocolPortProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetPrivateIpNextHopProtocolPortProtocolEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PrivateIpNextHopProtocolPortProtocolEnum Enum with underlying type: string
type PrivateIpNextHopProtocolPortProtocolEnum string

// Set of constants representing the allowable values for PrivateIpNextHopProtocolPortProtocolEnum
const (
	PrivateIpNextHopProtocolPortProtocolTcp PrivateIpNextHopProtocolPortProtocolEnum = "TCP"
	PrivateIpNextHopProtocolPortProtocolUdp PrivateIpNextHopProtocolPortProtocolEnum = "UDP"
	PrivateIpNextHopProtocolPortProtocolAll PrivateIpNextHopProtocolPortProtocolEnum = "ALL"
)

var mappingPrivateIpNextHopProtocolPortProtocolEnum = map[string]PrivateIpNextHopProtocolPortProtocolEnum{
	"TCP": PrivateIpNextHopProtocolPortProtocolTcp,
	"UDP": PrivateIpNextHopProtocolPortProtocolUdp,
	"ALL": PrivateIpNextHopProtocolPortProtocolAll,
}

var mappingPrivateIpNextHopProtocolPortProtocolEnumLowerCase = map[string]PrivateIpNextHopProtocolPortProtocolEnum{
	"tcp": PrivateIpNextHopProtocolPortProtocolTcp,
	"udp": PrivateIpNextHopProtocolPortProtocolUdp,
	"all": PrivateIpNextHopProtocolPortProtocolAll,
}

// GetPrivateIpNextHopProtocolPortProtocolEnumValues Enumerates the set of values for PrivateIpNextHopProtocolPortProtocolEnum
func GetPrivateIpNextHopProtocolPortProtocolEnumValues() []PrivateIpNextHopProtocolPortProtocolEnum {
	values := make([]PrivateIpNextHopProtocolPortProtocolEnum, 0)
	for _, v := range mappingPrivateIpNextHopProtocolPortProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetPrivateIpNextHopProtocolPortProtocolEnumStringValues Enumerates the set of values in String for PrivateIpNextHopProtocolPortProtocolEnum
func GetPrivateIpNextHopProtocolPortProtocolEnumStringValues() []string {
	return []string{
		"TCP",
		"UDP",
		"ALL",
	}
}

// GetMappingPrivateIpNextHopProtocolPortProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPrivateIpNextHopProtocolPortProtocolEnum(val string) (PrivateIpNextHopProtocolPortProtocolEnum, bool) {
	enum, ok := mappingPrivateIpNextHopProtocolPortProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
