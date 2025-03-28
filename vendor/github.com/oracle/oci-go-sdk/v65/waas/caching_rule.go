// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CachingRule The representation of CachingRule
type CachingRule struct {

	// The name of the caching rule.
	Name *string `mandatory:"true" json:"name"`

	// The action to take when the criteria of a caching rule are met.
	// - **CACHE:** Caches requested content when the criteria of the rule are met.
	// - **BYPASS_CACHE:** Allows requests to bypass the cache and be directed to the origin when the criteria of the rule is met.
	Action CachingRuleActionEnum `mandatory:"true" json:"action"`

	// The array of the rule criteria with condition and value. The caching rule would be applied for the requests that matched any of the listed conditions.
	Criteria []CachingRuleCriteria `mandatory:"true" json:"criteria"`

	// The unique key for the caching rule.
	Key *string `mandatory:"false" json:"key"`

	// The duration to cache content for the caching rule, specified in ISO 8601 extended format. Supported units: seconds, minutes, hours, days, weeks, months. The maximum value that can be set for any unit is `99`. Mixing of multiple units is not supported. Only applies when the `action` is set to `CACHE`.
	// Example: `PT1H`
	CachingDuration *string `mandatory:"false" json:"cachingDuration"`

	// Enables or disables client caching.
	// Browsers use the `Cache-Control` header value for caching content locally in the browser. This setting overrides the addition of a `Cache-Control` header in responses.
	IsClientCachingEnabled *bool `mandatory:"false" json:"isClientCachingEnabled"`

	// The duration to cache content in the user's browser, specified in ISO 8601 extended format. Supported units: seconds, minutes, hours, days, weeks, months. The maximum value that can be set for any unit is `99`. Mixing of multiple units is not supported. Only applies when the `action` is set to `CACHE`.
	// Example: `PT1H`
	ClientCachingDuration *string `mandatory:"false" json:"clientCachingDuration"`
}

func (m CachingRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CachingRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCachingRuleActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetCachingRuleActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CachingRuleActionEnum Enum with underlying type: string
type CachingRuleActionEnum string

// Set of constants representing the allowable values for CachingRuleActionEnum
const (
	CachingRuleActionCache       CachingRuleActionEnum = "CACHE"
	CachingRuleActionBypassCache CachingRuleActionEnum = "BYPASS_CACHE"
)

var mappingCachingRuleActionEnum = map[string]CachingRuleActionEnum{
	"CACHE":        CachingRuleActionCache,
	"BYPASS_CACHE": CachingRuleActionBypassCache,
}

var mappingCachingRuleActionEnumLowerCase = map[string]CachingRuleActionEnum{
	"cache":        CachingRuleActionCache,
	"bypass_cache": CachingRuleActionBypassCache,
}

// GetCachingRuleActionEnumValues Enumerates the set of values for CachingRuleActionEnum
func GetCachingRuleActionEnumValues() []CachingRuleActionEnum {
	values := make([]CachingRuleActionEnum, 0)
	for _, v := range mappingCachingRuleActionEnum {
		values = append(values, v)
	}
	return values
}

// GetCachingRuleActionEnumStringValues Enumerates the set of values in String for CachingRuleActionEnum
func GetCachingRuleActionEnumStringValues() []string {
	return []string{
		"CACHE",
		"BYPASS_CACHE",
	}
}

// GetMappingCachingRuleActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCachingRuleActionEnum(val string) (CachingRuleActionEnum, bool) {
	enum, ok := mappingCachingRuleActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
