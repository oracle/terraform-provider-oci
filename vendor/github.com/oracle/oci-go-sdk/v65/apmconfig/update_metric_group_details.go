// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Configuration API
//
// Use the Application Performance Monitoring Configuration API to query and set Application Performance Monitoring
// configuration. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmconfig

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateMetricGroupDetails A metric group defines a set of metrics to collect from a span. It uses a span filter to specify which spans to
// process. The set is then published to a namespace, which is a product level subdivision of metrics.
type UpdateMetricGroupDetails struct {

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The name by which a configuration entity is displayed to the end user.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Span Filter. The filterId is mandatory for the creation
	// of MetricGroups. A filterId is generated when a Span Filter is created.
	FilterId *string `mandatory:"false" json:"filterId"`

	// The namespace to which the metrics are published. It must be one of several predefined namespaces.
	Namespace *string `mandatory:"false" json:"namespace"`

	// A list of dimensions for the metric. This variable should not be used.
	Dimensions []Dimension `mandatory:"false" json:"dimensions"`

	// The list of metrics in this group.
	Metrics []Metric `mandatory:"false" json:"metrics"`
}

// GetFreeformTags returns FreeformTags
func (m UpdateMetricGroupDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateMetricGroupDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateMetricGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateMetricGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
