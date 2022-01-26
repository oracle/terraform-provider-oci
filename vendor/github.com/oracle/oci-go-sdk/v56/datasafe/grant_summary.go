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

// GrantSummaryPrivilegeTypeEnum Enum with underlying type: string
type GrantSummaryPrivilegeTypeEnum string

// Set of constants representing the allowable values for GrantSummaryPrivilegeTypeEnum
const (
	GrantSummaryPrivilegeTypeSystemPrivilege GrantSummaryPrivilegeTypeEnum = "SYSTEM_PRIVILEGE"
	GrantSummaryPrivilegeTypeObjectPrivilege GrantSummaryPrivilegeTypeEnum = "OBJECT_PRIVILEGE"
	GrantSummaryPrivilegeTypeAdminPrivilege  GrantSummaryPrivilegeTypeEnum = "ADMIN_PRIVILEGE"
	GrantSummaryPrivilegeTypeRole            GrantSummaryPrivilegeTypeEnum = "ROLE"
)

var mappingGrantSummaryPrivilegeType = map[string]GrantSummaryPrivilegeTypeEnum{
	"SYSTEM_PRIVILEGE": GrantSummaryPrivilegeTypeSystemPrivilege,
	"OBJECT_PRIVILEGE": GrantSummaryPrivilegeTypeObjectPrivilege,
	"ADMIN_PRIVILEGE":  GrantSummaryPrivilegeTypeAdminPrivilege,
	"ROLE":             GrantSummaryPrivilegeTypeRole,
}

// GetGrantSummaryPrivilegeTypeEnumValues Enumerates the set of values for GrantSummaryPrivilegeTypeEnum
func GetGrantSummaryPrivilegeTypeEnumValues() []GrantSummaryPrivilegeTypeEnum {
	values := make([]GrantSummaryPrivilegeTypeEnum, 0)
	for _, v := range mappingGrantSummaryPrivilegeType {
		values = append(values, v)
	}
	return values
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

var mappingGrantSummaryPrivilegeCategory = map[string]GrantSummaryPrivilegeCategoryEnum{
	"CRITICAL": GrantSummaryPrivilegeCategoryCritical,
	"HIGH":     GrantSummaryPrivilegeCategoryHigh,
	"MEDIUM":   GrantSummaryPrivilegeCategoryMedium,
	"LOW":      GrantSummaryPrivilegeCategoryLow,
}

// GetGrantSummaryPrivilegeCategoryEnumValues Enumerates the set of values for GrantSummaryPrivilegeCategoryEnum
func GetGrantSummaryPrivilegeCategoryEnumValues() []GrantSummaryPrivilegeCategoryEnum {
	values := make([]GrantSummaryPrivilegeCategoryEnum, 0)
	for _, v := range mappingGrantSummaryPrivilegeCategory {
		values = append(values, v)
	}
	return values
}
