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

// DynamicInputField The type representing the dynamic field concept. Dynamic fields have a dynamic type handler to define how to generate the field.
type DynamicInputField struct {

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

	Type BaseType `mandatory:"false" json:"type"`

	// Labels are keywords or labels that you can add to data assets, dataflows and so on. You can define your own labels and use them to categorize content.
	Labels []string `mandatory:"false" json:"labels"`
}

// GetKey returns Key
func (m DynamicInputField) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m DynamicInputField) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m DynamicInputField) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetConfigValues returns ConfigValues
func (m DynamicInputField) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

// GetObjectStatus returns ObjectStatus
func (m DynamicInputField) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetName returns Name
func (m DynamicInputField) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m DynamicInputField) GetDescription() *string {
	return m.Description
}

func (m DynamicInputField) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DynamicInputField) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DynamicInputField) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDynamicInputField DynamicInputField
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDynamicInputField
	}{
		"DYNAMIC_INPUT_FIELD",
		(MarshalTypeDynamicInputField)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DynamicInputField) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key          *string          `json:"key"`
		ModelVersion *string          `json:"modelVersion"`
		ParentRef    *ParentReference `json:"parentRef"`
		ConfigValues *ConfigValues    `json:"configValues"`
		ObjectStatus *int             `json:"objectStatus"`
		Name         *string          `json:"name"`
		Description  *string          `json:"description"`
		Type         basetype         `json:"type"`
		Labels       []string         `json:"labels"`
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

	nn, e = model.Type.UnmarshalPolymorphicJSON(model.Type.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Type = nn.(BaseType)
	} else {
		m.Type = nil
	}

	m.Labels = make([]string, len(model.Labels))
	copy(m.Labels, model.Labels)
	return
}
