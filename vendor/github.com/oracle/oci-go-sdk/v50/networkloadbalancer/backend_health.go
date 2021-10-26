// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// A description of the network load balancer API
//

package networkloadbalancer

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// BackendHealth The health status of the specified backend server.
type BackendHealth struct {

	// The general health status of the specified backend server.
	// *   **OK:**  All health check probes return `OK`
	// *   **WARNING:** At least one of the health check probes does not return `OK`
	// *   **CRITICAL:** None of the health check probes return `OK`.
	// *
	// *   **UNKNOWN:** One of the health checks probes return `UNKNOWN`,
	// *   or the system is unable to retrieve metrics at this time.
	Status BackendHealthStatusEnum `mandatory:"true" json:"status"`

	// A list of the most recent health check results returned for the specified backend server.
	HealthCheckResults []HealthCheckResult `mandatory:"true" json:"healthCheckResults"`
}

func (m BackendHealth) String() string {
	return common.PointerString(m)
}

// BackendHealthStatusEnum Enum with underlying type: string
type BackendHealthStatusEnum string

// Set of constants representing the allowable values for BackendHealthStatusEnum
const (
	BackendHealthStatusOk       BackendHealthStatusEnum = "OK"
	BackendHealthStatusWarning  BackendHealthStatusEnum = "WARNING"
	BackendHealthStatusCritical BackendHealthStatusEnum = "CRITICAL"
	BackendHealthStatusUnknown  BackendHealthStatusEnum = "UNKNOWN"
)

var mappingBackendHealthStatus = map[string]BackendHealthStatusEnum{
	"OK":       BackendHealthStatusOk,
	"WARNING":  BackendHealthStatusWarning,
	"CRITICAL": BackendHealthStatusCritical,
	"UNKNOWN":  BackendHealthStatusUnknown,
}

// GetBackendHealthStatusEnumValues Enumerates the set of values for BackendHealthStatusEnum
func GetBackendHealthStatusEnumValues() []BackendHealthStatusEnum {
	values := make([]BackendHealthStatusEnum, 0)
	for _, v := range mappingBackendHealthStatus {
		values = append(values, v)
	}
	return values
}
