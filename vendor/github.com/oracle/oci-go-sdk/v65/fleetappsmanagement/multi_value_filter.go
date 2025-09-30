// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MultiValueFilter Content Source details.
type MultiValueFilter struct {

	// Name of report.
	FieldName *string `mandatory:"true" json:"fieldName"`

	// Array of values.
	Values []string `mandatory:"false" json:"values"`

	// Expression for Filter
	Expression MultiValueFilterExpressionEnum `mandatory:"false" json:"expression,omitempty"`
}

// GetFieldName returns FieldName
func (m MultiValueFilter) GetFieldName() *string {
	return m.FieldName
}

func (m MultiValueFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MultiValueFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMultiValueFilterExpressionEnum(string(m.Expression)); !ok && m.Expression != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Expression: %s. Supported values are: %s.", m.Expression, strings.Join(GetMultiValueFilterExpressionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MultiValueFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMultiValueFilter MultiValueFilter
	s := struct {
		DiscriminatorParam string `json:"filterType"`
		MarshalTypeMultiValueFilter
	}{
		"MULTI_VALUE",
		(MarshalTypeMultiValueFilter)(m),
	}

	return json.Marshal(&s)
}

// MultiValueFilterExpressionEnum Enum with underlying type: string
type MultiValueFilterExpressionEnum string

// Set of constants representing the allowable values for MultiValueFilterExpressionEnum
const (
	MultiValueFilterExpressionIn      MultiValueFilterExpressionEnum = "IN"
	MultiValueFilterExpressionNotIn   MultiValueFilterExpressionEnum = "NOT_IN"
	MultiValueFilterExpressionBetween MultiValueFilterExpressionEnum = "BETWEEN"
)

var mappingMultiValueFilterExpressionEnum = map[string]MultiValueFilterExpressionEnum{
	"IN":      MultiValueFilterExpressionIn,
	"NOT_IN":  MultiValueFilterExpressionNotIn,
	"BETWEEN": MultiValueFilterExpressionBetween,
}

var mappingMultiValueFilterExpressionEnumLowerCase = map[string]MultiValueFilterExpressionEnum{
	"in":      MultiValueFilterExpressionIn,
	"not_in":  MultiValueFilterExpressionNotIn,
	"between": MultiValueFilterExpressionBetween,
}

// GetMultiValueFilterExpressionEnumValues Enumerates the set of values for MultiValueFilterExpressionEnum
func GetMultiValueFilterExpressionEnumValues() []MultiValueFilterExpressionEnum {
	values := make([]MultiValueFilterExpressionEnum, 0)
	for _, v := range mappingMultiValueFilterExpressionEnum {
		values = append(values, v)
	}
	return values
}

// GetMultiValueFilterExpressionEnumStringValues Enumerates the set of values in String for MultiValueFilterExpressionEnum
func GetMultiValueFilterExpressionEnumStringValues() []string {
	return []string{
		"IN",
		"NOT_IN",
		"BETWEEN",
	}
}

// GetMappingMultiValueFilterExpressionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMultiValueFilterExpressionEnum(val string) (MultiValueFilterExpressionEnum, bool) {
	enum, ok := mappingMultiValueFilterExpressionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
