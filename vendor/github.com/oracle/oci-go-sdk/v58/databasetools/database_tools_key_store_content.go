// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Database Tools APIs to manage Connections and Private Endpoints.
//

package databasetools

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// DatabaseToolsKeyStoreContent The key store content.
type DatabaseToolsKeyStoreContent interface {
}

type databasetoolskeystorecontent struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorecontent) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorecontent databasetoolskeystorecontent
	s := struct {
		Model Unmarshalerdatabasetoolskeystorecontent
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorecontent) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStoreContentSecretId{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m databasetoolskeystorecontent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorecontent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStoreContentValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStoreContentValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStoreContentValueTypeEnum
const (
	DatabaseToolsKeyStoreContentValueTypeSecretid DatabaseToolsKeyStoreContentValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStoreContentValueTypeEnum = map[string]DatabaseToolsKeyStoreContentValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStoreContentValueTypeSecretid,
}

// GetDatabaseToolsKeyStoreContentValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStoreContentValueTypeEnum
func GetDatabaseToolsKeyStoreContentValueTypeEnumValues() []DatabaseToolsKeyStoreContentValueTypeEnum {
	values := make([]DatabaseToolsKeyStoreContentValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStoreContentValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStoreContentValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStoreContentValueTypeEnum
func GetDatabaseToolsKeyStoreContentValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStoreContentValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStoreContentValueTypeEnum(val string) (DatabaseToolsKeyStoreContentValueTypeEnum, bool) {
	mappingDatabaseToolsKeyStoreContentValueTypeEnumIgnoreCase := make(map[string]DatabaseToolsKeyStoreContentValueTypeEnum)
	for k, v := range mappingDatabaseToolsKeyStoreContentValueTypeEnum {
		mappingDatabaseToolsKeyStoreContentValueTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDatabaseToolsKeyStoreContentValueTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
