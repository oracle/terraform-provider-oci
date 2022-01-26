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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// MergeOperator Represents the start of a pipeline.
type MergeOperator struct {

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

	// The merge condition. The conditions are
	// ALL_SUCCESS - All the preceeding operators need to be successful.
	// ALL_FAILED - All the preceeding operators should have failed.
	// ALL_COMPLETE - All the preceeding operators should have completed. It could have executed successfully or failed.
	// ONE_SUCCESS - Atleast one of the preceeding operators should have succeeded.
	// ONE_FAILED - Atleast one of the preceeding operators should have failed.
	TriggerRule MergeOperatorTriggerRuleEnum `mandatory:"false" json:"triggerRule,omitempty"`
}

//GetKey returns Key
func (m MergeOperator) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m MergeOperator) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m MergeOperator) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m MergeOperator) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m MergeOperator) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m MergeOperator) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetInputPorts returns InputPorts
func (m MergeOperator) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m MergeOperator) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetObjectStatus returns ObjectStatus
func (m MergeOperator) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m MergeOperator) GetIdentifier() *string {
	return m.Identifier
}

//GetParameters returns Parameters
func (m MergeOperator) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m MergeOperator) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

func (m MergeOperator) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m MergeOperator) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMergeOperator MergeOperator
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeMergeOperator
	}{
		"MERGE_OPERATOR",
		(MarshalTypeMergeOperator)(m),
	}

	return json.Marshal(&s)
}

// MergeOperatorTriggerRuleEnum Enum with underlying type: string
type MergeOperatorTriggerRuleEnum string

// Set of constants representing the allowable values for MergeOperatorTriggerRuleEnum
const (
	MergeOperatorTriggerRuleAllSuccess  MergeOperatorTriggerRuleEnum = "ALL_SUCCESS"
	MergeOperatorTriggerRuleAllFailed   MergeOperatorTriggerRuleEnum = "ALL_FAILED"
	MergeOperatorTriggerRuleAllComplete MergeOperatorTriggerRuleEnum = "ALL_COMPLETE"
	MergeOperatorTriggerRuleOneSuccess  MergeOperatorTriggerRuleEnum = "ONE_SUCCESS"
	MergeOperatorTriggerRuleOneFailed   MergeOperatorTriggerRuleEnum = "ONE_FAILED"
)

var mappingMergeOperatorTriggerRule = map[string]MergeOperatorTriggerRuleEnum{
	"ALL_SUCCESS":  MergeOperatorTriggerRuleAllSuccess,
	"ALL_FAILED":   MergeOperatorTriggerRuleAllFailed,
	"ALL_COMPLETE": MergeOperatorTriggerRuleAllComplete,
	"ONE_SUCCESS":  MergeOperatorTriggerRuleOneSuccess,
	"ONE_FAILED":   MergeOperatorTriggerRuleOneFailed,
}

// GetMergeOperatorTriggerRuleEnumValues Enumerates the set of values for MergeOperatorTriggerRuleEnum
func GetMergeOperatorTriggerRuleEnumValues() []MergeOperatorTriggerRuleEnum {
	values := make([]MergeOperatorTriggerRuleEnum, 0)
	for _, v := range mappingMergeOperatorTriggerRule {
		values = append(values, v)
	}
	return values
}
