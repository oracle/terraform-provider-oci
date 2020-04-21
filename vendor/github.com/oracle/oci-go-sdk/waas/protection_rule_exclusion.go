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

// ProtectionRuleExclusion Allows specified types of requests to bypass the protection rule. If a request matches any of the criteria in the `exclusions` field, the protection rule will not be executed. Rules can have more than one exclusion and exclusions are applied to requests disjunctively, meaning the specified exclusion strings are independently matched against the specified targets of a request. The first target to match a specified string will trigger an exclusion. **Example:** If the following exclusions are defined for a protection rule:
//     "action": "BLOCK",
//     "exclusions": [
//         {
//             "target":"REQUEST_COOKIES",
//             "exclusions":["example.com", "12345", "219ffwef9w0f"]
//         },
//         {
//             "target":"REQUEST_COOKIE_NAMES",
//             "exclusions":["OAMAuthnCookie", "JSESSIONID", "HCM-PSJSESSIONID"]
//         }
//     ],
//     "key": "1000000",
// A request with the cookie name `sessionid` would trigger an exclusion. A request with the cookie name `yourcompany.com` would *not* trigger and exclusion.
type ProtectionRuleExclusion struct {

	// The target of the exclusion.
	Target ProtectionRuleExclusionTargetEnum `mandatory:"false" json:"target,omitempty"`

	Exclusions []string `mandatory:"false" json:"exclusions"`
}

func (m ProtectionRuleExclusion) String() string {
	return common.PointerString(m)
}

// ProtectionRuleExclusionTargetEnum Enum with underlying type: string
type ProtectionRuleExclusionTargetEnum string

// Set of constants representing the allowable values for ProtectionRuleExclusionTargetEnum
const (
	ProtectionRuleExclusionTargetRequestCookies     ProtectionRuleExclusionTargetEnum = "REQUEST_COOKIES"
	ProtectionRuleExclusionTargetRequestCookieNames ProtectionRuleExclusionTargetEnum = "REQUEST_COOKIE_NAMES"
	ProtectionRuleExclusionTargetArgs               ProtectionRuleExclusionTargetEnum = "ARGS"
	ProtectionRuleExclusionTargetArgsNames          ProtectionRuleExclusionTargetEnum = "ARGS_NAMES"
)

var mappingProtectionRuleExclusionTarget = map[string]ProtectionRuleExclusionTargetEnum{
	"REQUEST_COOKIES":      ProtectionRuleExclusionTargetRequestCookies,
	"REQUEST_COOKIE_NAMES": ProtectionRuleExclusionTargetRequestCookieNames,
	"ARGS":                 ProtectionRuleExclusionTargetArgs,
	"ARGS_NAMES":           ProtectionRuleExclusionTargetArgsNames,
}

// GetProtectionRuleExclusionTargetEnumValues Enumerates the set of values for ProtectionRuleExclusionTargetEnum
func GetProtectionRuleExclusionTargetEnumValues() []ProtectionRuleExclusionTargetEnum {
	values := make([]ProtectionRuleExclusionTargetEnum, 0)
	for _, v := range mappingProtectionRuleExclusionTarget {
		values = append(values, v)
	}
	return values
}
