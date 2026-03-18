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

// DistributedAutonomousDatabaseCatalog Globally distributed autonomous database catalog.
type DistributedAutonomousDatabaseCatalog interface {

	// The name of catalog.
	GetName() *string

	// The time the catalog was created. An RFC3339 formatted datetime string
	GetTimeCreated() *common.SDKTime

	// The time the catalog was last updated. An RFC3339 formatted datetime string
	GetTimeUpdated() *common.SDKTime
}

type distributedautonomousdatabasecatalog struct {
	JsonData    []byte
	Name        *string         `mandatory:"true" json:"name"`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`
	Source      string          `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *distributedautonomousdatabasecatalog) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdistributedautonomousdatabasecatalog distributedautonomousdatabasecatalog
	s := struct {
		Model Unmarshalerdistributedautonomousdatabasecatalog
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
func (m *distributedautonomousdatabasecatalog) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "ADB_D":
		mm := DistributedAutonomousDatabaseCatalogWithDedicatedInfra{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DistributedAutonomousDatabaseCatalog: %s.", m.Source)
		return *m, nil
	}
}

// GetName returns Name
func (m distributedautonomousdatabasecatalog) GetName() *string {
	return m.Name
}

// GetTimeCreated returns TimeCreated
func (m distributedautonomousdatabasecatalog) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m distributedautonomousdatabasecatalog) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m distributedautonomousdatabasecatalog) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m distributedautonomousdatabasecatalog) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DistributedAutonomousDatabaseCatalogSourceEnum Enum with underlying type: string
type DistributedAutonomousDatabaseCatalogSourceEnum string

// Set of constants representing the allowable values for DistributedAutonomousDatabaseCatalogSourceEnum
const (
	DistributedAutonomousDatabaseCatalogSourceAdbD DistributedAutonomousDatabaseCatalogSourceEnum = "ADB_D"
)

var mappingDistributedAutonomousDatabaseCatalogSourceEnum = map[string]DistributedAutonomousDatabaseCatalogSourceEnum{
	"ADB_D": DistributedAutonomousDatabaseCatalogSourceAdbD,
}

var mappingDistributedAutonomousDatabaseCatalogSourceEnumLowerCase = map[string]DistributedAutonomousDatabaseCatalogSourceEnum{
	"adb_d": DistributedAutonomousDatabaseCatalogSourceAdbD,
}

// GetDistributedAutonomousDatabaseCatalogSourceEnumValues Enumerates the set of values for DistributedAutonomousDatabaseCatalogSourceEnum
func GetDistributedAutonomousDatabaseCatalogSourceEnumValues() []DistributedAutonomousDatabaseCatalogSourceEnum {
	values := make([]DistributedAutonomousDatabaseCatalogSourceEnum, 0)
	for _, v := range mappingDistributedAutonomousDatabaseCatalogSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedAutonomousDatabaseCatalogSourceEnumStringValues Enumerates the set of values in String for DistributedAutonomousDatabaseCatalogSourceEnum
func GetDistributedAutonomousDatabaseCatalogSourceEnumStringValues() []string {
	return []string{
		"ADB_D",
	}
}

// GetMappingDistributedAutonomousDatabaseCatalogSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedAutonomousDatabaseCatalogSourceEnum(val string) (DistributedAutonomousDatabaseCatalogSourceEnum, bool) {
	enum, ok := mappingDistributedAutonomousDatabaseCatalogSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
