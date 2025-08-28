// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CloudDbSystemConnectionInfo The connection details required to connect to a cloud DB system component.
type CloudDbSystemConnectionInfo interface {
}

type clouddbsystemconnectioninfo struct {
	JsonData      []byte
	ComponentType string `json:"componentType"`
}

// UnmarshalJSON unmarshals json
func (m *clouddbsystemconnectioninfo) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerclouddbsystemconnectioninfo clouddbsystemconnectioninfo
	s := struct {
		Model Unmarshalerclouddbsystemconnectioninfo
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ComponentType = s.Model.ComponentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *clouddbsystemconnectioninfo) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ComponentType {
	case "ASM":
		mm := CloudAsmConnectionInfo{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATABASE":
		mm := CloudDatabaseConnectionInfo{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CloudDbSystemConnectionInfo: %s.", m.ComponentType)
		return *m, nil
	}
}

func (m clouddbsystemconnectioninfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m clouddbsystemconnectioninfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CloudDbSystemConnectionInfoComponentTypeEnum Enum with underlying type: string
type CloudDbSystemConnectionInfoComponentTypeEnum string

// Set of constants representing the allowable values for CloudDbSystemConnectionInfoComponentTypeEnum
const (
	CloudDbSystemConnectionInfoComponentTypeDatabase CloudDbSystemConnectionInfoComponentTypeEnum = "DATABASE"
	CloudDbSystemConnectionInfoComponentTypeAsm      CloudDbSystemConnectionInfoComponentTypeEnum = "ASM"
)

var mappingCloudDbSystemConnectionInfoComponentTypeEnum = map[string]CloudDbSystemConnectionInfoComponentTypeEnum{
	"DATABASE": CloudDbSystemConnectionInfoComponentTypeDatabase,
	"ASM":      CloudDbSystemConnectionInfoComponentTypeAsm,
}

var mappingCloudDbSystemConnectionInfoComponentTypeEnumLowerCase = map[string]CloudDbSystemConnectionInfoComponentTypeEnum{
	"database": CloudDbSystemConnectionInfoComponentTypeDatabase,
	"asm":      CloudDbSystemConnectionInfoComponentTypeAsm,
}

// GetCloudDbSystemConnectionInfoComponentTypeEnumValues Enumerates the set of values for CloudDbSystemConnectionInfoComponentTypeEnum
func GetCloudDbSystemConnectionInfoComponentTypeEnumValues() []CloudDbSystemConnectionInfoComponentTypeEnum {
	values := make([]CloudDbSystemConnectionInfoComponentTypeEnum, 0)
	for _, v := range mappingCloudDbSystemConnectionInfoComponentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudDbSystemConnectionInfoComponentTypeEnumStringValues Enumerates the set of values in String for CloudDbSystemConnectionInfoComponentTypeEnum
func GetCloudDbSystemConnectionInfoComponentTypeEnumStringValues() []string {
	return []string{
		"DATABASE",
		"ASM",
	}
}

// GetMappingCloudDbSystemConnectionInfoComponentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudDbSystemConnectionInfoComponentTypeEnum(val string) (CloudDbSystemConnectionInfoComponentTypeEnum, bool) {
	enum, ok := mappingCloudDbSystemConnectionInfoComponentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
