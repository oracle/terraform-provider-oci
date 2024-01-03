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

// SqlFirewallPolicy The SQL Firewall policy resource contains the firewall policy metadata for a single user.
type SqlFirewallPolicy struct {

	// The OCID of the SQL Firewall policy.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the SQL Firewall policy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the SQL Firewall policy.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the security policy corresponding to the SQL Firewall policy.
	SecurityPolicyId *string `mandatory:"true" json:"securityPolicyId"`

	// The database user name.
	DbUserName *string `mandatory:"true" json:"dbUserName"`

	// Specifies whether the SQL Firewall policy is enabled or disabled.
	Status SqlFirewallPolicyStatusEnum `mandatory:"true" json:"status"`

	// The time that the SQL Firewall policy was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the SQL Firewall policy.
	LifecycleState SqlFirewallPolicyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The description of the SQL Firewall policy.
	Description *string `mandatory:"false" json:"description"`

	// Specifies the level of SQL included for this SQL Firewall policy.
	// USER_ISSUED_SQL - User issued SQL statements only.
	// ALL_SQL - Includes all SQL statements including SQL statement issued inside PL/SQL units.
	SqlLevel SqlFirewallPolicySqlLevelEnum `mandatory:"false" json:"sqlLevel,omitempty"`

	// Specifies the SQL Firewall policy enforcement option.
	EnforcementScope SqlFirewallPolicyEnforcementScopeEnum `mandatory:"false" json:"enforcementScope,omitempty"`

	// Specifies the mode in which the SQL Firewall policy is enabled.
	ViolationAction SqlFirewallPolicyViolationActionEnum `mandatory:"false" json:"violationAction,omitempty"`

	// Specifies whether a unified audit policy should be enabled for auditing the SQL Firewall policy violations.
	ViolationAudit SqlFirewallPolicyViolationAuditEnum `mandatory:"false" json:"violationAudit,omitempty"`

	// The list of allowed ip addresses for the SQL Firewall policy.
	AllowedClientIps []string `mandatory:"false" json:"allowedClientIps"`

	// The list of allowed operating system user names for the SQL Firewall policy.
	AllowedClientOsUsernames []string `mandatory:"false" json:"allowedClientOsUsernames"`

	// The list of allowed client programs for the SQL Firewall policy.
	AllowedClientPrograms []string `mandatory:"false" json:"allowedClientPrograms"`

	// The date and time the SQL Firewall policy was last updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Details about the current state of the SQL Firewall policy in Data Safe.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m SqlFirewallPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlFirewallPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSqlFirewallPolicyStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetSqlFirewallPolicyStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlFirewallPolicyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSqlFirewallPolicyLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSqlFirewallPolicySqlLevelEnum(string(m.SqlLevel)); !ok && m.SqlLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SqlLevel: %s. Supported values are: %s.", m.SqlLevel, strings.Join(GetSqlFirewallPolicySqlLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlFirewallPolicyEnforcementScopeEnum(string(m.EnforcementScope)); !ok && m.EnforcementScope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EnforcementScope: %s. Supported values are: %s.", m.EnforcementScope, strings.Join(GetSqlFirewallPolicyEnforcementScopeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlFirewallPolicyViolationActionEnum(string(m.ViolationAction)); !ok && m.ViolationAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ViolationAction: %s. Supported values are: %s.", m.ViolationAction, strings.Join(GetSqlFirewallPolicyViolationActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlFirewallPolicyViolationAuditEnum(string(m.ViolationAudit)); !ok && m.ViolationAudit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ViolationAudit: %s. Supported values are: %s.", m.ViolationAudit, strings.Join(GetSqlFirewallPolicyViolationAuditEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SqlFirewallPolicySqlLevelEnum Enum with underlying type: string
type SqlFirewallPolicySqlLevelEnum string

// Set of constants representing the allowable values for SqlFirewallPolicySqlLevelEnum
const (
	SqlFirewallPolicySqlLevelUserIssuedSql SqlFirewallPolicySqlLevelEnum = "USER_ISSUED_SQL"
	SqlFirewallPolicySqlLevelAllSql        SqlFirewallPolicySqlLevelEnum = "ALL_SQL"
)

var mappingSqlFirewallPolicySqlLevelEnum = map[string]SqlFirewallPolicySqlLevelEnum{
	"USER_ISSUED_SQL": SqlFirewallPolicySqlLevelUserIssuedSql,
	"ALL_SQL":         SqlFirewallPolicySqlLevelAllSql,
}

var mappingSqlFirewallPolicySqlLevelEnumLowerCase = map[string]SqlFirewallPolicySqlLevelEnum{
	"user_issued_sql": SqlFirewallPolicySqlLevelUserIssuedSql,
	"all_sql":         SqlFirewallPolicySqlLevelAllSql,
}

// GetSqlFirewallPolicySqlLevelEnumValues Enumerates the set of values for SqlFirewallPolicySqlLevelEnum
func GetSqlFirewallPolicySqlLevelEnumValues() []SqlFirewallPolicySqlLevelEnum {
	values := make([]SqlFirewallPolicySqlLevelEnum, 0)
	for _, v := range mappingSqlFirewallPolicySqlLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallPolicySqlLevelEnumStringValues Enumerates the set of values in String for SqlFirewallPolicySqlLevelEnum
func GetSqlFirewallPolicySqlLevelEnumStringValues() []string {
	return []string{
		"USER_ISSUED_SQL",
		"ALL_SQL",
	}
}

// GetMappingSqlFirewallPolicySqlLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallPolicySqlLevelEnum(val string) (SqlFirewallPolicySqlLevelEnum, bool) {
	enum, ok := mappingSqlFirewallPolicySqlLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlFirewallPolicyStatusEnum Enum with underlying type: string
type SqlFirewallPolicyStatusEnum string

// Set of constants representing the allowable values for SqlFirewallPolicyStatusEnum
const (
	SqlFirewallPolicyStatusEnabled  SqlFirewallPolicyStatusEnum = "ENABLED"
	SqlFirewallPolicyStatusDisabled SqlFirewallPolicyStatusEnum = "DISABLED"
)

var mappingSqlFirewallPolicyStatusEnum = map[string]SqlFirewallPolicyStatusEnum{
	"ENABLED":  SqlFirewallPolicyStatusEnabled,
	"DISABLED": SqlFirewallPolicyStatusDisabled,
}

var mappingSqlFirewallPolicyStatusEnumLowerCase = map[string]SqlFirewallPolicyStatusEnum{
	"enabled":  SqlFirewallPolicyStatusEnabled,
	"disabled": SqlFirewallPolicyStatusDisabled,
}

// GetSqlFirewallPolicyStatusEnumValues Enumerates the set of values for SqlFirewallPolicyStatusEnum
func GetSqlFirewallPolicyStatusEnumValues() []SqlFirewallPolicyStatusEnum {
	values := make([]SqlFirewallPolicyStatusEnum, 0)
	for _, v := range mappingSqlFirewallPolicyStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallPolicyStatusEnumStringValues Enumerates the set of values in String for SqlFirewallPolicyStatusEnum
func GetSqlFirewallPolicyStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingSqlFirewallPolicyStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallPolicyStatusEnum(val string) (SqlFirewallPolicyStatusEnum, bool) {
	enum, ok := mappingSqlFirewallPolicyStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlFirewallPolicyEnforcementScopeEnum Enum with underlying type: string
type SqlFirewallPolicyEnforcementScopeEnum string

// Set of constants representing the allowable values for SqlFirewallPolicyEnforcementScopeEnum
const (
	SqlFirewallPolicyEnforcementScopeContext SqlFirewallPolicyEnforcementScopeEnum = "ENFORCE_CONTEXT"
	SqlFirewallPolicyEnforcementScopeSql     SqlFirewallPolicyEnforcementScopeEnum = "ENFORCE_SQL"
	SqlFirewallPolicyEnforcementScopeAll     SqlFirewallPolicyEnforcementScopeEnum = "ENFORCE_ALL"
)

var mappingSqlFirewallPolicyEnforcementScopeEnum = map[string]SqlFirewallPolicyEnforcementScopeEnum{
	"ENFORCE_CONTEXT": SqlFirewallPolicyEnforcementScopeContext,
	"ENFORCE_SQL":     SqlFirewallPolicyEnforcementScopeSql,
	"ENFORCE_ALL":     SqlFirewallPolicyEnforcementScopeAll,
}

var mappingSqlFirewallPolicyEnforcementScopeEnumLowerCase = map[string]SqlFirewallPolicyEnforcementScopeEnum{
	"enforce_context": SqlFirewallPolicyEnforcementScopeContext,
	"enforce_sql":     SqlFirewallPolicyEnforcementScopeSql,
	"enforce_all":     SqlFirewallPolicyEnforcementScopeAll,
}

// GetSqlFirewallPolicyEnforcementScopeEnumValues Enumerates the set of values for SqlFirewallPolicyEnforcementScopeEnum
func GetSqlFirewallPolicyEnforcementScopeEnumValues() []SqlFirewallPolicyEnforcementScopeEnum {
	values := make([]SqlFirewallPolicyEnforcementScopeEnum, 0)
	for _, v := range mappingSqlFirewallPolicyEnforcementScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallPolicyEnforcementScopeEnumStringValues Enumerates the set of values in String for SqlFirewallPolicyEnforcementScopeEnum
func GetSqlFirewallPolicyEnforcementScopeEnumStringValues() []string {
	return []string{
		"ENFORCE_CONTEXT",
		"ENFORCE_SQL",
		"ENFORCE_ALL",
	}
}

// GetMappingSqlFirewallPolicyEnforcementScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallPolicyEnforcementScopeEnum(val string) (SqlFirewallPolicyEnforcementScopeEnum, bool) {
	enum, ok := mappingSqlFirewallPolicyEnforcementScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlFirewallPolicyViolationActionEnum Enum with underlying type: string
type SqlFirewallPolicyViolationActionEnum string

// Set of constants representing the allowable values for SqlFirewallPolicyViolationActionEnum
const (
	SqlFirewallPolicyViolationActionBlock   SqlFirewallPolicyViolationActionEnum = "BLOCK"
	SqlFirewallPolicyViolationActionObserve SqlFirewallPolicyViolationActionEnum = "OBSERVE"
)

var mappingSqlFirewallPolicyViolationActionEnum = map[string]SqlFirewallPolicyViolationActionEnum{
	"BLOCK":   SqlFirewallPolicyViolationActionBlock,
	"OBSERVE": SqlFirewallPolicyViolationActionObserve,
}

var mappingSqlFirewallPolicyViolationActionEnumLowerCase = map[string]SqlFirewallPolicyViolationActionEnum{
	"block":   SqlFirewallPolicyViolationActionBlock,
	"observe": SqlFirewallPolicyViolationActionObserve,
}

// GetSqlFirewallPolicyViolationActionEnumValues Enumerates the set of values for SqlFirewallPolicyViolationActionEnum
func GetSqlFirewallPolicyViolationActionEnumValues() []SqlFirewallPolicyViolationActionEnum {
	values := make([]SqlFirewallPolicyViolationActionEnum, 0)
	for _, v := range mappingSqlFirewallPolicyViolationActionEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallPolicyViolationActionEnumStringValues Enumerates the set of values in String for SqlFirewallPolicyViolationActionEnum
func GetSqlFirewallPolicyViolationActionEnumStringValues() []string {
	return []string{
		"BLOCK",
		"OBSERVE",
	}
}

// GetMappingSqlFirewallPolicyViolationActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallPolicyViolationActionEnum(val string) (SqlFirewallPolicyViolationActionEnum, bool) {
	enum, ok := mappingSqlFirewallPolicyViolationActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlFirewallPolicyViolationAuditEnum Enum with underlying type: string
type SqlFirewallPolicyViolationAuditEnum string

// Set of constants representing the allowable values for SqlFirewallPolicyViolationAuditEnum
const (
	SqlFirewallPolicyViolationAuditEnabled  SqlFirewallPolicyViolationAuditEnum = "ENABLED"
	SqlFirewallPolicyViolationAuditDisabled SqlFirewallPolicyViolationAuditEnum = "DISABLED"
)

var mappingSqlFirewallPolicyViolationAuditEnum = map[string]SqlFirewallPolicyViolationAuditEnum{
	"ENABLED":  SqlFirewallPolicyViolationAuditEnabled,
	"DISABLED": SqlFirewallPolicyViolationAuditDisabled,
}

var mappingSqlFirewallPolicyViolationAuditEnumLowerCase = map[string]SqlFirewallPolicyViolationAuditEnum{
	"enabled":  SqlFirewallPolicyViolationAuditEnabled,
	"disabled": SqlFirewallPolicyViolationAuditDisabled,
}

// GetSqlFirewallPolicyViolationAuditEnumValues Enumerates the set of values for SqlFirewallPolicyViolationAuditEnum
func GetSqlFirewallPolicyViolationAuditEnumValues() []SqlFirewallPolicyViolationAuditEnum {
	values := make([]SqlFirewallPolicyViolationAuditEnum, 0)
	for _, v := range mappingSqlFirewallPolicyViolationAuditEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallPolicyViolationAuditEnumStringValues Enumerates the set of values in String for SqlFirewallPolicyViolationAuditEnum
func GetSqlFirewallPolicyViolationAuditEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingSqlFirewallPolicyViolationAuditEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallPolicyViolationAuditEnum(val string) (SqlFirewallPolicyViolationAuditEnum, bool) {
	enum, ok := mappingSqlFirewallPolicyViolationAuditEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
