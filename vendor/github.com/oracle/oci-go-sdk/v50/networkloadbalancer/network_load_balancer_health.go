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

// NetworkLoadBalancerHealth The health status details for the specified network load balancer.
// This object does not explicitly enumerate backend sets with a status of `OK`. However, the backend sets are included in the
// `totalBackendSetCount` sum.
type NetworkLoadBalancerHealth struct {

	// The overall health status of the network load balancer.
	// *  **OK:** All backend sets associated with the network load balancer return a status of `OK`.
	// *  **WARNING:** At least one of the backend sets associated with the network load balancer returns a status of `WARNING`,
	// no backend sets return a status of `CRITICAL`, and the network load balancer life cycle state is `ACTIVE`.
	// *  **CRITICAL:** One or more of the backend sets associated with the network load balancer return a status of `CRITICAL`.
	// *  **UNKNOWN:** If any one of the following conditions is true:
	//     *  The network load balancer life cycle state is not `ACTIVE`.
	//     *  No backend sets are defined for the network load balancer.
	//     *  More than half of the backend sets associated with the network load balancer return a status of `UNKNOWN`, none of the backend
	//        sets return a status of `WARNING` or `CRITICAL`, and the network load balancer life cycle state is `ACTIVE`.
	//     *  The system could not retrieve metrics for any reason.
	Status NetworkLoadBalancerHealthStatusEnum `mandatory:"true" json:"status"`

	// A list of backend sets that are currently in the `WARNING` health state. The list identifies each backend set by the
	// user-friendly name you assigned when you created the backend set.
	// Example: `example_backend_set3`
	WarningStateBackendSetNames []string `mandatory:"true" json:"warningStateBackendSetNames"`

	// A list of backend sets that are currently in the `CRITICAL` health state. The list identifies each backend set by the
	// user-friendly name you assigned when you created the backend set.
	// Example: `example_backend_set`
	CriticalStateBackendSetNames []string `mandatory:"true" json:"criticalStateBackendSetNames"`

	// A list of backend sets that are currently in the `UNKNOWN` health state. The list identifies each backend set by the
	// user-friendly name you assigned when you created the backend set.
	// Example: `example_backend_set2`
	UnknownStateBackendSetNames []string `mandatory:"true" json:"unknownStateBackendSetNames"`

	// The total number of backend sets associated with this network load balancer.
	// Example: `4`
	TotalBackendSetCount *int `mandatory:"true" json:"totalBackendSetCount"`
}

func (m NetworkLoadBalancerHealth) String() string {
	return common.PointerString(m)
}

// NetworkLoadBalancerHealthStatusEnum Enum with underlying type: string
type NetworkLoadBalancerHealthStatusEnum string

// Set of constants representing the allowable values for NetworkLoadBalancerHealthStatusEnum
const (
	NetworkLoadBalancerHealthStatusOk       NetworkLoadBalancerHealthStatusEnum = "OK"
	NetworkLoadBalancerHealthStatusWarning  NetworkLoadBalancerHealthStatusEnum = "WARNING"
	NetworkLoadBalancerHealthStatusCritical NetworkLoadBalancerHealthStatusEnum = "CRITICAL"
	NetworkLoadBalancerHealthStatusUnknown  NetworkLoadBalancerHealthStatusEnum = "UNKNOWN"
)

var mappingNetworkLoadBalancerHealthStatus = map[string]NetworkLoadBalancerHealthStatusEnum{
	"OK":       NetworkLoadBalancerHealthStatusOk,
	"WARNING":  NetworkLoadBalancerHealthStatusWarning,
	"CRITICAL": NetworkLoadBalancerHealthStatusCritical,
	"UNKNOWN":  NetworkLoadBalancerHealthStatusUnknown,
}

// GetNetworkLoadBalancerHealthStatusEnumValues Enumerates the set of values for NetworkLoadBalancerHealthStatusEnum
func GetNetworkLoadBalancerHealthStatusEnumValues() []NetworkLoadBalancerHealthStatusEnum {
	values := make([]NetworkLoadBalancerHealthStatusEnum, 0)
	for _, v := range mappingNetworkLoadBalancerHealthStatus {
		values = append(values, v)
	}
	return values
}
