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

// DistributedDatabaseShardDatabase Details of a distributed database shard.
type DistributedDatabaseShardDatabase interface {

	// Name of the shard.
	GetName() *string

	// The time the shard was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The time the shard was last updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	GetMetadata() *DistributedDbMetadata
}

type distributeddatabasesharddatabase struct {
	JsonData    []byte
	Metadata    *DistributedDbMetadata `mandatory:"false" json:"metadata"`
	Name        *string                `mandatory:"true" json:"name"`
	TimeCreated *common.SDKTime        `mandatory:"true" json:"timeCreated"`
	TimeUpdated *common.SDKTime        `mandatory:"true" json:"timeUpdated"`
	Source      string                 `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *distributeddatabasesharddatabase) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdistributeddatabasesharddatabase distributeddatabasesharddatabase
	s := struct {
		Model Unmarshalerdistributeddatabasesharddatabase
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
func (m *distributeddatabasesharddatabase) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "XS_NEW_CLUSTER":
		mm := DistributedDatabaseShardDatabaseWithExadbXsNewCluster{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "XS_NEW_VAULT_AND_CLUSTER":
		mm := DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndCluster{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "XD_EXISTING_CLUSTER":
		mm := DistributedDatabaseShardDatabaseWithExadbXd{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXISTING_DB_HOME":
		mm := DistributedDatabaseShardDatabaseWithDbHome{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "XS_EXISTING_CLUSTER":
		mm := DistributedDatabaseShardDatabaseWithExadbXs{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DistributedDatabaseShardDatabase: %s.", m.Source)
		return *m, nil
	}
}

// GetMetadata returns Metadata
func (m distributeddatabasesharddatabase) GetMetadata() *DistributedDbMetadata {
	return m.Metadata
}

// GetName returns Name
func (m distributeddatabasesharddatabase) GetName() *string {
	return m.Name
}

// GetTimeCreated returns TimeCreated
func (m distributeddatabasesharddatabase) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m distributeddatabasesharddatabase) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m distributeddatabasesharddatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m distributeddatabasesharddatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DistributedDatabaseShardDatabaseSourceEnum Enum with underlying type: string
type DistributedDatabaseShardDatabaseSourceEnum string

// Set of constants representing the allowable values for DistributedDatabaseShardDatabaseSourceEnum
const (
	DistributedDatabaseShardDatabaseSourceXsExistingCluster    DistributedDatabaseShardDatabaseSourceEnum = "XS_EXISTING_CLUSTER"
	DistributedDatabaseShardDatabaseSourceXsNewCluster         DistributedDatabaseShardDatabaseSourceEnum = "XS_NEW_CLUSTER"
	DistributedDatabaseShardDatabaseSourceXsNewVaultAndCluster DistributedDatabaseShardDatabaseSourceEnum = "XS_NEW_VAULT_AND_CLUSTER"
	DistributedDatabaseShardDatabaseSourceXdExistingCluster    DistributedDatabaseShardDatabaseSourceEnum = "XD_EXISTING_CLUSTER"
	DistributedDatabaseShardDatabaseSourceExistingDbHome       DistributedDatabaseShardDatabaseSourceEnum = "EXISTING_DB_HOME"
)

var mappingDistributedDatabaseShardDatabaseSourceEnum = map[string]DistributedDatabaseShardDatabaseSourceEnum{
	"XS_EXISTING_CLUSTER":      DistributedDatabaseShardDatabaseSourceXsExistingCluster,
	"XS_NEW_CLUSTER":           DistributedDatabaseShardDatabaseSourceXsNewCluster,
	"XS_NEW_VAULT_AND_CLUSTER": DistributedDatabaseShardDatabaseSourceXsNewVaultAndCluster,
	"XD_EXISTING_CLUSTER":      DistributedDatabaseShardDatabaseSourceXdExistingCluster,
	"EXISTING_DB_HOME":         DistributedDatabaseShardDatabaseSourceExistingDbHome,
}

var mappingDistributedDatabaseShardDatabaseSourceEnumLowerCase = map[string]DistributedDatabaseShardDatabaseSourceEnum{
	"xs_existing_cluster":      DistributedDatabaseShardDatabaseSourceXsExistingCluster,
	"xs_new_cluster":           DistributedDatabaseShardDatabaseSourceXsNewCluster,
	"xs_new_vault_and_cluster": DistributedDatabaseShardDatabaseSourceXsNewVaultAndCluster,
	"xd_existing_cluster":      DistributedDatabaseShardDatabaseSourceXdExistingCluster,
	"existing_db_home":         DistributedDatabaseShardDatabaseSourceExistingDbHome,
}

// GetDistributedDatabaseShardDatabaseSourceEnumValues Enumerates the set of values for DistributedDatabaseShardDatabaseSourceEnum
func GetDistributedDatabaseShardDatabaseSourceEnumValues() []DistributedDatabaseShardDatabaseSourceEnum {
	values := make([]DistributedDatabaseShardDatabaseSourceEnum, 0)
	for _, v := range mappingDistributedDatabaseShardDatabaseSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseShardDatabaseSourceEnumStringValues Enumerates the set of values in String for DistributedDatabaseShardDatabaseSourceEnum
func GetDistributedDatabaseShardDatabaseSourceEnumStringValues() []string {
	return []string{
		"XS_EXISTING_CLUSTER",
		"XS_NEW_CLUSTER",
		"XS_NEW_VAULT_AND_CLUSTER",
		"XD_EXISTING_CLUSTER",
		"EXISTING_DB_HOME",
	}
}

// GetMappingDistributedDatabaseShardDatabaseSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseShardDatabaseSourceEnum(val string) (DistributedDatabaseShardDatabaseSourceEnum, bool) {
	enum, ok := mappingDistributedDatabaseShardDatabaseSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
