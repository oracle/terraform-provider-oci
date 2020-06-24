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

// EntityShapeModelTypeEnum Enum with underlying type: string
type EntityShapeModelTypeEnum string

// Set of constants representing the allowable values for EntityShapeModelTypeEnum
const (
	EntityShapeModelTypeFileEntity EntityShapeModelTypeEnum = "FILE_ENTITY"
)

var mappingEntityShapeModelType = map[string]EntityShapeModelTypeEnum{
	"FILE_ENTITY": EntityShapeModelTypeFileEntity,
}

// GetEntityShapeModelTypeEnumValues Enumerates the set of values for EntityShapeModelTypeEnum
func GetEntityShapeModelTypeEnumValues() []EntityShapeModelTypeEnum {
	values := make([]EntityShapeModelTypeEnum, 0)
	for _, v := range mappingEntityShapeModelType {
		values = append(values, v)
	}
	return values
}
