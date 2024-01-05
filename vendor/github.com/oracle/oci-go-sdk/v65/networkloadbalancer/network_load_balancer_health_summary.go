// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// This describes the network load balancer API.
//

package networkloadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NetworkLoadBalancerHealthSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNetworkLoadBalancerHealthSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetNetworkLoadBalancerHealthSummaryStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingNetworkLoadBalancerHealthSummaryStatusEnum = map[string]NetworkLoadBalancerHealthSummaryStatusEnum{
	"OK":       NetworkLoadBalancerHealthSummaryStatusOk,
	"WARNING":  NetworkLoadBalancerHealthSummaryStatusWarning,
	"CRITICAL": NetworkLoadBalancerHealthSummaryStatusCritical,
	"UNKNOWN":  NetworkLoadBalancerHealthSummaryStatusUnknown,
}

var mappingNetworkLoadBalancerHealthSummaryStatusEnumLowerCase = map[string]NetworkLoadBalancerHealthSummaryStatusEnum{
	"ok":       NetworkLoadBalancerHealthSummaryStatusOk,
	"warning":  NetworkLoadBalancerHealthSummaryStatusWarning,
	"critical": NetworkLoadBalancerHealthSummaryStatusCritical,
	"unknown":  NetworkLoadBalancerHealthSummaryStatusUnknown,
}

// GetNetworkLoadBalancerHealthSummaryStatusEnumValues Enumerates the set of values for NetworkLoadBalancerHealthSummaryStatusEnum
func GetNetworkLoadBalancerHealthSummaryStatusEnumValues() []NetworkLoadBalancerHealthSummaryStatusEnum {
	values := make([]NetworkLoadBalancerHealthSummaryStatusEnum, 0)
	for _, v := range mappingNetworkLoadBalancerHealthSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkLoadBalancerHealthSummaryStatusEnumStringValues Enumerates the set of values in String for NetworkLoadBalancerHealthSummaryStatusEnum
func GetNetworkLoadBalancerHealthSummaryStatusEnumStringValues() []string {
	return []string{
		"OK",
		"WARNING",
		"CRITICAL",
		"UNKNOWN",
	}
}

// GetMappingNetworkLoadBalancerHealthSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkLoadBalancerHealthSummaryStatusEnum(val string) (NetworkLoadBalancerHealthSummaryStatusEnum, bool) {
	enum, ok := mappingNetworkLoadBalancerHealthSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
