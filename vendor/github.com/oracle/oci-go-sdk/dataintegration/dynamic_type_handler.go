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
	case "RULE_TYPE_CONFIGS":
		mm := RuleTypeConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m dynamictypehandler) String() string {
	return common.PointerString(m)
}

// DynamicTypeHandlerModelTypeEnum Enum with underlying type: string
type DynamicTypeHandlerModelTypeEnum string

// Set of constants representing the allowable values for DynamicTypeHandlerModelTypeEnum
const (
	DynamicTypeHandlerModelTypeRuleTypeConfigs DynamicTypeHandlerModelTypeEnum = "RULE_TYPE_CONFIGS"
)

var mappingDynamicTypeHandlerModelType = map[string]DynamicTypeHandlerModelTypeEnum{
	"RULE_TYPE_CONFIGS": DynamicTypeHandlerModelTypeRuleTypeConfigs,
}

// GetDynamicTypeHandlerModelTypeEnumValues Enumerates the set of values for DynamicTypeHandlerModelTypeEnum
func GetDynamicTypeHandlerModelTypeEnumValues() []DynamicTypeHandlerModelTypeEnum {
	values := make([]DynamicTypeHandlerModelTypeEnum, 0)
	for _, v := range mappingDynamicTypeHandlerModelType {
		values = append(values, v)
	}
	return values
}
