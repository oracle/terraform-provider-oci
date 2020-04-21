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
	"github.com/oracle/oci-go-sdk/common"
)

// CorsPolicy Enable CORS (Cross-Origin-Resource-Sharing) request handling.
type CorsPolicy struct {

	// The list of allowed origins that the CORS handler will use to respond to CORS requests. The gateway will
	// send the Access-Control-Allow-Origin header with the best origin match for the circumstances. '*' will match
	// any origins, and 'null' will match queries from 'file:' origins. All other origins must be qualified with the
	// scheme, full hostname, and port if necessary.
	AllowedOrigins []string `mandatory:"true" json:"allowedOrigins"`

	// The list of allowed HTTP methods that will be returned for the preflight OPTIONS request in the
	// Access-Control-Allow-Methods header. '*' will allow all methods.
	AllowedMethods []string `mandatory:"false" json:"allowedMethods"`

	// The list of headers that will be allowed from the client via the Access-Control-Allow-Headers header.
	// '*' will allow all headers.
	AllowedHeaders []string `mandatory:"false" json:"allowedHeaders"`

	// The list of headers that the client will be allowed to see from the response as indicated by the
	// Access-Control-Expose-Headers header. '*' will expose all headers.
	ExposedHeaders []string `mandatory:"false" json:"exposedHeaders"`

	// Whether to send the Access-Control-Allow-Credentials header to allow CORS requests with cookies.
	IsAllowCredentialsEnabled *bool `mandatory:"false" json:"isAllowCredentialsEnabled"`

	// The time in seconds for the client to cache preflight responses. This is sent as the Access-Control-Max-Age
	// if greater than 0.
	MaxAgeInSeconds *int `mandatory:"false" json:"maxAgeInSeconds"`
}

func (m CorsPolicy) String() string {
	return common.PointerString(m)
}
