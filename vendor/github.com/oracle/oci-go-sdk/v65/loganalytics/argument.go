// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Argument Generic queryString argument.
type Argument interface {
}

type argument struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *argument) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerargument argument
	s := struct {
		Model Unmarshalerargument
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *argument) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "LITERAL":
		mm := LiteralArgument{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FIELD":
		mm := FieldArgument{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Argument: %s.", m.Type)
		return *m, nil
	}
}

func (m argument) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m argument) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ArgumentTypeEnum Enum with underlying type: string
type ArgumentTypeEnum string

// Set of constants representing the allowable values for ArgumentTypeEnum
const (
	ArgumentTypeField   ArgumentTypeEnum = "FIELD"
	ArgumentTypeLiteral ArgumentTypeEnum = "LITERAL"
)

var mappingArgumentTypeEnum = map[string]ArgumentTypeEnum{
	"FIELD":   ArgumentTypeField,
	"LITERAL": ArgumentTypeLiteral,
}

var mappingArgumentTypeEnumLowerCase = map[string]ArgumentTypeEnum{
	"field":   ArgumentTypeField,
	"literal": ArgumentTypeLiteral,
}

// GetArgumentTypeEnumValues Enumerates the set of values for ArgumentTypeEnum
func GetArgumentTypeEnumValues() []ArgumentTypeEnum {
	values := make([]ArgumentTypeEnum, 0)
	for _, v := range mappingArgumentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetArgumentTypeEnumStringValues Enumerates the set of values in String for ArgumentTypeEnum
func GetArgumentTypeEnumStringValues() []string {
	return []string{
		"FIELD",
		"LITERAL",
	}
}

// GetMappingArgumentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingArgumentTypeEnum(val string) (ArgumentTypeEnum, bool) {
	enum, ok := mappingArgumentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
