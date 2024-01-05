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

// ProtectionRuleExclusion Allows specified types of requests to bypass the protection rule. If a request matches any of the criteria in the `exclusions` field, the protection rule will not be executed. Rules can have more than one exclusion and exclusions are applied to requests disjunctively, meaning the specified exclusion strings are independently matched against the specified targets of a request. The first target to match a specified string will trigger an exclusion. **Example:** If the following exclusions are defined for a protection rule:
//
//	"action": "BLOCK",
//	"exclusions": [
//	    {
//	        "target":"REQUEST_COOKIES",
//	        "exclusions":["example.com", "12345", "219ffwef9w0f"]
//	    },
//	    {
//	        "target":"REQUEST_COOKIE_NAMES",
//	        "exclusions":["OAMAuthnCookie", "JSESSIONID", "HCM-PSJSESSIONID"]
//	    }
//	],
//	"key": "1000000",
//
// A request with the cookie name `sessionid` would trigger an exclusion. A request with the cookie name `yourcompany.com` would *not* trigger and exclusion.
type ProtectionRuleExclusion struct {

	// The target of the exclusion.
	Target ProtectionRuleExclusionTargetEnum `mandatory:"false" json:"target,omitempty"`

	Exclusions []string `mandatory:"false" json:"exclusions"`
}

func (m ProtectionRuleExclusion) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProtectionRuleExclusion) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingProtectionRuleExclusionTargetEnum(string(m.Target)); !ok && m.Target != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Target: %s. Supported values are: %s.", m.Target, strings.Join(GetProtectionRuleExclusionTargetEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingProtectionRuleExclusionTargetEnum = map[string]ProtectionRuleExclusionTargetEnum{
	"REQUEST_COOKIES":      ProtectionRuleExclusionTargetRequestCookies,
	"REQUEST_COOKIE_NAMES": ProtectionRuleExclusionTargetRequestCookieNames,
	"ARGS":                 ProtectionRuleExclusionTargetArgs,
	"ARGS_NAMES":           ProtectionRuleExclusionTargetArgsNames,
}

var mappingProtectionRuleExclusionTargetEnumLowerCase = map[string]ProtectionRuleExclusionTargetEnum{
	"request_cookies":      ProtectionRuleExclusionTargetRequestCookies,
	"request_cookie_names": ProtectionRuleExclusionTargetRequestCookieNames,
	"args":                 ProtectionRuleExclusionTargetArgs,
	"args_names":           ProtectionRuleExclusionTargetArgsNames,
}

// GetProtectionRuleExclusionTargetEnumValues Enumerates the set of values for ProtectionRuleExclusionTargetEnum
func GetProtectionRuleExclusionTargetEnumValues() []ProtectionRuleExclusionTargetEnum {
	values := make([]ProtectionRuleExclusionTargetEnum, 0)
	for _, v := range mappingProtectionRuleExclusionTargetEnum {
		values = append(values, v)
	}
	return values
}

// GetProtectionRuleExclusionTargetEnumStringValues Enumerates the set of values in String for ProtectionRuleExclusionTargetEnum
func GetProtectionRuleExclusionTargetEnumStringValues() []string {
	return []string{
		"REQUEST_COOKIES",
		"REQUEST_COOKIE_NAMES",
		"ARGS",
		"ARGS_NAMES",
	}
}

// GetMappingProtectionRuleExclusionTargetEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProtectionRuleExclusionTargetEnum(val string) (ProtectionRuleExclusionTargetEnum, bool) {
	enum, ok := mappingProtectionRuleExclusionTargetEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
