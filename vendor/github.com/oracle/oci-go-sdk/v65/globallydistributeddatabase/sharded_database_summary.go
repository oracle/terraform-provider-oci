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

// ShardedDatabaseSummary Sharded Database resource summary.
type ShardedDatabaseSummary interface {

	// Sharded Database identifier
	GetId() *string

	// Identifier of the compartment where sharded database exists.
	GetCompartmentId() *string

	// Oracle sharded database display name.
	GetDisplayName() *string

	// The time the the Sharded Database was created. An RFC3339 formatted datetime string
	GetTimeCreated() *common.SDKTime

	// The time the Sharded Database was last updated. An RFC3339 formatted datetime string
	GetTimeUpdated() *common.SDKTime

	// Lifecycle state of sharded database.
	GetLifecycleState() ShardedDatabaseLifecycleStateEnum

	// Detailed message for the lifecycle state.
	GetLifecycleStateDetails() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type shardeddatabasesummary struct {
	JsonData              []byte
	FreeformTags          map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags           map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags            map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id                    *string                           `mandatory:"true" json:"id"`
	CompartmentId         *string                           `mandatory:"true" json:"compartmentId"`
	DisplayName           *string                           `mandatory:"true" json:"displayName"`
	TimeCreated           *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	TimeUpdated           *common.SDKTime                   `mandatory:"true" json:"timeUpdated"`
	LifecycleState        ShardedDatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	LifecycleStateDetails *string                           `mandatory:"true" json:"lifecycleStateDetails"`
	DbDeploymentType      string                            `json:"dbDeploymentType"`
}

// UnmarshalJSON unmarshals json
func (m *shardeddatabasesummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalershardeddatabasesummary shardeddatabasesummary
	s := struct {
		Model Unmarshalershardeddatabasesummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleState = s.Model.LifecycleState
	m.LifecycleStateDetails = s.Model.LifecycleStateDetails
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.DbDeploymentType = s.Model.DbDeploymentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *shardeddatabasesummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DbDeploymentType {
	case "DEDICATED":
		mm := DedicatedShardedDatabaseSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ShardedDatabaseSummary: %s.", m.DbDeploymentType)
		return *m, nil
	}
}

// GetFreeformTags returns FreeformTags
func (m shardeddatabasesummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m shardeddatabasesummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m shardeddatabasesummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m shardeddatabasesummary) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m shardeddatabasesummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m shardeddatabasesummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeCreated returns TimeCreated
func (m shardeddatabasesummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m shardeddatabasesummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m shardeddatabasesummary) GetLifecycleState() ShardedDatabaseLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleStateDetails returns LifecycleStateDetails
func (m shardeddatabasesummary) GetLifecycleStateDetails() *string {
	return m.LifecycleStateDetails
}

func (m shardeddatabasesummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m shardeddatabasesummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingShardedDatabaseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetShardedDatabaseLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ShardedDatabaseSummaryDbDeploymentTypeEnum Enum with underlying type: string
type ShardedDatabaseSummaryDbDeploymentTypeEnum string

// Set of constants representing the allowable values for ShardedDatabaseSummaryDbDeploymentTypeEnum
const (
	ShardedDatabaseSummaryDbDeploymentTypeDedicated ShardedDatabaseSummaryDbDeploymentTypeEnum = "DEDICATED"
)

var mappingShardedDatabaseSummaryDbDeploymentTypeEnum = map[string]ShardedDatabaseSummaryDbDeploymentTypeEnum{
	"DEDICATED": ShardedDatabaseSummaryDbDeploymentTypeDedicated,
}

var mappingShardedDatabaseSummaryDbDeploymentTypeEnumLowerCase = map[string]ShardedDatabaseSummaryDbDeploymentTypeEnum{
	"dedicated": ShardedDatabaseSummaryDbDeploymentTypeDedicated,
}

// GetShardedDatabaseSummaryDbDeploymentTypeEnumValues Enumerates the set of values for ShardedDatabaseSummaryDbDeploymentTypeEnum
func GetShardedDatabaseSummaryDbDeploymentTypeEnumValues() []ShardedDatabaseSummaryDbDeploymentTypeEnum {
	values := make([]ShardedDatabaseSummaryDbDeploymentTypeEnum, 0)
	for _, v := range mappingShardedDatabaseSummaryDbDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetShardedDatabaseSummaryDbDeploymentTypeEnumStringValues Enumerates the set of values in String for ShardedDatabaseSummaryDbDeploymentTypeEnum
func GetShardedDatabaseSummaryDbDeploymentTypeEnumStringValues() []string {
	return []string{
		"DEDICATED",
	}
}

// GetMappingShardedDatabaseSummaryDbDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingShardedDatabaseSummaryDbDeploymentTypeEnum(val string) (ShardedDatabaseSummaryDbDeploymentTypeEnum, bool) {
	enum, ok := mappingShardedDatabaseSummaryDbDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
