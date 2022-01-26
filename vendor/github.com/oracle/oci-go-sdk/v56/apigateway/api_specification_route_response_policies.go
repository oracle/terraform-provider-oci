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

// ApiSpecificationRouteResponsePolicies Behavior applied to any responses sent by the API for requests on this route.
type ApiSpecificationRouteResponsePolicies struct {
	HeaderTransformations *HeaderTransformationPolicy `mandatory:"false" json:"headerTransformations"`

	ResponseCacheStore ResponseCacheStorePolicy `mandatory:"false" json:"responseCacheStore"`
}

func (m ApiSpecificationRouteResponsePolicies) String() string {
	return common.PointerString(m)
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
