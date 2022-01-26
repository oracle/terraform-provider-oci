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

// Lookup The information about the lookup operator. The lookup operator has two input links, a primary input, and a lookup source input. It has an output link, fields of the lookup input are appended to the primary input and projected as the output fields.
type Lookup struct {

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

	LookupCondition *Expression `mandatory:"false" json:"lookupCondition"`

	// For the rows for which lookup condition does not satisfy, if set to true - do not return those rows of primary Input source and if set to false - create a row with primary input fields values and lookup field values as NULL.
	IsSkipNoMatch *bool `mandatory:"false" json:"isSkipNoMatch"`

	// this map is used for replacing NULL values in the record. Key is the column name and value is the NULL replacement.
	NullFillValues map[string]interface{} `mandatory:"false" json:"nullFillValues"`

	// if there are multiple records found in the lookup input what action should be performed. The default value for this field is RETURN_ANY.
	MultiMatchStrategy LookupMultiMatchStrategyEnum `mandatory:"false" json:"multiMatchStrategy,omitempty"`
}

//GetKey returns Key
func (m Lookup) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m Lookup) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m Lookup) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m Lookup) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m Lookup) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m Lookup) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetInputPorts returns InputPorts
func (m Lookup) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m Lookup) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetObjectStatus returns ObjectStatus
func (m Lookup) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m Lookup) GetIdentifier() *string {
	return m.Identifier
}

//GetParameters returns Parameters
func (m Lookup) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m Lookup) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

func (m Lookup) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m Lookup) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLookup Lookup
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeLookup
	}{
		"LOOKUP_OPERATOR",
		(MarshalTypeLookup)(m),
	}

	return json.Marshal(&s)
}

// LookupMultiMatchStrategyEnum Enum with underlying type: string
type LookupMultiMatchStrategyEnum string

// Set of constants representing the allowable values for LookupMultiMatchStrategyEnum
const (
	LookupMultiMatchStrategyAny   LookupMultiMatchStrategyEnum = "RETURN_ANY"
	LookupMultiMatchStrategyFirst LookupMultiMatchStrategyEnum = "RETURN_FIRST"
	LookupMultiMatchStrategyLast  LookupMultiMatchStrategyEnum = "RETURN_LAST"
	LookupMultiMatchStrategyAll   LookupMultiMatchStrategyEnum = "RETURN_ALL"
	LookupMultiMatchStrategyError LookupMultiMatchStrategyEnum = "RETURN_ERROR"
)

var mappingLookupMultiMatchStrategy = map[string]LookupMultiMatchStrategyEnum{
	"RETURN_ANY":   LookupMultiMatchStrategyAny,
	"RETURN_FIRST": LookupMultiMatchStrategyFirst,
	"RETURN_LAST":  LookupMultiMatchStrategyLast,
	"RETURN_ALL":   LookupMultiMatchStrategyAll,
	"RETURN_ERROR": LookupMultiMatchStrategyError,
}

// GetLookupMultiMatchStrategyEnumValues Enumerates the set of values for LookupMultiMatchStrategyEnum
func GetLookupMultiMatchStrategyEnumValues() []LookupMultiMatchStrategyEnum {
	values := make([]LookupMultiMatchStrategyEnum, 0)
	for _, v := range mappingLookupMultiMatchStrategy {
		values = append(values, v)
	}
	return values
}
