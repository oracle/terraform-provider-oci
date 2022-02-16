// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// UserSummary The summary of information about the database user. It includes details such as user type, account status,
// last login time, user creation time, authentication type, user profile, and time and date of the last password change.
// It also contains the user category derived from these user details, as well as granted privileges.
type UserSummary struct {

	// The unique user key. This is a system-generated identifier. Use ListUsers to get the user key for a user.
	Key *string `mandatory:"true" json:"key"`

	// The database user name.
	UserName *string `mandatory:"true" json:"userName"`

	// The OCID of the target database.
	TargetId *string `mandatory:"true" json:"targetId"`

	// The user category based on the privileges and other details of the user.
	UserCategory UserSummaryUserCategoryEnum `mandatory:"false" json:"userCategory,omitempty"`

	// The user account status.
	AccountStatus UserSummaryAccountStatusEnum `mandatory:"false" json:"accountStatus,omitempty"`

	// The date and time when the user last logged in, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeLastLogin *common.SDKTime `mandatory:"false" json:"timeLastLogin"`

	// The date and time when the user was created in the database, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUserCreated *common.SDKTime `mandatory:"false" json:"timeUserCreated"`

	// The user authentication method.
	AuthenticationType UserSummaryAuthenticationTypeEnum `mandatory:"false" json:"authenticationType,omitempty"`

	// The user profile name.
	UserProfile *string `mandatory:"false" json:"userProfile"`

	// The date and time when the user password was last changed, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimePasswordChanged *common.SDKTime `mandatory:"false" json:"timePasswordChanged"`

	// The user type, which can be a combination of the following:
	// 'Admin Privileged': The user has administrative privileges.
	// 'Application': The user is an Oracle E-Business Suite Applications (EBS) or Fusion Applications (FA) user.
	// 'Privileged': The user is a privileged user.
	// 'Schema': The user is EXPIRED & LOCKED / EXPIRED / LOCKED, or a schema-only account (authentication type is NONE).
	// 'Non-privileged': The user is a non-privileged user.
	UserTypes []UserSummaryUserTypesEnum `mandatory:"false" json:"userTypes,omitempty"`

	// The admin roles granted to the user.
	AdminRoles []UserSummaryAdminRolesEnum `mandatory:"false" json:"adminRoles,omitempty"`
}

func (m UserSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUserSummaryUserCategoryEnum(string(m.UserCategory)); !ok && m.UserCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UserCategory: %s. Supported values are: %s.", m.UserCategory, strings.Join(GetUserSummaryUserCategoryEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUserSummaryAccountStatusEnum(string(m.AccountStatus)); !ok && m.AccountStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccountStatus: %s. Supported values are: %s.", m.AccountStatus, strings.Join(GetUserSummaryAccountStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUserSummaryAuthenticationTypeEnum(string(m.AuthenticationType)); !ok && m.AuthenticationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticationType: %s. Supported values are: %s.", m.AuthenticationType, strings.Join(GetUserSummaryAuthenticationTypeEnumStringValues(), ",")))
	}
	for _, val := range m.UserTypes {
		if _, ok := GetMappingUserSummaryUserTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UserTypes: %s. Supported values are: %s.", val, strings.Join(GetUserSummaryUserTypesEnumStringValues(), ",")))
		}
	}

	for _, val := range m.AdminRoles {
		if _, ok := GetMappingUserSummaryAdminRolesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdminRoles: %s. Supported values are: %s.", val, strings.Join(GetUserSummaryAdminRolesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UserSummaryUserCategoryEnum Enum with underlying type: string
type UserSummaryUserCategoryEnum string

// Set of constants representing the allowable values for UserSummaryUserCategoryEnum
const (
	UserSummaryUserCategoryCritical UserSummaryUserCategoryEnum = "CRITICAL"
	UserSummaryUserCategoryHigh     UserSummaryUserCategoryEnum = "HIGH"
	UserSummaryUserCategoryMedium   UserSummaryUserCategoryEnum = "MEDIUM"
	UserSummaryUserCategoryLow      UserSummaryUserCategoryEnum = "LOW"
)

var mappingUserSummaryUserCategoryEnum = map[string]UserSummaryUserCategoryEnum{
	"CRITICAL": UserSummaryUserCategoryCritical,
	"HIGH":     UserSummaryUserCategoryHigh,
	"MEDIUM":   UserSummaryUserCategoryMedium,
	"LOW":      UserSummaryUserCategoryLow,
}

// GetUserSummaryUserCategoryEnumValues Enumerates the set of values for UserSummaryUserCategoryEnum
func GetUserSummaryUserCategoryEnumValues() []UserSummaryUserCategoryEnum {
	values := make([]UserSummaryUserCategoryEnum, 0)
	for _, v := range mappingUserSummaryUserCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetUserSummaryUserCategoryEnumStringValues Enumerates the set of values in String for UserSummaryUserCategoryEnum
func GetUserSummaryUserCategoryEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
	}
}

// GetMappingUserSummaryUserCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserSummaryUserCategoryEnum(val string) (UserSummaryUserCategoryEnum, bool) {
	mappingUserSummaryUserCategoryEnumIgnoreCase := make(map[string]UserSummaryUserCategoryEnum)
	for k, v := range mappingUserSummaryUserCategoryEnum {
		mappingUserSummaryUserCategoryEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUserSummaryUserCategoryEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// UserSummaryAccountStatusEnum Enum with underlying type: string
type UserSummaryAccountStatusEnum string

// Set of constants representing the allowable values for UserSummaryAccountStatusEnum
const (
	UserSummaryAccountStatusOpen             UserSummaryAccountStatusEnum = "OPEN"
	UserSummaryAccountStatusLocked           UserSummaryAccountStatusEnum = "LOCKED"
	UserSummaryAccountStatusExpired          UserSummaryAccountStatusEnum = "EXPIRED"
	UserSummaryAccountStatusExpiredAndLocked UserSummaryAccountStatusEnum = "EXPIRED_AND_LOCKED"
	UserSummaryAccountStatusNone             UserSummaryAccountStatusEnum = "NONE"
)

var mappingUserSummaryAccountStatusEnum = map[string]UserSummaryAccountStatusEnum{
	"OPEN":               UserSummaryAccountStatusOpen,
	"LOCKED":             UserSummaryAccountStatusLocked,
	"EXPIRED":            UserSummaryAccountStatusExpired,
	"EXPIRED_AND_LOCKED": UserSummaryAccountStatusExpiredAndLocked,
	"NONE":               UserSummaryAccountStatusNone,
}

// GetUserSummaryAccountStatusEnumValues Enumerates the set of values for UserSummaryAccountStatusEnum
func GetUserSummaryAccountStatusEnumValues() []UserSummaryAccountStatusEnum {
	values := make([]UserSummaryAccountStatusEnum, 0)
	for _, v := range mappingUserSummaryAccountStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetUserSummaryAccountStatusEnumStringValues Enumerates the set of values in String for UserSummaryAccountStatusEnum
func GetUserSummaryAccountStatusEnumStringValues() []string {
	return []string{
		"OPEN",
		"LOCKED",
		"EXPIRED",
		"EXPIRED_AND_LOCKED",
		"NONE",
	}
}

// GetMappingUserSummaryAccountStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserSummaryAccountStatusEnum(val string) (UserSummaryAccountStatusEnum, bool) {
	mappingUserSummaryAccountStatusEnumIgnoreCase := make(map[string]UserSummaryAccountStatusEnum)
	for k, v := range mappingUserSummaryAccountStatusEnum {
		mappingUserSummaryAccountStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUserSummaryAccountStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// UserSummaryAuthenticationTypeEnum Enum with underlying type: string
type UserSummaryAuthenticationTypeEnum string

// Set of constants representing the allowable values for UserSummaryAuthenticationTypeEnum
const (
	UserSummaryAuthenticationTypePassword UserSummaryAuthenticationTypeEnum = "PASSWORD"
	UserSummaryAuthenticationTypeNone     UserSummaryAuthenticationTypeEnum = "NONE"
)

var mappingUserSummaryAuthenticationTypeEnum = map[string]UserSummaryAuthenticationTypeEnum{
	"PASSWORD": UserSummaryAuthenticationTypePassword,
	"NONE":     UserSummaryAuthenticationTypeNone,
}

// GetUserSummaryAuthenticationTypeEnumValues Enumerates the set of values for UserSummaryAuthenticationTypeEnum
func GetUserSummaryAuthenticationTypeEnumValues() []UserSummaryAuthenticationTypeEnum {
	values := make([]UserSummaryAuthenticationTypeEnum, 0)
	for _, v := range mappingUserSummaryAuthenticationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUserSummaryAuthenticationTypeEnumStringValues Enumerates the set of values in String for UserSummaryAuthenticationTypeEnum
func GetUserSummaryAuthenticationTypeEnumStringValues() []string {
	return []string{
		"PASSWORD",
		"NONE",
	}
}

// GetMappingUserSummaryAuthenticationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserSummaryAuthenticationTypeEnum(val string) (UserSummaryAuthenticationTypeEnum, bool) {
	mappingUserSummaryAuthenticationTypeEnumIgnoreCase := make(map[string]UserSummaryAuthenticationTypeEnum)
	for k, v := range mappingUserSummaryAuthenticationTypeEnum {
		mappingUserSummaryAuthenticationTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUserSummaryAuthenticationTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// UserSummaryUserTypesEnum Enum with underlying type: string
type UserSummaryUserTypesEnum string

// Set of constants representing the allowable values for UserSummaryUserTypesEnum
const (
	UserSummaryUserTypesAdminPrivileged UserSummaryUserTypesEnum = "ADMIN_PRIVILEGED"
	UserSummaryUserTypesApplication     UserSummaryUserTypesEnum = "APPLICATION"
	UserSummaryUserTypesPrivileged      UserSummaryUserTypesEnum = "PRIVILEGED"
	UserSummaryUserTypesSchema          UserSummaryUserTypesEnum = "SCHEMA"
	UserSummaryUserTypesNonPrivileged   UserSummaryUserTypesEnum = "NON_PRIVILEGED"
)

var mappingUserSummaryUserTypesEnum = map[string]UserSummaryUserTypesEnum{
	"ADMIN_PRIVILEGED": UserSummaryUserTypesAdminPrivileged,
	"APPLICATION":      UserSummaryUserTypesApplication,
	"PRIVILEGED":       UserSummaryUserTypesPrivileged,
	"SCHEMA":           UserSummaryUserTypesSchema,
	"NON_PRIVILEGED":   UserSummaryUserTypesNonPrivileged,
}

// GetUserSummaryUserTypesEnumValues Enumerates the set of values for UserSummaryUserTypesEnum
func GetUserSummaryUserTypesEnumValues() []UserSummaryUserTypesEnum {
	values := make([]UserSummaryUserTypesEnum, 0)
	for _, v := range mappingUserSummaryUserTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetUserSummaryUserTypesEnumStringValues Enumerates the set of values in String for UserSummaryUserTypesEnum
func GetUserSummaryUserTypesEnumStringValues() []string {
	return []string{
		"ADMIN_PRIVILEGED",
		"APPLICATION",
		"PRIVILEGED",
		"SCHEMA",
		"NON_PRIVILEGED",
	}
}

// GetMappingUserSummaryUserTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserSummaryUserTypesEnum(val string) (UserSummaryUserTypesEnum, bool) {
	mappingUserSummaryUserTypesEnumIgnoreCase := make(map[string]UserSummaryUserTypesEnum)
	for k, v := range mappingUserSummaryUserTypesEnum {
		mappingUserSummaryUserTypesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUserSummaryUserTypesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// UserSummaryAdminRolesEnum Enum with underlying type: string
type UserSummaryAdminRolesEnum string

// Set of constants representing the allowable values for UserSummaryAdminRolesEnum
const (
	UserSummaryAdminRolesPdbDba     UserSummaryAdminRolesEnum = "PDB_DBA"
	UserSummaryAdminRolesDba        UserSummaryAdminRolesEnum = "DBA"
	UserSummaryAdminRolesDvAdmin    UserSummaryAdminRolesEnum = "DV_ADMIN"
	UserSummaryAdminRolesAuditAdmin UserSummaryAdminRolesEnum = "AUDIT_ADMIN"
)

var mappingUserSummaryAdminRolesEnum = map[string]UserSummaryAdminRolesEnum{
	"PDB_DBA":     UserSummaryAdminRolesPdbDba,
	"DBA":         UserSummaryAdminRolesDba,
	"DV_ADMIN":    UserSummaryAdminRolesDvAdmin,
	"AUDIT_ADMIN": UserSummaryAdminRolesAuditAdmin,
}

// GetUserSummaryAdminRolesEnumValues Enumerates the set of values for UserSummaryAdminRolesEnum
func GetUserSummaryAdminRolesEnumValues() []UserSummaryAdminRolesEnum {
	values := make([]UserSummaryAdminRolesEnum, 0)
	for _, v := range mappingUserSummaryAdminRolesEnum {
		values = append(values, v)
	}
	return values
}

// GetUserSummaryAdminRolesEnumStringValues Enumerates the set of values in String for UserSummaryAdminRolesEnum
func GetUserSummaryAdminRolesEnumStringValues() []string {
	return []string{
		"PDB_DBA",
		"DBA",
		"DV_ADMIN",
		"AUDIT_ADMIN",
	}
}

// GetMappingUserSummaryAdminRolesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserSummaryAdminRolesEnum(val string) (UserSummaryAdminRolesEnum, bool) {
	mappingUserSummaryAdminRolesEnumIgnoreCase := make(map[string]UserSummaryAdminRolesEnum)
	for k, v := range mappingUserSummaryAdminRolesEnum {
		mappingUserSummaryAdminRolesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUserSummaryAdminRolesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
