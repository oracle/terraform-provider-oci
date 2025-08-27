// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Scheduler API
//
// Use the Resource scheduler API to manage schedules, to perform actions on a collection of resources.
//

package resourcescheduler

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Parameter This is a generic input parameter to use when acting on the resource.
type Parameter interface {
}

type parameter struct {
	JsonData      []byte
	ParameterType string `json:"parameterType"`
}

// UnmarshalJSON unmarshals json
func (m *parameter) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerparameter parameter
	s := struct {
		Model Unmarshalerparameter
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ParameterType = s.Model.ParameterType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *parameter) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ParameterType {
	case "QUERY":
		mm := QueryParameter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HEADER":
		mm := HeaderParameter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BODY":
		mm := BodyParameter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PATH":
		mm := PathParameter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for Parameter: %s.", m.ParameterType)
		return *m, nil
	}
}

func (m parameter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m parameter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ParameterParameterTypeEnum Enum with underlying type: string
type ParameterParameterTypeEnum string

// Set of constants representing the allowable values for ParameterParameterTypeEnum
const (
	ParameterParameterTypeHeader ParameterParameterTypeEnum = "HEADER"
	ParameterParameterTypeBody   ParameterParameterTypeEnum = "BODY"
	ParameterParameterTypePath   ParameterParameterTypeEnum = "PATH"
	ParameterParameterTypeQuery  ParameterParameterTypeEnum = "QUERY"
)

var mappingParameterParameterTypeEnum = map[string]ParameterParameterTypeEnum{
	"HEADER": ParameterParameterTypeHeader,
	"BODY":   ParameterParameterTypeBody,
	"PATH":   ParameterParameterTypePath,
	"QUERY":  ParameterParameterTypeQuery,
}

var mappingParameterParameterTypeEnumLowerCase = map[string]ParameterParameterTypeEnum{
	"header": ParameterParameterTypeHeader,
	"body":   ParameterParameterTypeBody,
	"path":   ParameterParameterTypePath,
	"query":  ParameterParameterTypeQuery,
}

// GetParameterParameterTypeEnumValues Enumerates the set of values for ParameterParameterTypeEnum
func GetParameterParameterTypeEnumValues() []ParameterParameterTypeEnum {
	values := make([]ParameterParameterTypeEnum, 0)
	for _, v := range mappingParameterParameterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetParameterParameterTypeEnumStringValues Enumerates the set of values in String for ParameterParameterTypeEnum
func GetParameterParameterTypeEnumStringValues() []string {
	return []string{
		"HEADER",
		"BODY",
		"PATH",
		"QUERY",
	}
}

// GetMappingParameterParameterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingParameterParameterTypeEnum(val string) (ParameterParameterTypeEnum, bool) {
	enum, ok := mappingParameterParameterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
