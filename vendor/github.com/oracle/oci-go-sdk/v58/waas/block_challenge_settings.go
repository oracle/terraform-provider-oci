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

// BlockChallengeSettings The challenge settings if `action` is set to `BLOCK`.
type BlockChallengeSettings struct {

	// The method used to block requests that fail the challenge, if `action` is set to `BLOCK`. If unspecified, defaults to `SHOW_ERROR_PAGE`.
	BlockAction BlockChallengeSettingsBlockActionEnum `mandatory:"false" json:"blockAction,omitempty"`

	// The response status code to return when `action` is set to `BLOCK`, `blockAction` is set to `SET_RESPONSE_CODE` or `SHOW_ERROR_PAGE`, and the request is blocked. If unspecified, defaults to `403`. The list of available response codes: `200`, `201`, `202`, `204`, `206`, `300`, `301`, `302`, `303`, `304`, `307`, `400`, `401`, `403`, `404`, `405`, `408`, `409`, `411`, `412`, `413`, `414`, `415`, `416`, `422`, `444`, `494`, `495`, `496`, `497`, `499`, `500`, `501`, `502`, `503`, `504`, `507`.
	BlockResponseCode *int `mandatory:"false" json:"blockResponseCode"`

	// The message to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the request is blocked. If unspecified, defaults to `Access to the website is blocked`.
	BlockErrorPageMessage *string `mandatory:"false" json:"blockErrorPageMessage"`

	// The description text to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE`, and the request is blocked. If unspecified, defaults to `Access blocked by website owner. Please contact support.`
	BlockErrorPageDescription *string `mandatory:"false" json:"blockErrorPageDescription"`

	// The error code to show on the error page when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_ERROR_PAGE` and the request is blocked. If unspecified, defaults to `403`.
	BlockErrorPageCode *string `mandatory:"false" json:"blockErrorPageCode"`

	// The title used when showing a CAPTCHA challenge when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, defaults to `Are you human?`
	CaptchaTitle *string `mandatory:"false" json:"captchaTitle"`

	// The text to show in the header when showing a CAPTCHA challenge when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, defaults to `We have detected an increased number of attempts to access this webapp. To help us keep this webapp secure, please let us know that you are not a robot by entering the text from captcha below.`
	CaptchaHeader *string `mandatory:"false" json:"captchaHeader"`

	// The text to show in the footer when showing a CAPTCHA challenge when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, default to `Enter the letters and numbers as they are shown in image above`.
	CaptchaFooter *string `mandatory:"false" json:"captchaFooter"`

	// The text to show on the label of the CAPTCHA challenge submit button when `action` is set to `BLOCK`, `blockAction` is set to `SHOW_CAPTCHA`, and the request is blocked. If unspecified, defaults to `Yes, I am human`.
	CaptchaSubmitLabel *string `mandatory:"false" json:"captchaSubmitLabel"`
}

func (m BlockChallengeSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BlockChallengeSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBlockChallengeSettingsBlockActionEnum(string(m.BlockAction)); !ok && m.BlockAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BlockAction: %s. Supported values are: %s.", m.BlockAction, strings.Join(GetBlockChallengeSettingsBlockActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BlockChallengeSettingsBlockActionEnum Enum with underlying type: string
type BlockChallengeSettingsBlockActionEnum string

// Set of constants representing the allowable values for BlockChallengeSettingsBlockActionEnum
const (
	BlockChallengeSettingsBlockActionSetResponseCode BlockChallengeSettingsBlockActionEnum = "SET_RESPONSE_CODE"
	BlockChallengeSettingsBlockActionShowErrorPage   BlockChallengeSettingsBlockActionEnum = "SHOW_ERROR_PAGE"
	BlockChallengeSettingsBlockActionShowCaptcha     BlockChallengeSettingsBlockActionEnum = "SHOW_CAPTCHA"
)

var mappingBlockChallengeSettingsBlockActionEnum = map[string]BlockChallengeSettingsBlockActionEnum{
	"SET_RESPONSE_CODE": BlockChallengeSettingsBlockActionSetResponseCode,
	"SHOW_ERROR_PAGE":   BlockChallengeSettingsBlockActionShowErrorPage,
	"SHOW_CAPTCHA":      BlockChallengeSettingsBlockActionShowCaptcha,
}

// GetBlockChallengeSettingsBlockActionEnumValues Enumerates the set of values for BlockChallengeSettingsBlockActionEnum
func GetBlockChallengeSettingsBlockActionEnumValues() []BlockChallengeSettingsBlockActionEnum {
	values := make([]BlockChallengeSettingsBlockActionEnum, 0)
	for _, v := range mappingBlockChallengeSettingsBlockActionEnum {
		values = append(values, v)
	}
	return values
}

// GetBlockChallengeSettingsBlockActionEnumStringValues Enumerates the set of values in String for BlockChallengeSettingsBlockActionEnum
func GetBlockChallengeSettingsBlockActionEnumStringValues() []string {
	return []string{
		"SET_RESPONSE_CODE",
		"SHOW_ERROR_PAGE",
		"SHOW_CAPTCHA",
	}
}

// GetMappingBlockChallengeSettingsBlockActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBlockChallengeSettingsBlockActionEnum(val string) (BlockChallengeSettingsBlockActionEnum, bool) {
	mappingBlockChallengeSettingsBlockActionEnumIgnoreCase := make(map[string]BlockChallengeSettingsBlockActionEnum)
	for k, v := range mappingBlockChallengeSettingsBlockActionEnum {
		mappingBlockChallengeSettingsBlockActionEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingBlockChallengeSettingsBlockActionEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
