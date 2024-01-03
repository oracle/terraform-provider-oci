// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UserDetails The details of a particular user.
type UserDetails struct {

	// The name of the user.
	Name *string `mandatory:"false" json:"name"`

	// The status of the user account.
	Status *string `mandatory:"false" json:"status"`

	// The name of the profile assigned to the user.
	Profile *string `mandatory:"false" json:"profile"`

	// The default tablespace of the user.
	Tablespace *string `mandatory:"false" json:"tablespace"`

	// Indicates whether or not the user is predefined by ORACLE.
	IsUserPredefinedByOracle *bool `mandatory:"false" json:"isUserPredefinedByOracle"`

	// The authentication type of the user.
	AuthenticationType UserDetailsAuthenticationTypeEnum `mandatory:"false" json:"authenticationType,omitempty"`
}

func (m UserDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUserDetailsAuthenticationTypeEnum(string(m.AuthenticationType)); !ok && m.AuthenticationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticationType: %s. Supported values are: %s.", m.AuthenticationType, strings.Join(GetUserDetailsAuthenticationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UserDetailsAuthenticationTypeEnum Enum with underlying type: string
type UserDetailsAuthenticationTypeEnum string

// Set of constants representing the allowable values for UserDetailsAuthenticationTypeEnum
const (
	UserDetailsAuthenticationTypePassword UserDetailsAuthenticationTypeEnum = "PASSWORD"
	UserDetailsAuthenticationTypeNone     UserDetailsAuthenticationTypeEnum = "NONE"
)

var mappingUserDetailsAuthenticationTypeEnum = map[string]UserDetailsAuthenticationTypeEnum{
	"PASSWORD": UserDetailsAuthenticationTypePassword,
	"NONE":     UserDetailsAuthenticationTypeNone,
}

var mappingUserDetailsAuthenticationTypeEnumLowerCase = map[string]UserDetailsAuthenticationTypeEnum{
	"password": UserDetailsAuthenticationTypePassword,
	"none":     UserDetailsAuthenticationTypeNone,
}

// GetUserDetailsAuthenticationTypeEnumValues Enumerates the set of values for UserDetailsAuthenticationTypeEnum
func GetUserDetailsAuthenticationTypeEnumValues() []UserDetailsAuthenticationTypeEnum {
	values := make([]UserDetailsAuthenticationTypeEnum, 0)
	for _, v := range mappingUserDetailsAuthenticationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUserDetailsAuthenticationTypeEnumStringValues Enumerates the set of values in String for UserDetailsAuthenticationTypeEnum
func GetUserDetailsAuthenticationTypeEnumStringValues() []string {
	return []string{
		"PASSWORD",
		"NONE",
	}
}

// GetMappingUserDetailsAuthenticationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserDetailsAuthenticationTypeEnum(val string) (UserDetailsAuthenticationTypeEnum, bool) {
	enum, ok := mappingUserDetailsAuthenticationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
