// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDistributedAutonomousDatabaseShardDetails Globally distributed autonomous database shard.
type CreateDistributedAutonomousDatabaseShardDetails interface {
}

type createdistributedautonomousdatabasesharddetails struct {
	JsonData []byte
	Source   string `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *createdistributedautonomousdatabasesharddetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedistributedautonomousdatabasesharddetails createdistributedautonomousdatabasesharddetails
	s := struct {
		Model Unmarshalercreatedistributedautonomousdatabasesharddetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Source = s.Model.Source

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdistributedautonomousdatabasesharddetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "ADB_D":
		mm := CreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateDistributedAutonomousDatabaseShardDetails: %s.", m.Source)
		return *m, nil
	}
}

func (m createdistributedautonomousdatabasesharddetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdistributedautonomousdatabasesharddetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDistributedAutonomousDatabaseShardDetailsSourceEnum Enum with underlying type: string
type CreateDistributedAutonomousDatabaseShardDetailsSourceEnum string

// Set of constants representing the allowable values for CreateDistributedAutonomousDatabaseShardDetailsSourceEnum
const (
	CreateDistributedAutonomousDatabaseShardDetailsSourceAdbD CreateDistributedAutonomousDatabaseShardDetailsSourceEnum = "ADB_D"
)

var mappingCreateDistributedAutonomousDatabaseShardDetailsSourceEnum = map[string]CreateDistributedAutonomousDatabaseShardDetailsSourceEnum{
	"ADB_D": CreateDistributedAutonomousDatabaseShardDetailsSourceAdbD,
}

var mappingCreateDistributedAutonomousDatabaseShardDetailsSourceEnumLowerCase = map[string]CreateDistributedAutonomousDatabaseShardDetailsSourceEnum{
	"adb_d": CreateDistributedAutonomousDatabaseShardDetailsSourceAdbD,
}

// GetCreateDistributedAutonomousDatabaseShardDetailsSourceEnumValues Enumerates the set of values for CreateDistributedAutonomousDatabaseShardDetailsSourceEnum
func GetCreateDistributedAutonomousDatabaseShardDetailsSourceEnumValues() []CreateDistributedAutonomousDatabaseShardDetailsSourceEnum {
	values := make([]CreateDistributedAutonomousDatabaseShardDetailsSourceEnum, 0)
	for _, v := range mappingCreateDistributedAutonomousDatabaseShardDetailsSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDistributedAutonomousDatabaseShardDetailsSourceEnumStringValues Enumerates the set of values in String for CreateDistributedAutonomousDatabaseShardDetailsSourceEnum
func GetCreateDistributedAutonomousDatabaseShardDetailsSourceEnumStringValues() []string {
	return []string{
		"ADB_D",
	}
}

// GetMappingCreateDistributedAutonomousDatabaseShardDetailsSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDistributedAutonomousDatabaseShardDetailsSourceEnum(val string) (CreateDistributedAutonomousDatabaseShardDetailsSourceEnum, bool) {
	enum, ok := mappingCreateDistributedAutonomousDatabaseShardDetailsSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
