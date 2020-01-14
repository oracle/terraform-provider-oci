// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateDatabaseBase Details for creating a database.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateDatabaseBase interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Database Home.
	GetDbHomeId() *string

	// A valid Oracle Database version. To get a list of supported versions, use the ListDbVersions operation.
	GetDbVersion() *string
}

type createdatabasebase struct {
	JsonData  []byte
	DbHomeId  *string `mandatory:"true" json:"dbHomeId"`
	DbVersion *string `mandatory:"false" json:"dbVersion"`
	Source    string  `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *createdatabasebase) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedatabasebase createdatabasebase
	s := struct {
		Model Unmarshalercreatedatabasebase
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DbHomeId = s.Model.DbHomeId
	m.DbVersion = s.Model.DbVersion
	m.Source = s.Model.Source

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdatabasebase) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "NONE":
		mm := CreateNewDatabaseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDbHomeId returns DbHomeId
func (m createdatabasebase) GetDbHomeId() *string {
	return m.DbHomeId
}

//GetDbVersion returns DbVersion
func (m createdatabasebase) GetDbVersion() *string {
	return m.DbVersion
}

func (m createdatabasebase) String() string {
	return common.PointerString(m)
}

// CreateDatabaseBaseSourceEnum Enum with underlying type: string
type CreateDatabaseBaseSourceEnum string

// Set of constants representing the allowable values for CreateDatabaseBaseSourceEnum
const (
	CreateDatabaseBaseSourceNone CreateDatabaseBaseSourceEnum = "NONE"
)

var mappingCreateDatabaseBaseSource = map[string]CreateDatabaseBaseSourceEnum{
	"NONE": CreateDatabaseBaseSourceNone,
}

// GetCreateDatabaseBaseSourceEnumValues Enumerates the set of values for CreateDatabaseBaseSourceEnum
func GetCreateDatabaseBaseSourceEnumValues() []CreateDatabaseBaseSourceEnum {
	values := make([]CreateDatabaseBaseSourceEnum, 0)
	for _, v := range mappingCreateDatabaseBaseSource {
		values = append(values, v)
	}
	return values
}
