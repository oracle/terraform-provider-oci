// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RelatedObjectTypeDetails Related object details
type RelatedObjectTypeDetails interface {
}

type relatedobjecttypedetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *relatedobjecttypedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerrelatedobjecttypedetails relatedobjecttypedetails
	s := struct {
		Model Unmarshalerrelatedobjecttypedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *relatedobjecttypedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "SQL":
		mm := SqlTypeDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SCHEMA_OBJECT":
		mm := SchemaObjectTypeDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATABASE_PARAMETER":
		mm := DatabaseParameterTypeDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for RelatedObjectTypeDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m relatedobjecttypedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m relatedobjecttypedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RelatedObjectTypeDetailsTypeEnum Enum with underlying type: string
type RelatedObjectTypeDetailsTypeEnum string

// Set of constants representing the allowable values for RelatedObjectTypeDetailsTypeEnum
const (
	RelatedObjectTypeDetailsTypeSchemaObject      RelatedObjectTypeDetailsTypeEnum = "SCHEMA_OBJECT"
	RelatedObjectTypeDetailsTypeSql               RelatedObjectTypeDetailsTypeEnum = "SQL"
	RelatedObjectTypeDetailsTypeDatabaseParameter RelatedObjectTypeDetailsTypeEnum = "DATABASE_PARAMETER"
)

var mappingRelatedObjectTypeDetailsTypeEnum = map[string]RelatedObjectTypeDetailsTypeEnum{
	"SCHEMA_OBJECT":      RelatedObjectTypeDetailsTypeSchemaObject,
	"SQL":                RelatedObjectTypeDetailsTypeSql,
	"DATABASE_PARAMETER": RelatedObjectTypeDetailsTypeDatabaseParameter,
}

var mappingRelatedObjectTypeDetailsTypeEnumLowerCase = map[string]RelatedObjectTypeDetailsTypeEnum{
	"schema_object":      RelatedObjectTypeDetailsTypeSchemaObject,
	"sql":                RelatedObjectTypeDetailsTypeSql,
	"database_parameter": RelatedObjectTypeDetailsTypeDatabaseParameter,
}

// GetRelatedObjectTypeDetailsTypeEnumValues Enumerates the set of values for RelatedObjectTypeDetailsTypeEnum
func GetRelatedObjectTypeDetailsTypeEnumValues() []RelatedObjectTypeDetailsTypeEnum {
	values := make([]RelatedObjectTypeDetailsTypeEnum, 0)
	for _, v := range mappingRelatedObjectTypeDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRelatedObjectTypeDetailsTypeEnumStringValues Enumerates the set of values in String for RelatedObjectTypeDetailsTypeEnum
func GetRelatedObjectTypeDetailsTypeEnumStringValues() []string {
	return []string{
		"SCHEMA_OBJECT",
		"SQL",
		"DATABASE_PARAMETER",
	}
}

// GetMappingRelatedObjectTypeDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRelatedObjectTypeDetailsTypeEnum(val string) (RelatedObjectTypeDetailsTypeEnum, bool) {
	enum, ok := mappingRelatedObjectTypeDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
