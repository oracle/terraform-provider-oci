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

// DatabaseToolsKeyStoreContentGenericJdbcDetails The key store content.
type DatabaseToolsKeyStoreContentGenericJdbcDetails interface {
}

type databasetoolskeystorecontentgenericjdbcdetails struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorecontentgenericjdbcdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorecontentgenericjdbcdetails databasetoolskeystorecontentgenericjdbcdetails
	s := struct {
		Model Unmarshalerdatabasetoolskeystorecontentgenericjdbcdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorecontentgenericjdbcdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStoreContentSecretIdGenericJdbcDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DatabaseToolsKeyStoreContentGenericJdbcDetails: %s.", m.ValueType)
		return *m, nil
	}
}

func (m databasetoolskeystorecontentgenericjdbcdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorecontentgenericjdbcdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeEnum
const (
	DatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeSecretid DatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeEnum = map[string]DatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeSecretid,
}

var mappingDatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeEnumLowerCase = map[string]DatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeEnum{
	"secretid": DatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeSecretid,
}

// GetDatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeEnum
func GetDatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeEnumValues() []DatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeEnum {
	values := make([]DatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeEnum
func GetDatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeEnum(val string) (DatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsKeyStoreContentGenericJdbcDetailsValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
