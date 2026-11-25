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

// DistributedAutonomousDatabaseCatalogDatabase Details of the catalog database associated with the Globally distributed autonomous database.
type DistributedAutonomousDatabaseCatalogDatabase interface {

	// The name of the catalog.
	GetName() *string

	// The time the catalog was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The time the catalog was last updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	GetMetadata() *DistributedAutonomousDbMetadata
}

type distributedautonomousdatabasecatalogdatabase struct {
	JsonData    []byte
	Metadata    *DistributedAutonomousDbMetadata `mandatory:"false" json:"metadata"`
	Name        *string                          `mandatory:"true" json:"name"`
	TimeCreated *common.SDKTime                  `mandatory:"true" json:"timeCreated"`
	TimeUpdated *common.SDKTime                  `mandatory:"true" json:"timeUpdated"`
	Source      string                           `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *distributedautonomousdatabasecatalogdatabase) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdistributedautonomousdatabasecatalogdatabase distributedautonomousdatabasecatalogdatabase
	s := struct {
		Model Unmarshalerdistributedautonomousdatabasecatalogdatabase
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
func (m *distributedautonomousdatabasecatalogdatabase) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "ADBD_EXISTING_CLUSTER":
		mm := DistributedAutonomousDatabaseCatalogWithDedicatedInfra{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DistributedAutonomousDatabaseCatalogDatabase: %s.", m.Source)
		return *m, nil
	}
}

// GetMetadata returns Metadata
func (m distributedautonomousdatabasecatalogdatabase) GetMetadata() *DistributedAutonomousDbMetadata {
	return m.Metadata
}

// GetName returns Name
func (m distributedautonomousdatabasecatalogdatabase) GetName() *string {
	return m.Name
}

// GetTimeCreated returns TimeCreated
func (m distributedautonomousdatabasecatalogdatabase) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m distributedautonomousdatabasecatalogdatabase) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m distributedautonomousdatabasecatalogdatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m distributedautonomousdatabasecatalogdatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DistributedAutonomousDatabaseCatalogDatabaseSourceEnum Enum with underlying type: string
type DistributedAutonomousDatabaseCatalogDatabaseSourceEnum string

// Set of constants representing the allowable values for DistributedAutonomousDatabaseCatalogDatabaseSourceEnum
const (
	DistributedAutonomousDatabaseCatalogDatabaseSourceAdbdExistingCluster DistributedAutonomousDatabaseCatalogDatabaseSourceEnum = "ADBD_EXISTING_CLUSTER"
)

var mappingDistributedAutonomousDatabaseCatalogDatabaseSourceEnum = map[string]DistributedAutonomousDatabaseCatalogDatabaseSourceEnum{
	"ADBD_EXISTING_CLUSTER": DistributedAutonomousDatabaseCatalogDatabaseSourceAdbdExistingCluster,
}

var mappingDistributedAutonomousDatabaseCatalogDatabaseSourceEnumLowerCase = map[string]DistributedAutonomousDatabaseCatalogDatabaseSourceEnum{
	"adbd_existing_cluster": DistributedAutonomousDatabaseCatalogDatabaseSourceAdbdExistingCluster,
}

// GetDistributedAutonomousDatabaseCatalogDatabaseSourceEnumValues Enumerates the set of values for DistributedAutonomousDatabaseCatalogDatabaseSourceEnum
func GetDistributedAutonomousDatabaseCatalogDatabaseSourceEnumValues() []DistributedAutonomousDatabaseCatalogDatabaseSourceEnum {
	values := make([]DistributedAutonomousDatabaseCatalogDatabaseSourceEnum, 0)
	for _, v := range mappingDistributedAutonomousDatabaseCatalogDatabaseSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedAutonomousDatabaseCatalogDatabaseSourceEnumStringValues Enumerates the set of values in String for DistributedAutonomousDatabaseCatalogDatabaseSourceEnum
func GetDistributedAutonomousDatabaseCatalogDatabaseSourceEnumStringValues() []string {
	return []string{
		"ADBD_EXISTING_CLUSTER",
	}
}

// GetMappingDistributedAutonomousDatabaseCatalogDatabaseSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedAutonomousDatabaseCatalogDatabaseSourceEnum(val string) (DistributedAutonomousDatabaseCatalogDatabaseSourceEnum, bool) {
	enum, ok := mappingDistributedAutonomousDatabaseCatalogDatabaseSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
