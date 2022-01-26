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

// UpdateMetricGroupDetails A Metric Group.
type UpdateMetricGroupDetails struct {

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The name of this metric group
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of a Span Filter. The filterId is mandatory for the creation
	// of MetricGroups. A filterId will be generated when a Span Filter is created.
	FilterId *string `mandatory:"false" json:"filterId"`

	// The namespace to write the metrics to
	Namespace *string `mandatory:"false" json:"namespace"`

	// A list of dimensions for this metric
	Dimensions []Dimension `mandatory:"false" json:"dimensions"`

	Metrics []Metric `mandatory:"false" json:"metrics"`
}

//GetFreeformTags returns FreeformTags
func (m UpdateMetricGroupDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateMetricGroupDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateMetricGroupDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateMetricGroupDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateMetricGroupDetails UpdateMetricGroupDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeUpdateMetricGroupDetails
	}{
		"METRIC_GROUP",
		(MarshalTypeUpdateMetricGroupDetails)(m),
	}

	return json.Marshal(&s)
}
