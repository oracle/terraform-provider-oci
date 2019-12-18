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

// WorkRequestLog A log message from the execution of a work request.
type WorkRequestLog struct {

	// Human-readable log message.
	Message *string `mandatory:"true" json:"message"`

	// The time the log message was written. An RFC3339 formatted datetime string.
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`
}

func (m WorkRequestLog) String() string {
	return common.PointerString(m)
}
