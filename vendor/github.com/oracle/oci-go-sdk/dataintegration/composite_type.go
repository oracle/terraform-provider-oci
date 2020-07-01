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

// CompositeType A CompositeType represents a type that is composed of a list of sub-types, for example an "Address" type.   The sub-types can be simple DataType or other CompositeType objects. Thus in general a CompositeType may represent an arbitrarily deep hierarchy of types.
type CompositeType struct {

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

	ParentType *CompositeType `mandatory:"false" json:"parentType"`

	// elements
	Elements []TypedObject `mandatory:"false" json:"elements"`

	ConfigDefinition *ConfigDefinition `mandatory:"false" json:"configDefinition"`
}

//GetKey returns Key
func (m CompositeType) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m CompositeType) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m CompositeType) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m CompositeType) GetName() *string {
	return m.Name
}

//GetObjectStatus returns ObjectStatus
func (m CompositeType) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetDescription returns Description
func (m CompositeType) GetDescription() *string {
	return m.Description
}

func (m CompositeType) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CompositeType) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCompositeType CompositeType
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeCompositeType
	}{
		"COMPOSITE_TYPE",
		(MarshalTypeCompositeType)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CompositeType) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key              *string           `json:"key"`
		ModelVersion     *string           `json:"modelVersion"`
		ParentRef        *ParentReference  `json:"parentRef"`
		Name             *string           `json:"name"`
		ObjectStatus     *int              `json:"objectStatus"`
		Description      *string           `json:"description"`
		ParentType       *CompositeType    `json:"parentType"`
		Elements         []typedobject     `json:"elements"`
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

	m.ParentType = model.ParentType

	m.Elements = make([]TypedObject, len(model.Elements))
	for i, n := range model.Elements {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Elements[i] = nn.(TypedObject)
		} else {
			m.Elements[i] = nil
		}
	}

	m.ConfigDefinition = model.ConfigDefinition

	return
}
