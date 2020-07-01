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

// CompositeFieldMap A composite field map.
type CompositeFieldMap struct {

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// An array of field maps.
	FieldMaps []FieldMap `mandatory:"false" json:"fieldMaps"`
}

//GetDescription returns Description
func (m CompositeFieldMap) GetDescription() *string {
	return m.Description
}

func (m CompositeFieldMap) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CompositeFieldMap) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCompositeFieldMap CompositeFieldMap
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeCompositeFieldMap
	}{
		"COMPOSITE_FIELD_MAP",
		(MarshalTypeCompositeFieldMap)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CompositeFieldMap) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description  *string          `json:"description"`
		Key          *string          `json:"key"`
		ModelVersion *string          `json:"modelVersion"`
		ParentRef    *ParentReference `json:"parentRef"`
		ConfigValues *ConfigValues    `json:"configValues"`
		ObjectStatus *int             `json:"objectStatus"`
		FieldMaps    []fieldmap       `json:"fieldMaps"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.ConfigValues = model.ConfigValues

	m.ObjectStatus = model.ObjectStatus

	m.FieldMaps = make([]FieldMap, len(model.FieldMaps))
	for i, n := range model.FieldMaps {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.FieldMaps[i] = nn.(FieldMap)
		} else {
			m.FieldMaps[i] = nil
		}
	}

	return
}
