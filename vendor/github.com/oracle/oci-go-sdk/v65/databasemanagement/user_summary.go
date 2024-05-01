// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UserSummary The summary of a specific User.
type UserSummary struct {

	// The name of the User.
	Name *string `mandatory:"true" json:"name"`

	// The status of the user account.
	Status UserSummaryStatusEnum `mandatory:"true" json:"status"`

	// The default tablespace for data.
	DefaultTablespace *string `mandatory:"true" json:"defaultTablespace"`

	// The name of the default tablespace for temporary tables or the name of a tablespace group.
	TempTablespace *string `mandatory:"true" json:"tempTablespace"`

	// The date and time the user was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The profile name of the user.
	Profile *string `mandatory:"true" json:"profile"`

	// The date and time of the expiration of the user account.
	TimeExpiring *common.SDKTime `mandatory:"false" json:"timeExpiring"`

	// The date the account was locked, if the status of the account is LOCKED.
	TimeLocked *common.SDKTime `mandatory:"false" json:"timeLocked"`
}

func (m UserSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUserSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetUserSummaryStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UserSummaryStatusEnum Enum with underlying type: string
type UserSummaryStatusEnum string

// Set of constants representing the allowable values for UserSummaryStatusEnum
const (
	UserSummaryStatusOpen                          UserSummaryStatusEnum = "OPEN"
	UserSummaryStatusExpired                       UserSummaryStatusEnum = "EXPIRED"
	UserSummaryStatusExpiredGrace                  UserSummaryStatusEnum = "EXPIRED_GRACE"
	UserSummaryStatusLocked                        UserSummaryStatusEnum = "LOCKED"
	UserSummaryStatusLockedTimed                   UserSummaryStatusEnum = "LOCKED_TIMED"
	UserSummaryStatusExpiredAndLocked              UserSummaryStatusEnum = "EXPIRED_AND_LOCKED"
	UserSummaryStatusExpiredGraceAndLocked         UserSummaryStatusEnum = "EXPIRED_GRACE_AND_LOCKED"
	UserSummaryStatusExpiredAndLockedTimed         UserSummaryStatusEnum = "EXPIRED_AND_LOCKED_TIMED"
	UserSummaryStatusExpiredGraceAndLockedTimed    UserSummaryStatusEnum = "EXPIRED_GRACE_AND_LOCKED_TIMED"
	UserSummaryStatusOpenAndInRollover             UserSummaryStatusEnum = "OPEN_AND_IN_ROLLOVER"
	UserSummaryStatusExpiredAndInRollover          UserSummaryStatusEnum = "EXPIRED_AND_IN_ROLLOVER"
	UserSummaryStatusLockedAndInRollover           UserSummaryStatusEnum = "LOCKED_AND_IN_ROLLOVER"
	UserSummaryStatusExpiredAndLockedAndInRollover UserSummaryStatusEnum = "EXPIRED_AND_LOCKED_AND_IN_ROLLOVER"
	UserSummaryStatusLockedTimedAndInRollover      UserSummaryStatusEnum = "LOCKED_TIMED_AND_IN_ROLLOVER"
	UserSummaryStatusExpiredAndLockedTimedAndInRol UserSummaryStatusEnum = "EXPIRED_AND_LOCKED_TIMED_AND_IN_ROL"
)

var mappingUserSummaryStatusEnum = map[string]UserSummaryStatusEnum{
	"OPEN":                                UserSummaryStatusOpen,
	"EXPIRED":                             UserSummaryStatusExpired,
	"EXPIRED_GRACE":                       UserSummaryStatusExpiredGrace,
	"LOCKED":                              UserSummaryStatusLocked,
	"LOCKED_TIMED":                        UserSummaryStatusLockedTimed,
	"EXPIRED_AND_LOCKED":                  UserSummaryStatusExpiredAndLocked,
	"EXPIRED_GRACE_AND_LOCKED":            UserSummaryStatusExpiredGraceAndLocked,
	"EXPIRED_AND_LOCKED_TIMED":            UserSummaryStatusExpiredAndLockedTimed,
	"EXPIRED_GRACE_AND_LOCKED_TIMED":      UserSummaryStatusExpiredGraceAndLockedTimed,
	"OPEN_AND_IN_ROLLOVER":                UserSummaryStatusOpenAndInRollover,
	"EXPIRED_AND_IN_ROLLOVER":             UserSummaryStatusExpiredAndInRollover,
	"LOCKED_AND_IN_ROLLOVER":              UserSummaryStatusLockedAndInRollover,
	"EXPIRED_AND_LOCKED_AND_IN_ROLLOVER":  UserSummaryStatusExpiredAndLockedAndInRollover,
	"LOCKED_TIMED_AND_IN_ROLLOVER":        UserSummaryStatusLockedTimedAndInRollover,
	"EXPIRED_AND_LOCKED_TIMED_AND_IN_ROL": UserSummaryStatusExpiredAndLockedTimedAndInRol,
}

var mappingUserSummaryStatusEnumLowerCase = map[string]UserSummaryStatusEnum{
	"open":                                UserSummaryStatusOpen,
	"expired":                             UserSummaryStatusExpired,
	"expired_grace":                       UserSummaryStatusExpiredGrace,
	"locked":                              UserSummaryStatusLocked,
	"locked_timed":                        UserSummaryStatusLockedTimed,
	"expired_and_locked":                  UserSummaryStatusExpiredAndLocked,
	"expired_grace_and_locked":            UserSummaryStatusExpiredGraceAndLocked,
	"expired_and_locked_timed":            UserSummaryStatusExpiredAndLockedTimed,
	"expired_grace_and_locked_timed":      UserSummaryStatusExpiredGraceAndLockedTimed,
	"open_and_in_rollover":                UserSummaryStatusOpenAndInRollover,
	"expired_and_in_rollover":             UserSummaryStatusExpiredAndInRollover,
	"locked_and_in_rollover":              UserSummaryStatusLockedAndInRollover,
	"expired_and_locked_and_in_rollover":  UserSummaryStatusExpiredAndLockedAndInRollover,
	"locked_timed_and_in_rollover":        UserSummaryStatusLockedTimedAndInRollover,
	"expired_and_locked_timed_and_in_rol": UserSummaryStatusExpiredAndLockedTimedAndInRol,
}

// GetUserSummaryStatusEnumValues Enumerates the set of values for UserSummaryStatusEnum
func GetUserSummaryStatusEnumValues() []UserSummaryStatusEnum {
	values := make([]UserSummaryStatusEnum, 0)
	for _, v := range mappingUserSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetUserSummaryStatusEnumStringValues Enumerates the set of values in String for UserSummaryStatusEnum
func GetUserSummaryStatusEnumStringValues() []string {
	return []string{
		"OPEN",
		"EXPIRED",
		"EXPIRED_GRACE",
		"LOCKED",
		"LOCKED_TIMED",
		"EXPIRED_AND_LOCKED",
		"EXPIRED_GRACE_AND_LOCKED",
		"EXPIRED_AND_LOCKED_TIMED",
		"EXPIRED_GRACE_AND_LOCKED_TIMED",
		"OPEN_AND_IN_ROLLOVER",
		"EXPIRED_AND_IN_ROLLOVER",
		"LOCKED_AND_IN_ROLLOVER",
		"EXPIRED_AND_LOCKED_AND_IN_ROLLOVER",
		"LOCKED_TIMED_AND_IN_ROLLOVER",
		"EXPIRED_AND_LOCKED_TIMED_AND_IN_ROL",
	}
}

// GetMappingUserSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserSummaryStatusEnum(val string) (UserSummaryStatusEnum, bool) {
	enum, ok := mappingUserSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
