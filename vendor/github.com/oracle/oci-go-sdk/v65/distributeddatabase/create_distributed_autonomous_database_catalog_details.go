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

// CreateDistributedAutonomousDatabaseCatalogDetails Details of the Globally distributed autonomous database catalog.
type CreateDistributedAutonomousDatabaseCatalogDetails interface {
}

type createdistributedautonomousdatabasecatalogdetails struct {
	JsonData []byte
	Source   string `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *createdistributedautonomousdatabasecatalogdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedistributedautonomousdatabasecatalogdetails createdistributedautonomousdatabasecatalogdetails
	s := struct {
		Model Unmarshalercreatedistributedautonomousdatabasecatalogdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Source = s.Model.Source

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdistributedautonomousdatabasecatalogdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "ADB_D":
		mm := CreateDistributedAutonomousDatabaseCatalogWithDedicatedInfraDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateDistributedAutonomousDatabaseCatalogDetails: %s.", m.Source)
		return *m, nil
	}
}

func (m createdistributedautonomousdatabasecatalogdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdistributedautonomousdatabasecatalogdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDistributedAutonomousDatabaseCatalogDetailsSourceEnum Enum with underlying type: string
type CreateDistributedAutonomousDatabaseCatalogDetailsSourceEnum string

// Set of constants representing the allowable values for CreateDistributedAutonomousDatabaseCatalogDetailsSourceEnum
const (
	CreateDistributedAutonomousDatabaseCatalogDetailsSourceAdbD CreateDistributedAutonomousDatabaseCatalogDetailsSourceEnum = "ADB_D"
)

var mappingCreateDistributedAutonomousDatabaseCatalogDetailsSourceEnum = map[string]CreateDistributedAutonomousDatabaseCatalogDetailsSourceEnum{
	"ADB_D": CreateDistributedAutonomousDatabaseCatalogDetailsSourceAdbD,
}

var mappingCreateDistributedAutonomousDatabaseCatalogDetailsSourceEnumLowerCase = map[string]CreateDistributedAutonomousDatabaseCatalogDetailsSourceEnum{
	"adb_d": CreateDistributedAutonomousDatabaseCatalogDetailsSourceAdbD,
}

// GetCreateDistributedAutonomousDatabaseCatalogDetailsSourceEnumValues Enumerates the set of values for CreateDistributedAutonomousDatabaseCatalogDetailsSourceEnum
func GetCreateDistributedAutonomousDatabaseCatalogDetailsSourceEnumValues() []CreateDistributedAutonomousDatabaseCatalogDetailsSourceEnum {
	values := make([]CreateDistributedAutonomousDatabaseCatalogDetailsSourceEnum, 0)
	for _, v := range mappingCreateDistributedAutonomousDatabaseCatalogDetailsSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDistributedAutonomousDatabaseCatalogDetailsSourceEnumStringValues Enumerates the set of values in String for CreateDistributedAutonomousDatabaseCatalogDetailsSourceEnum
func GetCreateDistributedAutonomousDatabaseCatalogDetailsSourceEnumStringValues() []string {
	return []string{
		"ADB_D",
	}
}

// GetMappingCreateDistributedAutonomousDatabaseCatalogDetailsSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDistributedAutonomousDatabaseCatalogDetailsSourceEnum(val string) (CreateDistributedAutonomousDatabaseCatalogDetailsSourceEnum, bool) {
	enum, ok := mappingCreateDistributedAutonomousDatabaseCatalogDetailsSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
