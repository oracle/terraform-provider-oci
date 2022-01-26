// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
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
		return *m, nil
	}
}

func (m argument) String() string {
	return common.PointerString(m)
}

// ArgumentTypeEnum Enum with underlying type: string
type ArgumentTypeEnum string

// Set of constants representing the allowable values for ArgumentTypeEnum
const (
	ArgumentTypeField   ArgumentTypeEnum = "FIELD"
	ArgumentTypeLiteral ArgumentTypeEnum = "LITERAL"
)

var mappingArgumentType = map[string]ArgumentTypeEnum{
	"FIELD":   ArgumentTypeField,
	"LITERAL": ArgumentTypeLiteral,
}

// GetArgumentTypeEnumValues Enumerates the set of values for ArgumentTypeEnum
func GetArgumentTypeEnumValues() []ArgumentTypeEnum {
	values := make([]ArgumentTypeEnum, 0)
	for _, v := range mappingArgumentType {
		values = append(values, v)
	}
	return values
}
