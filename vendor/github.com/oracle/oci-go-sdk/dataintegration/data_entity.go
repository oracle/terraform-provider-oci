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
	case "VIEW_ENTITY":
		mm := DataEntityFromView{}
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

// DataEntityModelTypeEnum Enum with underlying type: string
type DataEntityModelTypeEnum string

// Set of constants representing the allowable values for DataEntityModelTypeEnum
const (
	DataEntityModelTypeViewEntity  DataEntityModelTypeEnum = "VIEW_ENTITY"
	DataEntityModelTypeTableEntity DataEntityModelTypeEnum = "TABLE_ENTITY"
	DataEntityModelTypeFileEntity  DataEntityModelTypeEnum = "FILE_ENTITY"
)

var mappingDataEntityModelType = map[string]DataEntityModelTypeEnum{
	"VIEW_ENTITY":  DataEntityModelTypeViewEntity,
	"TABLE_ENTITY": DataEntityModelTypeTableEntity,
	"FILE_ENTITY":  DataEntityModelTypeFileEntity,
}

// GetDataEntityModelTypeEnumValues Enumerates the set of values for DataEntityModelTypeEnum
func GetDataEntityModelTypeEnumValues() []DataEntityModelTypeEnum {
	values := make([]DataEntityModelTypeEnum, 0)
	for _, v := range mappingDataEntityModelType {
		values = append(values, v)
	}
	return values
}
