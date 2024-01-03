// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// DeviceFingerprintChallenge The device fingerprint challenge settings. The device fingerprint challenge generates hashed signatures of both virtual and real browsers to identify and block malicious bots.
type DeviceFingerprintChallenge struct {

	// Enables or disables the device fingerprint challenge Web Application Firewall feature.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The action to take on requests from detected bots. If unspecified, defaults to `DETECT`.
	Action DeviceFingerprintChallengeActionEnum `mandatory:"false" json:"action,omitempty"`

	// The number of failed requests allowed before taking action. If unspecified, defaults to `10`.
	FailureThreshold *int `mandatory:"false" json:"failureThreshold"`

	// The number of seconds between challenges for the same IP address. If unspecified, defaults to `60`.
	ActionExpirationInSeconds *int `mandatory:"false" json:"actionExpirationInSeconds"`

	// The number of seconds before the failure threshold resets. If unspecified, defaults to `60`.
	FailureThresholdExpirationInSeconds *int `mandatory:"false" json:"failureThresholdExpirationInSeconds"`

	// The maximum number of IP addresses permitted with the same device fingerprint. If unspecified, defaults to `20`.
	MaxAddressCount *int `mandatory:"false" json:"maxAddressCount"`

	// The number of seconds before the maximum addresses count resets. If unspecified, defaults to `60`.
	MaxAddressCountExpirationInSeconds *int `mandatory:"false" json:"maxAddressCountExpirationInSeconds"`

	ChallengeSettings *BlockChallengeSettings `mandatory:"false" json:"challengeSettings"`
}

func (m DeviceFingerprintChallenge) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeviceFingerprintChallenge) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDeviceFingerprintChallengeActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetDeviceFingerprintChallengeActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DeviceFingerprintChallengeActionEnum Enum with underlying type: string
type DeviceFingerprintChallengeActionEnum string

// Set of constants representing the allowable values for DeviceFingerprintChallengeActionEnum
const (
	DeviceFingerprintChallengeActionDetect DeviceFingerprintChallengeActionEnum = "DETECT"
	DeviceFingerprintChallengeActionBlock  DeviceFingerprintChallengeActionEnum = "BLOCK"
)

var mappingDeviceFingerprintChallengeActionEnum = map[string]DeviceFingerprintChallengeActionEnum{
	"DETECT": DeviceFingerprintChallengeActionDetect,
	"BLOCK":  DeviceFingerprintChallengeActionBlock,
}

var mappingDeviceFingerprintChallengeActionEnumLowerCase = map[string]DeviceFingerprintChallengeActionEnum{
	"detect": DeviceFingerprintChallengeActionDetect,
	"block":  DeviceFingerprintChallengeActionBlock,
}

// GetDeviceFingerprintChallengeActionEnumValues Enumerates the set of values for DeviceFingerprintChallengeActionEnum
func GetDeviceFingerprintChallengeActionEnumValues() []DeviceFingerprintChallengeActionEnum {
	values := make([]DeviceFingerprintChallengeActionEnum, 0)
	for _, v := range mappingDeviceFingerprintChallengeActionEnum {
		values = append(values, v)
	}
	return values
}

// GetDeviceFingerprintChallengeActionEnumStringValues Enumerates the set of values in String for DeviceFingerprintChallengeActionEnum
func GetDeviceFingerprintChallengeActionEnumStringValues() []string {
	return []string{
		"DETECT",
		"BLOCK",
	}
}

// GetMappingDeviceFingerprintChallengeActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeviceFingerprintChallengeActionEnum(val string) (DeviceFingerprintChallengeActionEnum, bool) {
	enum, ok := mappingDeviceFingerprintChallengeActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
