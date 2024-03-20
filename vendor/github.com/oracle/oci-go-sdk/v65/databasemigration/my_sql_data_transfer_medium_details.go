// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MySqlDataTransferMediumDetails Optional additional properties for data transfer.
type MySqlDataTransferMediumDetails interface {
}

type mysqldatatransfermediumdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *mysqldatatransfermediumdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermysqldatatransfermediumdetails mysqldatatransfermediumdetails
	s := struct {
		Model Unmarshalermysqldatatransfermediumdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *mysqldatatransfermediumdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "OBJECT_STORAGE":
		mm := MySqlObjectStorageDataTransferMediumDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for MySqlDataTransferMediumDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m mysqldatatransfermediumdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m mysqldatatransfermediumdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MySqlDataTransferMediumDetailsTypeEnum Enum with underlying type: string
type MySqlDataTransferMediumDetailsTypeEnum string

// Set of constants representing the allowable values for MySqlDataTransferMediumDetailsTypeEnum
const (
	MySqlDataTransferMediumDetailsTypeObjectStorage MySqlDataTransferMediumDetailsTypeEnum = "OBJECT_STORAGE"
)

var mappingMySqlDataTransferMediumDetailsTypeEnum = map[string]MySqlDataTransferMediumDetailsTypeEnum{
	"OBJECT_STORAGE": MySqlDataTransferMediumDetailsTypeObjectStorage,
}

var mappingMySqlDataTransferMediumDetailsTypeEnumLowerCase = map[string]MySqlDataTransferMediumDetailsTypeEnum{
	"object_storage": MySqlDataTransferMediumDetailsTypeObjectStorage,
}

// GetMySqlDataTransferMediumDetailsTypeEnumValues Enumerates the set of values for MySqlDataTransferMediumDetailsTypeEnum
func GetMySqlDataTransferMediumDetailsTypeEnumValues() []MySqlDataTransferMediumDetailsTypeEnum {
	values := make([]MySqlDataTransferMediumDetailsTypeEnum, 0)
	for _, v := range mappingMySqlDataTransferMediumDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMySqlDataTransferMediumDetailsTypeEnumStringValues Enumerates the set of values in String for MySqlDataTransferMediumDetailsTypeEnum
func GetMySqlDataTransferMediumDetailsTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
	}
}

// GetMappingMySqlDataTransferMediumDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMySqlDataTransferMediumDetailsTypeEnum(val string) (MySqlDataTransferMediumDetailsTypeEnum, bool) {
	enum, ok := mappingMySqlDataTransferMediumDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
