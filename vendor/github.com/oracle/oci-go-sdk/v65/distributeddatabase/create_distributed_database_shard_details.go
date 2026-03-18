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

// CreateDistributedDatabaseShardDetails Globally distributed database shard.
type CreateDistributedDatabaseShardDetails interface {
}

type createdistributeddatabasesharddetails struct {
	JsonData []byte
	Source   string `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *createdistributeddatabasesharddetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedistributeddatabasesharddetails createdistributeddatabasesharddetails
	s := struct {
		Model Unmarshalercreatedistributeddatabasesharddetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Source = s.Model.Source

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdistributeddatabasesharddetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "NEW_VAULT_AND_CLUSTER":
		mm := CreateDistributedDatabaseShardWithExadbXsNewVaultAndClusterDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXADB_XS":
		mm := CreateDistributedDatabaseShardWithExadbXsDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateDistributedDatabaseShardDetails: %s.", m.Source)
		return *m, nil
	}
}

func (m createdistributeddatabasesharddetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdistributeddatabasesharddetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDistributedDatabaseShardDetailsSourceEnum Enum with underlying type: string
type CreateDistributedDatabaseShardDetailsSourceEnum string

// Set of constants representing the allowable values for CreateDistributedDatabaseShardDetailsSourceEnum
const (
	CreateDistributedDatabaseShardDetailsSourceExadbXs            CreateDistributedDatabaseShardDetailsSourceEnum = "EXADB_XS"
	CreateDistributedDatabaseShardDetailsSourceNewVaultAndCluster CreateDistributedDatabaseShardDetailsSourceEnum = "NEW_VAULT_AND_CLUSTER"
	CreateDistributedDatabaseShardDetailsSourceExistingCluster    CreateDistributedDatabaseShardDetailsSourceEnum = "EXISTING_CLUSTER"
)

var mappingCreateDistributedDatabaseShardDetailsSourceEnum = map[string]CreateDistributedDatabaseShardDetailsSourceEnum{
	"EXADB_XS":              CreateDistributedDatabaseShardDetailsSourceExadbXs,
	"NEW_VAULT_AND_CLUSTER": CreateDistributedDatabaseShardDetailsSourceNewVaultAndCluster,
	"EXISTING_CLUSTER":      CreateDistributedDatabaseShardDetailsSourceExistingCluster,
}

var mappingCreateDistributedDatabaseShardDetailsSourceEnumLowerCase = map[string]CreateDistributedDatabaseShardDetailsSourceEnum{
	"exadb_xs":              CreateDistributedDatabaseShardDetailsSourceExadbXs,
	"new_vault_and_cluster": CreateDistributedDatabaseShardDetailsSourceNewVaultAndCluster,
	"existing_cluster":      CreateDistributedDatabaseShardDetailsSourceExistingCluster,
}

// GetCreateDistributedDatabaseShardDetailsSourceEnumValues Enumerates the set of values for CreateDistributedDatabaseShardDetailsSourceEnum
func GetCreateDistributedDatabaseShardDetailsSourceEnumValues() []CreateDistributedDatabaseShardDetailsSourceEnum {
	values := make([]CreateDistributedDatabaseShardDetailsSourceEnum, 0)
	for _, v := range mappingCreateDistributedDatabaseShardDetailsSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDistributedDatabaseShardDetailsSourceEnumStringValues Enumerates the set of values in String for CreateDistributedDatabaseShardDetailsSourceEnum
func GetCreateDistributedDatabaseShardDetailsSourceEnumStringValues() []string {
	return []string{
		"EXADB_XS",
		"NEW_VAULT_AND_CLUSTER",
		"EXISTING_CLUSTER",
	}
}

// GetMappingCreateDistributedDatabaseShardDetailsSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDistributedDatabaseShardDetailsSourceEnum(val string) (CreateDistributedDatabaseShardDetailsSourceEnum, bool) {
	enum, ok := mappingCreateDistributedDatabaseShardDetailsSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
