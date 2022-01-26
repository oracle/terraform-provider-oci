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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// NameListRule The name list rule which defines how fields are projected. For example, this may be all fields begining with STR.
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

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Names []string `mandatory:"false" json:"names"`

	// The pattern matching strategy.
	MatchingStrategy NameListRuleMatchingStrategyEnum `mandatory:"false" json:"matchingStrategy,omitempty"`

	// The rule type.
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
