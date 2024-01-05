// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ColumnSourceDetails The source of masking columns.
type ColumnSourceDetails interface {
}

type columnsourcedetails struct {
	JsonData     []byte
	ColumnSource string `json:"columnSource"`
}

// UnmarshalJSON unmarshals json
func (m *columnsourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercolumnsourcedetails columnsourcedetails
	s := struct {
		Model Unmarshalercolumnsourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ColumnSource = s.Model.ColumnSource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *columnsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ColumnSource {
	case "SENSITIVE_DATA_MODEL":
		mm := ColumnSourceFromSdmDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TARGET":
		mm := ColumnSourceFromTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ColumnSourceDetails: %s.", m.ColumnSource)
		return *m, nil
	}
}

func (m columnsourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m columnsourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ColumnSourceDetailsColumnSourceEnum Enum with underlying type: string
type ColumnSourceDetailsColumnSourceEnum string

// Set of constants representing the allowable values for ColumnSourceDetailsColumnSourceEnum
const (
	ColumnSourceDetailsColumnSourceTarget             ColumnSourceDetailsColumnSourceEnum = "TARGET"
	ColumnSourceDetailsColumnSourceSensitiveDataModel ColumnSourceDetailsColumnSourceEnum = "SENSITIVE_DATA_MODEL"
)

var mappingColumnSourceDetailsColumnSourceEnum = map[string]ColumnSourceDetailsColumnSourceEnum{
	"TARGET":               ColumnSourceDetailsColumnSourceTarget,
	"SENSITIVE_DATA_MODEL": ColumnSourceDetailsColumnSourceSensitiveDataModel,
}

var mappingColumnSourceDetailsColumnSourceEnumLowerCase = map[string]ColumnSourceDetailsColumnSourceEnum{
	"target":               ColumnSourceDetailsColumnSourceTarget,
	"sensitive_data_model": ColumnSourceDetailsColumnSourceSensitiveDataModel,
}

// GetColumnSourceDetailsColumnSourceEnumValues Enumerates the set of values for ColumnSourceDetailsColumnSourceEnum
func GetColumnSourceDetailsColumnSourceEnumValues() []ColumnSourceDetailsColumnSourceEnum {
	values := make([]ColumnSourceDetailsColumnSourceEnum, 0)
	for _, v := range mappingColumnSourceDetailsColumnSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetColumnSourceDetailsColumnSourceEnumStringValues Enumerates the set of values in String for ColumnSourceDetailsColumnSourceEnum
func GetColumnSourceDetailsColumnSourceEnumStringValues() []string {
	return []string{
		"TARGET",
		"SENSITIVE_DATA_MODEL",
	}
}

// GetMappingColumnSourceDetailsColumnSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingColumnSourceDetailsColumnSourceEnum(val string) (ColumnSourceDetailsColumnSourceEnum, bool) {
	enum, ok := mappingColumnSourceDetailsColumnSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
