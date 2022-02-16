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

// DatabaseToolsKeyStorePasswordDetails The key store password.
type DatabaseToolsKeyStorePasswordDetails interface {
}

type databasetoolskeystorepassworddetails struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorepassworddetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorepassworddetails databasetoolskeystorepassworddetails
	s := struct {
		Model Unmarshalerdatabasetoolskeystorepassworddetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorepassworddetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStorePasswordSecretIdDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m databasetoolskeystorepassworddetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorepassworddetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStorePasswordDetailsValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStorePasswordDetailsValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStorePasswordDetailsValueTypeEnum
const (
	DatabaseToolsKeyStorePasswordDetailsValueTypeSecretid DatabaseToolsKeyStorePasswordDetailsValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStorePasswordDetailsValueTypeEnum = map[string]DatabaseToolsKeyStorePasswordDetailsValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStorePasswordDetailsValueTypeSecretid,
}

// GetDatabaseToolsKeyStorePasswordDetailsValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStorePasswordDetailsValueTypeEnum
func GetDatabaseToolsKeyStorePasswordDetailsValueTypeEnumValues() []DatabaseToolsKeyStorePasswordDetailsValueTypeEnum {
	values := make([]DatabaseToolsKeyStorePasswordDetailsValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStorePasswordDetailsValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStorePasswordDetailsValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStorePasswordDetailsValueTypeEnum
func GetDatabaseToolsKeyStorePasswordDetailsValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStorePasswordDetailsValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStorePasswordDetailsValueTypeEnum(val string) (DatabaseToolsKeyStorePasswordDetailsValueTypeEnum, bool) {
	mappingDatabaseToolsKeyStorePasswordDetailsValueTypeEnumIgnoreCase := make(map[string]DatabaseToolsKeyStorePasswordDetailsValueTypeEnum)
	for k, v := range mappingDatabaseToolsKeyStorePasswordDetailsValueTypeEnum {
		mappingDatabaseToolsKeyStorePasswordDetailsValueTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDatabaseToolsKeyStorePasswordDetailsValueTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
