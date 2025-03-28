// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// DatabaseToolsKeyStoreContentSummary The key store content.
type DatabaseToolsKeyStoreContentSummary interface {
}

type databasetoolskeystorecontentsummary struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorecontentsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorecontentsummary databasetoolskeystorecontentsummary
	s := struct {
		Model Unmarshalerdatabasetoolskeystorecontentsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorecontentsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStoreContentSecretIdSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DatabaseToolsKeyStoreContentSummary: %s.", m.ValueType)
		return *m, nil
	}
}

func (m databasetoolskeystorecontentsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorecontentsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStoreContentSummaryValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStoreContentSummaryValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStoreContentSummaryValueTypeEnum
const (
	DatabaseToolsKeyStoreContentSummaryValueTypeSecretid DatabaseToolsKeyStoreContentSummaryValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStoreContentSummaryValueTypeEnum = map[string]DatabaseToolsKeyStoreContentSummaryValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStoreContentSummaryValueTypeSecretid,
}

var mappingDatabaseToolsKeyStoreContentSummaryValueTypeEnumLowerCase = map[string]DatabaseToolsKeyStoreContentSummaryValueTypeEnum{
	"secretid": DatabaseToolsKeyStoreContentSummaryValueTypeSecretid,
}

// GetDatabaseToolsKeyStoreContentSummaryValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStoreContentSummaryValueTypeEnum
func GetDatabaseToolsKeyStoreContentSummaryValueTypeEnumValues() []DatabaseToolsKeyStoreContentSummaryValueTypeEnum {
	values := make([]DatabaseToolsKeyStoreContentSummaryValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStoreContentSummaryValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStoreContentSummaryValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStoreContentSummaryValueTypeEnum
func GetDatabaseToolsKeyStoreContentSummaryValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStoreContentSummaryValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStoreContentSummaryValueTypeEnum(val string) (DatabaseToolsKeyStoreContentSummaryValueTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsKeyStoreContentSummaryValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
