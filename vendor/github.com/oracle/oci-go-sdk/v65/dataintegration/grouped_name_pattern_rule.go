// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GroupedNamePatternRule This rule projects fields as a group recognised as name pattern.
type GroupedNamePatternRule struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Specifies whether the rule uses a java regex syntax.
	IsJavaRegexSyntax *bool `mandatory:"false" json:"isJavaRegexSyntax"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// A user defined description for the object.
	Description *string `mandatory:"false" json:"description"`

	// Name of the group.
	Name *string `mandatory:"false" json:"name"`

	// Specifies whether to skip remaining rules when a match is found.
	IsSkipRemainingRulesOnMatch *bool `mandatory:"false" json:"isSkipRemainingRulesOnMatch"`

	// Reference to a typed object. This can be either a key value to an object within the document, a shall referenced to a `TypedObject`, or a full `TypedObject` definition.
	Scope *interface{} `mandatory:"false" json:"scope"`

	// Specifies whether to cascade or not.
	IsCascade *bool `mandatory:"false" json:"isCascade"`

	// Specifies if the rule is case sensitive.
	IsCaseSensitive *bool `mandatory:"false" json:"isCaseSensitive"`

	// The rule pattern.
	Pattern *string `mandatory:"false" json:"pattern"`

	// The pattern matching strategy.
	MatchingStrategy GroupedNamePatternRuleMatchingStrategyEnum `mandatory:"false" json:"matchingStrategy,omitempty"`

	// The rule type.
	RuleType GroupedNamePatternRuleRuleTypeEnum `mandatory:"false" json:"ruleType,omitempty"`
}

// GetKey returns Key
func (m GroupedNamePatternRule) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m GroupedNamePatternRule) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m GroupedNamePatternRule) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetIsJavaRegexSyntax returns IsJavaRegexSyntax
func (m GroupedNamePatternRule) GetIsJavaRegexSyntax() *bool {
	return m.IsJavaRegexSyntax
}

// GetConfigValues returns ConfigValues
func (m GroupedNamePatternRule) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

// GetObjectStatus returns ObjectStatus
func (m GroupedNamePatternRule) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetDescription returns Description
func (m GroupedNamePatternRule) GetDescription() *string {
	return m.Description
}

func (m GroupedNamePatternRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GroupedNamePatternRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGroupedNamePatternRuleMatchingStrategyEnum(string(m.MatchingStrategy)); !ok && m.MatchingStrategy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MatchingStrategy: %s. Supported values are: %s.", m.MatchingStrategy, strings.Join(GetGroupedNamePatternRuleMatchingStrategyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGroupedNamePatternRuleRuleTypeEnum(string(m.RuleType)); !ok && m.RuleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuleType: %s. Supported values are: %s.", m.RuleType, strings.Join(GetGroupedNamePatternRuleRuleTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GroupedNamePatternRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGroupedNamePatternRule GroupedNamePatternRule
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeGroupedNamePatternRule
	}{
		"GROUPED_NAME_PATTERN_RULE",
		(MarshalTypeGroupedNamePatternRule)(m),
	}

	return json.Marshal(&s)
}

// GroupedNamePatternRuleMatchingStrategyEnum Enum with underlying type: string
type GroupedNamePatternRuleMatchingStrategyEnum string

// Set of constants representing the allowable values for GroupedNamePatternRuleMatchingStrategyEnum
const (
	GroupedNamePatternRuleMatchingStrategyNameOrTags GroupedNamePatternRuleMatchingStrategyEnum = "NAME_OR_TAGS"
	GroupedNamePatternRuleMatchingStrategyTagsOnly   GroupedNamePatternRuleMatchingStrategyEnum = "TAGS_ONLY"
	GroupedNamePatternRuleMatchingStrategyNameOnly   GroupedNamePatternRuleMatchingStrategyEnum = "NAME_ONLY"
)

var mappingGroupedNamePatternRuleMatchingStrategyEnum = map[string]GroupedNamePatternRuleMatchingStrategyEnum{
	"NAME_OR_TAGS": GroupedNamePatternRuleMatchingStrategyNameOrTags,
	"TAGS_ONLY":    GroupedNamePatternRuleMatchingStrategyTagsOnly,
	"NAME_ONLY":    GroupedNamePatternRuleMatchingStrategyNameOnly,
}

var mappingGroupedNamePatternRuleMatchingStrategyEnumLowerCase = map[string]GroupedNamePatternRuleMatchingStrategyEnum{
	"name_or_tags": GroupedNamePatternRuleMatchingStrategyNameOrTags,
	"tags_only":    GroupedNamePatternRuleMatchingStrategyTagsOnly,
	"name_only":    GroupedNamePatternRuleMatchingStrategyNameOnly,
}

// GetGroupedNamePatternRuleMatchingStrategyEnumValues Enumerates the set of values for GroupedNamePatternRuleMatchingStrategyEnum
func GetGroupedNamePatternRuleMatchingStrategyEnumValues() []GroupedNamePatternRuleMatchingStrategyEnum {
	values := make([]GroupedNamePatternRuleMatchingStrategyEnum, 0)
	for _, v := range mappingGroupedNamePatternRuleMatchingStrategyEnum {
		values = append(values, v)
	}
	return values
}

// GetGroupedNamePatternRuleMatchingStrategyEnumStringValues Enumerates the set of values in String for GroupedNamePatternRuleMatchingStrategyEnum
func GetGroupedNamePatternRuleMatchingStrategyEnumStringValues() []string {
	return []string{
		"NAME_OR_TAGS",
		"TAGS_ONLY",
		"NAME_ONLY",
	}
}

// GetMappingGroupedNamePatternRuleMatchingStrategyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGroupedNamePatternRuleMatchingStrategyEnum(val string) (GroupedNamePatternRuleMatchingStrategyEnum, bool) {
	enum, ok := mappingGroupedNamePatternRuleMatchingStrategyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GroupedNamePatternRuleRuleTypeEnum Enum with underlying type: string
type GroupedNamePatternRuleRuleTypeEnum string

// Set of constants representing the allowable values for GroupedNamePatternRuleRuleTypeEnum
const (
	GroupedNamePatternRuleRuleTypeInclude GroupedNamePatternRuleRuleTypeEnum = "INCLUDE"
	GroupedNamePatternRuleRuleTypeExclude GroupedNamePatternRuleRuleTypeEnum = "EXCLUDE"
)

var mappingGroupedNamePatternRuleRuleTypeEnum = map[string]GroupedNamePatternRuleRuleTypeEnum{
	"INCLUDE": GroupedNamePatternRuleRuleTypeInclude,
	"EXCLUDE": GroupedNamePatternRuleRuleTypeExclude,
}

var mappingGroupedNamePatternRuleRuleTypeEnumLowerCase = map[string]GroupedNamePatternRuleRuleTypeEnum{
	"include": GroupedNamePatternRuleRuleTypeInclude,
	"exclude": GroupedNamePatternRuleRuleTypeExclude,
}

// GetGroupedNamePatternRuleRuleTypeEnumValues Enumerates the set of values for GroupedNamePatternRuleRuleTypeEnum
func GetGroupedNamePatternRuleRuleTypeEnumValues() []GroupedNamePatternRuleRuleTypeEnum {
	values := make([]GroupedNamePatternRuleRuleTypeEnum, 0)
	for _, v := range mappingGroupedNamePatternRuleRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGroupedNamePatternRuleRuleTypeEnumStringValues Enumerates the set of values in String for GroupedNamePatternRuleRuleTypeEnum
func GetGroupedNamePatternRuleRuleTypeEnumStringValues() []string {
	return []string{
		"INCLUDE",
		"EXCLUDE",
	}
}

// GetMappingGroupedNamePatternRuleRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGroupedNamePatternRuleRuleTypeEnum(val string) (GroupedNamePatternRuleRuleTypeEnum, bool) {
	enum, ok := mappingGroupedNamePatternRuleRuleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
