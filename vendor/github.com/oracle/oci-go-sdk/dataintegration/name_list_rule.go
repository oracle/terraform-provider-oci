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

// NameListRule The name list rule which defines how fields are projected. For example this may be all fields begining with STR.
type NameListRule struct {

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

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Names []string `mandatory:"false" json:"names"`

	// matchingStrategy
	MatchingStrategy NameListRuleMatchingStrategyEnum `mandatory:"false" json:"matchingStrategy,omitempty"`

	// ruleType
	RuleType NameListRuleRuleTypeEnum `mandatory:"false" json:"ruleType,omitempty"`
}

//GetKey returns Key
func (m NameListRule) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m NameListRule) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m NameListRule) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetIsJavaRegexSyntax returns IsJavaRegexSyntax
func (m NameListRule) GetIsJavaRegexSyntax() *bool {
	return m.IsJavaRegexSyntax
}

//GetConfigValues returns ConfigValues
func (m NameListRule) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

//GetObjectStatus returns ObjectStatus
func (m NameListRule) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetDescription returns Description
func (m NameListRule) GetDescription() *string {
	return m.Description
}

func (m NameListRule) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m NameListRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNameListRule NameListRule
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeNameListRule
	}{
		"NAME_LIST_RULE",
		(MarshalTypeNameListRule)(m),
	}

	return json.Marshal(&s)
}

// NameListRuleMatchingStrategyEnum Enum with underlying type: string
type NameListRuleMatchingStrategyEnum string

// Set of constants representing the allowable values for NameListRuleMatchingStrategyEnum
const (
	NameListRuleMatchingStrategyNameOrTags NameListRuleMatchingStrategyEnum = "NAME_OR_TAGS"
	NameListRuleMatchingStrategyTagsOnly   NameListRuleMatchingStrategyEnum = "TAGS_ONLY"
	NameListRuleMatchingStrategyNameOnly   NameListRuleMatchingStrategyEnum = "NAME_ONLY"
)

var mappingNameListRuleMatchingStrategy = map[string]NameListRuleMatchingStrategyEnum{
	"NAME_OR_TAGS": NameListRuleMatchingStrategyNameOrTags,
	"TAGS_ONLY":    NameListRuleMatchingStrategyTagsOnly,
	"NAME_ONLY":    NameListRuleMatchingStrategyNameOnly,
}

// GetNameListRuleMatchingStrategyEnumValues Enumerates the set of values for NameListRuleMatchingStrategyEnum
func GetNameListRuleMatchingStrategyEnumValues() []NameListRuleMatchingStrategyEnum {
	values := make([]NameListRuleMatchingStrategyEnum, 0)
	for _, v := range mappingNameListRuleMatchingStrategy {
		values = append(values, v)
	}
	return values
}

// NameListRuleRuleTypeEnum Enum with underlying type: string
type NameListRuleRuleTypeEnum string

// Set of constants representing the allowable values for NameListRuleRuleTypeEnum
const (
	NameListRuleRuleTypeInclude NameListRuleRuleTypeEnum = "INCLUDE"
	NameListRuleRuleTypeExclude NameListRuleRuleTypeEnum = "EXCLUDE"
)

var mappingNameListRuleRuleType = map[string]NameListRuleRuleTypeEnum{
	"INCLUDE": NameListRuleRuleTypeInclude,
	"EXCLUDE": NameListRuleRuleTypeExclude,
}

// GetNameListRuleRuleTypeEnumValues Enumerates the set of values for NameListRuleRuleTypeEnum
func GetNameListRuleRuleTypeEnumValues() []NameListRuleRuleTypeEnum {
	values := make([]NameListRuleRuleTypeEnum, 0)
	for _, v := range mappingNameListRuleRuleType {
		values = append(values, v)
	}
	return values
}
