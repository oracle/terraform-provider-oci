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

// UpdateSqlFirewallPolicyDetails Details to update the SQL Firewall policy.
type UpdateSqlFirewallPolicyDetails struct {

	// The display name of the SQL Firewall policy. The name does not have to be unique, and it is changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the SQL Firewall policy.
	Description *string `mandatory:"false" json:"description"`

	// Specifies whether the SQL Firewall policy is enabled or disabled.
	Status UpdateSqlFirewallPolicyDetailsStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Specifies the SQL Firewall policy enforcement option.
	EnforcementScope UpdateSqlFirewallPolicyDetailsEnforcementScopeEnum `mandatory:"false" json:"enforcementScope,omitempty"`

	// Specifies the SQL Firewall action based on detection of SQL Firewall violations.
	ViolationAction UpdateSqlFirewallPolicyDetailsViolationActionEnum `mandatory:"false" json:"violationAction,omitempty"`

	// Specifies whether a unified audit policy should be enabled for auditing the SQL Firewall policy violations.
	ViolationAudit UpdateSqlFirewallPolicyDetailsViolationAuditEnum `mandatory:"false" json:"violationAudit,omitempty"`

	// List of allowed ip addresses for the SQL Firewall policy.
	AllowedClientIps []string `mandatory:"false" json:"allowedClientIps"`

	// List of allowed operating system user names for the SQL Firewall policy.
	AllowedClientOsUsernames []string `mandatory:"false" json:"allowedClientOsUsernames"`

	// List of allowed client programs for the SQL Firewall policy.
	AllowedClientPrograms []string `mandatory:"false" json:"allowedClientPrograms"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateSqlFirewallPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSqlFirewallPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateSqlFirewallPolicyDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetUpdateSqlFirewallPolicyDetailsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateSqlFirewallPolicyDetailsEnforcementScopeEnum(string(m.EnforcementScope)); !ok && m.EnforcementScope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EnforcementScope: %s. Supported values are: %s.", m.EnforcementScope, strings.Join(GetUpdateSqlFirewallPolicyDetailsEnforcementScopeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateSqlFirewallPolicyDetailsViolationActionEnum(string(m.ViolationAction)); !ok && m.ViolationAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ViolationAction: %s. Supported values are: %s.", m.ViolationAction, strings.Join(GetUpdateSqlFirewallPolicyDetailsViolationActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateSqlFirewallPolicyDetailsViolationAuditEnum(string(m.ViolationAudit)); !ok && m.ViolationAudit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ViolationAudit: %s. Supported values are: %s.", m.ViolationAudit, strings.Join(GetUpdateSqlFirewallPolicyDetailsViolationAuditEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateSqlFirewallPolicyDetailsStatusEnum Enum with underlying type: string
type UpdateSqlFirewallPolicyDetailsStatusEnum string

// Set of constants representing the allowable values for UpdateSqlFirewallPolicyDetailsStatusEnum
const (
	UpdateSqlFirewallPolicyDetailsStatusEnabled  UpdateSqlFirewallPolicyDetailsStatusEnum = "ENABLED"
	UpdateSqlFirewallPolicyDetailsStatusDisabled UpdateSqlFirewallPolicyDetailsStatusEnum = "DISABLED"
)

var mappingUpdateSqlFirewallPolicyDetailsStatusEnum = map[string]UpdateSqlFirewallPolicyDetailsStatusEnum{
	"ENABLED":  UpdateSqlFirewallPolicyDetailsStatusEnabled,
	"DISABLED": UpdateSqlFirewallPolicyDetailsStatusDisabled,
}

var mappingUpdateSqlFirewallPolicyDetailsStatusEnumLowerCase = map[string]UpdateSqlFirewallPolicyDetailsStatusEnum{
	"enabled":  UpdateSqlFirewallPolicyDetailsStatusEnabled,
	"disabled": UpdateSqlFirewallPolicyDetailsStatusDisabled,
}

// GetUpdateSqlFirewallPolicyDetailsStatusEnumValues Enumerates the set of values for UpdateSqlFirewallPolicyDetailsStatusEnum
func GetUpdateSqlFirewallPolicyDetailsStatusEnumValues() []UpdateSqlFirewallPolicyDetailsStatusEnum {
	values := make([]UpdateSqlFirewallPolicyDetailsStatusEnum, 0)
	for _, v := range mappingUpdateSqlFirewallPolicyDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateSqlFirewallPolicyDetailsStatusEnumStringValues Enumerates the set of values in String for UpdateSqlFirewallPolicyDetailsStatusEnum
func GetUpdateSqlFirewallPolicyDetailsStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingUpdateSqlFirewallPolicyDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateSqlFirewallPolicyDetailsStatusEnum(val string) (UpdateSqlFirewallPolicyDetailsStatusEnum, bool) {
	enum, ok := mappingUpdateSqlFirewallPolicyDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateSqlFirewallPolicyDetailsEnforcementScopeEnum Enum with underlying type: string
type UpdateSqlFirewallPolicyDetailsEnforcementScopeEnum string

// Set of constants representing the allowable values for UpdateSqlFirewallPolicyDetailsEnforcementScopeEnum
const (
	UpdateSqlFirewallPolicyDetailsEnforcementScopeContext UpdateSqlFirewallPolicyDetailsEnforcementScopeEnum = "ENFORCE_CONTEXT"
	UpdateSqlFirewallPolicyDetailsEnforcementScopeSql     UpdateSqlFirewallPolicyDetailsEnforcementScopeEnum = "ENFORCE_SQL"
	UpdateSqlFirewallPolicyDetailsEnforcementScopeAll     UpdateSqlFirewallPolicyDetailsEnforcementScopeEnum = "ENFORCE_ALL"
)

var mappingUpdateSqlFirewallPolicyDetailsEnforcementScopeEnum = map[string]UpdateSqlFirewallPolicyDetailsEnforcementScopeEnum{
	"ENFORCE_CONTEXT": UpdateSqlFirewallPolicyDetailsEnforcementScopeContext,
	"ENFORCE_SQL":     UpdateSqlFirewallPolicyDetailsEnforcementScopeSql,
	"ENFORCE_ALL":     UpdateSqlFirewallPolicyDetailsEnforcementScopeAll,
}

var mappingUpdateSqlFirewallPolicyDetailsEnforcementScopeEnumLowerCase = map[string]UpdateSqlFirewallPolicyDetailsEnforcementScopeEnum{
	"enforce_context": UpdateSqlFirewallPolicyDetailsEnforcementScopeContext,
	"enforce_sql":     UpdateSqlFirewallPolicyDetailsEnforcementScopeSql,
	"enforce_all":     UpdateSqlFirewallPolicyDetailsEnforcementScopeAll,
}

// GetUpdateSqlFirewallPolicyDetailsEnforcementScopeEnumValues Enumerates the set of values for UpdateSqlFirewallPolicyDetailsEnforcementScopeEnum
func GetUpdateSqlFirewallPolicyDetailsEnforcementScopeEnumValues() []UpdateSqlFirewallPolicyDetailsEnforcementScopeEnum {
	values := make([]UpdateSqlFirewallPolicyDetailsEnforcementScopeEnum, 0)
	for _, v := range mappingUpdateSqlFirewallPolicyDetailsEnforcementScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateSqlFirewallPolicyDetailsEnforcementScopeEnumStringValues Enumerates the set of values in String for UpdateSqlFirewallPolicyDetailsEnforcementScopeEnum
func GetUpdateSqlFirewallPolicyDetailsEnforcementScopeEnumStringValues() []string {
	return []string{
		"ENFORCE_CONTEXT",
		"ENFORCE_SQL",
		"ENFORCE_ALL",
	}
}

// GetMappingUpdateSqlFirewallPolicyDetailsEnforcementScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateSqlFirewallPolicyDetailsEnforcementScopeEnum(val string) (UpdateSqlFirewallPolicyDetailsEnforcementScopeEnum, bool) {
	enum, ok := mappingUpdateSqlFirewallPolicyDetailsEnforcementScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateSqlFirewallPolicyDetailsViolationActionEnum Enum with underlying type: string
type UpdateSqlFirewallPolicyDetailsViolationActionEnum string

// Set of constants representing the allowable values for UpdateSqlFirewallPolicyDetailsViolationActionEnum
const (
	UpdateSqlFirewallPolicyDetailsViolationActionBlock   UpdateSqlFirewallPolicyDetailsViolationActionEnum = "BLOCK"
	UpdateSqlFirewallPolicyDetailsViolationActionObserve UpdateSqlFirewallPolicyDetailsViolationActionEnum = "OBSERVE"
)

var mappingUpdateSqlFirewallPolicyDetailsViolationActionEnum = map[string]UpdateSqlFirewallPolicyDetailsViolationActionEnum{
	"BLOCK":   UpdateSqlFirewallPolicyDetailsViolationActionBlock,
	"OBSERVE": UpdateSqlFirewallPolicyDetailsViolationActionObserve,
}

var mappingUpdateSqlFirewallPolicyDetailsViolationActionEnumLowerCase = map[string]UpdateSqlFirewallPolicyDetailsViolationActionEnum{
	"block":   UpdateSqlFirewallPolicyDetailsViolationActionBlock,
	"observe": UpdateSqlFirewallPolicyDetailsViolationActionObserve,
}

// GetUpdateSqlFirewallPolicyDetailsViolationActionEnumValues Enumerates the set of values for UpdateSqlFirewallPolicyDetailsViolationActionEnum
func GetUpdateSqlFirewallPolicyDetailsViolationActionEnumValues() []UpdateSqlFirewallPolicyDetailsViolationActionEnum {
	values := make([]UpdateSqlFirewallPolicyDetailsViolationActionEnum, 0)
	for _, v := range mappingUpdateSqlFirewallPolicyDetailsViolationActionEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateSqlFirewallPolicyDetailsViolationActionEnumStringValues Enumerates the set of values in String for UpdateSqlFirewallPolicyDetailsViolationActionEnum
func GetUpdateSqlFirewallPolicyDetailsViolationActionEnumStringValues() []string {
	return []string{
		"BLOCK",
		"OBSERVE",
	}
}

// GetMappingUpdateSqlFirewallPolicyDetailsViolationActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateSqlFirewallPolicyDetailsViolationActionEnum(val string) (UpdateSqlFirewallPolicyDetailsViolationActionEnum, bool) {
	enum, ok := mappingUpdateSqlFirewallPolicyDetailsViolationActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateSqlFirewallPolicyDetailsViolationAuditEnum Enum with underlying type: string
type UpdateSqlFirewallPolicyDetailsViolationAuditEnum string

// Set of constants representing the allowable values for UpdateSqlFirewallPolicyDetailsViolationAuditEnum
const (
	UpdateSqlFirewallPolicyDetailsViolationAuditEnabled  UpdateSqlFirewallPolicyDetailsViolationAuditEnum = "ENABLED"
	UpdateSqlFirewallPolicyDetailsViolationAuditDisabled UpdateSqlFirewallPolicyDetailsViolationAuditEnum = "DISABLED"
)

var mappingUpdateSqlFirewallPolicyDetailsViolationAuditEnum = map[string]UpdateSqlFirewallPolicyDetailsViolationAuditEnum{
	"ENABLED":  UpdateSqlFirewallPolicyDetailsViolationAuditEnabled,
	"DISABLED": UpdateSqlFirewallPolicyDetailsViolationAuditDisabled,
}

var mappingUpdateSqlFirewallPolicyDetailsViolationAuditEnumLowerCase = map[string]UpdateSqlFirewallPolicyDetailsViolationAuditEnum{
	"enabled":  UpdateSqlFirewallPolicyDetailsViolationAuditEnabled,
	"disabled": UpdateSqlFirewallPolicyDetailsViolationAuditDisabled,
}

// GetUpdateSqlFirewallPolicyDetailsViolationAuditEnumValues Enumerates the set of values for UpdateSqlFirewallPolicyDetailsViolationAuditEnum
func GetUpdateSqlFirewallPolicyDetailsViolationAuditEnumValues() []UpdateSqlFirewallPolicyDetailsViolationAuditEnum {
	values := make([]UpdateSqlFirewallPolicyDetailsViolationAuditEnum, 0)
	for _, v := range mappingUpdateSqlFirewallPolicyDetailsViolationAuditEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateSqlFirewallPolicyDetailsViolationAuditEnumStringValues Enumerates the set of values in String for UpdateSqlFirewallPolicyDetailsViolationAuditEnum
func GetUpdateSqlFirewallPolicyDetailsViolationAuditEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingUpdateSqlFirewallPolicyDetailsViolationAuditEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateSqlFirewallPolicyDetailsViolationAuditEnum(val string) (UpdateSqlFirewallPolicyDetailsViolationAuditEnum, bool) {
	enum, ok := mappingUpdateSqlFirewallPolicyDetailsViolationAuditEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
