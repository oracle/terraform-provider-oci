// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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
	case "NONE":
		mm := CreateDbSystemSourceFromNoneDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IMPORTURL":
		mm := CreateDbSystemSourceImportFromUrlDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PITR":
		mm := CreateDbSystemSourceFromPitrDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m createdbsystemsourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdbsystemsourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDbSystemSourceDetailsSourceTypeEnum Enum with underlying type: string
type CreateDbSystemSourceDetailsSourceTypeEnum string

// Set of constants representing the allowable values for CreateDbSystemSourceDetailsSourceTypeEnum
const (
	CreateDbSystemSourceDetailsSourceTypeNone      CreateDbSystemSourceDetailsSourceTypeEnum = "NONE"
	CreateDbSystemSourceDetailsSourceTypeBackup    CreateDbSystemSourceDetailsSourceTypeEnum = "BACKUP"
	CreateDbSystemSourceDetailsSourceTypePitr      CreateDbSystemSourceDetailsSourceTypeEnum = "PITR"
	CreateDbSystemSourceDetailsSourceTypeImporturl CreateDbSystemSourceDetailsSourceTypeEnum = "IMPORTURL"
)

var mappingCreateDbSystemSourceDetailsSourceTypeEnum = map[string]CreateDbSystemSourceDetailsSourceTypeEnum{
	"NONE":      CreateDbSystemSourceDetailsSourceTypeNone,
	"BACKUP":    CreateDbSystemSourceDetailsSourceTypeBackup,
	"PITR":      CreateDbSystemSourceDetailsSourceTypePitr,
	"IMPORTURL": CreateDbSystemSourceDetailsSourceTypeImporturl,
}

var mappingCreateDbSystemSourceDetailsSourceTypeEnumLowerCase = map[string]CreateDbSystemSourceDetailsSourceTypeEnum{
	"none":      CreateDbSystemSourceDetailsSourceTypeNone,
	"backup":    CreateDbSystemSourceDetailsSourceTypeBackup,
	"pitr":      CreateDbSystemSourceDetailsSourceTypePitr,
	"importurl": CreateDbSystemSourceDetailsSourceTypeImporturl,
}

// GetCreateDbSystemSourceDetailsSourceTypeEnumValues Enumerates the set of values for CreateDbSystemSourceDetailsSourceTypeEnum
func GetCreateDbSystemSourceDetailsSourceTypeEnumValues() []CreateDbSystemSourceDetailsSourceTypeEnum {
	values := make([]CreateDbSystemSourceDetailsSourceTypeEnum, 0)
	for _, v := range mappingCreateDbSystemSourceDetailsSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDbSystemSourceDetailsSourceTypeEnumStringValues Enumerates the set of values in String for CreateDbSystemSourceDetailsSourceTypeEnum
func GetCreateDbSystemSourceDetailsSourceTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"BACKUP",
		"PITR",
		"IMPORTURL",
	}
}

// GetMappingCreateDbSystemSourceDetailsSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDbSystemSourceDetailsSourceTypeEnum(val string) (CreateDbSystemSourceDetailsSourceTypeEnum, bool) {
	enum, ok := mappingCreateDbSystemSourceDetailsSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
