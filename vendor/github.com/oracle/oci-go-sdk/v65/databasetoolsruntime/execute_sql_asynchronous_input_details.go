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

// ExecuteSqlAsynchronousInputDetails Async request script input details
type ExecuteSqlAsynchronousInputDetails interface {
}

type executesqlasynchronousinputdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *executesqlasynchronousinputdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerexecutesqlasynchronousinputdetails executesqlasynchronousinputdetails
	s := struct {
		Model Unmarshalerexecutesqlasynchronousinputdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *executesqlasynchronousinputdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "OBJECT_STORAGE":
		mm := ExecuteSqlAsynchronousInputObjectStorageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INLINE":
		mm := ExecuteSqlAsynchronousInputInlineDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ExecuteSqlAsynchronousInputDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m executesqlasynchronousinputdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m executesqlasynchronousinputdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExecuteSqlAsynchronousInputDetailsTypeEnum Enum with underlying type: string
type ExecuteSqlAsynchronousInputDetailsTypeEnum string

// Set of constants representing the allowable values for ExecuteSqlAsynchronousInputDetailsTypeEnum
const (
	ExecuteSqlAsynchronousInputDetailsTypeObjectStorage ExecuteSqlAsynchronousInputDetailsTypeEnum = "OBJECT_STORAGE"
	ExecuteSqlAsynchronousInputDetailsTypeInline        ExecuteSqlAsynchronousInputDetailsTypeEnum = "INLINE"
)

var mappingExecuteSqlAsynchronousInputDetailsTypeEnum = map[string]ExecuteSqlAsynchronousInputDetailsTypeEnum{
	"OBJECT_STORAGE": ExecuteSqlAsynchronousInputDetailsTypeObjectStorage,
	"INLINE":         ExecuteSqlAsynchronousInputDetailsTypeInline,
}

var mappingExecuteSqlAsynchronousInputDetailsTypeEnumLowerCase = map[string]ExecuteSqlAsynchronousInputDetailsTypeEnum{
	"object_storage": ExecuteSqlAsynchronousInputDetailsTypeObjectStorage,
	"inline":         ExecuteSqlAsynchronousInputDetailsTypeInline,
}

// GetExecuteSqlAsynchronousInputDetailsTypeEnumValues Enumerates the set of values for ExecuteSqlAsynchronousInputDetailsTypeEnum
func GetExecuteSqlAsynchronousInputDetailsTypeEnumValues() []ExecuteSqlAsynchronousInputDetailsTypeEnum {
	values := make([]ExecuteSqlAsynchronousInputDetailsTypeEnum, 0)
	for _, v := range mappingExecuteSqlAsynchronousInputDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExecuteSqlAsynchronousInputDetailsTypeEnumStringValues Enumerates the set of values in String for ExecuteSqlAsynchronousInputDetailsTypeEnum
func GetExecuteSqlAsynchronousInputDetailsTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
		"INLINE",
	}
}

// GetMappingExecuteSqlAsynchronousInputDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecuteSqlAsynchronousInputDetailsTypeEnum(val string) (ExecuteSqlAsynchronousInputDetailsTypeEnum, bool) {
	enum, ok := mappingExecuteSqlAsynchronousInputDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
