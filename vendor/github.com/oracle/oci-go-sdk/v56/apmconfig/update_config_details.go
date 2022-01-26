// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Configuration API
//
// An API for the APM Configuration service. Use this API to query and set APM configuration.
//

package apmconfig

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateConfigDetails The request body used to update the Configuration entity. It must specify the data to update the item with.
type UpdateConfigDetails interface {

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type updateconfigdetails struct {
	JsonData     []byte
	FreeformTags map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags  map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	ConfigType   string                            `json:"configType"`
}

// UnmarshalJSON unmarshals json
func (m *updateconfigdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateconfigdetails updateconfigdetails
	s := struct {
		Model Unmarshalerupdateconfigdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.ConfigType = s.Model.ConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateconfigdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigType {
	case "METRIC_GROUP":
		mm := UpdateMetricGroupDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APDEX":
		mm := UpdateApdexRulesDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SPAN_FILTER":
		mm := UpdateSpanFilterDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetFreeformTags returns FreeformTags
func (m updateconfigdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m updateconfigdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m updateconfigdetails) String() string {
	return common.PointerString(m)
}
