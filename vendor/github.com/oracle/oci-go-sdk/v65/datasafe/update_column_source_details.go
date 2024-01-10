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

// UpdateColumnSourceDetails Details to update the column source of a masking policy.
type UpdateColumnSourceDetails interface {
}

type updatecolumnsourcedetails struct {
	JsonData     []byte
	ColumnSource string `json:"columnSource"`
}

// UnmarshalJSON unmarshals json
func (m *updatecolumnsourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatecolumnsourcedetails updatecolumnsourcedetails
	s := struct {
		Model Unmarshalerupdatecolumnsourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ColumnSource = s.Model.ColumnSource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatecolumnsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ColumnSource {
	case "SENSITIVE_DATA_MODEL":
		mm := UpdateColumnSourceSdmDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TARGET":
		mm := UpdateColumnSourceTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdateColumnSourceDetails: %s.", m.ColumnSource)
		return *m, nil
	}
}

func (m updatecolumnsourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatecolumnsourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateColumnSourceDetailsColumnSourceEnum Enum with underlying type: string
type UpdateColumnSourceDetailsColumnSourceEnum string

// Set of constants representing the allowable values for UpdateColumnSourceDetailsColumnSourceEnum
const (
	UpdateColumnSourceDetailsColumnSourceTarget             UpdateColumnSourceDetailsColumnSourceEnum = "TARGET"
	UpdateColumnSourceDetailsColumnSourceSensitiveDataModel UpdateColumnSourceDetailsColumnSourceEnum = "SENSITIVE_DATA_MODEL"
)

var mappingUpdateColumnSourceDetailsColumnSourceEnum = map[string]UpdateColumnSourceDetailsColumnSourceEnum{
	"TARGET":               UpdateColumnSourceDetailsColumnSourceTarget,
	"SENSITIVE_DATA_MODEL": UpdateColumnSourceDetailsColumnSourceSensitiveDataModel,
}

var mappingUpdateColumnSourceDetailsColumnSourceEnumLowerCase = map[string]UpdateColumnSourceDetailsColumnSourceEnum{
	"target":               UpdateColumnSourceDetailsColumnSourceTarget,
	"sensitive_data_model": UpdateColumnSourceDetailsColumnSourceSensitiveDataModel,
}

// GetUpdateColumnSourceDetailsColumnSourceEnumValues Enumerates the set of values for UpdateColumnSourceDetailsColumnSourceEnum
func GetUpdateColumnSourceDetailsColumnSourceEnumValues() []UpdateColumnSourceDetailsColumnSourceEnum {
	values := make([]UpdateColumnSourceDetailsColumnSourceEnum, 0)
	for _, v := range mappingUpdateColumnSourceDetailsColumnSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateColumnSourceDetailsColumnSourceEnumStringValues Enumerates the set of values in String for UpdateColumnSourceDetailsColumnSourceEnum
func GetUpdateColumnSourceDetailsColumnSourceEnumStringValues() []string {
	return []string{
		"TARGET",
		"SENSITIVE_DATA_MODEL",
	}
}

// GetMappingUpdateColumnSourceDetailsColumnSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateColumnSourceDetailsColumnSourceEnum(val string) (UpdateColumnSourceDetailsColumnSourceEnum, bool) {
	enum, ok := mappingUpdateColumnSourceDetailsColumnSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
