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
	"github.com/oracle/oci-go-sdk/v64/common"
	"strings"
)

// EntityShape The data entity shape object.
type EntityShape interface {
	GetMetadata() *ObjectMetadata
}

type entityshape struct {
	JsonData  []byte
	Metadata  *ObjectMetadata `mandatory:"false" json:"metadata"`
	ModelType string          `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *entityshape) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerentityshape entityshape
	s := struct {
		Model Unmarshalerentityshape
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
func (m *entityshape) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "DATA_STORE_ENTITY":
		mm := EntityShapeFromDataStore{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TABLE_ENTITY":
		mm := EntityShapeFromTable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SQL_ENTITY":
		mm := EntityShapeFromSql{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VIEW_ENTITY":
		mm := EntityShapeFromView{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FILE_ENTITY":
		mm := EntityShapeFromFile{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetMetadata returns Metadata
func (m entityshape) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m entityshape) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m entityshape) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EntityShapeModelTypeEnum Enum with underlying type: string
type EntityShapeModelTypeEnum string

// Set of constants representing the allowable values for EntityShapeModelTypeEnum
const (
	EntityShapeModelTypeViewEntity      EntityShapeModelTypeEnum = "VIEW_ENTITY"
	EntityShapeModelTypeTableEntity     EntityShapeModelTypeEnum = "TABLE_ENTITY"
	EntityShapeModelTypeFileEntity      EntityShapeModelTypeEnum = "FILE_ENTITY"
	EntityShapeModelTypeSqlEntity       EntityShapeModelTypeEnum = "SQL_ENTITY"
	EntityShapeModelTypeDataStoreEntity EntityShapeModelTypeEnum = "DATA_STORE_ENTITY"
)

var mappingEntityShapeModelTypeEnum = map[string]EntityShapeModelTypeEnum{
	"VIEW_ENTITY":       EntityShapeModelTypeViewEntity,
	"TABLE_ENTITY":      EntityShapeModelTypeTableEntity,
	"FILE_ENTITY":       EntityShapeModelTypeFileEntity,
	"SQL_ENTITY":        EntityShapeModelTypeSqlEntity,
	"DATA_STORE_ENTITY": EntityShapeModelTypeDataStoreEntity,
}

var mappingEntityShapeModelTypeEnumLowerCase = map[string]EntityShapeModelTypeEnum{
	"view_entity":       EntityShapeModelTypeViewEntity,
	"table_entity":      EntityShapeModelTypeTableEntity,
	"file_entity":       EntityShapeModelTypeFileEntity,
	"sql_entity":        EntityShapeModelTypeSqlEntity,
	"data_store_entity": EntityShapeModelTypeDataStoreEntity,
}

// GetEntityShapeModelTypeEnumValues Enumerates the set of values for EntityShapeModelTypeEnum
func GetEntityShapeModelTypeEnumValues() []EntityShapeModelTypeEnum {
	values := make([]EntityShapeModelTypeEnum, 0)
	for _, v := range mappingEntityShapeModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEntityShapeModelTypeEnumStringValues Enumerates the set of values in String for EntityShapeModelTypeEnum
func GetEntityShapeModelTypeEnumStringValues() []string {
	return []string{
		"VIEW_ENTITY",
		"TABLE_ENTITY",
		"FILE_ENTITY",
		"SQL_ENTITY",
		"DATA_STORE_ENTITY",
	}
}

// GetMappingEntityShapeModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEntityShapeModelTypeEnum(val string) (EntityShapeModelTypeEnum, bool) {
	enum, ok := mappingEntityShapeModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
