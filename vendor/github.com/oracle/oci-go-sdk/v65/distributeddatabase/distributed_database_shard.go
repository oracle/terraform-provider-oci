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

// DistributedDatabaseShard Globally distributed database shard.
type DistributedDatabaseShard interface {

	// Name of the shard.
	GetName() *string

	// The time the shard was created. An RFC3339 formatted datetime string
	GetTimeCreated() *common.SDKTime

	// The time the shard was last updated. An RFC3339 formatted datetime string
	GetTimeUpdated() *common.SDKTime
}

type distributeddatabaseshard struct {
	JsonData    []byte
	Name        *string         `mandatory:"true" json:"name"`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`
	Source      string          `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *distributeddatabaseshard) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdistributeddatabaseshard distributeddatabaseshard
	s := struct {
		Model Unmarshalerdistributeddatabaseshard
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
func (m *distributeddatabaseshard) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "NEW_VAULT_AND_CLUSTER":
		mm := DistributedDatabaseShardWithExadbXsNewVaultAndCluster{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXADB_XS":
		mm := DistributedDatabaseShardWithExadbXs{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DistributedDatabaseShard: %s.", m.Source)
		return *m, nil
	}
}

// GetName returns Name
func (m distributeddatabaseshard) GetName() *string {
	return m.Name
}

// GetTimeCreated returns TimeCreated
func (m distributeddatabaseshard) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m distributeddatabaseshard) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m distributeddatabaseshard) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m distributeddatabaseshard) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DistributedDatabaseShardSourceEnum Enum with underlying type: string
type DistributedDatabaseShardSourceEnum string

// Set of constants representing the allowable values for DistributedDatabaseShardSourceEnum
const (
	DistributedDatabaseShardSourceExadbXs            DistributedDatabaseShardSourceEnum = "EXADB_XS"
	DistributedDatabaseShardSourceNewVaultAndCluster DistributedDatabaseShardSourceEnum = "NEW_VAULT_AND_CLUSTER"
	DistributedDatabaseShardSourceExistingCluster    DistributedDatabaseShardSourceEnum = "EXISTING_CLUSTER"
)

var mappingDistributedDatabaseShardSourceEnum = map[string]DistributedDatabaseShardSourceEnum{
	"EXADB_XS":              DistributedDatabaseShardSourceExadbXs,
	"NEW_VAULT_AND_CLUSTER": DistributedDatabaseShardSourceNewVaultAndCluster,
	"EXISTING_CLUSTER":      DistributedDatabaseShardSourceExistingCluster,
}

var mappingDistributedDatabaseShardSourceEnumLowerCase = map[string]DistributedDatabaseShardSourceEnum{
	"exadb_xs":              DistributedDatabaseShardSourceExadbXs,
	"new_vault_and_cluster": DistributedDatabaseShardSourceNewVaultAndCluster,
	"existing_cluster":      DistributedDatabaseShardSourceExistingCluster,
}

// GetDistributedDatabaseShardSourceEnumValues Enumerates the set of values for DistributedDatabaseShardSourceEnum
func GetDistributedDatabaseShardSourceEnumValues() []DistributedDatabaseShardSourceEnum {
	values := make([]DistributedDatabaseShardSourceEnum, 0)
	for _, v := range mappingDistributedDatabaseShardSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseShardSourceEnumStringValues Enumerates the set of values in String for DistributedDatabaseShardSourceEnum
func GetDistributedDatabaseShardSourceEnumStringValues() []string {
	return []string{
		"EXADB_XS",
		"NEW_VAULT_AND_CLUSTER",
		"EXISTING_CLUSTER",
	}
}

// GetMappingDistributedDatabaseShardSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseShardSourceEnum(val string) (DistributedDatabaseShardSourceEnum, bool) {
	enum, ok := mappingDistributedDatabaseShardSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
