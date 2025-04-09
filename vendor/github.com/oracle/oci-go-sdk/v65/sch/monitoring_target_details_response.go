// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connector Hub API
//
// Use the Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Connector Hub, see
// the Connector Hub documentation (https://docs.oracle.com/iaas/Content/connector-hub/home.htm).
// Connector Hub is formerly known as Service Connector Hub.
//

package sch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MonitoringTargetDetailsResponse The destination metric for data transferred from the source.
// For configuration instructions, see
// Creating a Connector (https://docs.oracle.com/iaas/Content/connector-hub/create-service-connector.htm).
type MonitoringTargetDetailsResponse struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the metric.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The namespace of the metric.
	// Example: `oci_computeagent`
	MetricNamespace *string `mandatory:"true" json:"metricNamespace"`

	// The name of the metric.
	// Example: `CpuUtilization`
	Metric *string `mandatory:"true" json:"metric"`

	PrivateEndpointMetadata *PrivateEndpointMetadata `mandatory:"false" json:"privateEndpointMetadata"`

	// List of dimension names and values.
	Dimensions []DimensionDetails `mandatory:"false" json:"dimensions"`
}

// GetPrivateEndpointMetadata returns PrivateEndpointMetadata
func (m MonitoringTargetDetailsResponse) GetPrivateEndpointMetadata() *PrivateEndpointMetadata {
	return m.PrivateEndpointMetadata
}

func (m MonitoringTargetDetailsResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MonitoringTargetDetailsResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MonitoringTargetDetailsResponse) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMonitoringTargetDetailsResponse MonitoringTargetDetailsResponse
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeMonitoringTargetDetailsResponse
	}{
		"monitoring",
		(MarshalTypeMonitoringTargetDetailsResponse)(m),
	}

	return json.Marshal(&s)
}
