// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.cloud.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrometheusEmitterDataSource A Prometheus data source.
type PrometheusEmitterDataSource struct {

	// Identifier for DataSource. This represents the type and name for the data source associated with the Management Agent.
	Key *string `mandatory:"true" json:"key"`

	// Unique name of the DataSource.
	Name *string `mandatory:"true" json:"name"`

	// Compartment owning this DataSource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the DataSource was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the DataSource data was last received. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The url through which the Prometheus Exporter publishes its metrics. (http only)
	Url *string `mandatory:"true" json:"url"`

	// The OCI monitoring namespace to which scraped metrics should be uploaded.
	Namespace *string `mandatory:"true" json:"namespace"`

	// Comma separated metric name list. The complete set of desired scraped metrics. Use this property to limit the set of metrics uploaded if required.
	AllowMetrics *string `mandatory:"false" json:"allowMetrics"`

	// The url of the network proxy that provides access to the Prometheus Exporter's endpoint (url required property).
	ProxyUrl *string `mandatory:"false" json:"proxyUrl"`

	// Number in milliseconds. The timeout for connecting to the Prometheus Exporter's endpoint.
	ConnectionTimeout *int `mandatory:"false" json:"connectionTimeout"`

	// Number in milliseconds. The timeout for reading the response from the Prometheus Exporter's endpoint.
	ReadTimeout *int `mandatory:"false" json:"readTimeout"`

	// Number in kilobytes. The limit on the data being sent, not to exceed the agent's fixed limit of 400 (KB).
	ReadDataLimit *int `mandatory:"false" json:"readDataLimit"`

	// Number in minutes. The scraping occurs at the specified interval.
	ScheduleMins *int `mandatory:"false" json:"scheduleMins"`

	// OCI monitoring resource group to assign the metric to.
	ResourceGroup *string `mandatory:"false" json:"resourceGroup"`

	// The names of other user-supplied properties expressed as fixed values to be used as dimensions for every uploaded datapoint.
	MetricDimensions []MetricDimension `mandatory:"false" json:"metricDimensions"`

	// State of the DataSource.
	State LifecycleStatesEnum `mandatory:"true" json:"state"`
}

// GetKey returns Key
func (m PrometheusEmitterDataSource) GetKey() *string {
	return m.Key
}

// GetName returns Name
func (m PrometheusEmitterDataSource) GetName() *string {
	return m.Name
}

// GetCompartmentId returns CompartmentId
func (m PrometheusEmitterDataSource) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetState returns State
func (m PrometheusEmitterDataSource) GetState() LifecycleStatesEnum {
	return m.State
}

// GetTimeCreated returns TimeCreated
func (m PrometheusEmitterDataSource) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m PrometheusEmitterDataSource) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m PrometheusEmitterDataSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrometheusEmitterDataSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLifecycleStatesEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PrometheusEmitterDataSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePrometheusEmitterDataSource PrometheusEmitterDataSource
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypePrometheusEmitterDataSource
	}{
		"PROMETHEUS_EMITTER",
		(MarshalTypePrometheusEmitterDataSource)(m),
	}

	return json.Marshal(&s)
}
