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

	// A user defined description for the object.
	Description *string `mandatory:"false" json:"description"`

	// An array of types.
	Types []interface{} `mandatory:"false" json:"types"`

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

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Names []string `mandatory:"false" json:"names"`

	// The pattern matching strategy.
	MatchingStrategy TypedNamePatternRuleMatchingStrategyEnum `mandatory:"false" json:"matchingStrategy,omitempty"`

	// The rule type.
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TypedNamePatternRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTypedNamePatternRuleMatchingStrategyEnum(string(m.MatchingStrategy)); !ok && m.MatchingStrategy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MatchingStrategy: %s. Supported values are: %s.", m.MatchingStrategy, strings.Join(GetTypedNamePatternRuleMatchingStrategyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTypedNamePatternRuleRuleTypeEnum(string(m.RuleType)); !ok && m.RuleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuleType: %s. Supported values are: %s.", m.RuleType, strings.Join(GetTypedNamePatternRuleRuleTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

// TypedNamePatternRuleMatchingStrategyEnum Enum with underlying type: string
type TypedNamePatternRuleMatchingStrategyEnum string

// Set of constants representing the allowable values for TypedNamePatternRuleMatchingStrategyEnum
const (
	TypedNamePatternRuleMatchingStrategyNameOrTags TypedNamePatternRuleMatchingStrategyEnum = "NAME_OR_TAGS"
	TypedNamePatternRuleMatchingStrategyTagsOnly   TypedNamePatternRuleMatchingStrategyEnum = "TAGS_ONLY"
	TypedNamePatternRuleMatchingStrategyNameOnly   TypedNamePatternRuleMatchingStrategyEnum = "NAME_ONLY"
)

var mappingTypedNamePatternRuleMatchingStrategyEnum = map[string]TypedNamePatternRuleMatchingStrategyEnum{
	"NAME_OR_TAGS": TypedNamePatternRuleMatchingStrategyNameOrTags,
	"TAGS_ONLY":    TypedNamePatternRuleMatchingStrategyTagsOnly,
	"NAME_ONLY":    TypedNamePatternRuleMatchingStrategyNameOnly,
}

// GetTypedNamePatternRuleMatchingStrategyEnumValues Enumerates the set of values for TypedNamePatternRuleMatchingStrategyEnum
func GetTypedNamePatternRuleMatchingStrategyEnumValues() []TypedNamePatternRuleMatchingStrategyEnum {
	values := make([]TypedNamePatternRuleMatchingStrategyEnum, 0)
	for _, v := range mappingTypedNamePatternRuleMatchingStrategyEnum {
		values = append(values, v)
	}
	return values
}

// GetTypedNamePatternRuleMatchingStrategyEnumStringValues Enumerates the set of values in String for TypedNamePatternRuleMatchingStrategyEnum
func GetTypedNamePatternRuleMatchingStrategyEnumStringValues() []string {
	return []string{
		"NAME_OR_TAGS",
		"TAGS_ONLY",
		"NAME_ONLY",
	}
}

// GetMappingTypedNamePatternRuleMatchingStrategyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTypedNamePatternRuleMatchingStrategyEnum(val string) (TypedNamePatternRuleMatchingStrategyEnum, bool) {
	mappingTypedNamePatternRuleMatchingStrategyEnumIgnoreCase := make(map[string]TypedNamePatternRuleMatchingStrategyEnum)
	for k, v := range mappingTypedNamePatternRuleMatchingStrategyEnum {
		mappingTypedNamePatternRuleMatchingStrategyEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingTypedNamePatternRuleMatchingStrategyEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// TypedNamePatternRuleRuleTypeEnum Enum with underlying type: string
type TypedNamePatternRuleRuleTypeEnum string

// Set of constants representing the allowable values for TypedNamePatternRuleRuleTypeEnum
const (
	TypedNamePatternRuleRuleTypeInclude TypedNamePatternRuleRuleTypeEnum = "INCLUDE"
	TypedNamePatternRuleRuleTypeExclude TypedNamePatternRuleRuleTypeEnum = "EXCLUDE"
)

var mappingTypedNamePatternRuleRuleTypeEnum = map[string]TypedNamePatternRuleRuleTypeEnum{
	"INCLUDE": TypedNamePatternRuleRuleTypeInclude,
	"EXCLUDE": TypedNamePatternRuleRuleTypeExclude,
}

// GetTypedNamePatternRuleRuleTypeEnumValues Enumerates the set of values for TypedNamePatternRuleRuleTypeEnum
func GetTypedNamePatternRuleRuleTypeEnumValues() []TypedNamePatternRuleRuleTypeEnum {
	values := make([]TypedNamePatternRuleRuleTypeEnum, 0)
	for _, v := range mappingTypedNamePatternRuleRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTypedNamePatternRuleRuleTypeEnumStringValues Enumerates the set of values in String for TypedNamePatternRuleRuleTypeEnum
func GetTypedNamePatternRuleRuleTypeEnumStringValues() []string {
	return []string{
		"INCLUDE",
		"EXCLUDE",
	}
}

// GetMappingTypedNamePatternRuleRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTypedNamePatternRuleRuleTypeEnum(val string) (TypedNamePatternRuleRuleTypeEnum, bool) {
	mappingTypedNamePatternRuleRuleTypeEnumIgnoreCase := make(map[string]TypedNamePatternRuleRuleTypeEnum)
	for k, v := range mappingTypedNamePatternRuleRuleTypeEnum {
		mappingTypedNamePatternRuleRuleTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingTypedNamePatternRuleRuleTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
