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

// SubsettingRuleSummary Provides summary information for a subsetting rule, which defines how specific tables or schemas are subsetted
type SubsettingRuleSummary struct {

	// The unique key that identifies a subsetting rule. The key is numeric and unique within a subsetting policy
	Key *string `mandatory:"true" json:"key"`

	Scope SubsetScope `mandatory:"true" json:"scope"`

	SubsetRuleEntry SubsetRuleEntry `mandatory:"true" json:"subsetRuleEntry"`

	// Strategy to be applied while propagating subsetting rule to related tables
	RelatedTablesPropagation SubsettingRuleSummaryRelatedTablesPropagationEnum `mandatory:"true" json:"relatedTablesPropagation"`

	// Strategy to be applied while processing peer tables. Peer tables are the tables related to the subset table but not on the ancestor or descendant paths.
	// MINIMUM_ROWS option will retain minimum number of rows in peer tables which are required to maintain referential integrity
	// SUBSET option will propagate the subsetting rule to the peer tables
	// MAXIMUM_ROWS option will retain maximum number of rows in peer tables without breaking the referential integrity
	PeerTablesAction SubsettingRuleSummaryPeerTablesActionEnum `mandatory:"true" json:"peerTablesAction"`

	// The display name of the subsetting rule
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the subsetting rule
	Description *string `mandatory:"false" json:"description"`

	// Specifies how this rule combines with other rules. UNION evaluates this rule independently and adds matching rows to the result set.
	// SERIAL applies this rule sequentially to filter rows selected by a compatible preceding rule.
	RuleCombinationMode SubsettingRuleSummaryRuleCombinationModeEnum `mandatory:"false" json:"ruleCombinationMode,omitempty"`
}

func (m SubsettingRuleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubsettingRuleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSubsettingRuleSummaryRelatedTablesPropagationEnum(string(m.RelatedTablesPropagation)); !ok && m.RelatedTablesPropagation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RelatedTablesPropagation: %s. Supported values are: %s.", m.RelatedTablesPropagation, strings.Join(GetSubsettingRuleSummaryRelatedTablesPropagationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSubsettingRuleSummaryPeerTablesActionEnum(string(m.PeerTablesAction)); !ok && m.PeerTablesAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PeerTablesAction: %s. Supported values are: %s.", m.PeerTablesAction, strings.Join(GetSubsettingRuleSummaryPeerTablesActionEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSubsettingRuleSummaryRuleCombinationModeEnum(string(m.RuleCombinationMode)); !ok && m.RuleCombinationMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuleCombinationMode: %s. Supported values are: %s.", m.RuleCombinationMode, strings.Join(GetSubsettingRuleSummaryRuleCombinationModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *SubsettingRuleSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName              *string                                           `json:"displayName"`
		Description              *string                                           `json:"description"`
		RuleCombinationMode      SubsettingRuleSummaryRuleCombinationModeEnum      `json:"ruleCombinationMode"`
		Key                      *string                                           `json:"key"`
		Scope                    subsetscope                                       `json:"scope"`
		SubsetRuleEntry          subsetruleentry                                   `json:"subsetRuleEntry"`
		RelatedTablesPropagation SubsettingRuleSummaryRelatedTablesPropagationEnum `json:"relatedTablesPropagation"`
		PeerTablesAction         SubsettingRuleSummaryPeerTablesActionEnum         `json:"peerTablesAction"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

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

// SubsettingRuleSummaryRuleCombinationModeEnum Enum with underlying type: string
type SubsettingRuleSummaryRuleCombinationModeEnum string

// Set of constants representing the allowable values for SubsettingRuleSummaryRuleCombinationModeEnum
const (
	SubsettingRuleSummaryRuleCombinationModeUnion  SubsettingRuleSummaryRuleCombinationModeEnum = "UNION"
	SubsettingRuleSummaryRuleCombinationModeSerial SubsettingRuleSummaryRuleCombinationModeEnum = "SERIAL"
)

var mappingSubsettingRuleSummaryRuleCombinationModeEnum = map[string]SubsettingRuleSummaryRuleCombinationModeEnum{
	"UNION":  SubsettingRuleSummaryRuleCombinationModeUnion,
	"SERIAL": SubsettingRuleSummaryRuleCombinationModeSerial,
}

var mappingSubsettingRuleSummaryRuleCombinationModeEnumLowerCase = map[string]SubsettingRuleSummaryRuleCombinationModeEnum{
	"union":  SubsettingRuleSummaryRuleCombinationModeUnion,
	"serial": SubsettingRuleSummaryRuleCombinationModeSerial,
}

// GetSubsettingRuleSummaryRuleCombinationModeEnumValues Enumerates the set of values for SubsettingRuleSummaryRuleCombinationModeEnum
func GetSubsettingRuleSummaryRuleCombinationModeEnumValues() []SubsettingRuleSummaryRuleCombinationModeEnum {
	values := make([]SubsettingRuleSummaryRuleCombinationModeEnum, 0)
	for _, v := range mappingSubsettingRuleSummaryRuleCombinationModeEnum {
		values = append(values, v)
	}
	return values
}

// GetSubsettingRuleSummaryRuleCombinationModeEnumStringValues Enumerates the set of values in String for SubsettingRuleSummaryRuleCombinationModeEnum
func GetSubsettingRuleSummaryRuleCombinationModeEnumStringValues() []string {
	return []string{
		"UNION",
		"SERIAL",
	}
}

// GetMappingSubsettingRuleSummaryRuleCombinationModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubsettingRuleSummaryRuleCombinationModeEnum(val string) (SubsettingRuleSummaryRuleCombinationModeEnum, bool) {
	enum, ok := mappingSubsettingRuleSummaryRuleCombinationModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SubsettingRuleSummaryRelatedTablesPropagationEnum Enum with underlying type: string
type SubsettingRuleSummaryRelatedTablesPropagationEnum string

// Set of constants representing the allowable values for SubsettingRuleSummaryRelatedTablesPropagationEnum
const (
	SubsettingRuleSummaryRelatedTablesPropagationAncestorsAndDescendants SubsettingRuleSummaryRelatedTablesPropagationEnum = "ANCESTORS_AND_DESCENDANTS"
	SubsettingRuleSummaryRelatedTablesPropagationAncestors               SubsettingRuleSummaryRelatedTablesPropagationEnum = "ANCESTORS"
	SubsettingRuleSummaryRelatedTablesPropagationDescendants             SubsettingRuleSummaryRelatedTablesPropagationEnum = "DESCENDANTS"
	SubsettingRuleSummaryRelatedTablesPropagationNone                    SubsettingRuleSummaryRelatedTablesPropagationEnum = "NONE"
)

var mappingSubsettingRuleSummaryRelatedTablesPropagationEnum = map[string]SubsettingRuleSummaryRelatedTablesPropagationEnum{
	"ANCESTORS_AND_DESCENDANTS": SubsettingRuleSummaryRelatedTablesPropagationAncestorsAndDescendants,
	"ANCESTORS":                 SubsettingRuleSummaryRelatedTablesPropagationAncestors,
	"DESCENDANTS":               SubsettingRuleSummaryRelatedTablesPropagationDescendants,
	"NONE":                      SubsettingRuleSummaryRelatedTablesPropagationNone,
}

var mappingSubsettingRuleSummaryRelatedTablesPropagationEnumLowerCase = map[string]SubsettingRuleSummaryRelatedTablesPropagationEnum{
	"ancestors_and_descendants": SubsettingRuleSummaryRelatedTablesPropagationAncestorsAndDescendants,
	"ancestors":                 SubsettingRuleSummaryRelatedTablesPropagationAncestors,
	"descendants":               SubsettingRuleSummaryRelatedTablesPropagationDescendants,
	"none":                      SubsettingRuleSummaryRelatedTablesPropagationNone,
}

// GetSubsettingRuleSummaryRelatedTablesPropagationEnumValues Enumerates the set of values for SubsettingRuleSummaryRelatedTablesPropagationEnum
func GetSubsettingRuleSummaryRelatedTablesPropagationEnumValues() []SubsettingRuleSummaryRelatedTablesPropagationEnum {
	values := make([]SubsettingRuleSummaryRelatedTablesPropagationEnum, 0)
	for _, v := range mappingSubsettingRuleSummaryRelatedTablesPropagationEnum {
		values = append(values, v)
	}
	return values
}

// GetSubsettingRuleSummaryRelatedTablesPropagationEnumStringValues Enumerates the set of values in String for SubsettingRuleSummaryRelatedTablesPropagationEnum
func GetSubsettingRuleSummaryRelatedTablesPropagationEnumStringValues() []string {
	return []string{
		"ANCESTORS_AND_DESCENDANTS",
		"ANCESTORS",
		"DESCENDANTS",
		"NONE",
	}
}

// GetMappingSubsettingRuleSummaryRelatedTablesPropagationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubsettingRuleSummaryRelatedTablesPropagationEnum(val string) (SubsettingRuleSummaryRelatedTablesPropagationEnum, bool) {
	enum, ok := mappingSubsettingRuleSummaryRelatedTablesPropagationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SubsettingRuleSummaryPeerTablesActionEnum Enum with underlying type: string
type SubsettingRuleSummaryPeerTablesActionEnum string

// Set of constants representing the allowable values for SubsettingRuleSummaryPeerTablesActionEnum
const (
	SubsettingRuleSummaryPeerTablesActionMinimumRows SubsettingRuleSummaryPeerTablesActionEnum = "MINIMUM_ROWS"
	SubsettingRuleSummaryPeerTablesActionSubset      SubsettingRuleSummaryPeerTablesActionEnum = "SUBSET"
	SubsettingRuleSummaryPeerTablesActionMaximumRows SubsettingRuleSummaryPeerTablesActionEnum = "MAXIMUM_ROWS"
)

var mappingSubsettingRuleSummaryPeerTablesActionEnum = map[string]SubsettingRuleSummaryPeerTablesActionEnum{
	"MINIMUM_ROWS": SubsettingRuleSummaryPeerTablesActionMinimumRows,
	"SUBSET":       SubsettingRuleSummaryPeerTablesActionSubset,
	"MAXIMUM_ROWS": SubsettingRuleSummaryPeerTablesActionMaximumRows,
}

var mappingSubsettingRuleSummaryPeerTablesActionEnumLowerCase = map[string]SubsettingRuleSummaryPeerTablesActionEnum{
	"minimum_rows": SubsettingRuleSummaryPeerTablesActionMinimumRows,
	"subset":       SubsettingRuleSummaryPeerTablesActionSubset,
	"maximum_rows": SubsettingRuleSummaryPeerTablesActionMaximumRows,
}

// GetSubsettingRuleSummaryPeerTablesActionEnumValues Enumerates the set of values for SubsettingRuleSummaryPeerTablesActionEnum
func GetSubsettingRuleSummaryPeerTablesActionEnumValues() []SubsettingRuleSummaryPeerTablesActionEnum {
	values := make([]SubsettingRuleSummaryPeerTablesActionEnum, 0)
	for _, v := range mappingSubsettingRuleSummaryPeerTablesActionEnum {
		values = append(values, v)
	}
	return values
}

// GetSubsettingRuleSummaryPeerTablesActionEnumStringValues Enumerates the set of values in String for SubsettingRuleSummaryPeerTablesActionEnum
func GetSubsettingRuleSummaryPeerTablesActionEnumStringValues() []string {
	return []string{
		"MINIMUM_ROWS",
		"SUBSET",
		"MAXIMUM_ROWS",
	}
}

// GetMappingSubsettingRuleSummaryPeerTablesActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubsettingRuleSummaryPeerTablesActionEnum(val string) (SubsettingRuleSummaryPeerTablesActionEnum, bool) {
	enum, ok := mappingSubsettingRuleSummaryPeerTablesActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
