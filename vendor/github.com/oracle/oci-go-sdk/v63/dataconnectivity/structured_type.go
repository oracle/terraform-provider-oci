// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v63/common"
	"strings"
)

// StructuredType A `StructuredType` object represents a data type that exists in a physical data asset object such as a table column, but is more complex. For example, an Oracle database `OBJECT` type. It can be composed of multiple `DataType` objects.
type StructuredType struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// A user defined description for the object.
	Description *string `mandatory:"false" json:"description"`

	Schema BaseType `mandatory:"false" json:"schema"`
}

//GetKey returns Key
func (m StructuredType) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m StructuredType) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m StructuredType) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m StructuredType) GetName() *string {
	return m.Name
}

//GetObjectStatus returns ObjectStatus
func (m StructuredType) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetDescription returns Description
func (m StructuredType) GetDescription() *string {
	return m.Description
}

func (m StructuredType) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StructuredType) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m StructuredType) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStructuredType StructuredType
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeStructuredType
	}{
		"STRUCTURED_TYPE",
		(MarshalTypeStructuredType)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *StructuredType) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key          *string          `json:"key"`
		ModelVersion *string          `json:"modelVersion"`
		ParentRef    *ParentReference `json:"parentRef"`
		Name         *string          `json:"name"`
		ObjectStatus *int             `json:"objectStatus"`
		Description  *string          `json:"description"`
		Schema       basetype         `json:"schema"`
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

	nn, e = model.Schema.UnmarshalPolymorphicJSON(model.Schema.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Schema = nn.(BaseType)
	} else {
		m.Schema = nil
	}

	return
}
