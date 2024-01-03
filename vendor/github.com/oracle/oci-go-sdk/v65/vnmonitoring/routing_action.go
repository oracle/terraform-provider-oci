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

// RoutingAction Defines the details for routing actions taken on the traffic flow.
type RoutingAction interface {

	// The type of the routing support for the traffic flow.
	GetActionType() RoutingActionActionTypeEnum
}

type routingaction struct {
	JsonData   []byte
	ActionType RoutingActionActionTypeEnum `mandatory:"true" json:"actionType"`
	Action     string                      `json:"action"`
}

// UnmarshalJSON unmarshals json
func (m *routingaction) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerroutingaction routingaction
	s := struct {
		Model Unmarshalerroutingaction
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ActionType = s.Model.ActionType
	m.Action = s.Model.Action

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *routingaction) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Action {
	case "NO_ROUTE":
		mm := NoRouteRoutingAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INDETERMINATE":
		mm := IndeterminateRoutingAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FORWARDED":
		mm := ForwardedRoutingAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for RoutingAction: %s.", m.Action)
		return *m, nil
	}
}

// GetActionType returns ActionType
func (m routingaction) GetActionType() RoutingActionActionTypeEnum {
	return m.ActionType
}

func (m routingaction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m routingaction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRoutingActionActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetRoutingActionActionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RoutingActionActionTypeEnum Enum with underlying type: string
type RoutingActionActionTypeEnum string

// Set of constants representing the allowable values for RoutingActionActionTypeEnum
const (
	RoutingActionActionTypeExplicit     RoutingActionActionTypeEnum = "EXPLICIT"
	RoutingActionActionTypeImplicit     RoutingActionActionTypeEnum = "IMPLICIT"
	RoutingActionActionTypeNotSupported RoutingActionActionTypeEnum = "NOT_SUPPORTED"
)

var mappingRoutingActionActionTypeEnum = map[string]RoutingActionActionTypeEnum{
	"EXPLICIT":      RoutingActionActionTypeExplicit,
	"IMPLICIT":      RoutingActionActionTypeImplicit,
	"NOT_SUPPORTED": RoutingActionActionTypeNotSupported,
}

var mappingRoutingActionActionTypeEnumLowerCase = map[string]RoutingActionActionTypeEnum{
	"explicit":      RoutingActionActionTypeExplicit,
	"implicit":      RoutingActionActionTypeImplicit,
	"not_supported": RoutingActionActionTypeNotSupported,
}

// GetRoutingActionActionTypeEnumValues Enumerates the set of values for RoutingActionActionTypeEnum
func GetRoutingActionActionTypeEnumValues() []RoutingActionActionTypeEnum {
	values := make([]RoutingActionActionTypeEnum, 0)
	for _, v := range mappingRoutingActionActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRoutingActionActionTypeEnumStringValues Enumerates the set of values in String for RoutingActionActionTypeEnum
func GetRoutingActionActionTypeEnumStringValues() []string {
	return []string{
		"EXPLICIT",
		"IMPLICIT",
		"NOT_SUPPORTED",
	}
}

// GetMappingRoutingActionActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRoutingActionActionTypeEnum(val string) (RoutingActionActionTypeEnum, bool) {
	enum, ok := mappingRoutingActionActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RoutingActionActionEnum Enum with underlying type: string
type RoutingActionActionEnum string

// Set of constants representing the allowable values for RoutingActionActionEnum
const (
	RoutingActionActionForwarded     RoutingActionActionEnum = "FORWARDED"
	RoutingActionActionNoRoute       RoutingActionActionEnum = "NO_ROUTE"
	RoutingActionActionIndeterminate RoutingActionActionEnum = "INDETERMINATE"
)

var mappingRoutingActionActionEnum = map[string]RoutingActionActionEnum{
	"FORWARDED":     RoutingActionActionForwarded,
	"NO_ROUTE":      RoutingActionActionNoRoute,
	"INDETERMINATE": RoutingActionActionIndeterminate,
}

var mappingRoutingActionActionEnumLowerCase = map[string]RoutingActionActionEnum{
	"forwarded":     RoutingActionActionForwarded,
	"no_route":      RoutingActionActionNoRoute,
	"indeterminate": RoutingActionActionIndeterminate,
}

// GetRoutingActionActionEnumValues Enumerates the set of values for RoutingActionActionEnum
func GetRoutingActionActionEnumValues() []RoutingActionActionEnum {
	values := make([]RoutingActionActionEnum, 0)
	for _, v := range mappingRoutingActionActionEnum {
		values = append(values, v)
	}
	return values
}

// GetRoutingActionActionEnumStringValues Enumerates the set of values in String for RoutingActionActionEnum
func GetRoutingActionActionEnumStringValues() []string {
	return []string{
		"FORWARDED",
		"NO_ROUTE",
		"INDETERMINATE",
	}
}

// GetMappingRoutingActionActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRoutingActionActionEnum(val string) (RoutingActionActionEnum, bool) {
	enum, ok := mappingRoutingActionActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
