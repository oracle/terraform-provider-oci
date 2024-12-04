// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// RouteReflectorRoute Disintermediated routes
type RouteReflectorRoute struct {

	// CIDR.
	Cidr *int64 `mandatory:"true" json:"cidr"`

	// The RIB entries.
	RibEntries []string `mandatory:"true" json:"ribEntries"`

	// The route Type.
	RouteType *int64 `mandatory:"true" json:"routeType"`

	// The infobase
	Infobase RouteReflectorRouteInfobaseEnum `mandatory:"false" json:"infobase,omitempty"`

	// The vrf label
	Vrf *int `mandatory:"false" json:"vrf"`
}

func (m RouteReflectorRoute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RouteReflectorRoute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRouteReflectorRouteInfobaseEnum(string(m.Infobase)); !ok && m.Infobase != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Infobase: %s. Supported values are: %s.", m.Infobase, strings.Join(GetRouteReflectorRouteInfobaseEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RouteReflectorRouteInfobaseEnum Enum with underlying type: string
type RouteReflectorRouteInfobaseEnum string

// Set of constants representing the allowable values for RouteReflectorRouteInfobaseEnum
const (
	RouteReflectorRouteInfobaseFib RouteReflectorRouteInfobaseEnum = "FIB"
	RouteReflectorRouteInfobaseRib RouteReflectorRouteInfobaseEnum = "RIB"
)

var mappingRouteReflectorRouteInfobaseEnum = map[string]RouteReflectorRouteInfobaseEnum{
	"FIB": RouteReflectorRouteInfobaseFib,
	"RIB": RouteReflectorRouteInfobaseRib,
}

var mappingRouteReflectorRouteInfobaseEnumLowerCase = map[string]RouteReflectorRouteInfobaseEnum{
	"fib": RouteReflectorRouteInfobaseFib,
	"rib": RouteReflectorRouteInfobaseRib,
}

// GetRouteReflectorRouteInfobaseEnumValues Enumerates the set of values for RouteReflectorRouteInfobaseEnum
func GetRouteReflectorRouteInfobaseEnumValues() []RouteReflectorRouteInfobaseEnum {
	values := make([]RouteReflectorRouteInfobaseEnum, 0)
	for _, v := range mappingRouteReflectorRouteInfobaseEnum {
		values = append(values, v)
	}
	return values
}

// GetRouteReflectorRouteInfobaseEnumStringValues Enumerates the set of values in String for RouteReflectorRouteInfobaseEnum
func GetRouteReflectorRouteInfobaseEnumStringValues() []string {
	return []string{
		"FIB",
		"RIB",
	}
}

// GetMappingRouteReflectorRouteInfobaseEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRouteReflectorRouteInfobaseEnum(val string) (RouteReflectorRouteInfobaseEnum, bool) {
	enum, ok := mappingRouteReflectorRouteInfobaseEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
