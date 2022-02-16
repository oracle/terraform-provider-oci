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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m abstractdataoperationconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AbstractDataOperationConfigModelTypeEnum Enum with underlying type: string
type AbstractDataOperationConfigModelTypeEnum string

// Set of constants representing the allowable values for AbstractDataOperationConfigModelTypeEnum
const (
	AbstractDataOperationConfigModelTypeReadOperationConfig  AbstractDataOperationConfigModelTypeEnum = "READ_OPERATION_CONFIG"
	AbstractDataOperationConfigModelTypeWriteOperationConfig AbstractDataOperationConfigModelTypeEnum = "WRITE_OPERATION_CONFIG"
)

var mappingAbstractDataOperationConfigModelTypeEnum = map[string]AbstractDataOperationConfigModelTypeEnum{
	"READ_OPERATION_CONFIG":  AbstractDataOperationConfigModelTypeReadOperationConfig,
	"WRITE_OPERATION_CONFIG": AbstractDataOperationConfigModelTypeWriteOperationConfig,
}

// GetAbstractDataOperationConfigModelTypeEnumValues Enumerates the set of values for AbstractDataOperationConfigModelTypeEnum
func GetAbstractDataOperationConfigModelTypeEnumValues() []AbstractDataOperationConfigModelTypeEnum {
	values := make([]AbstractDataOperationConfigModelTypeEnum, 0)
	for _, v := range mappingAbstractDataOperationConfigModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAbstractDataOperationConfigModelTypeEnumStringValues Enumerates the set of values in String for AbstractDataOperationConfigModelTypeEnum
func GetAbstractDataOperationConfigModelTypeEnumStringValues() []string {
	return []string{
		"READ_OPERATION_CONFIG",
		"WRITE_OPERATION_CONFIG",
	}
}

// GetMappingAbstractDataOperationConfigModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAbstractDataOperationConfigModelTypeEnum(val string) (AbstractDataOperationConfigModelTypeEnum, bool) {
	mappingAbstractDataOperationConfigModelTypeEnumIgnoreCase := make(map[string]AbstractDataOperationConfigModelTypeEnum)
	for k, v := range mappingAbstractDataOperationConfigModelTypeEnum {
		mappingAbstractDataOperationConfigModelTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingAbstractDataOperationConfigModelTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
