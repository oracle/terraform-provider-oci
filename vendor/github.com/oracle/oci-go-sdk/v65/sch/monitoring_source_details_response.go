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

// MonitoringSourceDetailsResponse The Monitoring source.
// For configuration instructions, see
// Creating a Connector with a Monitoring Source (https://docs.oracle.com/iaas/Content/connector-hub/create-service-connector-monitoring-source.htm).
type MonitoringSourceDetailsResponse struct {

	// One or more compartment-specific lists of metric namespaces to retrieve data from.
	MonitoringSources []MonitoringSource `mandatory:"true" json:"monitoringSources"`

	PrivateEndpointMetadata *PrivateEndpointMetadata `mandatory:"false" json:"privateEndpointMetadata"`
}

// GetPrivateEndpointMetadata returns PrivateEndpointMetadata
func (m MonitoringSourceDetailsResponse) GetPrivateEndpointMetadata() *PrivateEndpointMetadata {
	return m.PrivateEndpointMetadata
}

func (m MonitoringSourceDetailsResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MonitoringSourceDetailsResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MonitoringSourceDetailsResponse) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMonitoringSourceDetailsResponse MonitoringSourceDetailsResponse
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeMonitoringSourceDetailsResponse
	}{
		"monitoring",
		(MarshalTypeMonitoringSourceDetailsResponse)(m),
	}

	return json.Marshal(&s)
}
