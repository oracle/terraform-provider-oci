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

// ConfiguredType A ConfiguraedType represents a type that has built-in configuration to the type itself. An example is a SSN type whose basic type is VARCHAR, but the type itself also has a built-in configuration like length=10
type ConfiguredType struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	WrappedType BaseType `mandatory:"false" json:"wrappedType"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	ConfigDefinition *ConfigDefinition `mandatory:"false" json:"configDefinition"`
}

//GetKey returns Key
func (m ConfiguredType) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m ConfiguredType) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m ConfiguredType) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m ConfiguredType) GetName() *string {
	return m.Name
}

//GetObjectStatus returns ObjectStatus
func (m ConfiguredType) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetDescription returns Description
func (m ConfiguredType) GetDescription() *string {
	return m.Description
}

func (m ConfiguredType) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ConfiguredType) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeConfiguredType ConfiguredType
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeConfiguredType
	}{
		"CONFIGURED_TYPE",
		(MarshalTypeConfiguredType)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ConfiguredType) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key              *string           `json:"key"`
		ModelVersion     *string           `json:"modelVersion"`
		ParentRef        *ParentReference  `json:"parentRef"`
		Name             *string           `json:"name"`
		ObjectStatus     *int              `json:"objectStatus"`
		Description      *string           `json:"description"`
		WrappedType      basetype          `json:"wrappedType"`
		ConfigValues     *ConfigValues     `json:"configValues"`
		ConfigDefinition *ConfigDefinition `json:"configDefinition"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.Name = model.Name

	m.ObjectStatus = model.ObjectStatus

	m.Description = model.Description

	nn, e = model.WrappedType.UnmarshalPolymorphicJSON(model.WrappedType.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.WrappedType = nn.(BaseType)
	} else {
		m.WrappedType = nil
	}

	m.ConfigValues = model.ConfigValues

	m.ConfigDefinition = model.ConfigDefinition

	return
}
