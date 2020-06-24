// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// NamePatternRule This rule projects fields by a name pattern, for example it may start with STR_ or end with _DATE, this is defined using a regular expression.
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

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// skipRemainingRulesOnMatch
	IsSkipRemainingRulesOnMatch *bool `mandatory:"false" json:"isSkipRemainingRulesOnMatch"`

	// Reference to a typed object, this can be either a key value to an object within the document, a shall referenced to a TypedObject or a full TypedObject definition.
	Scope *interface{} `mandatory:"false" json:"scope"`

	// cascade
	IsCascade *bool `mandatory:"false" json:"isCascade"`

	// caseSensitive
	IsCaseSensitive *bool `mandatory:"false" json:"isCaseSensitive"`

	// pattern
	Pattern *string `mandatory:"false" json:"pattern"`

	// matchingStrategy
	MatchingStrategy NamePatternRuleMatchingStrategyEnum `mandatory:"false" json:"matchingStrategy,omitempty"`

	// ruleType
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

var mappingNamePatternRuleMatchingStrategy = map[string]NamePatternRuleMatchingStrategyEnum{
	"NAME_OR_TAGS": NamePatternRuleMatchingStrategyNameOrTags,
	"TAGS_ONLY":    NamePatternRuleMatchingStrategyTagsOnly,
	"NAME_ONLY":    NamePatternRuleMatchingStrategyNameOnly,
}

// GetNamePatternRuleMatchingStrategyEnumValues Enumerates the set of values for NamePatternRuleMatchingStrategyEnum
func GetNamePatternRuleMatchingStrategyEnumValues() []NamePatternRuleMatchingStrategyEnum {
	values := make([]NamePatternRuleMatchingStrategyEnum, 0)
	for _, v := range mappingNamePatternRuleMatchingStrategy {
		values = append(values, v)
	}
	return values
}

// NamePatternRuleRuleTypeEnum Enum with underlying type: string
type NamePatternRuleRuleTypeEnum string

// Set of constants representing the allowable values for NamePatternRuleRuleTypeEnum
const (
	NamePatternRuleRuleTypeInclude NamePatternRuleRuleTypeEnum = "INCLUDE"
	NamePatternRuleRuleTypeExclude NamePatternRuleRuleTypeEnum = "EXCLUDE"
)

var mappingNamePatternRuleRuleType = map[string]NamePatternRuleRuleTypeEnum{
	"INCLUDE": NamePatternRuleRuleTypeInclude,
	"EXCLUDE": NamePatternRuleRuleTypeExclude,
}

// GetNamePatternRuleRuleTypeEnumValues Enumerates the set of values for NamePatternRuleRuleTypeEnum
func GetNamePatternRuleRuleTypeEnumValues() []NamePatternRuleRuleTypeEnum {
	values := make([]NamePatternRuleRuleTypeEnum, 0)
	for _, v := range mappingNamePatternRuleRuleType {
		values = append(values, v)
	}
	return values
}
