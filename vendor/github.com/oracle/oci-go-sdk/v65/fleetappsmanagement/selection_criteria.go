// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// SelectionCriteria Rule Selection Criteria for DYNAMIC resource selection for a GENERIC fleet.
// Rules define what resources are members of this fleet. All resources that meet the criteria are added automatically.
type SelectionCriteria struct {

	// Match condition for the rule selection.
	// Include resources that match all rules or any of the rules.
	MatchCondition SelectionCriteriaMatchConditionEnum `mandatory:"true" json:"matchCondition"`

	// Rules.
	Rules []Rule `mandatory:"true" json:"rules"`
}

func (m SelectionCriteria) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SelectionCriteria) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSelectionCriteriaMatchConditionEnum(string(m.MatchCondition)); !ok && m.MatchCondition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MatchCondition: %s. Supported values are: %s.", m.MatchCondition, strings.Join(GetSelectionCriteriaMatchConditionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SelectionCriteriaMatchConditionEnum Enum with underlying type: string
type SelectionCriteriaMatchConditionEnum string

// Set of constants representing the allowable values for SelectionCriteriaMatchConditionEnum
const (
	SelectionCriteriaMatchConditionMatchAll SelectionCriteriaMatchConditionEnum = "MATCH_ALL"
	SelectionCriteriaMatchConditionAny      SelectionCriteriaMatchConditionEnum = "ANY"
)

var mappingSelectionCriteriaMatchConditionEnum = map[string]SelectionCriteriaMatchConditionEnum{
	"MATCH_ALL": SelectionCriteriaMatchConditionMatchAll,
	"ANY":       SelectionCriteriaMatchConditionAny,
}

var mappingSelectionCriteriaMatchConditionEnumLowerCase = map[string]SelectionCriteriaMatchConditionEnum{
	"match_all": SelectionCriteriaMatchConditionMatchAll,
	"any":       SelectionCriteriaMatchConditionAny,
}

// GetSelectionCriteriaMatchConditionEnumValues Enumerates the set of values for SelectionCriteriaMatchConditionEnum
func GetSelectionCriteriaMatchConditionEnumValues() []SelectionCriteriaMatchConditionEnum {
	values := make([]SelectionCriteriaMatchConditionEnum, 0)
	for _, v := range mappingSelectionCriteriaMatchConditionEnum {
		values = append(values, v)
	}
	return values
}

// GetSelectionCriteriaMatchConditionEnumStringValues Enumerates the set of values in String for SelectionCriteriaMatchConditionEnum
func GetSelectionCriteriaMatchConditionEnumStringValues() []string {
	return []string{
		"MATCH_ALL",
		"ANY",
	}
}

// GetMappingSelectionCriteriaMatchConditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSelectionCriteriaMatchConditionEnum(val string) (SelectionCriteriaMatchConditionEnum, bool) {
	enum, ok := mappingSelectionCriteriaMatchConditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
