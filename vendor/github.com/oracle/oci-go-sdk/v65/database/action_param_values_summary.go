// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ActionParamValuesSummary Details of the action parameter and its possible values that is used in listParamsForActionType.
type ActionParamValuesSummary struct {

	// The name of this parameter.
	ParameterName *string `mandatory:"true" json:"parameterName"`

	// The type of the parameter.
	ParameterType ActionParamValuesSummaryParameterTypeEnum `mandatory:"true" json:"parameterType"`

	// Possible values for this parameter. In case of integer it's min and max values.
	ParameterValues []string `mandatory:"true" json:"parameterValues"`

	// Whether this parameter is required or not for this action type.、
	IsRequired *bool `mandatory:"true" json:"isRequired"`

	// The default value for this parameter.
	DefaultValue *string `mandatory:"false" json:"defaultValue"`
}

func (m ActionParamValuesSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ActionParamValuesSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingActionParamValuesSummaryParameterTypeEnum(string(m.ParameterType)); !ok && m.ParameterType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ParameterType: %s. Supported values are: %s.", m.ParameterType, strings.Join(GetActionParamValuesSummaryParameterTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ActionParamValuesSummaryParameterTypeEnum Enum with underlying type: string
type ActionParamValuesSummaryParameterTypeEnum string

// Set of constants representing the allowable values for ActionParamValuesSummaryParameterTypeEnum
const (
	ActionParamValuesSummaryParameterTypeBoolean ActionParamValuesSummaryParameterTypeEnum = "BOOLEAN"
	ActionParamValuesSummaryParameterTypeString  ActionParamValuesSummaryParameterTypeEnum = "STRING"
	ActionParamValuesSummaryParameterTypeInteger ActionParamValuesSummaryParameterTypeEnum = "INTEGER"
)

var mappingActionParamValuesSummaryParameterTypeEnum = map[string]ActionParamValuesSummaryParameterTypeEnum{
	"BOOLEAN": ActionParamValuesSummaryParameterTypeBoolean,
	"STRING":  ActionParamValuesSummaryParameterTypeString,
	"INTEGER": ActionParamValuesSummaryParameterTypeInteger,
}

var mappingActionParamValuesSummaryParameterTypeEnumLowerCase = map[string]ActionParamValuesSummaryParameterTypeEnum{
	"boolean": ActionParamValuesSummaryParameterTypeBoolean,
	"string":  ActionParamValuesSummaryParameterTypeString,
	"integer": ActionParamValuesSummaryParameterTypeInteger,
}

// GetActionParamValuesSummaryParameterTypeEnumValues Enumerates the set of values for ActionParamValuesSummaryParameterTypeEnum
func GetActionParamValuesSummaryParameterTypeEnumValues() []ActionParamValuesSummaryParameterTypeEnum {
	values := make([]ActionParamValuesSummaryParameterTypeEnum, 0)
	for _, v := range mappingActionParamValuesSummaryParameterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetActionParamValuesSummaryParameterTypeEnumStringValues Enumerates the set of values in String for ActionParamValuesSummaryParameterTypeEnum
func GetActionParamValuesSummaryParameterTypeEnumStringValues() []string {
	return []string{
		"BOOLEAN",
		"STRING",
		"INTEGER",
	}
}

// GetMappingActionParamValuesSummaryParameterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingActionParamValuesSummaryParameterTypeEnum(val string) (ActionParamValuesSummaryParameterTypeEnum, bool) {
	enum, ok := mappingActionParamValuesSummaryParameterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
