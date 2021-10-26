// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// RoutingPolicy A named ordered list of routing rules that is applied to a listener.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type RoutingPolicy struct {

	// The unique name for this list of routing rules. Avoid entering confidential information.
	// Example: `example_routing_policy`
	Name *string `mandatory:"true" json:"name"`

	// The version of the language in which `condition` of `rules` are composed.
	ConditionLanguageVersion RoutingPolicyConditionLanguageVersionEnum `mandatory:"true" json:"conditionLanguageVersion"`

	// The ordered list of routing rules.
	Rules []RoutingRule `mandatory:"true" json:"rules"`
}

func (m RoutingPolicy) String() string {
	return common.PointerString(m)
}

// RoutingPolicyConditionLanguageVersionEnum Enum with underlying type: string
type RoutingPolicyConditionLanguageVersionEnum string

// Set of constants representing the allowable values for RoutingPolicyConditionLanguageVersionEnum
const (
	RoutingPolicyConditionLanguageVersionV1 RoutingPolicyConditionLanguageVersionEnum = "V1"
)

var mappingRoutingPolicyConditionLanguageVersion = map[string]RoutingPolicyConditionLanguageVersionEnum{
	"V1": RoutingPolicyConditionLanguageVersionV1,
}

// GetRoutingPolicyConditionLanguageVersionEnumValues Enumerates the set of values for RoutingPolicyConditionLanguageVersionEnum
func GetRoutingPolicyConditionLanguageVersionEnumValues() []RoutingPolicyConditionLanguageVersionEnum {
	values := make([]RoutingPolicyConditionLanguageVersionEnum, 0)
	for _, v := range mappingRoutingPolicyConditionLanguageVersion {
		values = append(values, v)
	}
	return values
}
