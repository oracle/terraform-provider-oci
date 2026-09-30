// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// RegistrationPolicyFeaturesEnum Enum with underlying type: string
type RegistrationPolicyFeaturesEnum string

// Set of constants representing the allowable values for RegistrationPolicyFeaturesEnum
const (
	RegistrationPolicyFeaturesAssessment      RegistrationPolicyFeaturesEnum = "ASSESSMENT"
	RegistrationPolicyFeaturesAuditCollection RegistrationPolicyFeaturesEnum = "AUDIT_COLLECTION"
	RegistrationPolicyFeaturesAuditSetting    RegistrationPolicyFeaturesEnum = "AUDIT_SETTING"
	RegistrationPolicyFeaturesDataDiscovery   RegistrationPolicyFeaturesEnum = "DATA_DISCOVERY"
	RegistrationPolicyFeaturesMasking         RegistrationPolicyFeaturesEnum = "MASKING"
	RegistrationPolicyFeaturesSqlFirewall     RegistrationPolicyFeaturesEnum = "SQL_FIREWALL"
	RegistrationPolicyFeaturesAll             RegistrationPolicyFeaturesEnum = "ALL"
)

var mappingRegistrationPolicyFeaturesEnum = map[string]RegistrationPolicyFeaturesEnum{
	"ASSESSMENT":       RegistrationPolicyFeaturesAssessment,
	"AUDIT_COLLECTION": RegistrationPolicyFeaturesAuditCollection,
	"AUDIT_SETTING":    RegistrationPolicyFeaturesAuditSetting,
	"DATA_DISCOVERY":   RegistrationPolicyFeaturesDataDiscovery,
	"MASKING":          RegistrationPolicyFeaturesMasking,
	"SQL_FIREWALL":     RegistrationPolicyFeaturesSqlFirewall,
	"ALL":              RegistrationPolicyFeaturesAll,
}

var mappingRegistrationPolicyFeaturesEnumLowerCase = map[string]RegistrationPolicyFeaturesEnum{
	"assessment":       RegistrationPolicyFeaturesAssessment,
	"audit_collection": RegistrationPolicyFeaturesAuditCollection,
	"audit_setting":    RegistrationPolicyFeaturesAuditSetting,
	"data_discovery":   RegistrationPolicyFeaturesDataDiscovery,
	"masking":          RegistrationPolicyFeaturesMasking,
	"sql_firewall":     RegistrationPolicyFeaturesSqlFirewall,
	"all":              RegistrationPolicyFeaturesAll,
}

// GetRegistrationPolicyFeaturesEnumValues Enumerates the set of values for RegistrationPolicyFeaturesEnum
func GetRegistrationPolicyFeaturesEnumValues() []RegistrationPolicyFeaturesEnum {
	values := make([]RegistrationPolicyFeaturesEnum, 0)
	for _, v := range mappingRegistrationPolicyFeaturesEnum {
		values = append(values, v)
	}
	return values
}

// GetRegistrationPolicyFeaturesEnumStringValues Enumerates the set of values in String for RegistrationPolicyFeaturesEnum
func GetRegistrationPolicyFeaturesEnumStringValues() []string {
	return []string{
		"ASSESSMENT",
		"AUDIT_COLLECTION",
		"AUDIT_SETTING",
		"DATA_DISCOVERY",
		"MASKING",
		"SQL_FIREWALL",
		"ALL",
	}
}

// GetMappingRegistrationPolicyFeaturesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRegistrationPolicyFeaturesEnum(val string) (RegistrationPolicyFeaturesEnum, bool) {
	enum, ok := mappingRegistrationPolicyFeaturesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
