// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CachingRuleSummary The caching rule settings.
type CachingRuleSummary struct {

	// The name of the caching rule.
	Name *string `mandatory:"true" json:"name"`

	// The action to take when the criteria of a caching rule are met.
	// - **CACHE:** Caches requested content when the criteria of the rule are met.
	// - **BYPASS_CACHE:** Allows requests to bypass the cache and be directed to the origin when the criteria of the rule is met.
	Action CachingRuleSummaryActionEnum `mandatory:"true" json:"action"`

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

func (m CachingRuleSummary) String() string {
	return common.PointerString(m)
}

// CachingRuleSummaryActionEnum Enum with underlying type: string
type CachingRuleSummaryActionEnum string

// Set of constants representing the allowable values for CachingRuleSummaryActionEnum
const (
	CachingRuleSummaryActionCache       CachingRuleSummaryActionEnum = "CACHE"
	CachingRuleSummaryActionBypassCache CachingRuleSummaryActionEnum = "BYPASS_CACHE"
)

var mappingCachingRuleSummaryAction = map[string]CachingRuleSummaryActionEnum{
	"CACHE":        CachingRuleSummaryActionCache,
	"BYPASS_CACHE": CachingRuleSummaryActionBypassCache,
}

// GetCachingRuleSummaryActionEnumValues Enumerates the set of values for CachingRuleSummaryActionEnum
func GetCachingRuleSummaryActionEnumValues() []CachingRuleSummaryActionEnum {
	values := make([]CachingRuleSummaryActionEnum, 0)
	for _, v := range mappingCachingRuleSummaryAction {
		values = append(values, v)
	}
	return values
}
