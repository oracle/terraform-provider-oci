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

// CreateRoutingPolicyDetails An ordered list of routing rules.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateRoutingPolicyDetails struct {

	// The name for this list of routing rules. It must be unique and it cannot be changed. Avoid entering
	// confidential information.
	// Example: `example_routing_rules`
	Name *string `mandatory:"true" json:"name"`

	// The version of the language in which `condition` of `rules` are composed.
	ConditionLanguageVersion CreateRoutingPolicyDetailsConditionLanguageVersionEnum `mandatory:"true" json:"conditionLanguageVersion"`

	// The list of routing rules.
	Rules []RoutingRule `mandatory:"true" json:"rules"`
}

func (m CreateRoutingPolicyDetails) String() string {
	return common.PointerString(m)
}

// CreateRoutingPolicyDetailsConditionLanguageVersionEnum Enum with underlying type: string
type CreateRoutingPolicyDetailsConditionLanguageVersionEnum string

// Set of constants representing the allowable values for CreateRoutingPolicyDetailsConditionLanguageVersionEnum
const (
	CreateRoutingPolicyDetailsConditionLanguageVersionV1 CreateRoutingPolicyDetailsConditionLanguageVersionEnum = "V1"
)

var mappingCreateRoutingPolicyDetailsConditionLanguageVersion = map[string]CreateRoutingPolicyDetailsConditionLanguageVersionEnum{
	"V1": CreateRoutingPolicyDetailsConditionLanguageVersionV1,
}

// GetCreateRoutingPolicyDetailsConditionLanguageVersionEnumValues Enumerates the set of values for CreateRoutingPolicyDetailsConditionLanguageVersionEnum
func GetCreateRoutingPolicyDetailsConditionLanguageVersionEnumValues() []CreateRoutingPolicyDetailsConditionLanguageVersionEnum {
	values := make([]CreateRoutingPolicyDetailsConditionLanguageVersionEnum, 0)
	for _, v := range mappingCreateRoutingPolicyDetailsConditionLanguageVersion {
		values = append(values, v)
	}
	return values
}
