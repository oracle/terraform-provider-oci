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

// ReadOperationConfig The information about the read operation.
type ReadOperationConfig struct {

	// The object key.
	Key *string `mandatory:"false" json:"key"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// An array of operations.
	Operations []PushDownOperation `mandatory:"false" json:"operations"`

	DataFormat *DataFormat `mandatory:"false" json:"dataFormat"`

	PartitionConfig PartitionConfig `mandatory:"false" json:"partitionConfig"`

	ReadAttribute AbstractReadAttribute `mandatory:"false" json:"readAttribute"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`
}

func (m ReadOperationConfig) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ReadOperationConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeReadOperationConfig ReadOperationConfig
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeReadOperationConfig
	}{
		"READ_OPERATION_CONFIG",
		(MarshalTypeReadOperationConfig)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ReadOperationConfig) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key             *string               `json:"key"`
		ModelVersion    *string               `json:"modelVersion"`
		ParentRef       *ParentReference      `json:"parentRef"`
		Operations      []pushdownoperation   `json:"operations"`
		DataFormat      *DataFormat           `json:"dataFormat"`
		PartitionConfig partitionconfig       `json:"partitionConfig"`
		ReadAttribute   abstractreadattribute `json:"readAttribute"`
		ObjectStatus    *int                  `json:"objectStatus"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.Operations = make([]PushDownOperation, len(model.Operations))
	for i, n := range model.Operations {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Operations[i] = nn.(PushDownOperation)
		} else {
			m.Operations[i] = nil
		}
	}

	m.DataFormat = model.DataFormat

	nn, e = model.PartitionConfig.UnmarshalPolymorphicJSON(model.PartitionConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PartitionConfig = nn.(PartitionConfig)
	} else {
		m.PartitionConfig = nil
	}

	nn, e = model.ReadAttribute.UnmarshalPolymorphicJSON(model.ReadAttribute.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ReadAttribute = nn.(AbstractReadAttribute)
	} else {
		m.ReadAttribute = nil
	}

	m.ObjectStatus = model.ObjectStatus

	return
}
