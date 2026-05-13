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

// ExecuteSqlDatabaseToolsConnectionDetails A request
type ExecuteSqlDatabaseToolsConnectionDetails interface {
	GetOutput() ExecuteSqlOutputDetails
}

type executesqldatabasetoolsconnectiondetails struct {
	JsonData []byte
	Output   executesqloutputdetails `mandatory:"false" json:"output"`
	Type     string                  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *executesqldatabasetoolsconnectiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerexecutesqldatabasetoolsconnectiondetails executesqldatabasetoolsconnectiondetails
	s := struct {
		Model Unmarshalerexecutesqldatabasetoolsconnectiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Output = s.Model.Output
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *executesqldatabasetoolsconnectiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "ASYNCHRONOUS":
		mm := ExecuteSqlDatabaseToolsConnectionAsynchronousDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SYNCHRONOUS":
		mm := ExecuteSqlDatabaseToolsConnectionSynchronousDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ExecuteSqlDatabaseToolsConnectionDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetOutput returns Output
func (m executesqldatabasetoolsconnectiondetails) GetOutput() executesqloutputdetails {
	return m.Output
}

func (m executesqldatabasetoolsconnectiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m executesqldatabasetoolsconnectiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExecuteSqlDatabaseToolsConnectionDetailsTypeEnum Enum with underlying type: string
type ExecuteSqlDatabaseToolsConnectionDetailsTypeEnum string

// Set of constants representing the allowable values for ExecuteSqlDatabaseToolsConnectionDetailsTypeEnum
const (
	ExecuteSqlDatabaseToolsConnectionDetailsTypeSynchronous  ExecuteSqlDatabaseToolsConnectionDetailsTypeEnum = "SYNCHRONOUS"
	ExecuteSqlDatabaseToolsConnectionDetailsTypeAsynchronous ExecuteSqlDatabaseToolsConnectionDetailsTypeEnum = "ASYNCHRONOUS"
)

var mappingExecuteSqlDatabaseToolsConnectionDetailsTypeEnum = map[string]ExecuteSqlDatabaseToolsConnectionDetailsTypeEnum{
	"SYNCHRONOUS":  ExecuteSqlDatabaseToolsConnectionDetailsTypeSynchronous,
	"ASYNCHRONOUS": ExecuteSqlDatabaseToolsConnectionDetailsTypeAsynchronous,
}

var mappingExecuteSqlDatabaseToolsConnectionDetailsTypeEnumLowerCase = map[string]ExecuteSqlDatabaseToolsConnectionDetailsTypeEnum{
	"synchronous":  ExecuteSqlDatabaseToolsConnectionDetailsTypeSynchronous,
	"asynchronous": ExecuteSqlDatabaseToolsConnectionDetailsTypeAsynchronous,
}

// GetExecuteSqlDatabaseToolsConnectionDetailsTypeEnumValues Enumerates the set of values for ExecuteSqlDatabaseToolsConnectionDetailsTypeEnum
func GetExecuteSqlDatabaseToolsConnectionDetailsTypeEnumValues() []ExecuteSqlDatabaseToolsConnectionDetailsTypeEnum {
	values := make([]ExecuteSqlDatabaseToolsConnectionDetailsTypeEnum, 0)
	for _, v := range mappingExecuteSqlDatabaseToolsConnectionDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExecuteSqlDatabaseToolsConnectionDetailsTypeEnumStringValues Enumerates the set of values in String for ExecuteSqlDatabaseToolsConnectionDetailsTypeEnum
func GetExecuteSqlDatabaseToolsConnectionDetailsTypeEnumStringValues() []string {
	return []string{
		"SYNCHRONOUS",
		"ASYNCHRONOUS",
	}
}

// GetMappingExecuteSqlDatabaseToolsConnectionDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecuteSqlDatabaseToolsConnectionDetailsTypeEnum(val string) (ExecuteSqlDatabaseToolsConnectionDetailsTypeEnum, bool) {
	enum, ok := mappingExecuteSqlDatabaseToolsConnectionDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
