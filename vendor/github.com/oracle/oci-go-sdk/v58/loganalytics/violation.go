// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// Violation Violation
type Violation struct {

	// The indices associated with regular expression violations.
	Indexes []Indexes `mandatory:"false" json:"indexes"`

	// The rule description.
	RuleDescription *string `mandatory:"false" json:"ruleDescription"`

	// The rule name.
	RuleName *string `mandatory:"false" json:"ruleName"`

	// The rule remediation.
	RuleRemediation *string `mandatory:"false" json:"ruleRemediation"`

	// The rule type.  Either WARN or ERROR.
	RuleType ViolationRuleTypeEnum `mandatory:"false" json:"ruleType,omitempty"`
}

func (m Violation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Violation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingViolationRuleTypeEnum(string(m.RuleType)); !ok && m.RuleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuleType: %s. Supported values are: %s.", m.RuleType, strings.Join(GetViolationRuleTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ViolationRuleTypeEnum Enum with underlying type: string
type ViolationRuleTypeEnum string

// Set of constants representing the allowable values for ViolationRuleTypeEnum
const (
	ViolationRuleTypeWarn  ViolationRuleTypeEnum = "WARN"
	ViolationRuleTypeError ViolationRuleTypeEnum = "ERROR"
)

var mappingViolationRuleTypeEnum = map[string]ViolationRuleTypeEnum{
	"WARN":  ViolationRuleTypeWarn,
	"ERROR": ViolationRuleTypeError,
}

// GetViolationRuleTypeEnumValues Enumerates the set of values for ViolationRuleTypeEnum
func GetViolationRuleTypeEnumValues() []ViolationRuleTypeEnum {
	values := make([]ViolationRuleTypeEnum, 0)
	for _, v := range mappingViolationRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetViolationRuleTypeEnumStringValues Enumerates the set of values in String for ViolationRuleTypeEnum
func GetViolationRuleTypeEnumStringValues() []string {
	return []string{
		"WARN",
		"ERROR",
	}
}

// GetMappingViolationRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingViolationRuleTypeEnum(val string) (ViolationRuleTypeEnum, bool) {
	mappingViolationRuleTypeEnumIgnoreCase := make(map[string]ViolationRuleTypeEnum)
	for k, v := range mappingViolationRuleTypeEnum {
		mappingViolationRuleTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingViolationRuleTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
