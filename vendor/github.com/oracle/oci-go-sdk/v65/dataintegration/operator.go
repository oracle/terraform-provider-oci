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

// Operator An operator defines some data integration semantics in a data flow. It may be reading/writing data or transforming the data.
type Operator interface {

	// The key of the object.
	GetKey() *string

	// The model version of an object.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	GetName() *string

	// Details about the operator.
	GetDescription() *string

	// The version of the object that is used to track changes in the object instance.
	GetObjectVersion() *int

	// An array of input ports.
	GetInputPorts() []InputPort

	// An array of output ports.
	GetOutputPorts() []TypedObject

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	GetIdentifier() *string

	// An array of parameters used in the data flow.
	GetParameters() []Parameter

	GetOpConfigValues() *ConfigValues
}

type operator struct {
	JsonData       []byte
	Key            *string          `mandatory:"false" json:"key"`
	ModelVersion   *string          `mandatory:"false" json:"modelVersion"`
	ParentRef      *ParentReference `mandatory:"false" json:"parentRef"`
	Name           *string          `mandatory:"false" json:"name"`
	Description    *string          `mandatory:"false" json:"description"`
	ObjectVersion  *int             `mandatory:"false" json:"objectVersion"`
	InputPorts     []InputPort      `mandatory:"false" json:"inputPorts"`
	OutputPorts    json.RawMessage  `mandatory:"false" json:"outputPorts"`
	ObjectStatus   *int             `mandatory:"false" json:"objectStatus"`
	Identifier     *string          `mandatory:"false" json:"identifier"`
	Parameters     []Parameter      `mandatory:"false" json:"parameters"`
	OpConfigValues *ConfigValues    `mandatory:"false" json:"opConfigValues"`
	ModelType      string           `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *operator) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleroperator operator
	s := struct {
		Model Unmarshaleroperator
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.ModelVersion = s.Model.ModelVersion
	m.ParentRef = s.Model.ParentRef
	m.Name = s.Model.Name
	m.Description = s.Model.Description
	m.ObjectVersion = s.Model.ObjectVersion
	m.InputPorts = s.Model.InputPorts
	m.OutputPorts = s.Model.OutputPorts
	m.ObjectStatus = s.Model.ObjectStatus
	m.Identifier = s.Model.Identifier
	m.Parameters = s.Model.Parameters
	m.OpConfigValues = s.Model.OpConfigValues
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *operator) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "JOINER_OPERATOR":
		mm := Joiner{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TASK_OPERATOR":
		mm := TaskOperator{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FLATTEN_OPERATOR":
		mm := Flatten{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AGGREGATOR_OPERATOR":
		mm := Aggregator{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SORT_OPERATOR":
		mm := SortOper{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PROJECTION_OPERATOR":
		mm := Projection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "END_OPERATOR":
		mm := EndOperator{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SOURCE_OPERATOR":
		mm := Source{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "UNION_OPERATOR":
		mm := Union{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXPRESSION_OPERATOR":
		mm := ExpressionOperator{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FUNCTION_OPERATOR":
		mm := Function{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DECISION_OPERATOR":
		mm := DecisionOperator{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INTERSECT_OPERATOR":
		mm := Intersect{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TARGET_OPERATOR":
		mm := Target{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DISTINCT_OPERATOR":
		mm := Distinct{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FILTER_OPERATOR":
		mm := Filter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOOKUP_OPERATOR":
		mm := Lookup{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PIVOT_OPERATOR":
		mm := Pivot{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "START_OPERATOR":
		mm := StartOperator{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MERGE_OPERATOR":
		mm := MergeOperator{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SPLIT_OPERATOR":
		mm := Split{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MINUS_OPERATOR":
		mm := Minus{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Operator: %s.", m.ModelType)
		return *m, nil
	}
}

// GetKey returns Key
func (m operator) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m operator) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m operator) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m operator) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m operator) GetDescription() *string {
	return m.Description
}

// GetObjectVersion returns ObjectVersion
func (m operator) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetInputPorts returns InputPorts
func (m operator) GetInputPorts() []InputPort {
	return m.InputPorts
}

// GetOutputPorts returns OutputPorts
func (m operator) GetOutputPorts() json.RawMessage {
	return m.OutputPorts
}

// GetObjectStatus returns ObjectStatus
func (m operator) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m operator) GetIdentifier() *string {
	return m.Identifier
}

// GetParameters returns Parameters
func (m operator) GetParameters() []Parameter {
	return m.Parameters
}

// GetOpConfigValues returns OpConfigValues
func (m operator) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

func (m operator) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m operator) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OperatorModelTypeEnum Enum with underlying type: string
type OperatorModelTypeEnum string

// Set of constants representing the allowable values for OperatorModelTypeEnum
const (
	OperatorModelTypeSourceOperator     OperatorModelTypeEnum = "SOURCE_OPERATOR"
	OperatorModelTypeFilterOperator     OperatorModelTypeEnum = "FILTER_OPERATOR"
	OperatorModelTypeJoinerOperator     OperatorModelTypeEnum = "JOINER_OPERATOR"
	OperatorModelTypeAggregatorOperator OperatorModelTypeEnum = "AGGREGATOR_OPERATOR"
	OperatorModelTypeProjectionOperator OperatorModelTypeEnum = "PROJECTION_OPERATOR"
	OperatorModelTypeTargetOperator     OperatorModelTypeEnum = "TARGET_OPERATOR"
	OperatorModelTypeFlattenOperator    OperatorModelTypeEnum = "FLATTEN_OPERATOR"
	OperatorModelTypeDistinctOperator   OperatorModelTypeEnum = "DISTINCT_OPERATOR"
	OperatorModelTypeSortOperator       OperatorModelTypeEnum = "SORT_OPERATOR"
	OperatorModelTypeUnionOperator      OperatorModelTypeEnum = "UNION_OPERATOR"
	OperatorModelTypeIntersectOperator  OperatorModelTypeEnum = "INTERSECT_OPERATOR"
	OperatorModelTypeMinusOperator      OperatorModelTypeEnum = "MINUS_OPERATOR"
	OperatorModelTypeMergeOperator      OperatorModelTypeEnum = "MERGE_OPERATOR"
	OperatorModelTypeFunctionOperator   OperatorModelTypeEnum = "FUNCTION_OPERATOR"
	OperatorModelTypeSplitOperator      OperatorModelTypeEnum = "SPLIT_OPERATOR"
	OperatorModelTypeStartOperator      OperatorModelTypeEnum = "START_OPERATOR"
	OperatorModelTypeEndOperator        OperatorModelTypeEnum = "END_OPERATOR"
	OperatorModelTypePipelineOperator   OperatorModelTypeEnum = "PIPELINE_OPERATOR"
	OperatorModelTypeDecisionOperator   OperatorModelTypeEnum = "DECISION_OPERATOR"
	OperatorModelTypeTaskOperator       OperatorModelTypeEnum = "TASK_OPERATOR"
	OperatorModelTypeExpressionOperator OperatorModelTypeEnum = "EXPRESSION_OPERATOR"
	OperatorModelTypeLookupOperator     OperatorModelTypeEnum = "LOOKUP_OPERATOR"
	OperatorModelTypePivotOperator      OperatorModelTypeEnum = "PIVOT_OPERATOR"
)

var mappingOperatorModelTypeEnum = map[string]OperatorModelTypeEnum{
	"SOURCE_OPERATOR":     OperatorModelTypeSourceOperator,
	"FILTER_OPERATOR":     OperatorModelTypeFilterOperator,
	"JOINER_OPERATOR":     OperatorModelTypeJoinerOperator,
	"AGGREGATOR_OPERATOR": OperatorModelTypeAggregatorOperator,
	"PROJECTION_OPERATOR": OperatorModelTypeProjectionOperator,
	"TARGET_OPERATOR":     OperatorModelTypeTargetOperator,
	"FLATTEN_OPERATOR":    OperatorModelTypeFlattenOperator,
	"DISTINCT_OPERATOR":   OperatorModelTypeDistinctOperator,
	"SORT_OPERATOR":       OperatorModelTypeSortOperator,
	"UNION_OPERATOR":      OperatorModelTypeUnionOperator,
	"INTERSECT_OPERATOR":  OperatorModelTypeIntersectOperator,
	"MINUS_OPERATOR":      OperatorModelTypeMinusOperator,
	"MERGE_OPERATOR":      OperatorModelTypeMergeOperator,
	"FUNCTION_OPERATOR":   OperatorModelTypeFunctionOperator,
	"SPLIT_OPERATOR":      OperatorModelTypeSplitOperator,
	"START_OPERATOR":      OperatorModelTypeStartOperator,
	"END_OPERATOR":        OperatorModelTypeEndOperator,
	"PIPELINE_OPERATOR":   OperatorModelTypePipelineOperator,
	"DECISION_OPERATOR":   OperatorModelTypeDecisionOperator,
	"TASK_OPERATOR":       OperatorModelTypeTaskOperator,
	"EXPRESSION_OPERATOR": OperatorModelTypeExpressionOperator,
	"LOOKUP_OPERATOR":     OperatorModelTypeLookupOperator,
	"PIVOT_OPERATOR":      OperatorModelTypePivotOperator,
}

var mappingOperatorModelTypeEnumLowerCase = map[string]OperatorModelTypeEnum{
	"source_operator":     OperatorModelTypeSourceOperator,
	"filter_operator":     OperatorModelTypeFilterOperator,
	"joiner_operator":     OperatorModelTypeJoinerOperator,
	"aggregator_operator": OperatorModelTypeAggregatorOperator,
	"projection_operator": OperatorModelTypeProjectionOperator,
	"target_operator":     OperatorModelTypeTargetOperator,
	"flatten_operator":    OperatorModelTypeFlattenOperator,
	"distinct_operator":   OperatorModelTypeDistinctOperator,
	"sort_operator":       OperatorModelTypeSortOperator,
	"union_operator":      OperatorModelTypeUnionOperator,
	"intersect_operator":  OperatorModelTypeIntersectOperator,
	"minus_operator":      OperatorModelTypeMinusOperator,
	"merge_operator":      OperatorModelTypeMergeOperator,
	"function_operator":   OperatorModelTypeFunctionOperator,
	"split_operator":      OperatorModelTypeSplitOperator,
	"start_operator":      OperatorModelTypeStartOperator,
	"end_operator":        OperatorModelTypeEndOperator,
	"pipeline_operator":   OperatorModelTypePipelineOperator,
	"decision_operator":   OperatorModelTypeDecisionOperator,
	"task_operator":       OperatorModelTypeTaskOperator,
	"expression_operator": OperatorModelTypeExpressionOperator,
	"lookup_operator":     OperatorModelTypeLookupOperator,
	"pivot_operator":      OperatorModelTypePivotOperator,
}

// GetOperatorModelTypeEnumValues Enumerates the set of values for OperatorModelTypeEnum
func GetOperatorModelTypeEnumValues() []OperatorModelTypeEnum {
	values := make([]OperatorModelTypeEnum, 0)
	for _, v := range mappingOperatorModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperatorModelTypeEnumStringValues Enumerates the set of values in String for OperatorModelTypeEnum
func GetOperatorModelTypeEnumStringValues() []string {
	return []string{
		"SOURCE_OPERATOR",
		"FILTER_OPERATOR",
		"JOINER_OPERATOR",
		"AGGREGATOR_OPERATOR",
		"PROJECTION_OPERATOR",
		"TARGET_OPERATOR",
		"FLATTEN_OPERATOR",
		"DISTINCT_OPERATOR",
		"SORT_OPERATOR",
		"UNION_OPERATOR",
		"INTERSECT_OPERATOR",
		"MINUS_OPERATOR",
		"MERGE_OPERATOR",
		"FUNCTION_OPERATOR",
		"SPLIT_OPERATOR",
		"START_OPERATOR",
		"END_OPERATOR",
		"PIPELINE_OPERATOR",
		"DECISION_OPERATOR",
		"TASK_OPERATOR",
		"EXPRESSION_OPERATOR",
		"LOOKUP_OPERATOR",
		"PIVOT_OPERATOR",
	}
}

// GetMappingOperatorModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperatorModelTypeEnum(val string) (OperatorModelTypeEnum, bool) {
	enum, ok := mappingOperatorModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
