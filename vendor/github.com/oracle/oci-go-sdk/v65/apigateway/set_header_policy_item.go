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

// SetHeaderPolicyItem Set will add a new header if it was not in the original request.  If the header already exists on the
// request, you can choose to override, append, or skip it.
type SetHeaderPolicyItem struct {

	// The case-insensitive name of the header.  This name must be unique across transformation policies.
	Name *string `mandatory:"true" json:"name"`

	// A list of new values.  Each value can be a constant or may include one or more expressions enclosed within
	// ${} delimiters.
	Values []string `mandatory:"true" json:"values"`

	// If a header with the same name already exists in the request, OVERWRITE will overwrite the value,
	// APPEND will append to the existing value, or SKIP will keep the existing value.
	IfExists SetHeaderPolicyItemIfExistsEnum `mandatory:"false" json:"ifExists,omitempty"`
}

func (m SetHeaderPolicyItem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SetHeaderPolicyItem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSetHeaderPolicyItemIfExistsEnum(string(m.IfExists)); !ok && m.IfExists != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IfExists: %s. Supported values are: %s.", m.IfExists, strings.Join(GetSetHeaderPolicyItemIfExistsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SetHeaderPolicyItemIfExistsEnum Enum with underlying type: string
type SetHeaderPolicyItemIfExistsEnum string

// Set of constants representing the allowable values for SetHeaderPolicyItemIfExistsEnum
const (
	SetHeaderPolicyItemIfExistsOverwrite SetHeaderPolicyItemIfExistsEnum = "OVERWRITE"
	SetHeaderPolicyItemIfExistsAppend    SetHeaderPolicyItemIfExistsEnum = "APPEND"
	SetHeaderPolicyItemIfExistsSkip      SetHeaderPolicyItemIfExistsEnum = "SKIP"
)

var mappingSetHeaderPolicyItemIfExistsEnum = map[string]SetHeaderPolicyItemIfExistsEnum{
	"OVERWRITE": SetHeaderPolicyItemIfExistsOverwrite,
	"APPEND":    SetHeaderPolicyItemIfExistsAppend,
	"SKIP":      SetHeaderPolicyItemIfExistsSkip,
}

var mappingSetHeaderPolicyItemIfExistsEnumLowerCase = map[string]SetHeaderPolicyItemIfExistsEnum{
	"overwrite": SetHeaderPolicyItemIfExistsOverwrite,
	"append":    SetHeaderPolicyItemIfExistsAppend,
	"skip":      SetHeaderPolicyItemIfExistsSkip,
}

// GetSetHeaderPolicyItemIfExistsEnumValues Enumerates the set of values for SetHeaderPolicyItemIfExistsEnum
func GetSetHeaderPolicyItemIfExistsEnumValues() []SetHeaderPolicyItemIfExistsEnum {
	values := make([]SetHeaderPolicyItemIfExistsEnum, 0)
	for _, v := range mappingSetHeaderPolicyItemIfExistsEnum {
		values = append(values, v)
	}
	return values
}

// GetSetHeaderPolicyItemIfExistsEnumStringValues Enumerates the set of values in String for SetHeaderPolicyItemIfExistsEnum
func GetSetHeaderPolicyItemIfExistsEnumStringValues() []string {
	return []string{
		"OVERWRITE",
		"APPEND",
		"SKIP",
	}
}

// GetMappingSetHeaderPolicyItemIfExistsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSetHeaderPolicyItemIfExistsEnum(val string) (SetHeaderPolicyItemIfExistsEnum, bool) {
	enum, ok := mappingSetHeaderPolicyItemIfExistsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
