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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// FlowNode The flow node can be connected to other nodes in a data flow with input and output links and is bound to an opertor which defines the semantics of the node.
type FlowNode struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The type of the object.
	ModelType *string `mandatory:"false" json:"modelType"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// An array of input links.
	InputLinks []InputLink `mandatory:"false" json:"inputLinks"`

	// An array of output links.
	OutputLinks []OutputLink `mandatory:"false" json:"outputLinks"`

	Operator Operator `mandatory:"false" json:"operator"`

	UiProperties *UiProperties `mandatory:"false" json:"uiProperties"`

	ConfigProviderDelegate *ConfigProvider `mandatory:"false" json:"configProviderDelegate"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`
}

func (m FlowNode) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FlowNode) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *FlowNode) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key                    *string          `json:"key"`
		ModelType              *string          `json:"modelType"`
		ModelVersion           *string          `json:"modelVersion"`
		ParentRef              *ParentReference `json:"parentRef"`
		Name                   *string          `json:"name"`
		Description            *string          `json:"description"`
		InputLinks             []InputLink      `json:"inputLinks"`
		OutputLinks            []OutputLink     `json:"outputLinks"`
		Operator               operator         `json:"operator"`
		UiProperties           *UiProperties    `json:"uiProperties"`
		ConfigProviderDelegate *ConfigProvider  `json:"configProviderDelegate"`
		ObjectStatus           *int             `json:"objectStatus"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelType = model.ModelType

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.Name = model.Name

	m.Description = model.Description

	m.InputLinks = make([]InputLink, len(model.InputLinks))
	for i, n := range model.InputLinks {
		m.InputLinks[i] = n
	}

	m.OutputLinks = make([]OutputLink, len(model.OutputLinks))
	for i, n := range model.OutputLinks {
		m.OutputLinks[i] = n
	}

	nn, e = model.Operator.UnmarshalPolymorphicJSON(model.Operator.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Operator = nn.(Operator)
	} else {
		m.Operator = nil
	}

	m.UiProperties = model.UiProperties

	m.ConfigProviderDelegate = model.ConfigProviderDelegate

	m.ObjectStatus = model.ObjectStatus

	return
}
