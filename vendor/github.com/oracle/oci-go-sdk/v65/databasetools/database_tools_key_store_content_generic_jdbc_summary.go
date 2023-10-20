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

// DatabaseToolsKeyStoreContentGenericJdbcSummary The key store content.
type DatabaseToolsKeyStoreContentGenericJdbcSummary interface {
}

type databasetoolskeystorecontentgenericjdbcsummary struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorecontentgenericjdbcsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorecontentgenericjdbcsummary databasetoolskeystorecontentgenericjdbcsummary
	s := struct {
		Model Unmarshalerdatabasetoolskeystorecontentgenericjdbcsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorecontentgenericjdbcsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStoreContentSecretIdGenericJdbcSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DatabaseToolsKeyStoreContentGenericJdbcSummary: %s.", m.ValueType)
		return *m, nil
	}
}

func (m databasetoolskeystorecontentgenericjdbcsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorecontentgenericjdbcsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeEnum
const (
	DatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeSecretid DatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeEnum = map[string]DatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeSecretid,
}

var mappingDatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeEnumLowerCase = map[string]DatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeEnum{
	"secretid": DatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeSecretid,
}

// GetDatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeEnum
func GetDatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeEnumValues() []DatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeEnum {
	values := make([]DatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeEnum
func GetDatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeEnum(val string) (DatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsKeyStoreContentGenericJdbcSummaryValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
