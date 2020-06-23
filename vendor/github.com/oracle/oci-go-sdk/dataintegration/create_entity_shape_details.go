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
	case "FILE_ENTITY":
		mm := CreateEntityShapeFromFile{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m createentityshapedetails) String() string {
	return common.PointerString(m)
}

// CreateEntityShapeDetailsModelTypeEnum Enum with underlying type: string
type CreateEntityShapeDetailsModelTypeEnum string

// Set of constants representing the allowable values for CreateEntityShapeDetailsModelTypeEnum
const (
	CreateEntityShapeDetailsModelTypeFileEntity CreateEntityShapeDetailsModelTypeEnum = "FILE_ENTITY"
)

var mappingCreateEntityShapeDetailsModelType = map[string]CreateEntityShapeDetailsModelTypeEnum{
	"FILE_ENTITY": CreateEntityShapeDetailsModelTypeFileEntity,
}

// GetCreateEntityShapeDetailsModelTypeEnumValues Enumerates the set of values for CreateEntityShapeDetailsModelTypeEnum
func GetCreateEntityShapeDetailsModelTypeEnumValues() []CreateEntityShapeDetailsModelTypeEnum {
	values := make([]CreateEntityShapeDetailsModelTypeEnum, 0)
	for _, v := range mappingCreateEntityShapeDetailsModelType {
		values = append(values, v)
	}
	return values
}
