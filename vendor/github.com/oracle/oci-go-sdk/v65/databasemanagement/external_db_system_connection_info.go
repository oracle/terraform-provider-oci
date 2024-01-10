// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalDbSystemConnectionInfo The connection details required to connect to an external DB system component.
type ExternalDbSystemConnectionInfo interface {
}

type externaldbsystemconnectioninfo struct {
	JsonData      []byte
	ComponentType string `json:"componentType"`
}

// UnmarshalJSON unmarshals json
func (m *externaldbsystemconnectioninfo) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerexternaldbsystemconnectioninfo externaldbsystemconnectioninfo
	s := struct {
		Model Unmarshalerexternaldbsystemconnectioninfo
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ComponentType = s.Model.ComponentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *externaldbsystemconnectioninfo) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ComponentType {
	case "ASM":
		mm := ExternalAsmConnectionInfo{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATABASE":
		mm := ExternalDatabaseConnectionInfo{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ExternalDbSystemConnectionInfo: %s.", m.ComponentType)
		return *m, nil
	}
}

func (m externaldbsystemconnectioninfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m externaldbsystemconnectioninfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalDbSystemConnectionInfoComponentTypeEnum Enum with underlying type: string
type ExternalDbSystemConnectionInfoComponentTypeEnum string

// Set of constants representing the allowable values for ExternalDbSystemConnectionInfoComponentTypeEnum
const (
	ExternalDbSystemConnectionInfoComponentTypeDatabase ExternalDbSystemConnectionInfoComponentTypeEnum = "DATABASE"
	ExternalDbSystemConnectionInfoComponentTypeAsm      ExternalDbSystemConnectionInfoComponentTypeEnum = "ASM"
)

var mappingExternalDbSystemConnectionInfoComponentTypeEnum = map[string]ExternalDbSystemConnectionInfoComponentTypeEnum{
	"DATABASE": ExternalDbSystemConnectionInfoComponentTypeDatabase,
	"ASM":      ExternalDbSystemConnectionInfoComponentTypeAsm,
}

var mappingExternalDbSystemConnectionInfoComponentTypeEnumLowerCase = map[string]ExternalDbSystemConnectionInfoComponentTypeEnum{
	"database": ExternalDbSystemConnectionInfoComponentTypeDatabase,
	"asm":      ExternalDbSystemConnectionInfoComponentTypeAsm,
}

// GetExternalDbSystemConnectionInfoComponentTypeEnumValues Enumerates the set of values for ExternalDbSystemConnectionInfoComponentTypeEnum
func GetExternalDbSystemConnectionInfoComponentTypeEnumValues() []ExternalDbSystemConnectionInfoComponentTypeEnum {
	values := make([]ExternalDbSystemConnectionInfoComponentTypeEnum, 0)
	for _, v := range mappingExternalDbSystemConnectionInfoComponentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalDbSystemConnectionInfoComponentTypeEnumStringValues Enumerates the set of values in String for ExternalDbSystemConnectionInfoComponentTypeEnum
func GetExternalDbSystemConnectionInfoComponentTypeEnumStringValues() []string {
	return []string{
		"DATABASE",
		"ASM",
	}
}

// GetMappingExternalDbSystemConnectionInfoComponentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalDbSystemConnectionInfoComponentTypeEnum(val string) (ExternalDbSystemConnectionInfoComponentTypeEnum, bool) {
	enum, ok := mappingExternalDbSystemConnectionInfoComponentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
