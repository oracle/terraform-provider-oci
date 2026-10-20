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

// SchemaSourceForSummary The source of subsetting schemas
type SchemaSourceForSummary interface {
}

type schemasourceforsummary struct {
	JsonData     []byte
	SchemaSource string `json:"schemaSource"`
}

// UnmarshalJSON unmarshals json
func (m *schemasourceforsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerschemasourceforsummary schemasourceforsummary
	s := struct {
		Model Unmarshalerschemasourceforsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SchemaSource = s.Model.SchemaSource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *schemasourceforsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SchemaSource {
	case "TARGET":
		mm := SchemaSourceFromTargetForSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SENSITIVE_DATA_MODEL":
		mm := SchemaSourceFromSdmForSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for SchemaSourceForSummary: %s.", m.SchemaSource)
		return *m, nil
	}
}

func (m schemasourceforsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m schemasourceforsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SchemaSourceForSummarySchemaSourceEnum Enum with underlying type: string
type SchemaSourceForSummarySchemaSourceEnum string

// Set of constants representing the allowable values for SchemaSourceForSummarySchemaSourceEnum
const (
	SchemaSourceForSummarySchemaSourceTarget             SchemaSourceForSummarySchemaSourceEnum = "TARGET"
	SchemaSourceForSummarySchemaSourceSensitiveDataModel SchemaSourceForSummarySchemaSourceEnum = "SENSITIVE_DATA_MODEL"
)

var mappingSchemaSourceForSummarySchemaSourceEnum = map[string]SchemaSourceForSummarySchemaSourceEnum{
	"TARGET":               SchemaSourceForSummarySchemaSourceTarget,
	"SENSITIVE_DATA_MODEL": SchemaSourceForSummarySchemaSourceSensitiveDataModel,
}

var mappingSchemaSourceForSummarySchemaSourceEnumLowerCase = map[string]SchemaSourceForSummarySchemaSourceEnum{
	"target":               SchemaSourceForSummarySchemaSourceTarget,
	"sensitive_data_model": SchemaSourceForSummarySchemaSourceSensitiveDataModel,
}

// GetSchemaSourceForSummarySchemaSourceEnumValues Enumerates the set of values for SchemaSourceForSummarySchemaSourceEnum
func GetSchemaSourceForSummarySchemaSourceEnumValues() []SchemaSourceForSummarySchemaSourceEnum {
	values := make([]SchemaSourceForSummarySchemaSourceEnum, 0)
	for _, v := range mappingSchemaSourceForSummarySchemaSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetSchemaSourceForSummarySchemaSourceEnumStringValues Enumerates the set of values in String for SchemaSourceForSummarySchemaSourceEnum
func GetSchemaSourceForSummarySchemaSourceEnumStringValues() []string {
	return []string{
		"TARGET",
		"SENSITIVE_DATA_MODEL",
	}
}

// GetMappingSchemaSourceForSummarySchemaSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchemaSourceForSummarySchemaSourceEnum(val string) (SchemaSourceForSummarySchemaSourceEnum, bool) {
	enum, ok := mappingSchemaSourceForSummarySchemaSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
