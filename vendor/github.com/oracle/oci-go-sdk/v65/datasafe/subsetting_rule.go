// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SubsettingRule Defines a rule for subsetting data in specific tables or schemas, including scope and processing strategies
type SubsettingRule struct {

	// The unique key that identifies a subsetting rule. The key is numeric and unique within a subsetting policy
	Key *string `mandatory:"true" json:"key"`

	Scope SubsetScope `mandatory:"true" json:"scope"`

	SubsetRuleEntry SubsetRuleEntry `mandatory:"true" json:"subsetRuleEntry"`

	// Strategy to be applied while propagating subsetting rule to related tables
	RelatedTablesPropagation SubsettingRuleRelatedTablesPropagationEnum `mandatory:"true" json:"relatedTablesPropagation"`

	// Strategy to be applied while processing peer tables
	PeerTablesAction SubsettingRulePeerTablesActionEnum `mandatory:"true" json:"peerTablesAction"`

	// The description of the subset rule
	Description *string `mandatory:"false" json:"description"`

	// The display name of the subset rule
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Specifies how this rule combines with other rules. UNION evaluates this rule independently and adds matching rows to the result set.
	// SERIAL applies this rule sequentially to filter rows selected by a compatible preceding rule.
	RuleCombinationMode SubsettingRuleRuleCombinationModeEnum `mandatory:"false" json:"ruleCombinationMode,omitempty"`
}

func (m SubsettingRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubsettingRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSubsettingRuleRelatedTablesPropagationEnum(string(m.RelatedTablesPropagation)); !ok && m.RelatedTablesPropagation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RelatedTablesPropagation: %s. Supported values are: %s.", m.RelatedTablesPropagation, strings.Join(GetSubsettingRuleRelatedTablesPropagationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSubsettingRulePeerTablesActionEnum(string(m.PeerTablesAction)); !ok && m.PeerTablesAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PeerTablesAction: %s. Supported values are: %s.", m.PeerTablesAction, strings.Join(GetSubsettingRulePeerTablesActionEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSubsettingRuleRuleCombinationModeEnum(string(m.RuleCombinationMode)); !ok && m.RuleCombinationMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuleCombinationMode: %s. Supported values are: %s.", m.RuleCombinationMode, strings.Join(GetSubsettingRuleRuleCombinationModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *SubsettingRule) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description              *string                                    `json:"description"`
		DisplayName              *string                                    `json:"displayName"`
		RuleCombinationMode      SubsettingRuleRuleCombinationModeEnum      `json:"ruleCombinationMode"`
		Key                      *string                                    `json:"key"`
		Scope                    subsetscope                                `json:"scope"`
		SubsetRuleEntry          subsetruleentry                            `json:"subsetRuleEntry"`
		RelatedTablesPropagation SubsettingRuleRelatedTablesPropagationEnum `json:"relatedTablesPropagation"`
		PeerTablesAction         SubsettingRulePeerTablesActionEnum         `json:"peerTablesAction"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.RuleCombinationMode = model.RuleCombinationMode

	m.Key = model.Key

	nn, e = model.Scope.UnmarshalPolymorphicJSON(model.Scope.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Scope = nn.(SubsetScope)
	} else {
		m.Scope = nil
	}

	nn, e = model.SubsetRuleEntry.UnmarshalPolymorphicJSON(model.SubsetRuleEntry.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SubsetRuleEntry = nn.(SubsetRuleEntry)
	} else {
		m.SubsetRuleEntry = nil
	}

	m.RelatedTablesPropagation = model.RelatedTablesPropagation

	m.PeerTablesAction = model.PeerTablesAction

	return
}

// SubsettingRuleRuleCombinationModeEnum Enum with underlying type: string
type SubsettingRuleRuleCombinationModeEnum string

// Set of constants representing the allowable values for SubsettingRuleRuleCombinationModeEnum
const (
	SubsettingRuleRuleCombinationModeUnion  SubsettingRuleRuleCombinationModeEnum = "UNION"
	SubsettingRuleRuleCombinationModeSerial SubsettingRuleRuleCombinationModeEnum = "SERIAL"
)

var mappingSubsettingRuleRuleCombinationModeEnum = map[string]SubsettingRuleRuleCombinationModeEnum{
	"UNION":  SubsettingRuleRuleCombinationModeUnion,
	"SERIAL": SubsettingRuleRuleCombinationModeSerial,
}

var mappingSubsettingRuleRuleCombinationModeEnumLowerCase = map[string]SubsettingRuleRuleCombinationModeEnum{
	"union":  SubsettingRuleRuleCombinationModeUnion,
	"serial": SubsettingRuleRuleCombinationModeSerial,
}

// GetSubsettingRuleRuleCombinationModeEnumValues Enumerates the set of values for SubsettingRuleRuleCombinationModeEnum
func GetSubsettingRuleRuleCombinationModeEnumValues() []SubsettingRuleRuleCombinationModeEnum {
	values := make([]SubsettingRuleRuleCombinationModeEnum, 0)
	for _, v := range mappingSubsettingRuleRuleCombinationModeEnum {
		values = append(values, v)
	}
	return values
}

// GetSubsettingRuleRuleCombinationModeEnumStringValues Enumerates the set of values in String for SubsettingRuleRuleCombinationModeEnum
func GetSubsettingRuleRuleCombinationModeEnumStringValues() []string {
	return []string{
		"UNION",
		"SERIAL",
	}
}

// GetMappingSubsettingRuleRuleCombinationModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubsettingRuleRuleCombinationModeEnum(val string) (SubsettingRuleRuleCombinationModeEnum, bool) {
	enum, ok := mappingSubsettingRuleRuleCombinationModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SubsettingRuleRelatedTablesPropagationEnum Enum with underlying type: string
type SubsettingRuleRelatedTablesPropagationEnum string

// Set of constants representing the allowable values for SubsettingRuleRelatedTablesPropagationEnum
const (
	SubsettingRuleRelatedTablesPropagationAncestorsAndDescendants SubsettingRuleRelatedTablesPropagationEnum = "ANCESTORS_AND_DESCENDANTS"
	SubsettingRuleRelatedTablesPropagationAncestors               SubsettingRuleRelatedTablesPropagationEnum = "ANCESTORS"
	SubsettingRuleRelatedTablesPropagationDescendants             SubsettingRuleRelatedTablesPropagationEnum = "DESCENDANTS"
	SubsettingRuleRelatedTablesPropagationNone                    SubsettingRuleRelatedTablesPropagationEnum = "NONE"
)

var mappingSubsettingRuleRelatedTablesPropagationEnum = map[string]SubsettingRuleRelatedTablesPropagationEnum{
	"ANCESTORS_AND_DESCENDANTS": SubsettingRuleRelatedTablesPropagationAncestorsAndDescendants,
	"ANCESTORS":                 SubsettingRuleRelatedTablesPropagationAncestors,
	"DESCENDANTS":               SubsettingRuleRelatedTablesPropagationDescendants,
	"NONE":                      SubsettingRuleRelatedTablesPropagationNone,
}

var mappingSubsettingRuleRelatedTablesPropagationEnumLowerCase = map[string]SubsettingRuleRelatedTablesPropagationEnum{
	"ancestors_and_descendants": SubsettingRuleRelatedTablesPropagationAncestorsAndDescendants,
	"ancestors":                 SubsettingRuleRelatedTablesPropagationAncestors,
	"descendants":               SubsettingRuleRelatedTablesPropagationDescendants,
	"none":                      SubsettingRuleRelatedTablesPropagationNone,
}

// GetSubsettingRuleRelatedTablesPropagationEnumValues Enumerates the set of values for SubsettingRuleRelatedTablesPropagationEnum
func GetSubsettingRuleRelatedTablesPropagationEnumValues() []SubsettingRuleRelatedTablesPropagationEnum {
	values := make([]SubsettingRuleRelatedTablesPropagationEnum, 0)
	for _, v := range mappingSubsettingRuleRelatedTablesPropagationEnum {
		values = append(values, v)
	}
	return values
}

// GetSubsettingRuleRelatedTablesPropagationEnumStringValues Enumerates the set of values in String for SubsettingRuleRelatedTablesPropagationEnum
func GetSubsettingRuleRelatedTablesPropagationEnumStringValues() []string {
	return []string{
		"ANCESTORS_AND_DESCENDANTS",
		"ANCESTORS",
		"DESCENDANTS",
		"NONE",
	}
}

// GetMappingSubsettingRuleRelatedTablesPropagationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubsettingRuleRelatedTablesPropagationEnum(val string) (SubsettingRuleRelatedTablesPropagationEnum, bool) {
	enum, ok := mappingSubsettingRuleRelatedTablesPropagationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SubsettingRulePeerTablesActionEnum Enum with underlying type: string
type SubsettingRulePeerTablesActionEnum string

// Set of constants representing the allowable values for SubsettingRulePeerTablesActionEnum
const (
	SubsettingRulePeerTablesActionMinimumRows SubsettingRulePeerTablesActionEnum = "MINIMUM_ROWS"
	SubsettingRulePeerTablesActionSubset      SubsettingRulePeerTablesActionEnum = "SUBSET"
	SubsettingRulePeerTablesActionMaximumRows SubsettingRulePeerTablesActionEnum = "MAXIMUM_ROWS"
)

var mappingSubsettingRulePeerTablesActionEnum = map[string]SubsettingRulePeerTablesActionEnum{
	"MINIMUM_ROWS": SubsettingRulePeerTablesActionMinimumRows,
	"SUBSET":       SubsettingRulePeerTablesActionSubset,
	"MAXIMUM_ROWS": SubsettingRulePeerTablesActionMaximumRows,
}

var mappingSubsettingRulePeerTablesActionEnumLowerCase = map[string]SubsettingRulePeerTablesActionEnum{
	"minimum_rows": SubsettingRulePeerTablesActionMinimumRows,
	"subset":       SubsettingRulePeerTablesActionSubset,
	"maximum_rows": SubsettingRulePeerTablesActionMaximumRows,
}

// GetSubsettingRulePeerTablesActionEnumValues Enumerates the set of values for SubsettingRulePeerTablesActionEnum
func GetSubsettingRulePeerTablesActionEnumValues() []SubsettingRulePeerTablesActionEnum {
	values := make([]SubsettingRulePeerTablesActionEnum, 0)
	for _, v := range mappingSubsettingRulePeerTablesActionEnum {
		values = append(values, v)
	}
	return values
}

// GetSubsettingRulePeerTablesActionEnumStringValues Enumerates the set of values in String for SubsettingRulePeerTablesActionEnum
func GetSubsettingRulePeerTablesActionEnumStringValues() []string {
	return []string{
		"MINIMUM_ROWS",
		"SUBSET",
		"MAXIMUM_ROWS",
	}
}

// GetMappingSubsettingRulePeerTablesActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubsettingRulePeerTablesActionEnum(val string) (SubsettingRulePeerTablesActionEnum, bool) {
	enum, ok := mappingSubsettingRulePeerTablesActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
