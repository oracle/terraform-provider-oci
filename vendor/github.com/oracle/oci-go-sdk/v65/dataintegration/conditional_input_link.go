// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ConditionalInputLink The information about the conditional input link.
type ConditionalInputLink struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// Key of FlowPort reference
	Port *string `mandatory:"false" json:"port"`

	FromLink *OutputLink `mandatory:"false" json:"fromLink"`

	FieldMap FieldMap `mandatory:"false" json:"fieldMap"`

	Condition *Expression `mandatory:"false" json:"condition"`
}

// GetKey returns Key
func (m ConditionalInputLink) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m ConditionalInputLink) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m ConditionalInputLink) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetObjectStatus returns ObjectStatus
func (m ConditionalInputLink) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetDescription returns Description
func (m ConditionalInputLink) GetDescription() *string {
	return m.Description
}

// GetPort returns Port
func (m ConditionalInputLink) GetPort() *string {
	return m.Port
}

func (m ConditionalInputLink) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConditionalInputLink) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ConditionalInputLink) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeConditionalInputLink ConditionalInputLink
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeConditionalInputLink
	}{
		"CONDITIONAL_INPUT_LINK",
		(MarshalTypeConditionalInputLink)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ConditionalInputLink) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key          *string          `json:"key"`
		ModelVersion *string          `json:"modelVersion"`
		ParentRef    *ParentReference `json:"parentRef"`
		ObjectStatus *int             `json:"objectStatus"`
		Description  *string          `json:"description"`
		Port         *string          `json:"port"`
		FromLink     *OutputLink      `json:"fromLink"`
		FieldMap     fieldmap         `json:"fieldMap"`
		Condition    *Expression      `json:"condition"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.ObjectStatus = model.ObjectStatus

	m.Description = model.Description

	m.Port = model.Port

	m.FromLink = model.FromLink

	nn, e = model.FieldMap.UnmarshalPolymorphicJSON(model.FieldMap.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.FieldMap = nn.(FieldMap)
	} else {
		m.FieldMap = nil
	}

	m.Condition = model.Condition

	return
}
