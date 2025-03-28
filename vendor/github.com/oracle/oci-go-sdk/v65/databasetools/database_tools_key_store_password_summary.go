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

// DatabaseToolsKeyStorePasswordSummary The key store password.
type DatabaseToolsKeyStorePasswordSummary interface {
}

type databasetoolskeystorepasswordsummary struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorepasswordsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorepasswordsummary databasetoolskeystorepasswordsummary
	s := struct {
		Model Unmarshalerdatabasetoolskeystorepasswordsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorepasswordsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStorePasswordSecretIdSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DatabaseToolsKeyStorePasswordSummary: %s.", m.ValueType)
		return *m, nil
	}
}

func (m databasetoolskeystorepasswordsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorepasswordsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStorePasswordSummaryValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStorePasswordSummaryValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStorePasswordSummaryValueTypeEnum
const (
	DatabaseToolsKeyStorePasswordSummaryValueTypeSecretid DatabaseToolsKeyStorePasswordSummaryValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStorePasswordSummaryValueTypeEnum = map[string]DatabaseToolsKeyStorePasswordSummaryValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStorePasswordSummaryValueTypeSecretid,
}

var mappingDatabaseToolsKeyStorePasswordSummaryValueTypeEnumLowerCase = map[string]DatabaseToolsKeyStorePasswordSummaryValueTypeEnum{
	"secretid": DatabaseToolsKeyStorePasswordSummaryValueTypeSecretid,
}

// GetDatabaseToolsKeyStorePasswordSummaryValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStorePasswordSummaryValueTypeEnum
func GetDatabaseToolsKeyStorePasswordSummaryValueTypeEnumValues() []DatabaseToolsKeyStorePasswordSummaryValueTypeEnum {
	values := make([]DatabaseToolsKeyStorePasswordSummaryValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStorePasswordSummaryValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStorePasswordSummaryValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStorePasswordSummaryValueTypeEnum
func GetDatabaseToolsKeyStorePasswordSummaryValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStorePasswordSummaryValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStorePasswordSummaryValueTypeEnum(val string) (DatabaseToolsKeyStorePasswordSummaryValueTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsKeyStorePasswordSummaryValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
