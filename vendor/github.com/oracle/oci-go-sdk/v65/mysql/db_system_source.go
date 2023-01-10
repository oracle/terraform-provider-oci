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
	case "PITR":
		mm := DbSystemSourceFromPitr{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NONE":
		mm := DbSystemSourceFromNone{}
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dbsystemsource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbSystemSourceSourceTypeEnum Enum with underlying type: string
type DbSystemSourceSourceTypeEnum string

// Set of constants representing the allowable values for DbSystemSourceSourceTypeEnum
const (
	DbSystemSourceSourceTypeNone      DbSystemSourceSourceTypeEnum = "NONE"
	DbSystemSourceSourceTypeBackup    DbSystemSourceSourceTypeEnum = "BACKUP"
	DbSystemSourceSourceTypePitr      DbSystemSourceSourceTypeEnum = "PITR"
	DbSystemSourceSourceTypeImporturl DbSystemSourceSourceTypeEnum = "IMPORTURL"
)

var mappingDbSystemSourceSourceTypeEnum = map[string]DbSystemSourceSourceTypeEnum{
	"NONE":      DbSystemSourceSourceTypeNone,
	"BACKUP":    DbSystemSourceSourceTypeBackup,
	"PITR":      DbSystemSourceSourceTypePitr,
	"IMPORTURL": DbSystemSourceSourceTypeImporturl,
}

var mappingDbSystemSourceSourceTypeEnumLowerCase = map[string]DbSystemSourceSourceTypeEnum{
	"none":      DbSystemSourceSourceTypeNone,
	"backup":    DbSystemSourceSourceTypeBackup,
	"pitr":      DbSystemSourceSourceTypePitr,
	"importurl": DbSystemSourceSourceTypeImporturl,
}

// GetDbSystemSourceSourceTypeEnumValues Enumerates the set of values for DbSystemSourceSourceTypeEnum
func GetDbSystemSourceSourceTypeEnumValues() []DbSystemSourceSourceTypeEnum {
	values := make([]DbSystemSourceSourceTypeEnum, 0)
	for _, v := range mappingDbSystemSourceSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemSourceSourceTypeEnumStringValues Enumerates the set of values in String for DbSystemSourceSourceTypeEnum
func GetDbSystemSourceSourceTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"BACKUP",
		"PITR",
		"IMPORTURL",
	}
}

// GetMappingDbSystemSourceSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbSystemSourceSourceTypeEnum(val string) (DbSystemSourceSourceTypeEnum, bool) {
	enum, ok := mappingDbSystemSourceSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
