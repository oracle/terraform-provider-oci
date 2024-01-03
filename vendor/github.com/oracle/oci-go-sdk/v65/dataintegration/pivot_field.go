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

// PivotField The type representing the pivot field. Pivot fields have an expression to define a macro and a pattern to generate the column name
type PivotField struct {

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
}

// GetKey returns Key
func (m PivotField) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m PivotField) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m PivotField) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetConfigValues returns ConfigValues
func (m PivotField) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

// GetObjectStatus returns ObjectStatus
func (m PivotField) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetName returns Name
func (m PivotField) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m PivotField) GetDescription() *string {
	return m.Description
}

func (m PivotField) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PivotField) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PivotField) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePivotField PivotField
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypePivotField
	}{
		"PIVOT_FIELD",
		(MarshalTypePivotField)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *PivotField) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key               *string          `json:"key"`
		ModelVersion      *string          `json:"modelVersion"`
		ParentRef         *ParentReference `json:"parentRef"`
		ConfigValues      *ConfigValues    `json:"configValues"`
		ObjectStatus      *int             `json:"objectStatus"`
		Name              *string          `json:"name"`
		Description       *string          `json:"description"`
		Expr              *Expression      `json:"expr"`
		UseType           *ConfiguredType  `json:"useType"`
		Type              basetype         `json:"type"`
		ColumnNamePattern *string          `json:"columnNamePattern"`
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

	return
}
