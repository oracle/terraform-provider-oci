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

// TypedNamePatternRule The typed name rule for field projection.
type TypedNamePatternRule struct {

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

	// types
	Types []BaseType `mandatory:"false" json:"types"`

	// skipRemainingRulesOnMatch
	IsSkipRemainingRulesOnMatch *bool `mandatory:"false" json:"isSkipRemainingRulesOnMatch"`

	// Reference to a typed object, this can be either a key value to an object within the document, a shall referenced to a TypedObject or a full TypedObject definition.
	Scope *interface{} `mandatory:"false" json:"scope"`

	// cascade
	IsCascade *bool `mandatory:"false" json:"isCascade"`

	// caseSensitive
	IsCaseSensitive *bool `mandatory:"false" json:"isCaseSensitive"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Names []string `mandatory:"false" json:"names"`

	// matchingStrategy
	MatchingStrategy TypedNamePatternRuleMatchingStrategyEnum `mandatory:"false" json:"matchingStrategy,omitempty"`

	// ruleType
	RuleType TypedNamePatternRuleRuleTypeEnum `mandatory:"false" json:"ruleType,omitempty"`
}

//GetKey returns Key
func (m TypedNamePatternRule) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m TypedNamePatternRule) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m TypedNamePatternRule) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetIsJavaRegexSyntax returns IsJavaRegexSyntax
func (m TypedNamePatternRule) GetIsJavaRegexSyntax() *bool {
	return m.IsJavaRegexSyntax
}

//GetConfigValues returns ConfigValues
func (m TypedNamePatternRule) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

//GetObjectStatus returns ObjectStatus
func (m TypedNamePatternRule) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetDescription returns Description
func (m TypedNamePatternRule) GetDescription() *string {
	return m.Description
}

func (m TypedNamePatternRule) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m TypedNamePatternRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTypedNamePatternRule TypedNamePatternRule
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeTypedNamePatternRule
	}{
		"TYPED_NAME_PATTERN_RULE",
		(MarshalTypeTypedNamePatternRule)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *TypedNamePatternRule) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key                         *string                                  `json:"key"`
		ModelVersion                *string                                  `json:"modelVersion"`
		ParentRef                   *ParentReference                         `json:"parentRef"`
		IsJavaRegexSyntax           *bool                                    `json:"isJavaRegexSyntax"`
		ConfigValues                *ConfigValues                            `json:"configValues"`
		ObjectStatus                *int                                     `json:"objectStatus"`
		Description                 *string                                  `json:"description"`
		Types                       []basetype                               `json:"types"`
		IsSkipRemainingRulesOnMatch *bool                                    `json:"isSkipRemainingRulesOnMatch"`
		Scope                       *interface{}                             `json:"scope"`
		IsCascade                   *bool                                    `json:"isCascade"`
		MatchingStrategy            TypedNamePatternRuleMatchingStrategyEnum `json:"matchingStrategy"`
		IsCaseSensitive             *bool                                    `json:"isCaseSensitive"`
		RuleType                    TypedNamePatternRuleRuleTypeEnum         `json:"ruleType"`
		Names                       []string                                 `json:"names"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.IsJavaRegexSyntax = model.IsJavaRegexSyntax

	m.ConfigValues = model.ConfigValues

	m.ObjectStatus = model.ObjectStatus

	m.Description = model.Description

	m.Types = make([]BaseType, len(model.Types))
	for i, n := range model.Types {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Types[i] = nn.(BaseType)
		} else {
			m.Types[i] = nil
		}
	}

	m.IsSkipRemainingRulesOnMatch = model.IsSkipRemainingRulesOnMatch

	m.Scope = model.Scope

	m.IsCascade = model.IsCascade

	m.MatchingStrategy = model.MatchingStrategy

	m.IsCaseSensitive = model.IsCaseSensitive

	m.RuleType = model.RuleType

	m.Names = make([]string, len(model.Names))
	for i, n := range model.Names {
		m.Names[i] = n
	}

	return
}

// TypedNamePatternRuleMatchingStrategyEnum Enum with underlying type: string
type TypedNamePatternRuleMatchingStrategyEnum string

// Set of constants representing the allowable values for TypedNamePatternRuleMatchingStrategyEnum
const (
	TypedNamePatternRuleMatchingStrategyNameOrTags TypedNamePatternRuleMatchingStrategyEnum = "NAME_OR_TAGS"
	TypedNamePatternRuleMatchingStrategyTagsOnly   TypedNamePatternRuleMatchingStrategyEnum = "TAGS_ONLY"
	TypedNamePatternRuleMatchingStrategyNameOnly   TypedNamePatternRuleMatchingStrategyEnum = "NAME_ONLY"
)

var mappingTypedNamePatternRuleMatchingStrategy = map[string]TypedNamePatternRuleMatchingStrategyEnum{
	"NAME_OR_TAGS": TypedNamePatternRuleMatchingStrategyNameOrTags,
	"TAGS_ONLY":    TypedNamePatternRuleMatchingStrategyTagsOnly,
	"NAME_ONLY":    TypedNamePatternRuleMatchingStrategyNameOnly,
}

// GetTypedNamePatternRuleMatchingStrategyEnumValues Enumerates the set of values for TypedNamePatternRuleMatchingStrategyEnum
func GetTypedNamePatternRuleMatchingStrategyEnumValues() []TypedNamePatternRuleMatchingStrategyEnum {
	values := make([]TypedNamePatternRuleMatchingStrategyEnum, 0)
	for _, v := range mappingTypedNamePatternRuleMatchingStrategy {
		values = append(values, v)
	}
	return values
}

// TypedNamePatternRuleRuleTypeEnum Enum with underlying type: string
type TypedNamePatternRuleRuleTypeEnum string

// Set of constants representing the allowable values for TypedNamePatternRuleRuleTypeEnum
const (
	TypedNamePatternRuleRuleTypeInclude TypedNamePatternRuleRuleTypeEnum = "INCLUDE"
	TypedNamePatternRuleRuleTypeExclude TypedNamePatternRuleRuleTypeEnum = "EXCLUDE"
)

var mappingTypedNamePatternRuleRuleType = map[string]TypedNamePatternRuleRuleTypeEnum{
	"INCLUDE": TypedNamePatternRuleRuleTypeInclude,
	"EXCLUDE": TypedNamePatternRuleRuleTypeExclude,
}

// GetTypedNamePatternRuleRuleTypeEnumValues Enumerates the set of values for TypedNamePatternRuleRuleTypeEnum
func GetTypedNamePatternRuleRuleTypeEnumValues() []TypedNamePatternRuleRuleTypeEnum {
	values := make([]TypedNamePatternRuleRuleTypeEnum, 0)
	for _, v := range mappingTypedNamePatternRuleRuleType {
		values = append(values, v)
	}
	return values
}
