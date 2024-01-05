// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ApiSpecificationRoute A single route that forwards requests to a particular backend and may contain some additional policies.
type ApiSpecificationRoute struct {

	// A URL path pattern that must be matched on this route. The path pattern may contain a subset of RFC 6570 identifiers
	// to allow wildcard and parameterized matching.
	Path *string `mandatory:"true" json:"path"`

	Backend ApiSpecificationRouteBackend `mandatory:"true" json:"backend"`

	// A list of allowed methods on this route.
	Methods []ApiSpecificationRouteMethodsEnum `mandatory:"false" json:"methods,omitempty"`

	RequestPolicies *ApiSpecificationRouteRequestPolicies `mandatory:"false" json:"requestPolicies"`

	ResponsePolicies *ApiSpecificationRouteResponsePolicies `mandatory:"false" json:"responsePolicies"`

	LoggingPolicies *ApiSpecificationLoggingPolicies `mandatory:"false" json:"loggingPolicies"`
}

func (m ApiSpecificationRoute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApiSpecificationRoute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.Methods {
		if _, ok := GetMappingApiSpecificationRouteMethodsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Methods: %s. Supported values are: %s.", val, strings.Join(GetApiSpecificationRouteMethodsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ApiSpecificationRoute) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Methods          []ApiSpecificationRouteMethodsEnum     `json:"methods"`
		RequestPolicies  *ApiSpecificationRouteRequestPolicies  `json:"requestPolicies"`
		ResponsePolicies *ApiSpecificationRouteResponsePolicies `json:"responsePolicies"`
		LoggingPolicies  *ApiSpecificationLoggingPolicies       `json:"loggingPolicies"`
		Path             *string                                `json:"path"`
		Backend          apispecificationroutebackend           `json:"backend"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Methods = make([]ApiSpecificationRouteMethodsEnum, len(model.Methods))
	copy(m.Methods, model.Methods)
	m.RequestPolicies = model.RequestPolicies

	m.ResponsePolicies = model.ResponsePolicies

	m.LoggingPolicies = model.LoggingPolicies

	m.Path = model.Path

	nn, e = model.Backend.UnmarshalPolymorphicJSON(model.Backend.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Backend = nn.(ApiSpecificationRouteBackend)
	} else {
		m.Backend = nil
	}

	return
}

// ApiSpecificationRouteMethodsEnum Enum with underlying type: string
type ApiSpecificationRouteMethodsEnum string

// Set of constants representing the allowable values for ApiSpecificationRouteMethodsEnum
const (
	ApiSpecificationRouteMethodsAny     ApiSpecificationRouteMethodsEnum = "ANY"
	ApiSpecificationRouteMethodsHead    ApiSpecificationRouteMethodsEnum = "HEAD"
	ApiSpecificationRouteMethodsGet     ApiSpecificationRouteMethodsEnum = "GET"
	ApiSpecificationRouteMethodsPost    ApiSpecificationRouteMethodsEnum = "POST"
	ApiSpecificationRouteMethodsPut     ApiSpecificationRouteMethodsEnum = "PUT"
	ApiSpecificationRouteMethodsPatch   ApiSpecificationRouteMethodsEnum = "PATCH"
	ApiSpecificationRouteMethodsDelete  ApiSpecificationRouteMethodsEnum = "DELETE"
	ApiSpecificationRouteMethodsOptions ApiSpecificationRouteMethodsEnum = "OPTIONS"
)

var mappingApiSpecificationRouteMethodsEnum = map[string]ApiSpecificationRouteMethodsEnum{
	"ANY":     ApiSpecificationRouteMethodsAny,
	"HEAD":    ApiSpecificationRouteMethodsHead,
	"GET":     ApiSpecificationRouteMethodsGet,
	"POST":    ApiSpecificationRouteMethodsPost,
	"PUT":     ApiSpecificationRouteMethodsPut,
	"PATCH":   ApiSpecificationRouteMethodsPatch,
	"DELETE":  ApiSpecificationRouteMethodsDelete,
	"OPTIONS": ApiSpecificationRouteMethodsOptions,
}

var mappingApiSpecificationRouteMethodsEnumLowerCase = map[string]ApiSpecificationRouteMethodsEnum{
	"any":     ApiSpecificationRouteMethodsAny,
	"head":    ApiSpecificationRouteMethodsHead,
	"get":     ApiSpecificationRouteMethodsGet,
	"post":    ApiSpecificationRouteMethodsPost,
	"put":     ApiSpecificationRouteMethodsPut,
	"patch":   ApiSpecificationRouteMethodsPatch,
	"delete":  ApiSpecificationRouteMethodsDelete,
	"options": ApiSpecificationRouteMethodsOptions,
}

// GetApiSpecificationRouteMethodsEnumValues Enumerates the set of values for ApiSpecificationRouteMethodsEnum
func GetApiSpecificationRouteMethodsEnumValues() []ApiSpecificationRouteMethodsEnum {
	values := make([]ApiSpecificationRouteMethodsEnum, 0)
	for _, v := range mappingApiSpecificationRouteMethodsEnum {
		values = append(values, v)
	}
	return values
}

// GetApiSpecificationRouteMethodsEnumStringValues Enumerates the set of values in String for ApiSpecificationRouteMethodsEnum
func GetApiSpecificationRouteMethodsEnumStringValues() []string {
	return []string{
		"ANY",
		"HEAD",
		"GET",
		"POST",
		"PUT",
		"PATCH",
		"DELETE",
		"OPTIONS",
	}
}

// GetMappingApiSpecificationRouteMethodsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApiSpecificationRouteMethodsEnum(val string) (ApiSpecificationRouteMethodsEnum, bool) {
	enum, ok := mappingApiSpecificationRouteMethodsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
