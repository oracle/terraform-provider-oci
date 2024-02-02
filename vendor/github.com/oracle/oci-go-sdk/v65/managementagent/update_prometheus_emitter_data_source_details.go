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

// UpdatePrometheusEmitterDataSourceDetails A Prometheus data source.
type UpdatePrometheusEmitterDataSourceDetails struct {

	// The url through which the Prometheus Exporter publishes its metrics. (http only)
	Url *string `mandatory:"true" json:"url"`

	// Comma separated metric name list. The complete set of desired scraped metrics. Use this property to limit the set of metrics uploaded if required.
	AllowMetrics *string `mandatory:"false" json:"allowMetrics"`

	// The url of the network proxy that provides access to the Prometheus Exporter's endpoint (url required property).
	ProxyUrl *string `mandatory:"false" json:"proxyUrl"`

	// Number in milliseconds. The timeout for connecting to the Prometheus Exporter's endpoint.
	ConnectionTimeout *int `mandatory:"false" json:"connectionTimeout"`

	// Number in milliseconds. The timeout for reading the response from the Prometheus Exporter's endpoint.
	ReadTimeout *int `mandatory:"false" json:"readTimeout"`

	// Number in kilobytes. The limit on the data being sent, not to exceed the agent's fixed limit of 400 (KB).
	ReadDataLimitInKilobytes *int `mandatory:"false" json:"readDataLimitInKilobytes"`

	// Number in minutes. The scraping occurs at the specified interval.
	ScheduleMins *int `mandatory:"false" json:"scheduleMins"`

	// OCI monitoring resource group to assign the metric to.
	ResourceGroup *string `mandatory:"false" json:"resourceGroup"`

	// The names of other user-supplied properties expressed as fixed values to be used as dimensions for every uploaded datapoint.
	MetricDimensions []MetricDimension `mandatory:"false" json:"metricDimensions"`
}

func (m UpdatePrometheusEmitterDataSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdatePrometheusEmitterDataSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdatePrometheusEmitterDataSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdatePrometheusEmitterDataSourceDetails UpdatePrometheusEmitterDataSourceDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdatePrometheusEmitterDataSourceDetails
	}{
		"PROMETHEUS_EMITTER",
		(MarshalTypeUpdatePrometheusEmitterDataSourceDetails)(m),
	}

	return json.Marshal(&s)
}
