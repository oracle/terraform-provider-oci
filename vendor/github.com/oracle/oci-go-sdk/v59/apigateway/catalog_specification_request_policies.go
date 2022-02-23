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
	"github.com/oracle/oci-go-sdk/v59/common"
	"strings"
)

// CatalogSpecificationRequestPolicies Behavior applied to any requests received by the API Catalog endpoint.
type CatalogSpecificationRequestPolicies struct {
	Cors *CorsPolicy `mandatory:"false" json:"cors"`

	Authentication AuthenticationPolicy `mandatory:"false" json:"authentication"`

	RateLimiting *RateLimitingPolicy `mandatory:"false" json:"rateLimiting"`
}

func (m CatalogSpecificationRequestPolicies) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CatalogSpecificationRequestPolicies) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CatalogSpecificationRequestPolicies) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Cors           *CorsPolicy          `json:"cors"`
		Authentication authenticationpolicy `json:"authentication"`
		RateLimiting   *RateLimitingPolicy  `json:"rateLimiting"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Cors = model.Cors

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

	return
}
