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

// CreateAutonomousCompositeShardSpaceRaftClusterDetails Details required to create raft clusters for a composite shard space.
type CreateAutonomousCompositeShardSpaceRaftClusterDetails struct {

	// The name of raft cluster.
	// It must start with a letter and contain only letters, digits, and underscores.
	// Maximum length is 40 characters.
	Name *string `mandatory:"true" json:"name"`

	// Replication factor associated with the raft cluster.
	ReplicationFactor *int `mandatory:"true" json:"replicationFactor"`

	// Witnesses RU count associated with the raft cluster.
	WitnessCount *int `mandatory:"true" json:"witnessCount"`

	// Replication factor associated with the raft cluster.
	RuMode CreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnum `mandatory:"true" json:"ruMode"`

	// The details of databases associated with the composite shard space raft cluster.
	Databases []CreateDistributedAutonomousDatabaseShardDatabaseDetails `mandatory:"true" json:"databases"`
}

func (m CreateAutonomousCompositeShardSpaceRaftClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAutonomousCompositeShardSpaceRaftClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnum(string(m.RuMode)); !ok && m.RuMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuMode: %s. Supported values are: %s.", m.RuMode, strings.Join(GetCreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateAutonomousCompositeShardSpaceRaftClusterDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Name              *string                                                         `json:"name"`
		ReplicationFactor *int                                                            `json:"replicationFactor"`
		WitnessCount      *int                                                            `json:"witnessCount"`
		RuMode            CreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnum `json:"ruMode"`
		Databases         []createdistributedautonomousdatabasesharddatabasedetails       `json:"databases"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Name = model.Name

	m.ReplicationFactor = model.ReplicationFactor

	m.WitnessCount = model.WitnessCount

	m.RuMode = model.RuMode

	m.Databases = make([]CreateDistributedAutonomousDatabaseShardDatabaseDetails, len(model.Databases))
	for i, n := range model.Databases {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Databases[i] = nn.(CreateDistributedAutonomousDatabaseShardDatabaseDetails)
		} else {
			m.Databases[i] = nil
		}
	}
	return
}

// CreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnum Enum with underlying type: string
type CreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnum string

// Set of constants representing the allowable values for CreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnum
const (
	CreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeOnly  CreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnum = "READ_ONLY"
	CreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeWrite CreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnum = "READ_WRITE"
)

var mappingCreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnum = map[string]CreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnum{
	"READ_ONLY":  CreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeOnly,
	"READ_WRITE": CreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeWrite,
}

var mappingCreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnumLowerCase = map[string]CreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnum{
	"read_only":  CreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeOnly,
	"read_write": CreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeWrite,
}

// GetCreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnumValues Enumerates the set of values for CreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnum
func GetCreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnumValues() []CreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnum {
	values := make([]CreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnum, 0)
	for _, v := range mappingCreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnumStringValues Enumerates the set of values in String for CreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnum
func GetCreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnumStringValues() []string {
	return []string{
		"READ_ONLY",
		"READ_WRITE",
	}
}

// GetMappingCreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnum(val string) (CreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnum, bool) {
	enum, ok := mappingCreateAutonomousCompositeShardSpaceRaftClusterDetailsRuModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
