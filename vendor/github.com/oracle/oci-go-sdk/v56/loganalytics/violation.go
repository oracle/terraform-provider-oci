// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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

// ViolationRuleTypeEnum Enum with underlying type: string
type ViolationRuleTypeEnum string

// Set of constants representing the allowable values for ViolationRuleTypeEnum
const (
	ViolationRuleTypeWarn  ViolationRuleTypeEnum = "WARN"
	ViolationRuleTypeError ViolationRuleTypeEnum = "ERROR"
)

var mappingViolationRuleType = map[string]ViolationRuleTypeEnum{
	"WARN":  ViolationRuleTypeWarn,
	"ERROR": ViolationRuleTypeError,
}

// GetViolationRuleTypeEnumValues Enumerates the set of values for ViolationRuleTypeEnum
func GetViolationRuleTypeEnumValues() []ViolationRuleTypeEnum {
	values := make([]ViolationRuleTypeEnum, 0)
	for _, v := range mappingViolationRuleType {
		values = append(values, v)
	}
	return values
}
