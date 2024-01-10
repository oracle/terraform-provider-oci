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

// DatabaseToolsKeyStorePassword The key store password.
type DatabaseToolsKeyStorePassword interface {
}

type databasetoolskeystorepassword struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorepassword) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorepassword databasetoolskeystorepassword
	s := struct {
		Model Unmarshalerdatabasetoolskeystorepassword
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorepassword) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStorePasswordSecretId{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DatabaseToolsKeyStorePassword: %s.", m.ValueType)
		return *m, nil
	}
}

func (m databasetoolskeystorepassword) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorepassword) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStorePasswordValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStorePasswordValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStorePasswordValueTypeEnum
const (
	DatabaseToolsKeyStorePasswordValueTypeSecretid DatabaseToolsKeyStorePasswordValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStorePasswordValueTypeEnum = map[string]DatabaseToolsKeyStorePasswordValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStorePasswordValueTypeSecretid,
}

var mappingDatabaseToolsKeyStorePasswordValueTypeEnumLowerCase = map[string]DatabaseToolsKeyStorePasswordValueTypeEnum{
	"secretid": DatabaseToolsKeyStorePasswordValueTypeSecretid,
}

// GetDatabaseToolsKeyStorePasswordValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStorePasswordValueTypeEnum
func GetDatabaseToolsKeyStorePasswordValueTypeEnumValues() []DatabaseToolsKeyStorePasswordValueTypeEnum {
	values := make([]DatabaseToolsKeyStorePasswordValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStorePasswordValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStorePasswordValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStorePasswordValueTypeEnum
func GetDatabaseToolsKeyStorePasswordValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStorePasswordValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStorePasswordValueTypeEnum(val string) (DatabaseToolsKeyStorePasswordValueTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsKeyStorePasswordValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
