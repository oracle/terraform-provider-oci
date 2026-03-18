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

// DistributedDatabaseCatalog Globally distributed database catalog.
type DistributedDatabaseCatalog interface {

	// The name of catalog.
	GetName() *string

	// The time the catalog was created. An RFC3339 formatted datetime string
	GetTimeCreated() *common.SDKTime

	// The time the catalog was last updated. An RFC3339 formatted datetime string
	GetTimeUpdated() *common.SDKTime
}

type distributeddatabasecatalog struct {
	JsonData    []byte
	Name        *string         `mandatory:"true" json:"name"`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`
	Source      string          `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *distributeddatabasecatalog) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdistributeddatabasecatalog distributeddatabasecatalog
	s := struct {
		Model Unmarshalerdistributeddatabasecatalog
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
func (m *distributeddatabasecatalog) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "EXADB_XS":
		mm := DistributedDatabaseCatalogWithExadbXs{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NEW_VAULT_AND_CLUSTER":
		mm := DistributedDatabaseCatalogWithExadbXsNewVaultAndCluster{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DistributedDatabaseCatalog: %s.", m.Source)
		return *m, nil
	}
}

// GetName returns Name
func (m distributeddatabasecatalog) GetName() *string {
	return m.Name
}

// GetTimeCreated returns TimeCreated
func (m distributeddatabasecatalog) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m distributeddatabasecatalog) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m distributeddatabasecatalog) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m distributeddatabasecatalog) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DistributedDatabaseCatalogSourceEnum Enum with underlying type: string
type DistributedDatabaseCatalogSourceEnum string

// Set of constants representing the allowable values for DistributedDatabaseCatalogSourceEnum
const (
	DistributedDatabaseCatalogSourceExadbXs            DistributedDatabaseCatalogSourceEnum = "EXADB_XS"
	DistributedDatabaseCatalogSourceNewVaultAndCluster DistributedDatabaseCatalogSourceEnum = "NEW_VAULT_AND_CLUSTER"
	DistributedDatabaseCatalogSourceExistingCluster    DistributedDatabaseCatalogSourceEnum = "EXISTING_CLUSTER"
)

var mappingDistributedDatabaseCatalogSourceEnum = map[string]DistributedDatabaseCatalogSourceEnum{
	"EXADB_XS":              DistributedDatabaseCatalogSourceExadbXs,
	"NEW_VAULT_AND_CLUSTER": DistributedDatabaseCatalogSourceNewVaultAndCluster,
	"EXISTING_CLUSTER":      DistributedDatabaseCatalogSourceExistingCluster,
}

var mappingDistributedDatabaseCatalogSourceEnumLowerCase = map[string]DistributedDatabaseCatalogSourceEnum{
	"exadb_xs":              DistributedDatabaseCatalogSourceExadbXs,
	"new_vault_and_cluster": DistributedDatabaseCatalogSourceNewVaultAndCluster,
	"existing_cluster":      DistributedDatabaseCatalogSourceExistingCluster,
}

// GetDistributedDatabaseCatalogSourceEnumValues Enumerates the set of values for DistributedDatabaseCatalogSourceEnum
func GetDistributedDatabaseCatalogSourceEnumValues() []DistributedDatabaseCatalogSourceEnum {
	values := make([]DistributedDatabaseCatalogSourceEnum, 0)
	for _, v := range mappingDistributedDatabaseCatalogSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseCatalogSourceEnumStringValues Enumerates the set of values in String for DistributedDatabaseCatalogSourceEnum
func GetDistributedDatabaseCatalogSourceEnumStringValues() []string {
	return []string{
		"EXADB_XS",
		"NEW_VAULT_AND_CLUSTER",
		"EXISTING_CLUSTER",
	}
}

// GetMappingDistributedDatabaseCatalogSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseCatalogSourceEnum(val string) (DistributedDatabaseCatalogSourceEnum, bool) {
	enum, ok := mappingDistributedDatabaseCatalogSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
