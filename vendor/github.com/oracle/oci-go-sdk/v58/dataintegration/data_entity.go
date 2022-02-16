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

// DataEntity The data entity object.
type DataEntity interface {
	GetMetadata() *ObjectMetadata
}

type dataentity struct {
	JsonData  []byte
	Metadata  *ObjectMetadata `mandatory:"false" json:"metadata"`
	ModelType string          `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *dataentity) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdataentity dataentity
	s := struct {
		Model Unmarshalerdataentity
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Metadata = s.Model.Metadata
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dataentity) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "TABLE_ENTITY":
		mm := DataEntityFromTable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATA_STORE_ENTITY":
		mm := DataEntityFromDataStore{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VIEW_ENTITY":
		mm := DataEntityFromView{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SQL_ENTITY":
		mm := DataEntityFromSql{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FILE_ENTITY":
		mm := DataEntityFromFile{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetMetadata returns Metadata
func (m dataentity) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m dataentity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dataentity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DataEntityModelTypeEnum Enum with underlying type: string
type DataEntityModelTypeEnum string

// Set of constants representing the allowable values for DataEntityModelTypeEnum
const (
	DataEntityModelTypeViewEntity      DataEntityModelTypeEnum = "VIEW_ENTITY"
	DataEntityModelTypeTableEntity     DataEntityModelTypeEnum = "TABLE_ENTITY"
	DataEntityModelTypeFileEntity      DataEntityModelTypeEnum = "FILE_ENTITY"
	DataEntityModelTypeSqlEntity       DataEntityModelTypeEnum = "SQL_ENTITY"
	DataEntityModelTypeDataStoreEntity DataEntityModelTypeEnum = "DATA_STORE_ENTITY"
)

var mappingDataEntityModelTypeEnum = map[string]DataEntityModelTypeEnum{
	"VIEW_ENTITY":       DataEntityModelTypeViewEntity,
	"TABLE_ENTITY":      DataEntityModelTypeTableEntity,
	"FILE_ENTITY":       DataEntityModelTypeFileEntity,
	"SQL_ENTITY":        DataEntityModelTypeSqlEntity,
	"DATA_STORE_ENTITY": DataEntityModelTypeDataStoreEntity,
}

// GetDataEntityModelTypeEnumValues Enumerates the set of values for DataEntityModelTypeEnum
func GetDataEntityModelTypeEnumValues() []DataEntityModelTypeEnum {
	values := make([]DataEntityModelTypeEnum, 0)
	for _, v := range mappingDataEntityModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataEntityModelTypeEnumStringValues Enumerates the set of values in String for DataEntityModelTypeEnum
func GetDataEntityModelTypeEnumStringValues() []string {
	return []string{
		"VIEW_ENTITY",
		"TABLE_ENTITY",
		"FILE_ENTITY",
		"SQL_ENTITY",
		"DATA_STORE_ENTITY",
	}
}

// GetMappingDataEntityModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataEntityModelTypeEnum(val string) (DataEntityModelTypeEnum, bool) {
	mappingDataEntityModelTypeEnumIgnoreCase := make(map[string]DataEntityModelTypeEnum)
	for k, v := range mappingDataEntityModelTypeEnum {
		mappingDataEntityModelTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDataEntityModelTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
