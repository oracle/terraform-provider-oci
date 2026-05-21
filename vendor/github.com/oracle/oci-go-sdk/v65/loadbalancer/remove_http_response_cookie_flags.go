// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RemoveHttpResponseCookieFlags An object that represents the action of removing flags from response cookies that are backend generated. This rule
// applies only to HTTP listeners. This will clear existing flags on cookies set by backends. To modify existing flags,
// you first need to have this REMOVE_HTTP_RESPONSE_COOKIE_FLAGS rule for the flag and also an
// ADD_HTTP_RESPONSE_COOKIE_FLAGS to set the flag to the desired value.
type RemoveHttpResponseCookieFlags struct {

	// A list of cookie names that conforms to RFC 6265. In case it is not specified these flags will be applied by
	// default to all cookies. Additional flags can be set by adding specific rule for concrete cookie.
	// Example: ["example_cookie1", "example_cookie2"]
	MatchCookies []string `mandatory:"false" json:"matchCookies"`

	// The flags to be set on cookies.
	Flags []RemoveHttpResponseCookieFlagsFlagsEnum `mandatory:"true" json:"flags"`
}

func (m RemoveHttpResponseCookieFlags) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RemoveHttpResponseCookieFlags) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.Flags {
		if _, ok := GetMappingRemoveHttpResponseCookieFlagsFlagsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Flags: %s. Supported values are: %s.", val, strings.Join(GetRemoveHttpResponseCookieFlagsFlagsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RemoveHttpResponseCookieFlags) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRemoveHttpResponseCookieFlags RemoveHttpResponseCookieFlags
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeRemoveHttpResponseCookieFlags
	}{
		"REMOVE_HTTP_RESPONSE_COOKIE_FLAGS",
		(MarshalTypeRemoveHttpResponseCookieFlags)(m),
	}

	return json.Marshal(&s)
}

// RemoveHttpResponseCookieFlagsFlagsEnum Enum with underlying type: string
type RemoveHttpResponseCookieFlagsFlagsEnum string

// Set of constants representing the allowable values for RemoveHttpResponseCookieFlagsFlagsEnum
const (
	RemoveHttpResponseCookieFlagsFlagsSecure   RemoveHttpResponseCookieFlagsFlagsEnum = "Secure"
	RemoveHttpResponseCookieFlagsFlagsHttponly RemoveHttpResponseCookieFlagsFlagsEnum = "HttpOnly"
	RemoveHttpResponseCookieFlagsFlagsSamesite RemoveHttpResponseCookieFlagsFlagsEnum = "SameSite"
)

var mappingRemoveHttpResponseCookieFlagsFlagsEnum = map[string]RemoveHttpResponseCookieFlagsFlagsEnum{
	"Secure":   RemoveHttpResponseCookieFlagsFlagsSecure,
	"HttpOnly": RemoveHttpResponseCookieFlagsFlagsHttponly,
	"SameSite": RemoveHttpResponseCookieFlagsFlagsSamesite,
}

var mappingRemoveHttpResponseCookieFlagsFlagsEnumLowerCase = map[string]RemoveHttpResponseCookieFlagsFlagsEnum{
	"secure":   RemoveHttpResponseCookieFlagsFlagsSecure,
	"httponly": RemoveHttpResponseCookieFlagsFlagsHttponly,
	"samesite": RemoveHttpResponseCookieFlagsFlagsSamesite,
}

// GetRemoveHttpResponseCookieFlagsFlagsEnumValues Enumerates the set of values for RemoveHttpResponseCookieFlagsFlagsEnum
func GetRemoveHttpResponseCookieFlagsFlagsEnumValues() []RemoveHttpResponseCookieFlagsFlagsEnum {
	values := make([]RemoveHttpResponseCookieFlagsFlagsEnum, 0)
	for _, v := range mappingRemoveHttpResponseCookieFlagsFlagsEnum {
		values = append(values, v)
	}
	return values
}

// GetRemoveHttpResponseCookieFlagsFlagsEnumStringValues Enumerates the set of values in String for RemoveHttpResponseCookieFlagsFlagsEnum
func GetRemoveHttpResponseCookieFlagsFlagsEnumStringValues() []string {
	return []string{
		"Secure",
		"HttpOnly",
		"SameSite",
	}
}

// GetMappingRemoveHttpResponseCookieFlagsFlagsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRemoveHttpResponseCookieFlagsFlagsEnum(val string) (RemoveHttpResponseCookieFlagsFlagsEnum, bool) {
	enum, ok := mappingRemoveHttpResponseCookieFlagsFlagsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
