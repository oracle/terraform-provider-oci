// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CachingRuleCriteria A caching rule criteria condition and value.
type CachingRuleCriteria struct {

	// The condition of the caching rule criteria.
	// - **URL_IS:** Matches if the concatenation of request URL path and query is identical to the contents of the `value` field.
	// - **URL_STARTS_WITH:** Matches if the concatenation of request URL path and query starts with the contents of the `value` field.
	// - **URL_PART_ENDS_WITH:** Matches if the concatenation of request URL path and query ends with the contents of the `value` field.
	// - **URL_PART_CONTAINS:** Matches if the concatenation of request URL path and query contains the contents of the `value` field.
	// URL must start with /
	// URL can't contain restricted double slashes //
	// URL can't contain restricted ' & ? symbols
	Condition CachingRuleCriteriaConditionEnum `mandatory:"true" json:"condition"`

	// The value of the caching rule criteria.
	Value *string `mandatory:"true" json:"value"`
}

func (m CachingRuleCriteria) String() string {
	return common.PointerString(m)
}

// CachingRuleCriteriaConditionEnum Enum with underlying type: string
type CachingRuleCriteriaConditionEnum string

// Set of constants representing the allowable values for CachingRuleCriteriaConditionEnum
const (
	CachingRuleCriteriaConditionIs           CachingRuleCriteriaConditionEnum = "URL_IS"
	CachingRuleCriteriaConditionStartsWith   CachingRuleCriteriaConditionEnum = "URL_STARTS_WITH"
	CachingRuleCriteriaConditionPartEndsWith CachingRuleCriteriaConditionEnum = "URL_PART_ENDS_WITH"
	CachingRuleCriteriaConditionPartContains CachingRuleCriteriaConditionEnum = "URL_PART_CONTAINS"
)

var mappingCachingRuleCriteriaCondition = map[string]CachingRuleCriteriaConditionEnum{
	"URL_IS":             CachingRuleCriteriaConditionIs,
	"URL_STARTS_WITH":    CachingRuleCriteriaConditionStartsWith,
	"URL_PART_ENDS_WITH": CachingRuleCriteriaConditionPartEndsWith,
	"URL_PART_CONTAINS":  CachingRuleCriteriaConditionPartContains,
}

// GetCachingRuleCriteriaConditionEnumValues Enumerates the set of values for CachingRuleCriteriaConditionEnum
func GetCachingRuleCriteriaConditionEnumValues() []CachingRuleCriteriaConditionEnum {
	values := make([]CachingRuleCriteriaConditionEnum, 0)
	for _, v := range mappingCachingRuleCriteriaCondition {
		values = append(values, v)
	}
	return values
}
