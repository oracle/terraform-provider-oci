// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// DbSystemSource Parameters detailing how to provision the initial data of the DB System.
type DbSystemSource interface {
}

type dbsystemsource struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *dbsystemsource) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdbsystemsource dbsystemsource
	s := struct {
		Model Unmarshalerdbsystemsource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dbsystemsource) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "BACKUP":
		mm := DbSystemSourceFromBackup{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IMPORTURL":
		mm := DbSystemSourceImportFromUrl{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m dbsystemsource) String() string {
	return common.PointerString(m)
}

// DbSystemSourceSourceTypeEnum Enum with underlying type: string
type DbSystemSourceSourceTypeEnum string

// Set of constants representing the allowable values for DbSystemSourceSourceTypeEnum
const (
	DbSystemSourceSourceTypeNone      DbSystemSourceSourceTypeEnum = "NONE"
	DbSystemSourceSourceTypeBackup    DbSystemSourceSourceTypeEnum = "BACKUP"
	DbSystemSourceSourceTypeImporturl DbSystemSourceSourceTypeEnum = "IMPORTURL"
)

var mappingDbSystemSourceSourceType = map[string]DbSystemSourceSourceTypeEnum{
	"NONE":      DbSystemSourceSourceTypeNone,
	"BACKUP":    DbSystemSourceSourceTypeBackup,
	"IMPORTURL": DbSystemSourceSourceTypeImporturl,
}

// GetDbSystemSourceSourceTypeEnumValues Enumerates the set of values for DbSystemSourceSourceTypeEnum
func GetDbSystemSourceSourceTypeEnumValues() []DbSystemSourceSourceTypeEnum {
	values := make([]DbSystemSourceSourceTypeEnum, 0)
	for _, v := range mappingDbSystemSourceSourceType {
		values = append(values, v)
	}
	return values
}
