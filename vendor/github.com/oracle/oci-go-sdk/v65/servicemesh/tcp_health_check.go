// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TcpHealthCheck TCP based health check.
type TcpHealthCheck struct {

	// The time to wait for a health check response. If the timeout is reached the health check attempt
	// will be considered a failure.
	TimeoutInMs *int64 `mandatory:"true" json:"timeoutInMs"`

	// The interval between health checks.
	IntervalInMs *int64 `mandatory:"true" json:"intervalInMs"`

	// Base64 encoded content of the message which should be sent during the health checks.
	// Use empty string to do a connect-only health check.
	Send []byte `mandatory:"true" json:"send"`

	// The number of unhealthy health checks required before a host is marked unhealthy.
	UnhealthyThreshold *int `mandatory:"false" json:"unhealthyThreshold"`

	// The number of healthy health checks required before a host is marked healthy. Note that during startup,
	// only a single successful health check is required to mark a host healthy.
	HealthyThreshold *int `mandatory:"false" json:"healthyThreshold"`

	// Array of base64 encoded strings should be found in the returning message to be considered as healthy.
	// When checking the response, “fuzzy” matching is performed such that each payload block must be found,
	// and in the order specified, but not necessarily contiguous.
	Receive [][]byte `mandatory:"false" json:"receive"`
}

// GetTimeoutInMs returns TimeoutInMs
func (m TcpHealthCheck) GetTimeoutInMs() *int64 {
	return m.TimeoutInMs
}

// GetIntervalInMs returns IntervalInMs
func (m TcpHealthCheck) GetIntervalInMs() *int64 {
	return m.IntervalInMs
}

// GetUnhealthyThreshold returns UnhealthyThreshold
func (m TcpHealthCheck) GetUnhealthyThreshold() *int {
	return m.UnhealthyThreshold
}

// GetHealthyThreshold returns HealthyThreshold
func (m TcpHealthCheck) GetHealthyThreshold() *int {
	return m.HealthyThreshold
}

func (m TcpHealthCheck) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TcpHealthCheck) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TcpHealthCheck) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTcpHealthCheck TcpHealthCheck
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeTcpHealthCheck
	}{
		"TCP",
		(MarshalTypeTcpHealthCheck)(m),
	}

	return json.Marshal(&s)
}
