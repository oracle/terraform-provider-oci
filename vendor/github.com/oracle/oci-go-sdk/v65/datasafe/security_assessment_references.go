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

// SecurityAssessmentReferencesEnum Enum with underlying type: string
type SecurityAssessmentReferencesEnum string

// Set of constants representing the allowable values for SecurityAssessmentReferencesEnum
const (
	SecurityAssessmentReferencesStig SecurityAssessmentReferencesEnum = "STIG"
	SecurityAssessmentReferencesCis  SecurityAssessmentReferencesEnum = "CIS"
	SecurityAssessmentReferencesGdpr SecurityAssessmentReferencesEnum = "GDPR"
)

var mappingSecurityAssessmentReferencesEnum = map[string]SecurityAssessmentReferencesEnum{
	"STIG": SecurityAssessmentReferencesStig,
	"CIS":  SecurityAssessmentReferencesCis,
	"GDPR": SecurityAssessmentReferencesGdpr,
}

var mappingSecurityAssessmentReferencesEnumLowerCase = map[string]SecurityAssessmentReferencesEnum{
	"stig": SecurityAssessmentReferencesStig,
	"cis":  SecurityAssessmentReferencesCis,
	"gdpr": SecurityAssessmentReferencesGdpr,
}

// GetSecurityAssessmentReferencesEnumValues Enumerates the set of values for SecurityAssessmentReferencesEnum
func GetSecurityAssessmentReferencesEnumValues() []SecurityAssessmentReferencesEnum {
	values := make([]SecurityAssessmentReferencesEnum, 0)
	for _, v := range mappingSecurityAssessmentReferencesEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityAssessmentReferencesEnumStringValues Enumerates the set of values in String for SecurityAssessmentReferencesEnum
func GetSecurityAssessmentReferencesEnumStringValues() []string {
	return []string{
		"STIG",
		"CIS",
		"GDPR",
	}
}

// GetMappingSecurityAssessmentReferencesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityAssessmentReferencesEnum(val string) (SecurityAssessmentReferencesEnum, bool) {
	enum, ok := mappingSecurityAssessmentReferencesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
