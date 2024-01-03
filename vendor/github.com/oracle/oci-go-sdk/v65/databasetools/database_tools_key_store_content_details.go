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

// DatabaseToolsKeyStoreContentDetails The key store content.
type DatabaseToolsKeyStoreContentDetails interface {
}

type databasetoolskeystorecontentdetails struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorecontentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorecontentdetails databasetoolskeystorecontentdetails
	s := struct {
		Model Unmarshalerdatabasetoolskeystorecontentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorecontentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStoreContentSecretIdDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DatabaseToolsKeyStoreContentDetails: %s.", m.ValueType)
		return *m, nil
	}
}

func (m databasetoolskeystorecontentdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorecontentdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStoreContentDetailsValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStoreContentDetailsValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStoreContentDetailsValueTypeEnum
const (
	DatabaseToolsKeyStoreContentDetailsValueTypeSecretid DatabaseToolsKeyStoreContentDetailsValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStoreContentDetailsValueTypeEnum = map[string]DatabaseToolsKeyStoreContentDetailsValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStoreContentDetailsValueTypeSecretid,
}

var mappingDatabaseToolsKeyStoreContentDetailsValueTypeEnumLowerCase = map[string]DatabaseToolsKeyStoreContentDetailsValueTypeEnum{
	"secretid": DatabaseToolsKeyStoreContentDetailsValueTypeSecretid,
}

// GetDatabaseToolsKeyStoreContentDetailsValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStoreContentDetailsValueTypeEnum
func GetDatabaseToolsKeyStoreContentDetailsValueTypeEnumValues() []DatabaseToolsKeyStoreContentDetailsValueTypeEnum {
	values := make([]DatabaseToolsKeyStoreContentDetailsValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStoreContentDetailsValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStoreContentDetailsValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStoreContentDetailsValueTypeEnum
func GetDatabaseToolsKeyStoreContentDetailsValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStoreContentDetailsValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStoreContentDetailsValueTypeEnum(val string) (DatabaseToolsKeyStoreContentDetailsValueTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsKeyStoreContentDetailsValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
