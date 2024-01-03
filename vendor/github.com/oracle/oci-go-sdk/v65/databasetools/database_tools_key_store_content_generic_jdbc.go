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

// DatabaseToolsKeyStoreContentGenericJdbc The key store content.
type DatabaseToolsKeyStoreContentGenericJdbc interface {
}

type databasetoolskeystorecontentgenericjdbc struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorecontentgenericjdbc) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorecontentgenericjdbc databasetoolskeystorecontentgenericjdbc
	s := struct {
		Model Unmarshalerdatabasetoolskeystorecontentgenericjdbc
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorecontentgenericjdbc) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStoreContentSecretIdGenericJdbc{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DatabaseToolsKeyStoreContentGenericJdbc: %s.", m.ValueType)
		return *m, nil
	}
}

func (m databasetoolskeystorecontentgenericjdbc) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorecontentgenericjdbc) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStoreContentGenericJdbcValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStoreContentGenericJdbcValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStoreContentGenericJdbcValueTypeEnum
const (
	DatabaseToolsKeyStoreContentGenericJdbcValueTypeSecretid DatabaseToolsKeyStoreContentGenericJdbcValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStoreContentGenericJdbcValueTypeEnum = map[string]DatabaseToolsKeyStoreContentGenericJdbcValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStoreContentGenericJdbcValueTypeSecretid,
}

var mappingDatabaseToolsKeyStoreContentGenericJdbcValueTypeEnumLowerCase = map[string]DatabaseToolsKeyStoreContentGenericJdbcValueTypeEnum{
	"secretid": DatabaseToolsKeyStoreContentGenericJdbcValueTypeSecretid,
}

// GetDatabaseToolsKeyStoreContentGenericJdbcValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStoreContentGenericJdbcValueTypeEnum
func GetDatabaseToolsKeyStoreContentGenericJdbcValueTypeEnumValues() []DatabaseToolsKeyStoreContentGenericJdbcValueTypeEnum {
	values := make([]DatabaseToolsKeyStoreContentGenericJdbcValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStoreContentGenericJdbcValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStoreContentGenericJdbcValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStoreContentGenericJdbcValueTypeEnum
func GetDatabaseToolsKeyStoreContentGenericJdbcValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStoreContentGenericJdbcValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStoreContentGenericJdbcValueTypeEnum(val string) (DatabaseToolsKeyStoreContentGenericJdbcValueTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsKeyStoreContentGenericJdbcValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
