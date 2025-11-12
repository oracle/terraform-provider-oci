// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Rule Rule for DYNAMIC selection.
type Rule struct {

	// Compartment Id for which the rule is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The Compartment ID to dynamically search resources.
	// Provide the compartment ID to which the rule is applicable.
	ResourceCompartmentId *string `mandatory:"true" json:"resourceCompartmentId"`

	// Rule Conditions
	Conditions []Condition `mandatory:"true" json:"conditions"`

	// Based on what the rule is created.
	// It can be based on a resourceProperty or a tag.
	// If based on a tag, basis will be 'definedTagEquals'
	// If based on a resource property, basis will be 'inventoryProperties'
	Basis *string `mandatory:"false" json:"basis"`

	// Match condition for the rule selection.
	// Include resources that match all rules or any of the rules.
	// Default value for `matchCondition` is ANY
	MatchCondition RuleMatchConditionEnum `mandatory:"false" json:"matchCondition,omitempty"`

	// If set to true, resources will be returned for not only the provided compartment, but all compartments which
	// descend from it. Which resources are returned and their field contents depends on the value of accessLevel.
	// Default value for `compartmentIdInSubtree` is false
	CompartmentIdInSubtree *bool `mandatory:"false" json:"compartmentIdInSubtree"`
}

func (m Rule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Rule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRuleMatchConditionEnum(string(m.MatchCondition)); !ok && m.MatchCondition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MatchCondition: %s. Supported values are: %s.", m.MatchCondition, strings.Join(GetRuleMatchConditionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RuleMatchConditionEnum Enum with underlying type: string
type RuleMatchConditionEnum string

// Set of constants representing the allowable values for RuleMatchConditionEnum
const (
	RuleMatchConditionMatchAll RuleMatchConditionEnum = "MATCH_ALL"
	RuleMatchConditionAny      RuleMatchConditionEnum = "ANY"
)

var mappingRuleMatchConditionEnum = map[string]RuleMatchConditionEnum{
	"MATCH_ALL": RuleMatchConditionMatchAll,
	"ANY":       RuleMatchConditionAny,
}

var mappingRuleMatchConditionEnumLowerCase = map[string]RuleMatchConditionEnum{
	"match_all": RuleMatchConditionMatchAll,
	"any":       RuleMatchConditionAny,
}

// GetRuleMatchConditionEnumValues Enumerates the set of values for RuleMatchConditionEnum
func GetRuleMatchConditionEnumValues() []RuleMatchConditionEnum {
	values := make([]RuleMatchConditionEnum, 0)
	for _, v := range mappingRuleMatchConditionEnum {
		values = append(values, v)
	}
	return values
}

// GetRuleMatchConditionEnumStringValues Enumerates the set of values in String for RuleMatchConditionEnum
func GetRuleMatchConditionEnumStringValues() []string {
	return []string{
		"MATCH_ALL",
		"ANY",
	}
}

// GetMappingRuleMatchConditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuleMatchConditionEnum(val string) (RuleMatchConditionEnum, bool) {
	enum, ok := mappingRuleMatchConditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
