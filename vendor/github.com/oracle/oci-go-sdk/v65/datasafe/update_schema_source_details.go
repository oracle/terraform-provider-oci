// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateSchemaSourceDetails The source of subsetting schemas
type UpdateSchemaSourceDetails interface {
}

type updateschemasourcedetails struct {
	JsonData     []byte
	SchemaSource string `json:"schemaSource"`
}

// UnmarshalJSON unmarshals json
func (m *updateschemasourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateschemasourcedetails updateschemasourcedetails
	s := struct {
		Model Unmarshalerupdateschemasourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SchemaSource = s.Model.SchemaSource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateschemasourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SchemaSource {
	case "TARGET":
		mm := UpdateSchemaSourceFromTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SENSITIVE_DATA_MODEL":
		mm := UpdateSchemaSourceFromSdmDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateSchemaSourceDetails: %s.", m.SchemaSource)
		return *m, nil
	}
}

func (m updateschemasourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateschemasourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateSchemaSourceDetailsSchemaSourceEnum Enum with underlying type: string
type UpdateSchemaSourceDetailsSchemaSourceEnum string

// Set of constants representing the allowable values for UpdateSchemaSourceDetailsSchemaSourceEnum
const (
	UpdateSchemaSourceDetailsSchemaSourceTarget             UpdateSchemaSourceDetailsSchemaSourceEnum = "TARGET"
	UpdateSchemaSourceDetailsSchemaSourceSensitiveDataModel UpdateSchemaSourceDetailsSchemaSourceEnum = "SENSITIVE_DATA_MODEL"
)

var mappingUpdateSchemaSourceDetailsSchemaSourceEnum = map[string]UpdateSchemaSourceDetailsSchemaSourceEnum{
	"TARGET":               UpdateSchemaSourceDetailsSchemaSourceTarget,
	"SENSITIVE_DATA_MODEL": UpdateSchemaSourceDetailsSchemaSourceSensitiveDataModel,
}

var mappingUpdateSchemaSourceDetailsSchemaSourceEnumLowerCase = map[string]UpdateSchemaSourceDetailsSchemaSourceEnum{
	"target":               UpdateSchemaSourceDetailsSchemaSourceTarget,
	"sensitive_data_model": UpdateSchemaSourceDetailsSchemaSourceSensitiveDataModel,
}

// GetUpdateSchemaSourceDetailsSchemaSourceEnumValues Enumerates the set of values for UpdateSchemaSourceDetailsSchemaSourceEnum
func GetUpdateSchemaSourceDetailsSchemaSourceEnumValues() []UpdateSchemaSourceDetailsSchemaSourceEnum {
	values := make([]UpdateSchemaSourceDetailsSchemaSourceEnum, 0)
	for _, v := range mappingUpdateSchemaSourceDetailsSchemaSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateSchemaSourceDetailsSchemaSourceEnumStringValues Enumerates the set of values in String for UpdateSchemaSourceDetailsSchemaSourceEnum
func GetUpdateSchemaSourceDetailsSchemaSourceEnumStringValues() []string {
	return []string{
		"TARGET",
		"SENSITIVE_DATA_MODEL",
	}
}

// GetMappingUpdateSchemaSourceDetailsSchemaSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateSchemaSourceDetailsSchemaSourceEnum(val string) (UpdateSchemaSourceDetailsSchemaSourceEnum, bool) {
	enum, ok := mappingUpdateSchemaSourceDetailsSchemaSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
