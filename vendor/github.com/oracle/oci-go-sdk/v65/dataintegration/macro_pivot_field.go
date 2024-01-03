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

// MacroPivotField MacroPivotField is used for the PivotField with macro expressions. It can contain the rules according to the macro pattern/attribute added and create new fields according to the PivotKeyValues
type MacroPivotField struct {

	// The type of the types object.
	ModelType MacroPivotFieldModelTypeEnum `mandatory:"true" json:"modelType"`

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	Expr *Expression `mandatory:"false" json:"expr"`

	UseType *ConfiguredType `mandatory:"false" json:"useType"`

	Type BaseType `mandatory:"false" json:"type"`

	// column name pattern can be used to generate the name structure of the generated columns. By default column names are of %PIVOT_KEY_VALUE% or %MACRO_INPUT%_%PIVOT_KEY_VALUE%, but we can change it something by passing something like MY_PREFIX%PIVOT_KEY_VALUE%MY_SUFFIX or MY_PREFIX%MACRO_INPUT%_%PIVOT_KEY_VALUE%MY_SUFFIX which will add custom prefix and suffix to the column name.
	ColumnNamePattern *string `mandatory:"false" json:"columnNamePattern"`

	// Specifies whether the type of macro fields is inferred from an expression or useType (false) or the source field (true).
	IsUseSourceType *bool `mandatory:"false" json:"isUseSourceType"`
}

func (m MacroPivotField) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MacroPivotField) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMacroPivotFieldModelTypeEnum(string(m.ModelType)); !ok && m.ModelType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelType: %s. Supported values are: %s.", m.ModelType, strings.Join(GetMacroPivotFieldModelTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *MacroPivotField) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key               *string                      `json:"key"`
		ModelVersion      *string                      `json:"modelVersion"`
		ParentRef         *ParentReference             `json:"parentRef"`
		ConfigValues      *ConfigValues                `json:"configValues"`
		ObjectStatus      *int                         `json:"objectStatus"`
		Name              *string                      `json:"name"`
		Description       *string                      `json:"description"`
		Expr              *Expression                  `json:"expr"`
		UseType           *ConfiguredType              `json:"useType"`
		Type              basetype                     `json:"type"`
		ColumnNamePattern *string                      `json:"columnNamePattern"`
		IsUseSourceType   *bool                        `json:"isUseSourceType"`
		ModelType         MacroPivotFieldModelTypeEnum `json:"modelType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.ConfigValues = model.ConfigValues

	m.ObjectStatus = model.ObjectStatus

	m.Name = model.Name

	m.Description = model.Description

	m.Expr = model.Expr

	m.UseType = model.UseType

	nn, e = model.Type.UnmarshalPolymorphicJSON(model.Type.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Type = nn.(BaseType)
	} else {
		m.Type = nil
	}

	m.ColumnNamePattern = model.ColumnNamePattern

	m.IsUseSourceType = model.IsUseSourceType

	m.ModelType = model.ModelType

	return
}

// MacroPivotFieldModelTypeEnum Enum with underlying type: string
type MacroPivotFieldModelTypeEnum string

// Set of constants representing the allowable values for MacroPivotFieldModelTypeEnum
const (
	MacroPivotFieldModelTypeShape                    MacroPivotFieldModelTypeEnum = "SHAPE"
	MacroPivotFieldModelTypeInputPort                MacroPivotFieldModelTypeEnum = "INPUT_PORT"
	MacroPivotFieldModelTypeShapeField               MacroPivotFieldModelTypeEnum = "SHAPE_FIELD"
	MacroPivotFieldModelTypeInputField               MacroPivotFieldModelTypeEnum = "INPUT_FIELD"
	MacroPivotFieldModelTypeDerivedField             MacroPivotFieldModelTypeEnum = "DERIVED_FIELD"
	MacroPivotFieldModelTypeMacroField               MacroPivotFieldModelTypeEnum = "MACRO_FIELD"
	MacroPivotFieldModelTypeOutputField              MacroPivotFieldModelTypeEnum = "OUTPUT_FIELD"
	MacroPivotFieldModelTypeDynamicProxyField        MacroPivotFieldModelTypeEnum = "DYNAMIC_PROXY_FIELD"
	MacroPivotFieldModelTypeOutputPort               MacroPivotFieldModelTypeEnum = "OUTPUT_PORT"
	MacroPivotFieldModelTypeDynamicInputField        MacroPivotFieldModelTypeEnum = "DYNAMIC_INPUT_FIELD"
	MacroPivotFieldModelTypeProxyField               MacroPivotFieldModelTypeEnum = "PROXY_FIELD"
	MacroPivotFieldModelTypeParameter                MacroPivotFieldModelTypeEnum = "PARAMETER"
	MacroPivotFieldModelTypePivotField               MacroPivotFieldModelTypeEnum = "PIVOT_FIELD"
	MacroPivotFieldModelTypeMacroPivotField          MacroPivotFieldModelTypeEnum = "MACRO_PIVOT_FIELD"
	MacroPivotFieldModelTypeConditionalOutputPort    MacroPivotFieldModelTypeEnum = "CONDITIONAL_OUTPUT_PORT"
	MacroPivotFieldModelTypeInputProxyField          MacroPivotFieldModelTypeEnum = "INPUT_PROXY_FIELD"
	MacroPivotFieldModelTypeMaterializedDynamicField MacroPivotFieldModelTypeEnum = "MATERIALIZED_DYNAMIC_FIELD"
	MacroPivotFieldModelTypeDecisionOutputPort       MacroPivotFieldModelTypeEnum = "DECISION_OUTPUT_PORT"
)

var mappingMacroPivotFieldModelTypeEnum = map[string]MacroPivotFieldModelTypeEnum{
	"SHAPE":                      MacroPivotFieldModelTypeShape,
	"INPUT_PORT":                 MacroPivotFieldModelTypeInputPort,
	"SHAPE_FIELD":                MacroPivotFieldModelTypeShapeField,
	"INPUT_FIELD":                MacroPivotFieldModelTypeInputField,
	"DERIVED_FIELD":              MacroPivotFieldModelTypeDerivedField,
	"MACRO_FIELD":                MacroPivotFieldModelTypeMacroField,
	"OUTPUT_FIELD":               MacroPivotFieldModelTypeOutputField,
	"DYNAMIC_PROXY_FIELD":        MacroPivotFieldModelTypeDynamicProxyField,
	"OUTPUT_PORT":                MacroPivotFieldModelTypeOutputPort,
	"DYNAMIC_INPUT_FIELD":        MacroPivotFieldModelTypeDynamicInputField,
	"PROXY_FIELD":                MacroPivotFieldModelTypeProxyField,
	"PARAMETER":                  MacroPivotFieldModelTypeParameter,
	"PIVOT_FIELD":                MacroPivotFieldModelTypePivotField,
	"MACRO_PIVOT_FIELD":          MacroPivotFieldModelTypeMacroPivotField,
	"CONDITIONAL_OUTPUT_PORT":    MacroPivotFieldModelTypeConditionalOutputPort,
	"INPUT_PROXY_FIELD":          MacroPivotFieldModelTypeInputProxyField,
	"MATERIALIZED_DYNAMIC_FIELD": MacroPivotFieldModelTypeMaterializedDynamicField,
	"DECISION_OUTPUT_PORT":       MacroPivotFieldModelTypeDecisionOutputPort,
}

var mappingMacroPivotFieldModelTypeEnumLowerCase = map[string]MacroPivotFieldModelTypeEnum{
	"shape":                      MacroPivotFieldModelTypeShape,
	"input_port":                 MacroPivotFieldModelTypeInputPort,
	"shape_field":                MacroPivotFieldModelTypeShapeField,
	"input_field":                MacroPivotFieldModelTypeInputField,
	"derived_field":              MacroPivotFieldModelTypeDerivedField,
	"macro_field":                MacroPivotFieldModelTypeMacroField,
	"output_field":               MacroPivotFieldModelTypeOutputField,
	"dynamic_proxy_field":        MacroPivotFieldModelTypeDynamicProxyField,
	"output_port":                MacroPivotFieldModelTypeOutputPort,
	"dynamic_input_field":        MacroPivotFieldModelTypeDynamicInputField,
	"proxy_field":                MacroPivotFieldModelTypeProxyField,
	"parameter":                  MacroPivotFieldModelTypeParameter,
	"pivot_field":                MacroPivotFieldModelTypePivotField,
	"macro_pivot_field":          MacroPivotFieldModelTypeMacroPivotField,
	"conditional_output_port":    MacroPivotFieldModelTypeConditionalOutputPort,
	"input_proxy_field":          MacroPivotFieldModelTypeInputProxyField,
	"materialized_dynamic_field": MacroPivotFieldModelTypeMaterializedDynamicField,
	"decision_output_port":       MacroPivotFieldModelTypeDecisionOutputPort,
}

// GetMacroPivotFieldModelTypeEnumValues Enumerates the set of values for MacroPivotFieldModelTypeEnum
func GetMacroPivotFieldModelTypeEnumValues() []MacroPivotFieldModelTypeEnum {
	values := make([]MacroPivotFieldModelTypeEnum, 0)
	for _, v := range mappingMacroPivotFieldModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMacroPivotFieldModelTypeEnumStringValues Enumerates the set of values in String for MacroPivotFieldModelTypeEnum
func GetMacroPivotFieldModelTypeEnumStringValues() []string {
	return []string{
		"SHAPE",
		"INPUT_PORT",
		"SHAPE_FIELD",
		"INPUT_FIELD",
		"DERIVED_FIELD",
		"MACRO_FIELD",
		"OUTPUT_FIELD",
		"DYNAMIC_PROXY_FIELD",
		"OUTPUT_PORT",
		"DYNAMIC_INPUT_FIELD",
		"PROXY_FIELD",
		"PARAMETER",
		"PIVOT_FIELD",
		"MACRO_PIVOT_FIELD",
		"CONDITIONAL_OUTPUT_PORT",
		"INPUT_PROXY_FIELD",
		"MATERIALIZED_DYNAMIC_FIELD",
		"DECISION_OUTPUT_PORT",
	}
}

// GetMappingMacroPivotFieldModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMacroPivotFieldModelTypeEnum(val string) (MacroPivotFieldModelTypeEnum, bool) {
	enum, ok := mappingMacroPivotFieldModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
