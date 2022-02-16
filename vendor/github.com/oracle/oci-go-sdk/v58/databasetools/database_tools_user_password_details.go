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

// DatabaseToolsUserPasswordDetails The user password.
type DatabaseToolsUserPasswordDetails interface {
}

type databasetoolsuserpassworddetails struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolsuserpassworddetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolsuserpassworddetails databasetoolsuserpassworddetails
	s := struct {
		Model Unmarshalerdatabasetoolsuserpassworddetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolsuserpassworddetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsUserPasswordSecretIdDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m databasetoolsuserpassworddetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolsuserpassworddetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsUserPasswordDetailsValueTypeEnum Enum with underlying type: string
type DatabaseToolsUserPasswordDetailsValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsUserPasswordDetailsValueTypeEnum
const (
	DatabaseToolsUserPasswordDetailsValueTypeSecretid DatabaseToolsUserPasswordDetailsValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsUserPasswordDetailsValueTypeEnum = map[string]DatabaseToolsUserPasswordDetailsValueTypeEnum{
	"SECRETID": DatabaseToolsUserPasswordDetailsValueTypeSecretid,
}

// GetDatabaseToolsUserPasswordDetailsValueTypeEnumValues Enumerates the set of values for DatabaseToolsUserPasswordDetailsValueTypeEnum
func GetDatabaseToolsUserPasswordDetailsValueTypeEnumValues() []DatabaseToolsUserPasswordDetailsValueTypeEnum {
	values := make([]DatabaseToolsUserPasswordDetailsValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsUserPasswordDetailsValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsUserPasswordDetailsValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsUserPasswordDetailsValueTypeEnum
func GetDatabaseToolsUserPasswordDetailsValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsUserPasswordDetailsValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsUserPasswordDetailsValueTypeEnum(val string) (DatabaseToolsUserPasswordDetailsValueTypeEnum, bool) {
	mappingDatabaseToolsUserPasswordDetailsValueTypeEnumIgnoreCase := make(map[string]DatabaseToolsUserPasswordDetailsValueTypeEnum)
	for k, v := range mappingDatabaseToolsUserPasswordDetailsValueTypeEnum {
		mappingDatabaseToolsUserPasswordDetailsValueTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDatabaseToolsUserPasswordDetailsValueTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
