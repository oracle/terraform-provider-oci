// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExecuteSqlOutputDispositionDetails Describes how the response of a command is to be stored
type ExecuteSqlOutputDispositionDetails interface {
}

type executesqloutputdispositiondetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *executesqloutputdispositiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerexecutesqloutputdispositiondetails executesqloutputdispositiondetails
	s := struct {
		Model Unmarshalerexecutesqloutputdispositiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *executesqloutputdispositiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "OBJECT_STORAGE":
		mm := ExecuteSqlOutputDispositionObjectStorageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ExecuteSqlOutputDispositionDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m executesqloutputdispositiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m executesqloutputdispositiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExecuteSqlOutputDispositionDetailsTypeEnum Enum with underlying type: string
type ExecuteSqlOutputDispositionDetailsTypeEnum string

// Set of constants representing the allowable values for ExecuteSqlOutputDispositionDetailsTypeEnum
const (
	ExecuteSqlOutputDispositionDetailsTypeObjectStorage ExecuteSqlOutputDispositionDetailsTypeEnum = "OBJECT_STORAGE"
)

var mappingExecuteSqlOutputDispositionDetailsTypeEnum = map[string]ExecuteSqlOutputDispositionDetailsTypeEnum{
	"OBJECT_STORAGE": ExecuteSqlOutputDispositionDetailsTypeObjectStorage,
}

var mappingExecuteSqlOutputDispositionDetailsTypeEnumLowerCase = map[string]ExecuteSqlOutputDispositionDetailsTypeEnum{
	"object_storage": ExecuteSqlOutputDispositionDetailsTypeObjectStorage,
}

// GetExecuteSqlOutputDispositionDetailsTypeEnumValues Enumerates the set of values for ExecuteSqlOutputDispositionDetailsTypeEnum
func GetExecuteSqlOutputDispositionDetailsTypeEnumValues() []ExecuteSqlOutputDispositionDetailsTypeEnum {
	values := make([]ExecuteSqlOutputDispositionDetailsTypeEnum, 0)
	for _, v := range mappingExecuteSqlOutputDispositionDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExecuteSqlOutputDispositionDetailsTypeEnumStringValues Enumerates the set of values in String for ExecuteSqlOutputDispositionDetailsTypeEnum
func GetExecuteSqlOutputDispositionDetailsTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
	}
}

// GetMappingExecuteSqlOutputDispositionDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecuteSqlOutputDispositionDetailsTypeEnum(val string) (ExecuteSqlOutputDispositionDetailsTypeEnum, bool) {
	enum, ok := mappingExecuteSqlOutputDispositionDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
