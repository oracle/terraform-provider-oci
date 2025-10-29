// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// DatabaseToolsIdentity Manages credentials in a database to access service resources.
type DatabaseToolsIdentity interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools identity.
	GetId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools identity.
	GetCompartmentId() *string

	// A user-friendly name. Does not have to be unique and can be updated. Avoid entering confidential information.
	GetDisplayName() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Database Tools connection.
	GetDatabaseToolsConnectionId() *string

	// The current state of the Database Tools identity.
	GetLifecycleState() DatabaseToolsIdentityLifecycleStateEnum

	// The time the Database Tools identity was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The time the Database Tools identity was updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	// A message describing the current state in more detail. For example, this message can be used to provide actionable information for a resource in the Failed state.
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

type databasetoolsidentity struct {
	JsonData                  []byte
	LifecycleDetails          *string                                 `mandatory:"false" json:"lifecycleDetails"`
	DefinedTags               map[string]map[string]interface{}       `mandatory:"false" json:"definedTags"`
	FreeformTags              map[string]string                       `mandatory:"false" json:"freeformTags"`
	SystemTags                map[string]map[string]interface{}       `mandatory:"false" json:"systemTags"`
	Locks                     []ResourceLock                          `mandatory:"false" json:"locks"`
	Id                        *string                                 `mandatory:"true" json:"id"`
	CompartmentId             *string                                 `mandatory:"true" json:"compartmentId"`
	DisplayName               *string                                 `mandatory:"true" json:"displayName"`
	DatabaseToolsConnectionId *string                                 `mandatory:"true" json:"databaseToolsConnectionId"`
	LifecycleState            DatabaseToolsIdentityLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	TimeCreated               *common.SDKTime                         `mandatory:"true" json:"timeCreated"`
	TimeUpdated               *common.SDKTime                         `mandatory:"true" json:"timeUpdated"`
	Type                      string                                  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolsidentity) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolsidentity databasetoolsidentity
	s := struct {
		Model Unmarshalerdatabasetoolsidentity
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.DatabaseToolsConnectionId = s.Model.DatabaseToolsConnectionId
	m.LifecycleState = s.Model.LifecycleState
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.DefinedTags = s.Model.DefinedTags
	m.FreeformTags = s.Model.FreeformTags
	m.SystemTags = s.Model.SystemTags
	m.Locks = s.Model.Locks
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolsidentity) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "ORACLE_DATABASE_RESOURCE_PRINCIPAL":
		mm := DatabaseToolsIdentityOracleDatabaseResourcePrincipal{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DatabaseToolsIdentity: %s.", m.Type)
		return *m, nil
	}
}

// GetLifecycleDetails returns LifecycleDetails
func (m databasetoolsidentity) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetDefinedTags returns DefinedTags
func (m databasetoolsidentity) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m databasetoolsidentity) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetSystemTags returns SystemTags
func (m databasetoolsidentity) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLocks returns Locks
func (m databasetoolsidentity) GetLocks() []ResourceLock {
	return m.Locks
}

// GetId returns Id
func (m databasetoolsidentity) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m databasetoolsidentity) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m databasetoolsidentity) GetDisplayName() *string {
	return m.DisplayName
}

// GetDatabaseToolsConnectionId returns DatabaseToolsConnectionId
func (m databasetoolsidentity) GetDatabaseToolsConnectionId() *string {
	return m.DatabaseToolsConnectionId
}

// GetLifecycleState returns LifecycleState
func (m databasetoolsidentity) GetLifecycleState() DatabaseToolsIdentityLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m databasetoolsidentity) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m databasetoolsidentity) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m databasetoolsidentity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolsidentity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseToolsIdentityLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDatabaseToolsIdentityLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
