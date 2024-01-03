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

// DatabaseToolsKeyStorePasswordPostgresqlDetails The key store password.
type DatabaseToolsKeyStorePasswordPostgresqlDetails interface {
}

type databasetoolskeystorepasswordpostgresqldetails struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorepasswordpostgresqldetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorepasswordpostgresqldetails databasetoolskeystorepasswordpostgresqldetails
	s := struct {
		Model Unmarshalerdatabasetoolskeystorepasswordpostgresqldetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorepasswordpostgresqldetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStorePasswordSecretIdPostgresqlDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DatabaseToolsKeyStorePasswordPostgresqlDetails: %s.", m.ValueType)
		return *m, nil
	}
}

func (m databasetoolskeystorepasswordpostgresqldetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorepasswordpostgresqldetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeEnum
const (
	DatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeSecretid DatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeEnum = map[string]DatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeSecretid,
}

var mappingDatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeEnumLowerCase = map[string]DatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeEnum{
	"secretid": DatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeSecretid,
}

// GetDatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeEnum
func GetDatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeEnumValues() []DatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeEnum {
	values := make([]DatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeEnum
func GetDatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeEnum(val string) (DatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsKeyStorePasswordPostgresqlDetailsValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
