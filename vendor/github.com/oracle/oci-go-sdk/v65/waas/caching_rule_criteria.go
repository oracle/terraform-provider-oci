// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CachingRuleCriteria A caching rule criteria condition and value.
type CachingRuleCriteria struct {

	// The condition of the caching rule criteria.
	// - **URL_IS:** Matches if the concatenation of request URL path and query is identical to the contents of the `value` field.
	// - **URL_STARTS_WITH:** Matches if the concatenation of request URL path and query starts with the contents of the `value` field.
	// - **URL_PART_ENDS_WITH:** Matches if the concatenation of request URL path and query ends with the contents of the `value` field.
	// - **URL_PART_CONTAINS:** Matches if the concatenation of request URL path and query contains the contents of the `value` field.
	// URLs must start with a `/`. URLs can't contain restricted double slashes `//`. URLs can't contain the restricted `'` `&` `?` symbols. Resources to cache can only be specified by a URL, any query parameters are ignored.
	Condition CachingRuleCriteriaConditionEnum `mandatory:"true" json:"condition"`

	// The value of the caching rule criteria.
	Value *string `mandatory:"true" json:"value"`
}

func (m CachingRuleCriteria) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CachingRuleCriteria) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCachingRuleCriteriaConditionEnum(string(m.Condition)); !ok && m.Condition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Condition: %s. Supported values are: %s.", m.Condition, strings.Join(GetCachingRuleCriteriaConditionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingCachingRuleCriteriaConditionEnum = map[string]CachingRuleCriteriaConditionEnum{
	"URL_IS":             CachingRuleCriteriaConditionIs,
	"URL_STARTS_WITH":    CachingRuleCriteriaConditionStartsWith,
	"URL_PART_ENDS_WITH": CachingRuleCriteriaConditionPartEndsWith,
	"URL_PART_CONTAINS":  CachingRuleCriteriaConditionPartContains,
}

var mappingCachingRuleCriteriaConditionEnumLowerCase = map[string]CachingRuleCriteriaConditionEnum{
	"url_is":             CachingRuleCriteriaConditionIs,
	"url_starts_with":    CachingRuleCriteriaConditionStartsWith,
	"url_part_ends_with": CachingRuleCriteriaConditionPartEndsWith,
	"url_part_contains":  CachingRuleCriteriaConditionPartContains,
}

// GetCachingRuleCriteriaConditionEnumValues Enumerates the set of values for CachingRuleCriteriaConditionEnum
func GetCachingRuleCriteriaConditionEnumValues() []CachingRuleCriteriaConditionEnum {
	values := make([]CachingRuleCriteriaConditionEnum, 0)
	for _, v := range mappingCachingRuleCriteriaConditionEnum {
		values = append(values, v)
	}
	return values
}

// GetCachingRuleCriteriaConditionEnumStringValues Enumerates the set of values in String for CachingRuleCriteriaConditionEnum
func GetCachingRuleCriteriaConditionEnumStringValues() []string {
	return []string{
		"URL_IS",
		"URL_STARTS_WITH",
		"URL_PART_ENDS_WITH",
		"URL_PART_CONTAINS",
	}
}

// GetMappingCachingRuleCriteriaConditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCachingRuleCriteriaConditionEnum(val string) (CachingRuleCriteriaConditionEnum, bool) {
	enum, ok := mappingCachingRuleCriteriaConditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
