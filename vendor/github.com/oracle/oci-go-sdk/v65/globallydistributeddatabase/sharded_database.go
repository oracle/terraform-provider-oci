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

// ShardedDatabase Sharded Database resource.
type ShardedDatabase interface {

	// Sharded Database identifier
	GetId() *string

	// Identifier of the compartment in which sharded database exists.
	GetCompartmentId() *string

	// Oracle sharded database display name.
	GetDisplayName() *string

	// The time the the Sharded Database was created. An RFC3339 formatted datetime string
	GetTimeCreated() *common.SDKTime

	// The time the Sharded Database was last updated. An RFC3339 formatted datetime string
	GetTimeUpdated() *common.SDKTime

	// Lifecycle states for sharded databases.
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

type shardeddatabase struct {
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
func (m *shardeddatabase) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalershardeddatabase shardeddatabase
	s := struct {
		Model Unmarshalershardeddatabase
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
func (m *shardeddatabase) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DbDeploymentType {
	case "DEDICATED":
		mm := DedicatedShardedDatabase{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ShardedDatabase: %s.", m.DbDeploymentType)
		return *m, nil
	}
}

// GetFreeformTags returns FreeformTags
func (m shardeddatabase) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m shardeddatabase) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m shardeddatabase) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m shardeddatabase) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m shardeddatabase) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m shardeddatabase) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeCreated returns TimeCreated
func (m shardeddatabase) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m shardeddatabase) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m shardeddatabase) GetLifecycleState() ShardedDatabaseLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleStateDetails returns LifecycleStateDetails
func (m shardeddatabase) GetLifecycleStateDetails() *string {
	return m.LifecycleStateDetails
}

func (m shardeddatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m shardeddatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingShardedDatabaseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetShardedDatabaseLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ShardedDatabaseLifecycleStateEnum Enum with underlying type: string
type ShardedDatabaseLifecycleStateEnum string

// Set of constants representing the allowable values for ShardedDatabaseLifecycleStateEnum
const (
	ShardedDatabaseLifecycleStateActive         ShardedDatabaseLifecycleStateEnum = "ACTIVE"
	ShardedDatabaseLifecycleStateFailed         ShardedDatabaseLifecycleStateEnum = "FAILED"
	ShardedDatabaseLifecycleStateNeedsAttention ShardedDatabaseLifecycleStateEnum = "NEEDS_ATTENTION"
	ShardedDatabaseLifecycleStateInactive       ShardedDatabaseLifecycleStateEnum = "INACTIVE"
	ShardedDatabaseLifecycleStateDeleting       ShardedDatabaseLifecycleStateEnum = "DELETING"
	ShardedDatabaseLifecycleStateDeleted        ShardedDatabaseLifecycleStateEnum = "DELETED"
	ShardedDatabaseLifecycleStateUpdating       ShardedDatabaseLifecycleStateEnum = "UPDATING"
	ShardedDatabaseLifecycleStateCreating       ShardedDatabaseLifecycleStateEnum = "CREATING"
	ShardedDatabaseLifecycleStateUnavailable    ShardedDatabaseLifecycleStateEnum = "UNAVAILABLE"
)

var mappingShardedDatabaseLifecycleStateEnum = map[string]ShardedDatabaseLifecycleStateEnum{
	"ACTIVE":          ShardedDatabaseLifecycleStateActive,
	"FAILED":          ShardedDatabaseLifecycleStateFailed,
	"NEEDS_ATTENTION": ShardedDatabaseLifecycleStateNeedsAttention,
	"INACTIVE":        ShardedDatabaseLifecycleStateInactive,
	"DELETING":        ShardedDatabaseLifecycleStateDeleting,
	"DELETED":         ShardedDatabaseLifecycleStateDeleted,
	"UPDATING":        ShardedDatabaseLifecycleStateUpdating,
	"CREATING":        ShardedDatabaseLifecycleStateCreating,
	"UNAVAILABLE":     ShardedDatabaseLifecycleStateUnavailable,
}

var mappingShardedDatabaseLifecycleStateEnumLowerCase = map[string]ShardedDatabaseLifecycleStateEnum{
	"active":          ShardedDatabaseLifecycleStateActive,
	"failed":          ShardedDatabaseLifecycleStateFailed,
	"needs_attention": ShardedDatabaseLifecycleStateNeedsAttention,
	"inactive":        ShardedDatabaseLifecycleStateInactive,
	"deleting":        ShardedDatabaseLifecycleStateDeleting,
	"deleted":         ShardedDatabaseLifecycleStateDeleted,
	"updating":        ShardedDatabaseLifecycleStateUpdating,
	"creating":        ShardedDatabaseLifecycleStateCreating,
	"unavailable":     ShardedDatabaseLifecycleStateUnavailable,
}

// GetShardedDatabaseLifecycleStateEnumValues Enumerates the set of values for ShardedDatabaseLifecycleStateEnum
func GetShardedDatabaseLifecycleStateEnumValues() []ShardedDatabaseLifecycleStateEnum {
	values := make([]ShardedDatabaseLifecycleStateEnum, 0)
	for _, v := range mappingShardedDatabaseLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetShardedDatabaseLifecycleStateEnumStringValues Enumerates the set of values in String for ShardedDatabaseLifecycleStateEnum
func GetShardedDatabaseLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"FAILED",
		"NEEDS_ATTENTION",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"UPDATING",
		"CREATING",
		"UNAVAILABLE",
	}
}

// GetMappingShardedDatabaseLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingShardedDatabaseLifecycleStateEnum(val string) (ShardedDatabaseLifecycleStateEnum, bool) {
	enum, ok := mappingShardedDatabaseLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ShardedDatabaseDbDeploymentTypeEnum Enum with underlying type: string
type ShardedDatabaseDbDeploymentTypeEnum string

// Set of constants representing the allowable values for ShardedDatabaseDbDeploymentTypeEnum
const (
	ShardedDatabaseDbDeploymentTypeDedicated ShardedDatabaseDbDeploymentTypeEnum = "DEDICATED"
)

var mappingShardedDatabaseDbDeploymentTypeEnum = map[string]ShardedDatabaseDbDeploymentTypeEnum{
	"DEDICATED": ShardedDatabaseDbDeploymentTypeDedicated,
}

var mappingShardedDatabaseDbDeploymentTypeEnumLowerCase = map[string]ShardedDatabaseDbDeploymentTypeEnum{
	"dedicated": ShardedDatabaseDbDeploymentTypeDedicated,
}

// GetShardedDatabaseDbDeploymentTypeEnumValues Enumerates the set of values for ShardedDatabaseDbDeploymentTypeEnum
func GetShardedDatabaseDbDeploymentTypeEnumValues() []ShardedDatabaseDbDeploymentTypeEnum {
	values := make([]ShardedDatabaseDbDeploymentTypeEnum, 0)
	for _, v := range mappingShardedDatabaseDbDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetShardedDatabaseDbDeploymentTypeEnumStringValues Enumerates the set of values in String for ShardedDatabaseDbDeploymentTypeEnum
func GetShardedDatabaseDbDeploymentTypeEnumStringValues() []string {
	return []string{
		"DEDICATED",
	}
}

// GetMappingShardedDatabaseDbDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingShardedDatabaseDbDeploymentTypeEnum(val string) (ShardedDatabaseDbDeploymentTypeEnum, bool) {
	enum, ok := mappingShardedDatabaseDbDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
