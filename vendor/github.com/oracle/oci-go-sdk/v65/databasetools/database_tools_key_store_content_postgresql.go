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

// DatabaseToolsKeyStoreContentPostgresql The key store content.
type DatabaseToolsKeyStoreContentPostgresql interface {
}

type databasetoolskeystorecontentpostgresql struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorecontentpostgresql) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorecontentpostgresql databasetoolskeystorecontentpostgresql
	s := struct {
		Model Unmarshalerdatabasetoolskeystorecontentpostgresql
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorecontentpostgresql) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStoreContentSecretIdPostgresql{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DatabaseToolsKeyStoreContentPostgresql: %s.", m.ValueType)
		return *m, nil
	}
}

func (m databasetoolskeystorecontentpostgresql) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorecontentpostgresql) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStoreContentPostgresqlValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStoreContentPostgresqlValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStoreContentPostgresqlValueTypeEnum
const (
	DatabaseToolsKeyStoreContentPostgresqlValueTypeSecretid DatabaseToolsKeyStoreContentPostgresqlValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStoreContentPostgresqlValueTypeEnum = map[string]DatabaseToolsKeyStoreContentPostgresqlValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStoreContentPostgresqlValueTypeSecretid,
}

var mappingDatabaseToolsKeyStoreContentPostgresqlValueTypeEnumLowerCase = map[string]DatabaseToolsKeyStoreContentPostgresqlValueTypeEnum{
	"secretid": DatabaseToolsKeyStoreContentPostgresqlValueTypeSecretid,
}

// GetDatabaseToolsKeyStoreContentPostgresqlValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStoreContentPostgresqlValueTypeEnum
func GetDatabaseToolsKeyStoreContentPostgresqlValueTypeEnumValues() []DatabaseToolsKeyStoreContentPostgresqlValueTypeEnum {
	values := make([]DatabaseToolsKeyStoreContentPostgresqlValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStoreContentPostgresqlValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStoreContentPostgresqlValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStoreContentPostgresqlValueTypeEnum
func GetDatabaseToolsKeyStoreContentPostgresqlValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStoreContentPostgresqlValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStoreContentPostgresqlValueTypeEnum(val string) (DatabaseToolsKeyStoreContentPostgresqlValueTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsKeyStoreContentPostgresqlValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
