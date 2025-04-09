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

// MetricGroupSummary A metric group defines a set of metrics to collect from a span. It uses a span filter to specify which spans to
// process. The set is then published to a namespace, which is a product level subdivision of metrics.
type MetricGroupSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the configuration item. An OCID is generated
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

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a user.
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a user.
	UpdatedBy *string `mandatory:"false" json:"updatedBy"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `mandatory:"false" json:"etag"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

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

// GetId returns Id
func (m MetricGroupSummary) GetId() *string {
	return m.Id
}

// GetTimeCreated returns TimeCreated
func (m MetricGroupSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m MetricGroupSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetCreatedBy returns CreatedBy
func (m MetricGroupSummary) GetCreatedBy() *string {
	return m.CreatedBy
}

// GetUpdatedBy returns UpdatedBy
func (m MetricGroupSummary) GetUpdatedBy() *string {
	return m.UpdatedBy
}

// GetEtag returns Etag
func (m MetricGroupSummary) GetEtag() *string {
	return m.Etag
}

// GetFreeformTags returns FreeformTags
func (m MetricGroupSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m MetricGroupSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m MetricGroupSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m MetricGroupSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MetricGroupSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MetricGroupSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMetricGroupSummary MetricGroupSummary
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeMetricGroupSummary
	}{
		"METRIC_GROUP",
		(MarshalTypeMetricGroupSummary)(m),
	}

	return json.Marshal(&s)
}
