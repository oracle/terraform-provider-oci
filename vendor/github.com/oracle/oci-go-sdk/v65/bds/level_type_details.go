// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LevelTypeDetails Details of the type of level used to trigger the creation of a new node backup configuration or node replacement configuration.
type LevelTypeDetails interface {
}

type leveltypedetails struct {
	JsonData  []byte
	LevelType string `json:"levelType"`
}

// UnmarshalJSON unmarshals json
func (m *leveltypedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerleveltypedetails leveltypedetails
	s := struct {
		Model Unmarshalerleveltypedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.LevelType = s.Model.LevelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *leveltypedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.LevelType {
	case "NODE_TYPE_LEVEL":
		mm := NodeTypeLevelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NODE_LEVEL":
		mm := NodeLevelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for LevelTypeDetails: %s.", m.LevelType)
		return *m, nil
	}
}

func (m leveltypedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m leveltypedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LevelTypeDetailsLevelTypeEnum Enum with underlying type: string
type LevelTypeDetailsLevelTypeEnum string

// Set of constants representing the allowable values for LevelTypeDetailsLevelTypeEnum
const (
	LevelTypeDetailsLevelTypeLevel     LevelTypeDetailsLevelTypeEnum = "NODE_LEVEL"
	LevelTypeDetailsLevelTypeTypeLevel LevelTypeDetailsLevelTypeEnum = "NODE_TYPE_LEVEL"
)

var mappingLevelTypeDetailsLevelTypeEnum = map[string]LevelTypeDetailsLevelTypeEnum{
	"NODE_LEVEL":      LevelTypeDetailsLevelTypeLevel,
	"NODE_TYPE_LEVEL": LevelTypeDetailsLevelTypeTypeLevel,
}

var mappingLevelTypeDetailsLevelTypeEnumLowerCase = map[string]LevelTypeDetailsLevelTypeEnum{
	"node_level":      LevelTypeDetailsLevelTypeLevel,
	"node_type_level": LevelTypeDetailsLevelTypeTypeLevel,
}

// GetLevelTypeDetailsLevelTypeEnumValues Enumerates the set of values for LevelTypeDetailsLevelTypeEnum
func GetLevelTypeDetailsLevelTypeEnumValues() []LevelTypeDetailsLevelTypeEnum {
	values := make([]LevelTypeDetailsLevelTypeEnum, 0)
	for _, v := range mappingLevelTypeDetailsLevelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLevelTypeDetailsLevelTypeEnumStringValues Enumerates the set of values in String for LevelTypeDetailsLevelTypeEnum
func GetLevelTypeDetailsLevelTypeEnumStringValues() []string {
	return []string{
		"NODE_LEVEL",
		"NODE_TYPE_LEVEL",
	}
}

// GetMappingLevelTypeDetailsLevelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLevelTypeDetailsLevelTypeEnum(val string) (LevelTypeDetailsLevelTypeEnum, bool) {
	enum, ok := mappingLevelTypeDetailsLevelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
