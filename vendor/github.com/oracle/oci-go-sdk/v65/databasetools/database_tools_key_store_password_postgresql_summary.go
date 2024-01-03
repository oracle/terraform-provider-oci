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

// DatabaseToolsKeyStorePasswordPostgresqlSummary The key store password.
type DatabaseToolsKeyStorePasswordPostgresqlSummary interface {
}

type databasetoolskeystorepasswordpostgresqlsummary struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorepasswordpostgresqlsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorepasswordpostgresqlsummary databasetoolskeystorepasswordpostgresqlsummary
	s := struct {
		Model Unmarshalerdatabasetoolskeystorepasswordpostgresqlsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorepasswordpostgresqlsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStorePasswordSecretIdPostgresqlSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DatabaseToolsKeyStorePasswordPostgresqlSummary: %s.", m.ValueType)
		return *m, nil
	}
}

func (m databasetoolskeystorepasswordpostgresqlsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorepasswordpostgresqlsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeEnum
const (
	DatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeSecretid DatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeEnum = map[string]DatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeSecretid,
}

var mappingDatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeEnumLowerCase = map[string]DatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeEnum{
	"secretid": DatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeSecretid,
}

// GetDatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeEnum
func GetDatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeEnumValues() []DatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeEnum {
	values := make([]DatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeEnum
func GetDatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeEnum(val string) (DatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsKeyStorePasswordPostgresqlSummaryValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
