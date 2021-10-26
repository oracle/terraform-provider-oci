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

// UpdateRoutingPolicyDetails An updated list of routing rules that overwrites the existing list of routing rules.
type UpdateRoutingPolicyDetails struct {

	// The list of routing rules.
	Rules []RoutingRule `mandatory:"true" json:"rules"`

	// The version of the language in which `condition` of `rules` are composed.
	ConditionLanguageVersion UpdateRoutingPolicyDetailsConditionLanguageVersionEnum `mandatory:"false" json:"conditionLanguageVersion,omitempty"`
}

func (m UpdateRoutingPolicyDetails) String() string {
	return common.PointerString(m)
}

// UpdateRoutingPolicyDetailsConditionLanguageVersionEnum Enum with underlying type: string
type UpdateRoutingPolicyDetailsConditionLanguageVersionEnum string

// Set of constants representing the allowable values for UpdateRoutingPolicyDetailsConditionLanguageVersionEnum
const (
	UpdateRoutingPolicyDetailsConditionLanguageVersionV1 UpdateRoutingPolicyDetailsConditionLanguageVersionEnum = "V1"
)

var mappingUpdateRoutingPolicyDetailsConditionLanguageVersion = map[string]UpdateRoutingPolicyDetailsConditionLanguageVersionEnum{
	"V1": UpdateRoutingPolicyDetailsConditionLanguageVersionV1,
}

// GetUpdateRoutingPolicyDetailsConditionLanguageVersionEnumValues Enumerates the set of values for UpdateRoutingPolicyDetailsConditionLanguageVersionEnum
func GetUpdateRoutingPolicyDetailsConditionLanguageVersionEnumValues() []UpdateRoutingPolicyDetailsConditionLanguageVersionEnum {
	values := make([]UpdateRoutingPolicyDetailsConditionLanguageVersionEnum, 0)
	for _, v := range mappingUpdateRoutingPolicyDetailsConditionLanguageVersion {
		values = append(values, v)
	}
	return values
}
