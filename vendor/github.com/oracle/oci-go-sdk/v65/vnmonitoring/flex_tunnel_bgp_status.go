// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FlexTunnelBgpStatus Boarder Gateway Protocol (BGP) session status
type FlexTunnelBgpStatus struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the flex tunnel.
	FlexTunnelId *string `mandatory:"true" json:"flexTunnelId"`

	// The state of the IPv4 BGP session associated with the flex tunnel.
	BgpSessionStatus FlexTunnelBgpStatusBgpSessionStatusEnum `mandatory:"false" json:"bgpSessionStatus,omitempty"`

	// The state of the IPv6 BGP session associated with the flex tunnel.
	BgpSessionStatusIpv6 FlexTunnelBgpStatusBgpSessionStatusIpv6Enum `mandatory:"false" json:"bgpSessionStatusIpv6,omitempty"`
}

func (m FlexTunnelBgpStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FlexTunnelBgpStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFlexTunnelBgpStatusBgpSessionStatusEnum(string(m.BgpSessionStatus)); !ok && m.BgpSessionStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BgpSessionStatus: %s. Supported values are: %s.", m.BgpSessionStatus, strings.Join(GetFlexTunnelBgpStatusBgpSessionStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFlexTunnelBgpStatusBgpSessionStatusIpv6Enum(string(m.BgpSessionStatusIpv6)); !ok && m.BgpSessionStatusIpv6 != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BgpSessionStatusIpv6: %s. Supported values are: %s.", m.BgpSessionStatusIpv6, strings.Join(GetFlexTunnelBgpStatusBgpSessionStatusIpv6EnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FlexTunnelBgpStatusBgpSessionStatusEnum Enum with underlying type: string
type FlexTunnelBgpStatusBgpSessionStatusEnum string

// Set of constants representing the allowable values for FlexTunnelBgpStatusBgpSessionStatusEnum
const (
	FlexTunnelBgpStatusBgpSessionStatusUp   FlexTunnelBgpStatusBgpSessionStatusEnum = "UP"
	FlexTunnelBgpStatusBgpSessionStatusDown FlexTunnelBgpStatusBgpSessionStatusEnum = "DOWN"
)

var mappingFlexTunnelBgpStatusBgpSessionStatusEnum = map[string]FlexTunnelBgpStatusBgpSessionStatusEnum{
	"UP":   FlexTunnelBgpStatusBgpSessionStatusUp,
	"DOWN": FlexTunnelBgpStatusBgpSessionStatusDown,
}

var mappingFlexTunnelBgpStatusBgpSessionStatusEnumLowerCase = map[string]FlexTunnelBgpStatusBgpSessionStatusEnum{
	"up":   FlexTunnelBgpStatusBgpSessionStatusUp,
	"down": FlexTunnelBgpStatusBgpSessionStatusDown,
}

// GetFlexTunnelBgpStatusBgpSessionStatusEnumValues Enumerates the set of values for FlexTunnelBgpStatusBgpSessionStatusEnum
func GetFlexTunnelBgpStatusBgpSessionStatusEnumValues() []FlexTunnelBgpStatusBgpSessionStatusEnum {
	values := make([]FlexTunnelBgpStatusBgpSessionStatusEnum, 0)
	for _, v := range mappingFlexTunnelBgpStatusBgpSessionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetFlexTunnelBgpStatusBgpSessionStatusEnumStringValues Enumerates the set of values in String for FlexTunnelBgpStatusBgpSessionStatusEnum
func GetFlexTunnelBgpStatusBgpSessionStatusEnumStringValues() []string {
	return []string{
		"UP",
		"DOWN",
	}
}

// GetMappingFlexTunnelBgpStatusBgpSessionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFlexTunnelBgpStatusBgpSessionStatusEnum(val string) (FlexTunnelBgpStatusBgpSessionStatusEnum, bool) {
	enum, ok := mappingFlexTunnelBgpStatusBgpSessionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// FlexTunnelBgpStatusBgpSessionStatusIpv6Enum Enum with underlying type: string
type FlexTunnelBgpStatusBgpSessionStatusIpv6Enum string

// Set of constants representing the allowable values for FlexTunnelBgpStatusBgpSessionStatusIpv6Enum
const (
	FlexTunnelBgpStatusBgpSessionStatusIpv6Up   FlexTunnelBgpStatusBgpSessionStatusIpv6Enum = "UP"
	FlexTunnelBgpStatusBgpSessionStatusIpv6Down FlexTunnelBgpStatusBgpSessionStatusIpv6Enum = "DOWN"
)

var mappingFlexTunnelBgpStatusBgpSessionStatusIpv6Enum = map[string]FlexTunnelBgpStatusBgpSessionStatusIpv6Enum{
	"UP":   FlexTunnelBgpStatusBgpSessionStatusIpv6Up,
	"DOWN": FlexTunnelBgpStatusBgpSessionStatusIpv6Down,
}

var mappingFlexTunnelBgpStatusBgpSessionStatusIpv6EnumLowerCase = map[string]FlexTunnelBgpStatusBgpSessionStatusIpv6Enum{
	"up":   FlexTunnelBgpStatusBgpSessionStatusIpv6Up,
	"down": FlexTunnelBgpStatusBgpSessionStatusIpv6Down,
}

// GetFlexTunnelBgpStatusBgpSessionStatusIpv6EnumValues Enumerates the set of values for FlexTunnelBgpStatusBgpSessionStatusIpv6Enum
func GetFlexTunnelBgpStatusBgpSessionStatusIpv6EnumValues() []FlexTunnelBgpStatusBgpSessionStatusIpv6Enum {
	values := make([]FlexTunnelBgpStatusBgpSessionStatusIpv6Enum, 0)
	for _, v := range mappingFlexTunnelBgpStatusBgpSessionStatusIpv6Enum {
		values = append(values, v)
	}
	return values
}

// GetFlexTunnelBgpStatusBgpSessionStatusIpv6EnumStringValues Enumerates the set of values in String for FlexTunnelBgpStatusBgpSessionStatusIpv6Enum
func GetFlexTunnelBgpStatusBgpSessionStatusIpv6EnumStringValues() []string {
	return []string{
		"UP",
		"DOWN",
	}
}

// GetMappingFlexTunnelBgpStatusBgpSessionStatusIpv6Enum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFlexTunnelBgpStatusBgpSessionStatusIpv6Enum(val string) (FlexTunnelBgpStatusBgpSessionStatusIpv6Enum, bool) {
	enum, ok := mappingFlexTunnelBgpStatusBgpSessionStatusIpv6EnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
