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

// DynamicTypeHandler This type defines how to derived fields for the dynamic type itself.
type DynamicTypeHandler interface {
}

type dynamictypehandler struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *dynamictypehandler) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdynamictypehandler dynamictypehandler
	s := struct {
		Model Unmarshalerdynamictypehandler
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dynamictypehandler) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "FLATTEN_TYPE_HANDLER":
		mm := FlattenTypeHandler{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RULE_TYPE_CONFIGS":
		mm := RuleTypeConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DynamicTypeHandler: %s.", m.ModelType)
		return *m, nil
	}
}

func (m dynamictypehandler) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dynamictypehandler) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DynamicTypeHandlerModelTypeEnum Enum with underlying type: string
type DynamicTypeHandlerModelTypeEnum string

// Set of constants representing the allowable values for DynamicTypeHandlerModelTypeEnum
const (
	DynamicTypeHandlerModelTypeRuleTypeConfigs    DynamicTypeHandlerModelTypeEnum = "RULE_TYPE_CONFIGS"
	DynamicTypeHandlerModelTypeFlattenTypeHandler DynamicTypeHandlerModelTypeEnum = "FLATTEN_TYPE_HANDLER"
)

var mappingDynamicTypeHandlerModelTypeEnum = map[string]DynamicTypeHandlerModelTypeEnum{
	"RULE_TYPE_CONFIGS":    DynamicTypeHandlerModelTypeRuleTypeConfigs,
	"FLATTEN_TYPE_HANDLER": DynamicTypeHandlerModelTypeFlattenTypeHandler,
}

var mappingDynamicTypeHandlerModelTypeEnumLowerCase = map[string]DynamicTypeHandlerModelTypeEnum{
	"rule_type_configs":    DynamicTypeHandlerModelTypeRuleTypeConfigs,
	"flatten_type_handler": DynamicTypeHandlerModelTypeFlattenTypeHandler,
}

// GetDynamicTypeHandlerModelTypeEnumValues Enumerates the set of values for DynamicTypeHandlerModelTypeEnum
func GetDynamicTypeHandlerModelTypeEnumValues() []DynamicTypeHandlerModelTypeEnum {
	values := make([]DynamicTypeHandlerModelTypeEnum, 0)
	for _, v := range mappingDynamicTypeHandlerModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDynamicTypeHandlerModelTypeEnumStringValues Enumerates the set of values in String for DynamicTypeHandlerModelTypeEnum
func GetDynamicTypeHandlerModelTypeEnumStringValues() []string {
	return []string{
		"RULE_TYPE_CONFIGS",
		"FLATTEN_TYPE_HANDLER",
	}
}

// GetMappingDynamicTypeHandlerModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDynamicTypeHandlerModelTypeEnum(val string) (DynamicTypeHandlerModelTypeEnum, bool) {
	enum, ok := mappingDynamicTypeHandlerModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
