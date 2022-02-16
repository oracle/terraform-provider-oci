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

// SetQueryParameterPolicyItem Set will add a new query parameter if it was not in the original request.  If the parameter already exists on the
// request, you can choose to override, append, or skip it.
type SetQueryParameterPolicyItem struct {

	// The case-sensitive name of the query parameter.  This name must be unique across transformation policies.
	Name *string `mandatory:"true" json:"name"`

	// A list of new values.  Each value can be a constant or may include one or more expressions enclosed within
	// ${} delimiters.
	Values []string `mandatory:"true" json:"values"`

	// If a query parameter with the same name already exists in the request, OVERWRITE will overwrite the value,
	// APPEND will append to the existing value, or SKIP will keep the existing value.
	IfExists SetQueryParameterPolicyItemIfExistsEnum `mandatory:"false" json:"ifExists,omitempty"`
}

func (m SetQueryParameterPolicyItem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SetQueryParameterPolicyItem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSetQueryParameterPolicyItemIfExistsEnum(string(m.IfExists)); !ok && m.IfExists != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IfExists: %s. Supported values are: %s.", m.IfExists, strings.Join(GetSetQueryParameterPolicyItemIfExistsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SetQueryParameterPolicyItemIfExistsEnum Enum with underlying type: string
type SetQueryParameterPolicyItemIfExistsEnum string

// Set of constants representing the allowable values for SetQueryParameterPolicyItemIfExistsEnum
const (
	SetQueryParameterPolicyItemIfExistsOverwrite SetQueryParameterPolicyItemIfExistsEnum = "OVERWRITE"
	SetQueryParameterPolicyItemIfExistsAppend    SetQueryParameterPolicyItemIfExistsEnum = "APPEND"
	SetQueryParameterPolicyItemIfExistsSkip      SetQueryParameterPolicyItemIfExistsEnum = "SKIP"
)

var mappingSetQueryParameterPolicyItemIfExistsEnum = map[string]SetQueryParameterPolicyItemIfExistsEnum{
	"OVERWRITE": SetQueryParameterPolicyItemIfExistsOverwrite,
	"APPEND":    SetQueryParameterPolicyItemIfExistsAppend,
	"SKIP":      SetQueryParameterPolicyItemIfExistsSkip,
}

// GetSetQueryParameterPolicyItemIfExistsEnumValues Enumerates the set of values for SetQueryParameterPolicyItemIfExistsEnum
func GetSetQueryParameterPolicyItemIfExistsEnumValues() []SetQueryParameterPolicyItemIfExistsEnum {
	values := make([]SetQueryParameterPolicyItemIfExistsEnum, 0)
	for _, v := range mappingSetQueryParameterPolicyItemIfExistsEnum {
		values = append(values, v)
	}
	return values
}

// GetSetQueryParameterPolicyItemIfExistsEnumStringValues Enumerates the set of values in String for SetQueryParameterPolicyItemIfExistsEnum
func GetSetQueryParameterPolicyItemIfExistsEnumStringValues() []string {
	return []string{
		"OVERWRITE",
		"APPEND",
		"SKIP",
	}
}

// GetMappingSetQueryParameterPolicyItemIfExistsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSetQueryParameterPolicyItemIfExistsEnum(val string) (SetQueryParameterPolicyItemIfExistsEnum, bool) {
	mappingSetQueryParameterPolicyItemIfExistsEnumIgnoreCase := make(map[string]SetQueryParameterPolicyItemIfExistsEnum)
	for k, v := range mappingSetQueryParameterPolicyItemIfExistsEnum {
		mappingSetQueryParameterPolicyItemIfExistsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSetQueryParameterPolicyItemIfExistsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
