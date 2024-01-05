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

// DatabaseToolsKeyStorePasswordMySqlSummary The key store password.
type DatabaseToolsKeyStorePasswordMySqlSummary interface {
}

type databasetoolskeystorepasswordmysqlsummary struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorepasswordmysqlsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorepasswordmysqlsummary databasetoolskeystorepasswordmysqlsummary
	s := struct {
		Model Unmarshalerdatabasetoolskeystorepasswordmysqlsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorepasswordmysqlsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStorePasswordSecretIdMySqlSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DatabaseToolsKeyStorePasswordMySqlSummary: %s.", m.ValueType)
		return *m, nil
	}
}

func (m databasetoolskeystorepasswordmysqlsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorepasswordmysqlsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStorePasswordMySqlSummaryValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStorePasswordMySqlSummaryValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStorePasswordMySqlSummaryValueTypeEnum
const (
	DatabaseToolsKeyStorePasswordMySqlSummaryValueTypeSecretid DatabaseToolsKeyStorePasswordMySqlSummaryValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStorePasswordMySqlSummaryValueTypeEnum = map[string]DatabaseToolsKeyStorePasswordMySqlSummaryValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStorePasswordMySqlSummaryValueTypeSecretid,
}

var mappingDatabaseToolsKeyStorePasswordMySqlSummaryValueTypeEnumLowerCase = map[string]DatabaseToolsKeyStorePasswordMySqlSummaryValueTypeEnum{
	"secretid": DatabaseToolsKeyStorePasswordMySqlSummaryValueTypeSecretid,
}

// GetDatabaseToolsKeyStorePasswordMySqlSummaryValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStorePasswordMySqlSummaryValueTypeEnum
func GetDatabaseToolsKeyStorePasswordMySqlSummaryValueTypeEnumValues() []DatabaseToolsKeyStorePasswordMySqlSummaryValueTypeEnum {
	values := make([]DatabaseToolsKeyStorePasswordMySqlSummaryValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStorePasswordMySqlSummaryValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStorePasswordMySqlSummaryValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStorePasswordMySqlSummaryValueTypeEnum
func GetDatabaseToolsKeyStorePasswordMySqlSummaryValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStorePasswordMySqlSummaryValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStorePasswordMySqlSummaryValueTypeEnum(val string) (DatabaseToolsKeyStorePasswordMySqlSummaryValueTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsKeyStorePasswordMySqlSummaryValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
