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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DatabaseToolsUserPassword The user password.
type DatabaseToolsUserPassword interface {
}

type databasetoolsuserpassword struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolsuserpassword) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolsuserpassword databasetoolsuserpassword
	s := struct {
		Model Unmarshalerdatabasetoolsuserpassword
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolsuserpassword) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsUserPasswordSecretId{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m databasetoolsuserpassword) String() string {
	return common.PointerString(m)
}

// DatabaseToolsUserPasswordValueTypeEnum Enum with underlying type: string
type DatabaseToolsUserPasswordValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsUserPasswordValueTypeEnum
const (
	DatabaseToolsUserPasswordValueTypeSecretid DatabaseToolsUserPasswordValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsUserPasswordValueType = map[string]DatabaseToolsUserPasswordValueTypeEnum{
	"SECRETID": DatabaseToolsUserPasswordValueTypeSecretid,
}

// GetDatabaseToolsUserPasswordValueTypeEnumValues Enumerates the set of values for DatabaseToolsUserPasswordValueTypeEnum
func GetDatabaseToolsUserPasswordValueTypeEnumValues() []DatabaseToolsUserPasswordValueTypeEnum {
	values := make([]DatabaseToolsUserPasswordValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsUserPasswordValueType {
		values = append(values, v)
	}
	return values
}
