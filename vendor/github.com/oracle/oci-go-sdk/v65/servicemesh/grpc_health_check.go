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

// GrpcHealthCheck Grpc based health check.
type GrpcHealthCheck struct {

	// The time to wait for a health check response. If the timeout is reached the health check attempt
	// will be considered a failure.
	TimeoutInMs *int64 `mandatory:"true" json:"timeoutInMs"`

	// The interval between health checks.
	IntervalInMs *int64 `mandatory:"true" json:"intervalInMs"`

	// The grpc service name to request for health check. Use empty string to check the health of the system.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// The number of unhealthy health checks required before a host is marked unhealthy.
	UnhealthyThreshold *int `mandatory:"false" json:"unhealthyThreshold"`

	// The number of healthy health checks required before a host is marked healthy. Note that during startup,
	// only a single successful health check is required to mark a host healthy.
	HealthyThreshold *int `mandatory:"false" json:"healthyThreshold"`

	// The value of the :authority header in the gRPC health check request.
	Authority *string `mandatory:"false" json:"authority"`

	// List of key-value pairs to be added to the metadata in the health check request.
	Metadata []HttpHeader `mandatory:"false" json:"metadata"`
}

// GetTimeoutInMs returns TimeoutInMs
func (m GrpcHealthCheck) GetTimeoutInMs() *int64 {
	return m.TimeoutInMs
}

// GetIntervalInMs returns IntervalInMs
func (m GrpcHealthCheck) GetIntervalInMs() *int64 {
	return m.IntervalInMs
}

// GetUnhealthyThreshold returns UnhealthyThreshold
func (m GrpcHealthCheck) GetUnhealthyThreshold() *int {
	return m.UnhealthyThreshold
}

// GetHealthyThreshold returns HealthyThreshold
func (m GrpcHealthCheck) GetHealthyThreshold() *int {
	return m.HealthyThreshold
}

func (m GrpcHealthCheck) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GrpcHealthCheck) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GrpcHealthCheck) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGrpcHealthCheck GrpcHealthCheck
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeGrpcHealthCheck
	}{
		"GRPC",
		(MarshalTypeGrpcHealthCheck)(m),
	}

	return json.Marshal(&s)
}
