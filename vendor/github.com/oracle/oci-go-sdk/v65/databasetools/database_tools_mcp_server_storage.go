// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools API
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

// DatabaseToolsMcpServerStorage The storage option used when running a tool asynchronously.
type DatabaseToolsMcpServerStorage interface {
}

type databasetoolsmcpserverstorage struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolsmcpserverstorage) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolsmcpserverstorage databasetoolsmcpserverstorage
	s := struct {
		Model Unmarshalerdatabasetoolsmcpserverstorage
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolsmcpserverstorage) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "OBJECT_STORAGE":
		mm := DatabaseToolsMcpServerStorageObjectStorage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NONE":
		mm := DatabaseToolsMcpServerStorageNone{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DatabaseToolsMcpServerStorage: %s.", m.Type)
		return *m, nil
	}
}

func (m databasetoolsmcpserverstorage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolsmcpserverstorage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsMcpServerStorageTypeEnum Enum with underlying type: string
type DatabaseToolsMcpServerStorageTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsMcpServerStorageTypeEnum
const (
	DatabaseToolsMcpServerStorageTypeNone          DatabaseToolsMcpServerStorageTypeEnum = "NONE"
	DatabaseToolsMcpServerStorageTypeObjectStorage DatabaseToolsMcpServerStorageTypeEnum = "OBJECT_STORAGE"
)

var mappingDatabaseToolsMcpServerStorageTypeEnum = map[string]DatabaseToolsMcpServerStorageTypeEnum{
	"NONE":           DatabaseToolsMcpServerStorageTypeNone,
	"OBJECT_STORAGE": DatabaseToolsMcpServerStorageTypeObjectStorage,
}

var mappingDatabaseToolsMcpServerStorageTypeEnumLowerCase = map[string]DatabaseToolsMcpServerStorageTypeEnum{
	"none":           DatabaseToolsMcpServerStorageTypeNone,
	"object_storage": DatabaseToolsMcpServerStorageTypeObjectStorage,
}

// GetDatabaseToolsMcpServerStorageTypeEnumValues Enumerates the set of values for DatabaseToolsMcpServerStorageTypeEnum
func GetDatabaseToolsMcpServerStorageTypeEnumValues() []DatabaseToolsMcpServerStorageTypeEnum {
	values := make([]DatabaseToolsMcpServerStorageTypeEnum, 0)
	for _, v := range mappingDatabaseToolsMcpServerStorageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsMcpServerStorageTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsMcpServerStorageTypeEnum
func GetDatabaseToolsMcpServerStorageTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"OBJECT_STORAGE",
	}
}

// GetMappingDatabaseToolsMcpServerStorageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsMcpServerStorageTypeEnum(val string) (DatabaseToolsMcpServerStorageTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsMcpServerStorageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
