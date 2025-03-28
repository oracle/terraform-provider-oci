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

// DatabaseToolsKeyStorePasswordPostgresql The key store password.
type DatabaseToolsKeyStorePasswordPostgresql interface {
}

type databasetoolskeystorepasswordpostgresql struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorepasswordpostgresql) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorepasswordpostgresql databasetoolskeystorepasswordpostgresql
	s := struct {
		Model Unmarshalerdatabasetoolskeystorepasswordpostgresql
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorepasswordpostgresql) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStorePasswordSecretIdPostgresql{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DatabaseToolsKeyStorePasswordPostgresql: %s.", m.ValueType)
		return *m, nil
	}
}

func (m databasetoolskeystorepasswordpostgresql) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorepasswordpostgresql) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStorePasswordPostgresqlValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStorePasswordPostgresqlValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStorePasswordPostgresqlValueTypeEnum
const (
	DatabaseToolsKeyStorePasswordPostgresqlValueTypeSecretid DatabaseToolsKeyStorePasswordPostgresqlValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStorePasswordPostgresqlValueTypeEnum = map[string]DatabaseToolsKeyStorePasswordPostgresqlValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStorePasswordPostgresqlValueTypeSecretid,
}

var mappingDatabaseToolsKeyStorePasswordPostgresqlValueTypeEnumLowerCase = map[string]DatabaseToolsKeyStorePasswordPostgresqlValueTypeEnum{
	"secretid": DatabaseToolsKeyStorePasswordPostgresqlValueTypeSecretid,
}

// GetDatabaseToolsKeyStorePasswordPostgresqlValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStorePasswordPostgresqlValueTypeEnum
func GetDatabaseToolsKeyStorePasswordPostgresqlValueTypeEnumValues() []DatabaseToolsKeyStorePasswordPostgresqlValueTypeEnum {
	values := make([]DatabaseToolsKeyStorePasswordPostgresqlValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStorePasswordPostgresqlValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStorePasswordPostgresqlValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStorePasswordPostgresqlValueTypeEnum
func GetDatabaseToolsKeyStorePasswordPostgresqlValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStorePasswordPostgresqlValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStorePasswordPostgresqlValueTypeEnum(val string) (DatabaseToolsKeyStorePasswordPostgresqlValueTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsKeyStorePasswordPostgresqlValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
