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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ApiSpecificationRouteRequestPolicies Behavior applied to any requests received by the API on this route.
type ApiSpecificationRouteRequestPolicies struct {
	Authorization RouteAuthorizationPolicy `mandatory:"false" json:"authorization"`

	Cors *CorsPolicy `mandatory:"false" json:"cors"`

	QueryParameterValidations *QueryParameterValidationRequestPolicy `mandatory:"false" json:"queryParameterValidations"`

	HeaderValidations *HeaderValidationRequestPolicy `mandatory:"false" json:"headerValidations"`

	BodyValidation *BodyValidationRequestPolicy `mandatory:"false" json:"bodyValidation"`

	HeaderTransformations *HeaderTransformationPolicy `mandatory:"false" json:"headerTransformations"`

	QueryParameterTransformations *QueryParameterTransformationPolicy `mandatory:"false" json:"queryParameterTransformations"`

	ResponseCacheLookup ResponseCacheLookupPolicy `mandatory:"false" json:"responseCacheLookup"`
}

func (m ApiSpecificationRouteRequestPolicies) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ApiSpecificationRouteRequestPolicies) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Authorization                 routeauthorizationpolicy               `json:"authorization"`
		Cors                          *CorsPolicy                            `json:"cors"`
		QueryParameterValidations     *QueryParameterValidationRequestPolicy `json:"queryParameterValidations"`
		HeaderValidations             *HeaderValidationRequestPolicy         `json:"headerValidations"`
		BodyValidation                *BodyValidationRequestPolicy           `json:"bodyValidation"`
		HeaderTransformations         *HeaderTransformationPolicy            `json:"headerTransformations"`
		QueryParameterTransformations *QueryParameterTransformationPolicy    `json:"queryParameterTransformations"`
		ResponseCacheLookup           responsecachelookuppolicy              `json:"responseCacheLookup"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Authorization.UnmarshalPolymorphicJSON(model.Authorization.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Authorization = nn.(RouteAuthorizationPolicy)
	} else {
		m.Authorization = nil
	}

	m.Cors = model.Cors

	m.QueryParameterValidations = model.QueryParameterValidations

	m.HeaderValidations = model.HeaderValidations

	m.BodyValidation = model.BodyValidation

	m.HeaderTransformations = model.HeaderTransformations

	m.QueryParameterTransformations = model.QueryParameterTransformations

	nn, e = model.ResponseCacheLookup.UnmarshalPolymorphicJSON(model.ResponseCacheLookup.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ResponseCacheLookup = nn.(ResponseCacheLookupPolicy)
	} else {
		m.ResponseCacheLookup = nil
	}

	return
}
