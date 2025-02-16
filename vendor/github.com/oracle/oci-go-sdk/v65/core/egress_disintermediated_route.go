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

// EgressDisintermediatedRoute Egress Disintermediated routes
type EgressDisintermediatedRoute struct {

	// A singular IP address range in CIDR notation used for matching destination when routing traffic.
	IpPrefix *string `mandatory:"true" json:"ipPrefix"`

	// The IP address of the next hop, for the destination prefix.
	NextHopIp *string `mandatory:"true" json:"nextHopIp"`

	// The label used in the encapsulated packet. For MPLS_O_UDP packets, this will be the MPLS Label. For VXLAN
	// packets, this will be the VNI.
	EncapLabel *int64 `mandatory:"true" json:"encapLabel"`

	// Route table label
	Label *int `mandatory:"false" json:"label"`

	// ad Name
	AdName EgressDisintermediatedRouteAdNameEnum `mandatory:"false" json:"adName,omitempty"`

	// shard id
	ShardId *int `mandatory:"false" json:"shardId"`
}

func (m EgressDisintermediatedRoute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EgressDisintermediatedRoute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingEgressDisintermediatedRouteAdNameEnum(string(m.AdName)); !ok && m.AdName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdName: %s. Supported values are: %s.", m.AdName, strings.Join(GetEgressDisintermediatedRouteAdNameEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EgressDisintermediatedRouteAdNameEnum Enum with underlying type: string
type EgressDisintermediatedRouteAdNameEnum string

// Set of constants representing the allowable values for EgressDisintermediatedRouteAdNameEnum
const (
	EgressDisintermediatedRouteAdNameAd1  EgressDisintermediatedRouteAdNameEnum = "AD1"
	EgressDisintermediatedRouteAdNameAd2  EgressDisintermediatedRouteAdNameEnum = "AD2"
	EgressDisintermediatedRouteAdNameAd3  EgressDisintermediatedRouteAdNameEnum = "AD3"
	EgressDisintermediatedRouteAdNamePop1 EgressDisintermediatedRouteAdNameEnum = "POP1"
	EgressDisintermediatedRouteAdNamePop2 EgressDisintermediatedRouteAdNameEnum = "POP2"
)

var mappingEgressDisintermediatedRouteAdNameEnum = map[string]EgressDisintermediatedRouteAdNameEnum{
	"AD1":  EgressDisintermediatedRouteAdNameAd1,
	"AD2":  EgressDisintermediatedRouteAdNameAd2,
	"AD3":  EgressDisintermediatedRouteAdNameAd3,
	"POP1": EgressDisintermediatedRouteAdNamePop1,
	"POP2": EgressDisintermediatedRouteAdNamePop2,
}

var mappingEgressDisintermediatedRouteAdNameEnumLowerCase = map[string]EgressDisintermediatedRouteAdNameEnum{
	"ad1":  EgressDisintermediatedRouteAdNameAd1,
	"ad2":  EgressDisintermediatedRouteAdNameAd2,
	"ad3":  EgressDisintermediatedRouteAdNameAd3,
	"pop1": EgressDisintermediatedRouteAdNamePop1,
	"pop2": EgressDisintermediatedRouteAdNamePop2,
}

// GetEgressDisintermediatedRouteAdNameEnumValues Enumerates the set of values for EgressDisintermediatedRouteAdNameEnum
func GetEgressDisintermediatedRouteAdNameEnumValues() []EgressDisintermediatedRouteAdNameEnum {
	values := make([]EgressDisintermediatedRouteAdNameEnum, 0)
	for _, v := range mappingEgressDisintermediatedRouteAdNameEnum {
		values = append(values, v)
	}
	return values
}

// GetEgressDisintermediatedRouteAdNameEnumStringValues Enumerates the set of values in String for EgressDisintermediatedRouteAdNameEnum
func GetEgressDisintermediatedRouteAdNameEnumStringValues() []string {
	return []string{
		"AD1",
		"AD2",
		"AD3",
		"POP1",
		"POP2",
	}
}

// GetMappingEgressDisintermediatedRouteAdNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEgressDisintermediatedRouteAdNameEnum(val string) (EgressDisintermediatedRouteAdNameEnum, bool) {
	enum, ok := mappingEgressDisintermediatedRouteAdNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
