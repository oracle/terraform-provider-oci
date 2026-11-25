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

// CreateDistributedAutonomousDatabaseShardDatabaseDetails Details for creating a distributed autonomous database shard.
type CreateDistributedAutonomousDatabaseShardDatabaseDetails interface {
}

type createdistributedautonomousdatabasesharddatabasedetails struct {
	JsonData []byte
	Source   string `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *createdistributedautonomousdatabasesharddatabasedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedistributedautonomousdatabasesharddatabasedetails createdistributedautonomousdatabasesharddatabasedetails
	s := struct {
		Model Unmarshalercreatedistributedautonomousdatabasesharddatabasedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Source = s.Model.Source

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdistributedautonomousdatabasesharddatabasedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "ADBD_EXISTING_CLUSTER":
		mm := CreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateDistributedAutonomousDatabaseShardDatabaseDetails: %s.", m.Source)
		return *m, nil
	}
}

func (m createdistributedautonomousdatabasesharddatabasedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdistributedautonomousdatabasesharddatabasedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceEnum Enum with underlying type: string
type CreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceEnum string

// Set of constants representing the allowable values for CreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceEnum
const (
	CreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceAdbdExistingCluster CreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceEnum = "ADBD_EXISTING_CLUSTER"
)

var mappingCreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceEnum = map[string]CreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceEnum{
	"ADBD_EXISTING_CLUSTER": CreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceAdbdExistingCluster,
}

var mappingCreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceEnumLowerCase = map[string]CreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceEnum{
	"adbd_existing_cluster": CreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceAdbdExistingCluster,
}

// GetCreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceEnumValues Enumerates the set of values for CreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceEnum
func GetCreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceEnumValues() []CreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceEnum {
	values := make([]CreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceEnum, 0)
	for _, v := range mappingCreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceEnumStringValues Enumerates the set of values in String for CreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceEnum
func GetCreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceEnumStringValues() []string {
	return []string{
		"ADBD_EXISTING_CLUSTER",
	}
}

// GetMappingCreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceEnum(val string) (CreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceEnum, bool) {
	enum, ok := mappingCreateDistributedAutonomousDatabaseShardDatabaseDetailsSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
