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

// CachingRuleSummary The caching rule settings.
type CachingRuleSummary struct {

	// The name of the caching rule.
	Name *string `mandatory:"true" json:"name"`

	// The action to take on matched caching rules.
	// - **CACHE:** Allow to set caching rule, which would be cached.
	// - **BYPASS_CACHE:** Allow to set caching rule, which would never be cached. e.g. all requests would be passed directly to origin for those file types.
	Action CachingRuleSummaryActionEnum `mandatory:"true" json:"action"`

	// The array of the rule criteria with condition and value.
	Criteria []CachingRuleCriteria `mandatory:"true" json:"criteria"`

	// The unique key for the caching rule.
	Key *string `mandatory:"false" json:"key"`

	// The caching duration (applies only to `CACHE` action) specified in ISO 8601 extended format. Supported units: seconds, minutes, hours, days, weeks, months. Max value - 99. Mixing of multiple units is not supported.
	CachingDuration *string `mandatory:"false" json:"cachingDuration"`

	// Enables or disables the client caching.
	// Browsers use the Cache-Control header value for caching content locally, in the browser.
	// This setting will control the addition of a Cache-Control header to responses. It overrides existing Cache-Control headers.
	IsClientCachingEnabled *bool `mandatory:"false" json:"isClientCachingEnabled"`

	// The client caching duration (applies only to `CACHE` action) specified in ISO 8601 extended format, in case client caching enabled. It sets Cache-Control header max-age time, i.e. the local browser cache expire time. Supported units: seconds, minutes, hours, days, weeks, months. Max value - 99. Mixing of multiple units is not supported.
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
