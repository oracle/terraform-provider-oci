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

// CreateColumnSourceDetails Details to associate a column source with a masking policy.
type CreateColumnSourceDetails interface {
}

type createcolumnsourcedetails struct {
	JsonData     []byte
	ColumnSource string `json:"columnSource"`
}

// UnmarshalJSON unmarshals json
func (m *createcolumnsourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatecolumnsourcedetails createcolumnsourcedetails
	s := struct {
		Model Unmarshalercreatecolumnsourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ColumnSource = s.Model.ColumnSource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createcolumnsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ColumnSource {
	case "TARGET":
		mm := CreateColumnSourceFromTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SENSITIVE_DATA_MODEL":
		mm := CreateColumnSourceFromSdmDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateColumnSourceDetails: %s.", m.ColumnSource)
		return *m, nil
	}
}

func (m createcolumnsourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createcolumnsourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateColumnSourceDetailsColumnSourceEnum Enum with underlying type: string
type CreateColumnSourceDetailsColumnSourceEnum string

// Set of constants representing the allowable values for CreateColumnSourceDetailsColumnSourceEnum
const (
	CreateColumnSourceDetailsColumnSourceTarget             CreateColumnSourceDetailsColumnSourceEnum = "TARGET"
	CreateColumnSourceDetailsColumnSourceSensitiveDataModel CreateColumnSourceDetailsColumnSourceEnum = "SENSITIVE_DATA_MODEL"
)

var mappingCreateColumnSourceDetailsColumnSourceEnum = map[string]CreateColumnSourceDetailsColumnSourceEnum{
	"TARGET":               CreateColumnSourceDetailsColumnSourceTarget,
	"SENSITIVE_DATA_MODEL": CreateColumnSourceDetailsColumnSourceSensitiveDataModel,
}

var mappingCreateColumnSourceDetailsColumnSourceEnumLowerCase = map[string]CreateColumnSourceDetailsColumnSourceEnum{
	"target":               CreateColumnSourceDetailsColumnSourceTarget,
	"sensitive_data_model": CreateColumnSourceDetailsColumnSourceSensitiveDataModel,
}

// GetCreateColumnSourceDetailsColumnSourceEnumValues Enumerates the set of values for CreateColumnSourceDetailsColumnSourceEnum
func GetCreateColumnSourceDetailsColumnSourceEnumValues() []CreateColumnSourceDetailsColumnSourceEnum {
	values := make([]CreateColumnSourceDetailsColumnSourceEnum, 0)
	for _, v := range mappingCreateColumnSourceDetailsColumnSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateColumnSourceDetailsColumnSourceEnumStringValues Enumerates the set of values in String for CreateColumnSourceDetailsColumnSourceEnum
func GetCreateColumnSourceDetailsColumnSourceEnumStringValues() []string {
	return []string{
		"TARGET",
		"SENSITIVE_DATA_MODEL",
	}
}

// GetMappingCreateColumnSourceDetailsColumnSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateColumnSourceDetailsColumnSourceEnum(val string) (CreateColumnSourceDetailsColumnSourceEnum, bool) {
	enum, ok := mappingCreateColumnSourceDetailsColumnSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
