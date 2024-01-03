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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExecutionLogPolicy Configures the logging policies for the execution logs of an API Deployment.
type ExecutionLogPolicy struct {

	// Enables pushing of execution logs to the legacy OCI Object Storage log archival bucket.
	// Oracle recommends using the OCI Logging service to enable, retrieve, and query execution logs
	// for an API Deployment. If there is an active log object for the API Deployment and its
	// category is set to 'execution' in OCI Logging service, the logs will not be uploaded to the legacy
	// OCI Object Storage log archival bucket.
	// Please note that the functionality to push to the legacy OCI Object Storage log
	// archival bucket has been deprecated and will be removed in the future.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Specifies the log level used to control logging output of execution logs.
	// Enabling logging at a given level also enables logging at all higher levels.
	LogLevel ExecutionLogPolicyLogLevelEnum `mandatory:"false" json:"logLevel,omitempty"`
}

func (m ExecutionLogPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecutionLogPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExecutionLogPolicyLogLevelEnum(string(m.LogLevel)); !ok && m.LogLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LogLevel: %s. Supported values are: %s.", m.LogLevel, strings.Join(GetExecutionLogPolicyLogLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExecutionLogPolicyLogLevelEnum Enum with underlying type: string
type ExecutionLogPolicyLogLevelEnum string

// Set of constants representing the allowable values for ExecutionLogPolicyLogLevelEnum
const (
	ExecutionLogPolicyLogLevelInfo  ExecutionLogPolicyLogLevelEnum = "INFO"
	ExecutionLogPolicyLogLevelWarn  ExecutionLogPolicyLogLevelEnum = "WARN"
	ExecutionLogPolicyLogLevelError ExecutionLogPolicyLogLevelEnum = "ERROR"
)

var mappingExecutionLogPolicyLogLevelEnum = map[string]ExecutionLogPolicyLogLevelEnum{
	"INFO":  ExecutionLogPolicyLogLevelInfo,
	"WARN":  ExecutionLogPolicyLogLevelWarn,
	"ERROR": ExecutionLogPolicyLogLevelError,
}

var mappingExecutionLogPolicyLogLevelEnumLowerCase = map[string]ExecutionLogPolicyLogLevelEnum{
	"info":  ExecutionLogPolicyLogLevelInfo,
	"warn":  ExecutionLogPolicyLogLevelWarn,
	"error": ExecutionLogPolicyLogLevelError,
}

// GetExecutionLogPolicyLogLevelEnumValues Enumerates the set of values for ExecutionLogPolicyLogLevelEnum
func GetExecutionLogPolicyLogLevelEnumValues() []ExecutionLogPolicyLogLevelEnum {
	values := make([]ExecutionLogPolicyLogLevelEnum, 0)
	for _, v := range mappingExecutionLogPolicyLogLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetExecutionLogPolicyLogLevelEnumStringValues Enumerates the set of values in String for ExecutionLogPolicyLogLevelEnum
func GetExecutionLogPolicyLogLevelEnumStringValues() []string {
	return []string{
		"INFO",
		"WARN",
		"ERROR",
	}
}

// GetMappingExecutionLogPolicyLogLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecutionLogPolicyLogLevelEnum(val string) (ExecutionLogPolicyLogLevelEnum, bool) {
	enum, ok := mappingExecutionLogPolicyLogLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
