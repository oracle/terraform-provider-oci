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

// UpdateSubsettingRuleDetails Details to update the subsetting rule
type UpdateSubsettingRuleDetails struct {

	// The display name of the subset rule
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the subset rule
	Description *string `mandatory:"false" json:"description"`

	Scope SubsetScope `mandatory:"false" json:"scope"`

	SubsetRuleEntry SubsetRuleEntry `mandatory:"false" json:"subsetRuleEntry"`

	// Specifies how this rule combines with other rules. UNION evaluates this rule independently and adds matching rows to the result set.
	// SERIAL applies this rule sequentially to filter rows selected by a compatible preceding rule.
	RuleCombinationMode UpdateSubsettingRuleDetailsRuleCombinationModeEnum `mandatory:"false" json:"ruleCombinationMode,omitempty"`

	// Strategy to be applied while propagating subsetting rule to related tables
	RelatedTablesPropagation UpdateSubsettingRuleDetailsRelatedTablesPropagationEnum `mandatory:"false" json:"relatedTablesPropagation,omitempty"`

	// Strategy to be applied while processing peer tables
	PeerTablesAction UpdateSubsettingRuleDetailsPeerTablesActionEnum `mandatory:"false" json:"peerTablesAction,omitempty"`
}

func (m UpdateSubsettingRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSubsettingRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateSubsettingRuleDetailsRuleCombinationModeEnum(string(m.RuleCombinationMode)); !ok && m.RuleCombinationMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuleCombinationMode: %s. Supported values are: %s.", m.RuleCombinationMode, strings.Join(GetUpdateSubsettingRuleDetailsRuleCombinationModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateSubsettingRuleDetailsRelatedTablesPropagationEnum(string(m.RelatedTablesPropagation)); !ok && m.RelatedTablesPropagation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RelatedTablesPropagation: %s. Supported values are: %s.", m.RelatedTablesPropagation, strings.Join(GetUpdateSubsettingRuleDetailsRelatedTablesPropagationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateSubsettingRuleDetailsPeerTablesActionEnum(string(m.PeerTablesAction)); !ok && m.PeerTablesAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PeerTablesAction: %s. Supported values are: %s.", m.PeerTablesAction, strings.Join(GetUpdateSubsettingRuleDetailsPeerTablesActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateSubsettingRuleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName              *string                                                 `json:"displayName"`
		Description              *string                                                 `json:"description"`
		Scope                    subsetscope                                             `json:"scope"`
		SubsetRuleEntry          subsetruleentry                                         `json:"subsetRuleEntry"`
		RuleCombinationMode      UpdateSubsettingRuleDetailsRuleCombinationModeEnum      `json:"ruleCombinationMode"`
		RelatedTablesPropagation UpdateSubsettingRuleDetailsRelatedTablesPropagationEnum `json:"relatedTablesPropagation"`
		PeerTablesAction         UpdateSubsettingRuleDetailsPeerTablesActionEnum         `json:"peerTablesAction"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

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

	m.RuleCombinationMode = model.RuleCombinationMode

	m.RelatedTablesPropagation = model.RelatedTablesPropagation

	m.PeerTablesAction = model.PeerTablesAction

	return
}

// UpdateSubsettingRuleDetailsRuleCombinationModeEnum Enum with underlying type: string
type UpdateSubsettingRuleDetailsRuleCombinationModeEnum string

// Set of constants representing the allowable values for UpdateSubsettingRuleDetailsRuleCombinationModeEnum
const (
	UpdateSubsettingRuleDetailsRuleCombinationModeUnion  UpdateSubsettingRuleDetailsRuleCombinationModeEnum = "UNION"
	UpdateSubsettingRuleDetailsRuleCombinationModeSerial UpdateSubsettingRuleDetailsRuleCombinationModeEnum = "SERIAL"
)

var mappingUpdateSubsettingRuleDetailsRuleCombinationModeEnum = map[string]UpdateSubsettingRuleDetailsRuleCombinationModeEnum{
	"UNION":  UpdateSubsettingRuleDetailsRuleCombinationModeUnion,
	"SERIAL": UpdateSubsettingRuleDetailsRuleCombinationModeSerial,
}

var mappingUpdateSubsettingRuleDetailsRuleCombinationModeEnumLowerCase = map[string]UpdateSubsettingRuleDetailsRuleCombinationModeEnum{
	"union":  UpdateSubsettingRuleDetailsRuleCombinationModeUnion,
	"serial": UpdateSubsettingRuleDetailsRuleCombinationModeSerial,
}

// GetUpdateSubsettingRuleDetailsRuleCombinationModeEnumValues Enumerates the set of values for UpdateSubsettingRuleDetailsRuleCombinationModeEnum
func GetUpdateSubsettingRuleDetailsRuleCombinationModeEnumValues() []UpdateSubsettingRuleDetailsRuleCombinationModeEnum {
	values := make([]UpdateSubsettingRuleDetailsRuleCombinationModeEnum, 0)
	for _, v := range mappingUpdateSubsettingRuleDetailsRuleCombinationModeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateSubsettingRuleDetailsRuleCombinationModeEnumStringValues Enumerates the set of values in String for UpdateSubsettingRuleDetailsRuleCombinationModeEnum
func GetUpdateSubsettingRuleDetailsRuleCombinationModeEnumStringValues() []string {
	return []string{
		"UNION",
		"SERIAL",
	}
}

// GetMappingUpdateSubsettingRuleDetailsRuleCombinationModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateSubsettingRuleDetailsRuleCombinationModeEnum(val string) (UpdateSubsettingRuleDetailsRuleCombinationModeEnum, bool) {
	enum, ok := mappingUpdateSubsettingRuleDetailsRuleCombinationModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateSubsettingRuleDetailsRelatedTablesPropagationEnum Enum with underlying type: string
type UpdateSubsettingRuleDetailsRelatedTablesPropagationEnum string

// Set of constants representing the allowable values for UpdateSubsettingRuleDetailsRelatedTablesPropagationEnum
const (
	UpdateSubsettingRuleDetailsRelatedTablesPropagationAncestorsAndDescendants UpdateSubsettingRuleDetailsRelatedTablesPropagationEnum = "ANCESTORS_AND_DESCENDANTS"
	UpdateSubsettingRuleDetailsRelatedTablesPropagationAncestors               UpdateSubsettingRuleDetailsRelatedTablesPropagationEnum = "ANCESTORS"
	UpdateSubsettingRuleDetailsRelatedTablesPropagationDescendants             UpdateSubsettingRuleDetailsRelatedTablesPropagationEnum = "DESCENDANTS"
	UpdateSubsettingRuleDetailsRelatedTablesPropagationNone                    UpdateSubsettingRuleDetailsRelatedTablesPropagationEnum = "NONE"
)

var mappingUpdateSubsettingRuleDetailsRelatedTablesPropagationEnum = map[string]UpdateSubsettingRuleDetailsRelatedTablesPropagationEnum{
	"ANCESTORS_AND_DESCENDANTS": UpdateSubsettingRuleDetailsRelatedTablesPropagationAncestorsAndDescendants,
	"ANCESTORS":                 UpdateSubsettingRuleDetailsRelatedTablesPropagationAncestors,
	"DESCENDANTS":               UpdateSubsettingRuleDetailsRelatedTablesPropagationDescendants,
	"NONE":                      UpdateSubsettingRuleDetailsRelatedTablesPropagationNone,
}

var mappingUpdateSubsettingRuleDetailsRelatedTablesPropagationEnumLowerCase = map[string]UpdateSubsettingRuleDetailsRelatedTablesPropagationEnum{
	"ancestors_and_descendants": UpdateSubsettingRuleDetailsRelatedTablesPropagationAncestorsAndDescendants,
	"ancestors":                 UpdateSubsettingRuleDetailsRelatedTablesPropagationAncestors,
	"descendants":               UpdateSubsettingRuleDetailsRelatedTablesPropagationDescendants,
	"none":                      UpdateSubsettingRuleDetailsRelatedTablesPropagationNone,
}

// GetUpdateSubsettingRuleDetailsRelatedTablesPropagationEnumValues Enumerates the set of values for UpdateSubsettingRuleDetailsRelatedTablesPropagationEnum
func GetUpdateSubsettingRuleDetailsRelatedTablesPropagationEnumValues() []UpdateSubsettingRuleDetailsRelatedTablesPropagationEnum {
	values := make([]UpdateSubsettingRuleDetailsRelatedTablesPropagationEnum, 0)
	for _, v := range mappingUpdateSubsettingRuleDetailsRelatedTablesPropagationEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateSubsettingRuleDetailsRelatedTablesPropagationEnumStringValues Enumerates the set of values in String for UpdateSubsettingRuleDetailsRelatedTablesPropagationEnum
func GetUpdateSubsettingRuleDetailsRelatedTablesPropagationEnumStringValues() []string {
	return []string{
		"ANCESTORS_AND_DESCENDANTS",
		"ANCESTORS",
		"DESCENDANTS",
		"NONE",
	}
}

// GetMappingUpdateSubsettingRuleDetailsRelatedTablesPropagationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateSubsettingRuleDetailsRelatedTablesPropagationEnum(val string) (UpdateSubsettingRuleDetailsRelatedTablesPropagationEnum, bool) {
	enum, ok := mappingUpdateSubsettingRuleDetailsRelatedTablesPropagationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateSubsettingRuleDetailsPeerTablesActionEnum Enum with underlying type: string
type UpdateSubsettingRuleDetailsPeerTablesActionEnum string

// Set of constants representing the allowable values for UpdateSubsettingRuleDetailsPeerTablesActionEnum
const (
	UpdateSubsettingRuleDetailsPeerTablesActionMinimumRows UpdateSubsettingRuleDetailsPeerTablesActionEnum = "MINIMUM_ROWS"
	UpdateSubsettingRuleDetailsPeerTablesActionSubset      UpdateSubsettingRuleDetailsPeerTablesActionEnum = "SUBSET"
	UpdateSubsettingRuleDetailsPeerTablesActionMaximumRows UpdateSubsettingRuleDetailsPeerTablesActionEnum = "MAXIMUM_ROWS"
)

var mappingUpdateSubsettingRuleDetailsPeerTablesActionEnum = map[string]UpdateSubsettingRuleDetailsPeerTablesActionEnum{
	"MINIMUM_ROWS": UpdateSubsettingRuleDetailsPeerTablesActionMinimumRows,
	"SUBSET":       UpdateSubsettingRuleDetailsPeerTablesActionSubset,
	"MAXIMUM_ROWS": UpdateSubsettingRuleDetailsPeerTablesActionMaximumRows,
}

var mappingUpdateSubsettingRuleDetailsPeerTablesActionEnumLowerCase = map[string]UpdateSubsettingRuleDetailsPeerTablesActionEnum{
	"minimum_rows": UpdateSubsettingRuleDetailsPeerTablesActionMinimumRows,
	"subset":       UpdateSubsettingRuleDetailsPeerTablesActionSubset,
	"maximum_rows": UpdateSubsettingRuleDetailsPeerTablesActionMaximumRows,
}

// GetUpdateSubsettingRuleDetailsPeerTablesActionEnumValues Enumerates the set of values for UpdateSubsettingRuleDetailsPeerTablesActionEnum
func GetUpdateSubsettingRuleDetailsPeerTablesActionEnumValues() []UpdateSubsettingRuleDetailsPeerTablesActionEnum {
	values := make([]UpdateSubsettingRuleDetailsPeerTablesActionEnum, 0)
	for _, v := range mappingUpdateSubsettingRuleDetailsPeerTablesActionEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateSubsettingRuleDetailsPeerTablesActionEnumStringValues Enumerates the set of values in String for UpdateSubsettingRuleDetailsPeerTablesActionEnum
func GetUpdateSubsettingRuleDetailsPeerTablesActionEnumStringValues() []string {
	return []string{
		"MINIMUM_ROWS",
		"SUBSET",
		"MAXIMUM_ROWS",
	}
}

// GetMappingUpdateSubsettingRuleDetailsPeerTablesActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateSubsettingRuleDetailsPeerTablesActionEnum(val string) (UpdateSubsettingRuleDetailsPeerTablesActionEnum, bool) {
	enum, ok := mappingUpdateSubsettingRuleDetailsPeerTablesActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
