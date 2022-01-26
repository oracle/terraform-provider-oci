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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ApiValidationDetail Detail of a single error or warning.
type ApiValidationDetail struct {

	// Description of the warning/error.
	Msg *string `mandatory:"false" json:"msg"`

	// Severity of the issue.
	Severity ApiValidationDetailSeverityEnum `mandatory:"false" json:"severity,omitempty"`

	// Position of the issue in the specification file (line, column).
	Src [][]float32 `mandatory:"false" json:"src"`
}

func (m ApiValidationDetail) String() string {
	return common.PointerString(m)
}

// ApiValidationDetailSeverityEnum Enum with underlying type: string
type ApiValidationDetailSeverityEnum string

// Set of constants representing the allowable values for ApiValidationDetailSeverityEnum
const (
	ApiValidationDetailSeverityInfo    ApiValidationDetailSeverityEnum = "INFO"
	ApiValidationDetailSeverityWarning ApiValidationDetailSeverityEnum = "WARNING"
	ApiValidationDetailSeverityError   ApiValidationDetailSeverityEnum = "ERROR"
)

var mappingApiValidationDetailSeverity = map[string]ApiValidationDetailSeverityEnum{
	"INFO":    ApiValidationDetailSeverityInfo,
	"WARNING": ApiValidationDetailSeverityWarning,
	"ERROR":   ApiValidationDetailSeverityError,
}

// GetApiValidationDetailSeverityEnumValues Enumerates the set of values for ApiValidationDetailSeverityEnum
func GetApiValidationDetailSeverityEnumValues() []ApiValidationDetailSeverityEnum {
	values := make([]ApiValidationDetailSeverityEnum, 0)
	for _, v := range mappingApiValidationDetailSeverity {
		values = append(values, v)
	}
	return values
}
