// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// WafBlockedRequest The representation of WafBlockedRequest
type WafBlockedRequest struct {

	// The date and time the blocked requests were observed, expressed in RFC 3339 timestamp format.
	TimeObserved *common.SDKTime `mandatory:"false" json:"timeObserved"`

	// The number of seconds the data covers.
	TimeRangeInSeconds *int `mandatory:"false" json:"timeRangeInSeconds"`

	// The specific Web Application Firewall feature that blocked the requests, such as JavaScript Challenge or Access Control.
	WafFeature WafBlockedRequestWafFeatureEnum `mandatory:"false" json:"wafFeature,omitempty"`

	// The count of blocked requests.
	Count *int `mandatory:"false" json:"count"`
}

func (m WafBlockedRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WafBlockedRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingWafBlockedRequestWafFeatureEnum(string(m.WafFeature)); !ok && m.WafFeature != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WafFeature: %s. Supported values are: %s.", m.WafFeature, strings.Join(GetWafBlockedRequestWafFeatureEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WafBlockedRequestWafFeatureEnum Enum with underlying type: string
type WafBlockedRequestWafFeatureEnum string

// Set of constants representing the allowable values for WafBlockedRequestWafFeatureEnum
const (
	WafBlockedRequestWafFeatureProtectionRules            WafBlockedRequestWafFeatureEnum = "PROTECTION_RULES"
	WafBlockedRequestWafFeatureJsChallenge                WafBlockedRequestWafFeatureEnum = "JS_CHALLENGE"
	WafBlockedRequestWafFeatureAccessRules                WafBlockedRequestWafFeatureEnum = "ACCESS_RULES"
	WafBlockedRequestWafFeatureThreatFeeds                WafBlockedRequestWafFeatureEnum = "THREAT_FEEDS"
	WafBlockedRequestWafFeatureHumanInteractionChallenge  WafBlockedRequestWafFeatureEnum = "HUMAN_INTERACTION_CHALLENGE"
	WafBlockedRequestWafFeatureDeviceFingerprintChallenge WafBlockedRequestWafFeatureEnum = "DEVICE_FINGERPRINT_CHALLENGE"
	WafBlockedRequestWafFeatureCaptcha                    WafBlockedRequestWafFeatureEnum = "CAPTCHA"
	WafBlockedRequestWafFeatureAddressRateLimiting        WafBlockedRequestWafFeatureEnum = "ADDRESS_RATE_LIMITING"
)

var mappingWafBlockedRequestWafFeatureEnum = map[string]WafBlockedRequestWafFeatureEnum{
	"PROTECTION_RULES":             WafBlockedRequestWafFeatureProtectionRules,
	"JS_CHALLENGE":                 WafBlockedRequestWafFeatureJsChallenge,
	"ACCESS_RULES":                 WafBlockedRequestWafFeatureAccessRules,
	"THREAT_FEEDS":                 WafBlockedRequestWafFeatureThreatFeeds,
	"HUMAN_INTERACTION_CHALLENGE":  WafBlockedRequestWafFeatureHumanInteractionChallenge,
	"DEVICE_FINGERPRINT_CHALLENGE": WafBlockedRequestWafFeatureDeviceFingerprintChallenge,
	"CAPTCHA":                      WafBlockedRequestWafFeatureCaptcha,
	"ADDRESS_RATE_LIMITING":        WafBlockedRequestWafFeatureAddressRateLimiting,
}

// GetWafBlockedRequestWafFeatureEnumValues Enumerates the set of values for WafBlockedRequestWafFeatureEnum
func GetWafBlockedRequestWafFeatureEnumValues() []WafBlockedRequestWafFeatureEnum {
	values := make([]WafBlockedRequestWafFeatureEnum, 0)
	for _, v := range mappingWafBlockedRequestWafFeatureEnum {
		values = append(values, v)
	}
	return values
}

// GetWafBlockedRequestWafFeatureEnumStringValues Enumerates the set of values in String for WafBlockedRequestWafFeatureEnum
func GetWafBlockedRequestWafFeatureEnumStringValues() []string {
	return []string{
		"PROTECTION_RULES",
		"JS_CHALLENGE",
		"ACCESS_RULES",
		"THREAT_FEEDS",
		"HUMAN_INTERACTION_CHALLENGE",
		"DEVICE_FINGERPRINT_CHALLENGE",
		"CAPTCHA",
		"ADDRESS_RATE_LIMITING",
	}
}

// GetMappingWafBlockedRequestWafFeatureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWafBlockedRequestWafFeatureEnum(val string) (WafBlockedRequestWafFeatureEnum, bool) {
	mappingWafBlockedRequestWafFeatureEnumIgnoreCase := make(map[string]WafBlockedRequestWafFeatureEnum)
	for k, v := range mappingWafBlockedRequestWafFeatureEnum {
		mappingWafBlockedRequestWafFeatureEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingWafBlockedRequestWafFeatureEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
