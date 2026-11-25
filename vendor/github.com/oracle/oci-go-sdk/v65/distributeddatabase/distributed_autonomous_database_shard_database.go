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

// DistributedAutonomousDatabaseShardDatabase Globally distributed autonomous database shard.
type DistributedAutonomousDatabaseShardDatabase interface {

	// Name of the shard.
	GetName() *string

	// The time the shard was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The time the shard was last updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	GetMetadata() *DistributedAutonomousDbMetadata
}

type distributedautonomousdatabasesharddatabase struct {
	JsonData    []byte
	Metadata    *DistributedAutonomousDbMetadata `mandatory:"false" json:"metadata"`
	Name        *string                          `mandatory:"true" json:"name"`
	TimeCreated *common.SDKTime                  `mandatory:"true" json:"timeCreated"`
	TimeUpdated *common.SDKTime                  `mandatory:"true" json:"timeUpdated"`
	Source      string                           `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *distributedautonomousdatabasesharddatabase) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdistributedautonomousdatabasesharddatabase distributedautonomousdatabasesharddatabase
	s := struct {
		Model Unmarshalerdistributedautonomousdatabasesharddatabase
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.Metadata = s.Model.Metadata
	m.Source = s.Model.Source

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *distributedautonomousdatabasesharddatabase) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "ADBD_EXISTING_CLUSTER":
		mm := DistributedAutonomousDatabaseShardWithDedicatedInfra{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DistributedAutonomousDatabaseShardDatabase: %s.", m.Source)
		return *m, nil
	}
}

// GetMetadata returns Metadata
func (m distributedautonomousdatabasesharddatabase) GetMetadata() *DistributedAutonomousDbMetadata {
	return m.Metadata
}

// GetName returns Name
func (m distributedautonomousdatabasesharddatabase) GetName() *string {
	return m.Name
}

// GetTimeCreated returns TimeCreated
func (m distributedautonomousdatabasesharddatabase) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m distributedautonomousdatabasesharddatabase) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m distributedautonomousdatabasesharddatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m distributedautonomousdatabasesharddatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DistributedAutonomousDatabaseShardDatabaseSourceEnum Enum with underlying type: string
type DistributedAutonomousDatabaseShardDatabaseSourceEnum string

// Set of constants representing the allowable values for DistributedAutonomousDatabaseShardDatabaseSourceEnum
const (
	DistributedAutonomousDatabaseShardDatabaseSourceAdbdExistingCluster DistributedAutonomousDatabaseShardDatabaseSourceEnum = "ADBD_EXISTING_CLUSTER"
)

var mappingDistributedAutonomousDatabaseShardDatabaseSourceEnum = map[string]DistributedAutonomousDatabaseShardDatabaseSourceEnum{
	"ADBD_EXISTING_CLUSTER": DistributedAutonomousDatabaseShardDatabaseSourceAdbdExistingCluster,
}

var mappingDistributedAutonomousDatabaseShardDatabaseSourceEnumLowerCase = map[string]DistributedAutonomousDatabaseShardDatabaseSourceEnum{
	"adbd_existing_cluster": DistributedAutonomousDatabaseShardDatabaseSourceAdbdExistingCluster,
}

// GetDistributedAutonomousDatabaseShardDatabaseSourceEnumValues Enumerates the set of values for DistributedAutonomousDatabaseShardDatabaseSourceEnum
func GetDistributedAutonomousDatabaseShardDatabaseSourceEnumValues() []DistributedAutonomousDatabaseShardDatabaseSourceEnum {
	values := make([]DistributedAutonomousDatabaseShardDatabaseSourceEnum, 0)
	for _, v := range mappingDistributedAutonomousDatabaseShardDatabaseSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedAutonomousDatabaseShardDatabaseSourceEnumStringValues Enumerates the set of values in String for DistributedAutonomousDatabaseShardDatabaseSourceEnum
func GetDistributedAutonomousDatabaseShardDatabaseSourceEnumStringValues() []string {
	return []string{
		"ADBD_EXISTING_CLUSTER",
	}
}

// GetMappingDistributedAutonomousDatabaseShardDatabaseSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedAutonomousDatabaseShardDatabaseSourceEnum(val string) (DistributedAutonomousDatabaseShardDatabaseSourceEnum, bool) {
	enum, ok := mappingDistributedAutonomousDatabaseShardDatabaseSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
