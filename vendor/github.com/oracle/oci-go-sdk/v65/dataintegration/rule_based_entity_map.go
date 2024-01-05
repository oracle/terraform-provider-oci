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

// RuleBasedEntityMap A map of rule patterns.
type RuleBasedEntityMap struct {

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The object key.
	Key *string `mandatory:"false" json:"key"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// The pattern to map from.
	FromPattern *string `mandatory:"false" json:"fromPattern"`

	// The pattern to map to.
	ToPattern *string `mandatory:"false" json:"toPattern"`

	// Specifies whether the rule uses a java regex syntax.
	IsJavaRegexSyntax *bool `mandatory:"false" json:"isJavaRegexSyntax"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// mapType
	MapType RuleBasedEntityMapMapTypeEnum `mandatory:"false" json:"mapType,omitempty"`
}

// GetDescription returns Description
func (m RuleBasedEntityMap) GetDescription() *string {
	return m.Description
}

func (m RuleBasedEntityMap) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RuleBasedEntityMap) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRuleBasedEntityMapMapTypeEnum(string(m.MapType)); !ok && m.MapType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MapType: %s. Supported values are: %s.", m.MapType, strings.Join(GetRuleBasedEntityMapMapTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RuleBasedEntityMap) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRuleBasedEntityMap RuleBasedEntityMap
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeRuleBasedEntityMap
	}{
		"RULE_BASED_ENTITY_MAP",
		(MarshalTypeRuleBasedEntityMap)(m),
	}

	return json.Marshal(&s)
}

// RuleBasedEntityMapMapTypeEnum Enum with underlying type: string
type RuleBasedEntityMapMapTypeEnum string

// Set of constants representing the allowable values for RuleBasedEntityMapMapTypeEnum
const (
	RuleBasedEntityMapMapTypeMapbyname    RuleBasedEntityMapMapTypeEnum = "MAPBYNAME"
	RuleBasedEntityMapMapTypeMapbypattern RuleBasedEntityMapMapTypeEnum = "MAPBYPATTERN"
)

var mappingRuleBasedEntityMapMapTypeEnum = map[string]RuleBasedEntityMapMapTypeEnum{
	"MAPBYNAME":    RuleBasedEntityMapMapTypeMapbyname,
	"MAPBYPATTERN": RuleBasedEntityMapMapTypeMapbypattern,
}

var mappingRuleBasedEntityMapMapTypeEnumLowerCase = map[string]RuleBasedEntityMapMapTypeEnum{
	"mapbyname":    RuleBasedEntityMapMapTypeMapbyname,
	"mapbypattern": RuleBasedEntityMapMapTypeMapbypattern,
}

// GetRuleBasedEntityMapMapTypeEnumValues Enumerates the set of values for RuleBasedEntityMapMapTypeEnum
func GetRuleBasedEntityMapMapTypeEnumValues() []RuleBasedEntityMapMapTypeEnum {
	values := make([]RuleBasedEntityMapMapTypeEnum, 0)
	for _, v := range mappingRuleBasedEntityMapMapTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRuleBasedEntityMapMapTypeEnumStringValues Enumerates the set of values in String for RuleBasedEntityMapMapTypeEnum
func GetRuleBasedEntityMapMapTypeEnumStringValues() []string {
	return []string{
		"MAPBYNAME",
		"MAPBYPATTERN",
	}
}

// GetMappingRuleBasedEntityMapMapTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuleBasedEntityMapMapTypeEnum(val string) (RuleBasedEntityMapMapTypeEnum, bool) {
	enum, ok := mappingRuleBasedEntityMapMapTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
