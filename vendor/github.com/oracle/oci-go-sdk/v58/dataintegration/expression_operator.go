// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ExpressionOperator An operator for expressions
type ExpressionOperator struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Details about the operator.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// An array of input ports.
	InputPorts []InputPort `mandatory:"false" json:"inputPorts"`

	// An array of output ports.
	OutputPorts []OutputPort `mandatory:"false" json:"outputPorts"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// An array of parameters used in the data flow.
	Parameters []Parameter `mandatory:"false" json:"parameters"`

	OpConfigValues *ConfigValues `mandatory:"false" json:"opConfigValues"`

	ConfigProviderDelegate *ConfigProvider `mandatory:"false" json:"configProviderDelegate"`

	// The merge condition. The conditions are
	// ALL_SUCCESS - All the preceeding operators need to be successful.
	// ALL_FAILED - All the preceeding operators should have failed.
	// ALL_COMPLETE - All the preceeding operators should have completed. It could have executed successfully or failed.
	TriggerRule ExpressionOperatorTriggerRuleEnum `mandatory:"false" json:"triggerRule,omitempty"`
}

//GetKey returns Key
func (m ExpressionOperator) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m ExpressionOperator) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m ExpressionOperator) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m ExpressionOperator) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m ExpressionOperator) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m ExpressionOperator) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetInputPorts returns InputPorts
func (m ExpressionOperator) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m ExpressionOperator) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetObjectStatus returns ObjectStatus
func (m ExpressionOperator) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m ExpressionOperator) GetIdentifier() *string {
	return m.Identifier
}

//GetParameters returns Parameters
func (m ExpressionOperator) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m ExpressionOperator) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

func (m ExpressionOperator) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExpressionOperator) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExpressionOperatorTriggerRuleEnum(string(m.TriggerRule)); !ok && m.TriggerRule != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TriggerRule: %s. Supported values are: %s.", m.TriggerRule, strings.Join(GetExpressionOperatorTriggerRuleEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExpressionOperator) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExpressionOperator ExpressionOperator
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeExpressionOperator
	}{
		"EXPRESSION_OPERATOR",
		(MarshalTypeExpressionOperator)(m),
	}

	return json.Marshal(&s)
}

// ExpressionOperatorTriggerRuleEnum Enum with underlying type: string
type ExpressionOperatorTriggerRuleEnum string

// Set of constants representing the allowable values for ExpressionOperatorTriggerRuleEnum
const (
	ExpressionOperatorTriggerRuleSuccess  ExpressionOperatorTriggerRuleEnum = "ALL_SUCCESS"
	ExpressionOperatorTriggerRuleFailed   ExpressionOperatorTriggerRuleEnum = "ALL_FAILED"
	ExpressionOperatorTriggerRuleComplete ExpressionOperatorTriggerRuleEnum = "ALL_COMPLETE"
)

var mappingExpressionOperatorTriggerRuleEnum = map[string]ExpressionOperatorTriggerRuleEnum{
	"ALL_SUCCESS":  ExpressionOperatorTriggerRuleSuccess,
	"ALL_FAILED":   ExpressionOperatorTriggerRuleFailed,
	"ALL_COMPLETE": ExpressionOperatorTriggerRuleComplete,
}

// GetExpressionOperatorTriggerRuleEnumValues Enumerates the set of values for ExpressionOperatorTriggerRuleEnum
func GetExpressionOperatorTriggerRuleEnumValues() []ExpressionOperatorTriggerRuleEnum {
	values := make([]ExpressionOperatorTriggerRuleEnum, 0)
	for _, v := range mappingExpressionOperatorTriggerRuleEnum {
		values = append(values, v)
	}
	return values
}

// GetExpressionOperatorTriggerRuleEnumStringValues Enumerates the set of values in String for ExpressionOperatorTriggerRuleEnum
func GetExpressionOperatorTriggerRuleEnumStringValues() []string {
	return []string{
		"ALL_SUCCESS",
		"ALL_FAILED",
		"ALL_COMPLETE",
	}
}

// GetMappingExpressionOperatorTriggerRuleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExpressionOperatorTriggerRuleEnum(val string) (ExpressionOperatorTriggerRuleEnum, bool) {
	mappingExpressionOperatorTriggerRuleEnumIgnoreCase := make(map[string]ExpressionOperatorTriggerRuleEnum)
	for k, v := range mappingExpressionOperatorTriggerRuleEnum {
		mappingExpressionOperatorTriggerRuleEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingExpressionOperatorTriggerRuleEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
