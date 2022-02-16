// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RoutingPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRoutingPolicyConditionLanguageVersionEnum(string(m.ConditionLanguageVersion)); !ok && m.ConditionLanguageVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConditionLanguageVersion: %s. Supported values are: %s.", m.ConditionLanguageVersion, strings.Join(GetRoutingPolicyConditionLanguageVersionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RoutingPolicyConditionLanguageVersionEnum Enum with underlying type: string
type RoutingPolicyConditionLanguageVersionEnum string

// Set of constants representing the allowable values for RoutingPolicyConditionLanguageVersionEnum
const (
	RoutingPolicyConditionLanguageVersionV1 RoutingPolicyConditionLanguageVersionEnum = "V1"
)

var mappingRoutingPolicyConditionLanguageVersionEnum = map[string]RoutingPolicyConditionLanguageVersionEnum{
	"V1": RoutingPolicyConditionLanguageVersionV1,
}

// GetRoutingPolicyConditionLanguageVersionEnumValues Enumerates the set of values for RoutingPolicyConditionLanguageVersionEnum
func GetRoutingPolicyConditionLanguageVersionEnumValues() []RoutingPolicyConditionLanguageVersionEnum {
	values := make([]RoutingPolicyConditionLanguageVersionEnum, 0)
	for _, v := range mappingRoutingPolicyConditionLanguageVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetRoutingPolicyConditionLanguageVersionEnumStringValues Enumerates the set of values in String for RoutingPolicyConditionLanguageVersionEnum
func GetRoutingPolicyConditionLanguageVersionEnumStringValues() []string {
	return []string{
		"V1",
	}
}

// GetMappingRoutingPolicyConditionLanguageVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRoutingPolicyConditionLanguageVersionEnum(val string) (RoutingPolicyConditionLanguageVersionEnum, bool) {
	mappingRoutingPolicyConditionLanguageVersionEnumIgnoreCase := make(map[string]RoutingPolicyConditionLanguageVersionEnum)
	for k, v := range mappingRoutingPolicyConditionLanguageVersionEnum {
		mappingRoutingPolicyConditionLanguageVersionEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingRoutingPolicyConditionLanguageVersionEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
