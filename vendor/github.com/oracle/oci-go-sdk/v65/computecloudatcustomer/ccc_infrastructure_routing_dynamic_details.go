// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.cloud.oracle.com/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CccInfrastructureRoutingDynamicDetails Dynamic routing information for the Compute Cloud@Customer infrastructure.
type CccInfrastructureRoutingDynamicDetails struct {

	// The list of peer devices in the dynamic routing configuration.
	PeerInformation []PeerInformation `mandatory:"false" json:"peerInformation"`

	// The Oracle Autonomous System Number (ASN) to control routing and exchange information
	// within the dynamic routing configuration.
	OracleAsn *int `mandatory:"false" json:"oracleAsn"`

	// The topology in use for the Border Gateway Protocol (BGP) configuration.
	BgpTopology CccInfrastructureRoutingDynamicDetailsBgpTopologyEnum `mandatory:"false" json:"bgpTopology,omitempty"`
}

func (m CccInfrastructureRoutingDynamicDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CccInfrastructureRoutingDynamicDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCccInfrastructureRoutingDynamicDetailsBgpTopologyEnum(string(m.BgpTopology)); !ok && m.BgpTopology != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BgpTopology: %s. Supported values are: %s.", m.BgpTopology, strings.Join(GetCccInfrastructureRoutingDynamicDetailsBgpTopologyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CccInfrastructureRoutingDynamicDetailsBgpTopologyEnum Enum with underlying type: string
type CccInfrastructureRoutingDynamicDetailsBgpTopologyEnum string

// Set of constants representing the allowable values for CccInfrastructureRoutingDynamicDetailsBgpTopologyEnum
const (
	CccInfrastructureRoutingDynamicDetailsBgpTopologyTriangle CccInfrastructureRoutingDynamicDetailsBgpTopologyEnum = "TRIANGLE"
	CccInfrastructureRoutingDynamicDetailsBgpTopologySquare   CccInfrastructureRoutingDynamicDetailsBgpTopologyEnum = "SQUARE"
	CccInfrastructureRoutingDynamicDetailsBgpTopologyMesh     CccInfrastructureRoutingDynamicDetailsBgpTopologyEnum = "MESH"
)

var mappingCccInfrastructureRoutingDynamicDetailsBgpTopologyEnum = map[string]CccInfrastructureRoutingDynamicDetailsBgpTopologyEnum{
	"TRIANGLE": CccInfrastructureRoutingDynamicDetailsBgpTopologyTriangle,
	"SQUARE":   CccInfrastructureRoutingDynamicDetailsBgpTopologySquare,
	"MESH":     CccInfrastructureRoutingDynamicDetailsBgpTopologyMesh,
}

var mappingCccInfrastructureRoutingDynamicDetailsBgpTopologyEnumLowerCase = map[string]CccInfrastructureRoutingDynamicDetailsBgpTopologyEnum{
	"triangle": CccInfrastructureRoutingDynamicDetailsBgpTopologyTriangle,
	"square":   CccInfrastructureRoutingDynamicDetailsBgpTopologySquare,
	"mesh":     CccInfrastructureRoutingDynamicDetailsBgpTopologyMesh,
}

// GetCccInfrastructureRoutingDynamicDetailsBgpTopologyEnumValues Enumerates the set of values for CccInfrastructureRoutingDynamicDetailsBgpTopologyEnum
func GetCccInfrastructureRoutingDynamicDetailsBgpTopologyEnumValues() []CccInfrastructureRoutingDynamicDetailsBgpTopologyEnum {
	values := make([]CccInfrastructureRoutingDynamicDetailsBgpTopologyEnum, 0)
	for _, v := range mappingCccInfrastructureRoutingDynamicDetailsBgpTopologyEnum {
		values = append(values, v)
	}
	return values
}

// GetCccInfrastructureRoutingDynamicDetailsBgpTopologyEnumStringValues Enumerates the set of values in String for CccInfrastructureRoutingDynamicDetailsBgpTopologyEnum
func GetCccInfrastructureRoutingDynamicDetailsBgpTopologyEnumStringValues() []string {
	return []string{
		"TRIANGLE",
		"SQUARE",
		"MESH",
	}
}

// GetMappingCccInfrastructureRoutingDynamicDetailsBgpTopologyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCccInfrastructureRoutingDynamicDetailsBgpTopologyEnum(val string) (CccInfrastructureRoutingDynamicDetailsBgpTopologyEnum, bool) {
	enum, ok := mappingCccInfrastructureRoutingDynamicDetailsBgpTopologyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
