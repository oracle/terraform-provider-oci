// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// PathMatchCondition The path string and match condition to apply when evaluating an incoming URI for redirection.
type PathMatchCondition struct {

	// The path string that the redirection rule applies to.
	// Example: `/example`
	AttributeValue *string `mandatory:"true" json:"attributeValue"`

	// A string that specifies how to compare the PathMatchCondition object's `attributeValue` string to the
	// incoming URI.
	// *  **EXACT_MATCH** - The incoming URI path must exactly and completely match the `attributeValue` string.
	// *  **FORCE_LONGEST_PREFIX_MATCH** - The system looks for the `attributeValue` string with the best,
	//    longest match of the beginning portion of the incoming URI path.
	// *  **PREFIX_MATCH** - The beginning portion of the incoming URI path must exactly match the
	//    `attributeValue` string.
	// *  **SUFFIX_MATCH** - The ending portion of the incoming URI path must exactly match the `attributeValue`
	//    string.
	Operator PathMatchConditionOperatorEnum `mandatory:"true" json:"operator"`
}

func (m PathMatchCondition) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m PathMatchCondition) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePathMatchCondition PathMatchCondition
	s := struct {
		DiscriminatorParam string `json:"attributeName"`
		MarshalTypePathMatchCondition
	}{
		"PATH",
		(MarshalTypePathMatchCondition)(m),
	}

	return json.Marshal(&s)
}

// PathMatchConditionOperatorEnum Enum with underlying type: string
type PathMatchConditionOperatorEnum string

// Set of constants representing the allowable values for PathMatchConditionOperatorEnum
const (
	PathMatchConditionOperatorExactMatch              PathMatchConditionOperatorEnum = "EXACT_MATCH"
	PathMatchConditionOperatorForceLongestPrefixMatch PathMatchConditionOperatorEnum = "FORCE_LONGEST_PREFIX_MATCH"
	PathMatchConditionOperatorPrefixMatch             PathMatchConditionOperatorEnum = "PREFIX_MATCH"
	PathMatchConditionOperatorSuffixMatch             PathMatchConditionOperatorEnum = "SUFFIX_MATCH"
)

var mappingPathMatchConditionOperator = map[string]PathMatchConditionOperatorEnum{
	"EXACT_MATCH":                PathMatchConditionOperatorExactMatch,
	"FORCE_LONGEST_PREFIX_MATCH": PathMatchConditionOperatorForceLongestPrefixMatch,
	"PREFIX_MATCH":               PathMatchConditionOperatorPrefixMatch,
	"SUFFIX_MATCH":               PathMatchConditionOperatorSuffixMatch,
}

// GetPathMatchConditionOperatorEnumValues Enumerates the set of values for PathMatchConditionOperatorEnum
func GetPathMatchConditionOperatorEnumValues() []PathMatchConditionOperatorEnum {
	values := make([]PathMatchConditionOperatorEnum, 0)
	for _, v := range mappingPathMatchConditionOperator {
		values = append(values, v)
	}
	return values
}
