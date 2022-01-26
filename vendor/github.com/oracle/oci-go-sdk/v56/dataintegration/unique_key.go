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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UniqueKey The unqique key object.
type UniqueKey interface {

	// The object key.
	GetKey() *string

	// The object's model version.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	GetName() *string

	// An array of attribute references.
	GetAttributeRefs() []KeyAttribute

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int
}

type uniquekey struct {
	JsonData      []byte
	Key           *string          `mandatory:"false" json:"key"`
	ModelVersion  *string          `mandatory:"false" json:"modelVersion"`
	ParentRef     *ParentReference `mandatory:"false" json:"parentRef"`
	Name          *string          `mandatory:"false" json:"name"`
	AttributeRefs []KeyAttribute   `mandatory:"false" json:"attributeRefs"`
	ObjectStatus  *int             `mandatory:"false" json:"objectStatus"`
	ModelType     string           `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *uniquekey) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleruniquekey uniquekey
	s := struct {
		Model Unmarshaleruniquekey
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.ModelVersion = s.Model.ModelVersion
	m.ParentRef = s.Model.ParentRef
	m.Name = s.Model.Name
	m.AttributeRefs = s.Model.AttributeRefs
	m.ObjectStatus = s.Model.ObjectStatus
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *uniquekey) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "PRIMARY_KEY":
		mm := PrimaryKey{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "UNIQUE_KEY":
		mm := UniqueDataKey{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetKey returns Key
func (m uniquekey) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m uniquekey) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m uniquekey) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m uniquekey) GetName() *string {
	return m.Name
}

//GetAttributeRefs returns AttributeRefs
func (m uniquekey) GetAttributeRefs() []KeyAttribute {
	return m.AttributeRefs
}

//GetObjectStatus returns ObjectStatus
func (m uniquekey) GetObjectStatus() *int {
	return m.ObjectStatus
}

func (m uniquekey) String() string {
	return common.PointerString(m)
}

// UniqueKeyModelTypeEnum Enum with underlying type: string
type UniqueKeyModelTypeEnum string

// Set of constants representing the allowable values for UniqueKeyModelTypeEnum
const (
	UniqueKeyModelTypePrimaryKey UniqueKeyModelTypeEnum = "PRIMARY_KEY"
	UniqueKeyModelTypeUniqueKey  UniqueKeyModelTypeEnum = "UNIQUE_KEY"
)

var mappingUniqueKeyModelType = map[string]UniqueKeyModelTypeEnum{
	"PRIMARY_KEY": UniqueKeyModelTypePrimaryKey,
	"UNIQUE_KEY":  UniqueKeyModelTypeUniqueKey,
}

// GetUniqueKeyModelTypeEnumValues Enumerates the set of values for UniqueKeyModelTypeEnum
func GetUniqueKeyModelTypeEnumValues() []UniqueKeyModelTypeEnum {
	values := make([]UniqueKeyModelTypeEnum, 0)
	for _, v := range mappingUniqueKeyModelType {
		values = append(values, v)
	}
	return values
}
