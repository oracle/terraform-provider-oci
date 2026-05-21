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

// AddHttpResponseCookieFlags An object that represents the action of setting flags on response cookies that are backend generated. This rule
// applies only to HTTP listeners. This won't clear existing flags on cookies set by backend, instead it will set
// additional flags if they are not present. To modify existing flags, you first need to have a corresponding
// REMOVE_HTTP_RESPONSE_COOKIE_FLAGS rule for the flag. The REMOVE rules are processed first, so the existing flag
// will be removed and then the new flag will be added
type AddHttpResponseCookieFlags struct {

	// A list of cookie names that conform to RFC 6265. In case it is not specified these flags will be applied by
	// default to all cookies. Additional flags can be set by adding specific rule for concrete cookie.
	// Example: ["example_cookie1", "example_cookie2"]
	MatchCookies []string `mandatory:"false" json:"matchCookies"`

	// The flags to be set on cookies.
	Flags []AddHttpResponseCookieFlagsFlagsEnum `mandatory:"true" json:"flags"`
}

func (m AddHttpResponseCookieFlags) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddHttpResponseCookieFlags) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.Flags {
		if _, ok := GetMappingAddHttpResponseCookieFlagsFlagsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Flags: %s. Supported values are: %s.", val, strings.Join(GetAddHttpResponseCookieFlagsFlagsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AddHttpResponseCookieFlags) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAddHttpResponseCookieFlags AddHttpResponseCookieFlags
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeAddHttpResponseCookieFlags
	}{
		"ADD_HTTP_RESPONSE_COOKIE_FLAGS",
		(MarshalTypeAddHttpResponseCookieFlags)(m),
	}

	return json.Marshal(&s)
}

// AddHttpResponseCookieFlagsFlagsEnum Enum with underlying type: string
type AddHttpResponseCookieFlagsFlagsEnum string

// Set of constants representing the allowable values for AddHttpResponseCookieFlagsFlagsEnum
const (
	AddHttpResponseCookieFlagsFlagsSecure         AddHttpResponseCookieFlagsFlagsEnum = "Secure"
	AddHttpResponseCookieFlagsFlagsHttponly       AddHttpResponseCookieFlagsFlagsEnum = "HttpOnly"
	AddHttpResponseCookieFlagsFlagsSamesiteLax    AddHttpResponseCookieFlagsFlagsEnum = "SameSite_Lax"
	AddHttpResponseCookieFlagsFlagsSamesiteNone   AddHttpResponseCookieFlagsFlagsEnum = "SameSite_None"
	AddHttpResponseCookieFlagsFlagsSamesiteStrict AddHttpResponseCookieFlagsFlagsEnum = "SameSite_Strict"
)

var mappingAddHttpResponseCookieFlagsFlagsEnum = map[string]AddHttpResponseCookieFlagsFlagsEnum{
	"Secure":          AddHttpResponseCookieFlagsFlagsSecure,
	"HttpOnly":        AddHttpResponseCookieFlagsFlagsHttponly,
	"SameSite_Lax":    AddHttpResponseCookieFlagsFlagsSamesiteLax,
	"SameSite_None":   AddHttpResponseCookieFlagsFlagsSamesiteNone,
	"SameSite_Strict": AddHttpResponseCookieFlagsFlagsSamesiteStrict,
}

var mappingAddHttpResponseCookieFlagsFlagsEnumLowerCase = map[string]AddHttpResponseCookieFlagsFlagsEnum{
	"secure":          AddHttpResponseCookieFlagsFlagsSecure,
	"httponly":        AddHttpResponseCookieFlagsFlagsHttponly,
	"samesite_lax":    AddHttpResponseCookieFlagsFlagsSamesiteLax,
	"samesite_none":   AddHttpResponseCookieFlagsFlagsSamesiteNone,
	"samesite_strict": AddHttpResponseCookieFlagsFlagsSamesiteStrict,
}

// GetAddHttpResponseCookieFlagsFlagsEnumValues Enumerates the set of values for AddHttpResponseCookieFlagsFlagsEnum
func GetAddHttpResponseCookieFlagsFlagsEnumValues() []AddHttpResponseCookieFlagsFlagsEnum {
	values := make([]AddHttpResponseCookieFlagsFlagsEnum, 0)
	for _, v := range mappingAddHttpResponseCookieFlagsFlagsEnum {
		values = append(values, v)
	}
	return values
}

// GetAddHttpResponseCookieFlagsFlagsEnumStringValues Enumerates the set of values in String for AddHttpResponseCookieFlagsFlagsEnum
func GetAddHttpResponseCookieFlagsFlagsEnumStringValues() []string {
	return []string{
		"Secure",
		"HttpOnly",
		"SameSite_Lax",
		"SameSite_None",
		"SameSite_Strict",
	}
}

// GetMappingAddHttpResponseCookieFlagsFlagsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddHttpResponseCookieFlagsFlagsEnum(val string) (AddHttpResponseCookieFlagsFlagsEnum, bool) {
	enum, ok := mappingAddHttpResponseCookieFlagsFlagsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
