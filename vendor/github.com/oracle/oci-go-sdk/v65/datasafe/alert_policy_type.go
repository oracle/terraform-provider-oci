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

// AlertPolicyTypeEnum Enum with underlying type: string
type AlertPolicyTypeEnum string

// Set of constants representing the allowable values for AlertPolicyTypeEnum
const (
	AlertPolicyTypeAuditing           AlertPolicyTypeEnum = "AUDITING"
	AlertPolicyTypeSecurityAssessment AlertPolicyTypeEnum = "SECURITY_ASSESSMENT"
	AlertPolicyTypeUserAssessment     AlertPolicyTypeEnum = "USER_ASSESSMENT"
)

var mappingAlertPolicyTypeEnum = map[string]AlertPolicyTypeEnum{
	"AUDITING":            AlertPolicyTypeAuditing,
	"SECURITY_ASSESSMENT": AlertPolicyTypeSecurityAssessment,
	"USER_ASSESSMENT":     AlertPolicyTypeUserAssessment,
}

var mappingAlertPolicyTypeEnumLowerCase = map[string]AlertPolicyTypeEnum{
	"auditing":            AlertPolicyTypeAuditing,
	"security_assessment": AlertPolicyTypeSecurityAssessment,
	"user_assessment":     AlertPolicyTypeUserAssessment,
}

// GetAlertPolicyTypeEnumValues Enumerates the set of values for AlertPolicyTypeEnum
func GetAlertPolicyTypeEnumValues() []AlertPolicyTypeEnum {
	values := make([]AlertPolicyTypeEnum, 0)
	for _, v := range mappingAlertPolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAlertPolicyTypeEnumStringValues Enumerates the set of values in String for AlertPolicyTypeEnum
func GetAlertPolicyTypeEnumStringValues() []string {
	return []string{
		"AUDITING",
		"SECURITY_ASSESSMENT",
		"USER_ASSESSMENT",
	}
}

// GetMappingAlertPolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlertPolicyTypeEnum(val string) (AlertPolicyTypeEnum, bool) {
	enum, ok := mappingAlertPolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
