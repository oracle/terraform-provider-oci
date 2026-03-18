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

// DistributedAutonomousDatabaseShard Globally distributed autonomous database shard.
type DistributedAutonomousDatabaseShard interface {

	// Name of the shard.
	GetName() *string

	// The time the shard was created. An RFC3339 formatted datetime string
	GetTimeCreated() *common.SDKTime

	// The time the shard was last updated. An RFC3339 formatted datetime string
	GetTimeUpdated() *common.SDKTime
}

type distributedautonomousdatabaseshard struct {
	JsonData    []byte
	Name        *string         `mandatory:"true" json:"name"`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`
	Source      string          `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *distributedautonomousdatabaseshard) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdistributedautonomousdatabaseshard distributedautonomousdatabaseshard
	s := struct {
		Model Unmarshalerdistributedautonomousdatabaseshard
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.Source = s.Model.Source

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *distributedautonomousdatabaseshard) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "ADB_D":
		mm := DistributedAutonomousDatabaseShardWithDedicatedInfra{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DistributedAutonomousDatabaseShard: %s.", m.Source)
		return *m, nil
	}
}

// GetName returns Name
func (m distributedautonomousdatabaseshard) GetName() *string {
	return m.Name
}

// GetTimeCreated returns TimeCreated
func (m distributedautonomousdatabaseshard) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m distributedautonomousdatabaseshard) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m distributedautonomousdatabaseshard) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m distributedautonomousdatabaseshard) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DistributedAutonomousDatabaseShardSourceEnum Enum with underlying type: string
type DistributedAutonomousDatabaseShardSourceEnum string

// Set of constants representing the allowable values for DistributedAutonomousDatabaseShardSourceEnum
const (
	DistributedAutonomousDatabaseShardSourceAdbD DistributedAutonomousDatabaseShardSourceEnum = "ADB_D"
)

var mappingDistributedAutonomousDatabaseShardSourceEnum = map[string]DistributedAutonomousDatabaseShardSourceEnum{
	"ADB_D": DistributedAutonomousDatabaseShardSourceAdbD,
}

var mappingDistributedAutonomousDatabaseShardSourceEnumLowerCase = map[string]DistributedAutonomousDatabaseShardSourceEnum{
	"adb_d": DistributedAutonomousDatabaseShardSourceAdbD,
}

// GetDistributedAutonomousDatabaseShardSourceEnumValues Enumerates the set of values for DistributedAutonomousDatabaseShardSourceEnum
func GetDistributedAutonomousDatabaseShardSourceEnumValues() []DistributedAutonomousDatabaseShardSourceEnum {
	values := make([]DistributedAutonomousDatabaseShardSourceEnum, 0)
	for _, v := range mappingDistributedAutonomousDatabaseShardSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedAutonomousDatabaseShardSourceEnumStringValues Enumerates the set of values in String for DistributedAutonomousDatabaseShardSourceEnum
func GetDistributedAutonomousDatabaseShardSourceEnumStringValues() []string {
	return []string{
		"ADB_D",
	}
}

// GetMappingDistributedAutonomousDatabaseShardSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedAutonomousDatabaseShardSourceEnum(val string) (DistributedAutonomousDatabaseShardSourceEnum, bool) {
	enum, ok := mappingDistributedAutonomousDatabaseShardSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
