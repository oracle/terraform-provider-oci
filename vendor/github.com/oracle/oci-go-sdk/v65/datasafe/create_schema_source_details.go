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

// CreateSchemaSourceDetails The source of subsetting schemas
type CreateSchemaSourceDetails interface {
}

type createschemasourcedetails struct {
	JsonData     []byte
	SchemaSource string `json:"schemaSource"`
}

// UnmarshalJSON unmarshals json
func (m *createschemasourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateschemasourcedetails createschemasourcedetails
	s := struct {
		Model Unmarshalercreateschemasourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SchemaSource = s.Model.SchemaSource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createschemasourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SchemaSource {
	case "SENSITIVE_DATA_MODEL":
		mm := CreateSchemaSourceFromSdmDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TARGET":
		mm := CreateSchemaSourceFromTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateSchemaSourceDetails: %s.", m.SchemaSource)
		return *m, nil
	}
}

func (m createschemasourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createschemasourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateSchemaSourceDetailsSchemaSourceEnum Enum with underlying type: string
type CreateSchemaSourceDetailsSchemaSourceEnum string

// Set of constants representing the allowable values for CreateSchemaSourceDetailsSchemaSourceEnum
const (
	CreateSchemaSourceDetailsSchemaSourceTarget             CreateSchemaSourceDetailsSchemaSourceEnum = "TARGET"
	CreateSchemaSourceDetailsSchemaSourceSensitiveDataModel CreateSchemaSourceDetailsSchemaSourceEnum = "SENSITIVE_DATA_MODEL"
)

var mappingCreateSchemaSourceDetailsSchemaSourceEnum = map[string]CreateSchemaSourceDetailsSchemaSourceEnum{
	"TARGET":               CreateSchemaSourceDetailsSchemaSourceTarget,
	"SENSITIVE_DATA_MODEL": CreateSchemaSourceDetailsSchemaSourceSensitiveDataModel,
}

var mappingCreateSchemaSourceDetailsSchemaSourceEnumLowerCase = map[string]CreateSchemaSourceDetailsSchemaSourceEnum{
	"target":               CreateSchemaSourceDetailsSchemaSourceTarget,
	"sensitive_data_model": CreateSchemaSourceDetailsSchemaSourceSensitiveDataModel,
}

// GetCreateSchemaSourceDetailsSchemaSourceEnumValues Enumerates the set of values for CreateSchemaSourceDetailsSchemaSourceEnum
func GetCreateSchemaSourceDetailsSchemaSourceEnumValues() []CreateSchemaSourceDetailsSchemaSourceEnum {
	values := make([]CreateSchemaSourceDetailsSchemaSourceEnum, 0)
	for _, v := range mappingCreateSchemaSourceDetailsSchemaSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateSchemaSourceDetailsSchemaSourceEnumStringValues Enumerates the set of values in String for CreateSchemaSourceDetailsSchemaSourceEnum
func GetCreateSchemaSourceDetailsSchemaSourceEnumStringValues() []string {
	return []string{
		"TARGET",
		"SENSITIVE_DATA_MODEL",
	}
}

// GetMappingCreateSchemaSourceDetailsSchemaSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateSchemaSourceDetailsSchemaSourceEnum(val string) (CreateSchemaSourceDetailsSchemaSourceEnum, bool) {
	enum, ok := mappingCreateSchemaSourceDetailsSchemaSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
