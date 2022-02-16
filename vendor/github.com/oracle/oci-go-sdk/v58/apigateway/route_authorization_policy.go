// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// RouteAuthorizationPolicy If authentication has been performed, validate whether the request scope (if any) applies to this route.
// If no RouteAuthorizationPolicy is defined for a route, a policy with a type of AUTHENTICATION_ONLY is applied.
type RouteAuthorizationPolicy interface {
}

type routeauthorizationpolicy struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *routeauthorizationpolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerrouteauthorizationpolicy routeauthorizationpolicy
	s := struct {
		Model Unmarshalerrouteauthorizationpolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *routeauthorizationpolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "ANY_OF":
		mm := AnyOfRouteAuthorizationPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ANONYMOUS":
		mm := AnonymousRouteAuthorizationPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AUTHENTICATION_ONLY":
		mm := AuthenticationOnlyRouteAuthorizationPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m routeauthorizationpolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m routeauthorizationpolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RouteAuthorizationPolicyTypeEnum Enum with underlying type: string
type RouteAuthorizationPolicyTypeEnum string

// Set of constants representing the allowable values for RouteAuthorizationPolicyTypeEnum
const (
	RouteAuthorizationPolicyTypeAnonymous          RouteAuthorizationPolicyTypeEnum = "ANONYMOUS"
	RouteAuthorizationPolicyTypeAnyOf              RouteAuthorizationPolicyTypeEnum = "ANY_OF"
	RouteAuthorizationPolicyTypeAuthenticationOnly RouteAuthorizationPolicyTypeEnum = "AUTHENTICATION_ONLY"
)

var mappingRouteAuthorizationPolicyTypeEnum = map[string]RouteAuthorizationPolicyTypeEnum{
	"ANONYMOUS":           RouteAuthorizationPolicyTypeAnonymous,
	"ANY_OF":              RouteAuthorizationPolicyTypeAnyOf,
	"AUTHENTICATION_ONLY": RouteAuthorizationPolicyTypeAuthenticationOnly,
}

// GetRouteAuthorizationPolicyTypeEnumValues Enumerates the set of values for RouteAuthorizationPolicyTypeEnum
func GetRouteAuthorizationPolicyTypeEnumValues() []RouteAuthorizationPolicyTypeEnum {
	values := make([]RouteAuthorizationPolicyTypeEnum, 0)
	for _, v := range mappingRouteAuthorizationPolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRouteAuthorizationPolicyTypeEnumStringValues Enumerates the set of values in String for RouteAuthorizationPolicyTypeEnum
func GetRouteAuthorizationPolicyTypeEnumStringValues() []string {
	return []string{
		"ANONYMOUS",
		"ANY_OF",
		"AUTHENTICATION_ONLY",
	}
}

// GetMappingRouteAuthorizationPolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRouteAuthorizationPolicyTypeEnum(val string) (RouteAuthorizationPolicyTypeEnum, bool) {
	mappingRouteAuthorizationPolicyTypeEnumIgnoreCase := make(map[string]RouteAuthorizationPolicyTypeEnum)
	for k, v := range mappingRouteAuthorizationPolicyTypeEnum {
		mappingRouteAuthorizationPolicyTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingRouteAuthorizationPolicyTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
