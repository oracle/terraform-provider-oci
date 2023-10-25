// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseToolsConnectionSummary Summary of the Database Tools connection.
type DatabaseToolsConnectionSummary interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the `DatabaseToolsConnection`.
	GetId() *string

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	GetDisplayName() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools connection.
	GetCompartmentId() *string

	// The current state of the Database Tools connection.
	GetLifecycleState() LifecycleStateEnum

	// The time the Database Tools connection was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The time the Database Tools connection was updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	// Specifies whether this connection is supported by the Database Tools Runtime.
	GetRuntimeSupport() RuntimeSupportEnum

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}

	// Locks associated with this resource.
	GetLocks() []ResourceLock
}

type databasetoolsconnectionsummary struct {
	JsonData         []byte
	LifecycleDetails *string                           `mandatory:"false" json:"lifecycleDetails"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	FreeformTags     map[string]string                 `mandatory:"false" json:"freeformTags"`
	SystemTags       map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Locks            []ResourceLock                    `mandatory:"false" json:"locks"`
	Id               *string                           `mandatory:"true" json:"id"`
	DisplayName      *string                           `mandatory:"true" json:"displayName"`
	CompartmentId    *string                           `mandatory:"true" json:"compartmentId"`
	LifecycleState   LifecycleStateEnum                `mandatory:"true" json:"lifecycleState"`
	TimeCreated      *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	TimeUpdated      *common.SDKTime                   `mandatory:"true" json:"timeUpdated"`
	RuntimeSupport   RuntimeSupportEnum                `mandatory:"true" json:"runtimeSupport"`
	Type             string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolsconnectionsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolsconnectionsummary databasetoolsconnectionsummary
	s := struct {
		Model Unmarshalerdatabasetoolsconnectionsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.LifecycleState = s.Model.LifecycleState
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.RuntimeSupport = s.Model.RuntimeSupport
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.DefinedTags = s.Model.DefinedTags
	m.FreeformTags = s.Model.FreeformTags
	m.SystemTags = s.Model.SystemTags
	m.Locks = s.Model.Locks
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolsconnectionsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "POSTGRESQL":
		mm := DatabaseToolsConnectionPostgresqlSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_DATABASE":
		mm := DatabaseToolsConnectionOracleDatabaseSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL":
		mm := DatabaseToolsConnectionMySqlSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GENERIC_JDBC":
		mm := DatabaseToolsConnectionGenericJdbcSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DatabaseToolsConnectionSummary: %s.", m.Type)
		return *m, nil
	}
}

// GetLifecycleDetails returns LifecycleDetails
func (m databasetoolsconnectionsummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetDefinedTags returns DefinedTags
func (m databasetoolsconnectionsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m databasetoolsconnectionsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetSystemTags returns SystemTags
func (m databasetoolsconnectionsummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLocks returns Locks
func (m databasetoolsconnectionsummary) GetLocks() []ResourceLock {
	return m.Locks
}

// GetId returns Id
func (m databasetoolsconnectionsummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m databasetoolsconnectionsummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m databasetoolsconnectionsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetLifecycleState returns LifecycleState
func (m databasetoolsconnectionsummary) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m databasetoolsconnectionsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m databasetoolsconnectionsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetRuntimeSupport returns RuntimeSupport
func (m databasetoolsconnectionsummary) GetRuntimeSupport() RuntimeSupportEnum {
	return m.RuntimeSupport
}

func (m databasetoolsconnectionsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolsconnectionsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRuntimeSupportEnum(string(m.RuntimeSupport)); !ok && m.RuntimeSupport != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuntimeSupport: %s. Supported values are: %s.", m.RuntimeSupport, strings.Join(GetRuntimeSupportEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
