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

// CreateMySqlDataTransferMediumDetails Optional additional properties for data transfer.
type CreateMySqlDataTransferMediumDetails interface {
}

type createmysqldatatransfermediumdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createmysqldatatransfermediumdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatemysqldatatransfermediumdetails createmysqldatatransfermediumdetails
	s := struct {
		Model Unmarshalercreatemysqldatatransfermediumdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createmysqldatatransfermediumdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "OBJECT_STORAGE":
		mm := CreateMySqlObjectStorageDataTransferMediumDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateMySqlDataTransferMediumDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m createmysqldatatransfermediumdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createmysqldatatransfermediumdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateMySqlDataTransferMediumDetailsTypeEnum Enum with underlying type: string
type CreateMySqlDataTransferMediumDetailsTypeEnum string

// Set of constants representing the allowable values for CreateMySqlDataTransferMediumDetailsTypeEnum
const (
	CreateMySqlDataTransferMediumDetailsTypeObjectStorage CreateMySqlDataTransferMediumDetailsTypeEnum = "OBJECT_STORAGE"
)

var mappingCreateMySqlDataTransferMediumDetailsTypeEnum = map[string]CreateMySqlDataTransferMediumDetailsTypeEnum{
	"OBJECT_STORAGE": CreateMySqlDataTransferMediumDetailsTypeObjectStorage,
}

var mappingCreateMySqlDataTransferMediumDetailsTypeEnumLowerCase = map[string]CreateMySqlDataTransferMediumDetailsTypeEnum{
	"object_storage": CreateMySqlDataTransferMediumDetailsTypeObjectStorage,
}

// GetCreateMySqlDataTransferMediumDetailsTypeEnumValues Enumerates the set of values for CreateMySqlDataTransferMediumDetailsTypeEnum
func GetCreateMySqlDataTransferMediumDetailsTypeEnumValues() []CreateMySqlDataTransferMediumDetailsTypeEnum {
	values := make([]CreateMySqlDataTransferMediumDetailsTypeEnum, 0)
	for _, v := range mappingCreateMySqlDataTransferMediumDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateMySqlDataTransferMediumDetailsTypeEnumStringValues Enumerates the set of values in String for CreateMySqlDataTransferMediumDetailsTypeEnum
func GetCreateMySqlDataTransferMediumDetailsTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
	}
}

// GetMappingCreateMySqlDataTransferMediumDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateMySqlDataTransferMediumDetailsTypeEnum(val string) (CreateMySqlDataTransferMediumDetailsTypeEnum, bool) {
	enum, ok := mappingCreateMySqlDataTransferMediumDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
