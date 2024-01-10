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

// TypedObject The `TypedObject` class is a base class for any model object that has a type.
type TypedObject interface {

	// The key of the object.
	GetKey() *string

	// The model version of an object.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	GetConfigValues() *ConfigValues

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	GetName() *string

	// Detailed description for the object.
	GetDescription() *string
}

type typedobject struct {
	JsonData     []byte
	Key          *string          `mandatory:"false" json:"key"`
	ModelVersion *string          `mandatory:"false" json:"modelVersion"`
	ParentRef    *ParentReference `mandatory:"false" json:"parentRef"`
	ConfigValues *ConfigValues    `mandatory:"false" json:"configValues"`
	ObjectStatus *int             `mandatory:"false" json:"objectStatus"`
	Name         *string          `mandatory:"false" json:"name"`
	Description  *string          `mandatory:"false" json:"description"`
	ModelType    string           `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *typedobject) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertypedobject typedobject
	s := struct {
		Model Unmarshalertypedobject
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.ModelVersion = s.Model.ModelVersion
	m.ParentRef = s.Model.ParentRef
	m.ConfigValues = s.Model.ConfigValues
	m.ObjectStatus = s.Model.ObjectStatus
	m.Name = s.Model.Name
	m.Description = s.Model.Description
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *typedobject) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "OUTPUT_PORT":
		mm := OutputPort{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DYNAMIC_INPUT_FIELD":
		mm := DynamicInputField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FIELD":
		mm := AbstractField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INPUT_FIELD":
		mm := InputField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SHAPE":
		mm := Shape{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INPUT_PORT":
		mm := InputPort{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CONDITIONAL_OUTPUT_PORT":
		mm := ConditionalOutputPort{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DECISION_OUTPUT_PORT":
		mm := DecisionOutputPort{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MATERIALIZED_DYNAMIC_FIELD":
		mm := MaterializedDynamicField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PROXY_FIELD":
		mm := ProxyField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DYNAMIC_PROXY_FIELD":
		mm := DynamicProxyField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SHAPE_FIELD":
		mm := ShapeField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INPUT_PROXY_FIELD":
		mm := InputProxyField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PARAMETER":
		mm := Parameter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PIVOT_FIELD":
		mm := PivotField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OUTPUT_FIELD":
		mm := OutputField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MACRO_FIELD":
		mm := MacroField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DERIVED_FIELD":
		mm := DerivedField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TYPED_EXPRESSION":
		mm := TypedExpression{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FLOW_PORT":
		mm := FlowPort{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for TypedObject: %s.", m.ModelType)
		return *m, nil
	}
}

// GetKey returns Key
func (m typedobject) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m typedobject) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m typedobject) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetConfigValues returns ConfigValues
func (m typedobject) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

// GetObjectStatus returns ObjectStatus
func (m typedobject) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetName returns Name
func (m typedobject) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m typedobject) GetDescription() *string {
	return m.Description
}

func (m typedobject) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m typedobject) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TypedObjectModelTypeEnum Enum with underlying type: string
type TypedObjectModelTypeEnum string

// Set of constants representing the allowable values for TypedObjectModelTypeEnum
const (
	TypedObjectModelTypeShape                    TypedObjectModelTypeEnum = "SHAPE"
	TypedObjectModelTypeInputPort                TypedObjectModelTypeEnum = "INPUT_PORT"
	TypedObjectModelTypeShapeField               TypedObjectModelTypeEnum = "SHAPE_FIELD"
	TypedObjectModelTypeInputField               TypedObjectModelTypeEnum = "INPUT_FIELD"
	TypedObjectModelTypeDerivedField             TypedObjectModelTypeEnum = "DERIVED_FIELD"
	TypedObjectModelTypeMacroField               TypedObjectModelTypeEnum = "MACRO_FIELD"
	TypedObjectModelTypeOutputField              TypedObjectModelTypeEnum = "OUTPUT_FIELD"
	TypedObjectModelTypeDynamicProxyField        TypedObjectModelTypeEnum = "DYNAMIC_PROXY_FIELD"
	TypedObjectModelTypeOutputPort               TypedObjectModelTypeEnum = "OUTPUT_PORT"
	TypedObjectModelTypeDynamicInputField        TypedObjectModelTypeEnum = "DYNAMIC_INPUT_FIELD"
	TypedObjectModelTypeProxyField               TypedObjectModelTypeEnum = "PROXY_FIELD"
	TypedObjectModelTypeParameter                TypedObjectModelTypeEnum = "PARAMETER"
	TypedObjectModelTypePivotField               TypedObjectModelTypeEnum = "PIVOT_FIELD"
	TypedObjectModelTypeMacroPivotField          TypedObjectModelTypeEnum = "MACRO_PIVOT_FIELD"
	TypedObjectModelTypeConditionalOutputPort    TypedObjectModelTypeEnum = "CONDITIONAL_OUTPUT_PORT"
	TypedObjectModelTypeInputProxyField          TypedObjectModelTypeEnum = "INPUT_PROXY_FIELD"
	TypedObjectModelTypeMaterializedDynamicField TypedObjectModelTypeEnum = "MATERIALIZED_DYNAMIC_FIELD"
	TypedObjectModelTypeDecisionOutputPort       TypedObjectModelTypeEnum = "DECISION_OUTPUT_PORT"
)

var mappingTypedObjectModelTypeEnum = map[string]TypedObjectModelTypeEnum{
	"SHAPE":                      TypedObjectModelTypeShape,
	"INPUT_PORT":                 TypedObjectModelTypeInputPort,
	"SHAPE_FIELD":                TypedObjectModelTypeShapeField,
	"INPUT_FIELD":                TypedObjectModelTypeInputField,
	"DERIVED_FIELD":              TypedObjectModelTypeDerivedField,
	"MACRO_FIELD":                TypedObjectModelTypeMacroField,
	"OUTPUT_FIELD":               TypedObjectModelTypeOutputField,
	"DYNAMIC_PROXY_FIELD":        TypedObjectModelTypeDynamicProxyField,
	"OUTPUT_PORT":                TypedObjectModelTypeOutputPort,
	"DYNAMIC_INPUT_FIELD":        TypedObjectModelTypeDynamicInputField,
	"PROXY_FIELD":                TypedObjectModelTypeProxyField,
	"PARAMETER":                  TypedObjectModelTypeParameter,
	"PIVOT_FIELD":                TypedObjectModelTypePivotField,
	"MACRO_PIVOT_FIELD":          TypedObjectModelTypeMacroPivotField,
	"CONDITIONAL_OUTPUT_PORT":    TypedObjectModelTypeConditionalOutputPort,
	"INPUT_PROXY_FIELD":          TypedObjectModelTypeInputProxyField,
	"MATERIALIZED_DYNAMIC_FIELD": TypedObjectModelTypeMaterializedDynamicField,
	"DECISION_OUTPUT_PORT":       TypedObjectModelTypeDecisionOutputPort,
}

var mappingTypedObjectModelTypeEnumLowerCase = map[string]TypedObjectModelTypeEnum{
	"shape":                      TypedObjectModelTypeShape,
	"input_port":                 TypedObjectModelTypeInputPort,
	"shape_field":                TypedObjectModelTypeShapeField,
	"input_field":                TypedObjectModelTypeInputField,
	"derived_field":              TypedObjectModelTypeDerivedField,
	"macro_field":                TypedObjectModelTypeMacroField,
	"output_field":               TypedObjectModelTypeOutputField,
	"dynamic_proxy_field":        TypedObjectModelTypeDynamicProxyField,
	"output_port":                TypedObjectModelTypeOutputPort,
	"dynamic_input_field":        TypedObjectModelTypeDynamicInputField,
	"proxy_field":                TypedObjectModelTypeProxyField,
	"parameter":                  TypedObjectModelTypeParameter,
	"pivot_field":                TypedObjectModelTypePivotField,
	"macro_pivot_field":          TypedObjectModelTypeMacroPivotField,
	"conditional_output_port":    TypedObjectModelTypeConditionalOutputPort,
	"input_proxy_field":          TypedObjectModelTypeInputProxyField,
	"materialized_dynamic_field": TypedObjectModelTypeMaterializedDynamicField,
	"decision_output_port":       TypedObjectModelTypeDecisionOutputPort,
}

// GetTypedObjectModelTypeEnumValues Enumerates the set of values for TypedObjectModelTypeEnum
func GetTypedObjectModelTypeEnumValues() []TypedObjectModelTypeEnum {
	values := make([]TypedObjectModelTypeEnum, 0)
	for _, v := range mappingTypedObjectModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTypedObjectModelTypeEnumStringValues Enumerates the set of values in String for TypedObjectModelTypeEnum
func GetTypedObjectModelTypeEnumStringValues() []string {
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

// GetMappingTypedObjectModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTypedObjectModelTypeEnum(val string) (TypedObjectModelTypeEnum, bool) {
	enum, ok := mappingTypedObjectModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
