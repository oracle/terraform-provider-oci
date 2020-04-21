// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
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
	case "DB_BACKUP":
		mm := CreateDatabaseFromBackup{}
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
	CreateDatabaseBaseSourceNone     CreateDatabaseBaseSourceEnum = "NONE"
	CreateDatabaseBaseSourceDbBackup CreateDatabaseBaseSourceEnum = "DB_BACKUP"
)

var mappingCreateDatabaseBaseSource = map[string]CreateDatabaseBaseSourceEnum{
	"NONE":      CreateDatabaseBaseSourceNone,
	"DB_BACKUP": CreateDatabaseBaseSourceDbBackup,
}

// GetCreateDatabaseBaseSourceEnumValues Enumerates the set of values for CreateDatabaseBaseSourceEnum
func GetCreateDatabaseBaseSourceEnumValues() []CreateDatabaseBaseSourceEnum {
	values := make([]CreateDatabaseBaseSourceEnum, 0)
	for _, v := range mappingCreateDatabaseBaseSource {
		values = append(values, v)
	}
	return values
}
