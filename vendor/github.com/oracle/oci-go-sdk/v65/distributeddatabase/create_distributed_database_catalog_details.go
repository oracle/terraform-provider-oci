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

// CreateDistributedDatabaseCatalogDetails Details of the Globally distributed database catalog.
type CreateDistributedDatabaseCatalogDetails interface {
}

type createdistributeddatabasecatalogdetails struct {
	JsonData []byte
	Source   string `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *createdistributeddatabasecatalogdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedistributeddatabasecatalogdetails createdistributeddatabasecatalogdetails
	s := struct {
		Model Unmarshalercreatedistributeddatabasecatalogdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Source = s.Model.Source

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdistributeddatabasecatalogdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "NEW_VAULT_AND_CLUSTER":
		mm := CreateDistributedDatabaseCatalogWithExadbXsNewVaultAndClusterDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXADB_XS":
		mm := CreateDistributedDatabaseCatalogWithExadbXsDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateDistributedDatabaseCatalogDetails: %s.", m.Source)
		return *m, nil
	}
}

func (m createdistributeddatabasecatalogdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdistributeddatabasecatalogdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDistributedDatabaseCatalogDetailsSourceEnum Enum with underlying type: string
type CreateDistributedDatabaseCatalogDetailsSourceEnum string

// Set of constants representing the allowable values for CreateDistributedDatabaseCatalogDetailsSourceEnum
const (
	CreateDistributedDatabaseCatalogDetailsSourceExadbXs            CreateDistributedDatabaseCatalogDetailsSourceEnum = "EXADB_XS"
	CreateDistributedDatabaseCatalogDetailsSourceNewVaultAndCluster CreateDistributedDatabaseCatalogDetailsSourceEnum = "NEW_VAULT_AND_CLUSTER"
	CreateDistributedDatabaseCatalogDetailsSourceExistingCluster    CreateDistributedDatabaseCatalogDetailsSourceEnum = "EXISTING_CLUSTER"
)

var mappingCreateDistributedDatabaseCatalogDetailsSourceEnum = map[string]CreateDistributedDatabaseCatalogDetailsSourceEnum{
	"EXADB_XS":              CreateDistributedDatabaseCatalogDetailsSourceExadbXs,
	"NEW_VAULT_AND_CLUSTER": CreateDistributedDatabaseCatalogDetailsSourceNewVaultAndCluster,
	"EXISTING_CLUSTER":      CreateDistributedDatabaseCatalogDetailsSourceExistingCluster,
}

var mappingCreateDistributedDatabaseCatalogDetailsSourceEnumLowerCase = map[string]CreateDistributedDatabaseCatalogDetailsSourceEnum{
	"exadb_xs":              CreateDistributedDatabaseCatalogDetailsSourceExadbXs,
	"new_vault_and_cluster": CreateDistributedDatabaseCatalogDetailsSourceNewVaultAndCluster,
	"existing_cluster":      CreateDistributedDatabaseCatalogDetailsSourceExistingCluster,
}

// GetCreateDistributedDatabaseCatalogDetailsSourceEnumValues Enumerates the set of values for CreateDistributedDatabaseCatalogDetailsSourceEnum
func GetCreateDistributedDatabaseCatalogDetailsSourceEnumValues() []CreateDistributedDatabaseCatalogDetailsSourceEnum {
	values := make([]CreateDistributedDatabaseCatalogDetailsSourceEnum, 0)
	for _, v := range mappingCreateDistributedDatabaseCatalogDetailsSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDistributedDatabaseCatalogDetailsSourceEnumStringValues Enumerates the set of values in String for CreateDistributedDatabaseCatalogDetailsSourceEnum
func GetCreateDistributedDatabaseCatalogDetailsSourceEnumStringValues() []string {
	return []string{
		"EXADB_XS",
		"NEW_VAULT_AND_CLUSTER",
		"EXISTING_CLUSTER",
	}
}

// GetMappingCreateDistributedDatabaseCatalogDetailsSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDistributedDatabaseCatalogDetailsSourceEnum(val string) (CreateDistributedDatabaseCatalogDetailsSourceEnum, bool) {
	enum, ok := mappingCreateDistributedDatabaseCatalogDetailsSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
