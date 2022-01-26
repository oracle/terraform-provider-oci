// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v56/common"
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
		return *m, nil
	}
}

func (m dimensionvaluedetails) String() string {
	return common.PointerString(m)
}

// DimensionValueDetailsKindEnum Enum with underlying type: string
type DimensionValueDetailsKindEnum string

// Set of constants representing the allowable values for DimensionValueDetailsKindEnum
const (
	DimensionValueDetailsKindJmespath DimensionValueDetailsKindEnum = "jmesPath"
	DimensionValueDetailsKindStatic   DimensionValueDetailsKindEnum = "static"
)

var mappingDimensionValueDetailsKind = map[string]DimensionValueDetailsKindEnum{
	"jmesPath": DimensionValueDetailsKindJmespath,
	"static":   DimensionValueDetailsKindStatic,
}

// GetDimensionValueDetailsKindEnumValues Enumerates the set of values for DimensionValueDetailsKindEnum
func GetDimensionValueDetailsKindEnumValues() []DimensionValueDetailsKindEnum {
	values := make([]DimensionValueDetailsKindEnum, 0)
	for _, v := range mappingDimensionValueDetailsKind {
		values = append(values, v)
	}
	return values
}
