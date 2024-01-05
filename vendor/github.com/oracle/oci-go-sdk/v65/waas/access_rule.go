// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AccessRule A content access rule. An access rule specifies an action to take if a set of criteria is matched by a request.
type AccessRule struct {

	// The unique name of the access rule.
	Name *string `mandatory:"true" json:"name"`

	// The list of access rule criteria. The rule would be applied only for the requests that matched all the listed conditions.
	Criteria []AccessRuleCriteria `mandatory:"true" json:"criteria"`

	// The action to take when the access criteria are met for a rule. If unspecified, defaults to `ALLOW`.
	// - **ALLOW:** Takes no action, just logs the request.
	// - **DETECT:** Takes no action, but creates an alert for the request.
	// - **BLOCK:** Blocks the request by returning specified response code or showing error page.
	// - **BYPASS:** Bypasses some or all challenges.
	// - **REDIRECT:** Redirects the request to the specified URL. These fields are required when `REDIRECT` is selected: `redirectUrl`, `redirectResponseCode`.
	// - **SHOW_CAPTCHA:** Show a CAPTCHA Challenge page instead of the requested page.
	// Regardless of action, no further rules are processed once a rule is matched.
	Action AccessRuleActionEnum `mandatory:"true" json:"action"`

	// The method used to block requests if `action` is set to `BLOCK` and the access criteria are met. If unspecified, defaults to `SET_RESPONSE_CODE`.
	BlockAction AccessRuleBlockActionEnum `mandatory:"false" json:"blockAction,omitempty"`

	// The response status code to return when `action` is set to `BLOCK`, `blockAction` is set to `SET_RESPONSE_CODE`, and the access criteria are met. If unspecified, defaults to `403`. The list of available response codes: `200`, `201`, `202`, `204`, `206`, `300`, `301`, `302`, `303`, `304`, `307`, `400`, `401`, `403`, `404`, `405`, `408`, `409`, `411`, `412`, `413`, `414`, `415`, `416`, `422`, `444`, `494`, `495`, `496`, `497`, `499`, `500`, `501`, `502`, `503`, `504`, `507`.
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

	// The target to which the request should be redirected, represented as a URI reference. Required when `action` is `REDIRECT`.
	RedirectUrl *string `mandatory:"false" json:"redirectUrl"`

	// The response status code to return when `action` is set to `REDIRECT`.
	// - **MOVED_PERMANENTLY:** Used for designating the permanent movement of a page (numerical code - 301).
	// - **FOUND:** Used for designating the temporary movement of a page (numerical code - 302).
	RedirectResponseCode AccessRuleRedirectResponseCodeEnum `mandatory:"false" json:"redirectResponseCode,omitempty"`

	// The title used when showing a CAPTCHA challenge when `action` is set to `SHOW_CAPTCHA` and the request is challenged.
	CaptchaTitle *string `mandatory:"false" json:"captchaTitle"`

	// The text to show in the header when showing a CAPTCHA challenge when `action` is set to `SHOW_CAPTCHA` and the request is challenged.
	CaptchaHeader *string `mandatory:"false" json:"captchaHeader"`

	// The text to show in the footer when showing a CAPTCHA challenge when `action` is set to `SHOW_CAPTCHA` and the request is challenged.
	CaptchaFooter *string `mandatory:"false" json:"captchaFooter"`

	// The text to show on the label of the CAPTCHA challenge submit button when `action` is set to `SHOW_CAPTCHA` and the request is challenged.
	CaptchaSubmitLabel *string `mandatory:"false" json:"captchaSubmitLabel"`

	// An object that represents an action to apply to an HTTP response headers if all rule criteria will be matched regardless of `action` value.
	ResponseHeaderManipulation []HeaderManipulationAction `mandatory:"false" json:"responseHeaderManipulation"`
}

func (m AccessRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AccessRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAccessRuleActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetAccessRuleActionEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAccessRuleBlockActionEnum(string(m.BlockAction)); !ok && m.BlockAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BlockAction: %s. Supported values are: %s.", m.BlockAction, strings.Join(GetAccessRuleBlockActionEnumStringValues(), ",")))
	}
	for _, val := range m.BypassChallenges {
		if _, ok := GetMappingAccessRuleBypassChallengesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BypassChallenges: %s. Supported values are: %s.", val, strings.Join(GetAccessRuleBypassChallengesEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingAccessRuleRedirectResponseCodeEnum(string(m.RedirectResponseCode)); !ok && m.RedirectResponseCode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RedirectResponseCode: %s. Supported values are: %s.", m.RedirectResponseCode, strings.Join(GetAccessRuleRedirectResponseCodeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *AccessRule) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		BlockAction                AccessRuleBlockActionEnum          `json:"blockAction"`
		BlockResponseCode          *int                               `json:"blockResponseCode"`
		BlockErrorPageMessage      *string                            `json:"blockErrorPageMessage"`
		BlockErrorPageCode         *string                            `json:"blockErrorPageCode"`
		BlockErrorPageDescription  *string                            `json:"blockErrorPageDescription"`
		BypassChallenges           []AccessRuleBypassChallengesEnum   `json:"bypassChallenges"`
		RedirectUrl                *string                            `json:"redirectUrl"`
		RedirectResponseCode       AccessRuleRedirectResponseCodeEnum `json:"redirectResponseCode"`
		CaptchaTitle               *string                            `json:"captchaTitle"`
		CaptchaHeader              *string                            `json:"captchaHeader"`
		CaptchaFooter              *string                            `json:"captchaFooter"`
		CaptchaSubmitLabel         *string                            `json:"captchaSubmitLabel"`
		ResponseHeaderManipulation []headermanipulationaction         `json:"responseHeaderManipulation"`
		Name                       *string                            `json:"name"`
		Criteria                   []AccessRuleCriteria               `json:"criteria"`
		Action                     AccessRuleActionEnum               `json:"action"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.BlockAction = model.BlockAction

	m.BlockResponseCode = model.BlockResponseCode

	m.BlockErrorPageMessage = model.BlockErrorPageMessage

	m.BlockErrorPageCode = model.BlockErrorPageCode

	m.BlockErrorPageDescription = model.BlockErrorPageDescription

	m.BypassChallenges = make([]AccessRuleBypassChallengesEnum, len(model.BypassChallenges))
	copy(m.BypassChallenges, model.BypassChallenges)
	m.RedirectUrl = model.RedirectUrl

	m.RedirectResponseCode = model.RedirectResponseCode

	m.CaptchaTitle = model.CaptchaTitle

	m.CaptchaHeader = model.CaptchaHeader

	m.CaptchaFooter = model.CaptchaFooter

	m.CaptchaSubmitLabel = model.CaptchaSubmitLabel

	m.ResponseHeaderManipulation = make([]HeaderManipulationAction, len(model.ResponseHeaderManipulation))
	for i, n := range model.ResponseHeaderManipulation {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ResponseHeaderManipulation[i] = nn.(HeaderManipulationAction)
		} else {
			m.ResponseHeaderManipulation[i] = nil
		}
	}
	m.Name = model.Name

	m.Criteria = make([]AccessRuleCriteria, len(model.Criteria))
	copy(m.Criteria, model.Criteria)
	m.Action = model.Action

	return
}

// AccessRuleActionEnum Enum with underlying type: string
type AccessRuleActionEnum string

// Set of constants representing the allowable values for AccessRuleActionEnum
const (
	AccessRuleActionAllow       AccessRuleActionEnum = "ALLOW"
	AccessRuleActionDetect      AccessRuleActionEnum = "DETECT"
	AccessRuleActionBlock       AccessRuleActionEnum = "BLOCK"
	AccessRuleActionBypass      AccessRuleActionEnum = "BYPASS"
	AccessRuleActionRedirect    AccessRuleActionEnum = "REDIRECT"
	AccessRuleActionShowCaptcha AccessRuleActionEnum = "SHOW_CAPTCHA"
)

var mappingAccessRuleActionEnum = map[string]AccessRuleActionEnum{
	"ALLOW":        AccessRuleActionAllow,
	"DETECT":       AccessRuleActionDetect,
	"BLOCK":        AccessRuleActionBlock,
	"BYPASS":       AccessRuleActionBypass,
	"REDIRECT":     AccessRuleActionRedirect,
	"SHOW_CAPTCHA": AccessRuleActionShowCaptcha,
}

var mappingAccessRuleActionEnumLowerCase = map[string]AccessRuleActionEnum{
	"allow":        AccessRuleActionAllow,
	"detect":       AccessRuleActionDetect,
	"block":        AccessRuleActionBlock,
	"bypass":       AccessRuleActionBypass,
	"redirect":     AccessRuleActionRedirect,
	"show_captcha": AccessRuleActionShowCaptcha,
}

// GetAccessRuleActionEnumValues Enumerates the set of values for AccessRuleActionEnum
func GetAccessRuleActionEnumValues() []AccessRuleActionEnum {
	values := make([]AccessRuleActionEnum, 0)
	for _, v := range mappingAccessRuleActionEnum {
		values = append(values, v)
	}
	return values
}

// GetAccessRuleActionEnumStringValues Enumerates the set of values in String for AccessRuleActionEnum
func GetAccessRuleActionEnumStringValues() []string {
	return []string{
		"ALLOW",
		"DETECT",
		"BLOCK",
		"BYPASS",
		"REDIRECT",
		"SHOW_CAPTCHA",
	}
}

// GetMappingAccessRuleActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAccessRuleActionEnum(val string) (AccessRuleActionEnum, bool) {
	enum, ok := mappingAccessRuleActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AccessRuleBlockActionEnum Enum with underlying type: string
type AccessRuleBlockActionEnum string

// Set of constants representing the allowable values for AccessRuleBlockActionEnum
const (
	AccessRuleBlockActionSetResponseCode AccessRuleBlockActionEnum = "SET_RESPONSE_CODE"
	AccessRuleBlockActionShowErrorPage   AccessRuleBlockActionEnum = "SHOW_ERROR_PAGE"
)

var mappingAccessRuleBlockActionEnum = map[string]AccessRuleBlockActionEnum{
	"SET_RESPONSE_CODE": AccessRuleBlockActionSetResponseCode,
	"SHOW_ERROR_PAGE":   AccessRuleBlockActionShowErrorPage,
}

var mappingAccessRuleBlockActionEnumLowerCase = map[string]AccessRuleBlockActionEnum{
	"set_response_code": AccessRuleBlockActionSetResponseCode,
	"show_error_page":   AccessRuleBlockActionShowErrorPage,
}

// GetAccessRuleBlockActionEnumValues Enumerates the set of values for AccessRuleBlockActionEnum
func GetAccessRuleBlockActionEnumValues() []AccessRuleBlockActionEnum {
	values := make([]AccessRuleBlockActionEnum, 0)
	for _, v := range mappingAccessRuleBlockActionEnum {
		values = append(values, v)
	}
	return values
}

// GetAccessRuleBlockActionEnumStringValues Enumerates the set of values in String for AccessRuleBlockActionEnum
func GetAccessRuleBlockActionEnumStringValues() []string {
	return []string{
		"SET_RESPONSE_CODE",
		"SHOW_ERROR_PAGE",
	}
}

// GetMappingAccessRuleBlockActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAccessRuleBlockActionEnum(val string) (AccessRuleBlockActionEnum, bool) {
	enum, ok := mappingAccessRuleBlockActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingAccessRuleBypassChallengesEnum = map[string]AccessRuleBypassChallengesEnum{
	"JS_CHALLENGE":                 AccessRuleBypassChallengesJsChallenge,
	"DEVICE_FINGERPRINT_CHALLENGE": AccessRuleBypassChallengesDeviceFingerprintChallenge,
	"HUMAN_INTERACTION_CHALLENGE":  AccessRuleBypassChallengesHumanInteractionChallenge,
	"CAPTCHA":                      AccessRuleBypassChallengesCaptcha,
}

var mappingAccessRuleBypassChallengesEnumLowerCase = map[string]AccessRuleBypassChallengesEnum{
	"js_challenge":                 AccessRuleBypassChallengesJsChallenge,
	"device_fingerprint_challenge": AccessRuleBypassChallengesDeviceFingerprintChallenge,
	"human_interaction_challenge":  AccessRuleBypassChallengesHumanInteractionChallenge,
	"captcha":                      AccessRuleBypassChallengesCaptcha,
}

// GetAccessRuleBypassChallengesEnumValues Enumerates the set of values for AccessRuleBypassChallengesEnum
func GetAccessRuleBypassChallengesEnumValues() []AccessRuleBypassChallengesEnum {
	values := make([]AccessRuleBypassChallengesEnum, 0)
	for _, v := range mappingAccessRuleBypassChallengesEnum {
		values = append(values, v)
	}
	return values
}

// GetAccessRuleBypassChallengesEnumStringValues Enumerates the set of values in String for AccessRuleBypassChallengesEnum
func GetAccessRuleBypassChallengesEnumStringValues() []string {
	return []string{
		"JS_CHALLENGE",
		"DEVICE_FINGERPRINT_CHALLENGE",
		"HUMAN_INTERACTION_CHALLENGE",
		"CAPTCHA",
	}
}

// GetMappingAccessRuleBypassChallengesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAccessRuleBypassChallengesEnum(val string) (AccessRuleBypassChallengesEnum, bool) {
	enum, ok := mappingAccessRuleBypassChallengesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AccessRuleRedirectResponseCodeEnum Enum with underlying type: string
type AccessRuleRedirectResponseCodeEnum string

// Set of constants representing the allowable values for AccessRuleRedirectResponseCodeEnum
const (
	AccessRuleRedirectResponseCodeMovedPermanently AccessRuleRedirectResponseCodeEnum = "MOVED_PERMANENTLY"
	AccessRuleRedirectResponseCodeFound            AccessRuleRedirectResponseCodeEnum = "FOUND"
)

var mappingAccessRuleRedirectResponseCodeEnum = map[string]AccessRuleRedirectResponseCodeEnum{
	"MOVED_PERMANENTLY": AccessRuleRedirectResponseCodeMovedPermanently,
	"FOUND":             AccessRuleRedirectResponseCodeFound,
}

var mappingAccessRuleRedirectResponseCodeEnumLowerCase = map[string]AccessRuleRedirectResponseCodeEnum{
	"moved_permanently": AccessRuleRedirectResponseCodeMovedPermanently,
	"found":             AccessRuleRedirectResponseCodeFound,
}

// GetAccessRuleRedirectResponseCodeEnumValues Enumerates the set of values for AccessRuleRedirectResponseCodeEnum
func GetAccessRuleRedirectResponseCodeEnumValues() []AccessRuleRedirectResponseCodeEnum {
	values := make([]AccessRuleRedirectResponseCodeEnum, 0)
	for _, v := range mappingAccessRuleRedirectResponseCodeEnum {
		values = append(values, v)
	}
	return values
}

// GetAccessRuleRedirectResponseCodeEnumStringValues Enumerates the set of values in String for AccessRuleRedirectResponseCodeEnum
func GetAccessRuleRedirectResponseCodeEnumStringValues() []string {
	return []string{
		"MOVED_PERMANENTLY",
		"FOUND",
	}
}

// GetMappingAccessRuleRedirectResponseCodeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAccessRuleRedirectResponseCodeEnum(val string) (AccessRuleRedirectResponseCodeEnum, bool) {
	enum, ok := mappingAccessRuleRedirectResponseCodeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
