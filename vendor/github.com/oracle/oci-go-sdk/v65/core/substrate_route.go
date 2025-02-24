// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SubstrateRoute Describes how to encapsulate overlay packets to their next substrate hop
// for a particular overlay route.
type SubstrateRoute struct {

	// Substrate IPv4 address for this route, in dot-decimal notation.
	Ip *string `mandatory:"true" json:"ip"`

	// Encapsulation format to use for this substrate route.
	Encap SubstrateRouteEncapEnum `mandatory:"true" json:"encap"`

	// Destination MPLS label to use for this substrate route.
	Label *int `mandatory:"true" json:"label"`
}

func (m SubstrateRoute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubstrateRoute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSubstrateRouteEncapEnum(string(m.Encap)); !ok && m.Encap != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Encap: %s. Supported values are: %s.", m.Encap, strings.Join(GetSubstrateRouteEncapEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SubstrateRouteEncapEnum Enum with underlying type: string
type SubstrateRouteEncapEnum string

// Set of constants representing the allowable values for SubstrateRouteEncapEnum
const (
	SubstrateRouteEncapMplsOverGre SubstrateRouteEncapEnum = "MPLS_OVER_GRE"
	SubstrateRouteEncapMplsOverUdp SubstrateRouteEncapEnum = "MPLS_OVER_UDP"
	SubstrateRouteEncapGeneve      SubstrateRouteEncapEnum = "GENEVE"
)

var mappingSubstrateRouteEncapEnum = map[string]SubstrateRouteEncapEnum{
	"MPLS_OVER_GRE": SubstrateRouteEncapMplsOverGre,
	"MPLS_OVER_UDP": SubstrateRouteEncapMplsOverUdp,
	"GENEVE":        SubstrateRouteEncapGeneve,
}

var mappingSubstrateRouteEncapEnumLowerCase = map[string]SubstrateRouteEncapEnum{
	"mpls_over_gre": SubstrateRouteEncapMplsOverGre,
	"mpls_over_udp": SubstrateRouteEncapMplsOverUdp,
	"geneve":        SubstrateRouteEncapGeneve,
}

// GetSubstrateRouteEncapEnumValues Enumerates the set of values for SubstrateRouteEncapEnum
func GetSubstrateRouteEncapEnumValues() []SubstrateRouteEncapEnum {
	values := make([]SubstrateRouteEncapEnum, 0)
	for _, v := range mappingSubstrateRouteEncapEnum {
		values = append(values, v)
	}
	return values
}

// GetSubstrateRouteEncapEnumStringValues Enumerates the set of values in String for SubstrateRouteEncapEnum
func GetSubstrateRouteEncapEnumStringValues() []string {
	return []string{
		"MPLS_OVER_GRE",
		"MPLS_OVER_UDP",
		"GENEVE",
	}
}

// GetMappingSubstrateRouteEncapEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubstrateRouteEncapEnum(val string) (SubstrateRouteEncapEnum, bool) {
	enum, ok := mappingSubstrateRouteEncapEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
