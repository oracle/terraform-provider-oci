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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// User The summary of a specific User.
type User struct {

	// The name of the User.
	Name *string `mandatory:"true" json:"name"`

	// The status of the user account.
	Status UserStatusEnum `mandatory:"true" json:"status"`

	// The default tablespace for data.
	DefaultTablespace *string `mandatory:"true" json:"defaultTablespace"`

	// The name of the default tablespace for temporary tables or the name of a tablespace group.
	TempTablespace *string `mandatory:"true" json:"tempTablespace"`

	// The date and time the user was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The User resource profile name.
	Profile *string `mandatory:"true" json:"profile"`

	// The date the account was locked if account status was LOCKED.
	TimeLocked *common.SDKTime `mandatory:"false" json:"timeLocked"`

	// The date and time of the expiration of the user account.
	TimeExpiring *common.SDKTime `mandatory:"false" json:"timeExpiring"`

	// The default local temporary tablespace for the user.
	LocalTempTablespace *string `mandatory:"false" json:"localTempTablespace"`

	// The initial resource consumer group for the User.
	ConsumerGroup *string `mandatory:"false" json:"consumerGroup"`

	// The external name of the user.
	ExternalName *string `mandatory:"false" json:"externalName"`

	// The list of existing versions of the password hashes (also known as "verifiers") for the account.
	PasswordVersions *string `mandatory:"false" json:"passwordVersions"`

	// Indicates whether editions have been enabled for the corresponding user (Y) or not (N).
	EditionsEnabled UserEditionsEnabledEnum `mandatory:"false" json:"editionsEnabled,omitempty"`

	// The authentication mechanism for the user.
	Authentication UserAuthenticationEnum `mandatory:"false" json:"authentication,omitempty"`

	// Indicates whether a user can connect directly (N) or whether the account can only be proxied (Y) by users who have proxy privileges
	// for this account (that is, by users who have been granted the "connect through" privilege for this account).
	ProxyConnect UserProxyConnectEnum `mandatory:"false" json:"proxyConnect,omitempty"`

	// Indicates whether a given user is common(Y) or local(N).
	Common UserCommonEnum `mandatory:"false" json:"common,omitempty"`

	// The date and time of the last user login.
	// This column is not populated when a user connects to the database with administrative privileges, that is, AS { SYSASM | SYSBACKUP | SYSDBA | SYSDG | SYSOPER | SYSRAC | SYSKM }.
	TimeLastLogin *common.SDKTime `mandatory:"false" json:"timeLastLogin"`

	// Indicates whether the user was created and is maintained by Oracle-supplied scripts (such as catalog.sql or catproc.sql).
	OracleMaintained UserOracleMaintainedEnum `mandatory:"false" json:"oracleMaintained,omitempty"`

	// Indicates whether the user definition is inherited from another container (YES) or not (NO).
	Inherited UserInheritedEnum `mandatory:"false" json:"inherited,omitempty"`

	// The default collation for the user schema.
	DefaultCollation *string `mandatory:"false" json:"defaultCollation"`

	// Indicates whether the user is a common user created by an implicit application (YES) or not (NO).
	Implicit UserImplicitEnum `mandatory:"false" json:"implicit,omitempty"`

	// In a sharded database, indicates whether the user is created with shard DDL enabled (YES) or not (NO).
	AllShared UserAllSharedEnum `mandatory:"false" json:"allShared,omitempty"`

	// In a federated sharded database, indicates whether the user is an external shard user (YES) or not (NO).
	ExternalShared UserExternalSharedEnum `mandatory:"false" json:"externalShared,omitempty"`

	// The date and time when the user password was last set.
	// This column is populated only when the value of the AUTHENTICATION_TYPE column is PASSWORD. Otherwise, this column is null.
	TimePasswordChanged *common.SDKTime `mandatory:"false" json:"timePasswordChanged"`
}

func (m User) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m User) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUserStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetUserStatusEnumStringValues(), ",")))
	}

	if _, ok := GetMappingUserEditionsEnabledEnum(string(m.EditionsEnabled)); !ok && m.EditionsEnabled != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EditionsEnabled: %s. Supported values are: %s.", m.EditionsEnabled, strings.Join(GetUserEditionsEnabledEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUserAuthenticationEnum(string(m.Authentication)); !ok && m.Authentication != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Authentication: %s. Supported values are: %s.", m.Authentication, strings.Join(GetUserAuthenticationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUserProxyConnectEnum(string(m.ProxyConnect)); !ok && m.ProxyConnect != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProxyConnect: %s. Supported values are: %s.", m.ProxyConnect, strings.Join(GetUserProxyConnectEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUserCommonEnum(string(m.Common)); !ok && m.Common != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Common: %s. Supported values are: %s.", m.Common, strings.Join(GetUserCommonEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUserOracleMaintainedEnum(string(m.OracleMaintained)); !ok && m.OracleMaintained != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OracleMaintained: %s. Supported values are: %s.", m.OracleMaintained, strings.Join(GetUserOracleMaintainedEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUserInheritedEnum(string(m.Inherited)); !ok && m.Inherited != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Inherited: %s. Supported values are: %s.", m.Inherited, strings.Join(GetUserInheritedEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUserImplicitEnum(string(m.Implicit)); !ok && m.Implicit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Implicit: %s. Supported values are: %s.", m.Implicit, strings.Join(GetUserImplicitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUserAllSharedEnum(string(m.AllShared)); !ok && m.AllShared != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AllShared: %s. Supported values are: %s.", m.AllShared, strings.Join(GetUserAllSharedEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUserExternalSharedEnum(string(m.ExternalShared)); !ok && m.ExternalShared != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExternalShared: %s. Supported values are: %s.", m.ExternalShared, strings.Join(GetUserExternalSharedEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingUserStatusEnum = map[string]UserStatusEnum{
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
	for _, v := range mappingUserStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetUserStatusEnumStringValues Enumerates the set of values in String for UserStatusEnum
func GetUserStatusEnumStringValues() []string {
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

// GetMappingUserStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserStatusEnum(val string) (UserStatusEnum, bool) {
	mappingUserStatusEnumIgnoreCase := make(map[string]UserStatusEnum)
	for k, v := range mappingUserStatusEnum {
		mappingUserStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUserStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// UserEditionsEnabledEnum Enum with underlying type: string
type UserEditionsEnabledEnum string

// Set of constants representing the allowable values for UserEditionsEnabledEnum
const (
	UserEditionsEnabledYes UserEditionsEnabledEnum = "YES"
	UserEditionsEnabledNo  UserEditionsEnabledEnum = "NO"
)

var mappingUserEditionsEnabledEnum = map[string]UserEditionsEnabledEnum{
	"YES": UserEditionsEnabledYes,
	"NO":  UserEditionsEnabledNo,
}

// GetUserEditionsEnabledEnumValues Enumerates the set of values for UserEditionsEnabledEnum
func GetUserEditionsEnabledEnumValues() []UserEditionsEnabledEnum {
	values := make([]UserEditionsEnabledEnum, 0)
	for _, v := range mappingUserEditionsEnabledEnum {
		values = append(values, v)
	}
	return values
}

// GetUserEditionsEnabledEnumStringValues Enumerates the set of values in String for UserEditionsEnabledEnum
func GetUserEditionsEnabledEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingUserEditionsEnabledEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserEditionsEnabledEnum(val string) (UserEditionsEnabledEnum, bool) {
	mappingUserEditionsEnabledEnumIgnoreCase := make(map[string]UserEditionsEnabledEnum)
	for k, v := range mappingUserEditionsEnabledEnum {
		mappingUserEditionsEnabledEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUserEditionsEnabledEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
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

var mappingUserAuthenticationEnum = map[string]UserAuthenticationEnum{
	"NONE":     UserAuthenticationNone,
	"EXTERNAL": UserAuthenticationExternal,
	"GLOBAL":   UserAuthenticationGlobal,
	"PASSWORD": UserAuthenticationPassword,
}

// GetUserAuthenticationEnumValues Enumerates the set of values for UserAuthenticationEnum
func GetUserAuthenticationEnumValues() []UserAuthenticationEnum {
	values := make([]UserAuthenticationEnum, 0)
	for _, v := range mappingUserAuthenticationEnum {
		values = append(values, v)
	}
	return values
}

// GetUserAuthenticationEnumStringValues Enumerates the set of values in String for UserAuthenticationEnum
func GetUserAuthenticationEnumStringValues() []string {
	return []string{
		"NONE",
		"EXTERNAL",
		"GLOBAL",
		"PASSWORD",
	}
}

// GetMappingUserAuthenticationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserAuthenticationEnum(val string) (UserAuthenticationEnum, bool) {
	mappingUserAuthenticationEnumIgnoreCase := make(map[string]UserAuthenticationEnum)
	for k, v := range mappingUserAuthenticationEnum {
		mappingUserAuthenticationEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUserAuthenticationEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// UserProxyConnectEnum Enum with underlying type: string
type UserProxyConnectEnum string

// Set of constants representing the allowable values for UserProxyConnectEnum
const (
	UserProxyConnectYes UserProxyConnectEnum = "YES"
	UserProxyConnectNo  UserProxyConnectEnum = "NO"
)

var mappingUserProxyConnectEnum = map[string]UserProxyConnectEnum{
	"YES": UserProxyConnectYes,
	"NO":  UserProxyConnectNo,
}

// GetUserProxyConnectEnumValues Enumerates the set of values for UserProxyConnectEnum
func GetUserProxyConnectEnumValues() []UserProxyConnectEnum {
	values := make([]UserProxyConnectEnum, 0)
	for _, v := range mappingUserProxyConnectEnum {
		values = append(values, v)
	}
	return values
}

// GetUserProxyConnectEnumStringValues Enumerates the set of values in String for UserProxyConnectEnum
func GetUserProxyConnectEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingUserProxyConnectEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserProxyConnectEnum(val string) (UserProxyConnectEnum, bool) {
	mappingUserProxyConnectEnumIgnoreCase := make(map[string]UserProxyConnectEnum)
	for k, v := range mappingUserProxyConnectEnum {
		mappingUserProxyConnectEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUserProxyConnectEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// UserCommonEnum Enum with underlying type: string
type UserCommonEnum string

// Set of constants representing the allowable values for UserCommonEnum
const (
	UserCommonYes UserCommonEnum = "YES"
	UserCommonNo  UserCommonEnum = "NO"
)

var mappingUserCommonEnum = map[string]UserCommonEnum{
	"YES": UserCommonYes,
	"NO":  UserCommonNo,
}

// GetUserCommonEnumValues Enumerates the set of values for UserCommonEnum
func GetUserCommonEnumValues() []UserCommonEnum {
	values := make([]UserCommonEnum, 0)
	for _, v := range mappingUserCommonEnum {
		values = append(values, v)
	}
	return values
}

// GetUserCommonEnumStringValues Enumerates the set of values in String for UserCommonEnum
func GetUserCommonEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingUserCommonEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserCommonEnum(val string) (UserCommonEnum, bool) {
	mappingUserCommonEnumIgnoreCase := make(map[string]UserCommonEnum)
	for k, v := range mappingUserCommonEnum {
		mappingUserCommonEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUserCommonEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// UserOracleMaintainedEnum Enum with underlying type: string
type UserOracleMaintainedEnum string

// Set of constants representing the allowable values for UserOracleMaintainedEnum
const (
	UserOracleMaintainedYes UserOracleMaintainedEnum = "YES"
	UserOracleMaintainedNo  UserOracleMaintainedEnum = "NO"
)

var mappingUserOracleMaintainedEnum = map[string]UserOracleMaintainedEnum{
	"YES": UserOracleMaintainedYes,
	"NO":  UserOracleMaintainedNo,
}

// GetUserOracleMaintainedEnumValues Enumerates the set of values for UserOracleMaintainedEnum
func GetUserOracleMaintainedEnumValues() []UserOracleMaintainedEnum {
	values := make([]UserOracleMaintainedEnum, 0)
	for _, v := range mappingUserOracleMaintainedEnum {
		values = append(values, v)
	}
	return values
}

// GetUserOracleMaintainedEnumStringValues Enumerates the set of values in String for UserOracleMaintainedEnum
func GetUserOracleMaintainedEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingUserOracleMaintainedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserOracleMaintainedEnum(val string) (UserOracleMaintainedEnum, bool) {
	mappingUserOracleMaintainedEnumIgnoreCase := make(map[string]UserOracleMaintainedEnum)
	for k, v := range mappingUserOracleMaintainedEnum {
		mappingUserOracleMaintainedEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUserOracleMaintainedEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// UserInheritedEnum Enum with underlying type: string
type UserInheritedEnum string

// Set of constants representing the allowable values for UserInheritedEnum
const (
	UserInheritedYes UserInheritedEnum = "YES"
	UserInheritedNo  UserInheritedEnum = "NO"
)

var mappingUserInheritedEnum = map[string]UserInheritedEnum{
	"YES": UserInheritedYes,
	"NO":  UserInheritedNo,
}

// GetUserInheritedEnumValues Enumerates the set of values for UserInheritedEnum
func GetUserInheritedEnumValues() []UserInheritedEnum {
	values := make([]UserInheritedEnum, 0)
	for _, v := range mappingUserInheritedEnum {
		values = append(values, v)
	}
	return values
}

// GetUserInheritedEnumStringValues Enumerates the set of values in String for UserInheritedEnum
func GetUserInheritedEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingUserInheritedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserInheritedEnum(val string) (UserInheritedEnum, bool) {
	mappingUserInheritedEnumIgnoreCase := make(map[string]UserInheritedEnum)
	for k, v := range mappingUserInheritedEnum {
		mappingUserInheritedEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUserInheritedEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// UserImplicitEnum Enum with underlying type: string
type UserImplicitEnum string

// Set of constants representing the allowable values for UserImplicitEnum
const (
	UserImplicitYes UserImplicitEnum = "YES"
	UserImplicitNo  UserImplicitEnum = "NO"
)

var mappingUserImplicitEnum = map[string]UserImplicitEnum{
	"YES": UserImplicitYes,
	"NO":  UserImplicitNo,
}

// GetUserImplicitEnumValues Enumerates the set of values for UserImplicitEnum
func GetUserImplicitEnumValues() []UserImplicitEnum {
	values := make([]UserImplicitEnum, 0)
	for _, v := range mappingUserImplicitEnum {
		values = append(values, v)
	}
	return values
}

// GetUserImplicitEnumStringValues Enumerates the set of values in String for UserImplicitEnum
func GetUserImplicitEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingUserImplicitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserImplicitEnum(val string) (UserImplicitEnum, bool) {
	mappingUserImplicitEnumIgnoreCase := make(map[string]UserImplicitEnum)
	for k, v := range mappingUserImplicitEnum {
		mappingUserImplicitEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUserImplicitEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// UserAllSharedEnum Enum with underlying type: string
type UserAllSharedEnum string

// Set of constants representing the allowable values for UserAllSharedEnum
const (
	UserAllSharedYes UserAllSharedEnum = "YES"
	UserAllSharedNo  UserAllSharedEnum = "NO"
)

var mappingUserAllSharedEnum = map[string]UserAllSharedEnum{
	"YES": UserAllSharedYes,
	"NO":  UserAllSharedNo,
}

// GetUserAllSharedEnumValues Enumerates the set of values for UserAllSharedEnum
func GetUserAllSharedEnumValues() []UserAllSharedEnum {
	values := make([]UserAllSharedEnum, 0)
	for _, v := range mappingUserAllSharedEnum {
		values = append(values, v)
	}
	return values
}

// GetUserAllSharedEnumStringValues Enumerates the set of values in String for UserAllSharedEnum
func GetUserAllSharedEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingUserAllSharedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserAllSharedEnum(val string) (UserAllSharedEnum, bool) {
	mappingUserAllSharedEnumIgnoreCase := make(map[string]UserAllSharedEnum)
	for k, v := range mappingUserAllSharedEnum {
		mappingUserAllSharedEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUserAllSharedEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// UserExternalSharedEnum Enum with underlying type: string
type UserExternalSharedEnum string

// Set of constants representing the allowable values for UserExternalSharedEnum
const (
	UserExternalSharedYes UserExternalSharedEnum = "YES"
	UserExternalSharedNo  UserExternalSharedEnum = "NO"
)

var mappingUserExternalSharedEnum = map[string]UserExternalSharedEnum{
	"YES": UserExternalSharedYes,
	"NO":  UserExternalSharedNo,
}

// GetUserExternalSharedEnumValues Enumerates the set of values for UserExternalSharedEnum
func GetUserExternalSharedEnumValues() []UserExternalSharedEnum {
	values := make([]UserExternalSharedEnum, 0)
	for _, v := range mappingUserExternalSharedEnum {
		values = append(values, v)
	}
	return values
}

// GetUserExternalSharedEnumStringValues Enumerates the set of values in String for UserExternalSharedEnum
func GetUserExternalSharedEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingUserExternalSharedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserExternalSharedEnum(val string) (UserExternalSharedEnum, bool) {
	mappingUserExternalSharedEnumIgnoreCase := make(map[string]UserExternalSharedEnum)
	for k, v := range mappingUserExternalSharedEnum {
		mappingUserExternalSharedEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUserExternalSharedEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
