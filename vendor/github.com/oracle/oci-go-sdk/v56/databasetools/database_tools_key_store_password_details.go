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

// DatabaseToolsKeyStorePasswordDetails The key store password.
type DatabaseToolsKeyStorePasswordDetails interface {
}

type databasetoolskeystorepassworddetails struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorepassworddetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorepassworddetails databasetoolskeystorepassworddetails
	s := struct {
		Model Unmarshalerdatabasetoolskeystorepassworddetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorepassworddetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStorePasswordSecretIdDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m databasetoolskeystorepassworddetails) String() string {
	return common.PointerString(m)
}

// DatabaseToolsKeyStorePasswordDetailsValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStorePasswordDetailsValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStorePasswordDetailsValueTypeEnum
const (
	DatabaseToolsKeyStorePasswordDetailsValueTypeSecretid DatabaseToolsKeyStorePasswordDetailsValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStorePasswordDetailsValueType = map[string]DatabaseToolsKeyStorePasswordDetailsValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStorePasswordDetailsValueTypeSecretid,
}

// GetDatabaseToolsKeyStorePasswordDetailsValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStorePasswordDetailsValueTypeEnum
func GetDatabaseToolsKeyStorePasswordDetailsValueTypeEnumValues() []DatabaseToolsKeyStorePasswordDetailsValueTypeEnum {
	values := make([]DatabaseToolsKeyStorePasswordDetailsValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStorePasswordDetailsValueType {
		values = append(values, v)
	}
	return values
}
