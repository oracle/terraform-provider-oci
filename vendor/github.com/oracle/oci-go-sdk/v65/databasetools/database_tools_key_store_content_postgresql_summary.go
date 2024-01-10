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

// DatabaseToolsKeyStoreContentPostgresqlSummary The key store content.
type DatabaseToolsKeyStoreContentPostgresqlSummary interface {
}

type databasetoolskeystorecontentpostgresqlsummary struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorecontentpostgresqlsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorecontentpostgresqlsummary databasetoolskeystorecontentpostgresqlsummary
	s := struct {
		Model Unmarshalerdatabasetoolskeystorecontentpostgresqlsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorecontentpostgresqlsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStoreContentSecretIdPostgresqlSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DatabaseToolsKeyStoreContentPostgresqlSummary: %s.", m.ValueType)
		return *m, nil
	}
}

func (m databasetoolskeystorecontentpostgresqlsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorecontentpostgresqlsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeEnum
const (
	DatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeSecretid DatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeEnum = map[string]DatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeSecretid,
}

var mappingDatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeEnumLowerCase = map[string]DatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeEnum{
	"secretid": DatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeSecretid,
}

// GetDatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeEnum
func GetDatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeEnumValues() []DatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeEnum {
	values := make([]DatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeEnum
func GetDatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeEnum(val string) (DatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsKeyStoreContentPostgresqlSummaryValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
