// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LoadBalancerHealth The health status details for the specified load balancer.
// This object does not explicitly enumerate backend sets with a status of `OK`. However, they are included in the
// `totalBackendSetCount` sum.
type LoadBalancerHealth struct {

	// The overall health status of the load balancer.
	// *  **OK:** All backend sets associated with the load balancer return a status of `OK`.
	// *  **WARNING:** At least one of the backend sets associated with the load balancer returns a status of `WARNING`,
	// no backend sets return a status of `CRITICAL`, and the load balancer life cycle state is `ACTIVE`.
	// *  **CRITICAL:** One or more of the backend sets associated with the load balancer return a status of `CRITICAL`.
	// *  **UNKNOWN:** If any one of the following conditions is true:
	//     *  The load balancer life cycle state is not `ACTIVE`.
	//     *  No backend sets are defined for the load balancer.
	//     *  More than half of the backend sets associated with the load balancer return a status of `UNKNOWN`, none of the backend
	//        sets return a status of `WARNING` or `CRITICAL`, and the load balancer life cycle state is `ACTIVE`.
	//     *  The system could not retrieve metrics for any reason.
	Status LoadBalancerHealthStatusEnum `mandatory:"true" json:"status"`

	// A list of backend sets that are currently in the `WARNING` health state. The list identifies each backend set by the
	// friendly name you assigned when you created it.
	// Example: `example_backend_set3`
	WarningStateBackendSetNames []string `mandatory:"true" json:"warningStateBackendSetNames"`

	// A list of backend sets that are currently in the `CRITICAL` health state. The list identifies each backend set by the
	// friendly name you assigned when you created it.
	// Example: `example_backend_set`
	CriticalStateBackendSetNames []string `mandatory:"true" json:"criticalStateBackendSetNames"`

	// A list of backend sets that are currently in the `UNKNOWN` health state. The list identifies each backend set by the
	// friendly name you assigned when you created it.
	// Example: `example_backend_set2`
	UnknownStateBackendSetNames []string `mandatory:"true" json:"unknownStateBackendSetNames"`

	// The total number of backend sets associated with this load balancer.
	// Example: `4`
	TotalBackendSetCount *int `mandatory:"true" json:"totalBackendSetCount"`
}

func (m LoadBalancerHealth) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LoadBalancerHealth) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLoadBalancerHealthStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetLoadBalancerHealthStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LoadBalancerHealthStatusEnum Enum with underlying type: string
type LoadBalancerHealthStatusEnum string

// Set of constants representing the allowable values for LoadBalancerHealthStatusEnum
const (
	LoadBalancerHealthStatusOk       LoadBalancerHealthStatusEnum = "OK"
	LoadBalancerHealthStatusWarning  LoadBalancerHealthStatusEnum = "WARNING"
	LoadBalancerHealthStatusCritical LoadBalancerHealthStatusEnum = "CRITICAL"
	LoadBalancerHealthStatusUnknown  LoadBalancerHealthStatusEnum = "UNKNOWN"
)

var mappingLoadBalancerHealthStatusEnum = map[string]LoadBalancerHealthStatusEnum{
	"OK":       LoadBalancerHealthStatusOk,
	"WARNING":  LoadBalancerHealthStatusWarning,
	"CRITICAL": LoadBalancerHealthStatusCritical,
	"UNKNOWN":  LoadBalancerHealthStatusUnknown,
}

var mappingLoadBalancerHealthStatusEnumLowerCase = map[string]LoadBalancerHealthStatusEnum{
	"ok":       LoadBalancerHealthStatusOk,
	"warning":  LoadBalancerHealthStatusWarning,
	"critical": LoadBalancerHealthStatusCritical,
	"unknown":  LoadBalancerHealthStatusUnknown,
}

// GetLoadBalancerHealthStatusEnumValues Enumerates the set of values for LoadBalancerHealthStatusEnum
func GetLoadBalancerHealthStatusEnumValues() []LoadBalancerHealthStatusEnum {
	values := make([]LoadBalancerHealthStatusEnum, 0)
	for _, v := range mappingLoadBalancerHealthStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetLoadBalancerHealthStatusEnumStringValues Enumerates the set of values in String for LoadBalancerHealthStatusEnum
func GetLoadBalancerHealthStatusEnumStringValues() []string {
	return []string{
		"OK",
		"WARNING",
		"CRITICAL",
		"UNKNOWN",
	}
}

// GetMappingLoadBalancerHealthStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLoadBalancerHealthStatusEnum(val string) (LoadBalancerHealthStatusEnum, bool) {
	enum, ok := mappingLoadBalancerHealthStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
