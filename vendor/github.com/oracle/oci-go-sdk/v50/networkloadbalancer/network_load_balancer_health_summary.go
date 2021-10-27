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

// NetworkLoadBalancerHealthSummary A health status summary for the specified network load balancer
type NetworkLoadBalancerHealthSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network load balancer with which the health status is associated.
	NetworkLoadBalancerId *string `mandatory:"true" json:"networkLoadBalancerId"`

	// The overall health status of the network load balancer.
	// *  **OK:** All backend sets associated with the network load balancer return a status of `OK`.
	// *  **WARNING:** At least one of the backend sets associated with the network load balancer returns a status of `WARNING`,
	// no backend sets return a status of `CRITICAL`, and the network load balancer life cycle state is `ACTIVE`.
	// *  **CRITICAL:** One or more of the backend sets associated with the network load balancer returns a status of `CRITICAL`.
	// *  **UNKNOWN:** If any one of the following conditions is true:
	//     *  The network load balancer life cycle state is not `ACTIVE`.
	//     *  No backend sets are defined for the network load balancer.
	//     *  More than half of the backend sets associated with the network load balancer return a status of `UNKNOWN`, none of the backend
	//        sets returns a status of `WARNING` or `CRITICAL`, and the network load balancer life cycle state is `ACTIVE`.
	//     *  The system could not retrieve metrics for any reason.
	Status NetworkLoadBalancerHealthSummaryStatusEnum `mandatory:"true" json:"status"`
}

func (m NetworkLoadBalancerHealthSummary) String() string {
	return common.PointerString(m)
}

// NetworkLoadBalancerHealthSummaryStatusEnum Enum with underlying type: string
type NetworkLoadBalancerHealthSummaryStatusEnum string

// Set of constants representing the allowable values for NetworkLoadBalancerHealthSummaryStatusEnum
const (
	NetworkLoadBalancerHealthSummaryStatusOk       NetworkLoadBalancerHealthSummaryStatusEnum = "OK"
	NetworkLoadBalancerHealthSummaryStatusWarning  NetworkLoadBalancerHealthSummaryStatusEnum = "WARNING"
	NetworkLoadBalancerHealthSummaryStatusCritical NetworkLoadBalancerHealthSummaryStatusEnum = "CRITICAL"
	NetworkLoadBalancerHealthSummaryStatusUnknown  NetworkLoadBalancerHealthSummaryStatusEnum = "UNKNOWN"
)

var mappingNetworkLoadBalancerHealthSummaryStatus = map[string]NetworkLoadBalancerHealthSummaryStatusEnum{
	"OK":       NetworkLoadBalancerHealthSummaryStatusOk,
	"WARNING":  NetworkLoadBalancerHealthSummaryStatusWarning,
	"CRITICAL": NetworkLoadBalancerHealthSummaryStatusCritical,
	"UNKNOWN":  NetworkLoadBalancerHealthSummaryStatusUnknown,
}

// GetNetworkLoadBalancerHealthSummaryStatusEnumValues Enumerates the set of values for NetworkLoadBalancerHealthSummaryStatusEnum
func GetNetworkLoadBalancerHealthSummaryStatusEnumValues() []NetworkLoadBalancerHealthSummaryStatusEnum {
	values := make([]NetworkLoadBalancerHealthSummaryStatusEnum, 0)
	for _, v := range mappingNetworkLoadBalancerHealthSummaryStatus {
		values = append(values, v)
	}
	return values
}
