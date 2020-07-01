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

// InputLink The information about input links.
type InputLink struct {

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

	// From link reference.
	FromLink *string `mandatory:"false" json:"fromLink"`

	FieldMap FieldMap `mandatory:"false" json:"fieldMap"`
}

//GetKey returns Key
func (m InputLink) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m InputLink) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m InputLink) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetObjectStatus returns ObjectStatus
func (m InputLink) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetDescription returns Description
func (m InputLink) GetDescription() *string {
	return m.Description
}

//GetPort returns Port
func (m InputLink) GetPort() *string {
	return m.Port
}

func (m InputLink) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InputLink) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInputLink InputLink
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeInputLink
	}{
		"INPUT_LINK",
		(MarshalTypeInputLink)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *InputLink) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key          *string          `json:"key"`
		ModelVersion *string          `json:"modelVersion"`
		ParentRef    *ParentReference `json:"parentRef"`
		ObjectStatus *int             `json:"objectStatus"`
		Description  *string          `json:"description"`
		Port         *string          `json:"port"`
		FromLink     *string          `json:"fromLink"`
		FieldMap     fieldmap         `json:"fieldMap"`
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

	return
}
