// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// User The summary of a specific User.
type User struct {

	// The name of the User.
	Name *string `mandatory:"true" json:"name"`

	// The account status of the User
	Status UserStatusEnum `mandatory:"true" json:"status"`

	// The default tablespace for data.
	DefaultTablespace *string `mandatory:"true" json:"defaultTablespace"`

	// The name of the default tablespace for temporary tables or the name of a tablespace group.
	TempTablespace *string `mandatory:"true" json:"tempTablespace"`

	// The User creation date.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The User resource profile name.
	Profile *string `mandatory:"true" json:"profile"`

	// The date the account was locked if account status was LOCKED.
	TimeLocked *common.SDKTime `mandatory:"false" json:"timeLocked"`

	// The date of expiration of the account
	TimeExpiring *common.SDKTime `mandatory:"false" json:"timeExpiring"`

	// The default local temporary tablespace for the User.
	LocalTempTablespace *string `mandatory:"false" json:"localTempTablespace"`

	// The initial resource consumer group for the User.
	ConsumerGroup *string `mandatory:"false" json:"consumerGroup"`

	// The external name of User.
	ExternalName *string `mandatory:"false" json:"externalName"`

	// The list of versions of the password hashes (also known as "verifiers") existing for the account.
	PasswordVersions *string `mandatory:"false" json:"passwordVersions"`

	// Indicates whether editions have been enabled for the corresponding user (Y) or not (N)
	EditionsEnabled UserEditionsEnabledEnum `mandatory:"false" json:"editionsEnabled,omitempty"`

	// The authentication mechanism for the user
	Authentication UserAuthenticationEnum `mandatory:"false" json:"authentication,omitempty"`

	// Indicates whether a user can connect directly (N) or whether the account can only be proxied (Y) by users who have proxy privileges
	// for this account (that is, by users who have been granted the "connect through" privilege for this account).
	ProxyConnect UserProxyConnectEnum `mandatory:"false" json:"proxyConnect,omitempty"`

	// Indicates whether a given user is common(Y) or local(N).
	Common UserCommonEnum `mandatory:"false" json:"common,omitempty"`

	// The time of the last user login.
	// This column is not populated when a user connects to the database with administrative privileges, that is, AS { SYSASM | SYSBACKUP | SYSDBA | SYSDG | SYSOPER | SYSRAC | SYSKM }.
	TimeLastLogin *common.SDKTime `mandatory:"false" json:"timeLastLogin"`

	// Indicates whether the user was created, and is maintained, by Oracle-supplied scripts (such as catalog.sql or catproc.sql).
	OracleMaintained UserOracleMaintainedEnum `mandatory:"false" json:"oracleMaintained,omitempty"`

	// Indicates whether the user definition was inherited from another container (YES) or not (NO)
	Inherited UserInheritedEnum `mandatory:"false" json:"inherited,omitempty"`

	// The default collation for the user’s schema.
	DefaultCollation *string `mandatory:"false" json:"defaultCollation"`

	// Indicates whether this user is a common user created by an implicit application (YES) or not (NO)
	Implicit UserImplicitEnum `mandatory:"false" json:"implicit,omitempty"`

	// In a sharded database, the value in this column indicates whether the user was created with shard DDL enabled.
	AllShared UserAllSharedEnum `mandatory:"false" json:"allShared,omitempty"`

	// In a federated sharded database, indicates whether the user is an external shard user (YES) or not (NO).
	ExternalShared UserExternalSharedEnum `mandatory:"false" json:"externalShared,omitempty"`

	// The date on which the user's password was last set.
	// This column is populated only when the value of the AUTHENTICATION_TYPE column is PASSWORD. Otherwise, this column is null.
	TimePasswordChanged *common.SDKTime `mandatory:"false" json:"timePasswordChanged"`
}

func (m User) String() string {
	return common.PointerString(m)
}

// UserStatusEnum Enum with underlying type: string
type UserStatusEnum string

// Set of constants representing the allowable values for UserStatusEnum
const (
	UserStatusOpen                          UserStatusEnum = "OPEN"
	UserStatusExpired                       UserStatusEnum = "EXPIRED"
	UserStatusExpiredGrace                  UserStatusEnum = "EXPIRED_GRACE"
	UserStatusLocked                        UserStatusEnum = "LOCKED"
	UserStatusLockedTimed                   UserStatusEnum = "LOCKED_TIMED"
	UserStatusExpiredAndLocked              UserStatusEnum = "EXPIRED_AND_LOCKED"
	UserStatusExpiredGraceAndLocked         UserStatusEnum = "EXPIRED_GRACE_AND_LOCKED"
	UserStatusExpiredAndLockedTimed         UserStatusEnum = "EXPIRED_AND_LOCKED_TIMED"
	UserStatusExpiredGraceAndLockedTimed    UserStatusEnum = "EXPIRED_GRACE_AND_LOCKED_TIMED"
	UserStatusOpenAndInRollover             UserStatusEnum = "OPEN_AND_IN_ROLLOVER"
	UserStatusExpiredAndInRollover          UserStatusEnum = "EXPIRED_AND_IN_ROLLOVER"
	UserStatusLockedAndInRollover           UserStatusEnum = "LOCKED_AND_IN_ROLLOVER"
	UserStatusExpiredAndLockedAndInRollover UserStatusEnum = "EXPIRED_AND_LOCKED_AND_IN_ROLLOVER"
	UserStatusLockedTimedAndInRollover      UserStatusEnum = "LOCKED_TIMED_AND_IN_ROLLOVER"
	UserStatusExpiredAndLockedTimedAndInRol UserStatusEnum = "EXPIRED_AND_LOCKED_TIMED_AND_IN_ROL"
)

var mappingUserStatus = map[string]UserStatusEnum{
	"OPEN":                                UserStatusOpen,
	"EXPIRED":                             UserStatusExpired,
	"EXPIRED_GRACE":                       UserStatusExpiredGrace,
	"LOCKED":                              UserStatusLocked,
	"LOCKED_TIMED":                        UserStatusLockedTimed,
	"EXPIRED_AND_LOCKED":                  UserStatusExpiredAndLocked,
	"EXPIRED_GRACE_AND_LOCKED":            UserStatusExpiredGraceAndLocked,
	"EXPIRED_AND_LOCKED_TIMED":            UserStatusExpiredAndLockedTimed,
	"EXPIRED_GRACE_AND_LOCKED_TIMED":      UserStatusExpiredGraceAndLockedTimed,
	"OPEN_AND_IN_ROLLOVER":                UserStatusOpenAndInRollover,
	"EXPIRED_AND_IN_ROLLOVER":             UserStatusExpiredAndInRollover,
	"LOCKED_AND_IN_ROLLOVER":              UserStatusLockedAndInRollover,
	"EXPIRED_AND_LOCKED_AND_IN_ROLLOVER":  UserStatusExpiredAndLockedAndInRollover,
	"LOCKED_TIMED_AND_IN_ROLLOVER":        UserStatusLockedTimedAndInRollover,
	"EXPIRED_AND_LOCKED_TIMED_AND_IN_ROL": UserStatusExpiredAndLockedTimedAndInRol,
}

// GetUserStatusEnumValues Enumerates the set of values for UserStatusEnum
func GetUserStatusEnumValues() []UserStatusEnum {
	values := make([]UserStatusEnum, 0)
	for _, v := range mappingUserStatus {
		values = append(values, v)
	}
	return values
}

// UserEditionsEnabledEnum Enum with underlying type: string
type UserEditionsEnabledEnum string

// Set of constants representing the allowable values for UserEditionsEnabledEnum
const (
	UserEditionsEnabledYes UserEditionsEnabledEnum = "YES"
	UserEditionsEnabledNo  UserEditionsEnabledEnum = "NO"
)

var mappingUserEditionsEnabled = map[string]UserEditionsEnabledEnum{
	"YES": UserEditionsEnabledYes,
	"NO":  UserEditionsEnabledNo,
}

// GetUserEditionsEnabledEnumValues Enumerates the set of values for UserEditionsEnabledEnum
func GetUserEditionsEnabledEnumValues() []UserEditionsEnabledEnum {
	values := make([]UserEditionsEnabledEnum, 0)
	for _, v := range mappingUserEditionsEnabled {
		values = append(values, v)
	}
	return values
}

// UserAuthenticationEnum Enum with underlying type: string
type UserAuthenticationEnum string

// Set of constants representing the allowable values for UserAuthenticationEnum
const (
	UserAuthenticationNone     UserAuthenticationEnum = "NONE"
	UserAuthenticationExternal UserAuthenticationEnum = "EXTERNAL"
	UserAuthenticationGlobal   UserAuthenticationEnum = "GLOBAL"
	UserAuthenticationPassword UserAuthenticationEnum = "PASSWORD"
)

var mappingUserAuthentication = map[string]UserAuthenticationEnum{
	"NONE":     UserAuthenticationNone,
	"EXTERNAL": UserAuthenticationExternal,
	"GLOBAL":   UserAuthenticationGlobal,
	"PASSWORD": UserAuthenticationPassword,
}

// GetUserAuthenticationEnumValues Enumerates the set of values for UserAuthenticationEnum
func GetUserAuthenticationEnumValues() []UserAuthenticationEnum {
	values := make([]UserAuthenticationEnum, 0)
	for _, v := range mappingUserAuthentication {
		values = append(values, v)
	}
	return values
}

// UserProxyConnectEnum Enum with underlying type: string
type UserProxyConnectEnum string

// Set of constants representing the allowable values for UserProxyConnectEnum
const (
	UserProxyConnectYes UserProxyConnectEnum = "YES"
	UserProxyConnectNo  UserProxyConnectEnum = "NO"
)

var mappingUserProxyConnect = map[string]UserProxyConnectEnum{
	"YES": UserProxyConnectYes,
	"NO":  UserProxyConnectNo,
}

// GetUserProxyConnectEnumValues Enumerates the set of values for UserProxyConnectEnum
func GetUserProxyConnectEnumValues() []UserProxyConnectEnum {
	values := make([]UserProxyConnectEnum, 0)
	for _, v := range mappingUserProxyConnect {
		values = append(values, v)
	}
	return values
}

// UserCommonEnum Enum with underlying type: string
type UserCommonEnum string

// Set of constants representing the allowable values for UserCommonEnum
const (
	UserCommonYes UserCommonEnum = "YES"
	UserCommonNo  UserCommonEnum = "NO"
)

var mappingUserCommon = map[string]UserCommonEnum{
	"YES": UserCommonYes,
	"NO":  UserCommonNo,
}

// GetUserCommonEnumValues Enumerates the set of values for UserCommonEnum
func GetUserCommonEnumValues() []UserCommonEnum {
	values := make([]UserCommonEnum, 0)
	for _, v := range mappingUserCommon {
		values = append(values, v)
	}
	return values
}

// UserOracleMaintainedEnum Enum with underlying type: string
type UserOracleMaintainedEnum string

// Set of constants representing the allowable values for UserOracleMaintainedEnum
const (
	UserOracleMaintainedYes UserOracleMaintainedEnum = "YES"
	UserOracleMaintainedNo  UserOracleMaintainedEnum = "NO"
)

var mappingUserOracleMaintained = map[string]UserOracleMaintainedEnum{
	"YES": UserOracleMaintainedYes,
	"NO":  UserOracleMaintainedNo,
}

// GetUserOracleMaintainedEnumValues Enumerates the set of values for UserOracleMaintainedEnum
func GetUserOracleMaintainedEnumValues() []UserOracleMaintainedEnum {
	values := make([]UserOracleMaintainedEnum, 0)
	for _, v := range mappingUserOracleMaintained {
		values = append(values, v)
	}
	return values
}

// UserInheritedEnum Enum with underlying type: string
type UserInheritedEnum string

// Set of constants representing the allowable values for UserInheritedEnum
const (
	UserInheritedYes UserInheritedEnum = "YES"
	UserInheritedNo  UserInheritedEnum = "NO"
)

var mappingUserInherited = map[string]UserInheritedEnum{
	"YES": UserInheritedYes,
	"NO":  UserInheritedNo,
}

// GetUserInheritedEnumValues Enumerates the set of values for UserInheritedEnum
func GetUserInheritedEnumValues() []UserInheritedEnum {
	values := make([]UserInheritedEnum, 0)
	for _, v := range mappingUserInherited {
		values = append(values, v)
	}
	return values
}

// UserImplicitEnum Enum with underlying type: string
type UserImplicitEnum string

// Set of constants representing the allowable values for UserImplicitEnum
const (
	UserImplicitYes UserImplicitEnum = "YES"
	UserImplicitNo  UserImplicitEnum = "NO"
)

var mappingUserImplicit = map[string]UserImplicitEnum{
	"YES": UserImplicitYes,
	"NO":  UserImplicitNo,
}

// GetUserImplicitEnumValues Enumerates the set of values for UserImplicitEnum
func GetUserImplicitEnumValues() []UserImplicitEnum {
	values := make([]UserImplicitEnum, 0)
	for _, v := range mappingUserImplicit {
		values = append(values, v)
	}
	return values
}

// UserAllSharedEnum Enum with underlying type: string
type UserAllSharedEnum string

// Set of constants representing the allowable values for UserAllSharedEnum
const (
	UserAllSharedYes UserAllSharedEnum = "YES"
	UserAllSharedNo  UserAllSharedEnum = "NO"
)

var mappingUserAllShared = map[string]UserAllSharedEnum{
	"YES": UserAllSharedYes,
	"NO":  UserAllSharedNo,
}

// GetUserAllSharedEnumValues Enumerates the set of values for UserAllSharedEnum
func GetUserAllSharedEnumValues() []UserAllSharedEnum {
	values := make([]UserAllSharedEnum, 0)
	for _, v := range mappingUserAllShared {
		values = append(values, v)
	}
	return values
}

// UserExternalSharedEnum Enum with underlying type: string
type UserExternalSharedEnum string

// Set of constants representing the allowable values for UserExternalSharedEnum
const (
	UserExternalSharedYes UserExternalSharedEnum = "YES"
	UserExternalSharedNo  UserExternalSharedEnum = "NO"
)

var mappingUserExternalShared = map[string]UserExternalSharedEnum{
	"YES": UserExternalSharedYes,
	"NO":  UserExternalSharedNo,
}

// GetUserExternalSharedEnumValues Enumerates the set of values for UserExternalSharedEnum
func GetUserExternalSharedEnumValues() []UserExternalSharedEnum {
	values := make([]UserExternalSharedEnum, 0)
	for _, v := range mappingUserExternalShared {
		values = append(values, v)
	}
	return values
}
