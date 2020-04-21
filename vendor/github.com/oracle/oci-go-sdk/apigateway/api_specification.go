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
