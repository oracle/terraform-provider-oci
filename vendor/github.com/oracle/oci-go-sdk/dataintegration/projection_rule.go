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

// ProjectionRule Base type for how fields are projected, there are many different mechanisms for doing this such as by a name patter, datatype etc. See the modelType property for the types.
type ProjectionRule interface {

	// The key of the object.
	GetKey() *string

	// The model version of an object.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// Specifies whether the rule uses a java regex syntax.
	GetIsJavaRegexSyntax() *bool

	GetConfigValues() *ConfigValues

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Detailed description for the object.
	GetDescription() *string
}

type projectionrule struct {
	JsonData          []byte
	Key               *string          `mandatory:"false" json:"key"`
	ModelVersion      *string          `mandatory:"false" json:"modelVersion"`
	ParentRef         *ParentReference `mandatory:"false" json:"parentRef"`
	IsJavaRegexSyntax *bool            `mandatory:"false" json:"isJavaRegexSyntax"`
	ConfigValues      *ConfigValues    `mandatory:"false" json:"configValues"`
	ObjectStatus      *int             `mandatory:"false" json:"objectStatus"`
	Description       *string          `mandatory:"false" json:"description"`
	ModelType         string           `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *projectionrule) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerprojectionrule projectionrule
	s := struct {
		Model Unmarshalerprojectionrule
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.ModelVersion = s.Model.ModelVersion
	m.ParentRef = s.Model.ParentRef
	m.IsJavaRegexSyntax = s.Model.IsJavaRegexSyntax
	m.ConfigValues = s.Model.ConfigValues
	m.ObjectStatus = s.Model.ObjectStatus
	m.Description = s.Model.Description
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *projectionrule) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "RENAME_RULE":
		mm := RenameRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TYPE_LIST_RULE":
		mm := TypeListRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TYPED_NAME_PATTERN_RULE":
		mm := TypedNamePatternRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NAME_PATTERN_RULE":
		mm := NamePatternRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NAME_LIST_RULE":
		mm := NameListRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetKey returns Key
func (m projectionrule) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m projectionrule) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m projectionrule) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetIsJavaRegexSyntax returns IsJavaRegexSyntax
func (m projectionrule) GetIsJavaRegexSyntax() *bool {
	return m.IsJavaRegexSyntax
}

//GetConfigValues returns ConfigValues
func (m projectionrule) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

//GetObjectStatus returns ObjectStatus
func (m projectionrule) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetDescription returns Description
func (m projectionrule) GetDescription() *string {
	return m.Description
}

func (m projectionrule) String() string {
	return common.PointerString(m)
}

// ProjectionRuleModelTypeEnum Enum with underlying type: string
type ProjectionRuleModelTypeEnum string

// Set of constants representing the allowable values for ProjectionRuleModelTypeEnum
const (
	ProjectionRuleModelTypeNamePatternRule      ProjectionRuleModelTypeEnum = "NAME_PATTERN_RULE"
	ProjectionRuleModelTypeTypeListRule         ProjectionRuleModelTypeEnum = "TYPE_LIST_RULE"
	ProjectionRuleModelTypeNameListRule         ProjectionRuleModelTypeEnum = "NAME_LIST_RULE"
	ProjectionRuleModelTypeTypedNamePatternRule ProjectionRuleModelTypeEnum = "TYPED_NAME_PATTERN_RULE"
	ProjectionRuleModelTypeRenameRule           ProjectionRuleModelTypeEnum = "RENAME_RULE"
)

var mappingProjectionRuleModelType = map[string]ProjectionRuleModelTypeEnum{
	"NAME_PATTERN_RULE":       ProjectionRuleModelTypeNamePatternRule,
	"TYPE_LIST_RULE":          ProjectionRuleModelTypeTypeListRule,
	"NAME_LIST_RULE":          ProjectionRuleModelTypeNameListRule,
	"TYPED_NAME_PATTERN_RULE": ProjectionRuleModelTypeTypedNamePatternRule,
	"RENAME_RULE":             ProjectionRuleModelTypeRenameRule,
}

// GetProjectionRuleModelTypeEnumValues Enumerates the set of values for ProjectionRuleModelTypeEnum
func GetProjectionRuleModelTypeEnumValues() []ProjectionRuleModelTypeEnum {
	values := make([]ProjectionRuleModelTypeEnum, 0)
	for _, v := range mappingProjectionRuleModelType {
		values = append(values, v)
	}
	return values
}
