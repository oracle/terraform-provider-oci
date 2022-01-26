// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Service
//
// API for the Identity Dataplane
//

package identitydataplane

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// BadUserStateAuthenticateUserResult The representation of BadUserStateAuthenticateUserResult
type BadUserStateAuthenticateUserResult struct {

	// The tenant name.
	TenantInput *string `mandatory:"true" json:"tenantInput"`

	// The user name.
	UserInput *string `mandatory:"true" json:"userInput"`

	// The resolved tenant id.
	ResolvedTenantId *string `mandatory:"true" json:"resolvedTenantId"`

	// The resolved user id.
	ResolvedUserId *string `mandatory:"true" json:"resolvedUserId"`

	// The bad user state.
	UserState BadUserStateAuthenticateUserResultUserStateEnum `mandatory:"true" json:"userState"`
}

func (m BadUserStateAuthenticateUserResult) String() string {
	return common.PointerString(m)
}

// BadUserStateAuthenticateUserResultUserStateEnum Enum with underlying type: string
type BadUserStateAuthenticateUserResultUserStateEnum string

// Set of constants representing the allowable values for BadUserStateAuthenticateUserResultUserStateEnum
const (
	BadUserStateAuthenticateUserResultUserStateUserBlocked            BadUserStateAuthenticateUserResultUserStateEnum = "USER_BLOCKED"
	BadUserStateAuthenticateUserResultUserStateUserDisabled           BadUserStateAuthenticateUserResultUserStateEnum = "USER_DISABLED"
	BadUserStateAuthenticateUserResultUserStateOneTimePasswordExpired BadUserStateAuthenticateUserResultUserStateEnum = "ONE_TIME_PASSWORD_EXPIRED"
	BadUserStateAuthenticateUserResultUserStatePasswordInvalid        BadUserStateAuthenticateUserResultUserStateEnum = "PASSWORD_INVALID"
)

var mappingBadUserStateAuthenticateUserResultUserState = map[string]BadUserStateAuthenticateUserResultUserStateEnum{
	"USER_BLOCKED":              BadUserStateAuthenticateUserResultUserStateUserBlocked,
	"USER_DISABLED":             BadUserStateAuthenticateUserResultUserStateUserDisabled,
	"ONE_TIME_PASSWORD_EXPIRED": BadUserStateAuthenticateUserResultUserStateOneTimePasswordExpired,
	"PASSWORD_INVALID":          BadUserStateAuthenticateUserResultUserStatePasswordInvalid,
}

// GetBadUserStateAuthenticateUserResultUserStateEnumValues Enumerates the set of values for BadUserStateAuthenticateUserResultUserStateEnum
func GetBadUserStateAuthenticateUserResultUserStateEnumValues() []BadUserStateAuthenticateUserResultUserStateEnum {
	values := make([]BadUserStateAuthenticateUserResultUserStateEnum, 0)
	for _, v := range mappingBadUserStateAuthenticateUserResultUserState {
		values = append(values, v)
	}
	return values
}
