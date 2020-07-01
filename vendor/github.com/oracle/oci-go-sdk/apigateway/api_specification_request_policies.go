// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/common"
)

// ApiSpecificationRequestPolicies Global behavior applied to all requests received by the API.
type ApiSpecificationRequestPolicies struct {
	Authentication AuthenticationPolicy `mandatory:"false" json:"authentication"`

	RateLimiting *RateLimitingPolicy `mandatory:"false" json:"rateLimiting"`

	Cors *CorsPolicy `mandatory:"false" json:"cors"`
}

func (m ApiSpecificationRequestPolicies) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ApiSpecificationRequestPolicies) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Authentication authenticationpolicy `json:"authentication"`
		RateLimiting   *RateLimitingPolicy  `json:"rateLimiting"`
		Cors           *CorsPolicy          `json:"cors"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Authentication.UnmarshalPolymorphicJSON(model.Authentication.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Authentication = nn.(AuthenticationPolicy)
	} else {
		m.Authentication = nil
	}

	m.RateLimiting = model.RateLimiting

	m.Cors = model.Cors

	return
}
