// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConvertToPdbTargetBase Details of the container database in which the converted pluggable database will be located.
type ConvertToPdbTargetBase interface {
}

type converttopdbtargetbase struct {
	JsonData []byte
	Target   string `json:"target"`
}

// UnmarshalJSON unmarshals json
func (m *converttopdbtargetbase) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconverttopdbtargetbase converttopdbtargetbase
	s := struct {
		Model Unmarshalerconverttopdbtargetbase
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Target = s.Model.Target

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *converttopdbtargetbase) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Target {
	case "NEW_DATABASE":
		mm := PdbConversionToNewDatabaseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m converttopdbtargetbase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m converttopdbtargetbase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConvertToPdbTargetBaseTargetEnum Enum with underlying type: string
type ConvertToPdbTargetBaseTargetEnum string

// Set of constants representing the allowable values for ConvertToPdbTargetBaseTargetEnum
const (
	ConvertToPdbTargetBaseTargetNewDatabase ConvertToPdbTargetBaseTargetEnum = "NEW_DATABASE"
)

var mappingConvertToPdbTargetBaseTargetEnum = map[string]ConvertToPdbTargetBaseTargetEnum{
	"NEW_DATABASE": ConvertToPdbTargetBaseTargetNewDatabase,
}

var mappingConvertToPdbTargetBaseTargetEnumLowerCase = map[string]ConvertToPdbTargetBaseTargetEnum{
	"new_database": ConvertToPdbTargetBaseTargetNewDatabase,
}

// GetConvertToPdbTargetBaseTargetEnumValues Enumerates the set of values for ConvertToPdbTargetBaseTargetEnum
func GetConvertToPdbTargetBaseTargetEnumValues() []ConvertToPdbTargetBaseTargetEnum {
	values := make([]ConvertToPdbTargetBaseTargetEnum, 0)
	for _, v := range mappingConvertToPdbTargetBaseTargetEnum {
		values = append(values, v)
	}
	return values
}

// GetConvertToPdbTargetBaseTargetEnumStringValues Enumerates the set of values in String for ConvertToPdbTargetBaseTargetEnum
func GetConvertToPdbTargetBaseTargetEnumStringValues() []string {
	return []string{
		"NEW_DATABASE",
	}
}

// GetMappingConvertToPdbTargetBaseTargetEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConvertToPdbTargetBaseTargetEnum(val string) (ConvertToPdbTargetBaseTargetEnum, bool) {
	enum, ok := mappingConvertToPdbTargetBaseTargetEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
