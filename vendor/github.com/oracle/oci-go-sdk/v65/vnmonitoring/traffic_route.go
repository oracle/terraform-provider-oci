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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TrafficRoute Defines the traffic route taken in the path in `PathAnalysisResult`.
type TrafficRoute struct {

	// Reachability status for the given traffic route.
	ReachabilityStatus TrafficRouteReachabilityStatusEnum `mandatory:"true" json:"reachabilityStatus"`

	// The ordered sequence of nodes in the given the traffic route forming a path.
	Nodes []TrafficNode `mandatory:"true" json:"nodes"`

	// A description of the traffic route analysis. For example: "Traffic might not reach a destination
	// due to the LB backend being unhealthy".
	RouteAnalysisDescription *string `mandatory:"false" json:"routeAnalysisDescription"`
}

func (m TrafficRoute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TrafficRoute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTrafficRouteReachabilityStatusEnum(string(m.ReachabilityStatus)); !ok && m.ReachabilityStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReachabilityStatus: %s. Supported values are: %s.", m.ReachabilityStatus, strings.Join(GetTrafficRouteReachabilityStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *TrafficRoute) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		RouteAnalysisDescription *string                            `json:"routeAnalysisDescription"`
		ReachabilityStatus       TrafficRouteReachabilityStatusEnum `json:"reachabilityStatus"`
		Nodes                    []trafficnode                      `json:"nodes"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.RouteAnalysisDescription = model.RouteAnalysisDescription

	m.ReachabilityStatus = model.ReachabilityStatus

	m.Nodes = make([]TrafficNode, len(model.Nodes))
	for i, n := range model.Nodes {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Nodes[i] = nn.(TrafficNode)
		} else {
			m.Nodes[i] = nil
		}
	}
	return
}

// TrafficRouteReachabilityStatusEnum Enum with underlying type: string
type TrafficRouteReachabilityStatusEnum string

// Set of constants representing the allowable values for TrafficRouteReachabilityStatusEnum
const (
	TrafficRouteReachabilityStatusReachable     TrafficRouteReachabilityStatusEnum = "REACHABLE"
	TrafficRouteReachabilityStatusUnreachable   TrafficRouteReachabilityStatusEnum = "UNREACHABLE"
	TrafficRouteReachabilityStatusIndeterminate TrafficRouteReachabilityStatusEnum = "INDETERMINATE"
)

var mappingTrafficRouteReachabilityStatusEnum = map[string]TrafficRouteReachabilityStatusEnum{
	"REACHABLE":     TrafficRouteReachabilityStatusReachable,
	"UNREACHABLE":   TrafficRouteReachabilityStatusUnreachable,
	"INDETERMINATE": TrafficRouteReachabilityStatusIndeterminate,
}

var mappingTrafficRouteReachabilityStatusEnumLowerCase = map[string]TrafficRouteReachabilityStatusEnum{
	"reachable":     TrafficRouteReachabilityStatusReachable,
	"unreachable":   TrafficRouteReachabilityStatusUnreachable,
	"indeterminate": TrafficRouteReachabilityStatusIndeterminate,
}

// GetTrafficRouteReachabilityStatusEnumValues Enumerates the set of values for TrafficRouteReachabilityStatusEnum
func GetTrafficRouteReachabilityStatusEnumValues() []TrafficRouteReachabilityStatusEnum {
	values := make([]TrafficRouteReachabilityStatusEnum, 0)
	for _, v := range mappingTrafficRouteReachabilityStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetTrafficRouteReachabilityStatusEnumStringValues Enumerates the set of values in String for TrafficRouteReachabilityStatusEnum
func GetTrafficRouteReachabilityStatusEnumStringValues() []string {
	return []string{
		"REACHABLE",
		"UNREACHABLE",
		"INDETERMINATE",
	}
}

// GetMappingTrafficRouteReachabilityStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTrafficRouteReachabilityStatusEnum(val string) (TrafficRouteReachabilityStatusEnum, bool) {
	enum, ok := mappingTrafficRouteReachabilityStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
