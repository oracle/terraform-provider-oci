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

// ExecutionLogPolicy Configures the pushing of execution logs to OCI Public Logging.
type ExecutionLogPolicy struct {

	// Enables pushing of execution logs to OCI Public Logging.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Specifies the logging level, which affects the log entries pushed to
	// OCI Public Logging if `isEnabled` is set to True.
	LogLevel ExecutionLogPolicyLogLevelEnum `mandatory:"false" json:"logLevel,omitempty"`
}

func (m ExecutionLogPolicy) String() string {
	return common.PointerString(m)
}

// ExecutionLogPolicyLogLevelEnum Enum with underlying type: string
type ExecutionLogPolicyLogLevelEnum string

// Set of constants representing the allowable values for ExecutionLogPolicyLogLevelEnum
const (
	ExecutionLogPolicyLogLevelInfo  ExecutionLogPolicyLogLevelEnum = "INFO"
	ExecutionLogPolicyLogLevelWarn  ExecutionLogPolicyLogLevelEnum = "WARN"
	ExecutionLogPolicyLogLevelError ExecutionLogPolicyLogLevelEnum = "ERROR"
)

var mappingExecutionLogPolicyLogLevel = map[string]ExecutionLogPolicyLogLevelEnum{
	"INFO":  ExecutionLogPolicyLogLevelInfo,
	"WARN":  ExecutionLogPolicyLogLevelWarn,
	"ERROR": ExecutionLogPolicyLogLevelError,
}

// GetExecutionLogPolicyLogLevelEnumValues Enumerates the set of values for ExecutionLogPolicyLogLevelEnum
func GetExecutionLogPolicyLogLevelEnumValues() []ExecutionLogPolicyLogLevelEnum {
	values := make([]ExecutionLogPolicyLogLevelEnum, 0)
	for _, v := range mappingExecutionLogPolicyLogLevel {
		values = append(values, v)
	}
	return values
}
