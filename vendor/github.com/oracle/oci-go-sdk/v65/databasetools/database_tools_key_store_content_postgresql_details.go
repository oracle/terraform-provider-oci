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

// DatabaseToolsKeyStoreContentPostgresqlDetails The key store content.
type DatabaseToolsKeyStoreContentPostgresqlDetails interface {
}

type databasetoolskeystorecontentpostgresqldetails struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorecontentpostgresqldetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorecontentpostgresqldetails databasetoolskeystorecontentpostgresqldetails
	s := struct {
		Model Unmarshalerdatabasetoolskeystorecontentpostgresqldetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorecontentpostgresqldetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStoreContentSecretIdPostgresqlDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DatabaseToolsKeyStoreContentPostgresqlDetails: %s.", m.ValueType)
		return *m, nil
	}
}

func (m databasetoolskeystorecontentpostgresqldetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorecontentpostgresqldetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeEnum
const (
	DatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeSecretid DatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeEnum = map[string]DatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeSecretid,
}

var mappingDatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeEnumLowerCase = map[string]DatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeEnum{
	"secretid": DatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeSecretid,
}

// GetDatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeEnum
func GetDatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeEnumValues() []DatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeEnum {
	values := make([]DatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeEnum
func GetDatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeEnum(val string) (DatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsKeyStoreContentPostgresqlDetailsValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
