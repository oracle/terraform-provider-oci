// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AgentExtensionHandlerMetricMappingDetails Specific metric mapping configurations for Agent Extension Handlers.
type AgentExtensionHandlerMetricMappingDetails struct {

	// Metric name as defined by the collector.
	CollectorMetricName *string `mandatory:"true" json:"collectorMetricName"`

	// Metric name to be upload to telemetry.
	TelemetryMetricName *string `mandatory:"false" json:"telemetryMetricName"`

	// Is ignoring this metric.
	IsSkipUpload *bool `mandatory:"false" json:"isSkipUpload"`

	// Metric upload interval in seconds. Any metric sent by telegraf/collectd before the
	// configured interval expires will be dropped.
	MetricUploadIntervalInSeconds *int `mandatory:"false" json:"metricUploadIntervalInSeconds"`
}

func (m AgentExtensionHandlerMetricMappingDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AgentExtensionHandlerMetricMappingDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
