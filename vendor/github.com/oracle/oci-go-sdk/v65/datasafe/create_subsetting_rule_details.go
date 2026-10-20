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

// CreateSubsettingRuleDetails Details to create the subsetting rule
type CreateSubsettingRuleDetails struct {
	Scope SubsetScope `mandatory:"true" json:"scope"`

	SubsetRuleEntry SubsetRuleEntry `mandatory:"true" json:"subsetRuleEntry"`

	// The display name of the subset rule
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the subset rule
	Description *string `mandatory:"false" json:"description"`

	// Specifies how this rule combines with other rules. UNION evaluates this rule independently and adds matching rows to the result set.
	// SERIAL applies this rule sequentially to filter rows selected by a compatible preceding rule.
	RuleCombinationMode CreateSubsettingRuleDetailsRuleCombinationModeEnum `mandatory:"false" json:"ruleCombinationMode,omitempty"`

	// Strategy to be applied while propagating subsetting rule to related tables
	RelatedTablesPropagation CreateSubsettingRuleDetailsRelatedTablesPropagationEnum `mandatory:"false" json:"relatedTablesPropagation,omitempty"`

	// Strategy to be applied while processing peer tables
	PeerTablesAction CreateSubsettingRuleDetailsPeerTablesActionEnum `mandatory:"false" json:"peerTablesAction,omitempty"`
}

func (m CreateSubsettingRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSubsettingRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateSubsettingRuleDetailsRuleCombinationModeEnum(string(m.RuleCombinationMode)); !ok && m.RuleCombinationMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuleCombinationMode: %s. Supported values are: %s.", m.RuleCombinationMode, strings.Join(GetCreateSubsettingRuleDetailsRuleCombinationModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateSubsettingRuleDetailsRelatedTablesPropagationEnum(string(m.RelatedTablesPropagation)); !ok && m.RelatedTablesPropagation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RelatedTablesPropagation: %s. Supported values are: %s.", m.RelatedTablesPropagation, strings.Join(GetCreateSubsettingRuleDetailsRelatedTablesPropagationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateSubsettingRuleDetailsPeerTablesActionEnum(string(m.PeerTablesAction)); !ok && m.PeerTablesAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PeerTablesAction: %s. Supported values are: %s.", m.PeerTablesAction, strings.Join(GetCreateSubsettingRuleDetailsPeerTablesActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateSubsettingRuleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName              *string                                                 `json:"displayName"`
		Description              *string                                                 `json:"description"`
		RuleCombinationMode      CreateSubsettingRuleDetailsRuleCombinationModeEnum      `json:"ruleCombinationMode"`
		RelatedTablesPropagation CreateSubsettingRuleDetailsRelatedTablesPropagationEnum `json:"relatedTablesPropagation"`
		PeerTablesAction         CreateSubsettingRuleDetailsPeerTablesActionEnum         `json:"peerTablesAction"`
		Scope                    subsetscope                                             `json:"scope"`
		SubsetRuleEntry          subsetruleentry                                         `json:"subsetRuleEntry"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.RuleCombinationMode = model.RuleCombinationMode

	m.RelatedTablesPropagation = model.RelatedTablesPropagation

	m.PeerTablesAction = model.PeerTablesAction

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

	return
}

// CreateSubsettingRuleDetailsRuleCombinationModeEnum Enum with underlying type: string
type CreateSubsettingRuleDetailsRuleCombinationModeEnum string

// Set of constants representing the allowable values for CreateSubsettingRuleDetailsRuleCombinationModeEnum
const (
	CreateSubsettingRuleDetailsRuleCombinationModeUnion  CreateSubsettingRuleDetailsRuleCombinationModeEnum = "UNION"
	CreateSubsettingRuleDetailsRuleCombinationModeSerial CreateSubsettingRuleDetailsRuleCombinationModeEnum = "SERIAL"
)

var mappingCreateSubsettingRuleDetailsRuleCombinationModeEnum = map[string]CreateSubsettingRuleDetailsRuleCombinationModeEnum{
	"UNION":  CreateSubsettingRuleDetailsRuleCombinationModeUnion,
	"SERIAL": CreateSubsettingRuleDetailsRuleCombinationModeSerial,
}

var mappingCreateSubsettingRuleDetailsRuleCombinationModeEnumLowerCase = map[string]CreateSubsettingRuleDetailsRuleCombinationModeEnum{
	"union":  CreateSubsettingRuleDetailsRuleCombinationModeUnion,
	"serial": CreateSubsettingRuleDetailsRuleCombinationModeSerial,
}

// GetCreateSubsettingRuleDetailsRuleCombinationModeEnumValues Enumerates the set of values for CreateSubsettingRuleDetailsRuleCombinationModeEnum
func GetCreateSubsettingRuleDetailsRuleCombinationModeEnumValues() []CreateSubsettingRuleDetailsRuleCombinationModeEnum {
	values := make([]CreateSubsettingRuleDetailsRuleCombinationModeEnum, 0)
	for _, v := range mappingCreateSubsettingRuleDetailsRuleCombinationModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateSubsettingRuleDetailsRuleCombinationModeEnumStringValues Enumerates the set of values in String for CreateSubsettingRuleDetailsRuleCombinationModeEnum
func GetCreateSubsettingRuleDetailsRuleCombinationModeEnumStringValues() []string {
	return []string{
		"UNION",
		"SERIAL",
	}
}

// GetMappingCreateSubsettingRuleDetailsRuleCombinationModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateSubsettingRuleDetailsRuleCombinationModeEnum(val string) (CreateSubsettingRuleDetailsRuleCombinationModeEnum, bool) {
	enum, ok := mappingCreateSubsettingRuleDetailsRuleCombinationModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateSubsettingRuleDetailsRelatedTablesPropagationEnum Enum with underlying type: string
type CreateSubsettingRuleDetailsRelatedTablesPropagationEnum string

// Set of constants representing the allowable values for CreateSubsettingRuleDetailsRelatedTablesPropagationEnum
const (
	CreateSubsettingRuleDetailsRelatedTablesPropagationAncestorsAndDescendants CreateSubsettingRuleDetailsRelatedTablesPropagationEnum = "ANCESTORS_AND_DESCENDANTS"
	CreateSubsettingRuleDetailsRelatedTablesPropagationAncestors               CreateSubsettingRuleDetailsRelatedTablesPropagationEnum = "ANCESTORS"
	CreateSubsettingRuleDetailsRelatedTablesPropagationDescendants             CreateSubsettingRuleDetailsRelatedTablesPropagationEnum = "DESCENDANTS"
	CreateSubsettingRuleDetailsRelatedTablesPropagationNone                    CreateSubsettingRuleDetailsRelatedTablesPropagationEnum = "NONE"
)

var mappingCreateSubsettingRuleDetailsRelatedTablesPropagationEnum = map[string]CreateSubsettingRuleDetailsRelatedTablesPropagationEnum{
	"ANCESTORS_AND_DESCENDANTS": CreateSubsettingRuleDetailsRelatedTablesPropagationAncestorsAndDescendants,
	"ANCESTORS":                 CreateSubsettingRuleDetailsRelatedTablesPropagationAncestors,
	"DESCENDANTS":               CreateSubsettingRuleDetailsRelatedTablesPropagationDescendants,
	"NONE":                      CreateSubsettingRuleDetailsRelatedTablesPropagationNone,
}

var mappingCreateSubsettingRuleDetailsRelatedTablesPropagationEnumLowerCase = map[string]CreateSubsettingRuleDetailsRelatedTablesPropagationEnum{
	"ancestors_and_descendants": CreateSubsettingRuleDetailsRelatedTablesPropagationAncestorsAndDescendants,
	"ancestors":                 CreateSubsettingRuleDetailsRelatedTablesPropagationAncestors,
	"descendants":               CreateSubsettingRuleDetailsRelatedTablesPropagationDescendants,
	"none":                      CreateSubsettingRuleDetailsRelatedTablesPropagationNone,
}

// GetCreateSubsettingRuleDetailsRelatedTablesPropagationEnumValues Enumerates the set of values for CreateSubsettingRuleDetailsRelatedTablesPropagationEnum
func GetCreateSubsettingRuleDetailsRelatedTablesPropagationEnumValues() []CreateSubsettingRuleDetailsRelatedTablesPropagationEnum {
	values := make([]CreateSubsettingRuleDetailsRelatedTablesPropagationEnum, 0)
	for _, v := range mappingCreateSubsettingRuleDetailsRelatedTablesPropagationEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateSubsettingRuleDetailsRelatedTablesPropagationEnumStringValues Enumerates the set of values in String for CreateSubsettingRuleDetailsRelatedTablesPropagationEnum
func GetCreateSubsettingRuleDetailsRelatedTablesPropagationEnumStringValues() []string {
	return []string{
		"ANCESTORS_AND_DESCENDANTS",
		"ANCESTORS",
		"DESCENDANTS",
		"NONE",
	}
}

// GetMappingCreateSubsettingRuleDetailsRelatedTablesPropagationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateSubsettingRuleDetailsRelatedTablesPropagationEnum(val string) (CreateSubsettingRuleDetailsRelatedTablesPropagationEnum, bool) {
	enum, ok := mappingCreateSubsettingRuleDetailsRelatedTablesPropagationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateSubsettingRuleDetailsPeerTablesActionEnum Enum with underlying type: string
type CreateSubsettingRuleDetailsPeerTablesActionEnum string

// Set of constants representing the allowable values for CreateSubsettingRuleDetailsPeerTablesActionEnum
const (
	CreateSubsettingRuleDetailsPeerTablesActionMinimumRows CreateSubsettingRuleDetailsPeerTablesActionEnum = "MINIMUM_ROWS"
	CreateSubsettingRuleDetailsPeerTablesActionSubset      CreateSubsettingRuleDetailsPeerTablesActionEnum = "SUBSET"
	CreateSubsettingRuleDetailsPeerTablesActionMaximumRows CreateSubsettingRuleDetailsPeerTablesActionEnum = "MAXIMUM_ROWS"
)

var mappingCreateSubsettingRuleDetailsPeerTablesActionEnum = map[string]CreateSubsettingRuleDetailsPeerTablesActionEnum{
	"MINIMUM_ROWS": CreateSubsettingRuleDetailsPeerTablesActionMinimumRows,
	"SUBSET":       CreateSubsettingRuleDetailsPeerTablesActionSubset,
	"MAXIMUM_ROWS": CreateSubsettingRuleDetailsPeerTablesActionMaximumRows,
}

var mappingCreateSubsettingRuleDetailsPeerTablesActionEnumLowerCase = map[string]CreateSubsettingRuleDetailsPeerTablesActionEnum{
	"minimum_rows": CreateSubsettingRuleDetailsPeerTablesActionMinimumRows,
	"subset":       CreateSubsettingRuleDetailsPeerTablesActionSubset,
	"maximum_rows": CreateSubsettingRuleDetailsPeerTablesActionMaximumRows,
}

// GetCreateSubsettingRuleDetailsPeerTablesActionEnumValues Enumerates the set of values for CreateSubsettingRuleDetailsPeerTablesActionEnum
func GetCreateSubsettingRuleDetailsPeerTablesActionEnumValues() []CreateSubsettingRuleDetailsPeerTablesActionEnum {
	values := make([]CreateSubsettingRuleDetailsPeerTablesActionEnum, 0)
	for _, v := range mappingCreateSubsettingRuleDetailsPeerTablesActionEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateSubsettingRuleDetailsPeerTablesActionEnumStringValues Enumerates the set of values in String for CreateSubsettingRuleDetailsPeerTablesActionEnum
func GetCreateSubsettingRuleDetailsPeerTablesActionEnumStringValues() []string {
	return []string{
		"MINIMUM_ROWS",
		"SUBSET",
		"MAXIMUM_ROWS",
	}
}

// GetMappingCreateSubsettingRuleDetailsPeerTablesActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateSubsettingRuleDetailsPeerTablesActionEnum(val string) (CreateSubsettingRuleDetailsPeerTablesActionEnum, bool) {
	enum, ok := mappingCreateSubsettingRuleDetailsPeerTablesActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
