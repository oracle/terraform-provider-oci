// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LoadBalancerConfig This config is used to choose the load balancer service and bandwidth for OpenSearch and OpenDashboard load balancers.
type LoadBalancerConfig struct {

	// Load balancer service for OpenSearch and OpenDashboard load balancer. Default value is LOAD_BALANCER.
	LoadBalancerServiceType LoadBalancerConfigLoadBalancerServiceTypeEnum `mandatory:"true" json:"loadBalancerServiceType"`

	// Minimum bandwidth (Mbps) of OpenSearch load balancer. Not applicable for network load balancer service.
	LoadBalancerMinBandwidthInMbps *int `mandatory:"false" json:"loadBalancerMinBandwidthInMbps"`

	// Maximum bandwidth (Mbps) of OpenSearch load balancer. Not applicable for network load balancer service.
	LoadBalancerMaxBandwidthInMbps *int `mandatory:"false" json:"loadBalancerMaxBandwidthInMbps"`
}

func (m LoadBalancerConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LoadBalancerConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLoadBalancerConfigLoadBalancerServiceTypeEnum(string(m.LoadBalancerServiceType)); !ok && m.LoadBalancerServiceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LoadBalancerServiceType: %s. Supported values are: %s.", m.LoadBalancerServiceType, strings.Join(GetLoadBalancerConfigLoadBalancerServiceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LoadBalancerConfigLoadBalancerServiceTypeEnum Enum with underlying type: string
type LoadBalancerConfigLoadBalancerServiceTypeEnum string

// Set of constants representing the allowable values for LoadBalancerConfigLoadBalancerServiceTypeEnum
const (
	LoadBalancerConfigLoadBalancerServiceTypeLoadBalancer        LoadBalancerConfigLoadBalancerServiceTypeEnum = "LOAD_BALANCER"
	LoadBalancerConfigLoadBalancerServiceTypeNetworkLoadBalancer LoadBalancerConfigLoadBalancerServiceTypeEnum = "NETWORK_LOAD_BALANCER"
)

var mappingLoadBalancerConfigLoadBalancerServiceTypeEnum = map[string]LoadBalancerConfigLoadBalancerServiceTypeEnum{
	"LOAD_BALANCER":         LoadBalancerConfigLoadBalancerServiceTypeLoadBalancer,
	"NETWORK_LOAD_BALANCER": LoadBalancerConfigLoadBalancerServiceTypeNetworkLoadBalancer,
}

var mappingLoadBalancerConfigLoadBalancerServiceTypeEnumLowerCase = map[string]LoadBalancerConfigLoadBalancerServiceTypeEnum{
	"load_balancer":         LoadBalancerConfigLoadBalancerServiceTypeLoadBalancer,
	"network_load_balancer": LoadBalancerConfigLoadBalancerServiceTypeNetworkLoadBalancer,
}

// GetLoadBalancerConfigLoadBalancerServiceTypeEnumValues Enumerates the set of values for LoadBalancerConfigLoadBalancerServiceTypeEnum
func GetLoadBalancerConfigLoadBalancerServiceTypeEnumValues() []LoadBalancerConfigLoadBalancerServiceTypeEnum {
	values := make([]LoadBalancerConfigLoadBalancerServiceTypeEnum, 0)
	for _, v := range mappingLoadBalancerConfigLoadBalancerServiceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLoadBalancerConfigLoadBalancerServiceTypeEnumStringValues Enumerates the set of values in String for LoadBalancerConfigLoadBalancerServiceTypeEnum
func GetLoadBalancerConfigLoadBalancerServiceTypeEnumStringValues() []string {
	return []string{
		"LOAD_BALANCER",
		"NETWORK_LOAD_BALANCER",
	}
}

// GetMappingLoadBalancerConfigLoadBalancerServiceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLoadBalancerConfigLoadBalancerServiceTypeEnum(val string) (LoadBalancerConfigLoadBalancerServiceTypeEnum, bool) {
	enum, ok := mappingLoadBalancerConfigLoadBalancerServiceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
