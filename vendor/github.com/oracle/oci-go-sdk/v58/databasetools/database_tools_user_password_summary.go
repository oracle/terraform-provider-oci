// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Database Tools APIs to manage Connections and Private Endpoints.
//

package databasetools

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// DatabaseToolsUserPasswordSummary The user password.
type DatabaseToolsUserPasswordSummary interface {
}

type databasetoolsuserpasswordsummary struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolsuserpasswordsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolsuserpasswordsummary databasetoolsuserpasswordsummary
	s := struct {
		Model Unmarshalerdatabasetoolsuserpasswordsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolsuserpasswordsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsUserPasswordSecretIdSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m databasetoolsuserpasswordsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolsuserpasswordsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsUserPasswordSummaryValueTypeEnum Enum with underlying type: string
type DatabaseToolsUserPasswordSummaryValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsUserPasswordSummaryValueTypeEnum
const (
	DatabaseToolsUserPasswordSummaryValueTypeSecretid DatabaseToolsUserPasswordSummaryValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsUserPasswordSummaryValueTypeEnum = map[string]DatabaseToolsUserPasswordSummaryValueTypeEnum{
	"SECRETID": DatabaseToolsUserPasswordSummaryValueTypeSecretid,
}

// GetDatabaseToolsUserPasswordSummaryValueTypeEnumValues Enumerates the set of values for DatabaseToolsUserPasswordSummaryValueTypeEnum
func GetDatabaseToolsUserPasswordSummaryValueTypeEnumValues() []DatabaseToolsUserPasswordSummaryValueTypeEnum {
	values := make([]DatabaseToolsUserPasswordSummaryValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsUserPasswordSummaryValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsUserPasswordSummaryValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsUserPasswordSummaryValueTypeEnum
func GetDatabaseToolsUserPasswordSummaryValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsUserPasswordSummaryValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsUserPasswordSummaryValueTypeEnum(val string) (DatabaseToolsUserPasswordSummaryValueTypeEnum, bool) {
	mappingDatabaseToolsUserPasswordSummaryValueTypeEnumIgnoreCase := make(map[string]DatabaseToolsUserPasswordSummaryValueTypeEnum)
	for k, v := range mappingDatabaseToolsUserPasswordSummaryValueTypeEnum {
		mappingDatabaseToolsUserPasswordSummaryValueTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDatabaseToolsUserPasswordSummaryValueTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
