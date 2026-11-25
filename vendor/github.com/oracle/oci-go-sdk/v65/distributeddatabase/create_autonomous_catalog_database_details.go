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

// CreateAutonomousCatalogDatabaseDetails Details for creating a distributed autonomous database catalog.
type CreateAutonomousCatalogDatabaseDetails interface {
}

type createautonomouscatalogdatabasedetails struct {
	JsonData []byte
	Source   string `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *createautonomouscatalogdatabasedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateautonomouscatalogdatabasedetails createautonomouscatalogdatabasedetails
	s := struct {
		Model Unmarshalercreateautonomouscatalogdatabasedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Source = s.Model.Source

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createautonomouscatalogdatabasedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "ADBD_EXISTING_CLUSTER":
		mm := CreateDistributedAutonomousDatabaseCatalogWithDedicatedInfraDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateAutonomousCatalogDatabaseDetails: %s.", m.Source)
		return *m, nil
	}
}

func (m createautonomouscatalogdatabasedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createautonomouscatalogdatabasedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateAutonomousCatalogDatabaseDetailsSourceEnum Enum with underlying type: string
type CreateAutonomousCatalogDatabaseDetailsSourceEnum string

// Set of constants representing the allowable values for CreateAutonomousCatalogDatabaseDetailsSourceEnum
const (
	CreateAutonomousCatalogDatabaseDetailsSourceAdbdExistingCluster CreateAutonomousCatalogDatabaseDetailsSourceEnum = "ADBD_EXISTING_CLUSTER"
)

var mappingCreateAutonomousCatalogDatabaseDetailsSourceEnum = map[string]CreateAutonomousCatalogDatabaseDetailsSourceEnum{
	"ADBD_EXISTING_CLUSTER": CreateAutonomousCatalogDatabaseDetailsSourceAdbdExistingCluster,
}

var mappingCreateAutonomousCatalogDatabaseDetailsSourceEnumLowerCase = map[string]CreateAutonomousCatalogDatabaseDetailsSourceEnum{
	"adbd_existing_cluster": CreateAutonomousCatalogDatabaseDetailsSourceAdbdExistingCluster,
}

// GetCreateAutonomousCatalogDatabaseDetailsSourceEnumValues Enumerates the set of values for CreateAutonomousCatalogDatabaseDetailsSourceEnum
func GetCreateAutonomousCatalogDatabaseDetailsSourceEnumValues() []CreateAutonomousCatalogDatabaseDetailsSourceEnum {
	values := make([]CreateAutonomousCatalogDatabaseDetailsSourceEnum, 0)
	for _, v := range mappingCreateAutonomousCatalogDatabaseDetailsSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousCatalogDatabaseDetailsSourceEnumStringValues Enumerates the set of values in String for CreateAutonomousCatalogDatabaseDetailsSourceEnum
func GetCreateAutonomousCatalogDatabaseDetailsSourceEnumStringValues() []string {
	return []string{
		"ADBD_EXISTING_CLUSTER",
	}
}

// GetMappingCreateAutonomousCatalogDatabaseDetailsSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousCatalogDatabaseDetailsSourceEnum(val string) (CreateAutonomousCatalogDatabaseDetailsSourceEnum, bool) {
	enum, ok := mappingCreateAutonomousCatalogDatabaseDetailsSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
