// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// DatabaseToolsKeyStorePasswordMySql The key store password.
type DatabaseToolsKeyStorePasswordMySql interface {
}

type databasetoolskeystorepasswordmysql struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorepasswordmysql) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorepasswordmysql databasetoolskeystorepasswordmysql
	s := struct {
		Model Unmarshalerdatabasetoolskeystorepasswordmysql
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorepasswordmysql) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStorePasswordSecretIdMySql{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DatabaseToolsKeyStorePasswordMySql: %s.", m.ValueType)
		return *m, nil
	}
}

func (m databasetoolskeystorepasswordmysql) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorepasswordmysql) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStorePasswordMySqlValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStorePasswordMySqlValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStorePasswordMySqlValueTypeEnum
const (
	DatabaseToolsKeyStorePasswordMySqlValueTypeSecretid DatabaseToolsKeyStorePasswordMySqlValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStorePasswordMySqlValueTypeEnum = map[string]DatabaseToolsKeyStorePasswordMySqlValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStorePasswordMySqlValueTypeSecretid,
}

var mappingDatabaseToolsKeyStorePasswordMySqlValueTypeEnumLowerCase = map[string]DatabaseToolsKeyStorePasswordMySqlValueTypeEnum{
	"secretid": DatabaseToolsKeyStorePasswordMySqlValueTypeSecretid,
}

// GetDatabaseToolsKeyStorePasswordMySqlValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStorePasswordMySqlValueTypeEnum
func GetDatabaseToolsKeyStorePasswordMySqlValueTypeEnumValues() []DatabaseToolsKeyStorePasswordMySqlValueTypeEnum {
	values := make([]DatabaseToolsKeyStorePasswordMySqlValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStorePasswordMySqlValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStorePasswordMySqlValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStorePasswordMySqlValueTypeEnum
func GetDatabaseToolsKeyStorePasswordMySqlValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStorePasswordMySqlValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStorePasswordMySqlValueTypeEnum(val string) (DatabaseToolsKeyStorePasswordMySqlValueTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsKeyStorePasswordMySqlValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
