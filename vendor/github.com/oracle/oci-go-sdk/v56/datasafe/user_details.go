// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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

// UserDetailsAuthenticationTypeEnum Enum with underlying type: string
type UserDetailsAuthenticationTypeEnum string

// Set of constants representing the allowable values for UserDetailsAuthenticationTypeEnum
const (
	UserDetailsAuthenticationTypePassword UserDetailsAuthenticationTypeEnum = "PASSWORD"
	UserDetailsAuthenticationTypeNone     UserDetailsAuthenticationTypeEnum = "NONE"
)

var mappingUserDetailsAuthenticationType = map[string]UserDetailsAuthenticationTypeEnum{
	"PASSWORD": UserDetailsAuthenticationTypePassword,
	"NONE":     UserDetailsAuthenticationTypeNone,
}

// GetUserDetailsAuthenticationTypeEnumValues Enumerates the set of values for UserDetailsAuthenticationTypeEnum
func GetUserDetailsAuthenticationTypeEnumValues() []UserDetailsAuthenticationTypeEnum {
	values := make([]UserDetailsAuthenticationTypeEnum, 0)
	for _, v := range mappingUserDetailsAuthenticationType {
		values = append(values, v)
	}
	return values
}
