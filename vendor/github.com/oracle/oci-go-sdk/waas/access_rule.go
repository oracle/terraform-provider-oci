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

// AccessRule A content access rule. An access rule specifies an action to take if a set of criteria is matched by a request.
type AccessRule struct {

	// The unique name of the access rule.
	Name *string `mandatory:"true" json:"name"`

	// The list of access rule criteria.
	Criteria []AccessRuleCriteria `mandatory:"true" json:"criteria"`

	// The action to take when the access criteria are met for a rule. If unspecified, defaults to `ALLOW`.
	// - **ALLOW:** Takes no action, just logs the request.
	// - **DETECT:** Takes no action, but creates an alert for the request.
	// - **BLOCK:** Blocks the request by returning specified response code or showing error page.
	// - **BYPASS:** Bypasses some or all challenges.
	// - **REDIRECT:** Redirects the request to the specified URL.
	// Regardless of action, no further rules are processed once the rule is matched.
	Action AccessRuleActionEnum `mandatory:"true" json:"action"`

	// The method used to block requests if `action` is set to `BLOCK` and the access criteria are met. If unspecified, defaults to `SET_RESPONSE_CODE`.
	BlockAction AccessRuleBlockActionEnum `mandatory:"false" json:"blockAction,omitempty"`

	// The response status code to return when `action` is set to `BLOCK`, `blockAction` is set to `SET_RESPONSE_CODE`, and the access criteria are met. If unspecified, defaults to `403`.
	BlockResponseCode *int `mandatory:"false" json:"blockResponseCode"`

	// The message to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the access criteria are met. If unspecified, defaults to 'Access to the website is blocked.'
	BlockErrorPageMessage *string `mandatory:"false" json:"blockErrorPageMessage"`

	// The error code to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the access criteria are met. If unspecified, defaults to 'Access rules'.
	BlockErrorPageCode *string `mandatory:"false" json:"blockErrorPageCode"`

	// The description text to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the access criteria are met. If unspecified, defaults to 'Access blocked by website owner. Please contact support.'
	BlockErrorPageDescription *string `mandatory:"false" json:"blockErrorPageDescription"`

	// The list of challenges to bypass when `action` is set to `BYPASS`. If unspecified or empty, all challenges are bypassed.
	// - **JS_CHALLENGE:** Bypasses JavaScript Challenge.
	// - **DEVICE_FINGERPRINT_CHALLENGE:** Bypasses Device Fingerprint Challenge.
	// - **HUMAN_INTERACTION_CHALLENGE:** Bypasses Human Interaction Challenge.
	// - **CAPTCHA:** Bypasses CAPTCHA Challenge.
	BypassChallenges []AccessRuleBypassChallengesEnum `mandatory:"false" json:"bypassChallenges,omitempty"`

	// The target to which the request should be redirected, represented as a URI reference.
	RedirectUrl *string `mandatory:"false" json:"redirectUrl"`

	// The response status code to return when `action` is set to `REDIRECT`.
	// - **MOVED_PERMANENTLY:** Used for designating the permanent movement of a page (numerical code - 301).
	// - **FOUND:** Used for designating the temporary movement of a page (numerical code - 302).
	RedirectResponseCode AccessRuleRedirectResponseCodeEnum `mandatory:"false" json:"redirectResponseCode,omitempty"`
}

func (m AccessRule) String() string {
	return common.PointerString(m)
}

// AccessRuleActionEnum Enum with underlying type: string
type AccessRuleActionEnum string

// Set of constants representing the allowable values for AccessRuleActionEnum
const (
	AccessRuleActionAllow    AccessRuleActionEnum = "ALLOW"
	AccessRuleActionDetect   AccessRuleActionEnum = "DETECT"
	AccessRuleActionBlock    AccessRuleActionEnum = "BLOCK"
	AccessRuleActionBypass   AccessRuleActionEnum = "BYPASS"
	AccessRuleActionRedirect AccessRuleActionEnum = "REDIRECT"
)

var mappingAccessRuleAction = map[string]AccessRuleActionEnum{
	"ALLOW":    AccessRuleActionAllow,
	"DETECT":   AccessRuleActionDetect,
	"BLOCK":    AccessRuleActionBlock,
	"BYPASS":   AccessRuleActionBypass,
	"REDIRECT": AccessRuleActionRedirect,
}

// GetAccessRuleActionEnumValues Enumerates the set of values for AccessRuleActionEnum
func GetAccessRuleActionEnumValues() []AccessRuleActionEnum {
	values := make([]AccessRuleActionEnum, 0)
	for _, v := range mappingAccessRuleAction {
		values = append(values, v)
	}
	return values
}

// AccessRuleBlockActionEnum Enum with underlying type: string
type AccessRuleBlockActionEnum string

// Set of constants representing the allowable values for AccessRuleBlockActionEnum
const (
	AccessRuleBlockActionSetResponseCode AccessRuleBlockActionEnum = "SET_RESPONSE_CODE"
	AccessRuleBlockActionShowErrorPage   AccessRuleBlockActionEnum = "SHOW_ERROR_PAGE"
)

var mappingAccessRuleBlockAction = map[string]AccessRuleBlockActionEnum{
	"SET_RESPONSE_CODE": AccessRuleBlockActionSetResponseCode,
	"SHOW_ERROR_PAGE":   AccessRuleBlockActionShowErrorPage,
}

// GetAccessRuleBlockActionEnumValues Enumerates the set of values for AccessRuleBlockActionEnum
func GetAccessRuleBlockActionEnumValues() []AccessRuleBlockActionEnum {
	values := make([]AccessRuleBlockActionEnum, 0)
	for _, v := range mappingAccessRuleBlockAction {
		values = append(values, v)
	}
	return values
}

// AccessRuleBypassChallengesEnum Enum with underlying type: string
type AccessRuleBypassChallengesEnum string

// Set of constants representing the allowable values for AccessRuleBypassChallengesEnum
const (
	AccessRuleBypassChallengesJsChallenge                AccessRuleBypassChallengesEnum = "JS_CHALLENGE"
	AccessRuleBypassChallengesDeviceFingerprintChallenge AccessRuleBypassChallengesEnum = "DEVICE_FINGERPRINT_CHALLENGE"
	AccessRuleBypassChallengesHumanInteractionChallenge  AccessRuleBypassChallengesEnum = "HUMAN_INTERACTION_CHALLENGE"
	AccessRuleBypassChallengesCaptcha                    AccessRuleBypassChallengesEnum = "CAPTCHA"
)

var mappingAccessRuleBypassChallenges = map[string]AccessRuleBypassChallengesEnum{
	"JS_CHALLENGE":                 AccessRuleBypassChallengesJsChallenge,
	"DEVICE_FINGERPRINT_CHALLENGE": AccessRuleBypassChallengesDeviceFingerprintChallenge,
	"HUMAN_INTERACTION_CHALLENGE":  AccessRuleBypassChallengesHumanInteractionChallenge,
	"CAPTCHA":                      AccessRuleBypassChallengesCaptcha,
}

// GetAccessRuleBypassChallengesEnumValues Enumerates the set of values for AccessRuleBypassChallengesEnum
func GetAccessRuleBypassChallengesEnumValues() []AccessRuleBypassChallengesEnum {
	values := make([]AccessRuleBypassChallengesEnum, 0)
	for _, v := range mappingAccessRuleBypassChallenges {
		values = append(values, v)
	}
	return values
}

// AccessRuleRedirectResponseCodeEnum Enum with underlying type: string
type AccessRuleRedirectResponseCodeEnum string

// Set of constants representing the allowable values for AccessRuleRedirectResponseCodeEnum
const (
	AccessRuleRedirectResponseCodeMovedPermanently AccessRuleRedirectResponseCodeEnum = "MOVED_PERMANENTLY"
	AccessRuleRedirectResponseCodeFound            AccessRuleRedirectResponseCodeEnum = "FOUND"
)

var mappingAccessRuleRedirectResponseCode = map[string]AccessRuleRedirectResponseCodeEnum{
	"MOVED_PERMANENTLY": AccessRuleRedirectResponseCodeMovedPermanently,
	"FOUND":             AccessRuleRedirectResponseCodeFound,
}

// GetAccessRuleRedirectResponseCodeEnumValues Enumerates the set of values for AccessRuleRedirectResponseCodeEnum
func GetAccessRuleRedirectResponseCodeEnumValues() []AccessRuleRedirectResponseCodeEnum {
	values := make([]AccessRuleRedirectResponseCodeEnum, 0)
	for _, v := range mappingAccessRuleRedirectResponseCode {
		values = append(values, v)
	}
	return values
}
