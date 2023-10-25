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

// DatabaseToolsKeyStorePasswordGenericJdbcDetails The key store password.
type DatabaseToolsKeyStorePasswordGenericJdbcDetails interface {
}

type databasetoolskeystorepasswordgenericjdbcdetails struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorepasswordgenericjdbcdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorepasswordgenericjdbcdetails databasetoolskeystorepasswordgenericjdbcdetails
	s := struct {
		Model Unmarshalerdatabasetoolskeystorepasswordgenericjdbcdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorepasswordgenericjdbcdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStorePasswordSecretIdGenericJdbcDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DatabaseToolsKeyStorePasswordGenericJdbcDetails: %s.", m.ValueType)
		return *m, nil
	}
}

func (m databasetoolskeystorepasswordgenericjdbcdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorepasswordgenericjdbcdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeEnum
const (
	DatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeSecretid DatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeEnum = map[string]DatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeSecretid,
}

var mappingDatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeEnumLowerCase = map[string]DatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeEnum{
	"secretid": DatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeSecretid,
}

// GetDatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeEnum
func GetDatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeEnumValues() []DatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeEnum {
	values := make([]DatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeEnum
func GetDatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeEnum(val string) (DatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsKeyStorePasswordGenericJdbcDetailsValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
