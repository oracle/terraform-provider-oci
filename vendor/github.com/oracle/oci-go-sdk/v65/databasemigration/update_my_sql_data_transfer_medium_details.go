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

// UpdateMySqlDataTransferMediumDetails Optional additional properties for data transfer.
type UpdateMySqlDataTransferMediumDetails interface {
}

type updatemysqldatatransfermediumdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *updatemysqldatatransfermediumdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatemysqldatatransfermediumdetails updatemysqldatatransfermediumdetails
	s := struct {
		Model Unmarshalerupdatemysqldatatransfermediumdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatemysqldatatransfermediumdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "OBJECT_STORAGE":
		mm := UpdateMySqlObjectStorageDataTransferMediumDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdateMySqlDataTransferMediumDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m updatemysqldatatransfermediumdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatemysqldatatransfermediumdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateMySqlDataTransferMediumDetailsTypeEnum Enum with underlying type: string
type UpdateMySqlDataTransferMediumDetailsTypeEnum string

// Set of constants representing the allowable values for UpdateMySqlDataTransferMediumDetailsTypeEnum
const (
	UpdateMySqlDataTransferMediumDetailsTypeObjectStorage UpdateMySqlDataTransferMediumDetailsTypeEnum = "OBJECT_STORAGE"
)

var mappingUpdateMySqlDataTransferMediumDetailsTypeEnum = map[string]UpdateMySqlDataTransferMediumDetailsTypeEnum{
	"OBJECT_STORAGE": UpdateMySqlDataTransferMediumDetailsTypeObjectStorage,
}

var mappingUpdateMySqlDataTransferMediumDetailsTypeEnumLowerCase = map[string]UpdateMySqlDataTransferMediumDetailsTypeEnum{
	"object_storage": UpdateMySqlDataTransferMediumDetailsTypeObjectStorage,
}

// GetUpdateMySqlDataTransferMediumDetailsTypeEnumValues Enumerates the set of values for UpdateMySqlDataTransferMediumDetailsTypeEnum
func GetUpdateMySqlDataTransferMediumDetailsTypeEnumValues() []UpdateMySqlDataTransferMediumDetailsTypeEnum {
	values := make([]UpdateMySqlDataTransferMediumDetailsTypeEnum, 0)
	for _, v := range mappingUpdateMySqlDataTransferMediumDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateMySqlDataTransferMediumDetailsTypeEnumStringValues Enumerates the set of values in String for UpdateMySqlDataTransferMediumDetailsTypeEnum
func GetUpdateMySqlDataTransferMediumDetailsTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
	}
}

// GetMappingUpdateMySqlDataTransferMediumDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateMySqlDataTransferMediumDetailsTypeEnum(val string) (UpdateMySqlDataTransferMediumDetailsTypeEnum, bool) {
	enum, ok := mappingUpdateMySqlDataTransferMediumDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
