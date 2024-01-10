// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IngestTimeRuleAdditionalFieldCondition The additional field condition(s) to evaluate for an ingest time rule.
type IngestTimeRuleAdditionalFieldCondition struct {

	// The additional field name to be evaluated.
	ConditionField *string `mandatory:"true" json:"conditionField"`

	// The operator to be used for evaluating the additional field.
	ConditionOperator IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum `mandatory:"true" json:"conditionOperator"`

	// The additional field value to be evaluated.
	ConditionValue *string `mandatory:"true" json:"conditionValue"`
}

func (m IngestTimeRuleAdditionalFieldCondition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IngestTimeRuleAdditionalFieldCondition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIngestTimeRuleAdditionalFieldConditionConditionOperatorEnum(string(m.ConditionOperator)); !ok && m.ConditionOperator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConditionOperator: %s. Supported values are: %s.", m.ConditionOperator, strings.Join(GetIngestTimeRuleAdditionalFieldConditionConditionOperatorEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum Enum with underlying type: string
type IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum string

// Set of constants representing the allowable values for IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum
const (
	IngestTimeRuleAdditionalFieldConditionConditionOperatorContains             IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum = "CONTAINS"
	IngestTimeRuleAdditionalFieldConditionConditionOperatorContainsIgnoreCase   IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum = "CONTAINS_IGNORE_CASE"
	IngestTimeRuleAdditionalFieldConditionConditionOperatorContainsRegex        IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum = "CONTAINS_REGEX"
	IngestTimeRuleAdditionalFieldConditionConditionOperatorContainsOneofRegexes IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum = "CONTAINS_ONEOF_REGEXES"
	IngestTimeRuleAdditionalFieldConditionConditionOperatorEndsWith             IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum = "ENDS_WITH"
	IngestTimeRuleAdditionalFieldConditionConditionOperatorEqual                IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum = "EQUAL"
	IngestTimeRuleAdditionalFieldConditionConditionOperatorEqualIgnoreCase      IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum = "EQUAL_IGNORE_CASE"
	IngestTimeRuleAdditionalFieldConditionConditionOperatorIn                   IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum = "IN"
	IngestTimeRuleAdditionalFieldConditionConditionOperatorInIgnoreCase         IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum = "IN_IGNORE_CASE"
	IngestTimeRuleAdditionalFieldConditionConditionOperatorNotContains          IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum = "NOT_CONTAINS"
	IngestTimeRuleAdditionalFieldConditionConditionOperatorNotEqual             IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum = "NOT_EQUAL"
	IngestTimeRuleAdditionalFieldConditionConditionOperatorNotIn                IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum = "NOT_IN"
	IngestTimeRuleAdditionalFieldConditionConditionOperatorNotNull              IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum = "NOT_NULL"
	IngestTimeRuleAdditionalFieldConditionConditionOperatorStartsWith           IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum = "STARTS_WITH"
)

var mappingIngestTimeRuleAdditionalFieldConditionConditionOperatorEnum = map[string]IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum{
	"CONTAINS":               IngestTimeRuleAdditionalFieldConditionConditionOperatorContains,
	"CONTAINS_IGNORE_CASE":   IngestTimeRuleAdditionalFieldConditionConditionOperatorContainsIgnoreCase,
	"CONTAINS_REGEX":         IngestTimeRuleAdditionalFieldConditionConditionOperatorContainsRegex,
	"CONTAINS_ONEOF_REGEXES": IngestTimeRuleAdditionalFieldConditionConditionOperatorContainsOneofRegexes,
	"ENDS_WITH":              IngestTimeRuleAdditionalFieldConditionConditionOperatorEndsWith,
	"EQUAL":                  IngestTimeRuleAdditionalFieldConditionConditionOperatorEqual,
	"EQUAL_IGNORE_CASE":      IngestTimeRuleAdditionalFieldConditionConditionOperatorEqualIgnoreCase,
	"IN":                     IngestTimeRuleAdditionalFieldConditionConditionOperatorIn,
	"IN_IGNORE_CASE":         IngestTimeRuleAdditionalFieldConditionConditionOperatorInIgnoreCase,
	"NOT_CONTAINS":           IngestTimeRuleAdditionalFieldConditionConditionOperatorNotContains,
	"NOT_EQUAL":              IngestTimeRuleAdditionalFieldConditionConditionOperatorNotEqual,
	"NOT_IN":                 IngestTimeRuleAdditionalFieldConditionConditionOperatorNotIn,
	"NOT_NULL":               IngestTimeRuleAdditionalFieldConditionConditionOperatorNotNull,
	"STARTS_WITH":            IngestTimeRuleAdditionalFieldConditionConditionOperatorStartsWith,
}

var mappingIngestTimeRuleAdditionalFieldConditionConditionOperatorEnumLowerCase = map[string]IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum{
	"contains":               IngestTimeRuleAdditionalFieldConditionConditionOperatorContains,
	"contains_ignore_case":   IngestTimeRuleAdditionalFieldConditionConditionOperatorContainsIgnoreCase,
	"contains_regex":         IngestTimeRuleAdditionalFieldConditionConditionOperatorContainsRegex,
	"contains_oneof_regexes": IngestTimeRuleAdditionalFieldConditionConditionOperatorContainsOneofRegexes,
	"ends_with":              IngestTimeRuleAdditionalFieldConditionConditionOperatorEndsWith,
	"equal":                  IngestTimeRuleAdditionalFieldConditionConditionOperatorEqual,
	"equal_ignore_case":      IngestTimeRuleAdditionalFieldConditionConditionOperatorEqualIgnoreCase,
	"in":                     IngestTimeRuleAdditionalFieldConditionConditionOperatorIn,
	"in_ignore_case":         IngestTimeRuleAdditionalFieldConditionConditionOperatorInIgnoreCase,
	"not_contains":           IngestTimeRuleAdditionalFieldConditionConditionOperatorNotContains,
	"not_equal":              IngestTimeRuleAdditionalFieldConditionConditionOperatorNotEqual,
	"not_in":                 IngestTimeRuleAdditionalFieldConditionConditionOperatorNotIn,
	"not_null":               IngestTimeRuleAdditionalFieldConditionConditionOperatorNotNull,
	"starts_with":            IngestTimeRuleAdditionalFieldConditionConditionOperatorStartsWith,
}

// GetIngestTimeRuleAdditionalFieldConditionConditionOperatorEnumValues Enumerates the set of values for IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum
func GetIngestTimeRuleAdditionalFieldConditionConditionOperatorEnumValues() []IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum {
	values := make([]IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum, 0)
	for _, v := range mappingIngestTimeRuleAdditionalFieldConditionConditionOperatorEnum {
		values = append(values, v)
	}
	return values
}

// GetIngestTimeRuleAdditionalFieldConditionConditionOperatorEnumStringValues Enumerates the set of values in String for IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum
func GetIngestTimeRuleAdditionalFieldConditionConditionOperatorEnumStringValues() []string {
	return []string{
		"CONTAINS",
		"CONTAINS_IGNORE_CASE",
		"CONTAINS_REGEX",
		"CONTAINS_ONEOF_REGEXES",
		"ENDS_WITH",
		"EQUAL",
		"EQUAL_IGNORE_CASE",
		"IN",
		"IN_IGNORE_CASE",
		"NOT_CONTAINS",
		"NOT_EQUAL",
		"NOT_IN",
		"NOT_NULL",
		"STARTS_WITH",
	}
}

// GetMappingIngestTimeRuleAdditionalFieldConditionConditionOperatorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIngestTimeRuleAdditionalFieldConditionConditionOperatorEnum(val string) (IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum, bool) {
	enum, ok := mappingIngestTimeRuleAdditionalFieldConditionConditionOperatorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
