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

// SqlFirewallPolicyDimensions The dimensions available for SQL Firewall policy analytics.
type SqlFirewallPolicyDimensions struct {

	// The OCID of the security policy corresponding to the SQL Firewall policy.
	SecurityPolicyId *string `mandatory:"false" json:"securityPolicyId"`

	// Specifies the SQL Firewall policy enforcement option.
	EnforcementScope SqlFirewallPolicyDimensionsEnforcementScopeEnum `mandatory:"false" json:"enforcementScope,omitempty"`

	// Specifies the mode in which the SQL Firewall policy is enabled.
	ViolationAction SqlFirewallPolicyDimensionsViolationActionEnum `mandatory:"false" json:"violationAction,omitempty"`

	// The current state of the SQL Firewall policy.
	LifecycleState SqlFirewallPolicyLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m SqlFirewallPolicyDimensions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlFirewallPolicyDimensions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSqlFirewallPolicyDimensionsEnforcementScopeEnum(string(m.EnforcementScope)); !ok && m.EnforcementScope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EnforcementScope: %s. Supported values are: %s.", m.EnforcementScope, strings.Join(GetSqlFirewallPolicyDimensionsEnforcementScopeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlFirewallPolicyDimensionsViolationActionEnum(string(m.ViolationAction)); !ok && m.ViolationAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ViolationAction: %s. Supported values are: %s.", m.ViolationAction, strings.Join(GetSqlFirewallPolicyDimensionsViolationActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlFirewallPolicyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSqlFirewallPolicyLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SqlFirewallPolicyDimensionsEnforcementScopeEnum Enum with underlying type: string
type SqlFirewallPolicyDimensionsEnforcementScopeEnum string

// Set of constants representing the allowable values for SqlFirewallPolicyDimensionsEnforcementScopeEnum
const (
	SqlFirewallPolicyDimensionsEnforcementScopeContext SqlFirewallPolicyDimensionsEnforcementScopeEnum = "ENFORCE_CONTEXT"
	SqlFirewallPolicyDimensionsEnforcementScopeSql     SqlFirewallPolicyDimensionsEnforcementScopeEnum = "ENFORCE_SQL"
	SqlFirewallPolicyDimensionsEnforcementScopeAll     SqlFirewallPolicyDimensionsEnforcementScopeEnum = "ENFORCE_ALL"
)

var mappingSqlFirewallPolicyDimensionsEnforcementScopeEnum = map[string]SqlFirewallPolicyDimensionsEnforcementScopeEnum{
	"ENFORCE_CONTEXT": SqlFirewallPolicyDimensionsEnforcementScopeContext,
	"ENFORCE_SQL":     SqlFirewallPolicyDimensionsEnforcementScopeSql,
	"ENFORCE_ALL":     SqlFirewallPolicyDimensionsEnforcementScopeAll,
}

var mappingSqlFirewallPolicyDimensionsEnforcementScopeEnumLowerCase = map[string]SqlFirewallPolicyDimensionsEnforcementScopeEnum{
	"enforce_context": SqlFirewallPolicyDimensionsEnforcementScopeContext,
	"enforce_sql":     SqlFirewallPolicyDimensionsEnforcementScopeSql,
	"enforce_all":     SqlFirewallPolicyDimensionsEnforcementScopeAll,
}

// GetSqlFirewallPolicyDimensionsEnforcementScopeEnumValues Enumerates the set of values for SqlFirewallPolicyDimensionsEnforcementScopeEnum
func GetSqlFirewallPolicyDimensionsEnforcementScopeEnumValues() []SqlFirewallPolicyDimensionsEnforcementScopeEnum {
	values := make([]SqlFirewallPolicyDimensionsEnforcementScopeEnum, 0)
	for _, v := range mappingSqlFirewallPolicyDimensionsEnforcementScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallPolicyDimensionsEnforcementScopeEnumStringValues Enumerates the set of values in String for SqlFirewallPolicyDimensionsEnforcementScopeEnum
func GetSqlFirewallPolicyDimensionsEnforcementScopeEnumStringValues() []string {
	return []string{
		"ENFORCE_CONTEXT",
		"ENFORCE_SQL",
		"ENFORCE_ALL",
	}
}

// GetMappingSqlFirewallPolicyDimensionsEnforcementScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallPolicyDimensionsEnforcementScopeEnum(val string) (SqlFirewallPolicyDimensionsEnforcementScopeEnum, bool) {
	enum, ok := mappingSqlFirewallPolicyDimensionsEnforcementScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlFirewallPolicyDimensionsViolationActionEnum Enum with underlying type: string
type SqlFirewallPolicyDimensionsViolationActionEnum string

// Set of constants representing the allowable values for SqlFirewallPolicyDimensionsViolationActionEnum
const (
	SqlFirewallPolicyDimensionsViolationActionBlock   SqlFirewallPolicyDimensionsViolationActionEnum = "BLOCK"
	SqlFirewallPolicyDimensionsViolationActionObserve SqlFirewallPolicyDimensionsViolationActionEnum = "OBSERVE"
)

var mappingSqlFirewallPolicyDimensionsViolationActionEnum = map[string]SqlFirewallPolicyDimensionsViolationActionEnum{
	"BLOCK":   SqlFirewallPolicyDimensionsViolationActionBlock,
	"OBSERVE": SqlFirewallPolicyDimensionsViolationActionObserve,
}

var mappingSqlFirewallPolicyDimensionsViolationActionEnumLowerCase = map[string]SqlFirewallPolicyDimensionsViolationActionEnum{
	"block":   SqlFirewallPolicyDimensionsViolationActionBlock,
	"observe": SqlFirewallPolicyDimensionsViolationActionObserve,
}

// GetSqlFirewallPolicyDimensionsViolationActionEnumValues Enumerates the set of values for SqlFirewallPolicyDimensionsViolationActionEnum
func GetSqlFirewallPolicyDimensionsViolationActionEnumValues() []SqlFirewallPolicyDimensionsViolationActionEnum {
	values := make([]SqlFirewallPolicyDimensionsViolationActionEnum, 0)
	for _, v := range mappingSqlFirewallPolicyDimensionsViolationActionEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallPolicyDimensionsViolationActionEnumStringValues Enumerates the set of values in String for SqlFirewallPolicyDimensionsViolationActionEnum
func GetSqlFirewallPolicyDimensionsViolationActionEnumStringValues() []string {
	return []string{
		"BLOCK",
		"OBSERVE",
	}
}

// GetMappingSqlFirewallPolicyDimensionsViolationActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallPolicyDimensionsViolationActionEnum(val string) (SqlFirewallPolicyDimensionsViolationActionEnum, bool) {
	enum, ok := mappingSqlFirewallPolicyDimensionsViolationActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
