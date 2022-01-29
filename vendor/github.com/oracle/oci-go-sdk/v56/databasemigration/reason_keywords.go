// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

// ReasonKeywordsEnum Enum with underlying type: string
type ReasonKeywordsEnum string

// Set of constants representing the allowable values for ReasonKeywordsEnum
const (
	ReasonKeywordsOracleMaintained  ReasonKeywordsEnum = "ORACLE_MAINTAINED"
	ReasonKeywordsGgUnsupported     ReasonKeywordsEnum = "GG_UNSUPPORTED"
	ReasonKeywordsUserExcluded      ReasonKeywordsEnum = "USER_EXCLUDED"
	ReasonKeywordsMandatoryExcluded ReasonKeywordsEnum = "MANDATORY_EXCLUDED"
	ReasonKeywordsUserExcludedType  ReasonKeywordsEnum = "USER_EXCLUDED_TYPE"
)

var mappingReasonKeywordsEnum = map[string]ReasonKeywordsEnum{
	"ORACLE_MAINTAINED":  ReasonKeywordsOracleMaintained,
	"GG_UNSUPPORTED":     ReasonKeywordsGgUnsupported,
	"USER_EXCLUDED":      ReasonKeywordsUserExcluded,
	"MANDATORY_EXCLUDED": ReasonKeywordsMandatoryExcluded,
	"USER_EXCLUDED_TYPE": ReasonKeywordsUserExcludedType,
}

// GetReasonKeywordsEnumValues Enumerates the set of values for ReasonKeywordsEnum
func GetReasonKeywordsEnumValues() []ReasonKeywordsEnum {
	values := make([]ReasonKeywordsEnum, 0)
	for _, v := range mappingReasonKeywordsEnum {
		values = append(values, v)
	}
	return values
}

// GetReasonKeywordsEnumStringValues Enumerates the set of values in String for ReasonKeywordsEnum
func GetReasonKeywordsEnumStringValues() []string {
	return []string{
		"ORACLE_MAINTAINED",
		"GG_UNSUPPORTED",
		"USER_EXCLUDED",
		"MANDATORY_EXCLUDED",
		"USER_EXCLUDED_TYPE",
	}
}
