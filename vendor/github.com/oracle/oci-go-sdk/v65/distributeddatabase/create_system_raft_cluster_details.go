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

// CreateSystemRaftClusterDetails Details required to create system raft cluster.
type CreateSystemRaftClusterDetails struct {

	// The name of raft cluster.
	// It must start with a letter and contain only letters, digits, and underscores.
	// Maximum length is 40 characters.
	Name *string `mandatory:"true" json:"name"`

	// Replication factor associated with the raft cluster.
	ReplicationFactor *int `mandatory:"true" json:"replicationFactor"`

	// Witnesses RU count associated with the raft cluster.
	WitnessCount *int `mandatory:"true" json:"witnessCount"`

	// Replication factor associated with the raft cluster.
	RuMode CreateSystemRaftClusterDetailsRuModeEnum `mandatory:"true" json:"ruMode"`

	// Details of the databases associated with the system raft cluster.
	Databases []CreateDistributedDatabaseShardDatabaseDetails `mandatory:"true" json:"databases"`
}

func (m CreateSystemRaftClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSystemRaftClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateSystemRaftClusterDetailsRuModeEnum(string(m.RuMode)); !ok && m.RuMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuMode: %s. Supported values are: %s.", m.RuMode, strings.Join(GetCreateSystemRaftClusterDetailsRuModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateSystemRaftClusterDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Name              *string                                         `json:"name"`
		ReplicationFactor *int                                            `json:"replicationFactor"`
		WitnessCount      *int                                            `json:"witnessCount"`
		RuMode            CreateSystemRaftClusterDetailsRuModeEnum        `json:"ruMode"`
		Databases         []createdistributeddatabasesharddatabasedetails `json:"databases"`
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

// CreateSystemRaftClusterDetailsRuModeEnum Enum with underlying type: string
type CreateSystemRaftClusterDetailsRuModeEnum string

// Set of constants representing the allowable values for CreateSystemRaftClusterDetailsRuModeEnum
const (
	CreateSystemRaftClusterDetailsRuModeOnly  CreateSystemRaftClusterDetailsRuModeEnum = "READ_ONLY"
	CreateSystemRaftClusterDetailsRuModeWrite CreateSystemRaftClusterDetailsRuModeEnum = "READ_WRITE"
)

var mappingCreateSystemRaftClusterDetailsRuModeEnum = map[string]CreateSystemRaftClusterDetailsRuModeEnum{
	"READ_ONLY":  CreateSystemRaftClusterDetailsRuModeOnly,
	"READ_WRITE": CreateSystemRaftClusterDetailsRuModeWrite,
}

var mappingCreateSystemRaftClusterDetailsRuModeEnumLowerCase = map[string]CreateSystemRaftClusterDetailsRuModeEnum{
	"read_only":  CreateSystemRaftClusterDetailsRuModeOnly,
	"read_write": CreateSystemRaftClusterDetailsRuModeWrite,
}

// GetCreateSystemRaftClusterDetailsRuModeEnumValues Enumerates the set of values for CreateSystemRaftClusterDetailsRuModeEnum
func GetCreateSystemRaftClusterDetailsRuModeEnumValues() []CreateSystemRaftClusterDetailsRuModeEnum {
	values := make([]CreateSystemRaftClusterDetailsRuModeEnum, 0)
	for _, v := range mappingCreateSystemRaftClusterDetailsRuModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateSystemRaftClusterDetailsRuModeEnumStringValues Enumerates the set of values in String for CreateSystemRaftClusterDetailsRuModeEnum
func GetCreateSystemRaftClusterDetailsRuModeEnumStringValues() []string {
	return []string{
		"READ_ONLY",
		"READ_WRITE",
	}
}

// GetMappingCreateSystemRaftClusterDetailsRuModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateSystemRaftClusterDetailsRuModeEnum(val string) (CreateSystemRaftClusterDetailsRuModeEnum, bool) {
	enum, ok := mappingCreateSystemRaftClusterDetailsRuModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
