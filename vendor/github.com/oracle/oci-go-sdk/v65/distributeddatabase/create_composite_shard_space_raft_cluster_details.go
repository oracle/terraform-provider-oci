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

// CreateCompositeShardSpaceRaftClusterDetails Raft clusters details required for creation of composite shard spaces.
type CreateCompositeShardSpaceRaftClusterDetails struct {

	// The name of raft cluster.
	// It must start with a letter and contain only letters, digits, and underscores.
	// Maximum length is 40 characters.
	Name *string `mandatory:"true" json:"name"`

	// Replication factor associated with the raft cluster.
	ReplicationFactor *int `mandatory:"true" json:"replicationFactor"`

	// Witnesses RU count associated with the raft cluster.
	WitnessCount *int `mandatory:"true" json:"witnessCount"`

	// Replication factor associated with the raft cluster.
	RuMode CreateCompositeShardSpaceRaftClusterDetailsRuModeEnum `mandatory:"true" json:"ruMode"`

	// The details of databases associated with the composite shard space raft cluster.
	Databases []CreateDistributedDatabaseShardDatabaseDetails `mandatory:"true" json:"databases"`
}

func (m CreateCompositeShardSpaceRaftClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCompositeShardSpaceRaftClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateCompositeShardSpaceRaftClusterDetailsRuModeEnum(string(m.RuMode)); !ok && m.RuMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuMode: %s. Supported values are: %s.", m.RuMode, strings.Join(GetCreateCompositeShardSpaceRaftClusterDetailsRuModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateCompositeShardSpaceRaftClusterDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Name              *string                                               `json:"name"`
		ReplicationFactor *int                                                  `json:"replicationFactor"`
		WitnessCount      *int                                                  `json:"witnessCount"`
		RuMode            CreateCompositeShardSpaceRaftClusterDetailsRuModeEnum `json:"ruMode"`
		Databases         []createdistributeddatabasesharddatabasedetails       `json:"databases"`
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

	m.Databases = make([]CreateDistributedDatabaseShardDatabaseDetails, len(model.Databases))
	for i, n := range model.Databases {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Databases[i] = nn.(CreateDistributedDatabaseShardDatabaseDetails)
		} else {
			m.Databases[i] = nil
		}
	}
	return
}

// CreateCompositeShardSpaceRaftClusterDetailsRuModeEnum Enum with underlying type: string
type CreateCompositeShardSpaceRaftClusterDetailsRuModeEnum string

// Set of constants representing the allowable values for CreateCompositeShardSpaceRaftClusterDetailsRuModeEnum
const (
	CreateCompositeShardSpaceRaftClusterDetailsRuModeOnly  CreateCompositeShardSpaceRaftClusterDetailsRuModeEnum = "READ_ONLY"
	CreateCompositeShardSpaceRaftClusterDetailsRuModeWrite CreateCompositeShardSpaceRaftClusterDetailsRuModeEnum = "READ_WRITE"
)

var mappingCreateCompositeShardSpaceRaftClusterDetailsRuModeEnum = map[string]CreateCompositeShardSpaceRaftClusterDetailsRuModeEnum{
	"READ_ONLY":  CreateCompositeShardSpaceRaftClusterDetailsRuModeOnly,
	"READ_WRITE": CreateCompositeShardSpaceRaftClusterDetailsRuModeWrite,
}

var mappingCreateCompositeShardSpaceRaftClusterDetailsRuModeEnumLowerCase = map[string]CreateCompositeShardSpaceRaftClusterDetailsRuModeEnum{
	"read_only":  CreateCompositeShardSpaceRaftClusterDetailsRuModeOnly,
	"read_write": CreateCompositeShardSpaceRaftClusterDetailsRuModeWrite,
}

// GetCreateCompositeShardSpaceRaftClusterDetailsRuModeEnumValues Enumerates the set of values for CreateCompositeShardSpaceRaftClusterDetailsRuModeEnum
func GetCreateCompositeShardSpaceRaftClusterDetailsRuModeEnumValues() []CreateCompositeShardSpaceRaftClusterDetailsRuModeEnum {
	values := make([]CreateCompositeShardSpaceRaftClusterDetailsRuModeEnum, 0)
	for _, v := range mappingCreateCompositeShardSpaceRaftClusterDetailsRuModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateCompositeShardSpaceRaftClusterDetailsRuModeEnumStringValues Enumerates the set of values in String for CreateCompositeShardSpaceRaftClusterDetailsRuModeEnum
func GetCreateCompositeShardSpaceRaftClusterDetailsRuModeEnumStringValues() []string {
	return []string{
		"READ_ONLY",
		"READ_WRITE",
	}
}

// GetMappingCreateCompositeShardSpaceRaftClusterDetailsRuModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateCompositeShardSpaceRaftClusterDetailsRuModeEnum(val string) (CreateCompositeShardSpaceRaftClusterDetailsRuModeEnum, bool) {
	enum, ok := mappingCreateCompositeShardSpaceRaftClusterDetailsRuModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
