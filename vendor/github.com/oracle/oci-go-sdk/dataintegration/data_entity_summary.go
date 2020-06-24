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

// DataEntitySummary The data entity summary object.
type DataEntitySummary interface {
	GetMetadata() *ObjectMetadata
}

type dataentitysummary struct {
	JsonData  []byte
	Metadata  *ObjectMetadata `mandatory:"false" json:"metadata"`
	ModelType string          `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *dataentitysummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdataentitysummary dataentitysummary
	s := struct {
		Model Unmarshalerdataentitysummary
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
func (m *dataentitysummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "FILE_ENTITY":
		mm := DataEntitySummaryFromFile{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TABLE_ENTITY":
		mm := DataEntitySummaryFromTable{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VIEW_ENTITY":
		mm := DataEntitySummaryFromView{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetMetadata returns Metadata
func (m dataentitysummary) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m dataentitysummary) String() string {
	return common.PointerString(m)
}

// DataEntitySummaryModelTypeEnum Enum with underlying type: string
type DataEntitySummaryModelTypeEnum string

// Set of constants representing the allowable values for DataEntitySummaryModelTypeEnum
const (
	DataEntitySummaryModelTypeViewEntity  DataEntitySummaryModelTypeEnum = "VIEW_ENTITY"
	DataEntitySummaryModelTypeTableEntity DataEntitySummaryModelTypeEnum = "TABLE_ENTITY"
	DataEntitySummaryModelTypeFileEntity  DataEntitySummaryModelTypeEnum = "FILE_ENTITY"
)

var mappingDataEntitySummaryModelType = map[string]DataEntitySummaryModelTypeEnum{
	"VIEW_ENTITY":  DataEntitySummaryModelTypeViewEntity,
	"TABLE_ENTITY": DataEntitySummaryModelTypeTableEntity,
	"FILE_ENTITY":  DataEntitySummaryModelTypeFileEntity,
}

// GetDataEntitySummaryModelTypeEnumValues Enumerates the set of values for DataEntitySummaryModelTypeEnum
func GetDataEntitySummaryModelTypeEnumValues() []DataEntitySummaryModelTypeEnum {
	values := make([]DataEntitySummaryModelTypeEnum, 0)
	for _, v := range mappingDataEntitySummaryModelType {
		values = append(values, v)
	}
	return values
}
