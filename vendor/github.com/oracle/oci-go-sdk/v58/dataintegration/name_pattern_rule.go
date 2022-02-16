// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// NamePatternRule This rule projects fields by a name pattern, for example it may start with STR_ or end with _DATE. This is defined using a regular expression.
type NamePatternRule struct {

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
	MatchingStrategy NamePatternRuleMatchingStrategyEnum `mandatory:"false" json:"matchingStrategy,omitempty"`

	// The rule type.
	RuleType NamePatternRuleRuleTypeEnum `mandatory:"false" json:"ruleType,omitempty"`
}

//GetKey returns Key
func (m NamePatternRule) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m NamePatternRule) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m NamePatternRule) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetIsJavaRegexSyntax returns IsJavaRegexSyntax
func (m NamePatternRule) GetIsJavaRegexSyntax() *bool {
	return m.IsJavaRegexSyntax
}

//GetConfigValues returns ConfigValues
func (m NamePatternRule) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

//GetObjectStatus returns ObjectStatus
func (m NamePatternRule) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetDescription returns Description
func (m NamePatternRule) GetDescription() *string {
	return m.Description
}

func (m NamePatternRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NamePatternRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNamePatternRuleMatchingStrategyEnum(string(m.MatchingStrategy)); !ok && m.MatchingStrategy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MatchingStrategy: %s. Supported values are: %s.", m.MatchingStrategy, strings.Join(GetNamePatternRuleMatchingStrategyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNamePatternRuleRuleTypeEnum(string(m.RuleType)); !ok && m.RuleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuleType: %s. Supported values are: %s.", m.RuleType, strings.Join(GetNamePatternRuleRuleTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m NamePatternRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNamePatternRule NamePatternRule
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeNamePatternRule
	}{
		"NAME_PATTERN_RULE",
		(MarshalTypeNamePatternRule)(m),
	}

	return json.Marshal(&s)
}

// NamePatternRuleMatchingStrategyEnum Enum with underlying type: string
type NamePatternRuleMatchingStrategyEnum string

// Set of constants representing the allowable values for NamePatternRuleMatchingStrategyEnum
const (
	NamePatternRuleMatchingStrategyNameOrTags NamePatternRuleMatchingStrategyEnum = "NAME_OR_TAGS"
	NamePatternRuleMatchingStrategyTagsOnly   NamePatternRuleMatchingStrategyEnum = "TAGS_ONLY"
	NamePatternRuleMatchingStrategyNameOnly   NamePatternRuleMatchingStrategyEnum = "NAME_ONLY"
)

var mappingNamePatternRuleMatchingStrategyEnum = map[string]NamePatternRuleMatchingStrategyEnum{
	"NAME_OR_TAGS": NamePatternRuleMatchingStrategyNameOrTags,
	"TAGS_ONLY":    NamePatternRuleMatchingStrategyTagsOnly,
	"NAME_ONLY":    NamePatternRuleMatchingStrategyNameOnly,
}

// GetNamePatternRuleMatchingStrategyEnumValues Enumerates the set of values for NamePatternRuleMatchingStrategyEnum
func GetNamePatternRuleMatchingStrategyEnumValues() []NamePatternRuleMatchingStrategyEnum {
	values := make([]NamePatternRuleMatchingStrategyEnum, 0)
	for _, v := range mappingNamePatternRuleMatchingStrategyEnum {
		values = append(values, v)
	}
	return values
}

// GetNamePatternRuleMatchingStrategyEnumStringValues Enumerates the set of values in String for NamePatternRuleMatchingStrategyEnum
func GetNamePatternRuleMatchingStrategyEnumStringValues() []string {
	return []string{
		"NAME_OR_TAGS",
		"TAGS_ONLY",
		"NAME_ONLY",
	}
}

// GetMappingNamePatternRuleMatchingStrategyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNamePatternRuleMatchingStrategyEnum(val string) (NamePatternRuleMatchingStrategyEnum, bool) {
	mappingNamePatternRuleMatchingStrategyEnumIgnoreCase := make(map[string]NamePatternRuleMatchingStrategyEnum)
	for k, v := range mappingNamePatternRuleMatchingStrategyEnum {
		mappingNamePatternRuleMatchingStrategyEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingNamePatternRuleMatchingStrategyEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// NamePatternRuleRuleTypeEnum Enum with underlying type: string
type NamePatternRuleRuleTypeEnum string

// Set of constants representing the allowable values for NamePatternRuleRuleTypeEnum
const (
	NamePatternRuleRuleTypeInclude NamePatternRuleRuleTypeEnum = "INCLUDE"
	NamePatternRuleRuleTypeExclude NamePatternRuleRuleTypeEnum = "EXCLUDE"
)

var mappingNamePatternRuleRuleTypeEnum = map[string]NamePatternRuleRuleTypeEnum{
	"INCLUDE": NamePatternRuleRuleTypeInclude,
	"EXCLUDE": NamePatternRuleRuleTypeExclude,
}

// GetNamePatternRuleRuleTypeEnumValues Enumerates the set of values for NamePatternRuleRuleTypeEnum
func GetNamePatternRuleRuleTypeEnumValues() []NamePatternRuleRuleTypeEnum {
	values := make([]NamePatternRuleRuleTypeEnum, 0)
	for _, v := range mappingNamePatternRuleRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNamePatternRuleRuleTypeEnumStringValues Enumerates the set of values in String for NamePatternRuleRuleTypeEnum
func GetNamePatternRuleRuleTypeEnumStringValues() []string {
	return []string{
		"INCLUDE",
		"EXCLUDE",
	}
}

// GetMappingNamePatternRuleRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNamePatternRuleRuleTypeEnum(val string) (NamePatternRuleRuleTypeEnum, bool) {
	mappingNamePatternRuleRuleTypeEnumIgnoreCase := make(map[string]NamePatternRuleRuleTypeEnum)
	for k, v := range mappingNamePatternRuleRuleTypeEnum {
		mappingNamePatternRuleRuleTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingNamePatternRuleRuleTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
