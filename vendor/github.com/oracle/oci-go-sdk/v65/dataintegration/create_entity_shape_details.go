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

// CreateEntityShapeDetails The data entity shape object.
type CreateEntityShapeDetails interface {
}

type createentityshapedetails struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *createentityshapedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateentityshapedetails createentityshapedetails
	s := struct {
		Model Unmarshalercreateentityshapedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createentityshapedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "OBJECT_ENTITY":
		mm := CreateEntityShapeFromObject{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SQL_ENTITY":
		mm := CreateEntityShapeFromSql{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FILE_ENTITY":
		mm := CreateEntityShapeFromFile{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateEntityShapeDetails: %s.", m.ModelType)
		return *m, nil
	}
}

func (m createentityshapedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createentityshapedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateEntityShapeDetailsModelTypeEnum Enum with underlying type: string
type CreateEntityShapeDetailsModelTypeEnum string

// Set of constants representing the allowable values for CreateEntityShapeDetailsModelTypeEnum
const (
	CreateEntityShapeDetailsModelTypeFileEntity   CreateEntityShapeDetailsModelTypeEnum = "FILE_ENTITY"
	CreateEntityShapeDetailsModelTypeSqlEntity    CreateEntityShapeDetailsModelTypeEnum = "SQL_ENTITY"
	CreateEntityShapeDetailsModelTypeObjectEntity CreateEntityShapeDetailsModelTypeEnum = "OBJECT_ENTITY"
)

var mappingCreateEntityShapeDetailsModelTypeEnum = map[string]CreateEntityShapeDetailsModelTypeEnum{
	"FILE_ENTITY":   CreateEntityShapeDetailsModelTypeFileEntity,
	"SQL_ENTITY":    CreateEntityShapeDetailsModelTypeSqlEntity,
	"OBJECT_ENTITY": CreateEntityShapeDetailsModelTypeObjectEntity,
}

var mappingCreateEntityShapeDetailsModelTypeEnumLowerCase = map[string]CreateEntityShapeDetailsModelTypeEnum{
	"file_entity":   CreateEntityShapeDetailsModelTypeFileEntity,
	"sql_entity":    CreateEntityShapeDetailsModelTypeSqlEntity,
	"object_entity": CreateEntityShapeDetailsModelTypeObjectEntity,
}

// GetCreateEntityShapeDetailsModelTypeEnumValues Enumerates the set of values for CreateEntityShapeDetailsModelTypeEnum
func GetCreateEntityShapeDetailsModelTypeEnumValues() []CreateEntityShapeDetailsModelTypeEnum {
	values := make([]CreateEntityShapeDetailsModelTypeEnum, 0)
	for _, v := range mappingCreateEntityShapeDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateEntityShapeDetailsModelTypeEnumStringValues Enumerates the set of values in String for CreateEntityShapeDetailsModelTypeEnum
func GetCreateEntityShapeDetailsModelTypeEnumStringValues() []string {
	return []string{
		"FILE_ENTITY",
		"SQL_ENTITY",
		"OBJECT_ENTITY",
	}
}

// GetMappingCreateEntityShapeDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateEntityShapeDetailsModelTypeEnum(val string) (CreateEntityShapeDetailsModelTypeEnum, bool) {
	enum, ok := mappingCreateEntityShapeDetailsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
