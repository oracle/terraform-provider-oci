// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// HttpHealthCheck Http based health check.
type HttpHealthCheck struct {

	// The time to wait for a health check response. If the timeout is reached the health check attempt
	// will be considered a failure.
	TimeoutInMs *int64 `mandatory:"true" json:"timeoutInMs"`

	// The interval between health checks.
	IntervalInMs *int64 `mandatory:"true" json:"intervalInMs"`

	// The HTTP path to be requested during health check. E.g. /healthcheck.
	Path *string `mandatory:"true" json:"path"`

	// The number of unhealthy health checks required before a host is marked unhealthy.
	UnhealthyThreshold *int `mandatory:"false" json:"unhealthyThreshold"`

	// The number of healthy health checks required before a host is marked healthy. Note that during startup,
	// only a single successful health check is required to mark a host healthy.
	HealthyThreshold *int `mandatory:"false" json:"healthyThreshold"`

	// List of HTTP response statuses considered healthy. If provided, replaces the default 200-only policy.
	ExpectedStatuses []StatusCodeRange `mandatory:"false" json:"expectedStatuses"`

	// List of HTTP request headers to append to the health check request.
	RequestHeadersToAdd []HttpHeader `mandatory:"false" json:"requestHeadersToAdd"`
}

//GetTimeoutInMs returns TimeoutInMs
func (m HttpHealthCheck) GetTimeoutInMs() *int64 {
	return m.TimeoutInMs
}

//GetIntervalInMs returns IntervalInMs
func (m HttpHealthCheck) GetIntervalInMs() *int64 {
	return m.IntervalInMs
}

//GetUnhealthyThreshold returns UnhealthyThreshold
func (m HttpHealthCheck) GetUnhealthyThreshold() *int {
	return m.UnhealthyThreshold
}

//GetHealthyThreshold returns HealthyThreshold
func (m HttpHealthCheck) GetHealthyThreshold() *int {
	return m.HealthyThreshold
}

func (m HttpHealthCheck) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HttpHealthCheck) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HttpHealthCheck) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHttpHealthCheck HttpHealthCheck
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeHttpHealthCheck
	}{
		"HTTP",
		(MarshalTypeHttpHealthCheck)(m),
	}

	return json.Marshal(&s)
}
