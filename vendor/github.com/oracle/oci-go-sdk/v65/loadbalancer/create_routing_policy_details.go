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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateRoutingPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateRoutingPolicyDetailsConditionLanguageVersionEnum(string(m.ConditionLanguageVersion)); !ok && m.ConditionLanguageVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConditionLanguageVersion: %s. Supported values are: %s.", m.ConditionLanguageVersion, strings.Join(GetCreateRoutingPolicyDetailsConditionLanguageVersionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateRoutingPolicyDetailsConditionLanguageVersionEnum Enum with underlying type: string
type CreateRoutingPolicyDetailsConditionLanguageVersionEnum string

// Set of constants representing the allowable values for CreateRoutingPolicyDetailsConditionLanguageVersionEnum
const (
	CreateRoutingPolicyDetailsConditionLanguageVersionV1 CreateRoutingPolicyDetailsConditionLanguageVersionEnum = "V1"
)

var mappingCreateRoutingPolicyDetailsConditionLanguageVersionEnum = map[string]CreateRoutingPolicyDetailsConditionLanguageVersionEnum{
	"V1": CreateRoutingPolicyDetailsConditionLanguageVersionV1,
}

var mappingCreateRoutingPolicyDetailsConditionLanguageVersionEnumLowerCase = map[string]CreateRoutingPolicyDetailsConditionLanguageVersionEnum{
	"v1": CreateRoutingPolicyDetailsConditionLanguageVersionV1,
}

// GetCreateRoutingPolicyDetailsConditionLanguageVersionEnumValues Enumerates the set of values for CreateRoutingPolicyDetailsConditionLanguageVersionEnum
func GetCreateRoutingPolicyDetailsConditionLanguageVersionEnumValues() []CreateRoutingPolicyDetailsConditionLanguageVersionEnum {
	values := make([]CreateRoutingPolicyDetailsConditionLanguageVersionEnum, 0)
	for _, v := range mappingCreateRoutingPolicyDetailsConditionLanguageVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateRoutingPolicyDetailsConditionLanguageVersionEnumStringValues Enumerates the set of values in String for CreateRoutingPolicyDetailsConditionLanguageVersionEnum
func GetCreateRoutingPolicyDetailsConditionLanguageVersionEnumStringValues() []string {
	return []string{
		"V1",
	}
}

// GetMappingCreateRoutingPolicyDetailsConditionLanguageVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateRoutingPolicyDetailsConditionLanguageVersionEnum(val string) (CreateRoutingPolicyDetailsConditionLanguageVersionEnum, bool) {
	enum, ok := mappingCreateRoutingPolicyDetailsConditionLanguageVersionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
