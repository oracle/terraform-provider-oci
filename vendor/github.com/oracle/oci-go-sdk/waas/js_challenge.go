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

// JsChallengeActionEnum Enum with underlying type: string
type JsChallengeActionEnum string

// Set of constants representing the allowable values for JsChallengeActionEnum
const (
	JsChallengeActionDetect JsChallengeActionEnum = "DETECT"
	JsChallengeActionBlock  JsChallengeActionEnum = "BLOCK"
)

var mappingJsChallengeAction = map[string]JsChallengeActionEnum{
	"DETECT": JsChallengeActionDetect,
	"BLOCK":  JsChallengeActionBlock,
}

// GetJsChallengeActionEnumValues Enumerates the set of values for JsChallengeActionEnum
func GetJsChallengeActionEnumValues() []JsChallengeActionEnum {
	values := make([]JsChallengeActionEnum, 0)
	for _, v := range mappingJsChallengeAction {
		values = append(values, v)
	}
	return values
}
