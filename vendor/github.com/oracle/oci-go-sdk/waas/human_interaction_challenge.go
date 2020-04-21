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

// HumanInteractionChallenge The human interaction challenge settings. The human interaction challenge checks various event listeners in the user's browser to determine if there is a human user making a request.
type HumanInteractionChallenge struct {

	// Enables or disables the human interaction challenge Web Application Firewall feature.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The action to take against requests from detected bots. If unspecified, defaults to `DETECT`.
	Action HumanInteractionChallengeActionEnum `mandatory:"false" json:"action,omitempty"`

	// The number of failed requests before taking action. If unspecified, defaults to `10`.
	FailureThreshold *int `mandatory:"false" json:"failureThreshold"`

	// The number of seconds between challenges for the same IP address. If unspecified, defaults to `60`.
	ActionExpirationInSeconds *int `mandatory:"false" json:"actionExpirationInSeconds"`

	// The number of seconds before the failure threshold resets. If unspecified, defaults to  `60`.
	FailureThresholdExpirationInSeconds *int `mandatory:"false" json:"failureThresholdExpirationInSeconds"`

	// The number of interactions required to pass the challenge. If unspecified, defaults to `3`.
	InteractionThreshold *int `mandatory:"false" json:"interactionThreshold"`

	// The number of seconds to record the interactions from the user. If unspecified, defaults to `15`.
	RecordingPeriodInSeconds *int `mandatory:"false" json:"recordingPeriodInSeconds"`

	// Adds an additional HTTP header to requests that fail the challenge before being passed to the origin. Only applicable when the `action` is set to `DETECT`.
	SetHttpHeader *Header `mandatory:"false" json:"setHttpHeader"`

	ChallengeSettings *BlockChallengeSettings `mandatory:"false" json:"challengeSettings"`

	// When enabled, the user is identified not only by the IP address but also by an unique additional hash, which prevents blocking visitors with shared IP addresses.
	IsNatEnabled *bool `mandatory:"false" json:"isNatEnabled"`
}

func (m HumanInteractionChallenge) String() string {
	return common.PointerString(m)
}

// HumanInteractionChallengeActionEnum Enum with underlying type: string
type HumanInteractionChallengeActionEnum string

// Set of constants representing the allowable values for HumanInteractionChallengeActionEnum
const (
	HumanInteractionChallengeActionDetect HumanInteractionChallengeActionEnum = "DETECT"
	HumanInteractionChallengeActionBlock  HumanInteractionChallengeActionEnum = "BLOCK"
)

var mappingHumanInteractionChallengeAction = map[string]HumanInteractionChallengeActionEnum{
	"DETECT": HumanInteractionChallengeActionDetect,
	"BLOCK":  HumanInteractionChallengeActionBlock,
}

// GetHumanInteractionChallengeActionEnumValues Enumerates the set of values for HumanInteractionChallengeActionEnum
func GetHumanInteractionChallengeActionEnumValues() []HumanInteractionChallengeActionEnum {
	values := make([]HumanInteractionChallengeActionEnum, 0)
	for _, v := range mappingHumanInteractionChallengeAction {
		values = append(values, v)
	}
	return values
}
