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

// CreateDbSystemSourceDetails Parameters detailing how to provision the initial data of the system.
type CreateDbSystemSourceDetails interface {
}

type createdbsystemsourcedetails struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *createdbsystemsourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedbsystemsourcedetails createdbsystemsourcedetails
	s := struct {
		Model Unmarshalercreatedbsystemsourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdbsystemsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "BACKUP":
		mm := CreateDbSystemSourceFromBackupDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IMPORTURL":
		mm := CreateDbSystemSourceImportFromUrlDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m createdbsystemsourcedetails) String() string {
	return common.PointerString(m)
}

// CreateDbSystemSourceDetailsSourceTypeEnum Enum with underlying type: string
type CreateDbSystemSourceDetailsSourceTypeEnum string

// Set of constants representing the allowable values for CreateDbSystemSourceDetailsSourceTypeEnum
const (
	CreateDbSystemSourceDetailsSourceTypeNone      CreateDbSystemSourceDetailsSourceTypeEnum = "NONE"
	CreateDbSystemSourceDetailsSourceTypeBackup    CreateDbSystemSourceDetailsSourceTypeEnum = "BACKUP"
	CreateDbSystemSourceDetailsSourceTypeImporturl CreateDbSystemSourceDetailsSourceTypeEnum = "IMPORTURL"
)

var mappingCreateDbSystemSourceDetailsSourceType = map[string]CreateDbSystemSourceDetailsSourceTypeEnum{
	"NONE":      CreateDbSystemSourceDetailsSourceTypeNone,
	"BACKUP":    CreateDbSystemSourceDetailsSourceTypeBackup,
	"IMPORTURL": CreateDbSystemSourceDetailsSourceTypeImporturl,
}

// GetCreateDbSystemSourceDetailsSourceTypeEnumValues Enumerates the set of values for CreateDbSystemSourceDetailsSourceTypeEnum
func GetCreateDbSystemSourceDetailsSourceTypeEnumValues() []CreateDbSystemSourceDetailsSourceTypeEnum {
	values := make([]CreateDbSystemSourceDetailsSourceTypeEnum, 0)
	for _, v := range mappingCreateDbSystemSourceDetailsSourceType {
		values = append(values, v)
	}
	return values
}
