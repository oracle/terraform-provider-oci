// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LoadBalancerHealthSummary A health status summary for the specified load balancer.
type LoadBalancerHealthSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer the health status is associated with.
	LoadBalancerId *string `mandatory:"true" json:"loadBalancerId"`

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
	Status LoadBalancerHealthSummaryStatusEnum `mandatory:"true" json:"status"`
}

func (m LoadBalancerHealthSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LoadBalancerHealthSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLoadBalancerHealthSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetLoadBalancerHealthSummaryStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LoadBalancerHealthSummaryStatusEnum Enum with underlying type: string
type LoadBalancerHealthSummaryStatusEnum string

// Set of constants representing the allowable values for LoadBalancerHealthSummaryStatusEnum
const (
	LoadBalancerHealthSummaryStatusOk       LoadBalancerHealthSummaryStatusEnum = "OK"
	LoadBalancerHealthSummaryStatusWarning  LoadBalancerHealthSummaryStatusEnum = "WARNING"
	LoadBalancerHealthSummaryStatusCritical LoadBalancerHealthSummaryStatusEnum = "CRITICAL"
	LoadBalancerHealthSummaryStatusUnknown  LoadBalancerHealthSummaryStatusEnum = "UNKNOWN"
)

var mappingLoadBalancerHealthSummaryStatusEnum = map[string]LoadBalancerHealthSummaryStatusEnum{
	"OK":       LoadBalancerHealthSummaryStatusOk,
	"WARNING":  LoadBalancerHealthSummaryStatusWarning,
	"CRITICAL": LoadBalancerHealthSummaryStatusCritical,
	"UNKNOWN":  LoadBalancerHealthSummaryStatusUnknown,
}

var mappingLoadBalancerHealthSummaryStatusEnumLowerCase = map[string]LoadBalancerHealthSummaryStatusEnum{
	"ok":       LoadBalancerHealthSummaryStatusOk,
	"warning":  LoadBalancerHealthSummaryStatusWarning,
	"critical": LoadBalancerHealthSummaryStatusCritical,
	"unknown":  LoadBalancerHealthSummaryStatusUnknown,
}

// GetLoadBalancerHealthSummaryStatusEnumValues Enumerates the set of values for LoadBalancerHealthSummaryStatusEnum
func GetLoadBalancerHealthSummaryStatusEnumValues() []LoadBalancerHealthSummaryStatusEnum {
	values := make([]LoadBalancerHealthSummaryStatusEnum, 0)
	for _, v := range mappingLoadBalancerHealthSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetLoadBalancerHealthSummaryStatusEnumStringValues Enumerates the set of values in String for LoadBalancerHealthSummaryStatusEnum
func GetLoadBalancerHealthSummaryStatusEnumStringValues() []string {
	return []string{
		"OK",
		"WARNING",
		"CRITICAL",
		"UNKNOWN",
	}
}

// GetMappingLoadBalancerHealthSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLoadBalancerHealthSummaryStatusEnum(val string) (LoadBalancerHealthSummaryStatusEnum, bool) {
	enum, ok := mappingLoadBalancerHealthSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
