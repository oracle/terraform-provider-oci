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

// DatabaseToolsKeyStoreContentMySql The key store content.
type DatabaseToolsKeyStoreContentMySql interface {
}

type databasetoolskeystorecontentmysql struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorecontentmysql) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorecontentmysql databasetoolskeystorecontentmysql
	s := struct {
		Model Unmarshalerdatabasetoolskeystorecontentmysql
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorecontentmysql) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStoreContentSecretIdMySql{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DatabaseToolsKeyStoreContentMySql: %s.", m.ValueType)
		return *m, nil
	}
}

func (m databasetoolskeystorecontentmysql) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorecontentmysql) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStoreContentMySqlValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStoreContentMySqlValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStoreContentMySqlValueTypeEnum
const (
	DatabaseToolsKeyStoreContentMySqlValueTypeSecretid DatabaseToolsKeyStoreContentMySqlValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStoreContentMySqlValueTypeEnum = map[string]DatabaseToolsKeyStoreContentMySqlValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStoreContentMySqlValueTypeSecretid,
}

var mappingDatabaseToolsKeyStoreContentMySqlValueTypeEnumLowerCase = map[string]DatabaseToolsKeyStoreContentMySqlValueTypeEnum{
	"secretid": DatabaseToolsKeyStoreContentMySqlValueTypeSecretid,
}

// GetDatabaseToolsKeyStoreContentMySqlValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStoreContentMySqlValueTypeEnum
func GetDatabaseToolsKeyStoreContentMySqlValueTypeEnumValues() []DatabaseToolsKeyStoreContentMySqlValueTypeEnum {
	values := make([]DatabaseToolsKeyStoreContentMySqlValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStoreContentMySqlValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStoreContentMySqlValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStoreContentMySqlValueTypeEnum
func GetDatabaseToolsKeyStoreContentMySqlValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStoreContentMySqlValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStoreContentMySqlValueTypeEnum(val string) (DatabaseToolsKeyStoreContentMySqlValueTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsKeyStoreContentMySqlValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
