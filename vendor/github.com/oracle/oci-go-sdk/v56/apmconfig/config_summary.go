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

// ConfigSummary The request body used to update the Configuration entity. It must specify the data to update the item with.
type ConfigSummary interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the configuration item. An OCID will be generated
	// when the item is created.
	GetId() *string

	// The time the resource was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-12T22:47:12.613Z`
	GetTimeCreated() *common.SDKTime

	// The time the resource was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-13T22:47:12.613Z`
	GetTimeUpdated() *common.SDKTime

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type configsummary struct {
	JsonData     []byte
	Id           *string                           `mandatory:"false" json:"id"`
	TimeCreated  *common.SDKTime                   `mandatory:"false" json:"timeCreated"`
	TimeUpdated  *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	FreeformTags map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags  map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	ConfigType   string                            `json:"configType"`
}

// UnmarshalJSON unmarshals json
func (m *configsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconfigsummary configsummary
	s := struct {
		Model Unmarshalerconfigsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.ConfigType = s.Model.ConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *configsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigType {
	case "METRIC_GROUP":
		mm := MetricGroupSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APDEX":
		mm := ApdexRulesSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SPAN_FILTER":
		mm := SpanFilterSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m configsummary) GetId() *string {
	return m.Id
}

//GetTimeCreated returns TimeCreated
func (m configsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m configsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetFreeformTags returns FreeformTags
func (m configsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m configsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m configsummary) String() string {
	return common.PointerString(m)
}
