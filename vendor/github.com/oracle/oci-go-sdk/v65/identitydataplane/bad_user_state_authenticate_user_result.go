// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Data Plane API
//
// APIs for managing identity data plane services. For example, use this API to create a scoped-access security token. To manage identity domains (for example, creating or deleting an identity domain) or to manage resources (for example, users and groups) within the default identity domain, see IAM API (https://docs.oracle.com/iaas/api/#/en/identity/).
//

package identitydataplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BadUserStateAuthenticateUserResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBadUserStateAuthenticateUserResultUserStateEnum(string(m.UserState)); !ok && m.UserState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UserState: %s. Supported values are: %s.", m.UserState, strings.Join(GetBadUserStateAuthenticateUserResultUserStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingBadUserStateAuthenticateUserResultUserStateEnum = map[string]BadUserStateAuthenticateUserResultUserStateEnum{
	"USER_BLOCKED":              BadUserStateAuthenticateUserResultUserStateUserBlocked,
	"USER_DISABLED":             BadUserStateAuthenticateUserResultUserStateUserDisabled,
	"ONE_TIME_PASSWORD_EXPIRED": BadUserStateAuthenticateUserResultUserStateOneTimePasswordExpired,
	"PASSWORD_INVALID":          BadUserStateAuthenticateUserResultUserStatePasswordInvalid,
}

var mappingBadUserStateAuthenticateUserResultUserStateEnumLowerCase = map[string]BadUserStateAuthenticateUserResultUserStateEnum{
	"user_blocked":              BadUserStateAuthenticateUserResultUserStateUserBlocked,
	"user_disabled":             BadUserStateAuthenticateUserResultUserStateUserDisabled,
	"one_time_password_expired": BadUserStateAuthenticateUserResultUserStateOneTimePasswordExpired,
	"password_invalid":          BadUserStateAuthenticateUserResultUserStatePasswordInvalid,
}

// GetBadUserStateAuthenticateUserResultUserStateEnumValues Enumerates the set of values for BadUserStateAuthenticateUserResultUserStateEnum
func GetBadUserStateAuthenticateUserResultUserStateEnumValues() []BadUserStateAuthenticateUserResultUserStateEnum {
	values := make([]BadUserStateAuthenticateUserResultUserStateEnum, 0)
	for _, v := range mappingBadUserStateAuthenticateUserResultUserStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBadUserStateAuthenticateUserResultUserStateEnumStringValues Enumerates the set of values in String for BadUserStateAuthenticateUserResultUserStateEnum
func GetBadUserStateAuthenticateUserResultUserStateEnumStringValues() []string {
	return []string{
		"USER_BLOCKED",
		"USER_DISABLED",
		"ONE_TIME_PASSWORD_EXPIRED",
		"PASSWORD_INVALID",
	}
}

// GetMappingBadUserStateAuthenticateUserResultUserStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBadUserStateAuthenticateUserResultUserStateEnum(val string) (BadUserStateAuthenticateUserResultUserStateEnum, bool) {
	enum, ok := mappingBadUserStateAuthenticateUserResultUserStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
