// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// SqlFirewallPolicySummary The SQL Firewall policy resource contains the firewall policy metadata for a single user.
type SqlFirewallPolicySummary struct {

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
	Status SqlFirewallPolicySummaryStatusEnum `mandatory:"true" json:"status"`

	// The time that the SQL Firewall policy was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the SQL Firewall policy.
	LifecycleState SqlFirewallPolicyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The description of the SQL Firewall policy.
	Description *string `mandatory:"false" json:"description"`

	// Specifies the level of SQL included for this SQL Firewall policy.
	// USER_ISSUED_SQL - User issued SQL statements only.
	// ALL_SQL - Includes all SQL statements including SQL statement issued inside PL/SQL units.
	SqlLevel SqlFirewallPolicySummarySqlLevelEnum `mandatory:"false" json:"sqlLevel,omitempty"`

	// Specifies the SQL Firewall policy enforcement option.
	EnforcementScope SqlFirewallPolicySummaryEnforcementScopeEnum `mandatory:"false" json:"enforcementScope,omitempty"`

	// Specifies the SQL Firewall action based on detection of SQL Firewall violations.
	ViolationAction SqlFirewallPolicySummaryViolationActionEnum `mandatory:"false" json:"violationAction,omitempty"`

	// Specifies whether a unified audit policy should be enabled for auditing the SQL Firewall policy violations.
	ViolationAudit SqlFirewallPolicySummaryViolationAuditEnum `mandatory:"false" json:"violationAudit,omitempty"`

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
}

func (m SqlFirewallPolicySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlFirewallPolicySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSqlFirewallPolicySummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetSqlFirewallPolicySummaryStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlFirewallPolicyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSqlFirewallPolicyLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSqlFirewallPolicySummarySqlLevelEnum(string(m.SqlLevel)); !ok && m.SqlLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SqlLevel: %s. Supported values are: %s.", m.SqlLevel, strings.Join(GetSqlFirewallPolicySummarySqlLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlFirewallPolicySummaryEnforcementScopeEnum(string(m.EnforcementScope)); !ok && m.EnforcementScope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EnforcementScope: %s. Supported values are: %s.", m.EnforcementScope, strings.Join(GetSqlFirewallPolicySummaryEnforcementScopeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlFirewallPolicySummaryViolationActionEnum(string(m.ViolationAction)); !ok && m.ViolationAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ViolationAction: %s. Supported values are: %s.", m.ViolationAction, strings.Join(GetSqlFirewallPolicySummaryViolationActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlFirewallPolicySummaryViolationAuditEnum(string(m.ViolationAudit)); !ok && m.ViolationAudit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ViolationAudit: %s. Supported values are: %s.", m.ViolationAudit, strings.Join(GetSqlFirewallPolicySummaryViolationAuditEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SqlFirewallPolicySummarySqlLevelEnum Enum with underlying type: string
type SqlFirewallPolicySummarySqlLevelEnum string

// Set of constants representing the allowable values for SqlFirewallPolicySummarySqlLevelEnum
const (
	SqlFirewallPolicySummarySqlLevelUserIssuedSql SqlFirewallPolicySummarySqlLevelEnum = "USER_ISSUED_SQL"
	SqlFirewallPolicySummarySqlLevelAllSql        SqlFirewallPolicySummarySqlLevelEnum = "ALL_SQL"
)

var mappingSqlFirewallPolicySummarySqlLevelEnum = map[string]SqlFirewallPolicySummarySqlLevelEnum{
	"USER_ISSUED_SQL": SqlFirewallPolicySummarySqlLevelUserIssuedSql,
	"ALL_SQL":         SqlFirewallPolicySummarySqlLevelAllSql,
}

var mappingSqlFirewallPolicySummarySqlLevelEnumLowerCase = map[string]SqlFirewallPolicySummarySqlLevelEnum{
	"user_issued_sql": SqlFirewallPolicySummarySqlLevelUserIssuedSql,
	"all_sql":         SqlFirewallPolicySummarySqlLevelAllSql,
}

// GetSqlFirewallPolicySummarySqlLevelEnumValues Enumerates the set of values for SqlFirewallPolicySummarySqlLevelEnum
func GetSqlFirewallPolicySummarySqlLevelEnumValues() []SqlFirewallPolicySummarySqlLevelEnum {
	values := make([]SqlFirewallPolicySummarySqlLevelEnum, 0)
	for _, v := range mappingSqlFirewallPolicySummarySqlLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallPolicySummarySqlLevelEnumStringValues Enumerates the set of values in String for SqlFirewallPolicySummarySqlLevelEnum
func GetSqlFirewallPolicySummarySqlLevelEnumStringValues() []string {
	return []string{
		"USER_ISSUED_SQL",
		"ALL_SQL",
	}
}

// GetMappingSqlFirewallPolicySummarySqlLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallPolicySummarySqlLevelEnum(val string) (SqlFirewallPolicySummarySqlLevelEnum, bool) {
	enum, ok := mappingSqlFirewallPolicySummarySqlLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlFirewallPolicySummaryStatusEnum Enum with underlying type: string
type SqlFirewallPolicySummaryStatusEnum string

// Set of constants representing the allowable values for SqlFirewallPolicySummaryStatusEnum
const (
	SqlFirewallPolicySummaryStatusEnabled  SqlFirewallPolicySummaryStatusEnum = "ENABLED"
	SqlFirewallPolicySummaryStatusDisabled SqlFirewallPolicySummaryStatusEnum = "DISABLED"
)

var mappingSqlFirewallPolicySummaryStatusEnum = map[string]SqlFirewallPolicySummaryStatusEnum{
	"ENABLED":  SqlFirewallPolicySummaryStatusEnabled,
	"DISABLED": SqlFirewallPolicySummaryStatusDisabled,
}

var mappingSqlFirewallPolicySummaryStatusEnumLowerCase = map[string]SqlFirewallPolicySummaryStatusEnum{
	"enabled":  SqlFirewallPolicySummaryStatusEnabled,
	"disabled": SqlFirewallPolicySummaryStatusDisabled,
}

// GetSqlFirewallPolicySummaryStatusEnumValues Enumerates the set of values for SqlFirewallPolicySummaryStatusEnum
func GetSqlFirewallPolicySummaryStatusEnumValues() []SqlFirewallPolicySummaryStatusEnum {
	values := make([]SqlFirewallPolicySummaryStatusEnum, 0)
	for _, v := range mappingSqlFirewallPolicySummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallPolicySummaryStatusEnumStringValues Enumerates the set of values in String for SqlFirewallPolicySummaryStatusEnum
func GetSqlFirewallPolicySummaryStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingSqlFirewallPolicySummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallPolicySummaryStatusEnum(val string) (SqlFirewallPolicySummaryStatusEnum, bool) {
	enum, ok := mappingSqlFirewallPolicySummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlFirewallPolicySummaryEnforcementScopeEnum Enum with underlying type: string
type SqlFirewallPolicySummaryEnforcementScopeEnum string

// Set of constants representing the allowable values for SqlFirewallPolicySummaryEnforcementScopeEnum
const (
	SqlFirewallPolicySummaryEnforcementScopeContext SqlFirewallPolicySummaryEnforcementScopeEnum = "ENFORCE_CONTEXT"
	SqlFirewallPolicySummaryEnforcementScopeSql     SqlFirewallPolicySummaryEnforcementScopeEnum = "ENFORCE_SQL"
	SqlFirewallPolicySummaryEnforcementScopeAll     SqlFirewallPolicySummaryEnforcementScopeEnum = "ENFORCE_ALL"
)

var mappingSqlFirewallPolicySummaryEnforcementScopeEnum = map[string]SqlFirewallPolicySummaryEnforcementScopeEnum{
	"ENFORCE_CONTEXT": SqlFirewallPolicySummaryEnforcementScopeContext,
	"ENFORCE_SQL":     SqlFirewallPolicySummaryEnforcementScopeSql,
	"ENFORCE_ALL":     SqlFirewallPolicySummaryEnforcementScopeAll,
}

var mappingSqlFirewallPolicySummaryEnforcementScopeEnumLowerCase = map[string]SqlFirewallPolicySummaryEnforcementScopeEnum{
	"enforce_context": SqlFirewallPolicySummaryEnforcementScopeContext,
	"enforce_sql":     SqlFirewallPolicySummaryEnforcementScopeSql,
	"enforce_all":     SqlFirewallPolicySummaryEnforcementScopeAll,
}

// GetSqlFirewallPolicySummaryEnforcementScopeEnumValues Enumerates the set of values for SqlFirewallPolicySummaryEnforcementScopeEnum
func GetSqlFirewallPolicySummaryEnforcementScopeEnumValues() []SqlFirewallPolicySummaryEnforcementScopeEnum {
	values := make([]SqlFirewallPolicySummaryEnforcementScopeEnum, 0)
	for _, v := range mappingSqlFirewallPolicySummaryEnforcementScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallPolicySummaryEnforcementScopeEnumStringValues Enumerates the set of values in String for SqlFirewallPolicySummaryEnforcementScopeEnum
func GetSqlFirewallPolicySummaryEnforcementScopeEnumStringValues() []string {
	return []string{
		"ENFORCE_CONTEXT",
		"ENFORCE_SQL",
		"ENFORCE_ALL",
	}
}

// GetMappingSqlFirewallPolicySummaryEnforcementScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallPolicySummaryEnforcementScopeEnum(val string) (SqlFirewallPolicySummaryEnforcementScopeEnum, bool) {
	enum, ok := mappingSqlFirewallPolicySummaryEnforcementScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlFirewallPolicySummaryViolationActionEnum Enum with underlying type: string
type SqlFirewallPolicySummaryViolationActionEnum string

// Set of constants representing the allowable values for SqlFirewallPolicySummaryViolationActionEnum
const (
	SqlFirewallPolicySummaryViolationActionBlock   SqlFirewallPolicySummaryViolationActionEnum = "BLOCK"
	SqlFirewallPolicySummaryViolationActionObserve SqlFirewallPolicySummaryViolationActionEnum = "OBSERVE"
)

var mappingSqlFirewallPolicySummaryViolationActionEnum = map[string]SqlFirewallPolicySummaryViolationActionEnum{
	"BLOCK":   SqlFirewallPolicySummaryViolationActionBlock,
	"OBSERVE": SqlFirewallPolicySummaryViolationActionObserve,
}

var mappingSqlFirewallPolicySummaryViolationActionEnumLowerCase = map[string]SqlFirewallPolicySummaryViolationActionEnum{
	"block":   SqlFirewallPolicySummaryViolationActionBlock,
	"observe": SqlFirewallPolicySummaryViolationActionObserve,
}

// GetSqlFirewallPolicySummaryViolationActionEnumValues Enumerates the set of values for SqlFirewallPolicySummaryViolationActionEnum
func GetSqlFirewallPolicySummaryViolationActionEnumValues() []SqlFirewallPolicySummaryViolationActionEnum {
	values := make([]SqlFirewallPolicySummaryViolationActionEnum, 0)
	for _, v := range mappingSqlFirewallPolicySummaryViolationActionEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallPolicySummaryViolationActionEnumStringValues Enumerates the set of values in String for SqlFirewallPolicySummaryViolationActionEnum
func GetSqlFirewallPolicySummaryViolationActionEnumStringValues() []string {
	return []string{
		"BLOCK",
		"OBSERVE",
	}
}

// GetMappingSqlFirewallPolicySummaryViolationActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallPolicySummaryViolationActionEnum(val string) (SqlFirewallPolicySummaryViolationActionEnum, bool) {
	enum, ok := mappingSqlFirewallPolicySummaryViolationActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlFirewallPolicySummaryViolationAuditEnum Enum with underlying type: string
type SqlFirewallPolicySummaryViolationAuditEnum string

// Set of constants representing the allowable values for SqlFirewallPolicySummaryViolationAuditEnum
const (
	SqlFirewallPolicySummaryViolationAuditEnabled  SqlFirewallPolicySummaryViolationAuditEnum = "ENABLED"
	SqlFirewallPolicySummaryViolationAuditDisabled SqlFirewallPolicySummaryViolationAuditEnum = "DISABLED"
)

var mappingSqlFirewallPolicySummaryViolationAuditEnum = map[string]SqlFirewallPolicySummaryViolationAuditEnum{
	"ENABLED":  SqlFirewallPolicySummaryViolationAuditEnabled,
	"DISABLED": SqlFirewallPolicySummaryViolationAuditDisabled,
}

var mappingSqlFirewallPolicySummaryViolationAuditEnumLowerCase = map[string]SqlFirewallPolicySummaryViolationAuditEnum{
	"enabled":  SqlFirewallPolicySummaryViolationAuditEnabled,
	"disabled": SqlFirewallPolicySummaryViolationAuditDisabled,
}

// GetSqlFirewallPolicySummaryViolationAuditEnumValues Enumerates the set of values for SqlFirewallPolicySummaryViolationAuditEnum
func GetSqlFirewallPolicySummaryViolationAuditEnumValues() []SqlFirewallPolicySummaryViolationAuditEnum {
	values := make([]SqlFirewallPolicySummaryViolationAuditEnum, 0)
	for _, v := range mappingSqlFirewallPolicySummaryViolationAuditEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallPolicySummaryViolationAuditEnumStringValues Enumerates the set of values in String for SqlFirewallPolicySummaryViolationAuditEnum
func GetSqlFirewallPolicySummaryViolationAuditEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingSqlFirewallPolicySummaryViolationAuditEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallPolicySummaryViolationAuditEnum(val string) (SqlFirewallPolicySummaryViolationAuditEnum, bool) {
	enum, ok := mappingSqlFirewallPolicySummaryViolationAuditEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
