// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// DatabaseToolsKeyStorePasswordMySqlDetails The key store password.
type DatabaseToolsKeyStorePasswordMySqlDetails interface {
}

type databasetoolskeystorepasswordmysqldetails struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorepasswordmysqldetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorepasswordmysqldetails databasetoolskeystorepasswordmysqldetails
	s := struct {
		Model Unmarshalerdatabasetoolskeystorepasswordmysqldetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorepasswordmysqldetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStorePasswordSecretIdMySqlDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DatabaseToolsKeyStorePasswordMySqlDetails: %s.", m.ValueType)
		return *m, nil
	}
}

func (m databasetoolskeystorepasswordmysqldetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorepasswordmysqldetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStorePasswordMySqlDetailsValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStorePasswordMySqlDetailsValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStorePasswordMySqlDetailsValueTypeEnum
const (
	DatabaseToolsKeyStorePasswordMySqlDetailsValueTypeSecretid DatabaseToolsKeyStorePasswordMySqlDetailsValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStorePasswordMySqlDetailsValueTypeEnum = map[string]DatabaseToolsKeyStorePasswordMySqlDetailsValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStorePasswordMySqlDetailsValueTypeSecretid,
}

var mappingDatabaseToolsKeyStorePasswordMySqlDetailsValueTypeEnumLowerCase = map[string]DatabaseToolsKeyStorePasswordMySqlDetailsValueTypeEnum{
	"secretid": DatabaseToolsKeyStorePasswordMySqlDetailsValueTypeSecretid,
}

// GetDatabaseToolsKeyStorePasswordMySqlDetailsValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStorePasswordMySqlDetailsValueTypeEnum
func GetDatabaseToolsKeyStorePasswordMySqlDetailsValueTypeEnumValues() []DatabaseToolsKeyStorePasswordMySqlDetailsValueTypeEnum {
	values := make([]DatabaseToolsKeyStorePasswordMySqlDetailsValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStorePasswordMySqlDetailsValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStorePasswordMySqlDetailsValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStorePasswordMySqlDetailsValueTypeEnum
func GetDatabaseToolsKeyStorePasswordMySqlDetailsValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStorePasswordMySqlDetailsValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStorePasswordMySqlDetailsValueTypeEnum(val string) (DatabaseToolsKeyStorePasswordMySqlDetailsValueTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsKeyStorePasswordMySqlDetailsValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
