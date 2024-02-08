// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Connector Hub API
//
// Use the Service Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Service Connector Hub, see
// Service Connector Hub Overview (https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/overview.htm).
//

package sch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DimensionValueDetails Instructions for extracting the value corresponding to the specified dimension key: Either extract the value as-is (static) or derive the value from a path (evaluated).
type DimensionValueDetails interface {
}

type dimensionvaluedetails struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *dimensionvaluedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdimensionvaluedetails dimensionvaluedetails
	s := struct {
		Model Unmarshalerdimensionvaluedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dimensionvaluedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "static":
		mm := StaticDimensionValue{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "jmesPath":
		mm := JmesPathDimensionValue{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DimensionValueDetails: %s.", m.Kind)
		return *m, nil
	}
}

func (m dimensionvaluedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dimensionvaluedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DimensionValueDetailsKindEnum Enum with underlying type: string
type DimensionValueDetailsKindEnum string

// Set of constants representing the allowable values for DimensionValueDetailsKindEnum
const (
	DimensionValueDetailsKindJmespath DimensionValueDetailsKindEnum = "jmesPath"
	DimensionValueDetailsKindStatic   DimensionValueDetailsKindEnum = "static"
)

var mappingDimensionValueDetailsKindEnum = map[string]DimensionValueDetailsKindEnum{
	"jmesPath": DimensionValueDetailsKindJmespath,
	"static":   DimensionValueDetailsKindStatic,
}

var mappingDimensionValueDetailsKindEnumLowerCase = map[string]DimensionValueDetailsKindEnum{
	"jmespath": DimensionValueDetailsKindJmespath,
	"static":   DimensionValueDetailsKindStatic,
}

// GetDimensionValueDetailsKindEnumValues Enumerates the set of values for DimensionValueDetailsKindEnum
func GetDimensionValueDetailsKindEnumValues() []DimensionValueDetailsKindEnum {
	values := make([]DimensionValueDetailsKindEnum, 0)
	for _, v := range mappingDimensionValueDetailsKindEnum {
		values = append(values, v)
	}
	return values
}

// GetDimensionValueDetailsKindEnumStringValues Enumerates the set of values in String for DimensionValueDetailsKindEnum
func GetDimensionValueDetailsKindEnumStringValues() []string {
	return []string{
		"jmesPath",
		"static",
	}
}

// GetMappingDimensionValueDetailsKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDimensionValueDetailsKindEnum(val string) (DimensionValueDetailsKindEnum, bool) {
	enum, ok := mappingDimensionValueDetailsKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
