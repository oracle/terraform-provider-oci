// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Guarded Data Pipelines API
//
// Use Guarded Data Pipelines to facilitate data transfer between different security domains. The service provides physical, network, and logistical isolation between security domains, malware and vulnerability scanning, auditing, and logging, with deep content inspection capabilities.
//

package gdp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GdpMpPeeringRegion Region and state information for a GDP MP pipeline peering region.
type GdpMpPeeringRegion struct {

	// Public region name where a peered pipeline exists.
	Region *string `mandatory:"true" json:"region"`

	// The current state of the peering region.
	RegionState GdpMpPeeringRegionRegionStateEnum `mandatory:"true" json:"regionState"`
}

func (m GdpMpPeeringRegion) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GdpMpPeeringRegion) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGdpMpPeeringRegionRegionStateEnum(string(m.RegionState)); !ok && m.RegionState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RegionState: %s. Supported values are: %s.", m.RegionState, strings.Join(GetGdpMpPeeringRegionRegionStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GdpMpPeeringRegionRegionStateEnum Enum with underlying type: string
type GdpMpPeeringRegionRegionStateEnum string

// Set of constants representing the allowable values for GdpMpPeeringRegionRegionStateEnum
const (
	GdpMpPeeringRegionRegionStateActive   GdpMpPeeringRegionRegionStateEnum = "ACTIVE"
	GdpMpPeeringRegionRegionStateInactive GdpMpPeeringRegionRegionStateEnum = "INACTIVE"
)

var mappingGdpMpPeeringRegionRegionStateEnum = map[string]GdpMpPeeringRegionRegionStateEnum{
	"ACTIVE":   GdpMpPeeringRegionRegionStateActive,
	"INACTIVE": GdpMpPeeringRegionRegionStateInactive,
}

var mappingGdpMpPeeringRegionRegionStateEnumLowerCase = map[string]GdpMpPeeringRegionRegionStateEnum{
	"active":   GdpMpPeeringRegionRegionStateActive,
	"inactive": GdpMpPeeringRegionRegionStateInactive,
}

// GetGdpMpPeeringRegionRegionStateEnumValues Enumerates the set of values for GdpMpPeeringRegionRegionStateEnum
func GetGdpMpPeeringRegionRegionStateEnumValues() []GdpMpPeeringRegionRegionStateEnum {
	values := make([]GdpMpPeeringRegionRegionStateEnum, 0)
	for _, v := range mappingGdpMpPeeringRegionRegionStateEnum {
		values = append(values, v)
	}
	return values
}

// GetGdpMpPeeringRegionRegionStateEnumStringValues Enumerates the set of values in String for GdpMpPeeringRegionRegionStateEnum
func GetGdpMpPeeringRegionRegionStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingGdpMpPeeringRegionRegionStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGdpMpPeeringRegionRegionStateEnum(val string) (GdpMpPeeringRegionRegionStateEnum, bool) {
	enum, ok := mappingGdpMpPeeringRegionRegionStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
