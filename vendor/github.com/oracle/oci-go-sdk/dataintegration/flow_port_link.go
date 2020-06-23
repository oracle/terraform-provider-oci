// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// FlowPortLink The details of the flow port links.
type FlowPortLink interface {

	// The key of the object.
	GetKey() *string

	// The model version of an object.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Detailed description for the object.
	GetDescription() *string

	// Key of FlowPort reference
	GetPort() *string
}

type flowportlink struct {
	JsonData     []byte
	Key          *string          `mandatory:"false" json:"key"`
	ModelVersion *string          `mandatory:"false" json:"modelVersion"`
	ParentRef    *ParentReference `mandatory:"false" json:"parentRef"`
	ObjectStatus *int             `mandatory:"false" json:"objectStatus"`
	Description  *string          `mandatory:"false" json:"description"`
	Port         *string          `mandatory:"false" json:"port"`
	ModelType    string           `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *flowportlink) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerflowportlink flowportlink
	s := struct {
		Model Unmarshalerflowportlink
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.ModelVersion = s.Model.ModelVersion
	m.ParentRef = s.Model.ParentRef
	m.ObjectStatus = s.Model.ObjectStatus
	m.Description = s.Model.Description
	m.Port = s.Model.Port
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *flowportlink) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "INPUT_LINK":
		mm := InputLink{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OUTPUT_LINK":
		mm := OutputLink{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CONDITIONAL_INPUT_LINK":
		mm := ConditionalInputLink{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetKey returns Key
func (m flowportlink) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m flowportlink) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m flowportlink) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetObjectStatus returns ObjectStatus
func (m flowportlink) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetDescription returns Description
func (m flowportlink) GetDescription() *string {
	return m.Description
}

//GetPort returns Port
func (m flowportlink) GetPort() *string {
	return m.Port
}

func (m flowportlink) String() string {
	return common.PointerString(m)
}

// FlowPortLinkModelTypeEnum Enum with underlying type: string
type FlowPortLinkModelTypeEnum string

// Set of constants representing the allowable values for FlowPortLinkModelTypeEnum
const (
	FlowPortLinkModelTypeConditionalInputLink FlowPortLinkModelTypeEnum = "CONDITIONAL_INPUT_LINK"
	FlowPortLinkModelTypeOutputLink           FlowPortLinkModelTypeEnum = "OUTPUT_LINK"
	FlowPortLinkModelTypeInputLink            FlowPortLinkModelTypeEnum = "INPUT_LINK"
)

var mappingFlowPortLinkModelType = map[string]FlowPortLinkModelTypeEnum{
	"CONDITIONAL_INPUT_LINK": FlowPortLinkModelTypeConditionalInputLink,
	"OUTPUT_LINK":            FlowPortLinkModelTypeOutputLink,
	"INPUT_LINK":             FlowPortLinkModelTypeInputLink,
}

// GetFlowPortLinkModelTypeEnumValues Enumerates the set of values for FlowPortLinkModelTypeEnum
func GetFlowPortLinkModelTypeEnumValues() []FlowPortLinkModelTypeEnum {
	values := make([]FlowPortLinkModelTypeEnum, 0)
	for _, v := range mappingFlowPortLinkModelType {
		values = append(values, v)
	}
	return values
}
