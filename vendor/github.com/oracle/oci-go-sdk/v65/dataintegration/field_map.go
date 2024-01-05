// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FieldMap A field map is a way to map a source row shape to a target row shape that may be different.
type FieldMap interface {

	// Detailed description for the object.
	GetDescription() *string
}

type fieldmap struct {
	JsonData    []byte
	Description *string `mandatory:"false" json:"description"`
	ModelType   string  `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *fieldmap) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerfieldmap fieldmap
	s := struct {
		Model Unmarshalerfieldmap
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Description = s.Model.Description
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *fieldmap) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "RULE_BASED_FIELD_MAP":
		mm := RuleBasedFieldMap{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RULE_BASED_ENTITY_MAP":
		mm := RuleBasedEntityMap{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NAMED_ENTITY_MAP":
		mm := NamedEntityMap{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DIRECT_FIELD_MAP":
		mm := DirectFieldMap{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPOSITE_FIELD_MAP":
		mm := CompositeFieldMap{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CONDITIONAL_COMPOSITE_FIELD_MAP":
		mm := ConditionalCompositeFieldMap{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DIRECT_NAMED_FIELD_MAP":
		mm := DirectNamedFieldMap{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for FieldMap: %s.", m.ModelType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m fieldmap) GetDescription() *string {
	return m.Description
}

func (m fieldmap) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m fieldmap) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FieldMapModelTypeEnum Enum with underlying type: string
type FieldMapModelTypeEnum string

// Set of constants representing the allowable values for FieldMapModelTypeEnum
const (
	FieldMapModelTypeDirectNamedFieldMap          FieldMapModelTypeEnum = "DIRECT_NAMED_FIELD_MAP"
	FieldMapModelTypeCompositeFieldMap            FieldMapModelTypeEnum = "COMPOSITE_FIELD_MAP"
	FieldMapModelTypeDirectFieldMap               FieldMapModelTypeEnum = "DIRECT_FIELD_MAP"
	FieldMapModelTypeRuleBasedFieldMap            FieldMapModelTypeEnum = "RULE_BASED_FIELD_MAP"
	FieldMapModelTypeConditionalCompositeFieldMap FieldMapModelTypeEnum = "CONDITIONAL_COMPOSITE_FIELD_MAP"
	FieldMapModelTypeNamedEntityMap               FieldMapModelTypeEnum = "NAMED_ENTITY_MAP"
	FieldMapModelTypeRuleBasedEntityMap           FieldMapModelTypeEnum = "RULE_BASED_ENTITY_MAP"
)

var mappingFieldMapModelTypeEnum = map[string]FieldMapModelTypeEnum{
	"DIRECT_NAMED_FIELD_MAP":          FieldMapModelTypeDirectNamedFieldMap,
	"COMPOSITE_FIELD_MAP":             FieldMapModelTypeCompositeFieldMap,
	"DIRECT_FIELD_MAP":                FieldMapModelTypeDirectFieldMap,
	"RULE_BASED_FIELD_MAP":            FieldMapModelTypeRuleBasedFieldMap,
	"CONDITIONAL_COMPOSITE_FIELD_MAP": FieldMapModelTypeConditionalCompositeFieldMap,
	"NAMED_ENTITY_MAP":                FieldMapModelTypeNamedEntityMap,
	"RULE_BASED_ENTITY_MAP":           FieldMapModelTypeRuleBasedEntityMap,
}

var mappingFieldMapModelTypeEnumLowerCase = map[string]FieldMapModelTypeEnum{
	"direct_named_field_map":          FieldMapModelTypeDirectNamedFieldMap,
	"composite_field_map":             FieldMapModelTypeCompositeFieldMap,
	"direct_field_map":                FieldMapModelTypeDirectFieldMap,
	"rule_based_field_map":            FieldMapModelTypeRuleBasedFieldMap,
	"conditional_composite_field_map": FieldMapModelTypeConditionalCompositeFieldMap,
	"named_entity_map":                FieldMapModelTypeNamedEntityMap,
	"rule_based_entity_map":           FieldMapModelTypeRuleBasedEntityMap,
}

// GetFieldMapModelTypeEnumValues Enumerates the set of values for FieldMapModelTypeEnum
func GetFieldMapModelTypeEnumValues() []FieldMapModelTypeEnum {
	values := make([]FieldMapModelTypeEnum, 0)
	for _, v := range mappingFieldMapModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFieldMapModelTypeEnumStringValues Enumerates the set of values in String for FieldMapModelTypeEnum
func GetFieldMapModelTypeEnumStringValues() []string {
	return []string{
		"DIRECT_NAMED_FIELD_MAP",
		"COMPOSITE_FIELD_MAP",
		"DIRECT_FIELD_MAP",
		"RULE_BASED_FIELD_MAP",
		"CONDITIONAL_COMPOSITE_FIELD_MAP",
		"NAMED_ENTITY_MAP",
		"RULE_BASED_ENTITY_MAP",
	}
}

// GetMappingFieldMapModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFieldMapModelTypeEnum(val string) (FieldMapModelTypeEnum, bool) {
	enum, ok := mappingFieldMapModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
