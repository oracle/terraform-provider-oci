// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseToolsKeyStorePasswordGenericJdbc The key store password.
type DatabaseToolsKeyStorePasswordGenericJdbc interface {
}

type databasetoolskeystorepasswordgenericjdbc struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorepasswordgenericjdbc) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorepasswordgenericjdbc databasetoolskeystorepasswordgenericjdbc
	s := struct {
		Model Unmarshalerdatabasetoolskeystorepasswordgenericjdbc
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorepasswordgenericjdbc) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStorePasswordSecretIdGenericJdbc{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DatabaseToolsKeyStorePasswordGenericJdbc: %s.", m.ValueType)
		return *m, nil
	}
}

func (m databasetoolskeystorepasswordgenericjdbc) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorepasswordgenericjdbc) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStorePasswordGenericJdbcValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStorePasswordGenericJdbcValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStorePasswordGenericJdbcValueTypeEnum
const (
	DatabaseToolsKeyStorePasswordGenericJdbcValueTypeSecretid DatabaseToolsKeyStorePasswordGenericJdbcValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStorePasswordGenericJdbcValueTypeEnum = map[string]DatabaseToolsKeyStorePasswordGenericJdbcValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStorePasswordGenericJdbcValueTypeSecretid,
}

var mappingDatabaseToolsKeyStorePasswordGenericJdbcValueTypeEnumLowerCase = map[string]DatabaseToolsKeyStorePasswordGenericJdbcValueTypeEnum{
	"secretid": DatabaseToolsKeyStorePasswordGenericJdbcValueTypeSecretid,
}

// GetDatabaseToolsKeyStorePasswordGenericJdbcValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStorePasswordGenericJdbcValueTypeEnum
func GetDatabaseToolsKeyStorePasswordGenericJdbcValueTypeEnumValues() []DatabaseToolsKeyStorePasswordGenericJdbcValueTypeEnum {
	values := make([]DatabaseToolsKeyStorePasswordGenericJdbcValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStorePasswordGenericJdbcValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStorePasswordGenericJdbcValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStorePasswordGenericJdbcValueTypeEnum
func GetDatabaseToolsKeyStorePasswordGenericJdbcValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStorePasswordGenericJdbcValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStorePasswordGenericJdbcValueTypeEnum(val string) (DatabaseToolsKeyStorePasswordGenericJdbcValueTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsKeyStorePasswordGenericJdbcValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
