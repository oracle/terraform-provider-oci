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

// SpanFilter A span filter is a named setting that specifies filter criteria to match a subset of the spans.
type SpanFilter struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the configuration item. An OCID will be generated
	// when the item is created.
	Id *string `mandatory:"false" json:"id"`

	// The time the resource was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-12T22:47:12.613Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the resource was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-13T22:47:12.613Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The name by which this filter can be displayed in the UI.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The string that defines the Span Filter expression.
	FilterText *string `mandatory:"false" json:"filterText"`

	// An optional string that describes what the filter is intended or used for.
	Description *string `mandatory:"false" json:"description"`
}

//GetId returns Id
func (m SpanFilter) GetId() *string {
	return m.Id
}

//GetTimeCreated returns TimeCreated
func (m SpanFilter) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m SpanFilter) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetFreeformTags returns FreeformTags
func (m SpanFilter) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m SpanFilter) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m SpanFilter) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m SpanFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSpanFilter SpanFilter
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeSpanFilter
	}{
		"SPAN_FILTER",
		(MarshalTypeSpanFilter)(m),
	}

	return json.Marshal(&s)
}
