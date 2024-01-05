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

// DatabaseToolsUserPassword The user password.
type DatabaseToolsUserPassword interface {
}

type databasetoolsuserpassword struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolsuserpassword) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolsuserpassword databasetoolsuserpassword
	s := struct {
		Model Unmarshalerdatabasetoolsuserpassword
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolsuserpassword) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsUserPasswordSecretId{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DatabaseToolsUserPassword: %s.", m.ValueType)
		return *m, nil
	}
}

func (m databasetoolsuserpassword) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolsuserpassword) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsUserPasswordValueTypeEnum Enum with underlying type: string
type DatabaseToolsUserPasswordValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsUserPasswordValueTypeEnum
const (
	DatabaseToolsUserPasswordValueTypeSecretid DatabaseToolsUserPasswordValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsUserPasswordValueTypeEnum = map[string]DatabaseToolsUserPasswordValueTypeEnum{
	"SECRETID": DatabaseToolsUserPasswordValueTypeSecretid,
}

var mappingDatabaseToolsUserPasswordValueTypeEnumLowerCase = map[string]DatabaseToolsUserPasswordValueTypeEnum{
	"secretid": DatabaseToolsUserPasswordValueTypeSecretid,
}

// GetDatabaseToolsUserPasswordValueTypeEnumValues Enumerates the set of values for DatabaseToolsUserPasswordValueTypeEnum
func GetDatabaseToolsUserPasswordValueTypeEnumValues() []DatabaseToolsUserPasswordValueTypeEnum {
	values := make([]DatabaseToolsUserPasswordValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsUserPasswordValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsUserPasswordValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsUserPasswordValueTypeEnum
func GetDatabaseToolsUserPasswordValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsUserPasswordValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsUserPasswordValueTypeEnum(val string) (DatabaseToolsUserPasswordValueTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsUserPasswordValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
