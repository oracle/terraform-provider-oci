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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateRoutingPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateRoutingPolicyDetailsConditionLanguageVersionEnum(string(m.ConditionLanguageVersion)); !ok && m.ConditionLanguageVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConditionLanguageVersion: %s. Supported values are: %s.", m.ConditionLanguageVersion, strings.Join(GetUpdateRoutingPolicyDetailsConditionLanguageVersionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateRoutingPolicyDetailsConditionLanguageVersionEnum Enum with underlying type: string
type UpdateRoutingPolicyDetailsConditionLanguageVersionEnum string

// Set of constants representing the allowable values for UpdateRoutingPolicyDetailsConditionLanguageVersionEnum
const (
	UpdateRoutingPolicyDetailsConditionLanguageVersionV1 UpdateRoutingPolicyDetailsConditionLanguageVersionEnum = "V1"
)

var mappingUpdateRoutingPolicyDetailsConditionLanguageVersionEnum = map[string]UpdateRoutingPolicyDetailsConditionLanguageVersionEnum{
	"V1": UpdateRoutingPolicyDetailsConditionLanguageVersionV1,
}

// GetUpdateRoutingPolicyDetailsConditionLanguageVersionEnumValues Enumerates the set of values for UpdateRoutingPolicyDetailsConditionLanguageVersionEnum
func GetUpdateRoutingPolicyDetailsConditionLanguageVersionEnumValues() []UpdateRoutingPolicyDetailsConditionLanguageVersionEnum {
	values := make([]UpdateRoutingPolicyDetailsConditionLanguageVersionEnum, 0)
	for _, v := range mappingUpdateRoutingPolicyDetailsConditionLanguageVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateRoutingPolicyDetailsConditionLanguageVersionEnumStringValues Enumerates the set of values in String for UpdateRoutingPolicyDetailsConditionLanguageVersionEnum
func GetUpdateRoutingPolicyDetailsConditionLanguageVersionEnumStringValues() []string {
	return []string{
		"V1",
	}
}

// GetMappingUpdateRoutingPolicyDetailsConditionLanguageVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateRoutingPolicyDetailsConditionLanguageVersionEnum(val string) (UpdateRoutingPolicyDetailsConditionLanguageVersionEnum, bool) {
	mappingUpdateRoutingPolicyDetailsConditionLanguageVersionEnumIgnoreCase := make(map[string]UpdateRoutingPolicyDetailsConditionLanguageVersionEnum)
	for k, v := range mappingUpdateRoutingPolicyDetailsConditionLanguageVersionEnum {
		mappingUpdateRoutingPolicyDetailsConditionLanguageVersionEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUpdateRoutingPolicyDetailsConditionLanguageVersionEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
