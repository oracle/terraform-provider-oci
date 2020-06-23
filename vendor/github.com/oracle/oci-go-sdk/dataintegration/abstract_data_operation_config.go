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

// AbstractDataOperationConfig The information about the data operation.
type AbstractDataOperationConfig interface {
}

type abstractdataoperationconfig struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *abstractdataoperationconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerabstractdataoperationconfig abstractdataoperationconfig
	s := struct {
		Model Unmarshalerabstractdataoperationconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *abstractdataoperationconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "WRITE_OPERATION_CONFIG":
		mm := WriteOperationConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "READ_OPERATION_CONFIG":
		mm := ReadOperationConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m abstractdataoperationconfig) String() string {
	return common.PointerString(m)
}

// AbstractDataOperationConfigModelTypeEnum Enum with underlying type: string
type AbstractDataOperationConfigModelTypeEnum string

// Set of constants representing the allowable values for AbstractDataOperationConfigModelTypeEnum
const (
	AbstractDataOperationConfigModelTypeReadOperationConfig  AbstractDataOperationConfigModelTypeEnum = "READ_OPERATION_CONFIG"
	AbstractDataOperationConfigModelTypeWriteOperationConfig AbstractDataOperationConfigModelTypeEnum = "WRITE_OPERATION_CONFIG"
)

var mappingAbstractDataOperationConfigModelType = map[string]AbstractDataOperationConfigModelTypeEnum{
	"READ_OPERATION_CONFIG":  AbstractDataOperationConfigModelTypeReadOperationConfig,
	"WRITE_OPERATION_CONFIG": AbstractDataOperationConfigModelTypeWriteOperationConfig,
}

// GetAbstractDataOperationConfigModelTypeEnumValues Enumerates the set of values for AbstractDataOperationConfigModelTypeEnum
func GetAbstractDataOperationConfigModelTypeEnumValues() []AbstractDataOperationConfigModelTypeEnum {
	values := make([]AbstractDataOperationConfigModelTypeEnum, 0)
	for _, v := range mappingAbstractDataOperationConfigModelType {
		values = append(values, v)
	}
	return values
}
