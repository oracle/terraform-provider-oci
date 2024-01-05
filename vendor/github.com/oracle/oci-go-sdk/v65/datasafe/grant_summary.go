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

// GrantSummary The summary of user grants.
type GrantSummary struct {

	// The unique key of a user grant.
	Key *string `mandatory:"true" json:"key"`

	// The name of a user grant.
	GrantName *string `mandatory:"false" json:"grantName"`

	// The type of a user grant.
	PrivilegeType GrantSummaryPrivilegeTypeEnum `mandatory:"false" json:"privilegeType,omitempty"`

	// The privilege category.
	PrivilegeCategory GrantSummaryPrivilegeCategoryEnum `mandatory:"false" json:"privilegeCategory,omitempty"`

	// The grant depth level of the indirect grant.
	// An indirectly granted role/privilege is granted to the user through another role.
	// The depth level indicates how deep a privilege is within the grant hierarchy.
	DepthLevel *int `mandatory:"false" json:"depthLevel"`
}

func (m GrantSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GrantSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingGrantSummaryPrivilegeTypeEnum(string(m.PrivilegeType)); !ok && m.PrivilegeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrivilegeType: %s. Supported values are: %s.", m.PrivilegeType, strings.Join(GetGrantSummaryPrivilegeTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGrantSummaryPrivilegeCategoryEnum(string(m.PrivilegeCategory)); !ok && m.PrivilegeCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrivilegeCategory: %s. Supported values are: %s.", m.PrivilegeCategory, strings.Join(GetGrantSummaryPrivilegeCategoryEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GrantSummaryPrivilegeTypeEnum Enum with underlying type: string
type GrantSummaryPrivilegeTypeEnum string

// Set of constants representing the allowable values for GrantSummaryPrivilegeTypeEnum
const (
	GrantSummaryPrivilegeTypeSystemPrivilege GrantSummaryPrivilegeTypeEnum = "SYSTEM_PRIVILEGE"
	GrantSummaryPrivilegeTypeObjectPrivilege GrantSummaryPrivilegeTypeEnum = "OBJECT_PRIVILEGE"
	GrantSummaryPrivilegeTypeAdminPrivilege  GrantSummaryPrivilegeTypeEnum = "ADMIN_PRIVILEGE"
	GrantSummaryPrivilegeTypeRole            GrantSummaryPrivilegeTypeEnum = "ROLE"
)

var mappingGrantSummaryPrivilegeTypeEnum = map[string]GrantSummaryPrivilegeTypeEnum{
	"SYSTEM_PRIVILEGE": GrantSummaryPrivilegeTypeSystemPrivilege,
	"OBJECT_PRIVILEGE": GrantSummaryPrivilegeTypeObjectPrivilege,
	"ADMIN_PRIVILEGE":  GrantSummaryPrivilegeTypeAdminPrivilege,
	"ROLE":             GrantSummaryPrivilegeTypeRole,
}

var mappingGrantSummaryPrivilegeTypeEnumLowerCase = map[string]GrantSummaryPrivilegeTypeEnum{
	"system_privilege": GrantSummaryPrivilegeTypeSystemPrivilege,
	"object_privilege": GrantSummaryPrivilegeTypeObjectPrivilege,
	"admin_privilege":  GrantSummaryPrivilegeTypeAdminPrivilege,
	"role":             GrantSummaryPrivilegeTypeRole,
}

// GetGrantSummaryPrivilegeTypeEnumValues Enumerates the set of values for GrantSummaryPrivilegeTypeEnum
func GetGrantSummaryPrivilegeTypeEnumValues() []GrantSummaryPrivilegeTypeEnum {
	values := make([]GrantSummaryPrivilegeTypeEnum, 0)
	for _, v := range mappingGrantSummaryPrivilegeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGrantSummaryPrivilegeTypeEnumStringValues Enumerates the set of values in String for GrantSummaryPrivilegeTypeEnum
func GetGrantSummaryPrivilegeTypeEnumStringValues() []string {
	return []string{
		"SYSTEM_PRIVILEGE",
		"OBJECT_PRIVILEGE",
		"ADMIN_PRIVILEGE",
		"ROLE",
	}
}

// GetMappingGrantSummaryPrivilegeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGrantSummaryPrivilegeTypeEnum(val string) (GrantSummaryPrivilegeTypeEnum, bool) {
	enum, ok := mappingGrantSummaryPrivilegeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GrantSummaryPrivilegeCategoryEnum Enum with underlying type: string
type GrantSummaryPrivilegeCategoryEnum string

// Set of constants representing the allowable values for GrantSummaryPrivilegeCategoryEnum
const (
	GrantSummaryPrivilegeCategoryCritical GrantSummaryPrivilegeCategoryEnum = "CRITICAL"
	GrantSummaryPrivilegeCategoryHigh     GrantSummaryPrivilegeCategoryEnum = "HIGH"
	GrantSummaryPrivilegeCategoryMedium   GrantSummaryPrivilegeCategoryEnum = "MEDIUM"
	GrantSummaryPrivilegeCategoryLow      GrantSummaryPrivilegeCategoryEnum = "LOW"
)

var mappingGrantSummaryPrivilegeCategoryEnum = map[string]GrantSummaryPrivilegeCategoryEnum{
	"CRITICAL": GrantSummaryPrivilegeCategoryCritical,
	"HIGH":     GrantSummaryPrivilegeCategoryHigh,
	"MEDIUM":   GrantSummaryPrivilegeCategoryMedium,
	"LOW":      GrantSummaryPrivilegeCategoryLow,
}

var mappingGrantSummaryPrivilegeCategoryEnumLowerCase = map[string]GrantSummaryPrivilegeCategoryEnum{
	"critical": GrantSummaryPrivilegeCategoryCritical,
	"high":     GrantSummaryPrivilegeCategoryHigh,
	"medium":   GrantSummaryPrivilegeCategoryMedium,
	"low":      GrantSummaryPrivilegeCategoryLow,
}

// GetGrantSummaryPrivilegeCategoryEnumValues Enumerates the set of values for GrantSummaryPrivilegeCategoryEnum
func GetGrantSummaryPrivilegeCategoryEnumValues() []GrantSummaryPrivilegeCategoryEnum {
	values := make([]GrantSummaryPrivilegeCategoryEnum, 0)
	for _, v := range mappingGrantSummaryPrivilegeCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetGrantSummaryPrivilegeCategoryEnumStringValues Enumerates the set of values in String for GrantSummaryPrivilegeCategoryEnum
func GetGrantSummaryPrivilegeCategoryEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
	}
}

// GetMappingGrantSummaryPrivilegeCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGrantSummaryPrivilegeCategoryEnum(val string) (GrantSummaryPrivilegeCategoryEnum, bool) {
	enum, ok := mappingGrantSummaryPrivilegeCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
