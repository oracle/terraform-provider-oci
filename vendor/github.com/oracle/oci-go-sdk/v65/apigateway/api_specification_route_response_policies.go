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

// ApiSpecificationRouteResponsePolicies Behavior applied to any responses sent by the API for requests on this route.
type ApiSpecificationRouteResponsePolicies struct {
	HeaderTransformations *HeaderTransformationPolicy `mandatory:"false" json:"headerTransformations"`

	ResponseCacheStore ResponseCacheStorePolicy `mandatory:"false" json:"responseCacheStore"`
}

func (m ApiSpecificationRouteResponsePolicies) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApiSpecificationRouteResponsePolicies) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ApiSpecificationRouteResponsePolicies) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		HeaderTransformations *HeaderTransformationPolicy `json:"headerTransformations"`
		ResponseCacheStore    responsecachestorepolicy    `json:"responseCacheStore"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.HeaderTransformations = model.HeaderTransformations

	nn, e = model.ResponseCacheStore.UnmarshalPolymorphicJSON(model.ResponseCacheStore.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ResponseCacheStore = nn.(ResponseCacheStorePolicy)
	} else {
		m.ResponseCacheStore = nil
	}

	return
}
