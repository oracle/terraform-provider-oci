// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// PrimaryKey The primary key object.
type PrimaryKey struct {

	// The object key.
	Key *string `mandatory:"false" json:"key"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// An array of attribute references.
	AttributeRefs []KeyAttribute `mandatory:"false" json:"attributeRefs"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`
}

//GetKey returns Key
func (m PrimaryKey) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m PrimaryKey) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m PrimaryKey) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m PrimaryKey) GetName() *string {
	return m.Name
}

//GetAttributeRefs returns AttributeRefs
func (m PrimaryKey) GetAttributeRefs() []KeyAttribute {
	return m.AttributeRefs
}

//GetObjectStatus returns ObjectStatus
func (m PrimaryKey) GetObjectStatus() *int {
	return m.ObjectStatus
}

func (m PrimaryKey) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m PrimaryKey) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePrimaryKey PrimaryKey
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypePrimaryKey
	}{
		"PRIMARY_KEY",
		(MarshalTypePrimaryKey)(m),
	}

	return json.Marshal(&s)
}
