// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// AuditPolicyCategoryEnum Enum with underlying type: string
type AuditPolicyCategoryEnum string

// Set of constants representing the allowable values for AuditPolicyCategoryEnum
const (
	AuditPolicyCategoryBasicActivity       AuditPolicyCategoryEnum = "BASIC_ACTIVITY"
	AuditPolicyCategoryAdminUserActivity   AuditPolicyCategoryEnum = "ADMIN_USER_ACTIVITY"
	AuditPolicyCategoryUserActivity        AuditPolicyCategoryEnum = "USER_ACTIVITY"
	AuditPolicyCategoryOraclePredefined    AuditPolicyCategoryEnum = "ORACLE_PREDEFINED"
	AuditPolicyCategoryComplianceStandard  AuditPolicyCategoryEnum = "COMPLIANCE_STANDARD"
	AuditPolicyCategoryCustom              AuditPolicyCategoryEnum = "CUSTOM"
	AuditPolicyCategorySqlFirewallAuditing AuditPolicyCategoryEnum = "SQL_FIREWALL_AUDITING"
)

var mappingAuditPolicyCategoryEnum = map[string]AuditPolicyCategoryEnum{
	"BASIC_ACTIVITY":        AuditPolicyCategoryBasicActivity,
	"ADMIN_USER_ACTIVITY":   AuditPolicyCategoryAdminUserActivity,
	"USER_ACTIVITY":         AuditPolicyCategoryUserActivity,
	"ORACLE_PREDEFINED":     AuditPolicyCategoryOraclePredefined,
	"COMPLIANCE_STANDARD":   AuditPolicyCategoryComplianceStandard,
	"CUSTOM":                AuditPolicyCategoryCustom,
	"SQL_FIREWALL_AUDITING": AuditPolicyCategorySqlFirewallAuditing,
}

var mappingAuditPolicyCategoryEnumLowerCase = map[string]AuditPolicyCategoryEnum{
	"basic_activity":        AuditPolicyCategoryBasicActivity,
	"admin_user_activity":   AuditPolicyCategoryAdminUserActivity,
	"user_activity":         AuditPolicyCategoryUserActivity,
	"oracle_predefined":     AuditPolicyCategoryOraclePredefined,
	"compliance_standard":   AuditPolicyCategoryComplianceStandard,
	"custom":                AuditPolicyCategoryCustom,
	"sql_firewall_auditing": AuditPolicyCategorySqlFirewallAuditing,
}

// GetAuditPolicyCategoryEnumValues Enumerates the set of values for AuditPolicyCategoryEnum
func GetAuditPolicyCategoryEnumValues() []AuditPolicyCategoryEnum {
	values := make([]AuditPolicyCategoryEnum, 0)
	for _, v := range mappingAuditPolicyCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetAuditPolicyCategoryEnumStringValues Enumerates the set of values in String for AuditPolicyCategoryEnum
func GetAuditPolicyCategoryEnumStringValues() []string {
	return []string{
		"BASIC_ACTIVITY",
		"ADMIN_USER_ACTIVITY",
		"USER_ACTIVITY",
		"ORACLE_PREDEFINED",
		"COMPLIANCE_STANDARD",
		"CUSTOM",
		"SQL_FIREWALL_AUDITING",
	}
}

// GetMappingAuditPolicyCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuditPolicyCategoryEnum(val string) (AuditPolicyCategoryEnum, bool) {
	enum, ok := mappingAuditPolicyCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
