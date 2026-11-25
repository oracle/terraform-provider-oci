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

// DistributedDatabaseCatalogDatabase Details of a distributed database catalog.
type DistributedDatabaseCatalogDatabase interface {

	// The name of the catalog.
	GetName() *string

	// The time the catalog was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The time the catalog was last updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	GetMetadata() *DistributedDbMetadata
}

type distributeddatabasecatalogdatabase struct {
	JsonData    []byte
	Metadata    *DistributedDbMetadata `mandatory:"false" json:"metadata"`
	Name        *string                `mandatory:"true" json:"name"`
	TimeCreated *common.SDKTime        `mandatory:"true" json:"timeCreated"`
	TimeUpdated *common.SDKTime        `mandatory:"true" json:"timeUpdated"`
	Source      string                 `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *distributeddatabasecatalogdatabase) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdistributeddatabasecatalogdatabase distributeddatabasecatalogdatabase
	s := struct {
		Model Unmarshalerdistributeddatabasecatalogdatabase
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
func (m *distributeddatabasecatalogdatabase) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "XS_NEW_CLUSTER":
		mm := DistributedDatabaseCatalogDatabaseWithExadbXsNewCluster{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "XD_EXISTING_CLUSTER":
		mm := DistributedDatabaseCatalogDatabaseWithExadbXd{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "XS_NEW_VAULT_AND_CLUSTER":
		mm := DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndCluster{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXISTING_DB_HOME":
		mm := DistributedDatabaseCatalogDatabaseWithDbHome{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "XS_EXISTING_CLUSTER":
		mm := DistributedDatabaseCatalogDatabaseWithExadbXs{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DistributedDatabaseCatalogDatabase: %s.", m.Source)
		return *m, nil
	}
}

// GetMetadata returns Metadata
func (m distributeddatabasecatalogdatabase) GetMetadata() *DistributedDbMetadata {
	return m.Metadata
}

// GetName returns Name
func (m distributeddatabasecatalogdatabase) GetName() *string {
	return m.Name
}

// GetTimeCreated returns TimeCreated
func (m distributeddatabasecatalogdatabase) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m distributeddatabasecatalogdatabase) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m distributeddatabasecatalogdatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m distributeddatabasecatalogdatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DistributedDatabaseCatalogDatabaseSourceEnum Enum with underlying type: string
type DistributedDatabaseCatalogDatabaseSourceEnum string

// Set of constants representing the allowable values for DistributedDatabaseCatalogDatabaseSourceEnum
const (
	DistributedDatabaseCatalogDatabaseSourceXsExistingCluster    DistributedDatabaseCatalogDatabaseSourceEnum = "XS_EXISTING_CLUSTER"
	DistributedDatabaseCatalogDatabaseSourceXsNewCluster         DistributedDatabaseCatalogDatabaseSourceEnum = "XS_NEW_CLUSTER"
	DistributedDatabaseCatalogDatabaseSourceXsNewVaultAndCluster DistributedDatabaseCatalogDatabaseSourceEnum = "XS_NEW_VAULT_AND_CLUSTER"
	DistributedDatabaseCatalogDatabaseSourceXdExistingCluster    DistributedDatabaseCatalogDatabaseSourceEnum = "XD_EXISTING_CLUSTER"
	DistributedDatabaseCatalogDatabaseSourceExistingDbHome       DistributedDatabaseCatalogDatabaseSourceEnum = "EXISTING_DB_HOME"
)

var mappingDistributedDatabaseCatalogDatabaseSourceEnum = map[string]DistributedDatabaseCatalogDatabaseSourceEnum{
	"XS_EXISTING_CLUSTER":      DistributedDatabaseCatalogDatabaseSourceXsExistingCluster,
	"XS_NEW_CLUSTER":           DistributedDatabaseCatalogDatabaseSourceXsNewCluster,
	"XS_NEW_VAULT_AND_CLUSTER": DistributedDatabaseCatalogDatabaseSourceXsNewVaultAndCluster,
	"XD_EXISTING_CLUSTER":      DistributedDatabaseCatalogDatabaseSourceXdExistingCluster,
	"EXISTING_DB_HOME":         DistributedDatabaseCatalogDatabaseSourceExistingDbHome,
}

var mappingDistributedDatabaseCatalogDatabaseSourceEnumLowerCase = map[string]DistributedDatabaseCatalogDatabaseSourceEnum{
	"xs_existing_cluster":      DistributedDatabaseCatalogDatabaseSourceXsExistingCluster,
	"xs_new_cluster":           DistributedDatabaseCatalogDatabaseSourceXsNewCluster,
	"xs_new_vault_and_cluster": DistributedDatabaseCatalogDatabaseSourceXsNewVaultAndCluster,
	"xd_existing_cluster":      DistributedDatabaseCatalogDatabaseSourceXdExistingCluster,
	"existing_db_home":         DistributedDatabaseCatalogDatabaseSourceExistingDbHome,
}

// GetDistributedDatabaseCatalogDatabaseSourceEnumValues Enumerates the set of values for DistributedDatabaseCatalogDatabaseSourceEnum
func GetDistributedDatabaseCatalogDatabaseSourceEnumValues() []DistributedDatabaseCatalogDatabaseSourceEnum {
	values := make([]DistributedDatabaseCatalogDatabaseSourceEnum, 0)
	for _, v := range mappingDistributedDatabaseCatalogDatabaseSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseCatalogDatabaseSourceEnumStringValues Enumerates the set of values in String for DistributedDatabaseCatalogDatabaseSourceEnum
func GetDistributedDatabaseCatalogDatabaseSourceEnumStringValues() []string {
	return []string{
		"XS_EXISTING_CLUSTER",
		"XS_NEW_CLUSTER",
		"XS_NEW_VAULT_AND_CLUSTER",
		"XD_EXISTING_CLUSTER",
		"EXISTING_DB_HOME",
	}
}

// GetMappingDistributedDatabaseCatalogDatabaseSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseCatalogDatabaseSourceEnum(val string) (DistributedDatabaseCatalogDatabaseSourceEnum, bool) {
	enum, ok := mappingDistributedDatabaseCatalogDatabaseSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
