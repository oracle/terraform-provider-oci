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

// ShapeField The shape field object.
type ShapeField struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The reference to the type.
	Type *string `mandatory:"false" json:"type"`

	// Labels are keywords or labels that you can add to data assets, dataflows etc. You can define your own labels and use them to categorize content.
	Labels []string `mandatory:"false" json:"labels"`

	NativeShapeField *NativeShapeField `mandatory:"false" json:"nativeShapeField"`
}

//GetKey returns Key
func (m ShapeField) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m ShapeField) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m ShapeField) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetConfigValues returns ConfigValues
func (m ShapeField) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

//GetObjectStatus returns ObjectStatus
func (m ShapeField) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetName returns Name
func (m ShapeField) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m ShapeField) GetDescription() *string {
	return m.Description
}

func (m ShapeField) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ShapeField) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeShapeField ShapeField
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeShapeField
	}{
		"SHAPE_FIELD",
		(MarshalTypeShapeField)(m),
	}

	return json.Marshal(&s)
}
