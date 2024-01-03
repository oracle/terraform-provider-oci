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

// AuditSpecification Represents an audit policy relevant for the target database.The audit policy could be in any one of the following 3 states in the target database
// 1) Created and enabled
// 2) Created but not enabled
// 3) Not created
// For more details on relevant audit policies for the target database, refer to documentation (https://docs.oracle.com/en/cloud/paas/data-safe/udscs/audit-policies.html#GUID-361A9A9A-7C21-4F5A-8945-9B3A0C472827).
type AuditSpecification struct {

	// Indicates the audit policy name. Refer to the documentation (https://docs.oracle.com/en/cloud/paas/data-safe/udscs/audit-policies.html#GUID-361A9A9A-7C21-4F5A-8945-9B3A0C472827) for seeded audit policy names. For custom policies, refer to the user-defined policy name created in the target database.
	AuditPolicyName *string `mandatory:"true" json:"auditPolicyName"`

	// Indicates the names of corresponding database policy ( or policies) in the target database.
	DatabasePolicyNames []string `mandatory:"true" json:"databasePolicyNames"`

	// The category to which the audit policy belongs.
	AuditPolicyCategory AuditPolicyCategoryEnum `mandatory:"true" json:"auditPolicyCategory"`

	// Indicates whether the policy has been enabled, disabled or partially enabled in the target database. The status is PARTIALLY_ENABLED if any of the constituent database audit policies is not enabled.
	EnableStatus AuditSpecificationEnableStatusEnum `mandatory:"true" json:"enableStatus"`

	// Indicates whether the policy by default is enabled for all users with no flexibility to alter the enablement conditions.
	IsEnabledForAllUsers *bool `mandatory:"true" json:"isEnabledForAllUsers"`

	// Indicates whether the audit policy is available for provisioning/ de-provisioning from Oracle Data Safe, or is only available for displaying the current provisioning status from the target.
	IsViewOnly *bool `mandatory:"true" json:"isViewOnly"`

	// Indicates whether the audit policy is one of the predefined policies provided by Oracle Database.
	IsSeededInTarget *bool `mandatory:"true" json:"isSeededInTarget"`

	// Indicates whether the audit policy is one of the seeded policies provided by Oracle Data Safe.
	IsSeededInDataSafe *bool `mandatory:"true" json:"isSeededInDataSafe"`

	// Indicates whether the policy is already created on the target database.
	IsCreated *bool `mandatory:"true" json:"isCreated"`

	// Indicates on whom the audit policy is enabled.
	EnabledEntities AuditSpecificationEnabledEntitiesEnum `mandatory:"true" json:"enabledEntities"`

	// Provides information about the policy that has been only partially enabled.
	PartiallyEnabledMsg *string `mandatory:"false" json:"partiallyEnabledMsg"`
}

func (m AuditSpecification) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuditSpecification) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAuditPolicyCategoryEnum(string(m.AuditPolicyCategory)); !ok && m.AuditPolicyCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuditPolicyCategory: %s. Supported values are: %s.", m.AuditPolicyCategory, strings.Join(GetAuditPolicyCategoryEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAuditSpecificationEnableStatusEnum(string(m.EnableStatus)); !ok && m.EnableStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EnableStatus: %s. Supported values are: %s.", m.EnableStatus, strings.Join(GetAuditSpecificationEnableStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAuditSpecificationEnabledEntitiesEnum(string(m.EnabledEntities)); !ok && m.EnabledEntities != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EnabledEntities: %s. Supported values are: %s.", m.EnabledEntities, strings.Join(GetAuditSpecificationEnabledEntitiesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AuditSpecificationEnableStatusEnum Enum with underlying type: string
type AuditSpecificationEnableStatusEnum string

// Set of constants representing the allowable values for AuditSpecificationEnableStatusEnum
const (
	AuditSpecificationEnableStatusEnabled          AuditSpecificationEnableStatusEnum = "ENABLED"
	AuditSpecificationEnableStatusDisabled         AuditSpecificationEnableStatusEnum = "DISABLED"
	AuditSpecificationEnableStatusPartiallyEnabled AuditSpecificationEnableStatusEnum = "PARTIALLY_ENABLED"
)

var mappingAuditSpecificationEnableStatusEnum = map[string]AuditSpecificationEnableStatusEnum{
	"ENABLED":           AuditSpecificationEnableStatusEnabled,
	"DISABLED":          AuditSpecificationEnableStatusDisabled,
	"PARTIALLY_ENABLED": AuditSpecificationEnableStatusPartiallyEnabled,
}

var mappingAuditSpecificationEnableStatusEnumLowerCase = map[string]AuditSpecificationEnableStatusEnum{
	"enabled":           AuditSpecificationEnableStatusEnabled,
	"disabled":          AuditSpecificationEnableStatusDisabled,
	"partially_enabled": AuditSpecificationEnableStatusPartiallyEnabled,
}

// GetAuditSpecificationEnableStatusEnumValues Enumerates the set of values for AuditSpecificationEnableStatusEnum
func GetAuditSpecificationEnableStatusEnumValues() []AuditSpecificationEnableStatusEnum {
	values := make([]AuditSpecificationEnableStatusEnum, 0)
	for _, v := range mappingAuditSpecificationEnableStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAuditSpecificationEnableStatusEnumStringValues Enumerates the set of values in String for AuditSpecificationEnableStatusEnum
func GetAuditSpecificationEnableStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"PARTIALLY_ENABLED",
	}
}

// GetMappingAuditSpecificationEnableStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuditSpecificationEnableStatusEnum(val string) (AuditSpecificationEnableStatusEnum, bool) {
	enum, ok := mappingAuditSpecificationEnableStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AuditSpecificationEnabledEntitiesEnum Enum with underlying type: string
type AuditSpecificationEnabledEntitiesEnum string

// Set of constants representing the allowable values for AuditSpecificationEnabledEntitiesEnum
const (
	AuditSpecificationEnabledEntitiesAllUsers          AuditSpecificationEnabledEntitiesEnum = "ALL_USERS"
	AuditSpecificationEnabledEntitiesIncludeUsers      AuditSpecificationEnabledEntitiesEnum = "INCLUDE_USERS"
	AuditSpecificationEnabledEntitiesIncludeRoles      AuditSpecificationEnabledEntitiesEnum = "INCLUDE_ROLES"
	AuditSpecificationEnabledEntitiesExcludeUsers      AuditSpecificationEnabledEntitiesEnum = "EXCLUDE_USERS"
	AuditSpecificationEnabledEntitiesIncludeUsersRoles AuditSpecificationEnabledEntitiesEnum = "INCLUDE_USERS_ROLES"
	AuditSpecificationEnabledEntitiesDisabled          AuditSpecificationEnabledEntitiesEnum = "DISABLED"
)

var mappingAuditSpecificationEnabledEntitiesEnum = map[string]AuditSpecificationEnabledEntitiesEnum{
	"ALL_USERS":           AuditSpecificationEnabledEntitiesAllUsers,
	"INCLUDE_USERS":       AuditSpecificationEnabledEntitiesIncludeUsers,
	"INCLUDE_ROLES":       AuditSpecificationEnabledEntitiesIncludeRoles,
	"EXCLUDE_USERS":       AuditSpecificationEnabledEntitiesExcludeUsers,
	"INCLUDE_USERS_ROLES": AuditSpecificationEnabledEntitiesIncludeUsersRoles,
	"DISABLED":            AuditSpecificationEnabledEntitiesDisabled,
}

var mappingAuditSpecificationEnabledEntitiesEnumLowerCase = map[string]AuditSpecificationEnabledEntitiesEnum{
	"all_users":           AuditSpecificationEnabledEntitiesAllUsers,
	"include_users":       AuditSpecificationEnabledEntitiesIncludeUsers,
	"include_roles":       AuditSpecificationEnabledEntitiesIncludeRoles,
	"exclude_users":       AuditSpecificationEnabledEntitiesExcludeUsers,
	"include_users_roles": AuditSpecificationEnabledEntitiesIncludeUsersRoles,
	"disabled":            AuditSpecificationEnabledEntitiesDisabled,
}

// GetAuditSpecificationEnabledEntitiesEnumValues Enumerates the set of values for AuditSpecificationEnabledEntitiesEnum
func GetAuditSpecificationEnabledEntitiesEnumValues() []AuditSpecificationEnabledEntitiesEnum {
	values := make([]AuditSpecificationEnabledEntitiesEnum, 0)
	for _, v := range mappingAuditSpecificationEnabledEntitiesEnum {
		values = append(values, v)
	}
	return values
}

// GetAuditSpecificationEnabledEntitiesEnumStringValues Enumerates the set of values in String for AuditSpecificationEnabledEntitiesEnum
func GetAuditSpecificationEnabledEntitiesEnumStringValues() []string {
	return []string{
		"ALL_USERS",
		"INCLUDE_USERS",
		"INCLUDE_ROLES",
		"EXCLUDE_USERS",
		"INCLUDE_USERS_ROLES",
		"DISABLED",
	}
}

// GetMappingAuditSpecificationEnabledEntitiesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuditSpecificationEnabledEntitiesEnum(val string) (AuditSpecificationEnabledEntitiesEnum, bool) {
	enum, ok := mappingAuditSpecificationEnabledEntitiesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
