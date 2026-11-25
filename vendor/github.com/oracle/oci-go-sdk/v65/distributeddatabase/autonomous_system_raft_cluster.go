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

// AutonomousSystemRaftCluster Details required to create system raft cluster.
type AutonomousSystemRaftCluster struct {

	// The name of raft cluster.
	// It must start with a letter and contain only letters, digits, and underscores.
	// Maximum length is 40 characters.
	Name *string `mandatory:"true" json:"name"`

	// Replication factor associated with the raft cluster.
	ReplicationFactor *int `mandatory:"true" json:"replicationFactor"`

	// Witnesses RU count associated with the raft cluster.
	WitnessCount *int `mandatory:"true" json:"witnessCount"`

	// Replication factor associated with the raft cluster.
	RuMode AutonomousSystemRaftClusterRuModeEnum `mandatory:"true" json:"ruMode"`

	// Details of the databases associated with the system raft cluster.
	Databases []DistributedAutonomousDatabaseShardDatabase `mandatory:"true" json:"databases"`
}

func (m AutonomousSystemRaftCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousSystemRaftCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousSystemRaftClusterRuModeEnum(string(m.RuMode)); !ok && m.RuMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuMode: %s. Supported values are: %s.", m.RuMode, strings.Join(GetAutonomousSystemRaftClusterRuModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *AutonomousSystemRaftCluster) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Name              *string                                      `json:"name"`
		ReplicationFactor *int                                         `json:"replicationFactor"`
		WitnessCount      *int                                         `json:"witnessCount"`
		RuMode            AutonomousSystemRaftClusterRuModeEnum        `json:"ruMode"`
		Databases         []distributedautonomousdatabasesharddatabase `json:"databases"`
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

	m.Databases = make([]DistributedAutonomousDatabaseShardDatabase, len(model.Databases))
	for i, n := range model.Databases {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Databases[i] = nn.(DistributedAutonomousDatabaseShardDatabase)
		} else {
			m.Databases[i] = nil
		}
	}
	return
}

// AutonomousSystemRaftClusterRuModeEnum Enum with underlying type: string
type AutonomousSystemRaftClusterRuModeEnum string

// Set of constants representing the allowable values for AutonomousSystemRaftClusterRuModeEnum
const (
	AutonomousSystemRaftClusterRuModeOnly  AutonomousSystemRaftClusterRuModeEnum = "READ_ONLY"
	AutonomousSystemRaftClusterRuModeWrite AutonomousSystemRaftClusterRuModeEnum = "READ_WRITE"
)

var mappingAutonomousSystemRaftClusterRuModeEnum = map[string]AutonomousSystemRaftClusterRuModeEnum{
	"READ_ONLY":  AutonomousSystemRaftClusterRuModeOnly,
	"READ_WRITE": AutonomousSystemRaftClusterRuModeWrite,
}

var mappingAutonomousSystemRaftClusterRuModeEnumLowerCase = map[string]AutonomousSystemRaftClusterRuModeEnum{
	"read_only":  AutonomousSystemRaftClusterRuModeOnly,
	"read_write": AutonomousSystemRaftClusterRuModeWrite,
}

// GetAutonomousSystemRaftClusterRuModeEnumValues Enumerates the set of values for AutonomousSystemRaftClusterRuModeEnum
func GetAutonomousSystemRaftClusterRuModeEnumValues() []AutonomousSystemRaftClusterRuModeEnum {
	values := make([]AutonomousSystemRaftClusterRuModeEnum, 0)
	for _, v := range mappingAutonomousSystemRaftClusterRuModeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousSystemRaftClusterRuModeEnumStringValues Enumerates the set of values in String for AutonomousSystemRaftClusterRuModeEnum
func GetAutonomousSystemRaftClusterRuModeEnumStringValues() []string {
	return []string{
		"READ_ONLY",
		"READ_WRITE",
	}
}

// GetMappingAutonomousSystemRaftClusterRuModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousSystemRaftClusterRuModeEnum(val string) (AutonomousSystemRaftClusterRuModeEnum, bool) {
	enum, ok := mappingAutonomousSystemRaftClusterRuModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
