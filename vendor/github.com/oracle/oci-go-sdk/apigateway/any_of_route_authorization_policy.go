// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
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

// AnyOfRouteAuthorizationPolicy If authentication has been performed, validate whether the request scope (if any) applies to this route.
type AnyOfRouteAuthorizationPolicy struct {

	// A user whose scope includes any of these access ranges is allowed on
	// this route. Access ranges are case-sensitive.
	AllowedScope []string `mandatory:"true" json:"allowedScope"`
}

func (m AnyOfRouteAuthorizationPolicy) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AnyOfRouteAuthorizationPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAnyOfRouteAuthorizationPolicy AnyOfRouteAuthorizationPolicy
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAnyOfRouteAuthorizationPolicy
	}{
		"ANY_OF",
		(MarshalTypeAnyOfRouteAuthorizationPolicy)(m),
	}

	return json.Marshal(&s)
}
