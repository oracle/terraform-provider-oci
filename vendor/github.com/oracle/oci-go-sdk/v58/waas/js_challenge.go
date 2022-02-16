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

// JsChallenge The JavaScript challenge settings. JavaScript Challenge is the function to filter abnormal or malicious bots and allow access to real clients.
type JsChallenge struct {

	// Enables or disables the JavaScript challenge Web Application Firewall feature.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The action to take against requests from detected bots. If unspecified, defaults to `DETECT`.
	Action JsChallengeActionEnum `mandatory:"false" json:"action,omitempty"`

	// The number of failed requests before taking action. If unspecified, defaults to `10`.
	FailureThreshold *int `mandatory:"false" json:"failureThreshold"`

	// The number of seconds between challenges from the same IP address. If unspecified, defaults to `60`.
	ActionExpirationInSeconds *int `mandatory:"false" json:"actionExpirationInSeconds"`

	// Adds an additional HTTP header to requests that fail the challenge before being passed to the origin. Only applicable when the `action` is set to `DETECT`.
	SetHttpHeader *Header `mandatory:"false" json:"setHttpHeader"`

	ChallengeSettings *BlockChallengeSettings `mandatory:"false" json:"challengeSettings"`

	// When enabled, redirect responses from the origin will also be challenged. This will change HTTP 301/302 responses from origin to HTTP 200 with an HTML body containing JavaScript page redirection.
	AreRedirectsChallenged *bool `mandatory:"false" json:"areRedirectsChallenged"`

	// When defined, the JavaScript Challenge would be applied only for the requests that matched all the listed conditions.
	Criteria []AccessRuleCriteria `mandatory:"false" json:"criteria"`

	// When enabled, the user is identified not only by the IP address but also by an unique additional hash, which prevents blocking visitors with shared IP addresses.
	IsNatEnabled *bool `mandatory:"false" json:"isNatEnabled"`
}

func (m JsChallenge) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JsChallenge) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingJsChallengeActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetJsChallengeActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// JsChallengeActionEnum Enum with underlying type: string
type JsChallengeActionEnum string

// Set of constants representing the allowable values for JsChallengeActionEnum
const (
	JsChallengeActionDetect JsChallengeActionEnum = "DETECT"
	JsChallengeActionBlock  JsChallengeActionEnum = "BLOCK"
)

var mappingJsChallengeActionEnum = map[string]JsChallengeActionEnum{
	"DETECT": JsChallengeActionDetect,
	"BLOCK":  JsChallengeActionBlock,
}

// GetJsChallengeActionEnumValues Enumerates the set of values for JsChallengeActionEnum
func GetJsChallengeActionEnumValues() []JsChallengeActionEnum {
	values := make([]JsChallengeActionEnum, 0)
	for _, v := range mappingJsChallengeActionEnum {
		values = append(values, v)
	}
	return values
}

// GetJsChallengeActionEnumStringValues Enumerates the set of values in String for JsChallengeActionEnum
func GetJsChallengeActionEnumStringValues() []string {
	return []string{
		"DETECT",
		"BLOCK",
	}
}

// GetMappingJsChallengeActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJsChallengeActionEnum(val string) (JsChallengeActionEnum, bool) {
	mappingJsChallengeActionEnumIgnoreCase := make(map[string]JsChallengeActionEnum)
	for k, v := range mappingJsChallengeActionEnum {
		mappingJsChallengeActionEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingJsChallengeActionEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
