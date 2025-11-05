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

// FlexTunnelRouteSummary A summary of the routes advertised to and received from the SD-WAN.
type FlexTunnelRouteSummary struct {

	// The BGP network layer reachability information.
	Prefix *string `mandatory:"false" json:"prefix"`

	// The age of the route.
	Age *int64 `mandatory:"false" json:"age"`

	// Indicates this is the best route.
	IsBestPath *bool `mandatory:"false" json:"isBestPath"`

	// A list of ASNs in AS_Path.
	AsPath []int `mandatory:"false" json:"asPath"`

	// The source of the route advertisement.
	Advertiser FlexTunnelRouteSummaryAdvertiserEnum `mandatory:"false" json:"advertiser,omitempty"`
}

func (m FlexTunnelRouteSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FlexTunnelRouteSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFlexTunnelRouteSummaryAdvertiserEnum(string(m.Advertiser)); !ok && m.Advertiser != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Advertiser: %s. Supported values are: %s.", m.Advertiser, strings.Join(GetFlexTunnelRouteSummaryAdvertiserEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FlexTunnelRouteSummaryAdvertiserEnum Enum with underlying type: string
type FlexTunnelRouteSummaryAdvertiserEnum string

// Set of constants representing the allowable values for FlexTunnelRouteSummaryAdvertiserEnum
const (
	FlexTunnelRouteSummaryAdvertiserCustomer FlexTunnelRouteSummaryAdvertiserEnum = "CUSTOMER"
	FlexTunnelRouteSummaryAdvertiserOracle   FlexTunnelRouteSummaryAdvertiserEnum = "ORACLE"
)

var mappingFlexTunnelRouteSummaryAdvertiserEnum = map[string]FlexTunnelRouteSummaryAdvertiserEnum{
	"CUSTOMER": FlexTunnelRouteSummaryAdvertiserCustomer,
	"ORACLE":   FlexTunnelRouteSummaryAdvertiserOracle,
}

var mappingFlexTunnelRouteSummaryAdvertiserEnumLowerCase = map[string]FlexTunnelRouteSummaryAdvertiserEnum{
	"customer": FlexTunnelRouteSummaryAdvertiserCustomer,
	"oracle":   FlexTunnelRouteSummaryAdvertiserOracle,
}

// GetFlexTunnelRouteSummaryAdvertiserEnumValues Enumerates the set of values for FlexTunnelRouteSummaryAdvertiserEnum
func GetFlexTunnelRouteSummaryAdvertiserEnumValues() []FlexTunnelRouteSummaryAdvertiserEnum {
	values := make([]FlexTunnelRouteSummaryAdvertiserEnum, 0)
	for _, v := range mappingFlexTunnelRouteSummaryAdvertiserEnum {
		values = append(values, v)
	}
	return values
}

// GetFlexTunnelRouteSummaryAdvertiserEnumStringValues Enumerates the set of values in String for FlexTunnelRouteSummaryAdvertiserEnum
func GetFlexTunnelRouteSummaryAdvertiserEnumStringValues() []string {
	return []string{
		"CUSTOMER",
		"ORACLE",
	}
}

// GetMappingFlexTunnelRouteSummaryAdvertiserEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFlexTunnelRouteSummaryAdvertiserEnum(val string) (FlexTunnelRouteSummaryAdvertiserEnum, bool) {
	enum, ok := mappingFlexTunnelRouteSummaryAdvertiserEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
