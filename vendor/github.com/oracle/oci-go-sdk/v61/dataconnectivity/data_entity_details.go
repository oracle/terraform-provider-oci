// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v61/common"
	"strings"
)

// DataEntityDetails The data entity details object.
type DataEntityDetails interface {
}

type dataentitydetails struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *dataentitydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdataentitydetails dataentitydetails
	s := struct {
		Model Unmarshalerdataentitydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dataentitydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "FILE_ENTITY":
		mm := DataEntityFromFileEntityDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VIEW_ENTITY":
		mm := DataEntityFromViewEntityDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SQL_ENTITY":
		mm := DataEntityFromSqlEntityDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATA_STORE_ENTITY":
		mm := DataEntityFromDataStoreEntityDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TABLE_ENTITY":
		mm := DataEntityFromTableEntityDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m dataentitydetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dataentitydetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DataEntityDetailsModelTypeEnum Enum with underlying type: string
type DataEntityDetailsModelTypeEnum string

// Set of constants representing the allowable values for DataEntityDetailsModelTypeEnum
const (
	DataEntityDetailsModelTypeViewEntity      DataEntityDetailsModelTypeEnum = "VIEW_ENTITY"
	DataEntityDetailsModelTypeTableEntity     DataEntityDetailsModelTypeEnum = "TABLE_ENTITY"
	DataEntityDetailsModelTypeFileEntity      DataEntityDetailsModelTypeEnum = "FILE_ENTITY"
	DataEntityDetailsModelTypeDataStoreEntity DataEntityDetailsModelTypeEnum = "DATA_STORE_ENTITY"
	DataEntityDetailsModelTypeSqlEntity       DataEntityDetailsModelTypeEnum = "SQL_ENTITY"
)

var mappingDataEntityDetailsModelTypeEnum = map[string]DataEntityDetailsModelTypeEnum{
	"VIEW_ENTITY":       DataEntityDetailsModelTypeViewEntity,
	"TABLE_ENTITY":      DataEntityDetailsModelTypeTableEntity,
	"FILE_ENTITY":       DataEntityDetailsModelTypeFileEntity,
	"DATA_STORE_ENTITY": DataEntityDetailsModelTypeDataStoreEntity,
	"SQL_ENTITY":        DataEntityDetailsModelTypeSqlEntity,
}

var mappingDataEntityDetailsModelTypeEnumLowerCase = map[string]DataEntityDetailsModelTypeEnum{
	"view_entity":       DataEntityDetailsModelTypeViewEntity,
	"table_entity":      DataEntityDetailsModelTypeTableEntity,
	"file_entity":       DataEntityDetailsModelTypeFileEntity,
	"data_store_entity": DataEntityDetailsModelTypeDataStoreEntity,
	"sql_entity":        DataEntityDetailsModelTypeSqlEntity,
}

// GetDataEntityDetailsModelTypeEnumValues Enumerates the set of values for DataEntityDetailsModelTypeEnum
func GetDataEntityDetailsModelTypeEnumValues() []DataEntityDetailsModelTypeEnum {
	values := make([]DataEntityDetailsModelTypeEnum, 0)
	for _, v := range mappingDataEntityDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataEntityDetailsModelTypeEnumStringValues Enumerates the set of values in String for DataEntityDetailsModelTypeEnum
func GetDataEntityDetailsModelTypeEnumStringValues() []string {
	return []string{
		"VIEW_ENTITY",
		"TABLE_ENTITY",
		"FILE_ENTITY",
		"DATA_STORE_ENTITY",
		"SQL_ENTITY",
	}
}

// GetMappingDataEntityDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataEntityDetailsModelTypeEnum(val string) (DataEntityDetailsModelTypeEnum, bool) {
	enum, ok := mappingDataEntityDetailsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
