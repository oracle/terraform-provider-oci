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

// DatabaseToolsKeyStorePasswordGenericJdbcSummary The key store password.
type DatabaseToolsKeyStorePasswordGenericJdbcSummary interface {
}

type databasetoolskeystorepasswordgenericjdbcsummary struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorepasswordgenericjdbcsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorepasswordgenericjdbcsummary databasetoolskeystorepasswordgenericjdbcsummary
	s := struct {
		Model Unmarshalerdatabasetoolskeystorepasswordgenericjdbcsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorepasswordgenericjdbcsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStorePasswordSecretIdGenericJdbcSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DatabaseToolsKeyStorePasswordGenericJdbcSummary: %s.", m.ValueType)
		return *m, nil
	}
}

func (m databasetoolskeystorepasswordgenericjdbcsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorepasswordgenericjdbcsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeEnum
const (
	DatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeSecretid DatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeEnum = map[string]DatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeSecretid,
}

var mappingDatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeEnumLowerCase = map[string]DatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeEnum{
	"secretid": DatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeSecretid,
}

// GetDatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeEnum
func GetDatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeEnumValues() []DatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeEnum {
	values := make([]DatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeEnum
func GetDatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeEnum(val string) (DatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsKeyStorePasswordGenericJdbcSummaryValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
