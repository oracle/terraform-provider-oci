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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApiValidationDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingApiValidationDetailSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetApiValidationDetailSeverityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApiValidationDetailSeverityEnum Enum with underlying type: string
type ApiValidationDetailSeverityEnum string

// Set of constants representing the allowable values for ApiValidationDetailSeverityEnum
const (
	ApiValidationDetailSeverityInfo    ApiValidationDetailSeverityEnum = "INFO"
	ApiValidationDetailSeverityWarning ApiValidationDetailSeverityEnum = "WARNING"
	ApiValidationDetailSeverityError   ApiValidationDetailSeverityEnum = "ERROR"
)

var mappingApiValidationDetailSeverityEnum = map[string]ApiValidationDetailSeverityEnum{
	"INFO":    ApiValidationDetailSeverityInfo,
	"WARNING": ApiValidationDetailSeverityWarning,
	"ERROR":   ApiValidationDetailSeverityError,
}

// GetApiValidationDetailSeverityEnumValues Enumerates the set of values for ApiValidationDetailSeverityEnum
func GetApiValidationDetailSeverityEnumValues() []ApiValidationDetailSeverityEnum {
	values := make([]ApiValidationDetailSeverityEnum, 0)
	for _, v := range mappingApiValidationDetailSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetApiValidationDetailSeverityEnumStringValues Enumerates the set of values in String for ApiValidationDetailSeverityEnum
func GetApiValidationDetailSeverityEnumStringValues() []string {
	return []string{
		"INFO",
		"WARNING",
		"ERROR",
	}
}

// GetMappingApiValidationDetailSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApiValidationDetailSeverityEnum(val string) (ApiValidationDetailSeverityEnum, bool) {
	mappingApiValidationDetailSeverityEnumIgnoreCase := make(map[string]ApiValidationDetailSeverityEnum)
	for k, v := range mappingApiValidationDetailSeverityEnum {
		mappingApiValidationDetailSeverityEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingApiValidationDetailSeverityEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
