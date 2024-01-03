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

// FilterQueryParameterPolicy Filter parameters from the query string as they pass through the gateway.  The gateway applies filters after other
// transformations, so any parameters set or renamed must also be listed here when using an ALLOW type policy.
type FilterQueryParameterPolicy struct {

	// BLOCK drops any query parameters that are in the list of items, so it acts as an exclusion list.  ALLOW
	// permits only the parameters in the list and removes all others, so it acts as an inclusion list.
	Type FilterQueryParameterPolicyTypeEnum `mandatory:"true" json:"type"`

	// The list of query parameters.
	Items []FilterQueryParameterPolicyItem `mandatory:"true" json:"items"`
}

func (m FilterQueryParameterPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FilterQueryParameterPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFilterQueryParameterPolicyTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetFilterQueryParameterPolicyTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FilterQueryParameterPolicyTypeEnum Enum with underlying type: string
type FilterQueryParameterPolicyTypeEnum string

// Set of constants representing the allowable values for FilterQueryParameterPolicyTypeEnum
const (
	FilterQueryParameterPolicyTypeAllow FilterQueryParameterPolicyTypeEnum = "ALLOW"
	FilterQueryParameterPolicyTypeBlock FilterQueryParameterPolicyTypeEnum = "BLOCK"
)

var mappingFilterQueryParameterPolicyTypeEnum = map[string]FilterQueryParameterPolicyTypeEnum{
	"ALLOW": FilterQueryParameterPolicyTypeAllow,
	"BLOCK": FilterQueryParameterPolicyTypeBlock,
}

var mappingFilterQueryParameterPolicyTypeEnumLowerCase = map[string]FilterQueryParameterPolicyTypeEnum{
	"allow": FilterQueryParameterPolicyTypeAllow,
	"block": FilterQueryParameterPolicyTypeBlock,
}

// GetFilterQueryParameterPolicyTypeEnumValues Enumerates the set of values for FilterQueryParameterPolicyTypeEnum
func GetFilterQueryParameterPolicyTypeEnumValues() []FilterQueryParameterPolicyTypeEnum {
	values := make([]FilterQueryParameterPolicyTypeEnum, 0)
	for _, v := range mappingFilterQueryParameterPolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFilterQueryParameterPolicyTypeEnumStringValues Enumerates the set of values in String for FilterQueryParameterPolicyTypeEnum
func GetFilterQueryParameterPolicyTypeEnumStringValues() []string {
	return []string{
		"ALLOW",
		"BLOCK",
	}
}

// GetMappingFilterQueryParameterPolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFilterQueryParameterPolicyTypeEnum(val string) (FilterQueryParameterPolicyTypeEnum, bool) {
	enum, ok := mappingFilterQueryParameterPolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
