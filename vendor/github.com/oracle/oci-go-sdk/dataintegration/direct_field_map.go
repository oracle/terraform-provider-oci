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

// DirectFieldMap The information about a field map.
type DirectFieldMap struct {

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// Reference to a typed object
	SourceTypedObject *string `mandatory:"false" json:"sourceTypedObject"`

	// Reference to a typed object
	TargetTypedObject *string `mandatory:"false" json:"targetTypedObject"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`
}

//GetDescription returns Description
func (m DirectFieldMap) GetDescription() *string {
	return m.Description
}

func (m DirectFieldMap) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DirectFieldMap) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDirectFieldMap DirectFieldMap
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDirectFieldMap
	}{
		"DIRECT_FIELD_MAP",
		(MarshalTypeDirectFieldMap)(m),
	}

	return json.Marshal(&s)
}
