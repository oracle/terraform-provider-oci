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

// DataEntityDetailsModelTypeEnum Enum with underlying type: string
type DataEntityDetailsModelTypeEnum string

// Set of constants representing the allowable values for DataEntityDetailsModelTypeEnum
const (
	DataEntityDetailsModelTypeViewEntity  DataEntityDetailsModelTypeEnum = "VIEW_ENTITY"
	DataEntityDetailsModelTypeTableEntity DataEntityDetailsModelTypeEnum = "TABLE_ENTITY"
	DataEntityDetailsModelTypeFileEntity  DataEntityDetailsModelTypeEnum = "FILE_ENTITY"
)

var mappingDataEntityDetailsModelType = map[string]DataEntityDetailsModelTypeEnum{
	"VIEW_ENTITY":  DataEntityDetailsModelTypeViewEntity,
	"TABLE_ENTITY": DataEntityDetailsModelTypeTableEntity,
	"FILE_ENTITY":  DataEntityDetailsModelTypeFileEntity,
}

// GetDataEntityDetailsModelTypeEnumValues Enumerates the set of values for DataEntityDetailsModelTypeEnum
func GetDataEntityDetailsModelTypeEnumValues() []DataEntityDetailsModelTypeEnum {
	values := make([]DataEntityDetailsModelTypeEnum, 0)
	for _, v := range mappingDataEntityDetailsModelType {
		values = append(values, v)
	}
	return values
}
