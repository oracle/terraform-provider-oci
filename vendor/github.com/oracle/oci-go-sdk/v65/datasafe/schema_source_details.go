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

// SchemaSourceDetails The source of subsetting schemas
type SchemaSourceDetails interface {

	// The schemas to be subsetted
	GetSchemasForSubsetting() []string

	// The schemas which are related to the input list of schemas in 'schemasForSubsetting'.
	// These schemas can also be impacted from the subsetting process due to their relations with the schemas in 'schemasForSubsetting'
	GetDerivedSchemas() []string
}

type schemasourcedetails struct {
	JsonData             []byte
	SchemasForSubsetting []string `mandatory:"false" json:"schemasForSubsetting"`
	DerivedSchemas       []string `mandatory:"false" json:"derivedSchemas"`
	SchemaSource         string   `json:"schemaSource"`
}

// UnmarshalJSON unmarshals json
func (m *schemasourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerschemasourcedetails schemasourcedetails
	s := struct {
		Model Unmarshalerschemasourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SchemasForSubsetting = s.Model.SchemasForSubsetting
	m.DerivedSchemas = s.Model.DerivedSchemas
	m.SchemaSource = s.Model.SchemaSource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *schemasourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SchemaSource {
	case "TARGET":
		mm := SchemaSourceFromTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SENSITIVE_DATA_MODEL":
		mm := SchemaSourceFromSdmDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for SchemaSourceDetails: %s.", m.SchemaSource)
		return *m, nil
	}
}

// GetSchemasForSubsetting returns SchemasForSubsetting
func (m schemasourcedetails) GetSchemasForSubsetting() []string {
	return m.SchemasForSubsetting
}

// GetDerivedSchemas returns DerivedSchemas
func (m schemasourcedetails) GetDerivedSchemas() []string {
	return m.DerivedSchemas
}

func (m schemasourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m schemasourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SchemaSourceDetailsSchemaSourceEnum Enum with underlying type: string
type SchemaSourceDetailsSchemaSourceEnum string

// Set of constants representing the allowable values for SchemaSourceDetailsSchemaSourceEnum
const (
	SchemaSourceDetailsSchemaSourceTarget             SchemaSourceDetailsSchemaSourceEnum = "TARGET"
	SchemaSourceDetailsSchemaSourceSensitiveDataModel SchemaSourceDetailsSchemaSourceEnum = "SENSITIVE_DATA_MODEL"
)

var mappingSchemaSourceDetailsSchemaSourceEnum = map[string]SchemaSourceDetailsSchemaSourceEnum{
	"TARGET":               SchemaSourceDetailsSchemaSourceTarget,
	"SENSITIVE_DATA_MODEL": SchemaSourceDetailsSchemaSourceSensitiveDataModel,
}

var mappingSchemaSourceDetailsSchemaSourceEnumLowerCase = map[string]SchemaSourceDetailsSchemaSourceEnum{
	"target":               SchemaSourceDetailsSchemaSourceTarget,
	"sensitive_data_model": SchemaSourceDetailsSchemaSourceSensitiveDataModel,
}

// GetSchemaSourceDetailsSchemaSourceEnumValues Enumerates the set of values for SchemaSourceDetailsSchemaSourceEnum
func GetSchemaSourceDetailsSchemaSourceEnumValues() []SchemaSourceDetailsSchemaSourceEnum {
	values := make([]SchemaSourceDetailsSchemaSourceEnum, 0)
	for _, v := range mappingSchemaSourceDetailsSchemaSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetSchemaSourceDetailsSchemaSourceEnumStringValues Enumerates the set of values in String for SchemaSourceDetailsSchemaSourceEnum
func GetSchemaSourceDetailsSchemaSourceEnumStringValues() []string {
	return []string{
		"TARGET",
		"SENSITIVE_DATA_MODEL",
	}
}

// GetMappingSchemaSourceDetailsSchemaSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchemaSourceDetailsSchemaSourceEnum(val string) (SchemaSourceDetailsSchemaSourceEnum, bool) {
	enum, ok := mappingSchemaSourceDetailsSchemaSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
