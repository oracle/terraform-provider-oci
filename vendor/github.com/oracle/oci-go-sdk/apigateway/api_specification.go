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
	"github.com/oracle/oci-go-sdk/common"
)

// ApiSpecification The logical configuration of the API exposed by a deployment.
type ApiSpecification struct {
	RequestPolicies *ApiSpecificationRequestPolicies `mandatory:"false" json:"requestPolicies"`

	LoggingPolicies *ApiSpecificationLoggingPolicies `mandatory:"false" json:"loggingPolicies"`

	// A list of routes that this API exposes.
	Routes []ApiSpecificationRoute `mandatory:"false" json:"routes"`
}

func (m ApiSpecification) String() string {
	return common.PointerString(m)
}
