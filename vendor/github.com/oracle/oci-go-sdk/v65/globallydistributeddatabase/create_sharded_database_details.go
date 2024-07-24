// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage distributed databases.
//

package globallydistributeddatabase

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateShardedDatabaseDetails Details required for Sharded database creation.
type CreateShardedDatabaseDetails interface {

	// Identifier of the compartment where sharded database is to be created.
	GetCompartmentId() *string

	// Oracle sharded database display name.
	GetDisplayName() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createshardeddatabasedetails struct {
	JsonData         []byte
	FreeformTags     map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	CompartmentId    *string                           `mandatory:"true" json:"compartmentId"`
	DisplayName      *string                           `mandatory:"true" json:"displayName"`
	DbDeploymentType string                            `json:"dbDeploymentType"`
}

// UnmarshalJSON unmarshals json
func (m *createshardeddatabasedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateshardeddatabasedetails createshardeddatabasedetails
	s := struct {
		Model Unmarshalercreateshardeddatabasedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.DbDeploymentType = s.Model.DbDeploymentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createshardeddatabasedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DbDeploymentType {
	case "DEDICATED":
		mm := CreateDedicatedShardedDatabase{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateShardedDatabaseDetails: %s.", m.DbDeploymentType)
		return *m, nil
	}
}

// GetFreeformTags returns FreeformTags
func (m createshardeddatabasedetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createshardeddatabasedetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetCompartmentId returns CompartmentId
func (m createshardeddatabasedetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m createshardeddatabasedetails) GetDisplayName() *string {
	return m.DisplayName
}

func (m createshardeddatabasedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createshardeddatabasedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateShardedDatabaseDetailsDbDeploymentTypeEnum Enum with underlying type: string
type CreateShardedDatabaseDetailsDbDeploymentTypeEnum string

// Set of constants representing the allowable values for CreateShardedDatabaseDetailsDbDeploymentTypeEnum
const (
	CreateShardedDatabaseDetailsDbDeploymentTypeDedicated CreateShardedDatabaseDetailsDbDeploymentTypeEnum = "DEDICATED"
)

var mappingCreateShardedDatabaseDetailsDbDeploymentTypeEnum = map[string]CreateShardedDatabaseDetailsDbDeploymentTypeEnum{
	"DEDICATED": CreateShardedDatabaseDetailsDbDeploymentTypeDedicated,
}

var mappingCreateShardedDatabaseDetailsDbDeploymentTypeEnumLowerCase = map[string]CreateShardedDatabaseDetailsDbDeploymentTypeEnum{
	"dedicated": CreateShardedDatabaseDetailsDbDeploymentTypeDedicated,
}

// GetCreateShardedDatabaseDetailsDbDeploymentTypeEnumValues Enumerates the set of values for CreateShardedDatabaseDetailsDbDeploymentTypeEnum
func GetCreateShardedDatabaseDetailsDbDeploymentTypeEnumValues() []CreateShardedDatabaseDetailsDbDeploymentTypeEnum {
	values := make([]CreateShardedDatabaseDetailsDbDeploymentTypeEnum, 0)
	for _, v := range mappingCreateShardedDatabaseDetailsDbDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateShardedDatabaseDetailsDbDeploymentTypeEnumStringValues Enumerates the set of values in String for CreateShardedDatabaseDetailsDbDeploymentTypeEnum
func GetCreateShardedDatabaseDetailsDbDeploymentTypeEnumStringValues() []string {
	return []string{
		"DEDICATED",
	}
}

// GetMappingCreateShardedDatabaseDetailsDbDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateShardedDatabaseDetailsDbDeploymentTypeEnum(val string) (CreateShardedDatabaseDetailsDbDeploymentTypeEnum, bool) {
	enum, ok := mappingCreateShardedDatabaseDetailsDbDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
