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

// ExecuteSqlOutputDetails Defines how the script result should be stored
type ExecuteSqlOutputDetails interface {

	// Defines how the result of commands in a script should be stored.
	// If the command does not match any template filter, the result will be inline.
	GetResultDispositionTemplates() []ExecuteSqlOutputResultDispositionTemplate
}

type executesqloutputdetails struct {
	JsonData                   []byte
	ResultDispositionTemplates []ExecuteSqlOutputResultDispositionTemplate `mandatory:"false" json:"resultDispositionTemplates"`
	Type                       string                                      `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *executesqloutputdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerexecutesqloutputdetails executesqloutputdetails
	s := struct {
		Model Unmarshalerexecutesqloutputdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ResultDispositionTemplates = s.Model.ResultDispositionTemplates
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *executesqloutputdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "OBJECT_STORAGE":
		mm := ExecuteSqlOutputObjectStorageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ExecuteSqlOutputDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetResultDispositionTemplates returns ResultDispositionTemplates
func (m executesqloutputdetails) GetResultDispositionTemplates() []ExecuteSqlOutputResultDispositionTemplate {
	return m.ResultDispositionTemplates
}

func (m executesqloutputdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m executesqloutputdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExecuteSqlOutputDetailsTypeEnum Enum with underlying type: string
type ExecuteSqlOutputDetailsTypeEnum string

// Set of constants representing the allowable values for ExecuteSqlOutputDetailsTypeEnum
const (
	ExecuteSqlOutputDetailsTypeObjectStorage ExecuteSqlOutputDetailsTypeEnum = "OBJECT_STORAGE"
)

var mappingExecuteSqlOutputDetailsTypeEnum = map[string]ExecuteSqlOutputDetailsTypeEnum{
	"OBJECT_STORAGE": ExecuteSqlOutputDetailsTypeObjectStorage,
}

var mappingExecuteSqlOutputDetailsTypeEnumLowerCase = map[string]ExecuteSqlOutputDetailsTypeEnum{
	"object_storage": ExecuteSqlOutputDetailsTypeObjectStorage,
}

// GetExecuteSqlOutputDetailsTypeEnumValues Enumerates the set of values for ExecuteSqlOutputDetailsTypeEnum
func GetExecuteSqlOutputDetailsTypeEnumValues() []ExecuteSqlOutputDetailsTypeEnum {
	values := make([]ExecuteSqlOutputDetailsTypeEnum, 0)
	for _, v := range mappingExecuteSqlOutputDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExecuteSqlOutputDetailsTypeEnumStringValues Enumerates the set of values in String for ExecuteSqlOutputDetailsTypeEnum
func GetExecuteSqlOutputDetailsTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
	}
}

// GetMappingExecuteSqlOutputDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecuteSqlOutputDetailsTypeEnum(val string) (ExecuteSqlOutputDetailsTypeEnum, bool) {
	enum, ok := mappingExecuteSqlOutputDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
