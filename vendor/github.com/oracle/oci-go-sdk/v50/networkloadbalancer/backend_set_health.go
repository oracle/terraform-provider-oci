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

// BackendSetHealth The health status details for a backend set.
// This object does not explicitly enumerate backend servers with a status of `OK`. However, the backend sets are included in the
// `totalBackendCount` sum.
type BackendSetHealth struct {

	// Overall health status of the backend set.
	// *  **OK:** All backend servers in the backend set return a status of `OK`.
	// *  **WARNING:** Half or more of the backend servers in a backend set return a status of `OK` and at least one backend
	// server returns a status of `WARNING`, `CRITICAL`, or `UNKNOWN`.
	// *  **CRITICAL:** Fewer than half of the backend servers in a backend set return a status of `OK`.
	// *  **UNKNOWN:** If no probes have yet been sent to the backends, or the system is
	// unable to retrieve metrics from the backends.
	Status BackendSetHealthStatusEnum `mandatory:"true" json:"status"`

	// A list of backend servers that are currently in the `WARNING` health state. The list identifies each backend server by
	// IP address or OCID and port.
	// Example: `10.0.0.3:8080` or `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>:8080`
	WarningStateBackendNames []string `mandatory:"true" json:"warningStateBackendNames"`

	// A list of backend servers that are currently in the `CRITICAL` health state. The list identifies each backend server by
	// IP address and port.
	// Example: `10.0.0.4:8080`
	CriticalStateBackendNames []string `mandatory:"true" json:"criticalStateBackendNames"`

	// A list of backend servers that are currently in the `UNKNOWN` health state. The list identifies each backend server by
	// IP address and port.
	// Example: `10.0.0.5:8080`
	UnknownStateBackendNames []string `mandatory:"true" json:"unknownStateBackendNames"`

	// The total number of backend servers in this backend set.
	// Example: `7`
	TotalBackendCount *int `mandatory:"true" json:"totalBackendCount"`
}

func (m BackendSetHealth) String() string {
	return common.PointerString(m)
}

// BackendSetHealthStatusEnum Enum with underlying type: string
type BackendSetHealthStatusEnum string

// Set of constants representing the allowable values for BackendSetHealthStatusEnum
const (
	BackendSetHealthStatusOk       BackendSetHealthStatusEnum = "OK"
	BackendSetHealthStatusWarning  BackendSetHealthStatusEnum = "WARNING"
	BackendSetHealthStatusCritical BackendSetHealthStatusEnum = "CRITICAL"
	BackendSetHealthStatusUnknown  BackendSetHealthStatusEnum = "UNKNOWN"
)

var mappingBackendSetHealthStatus = map[string]BackendSetHealthStatusEnum{
	"OK":       BackendSetHealthStatusOk,
	"WARNING":  BackendSetHealthStatusWarning,
	"CRITICAL": BackendSetHealthStatusCritical,
	"UNKNOWN":  BackendSetHealthStatusUnknown,
}

// GetBackendSetHealthStatusEnumValues Enumerates the set of values for BackendSetHealthStatusEnum
func GetBackendSetHealthStatusEnumValues() []BackendSetHealthStatusEnum {
	values := make([]BackendSetHealthStatusEnum, 0)
	for _, v := range mappingBackendSetHealthStatus {
		values = append(values, v)
	}
	return values
}
