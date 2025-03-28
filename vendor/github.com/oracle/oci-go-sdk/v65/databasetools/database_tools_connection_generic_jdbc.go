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

// DatabaseToolsConnectionGenericJdbc Database Tools connection of a Generic JDBC database system.
type DatabaseToolsConnectionGenericJdbc struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools connection.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools connection.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the Database Tools connection was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the DatabaseToolsConnection was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The JDBC URL used to connect to the Generic JDBC database system.
	Url *string `mandatory:"true" json:"url"`

	// A message describing the current state in more detail. For example, this message can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// The user name.
	UserName *string `mandatory:"false" json:"userName"`

	UserPassword DatabaseToolsUserPassword `mandatory:"false" json:"userPassword"`

	// The advanced connection properties key-value pair.
	AdvancedProperties map[string]string `mandatory:"false" json:"advancedProperties"`

	// The CA certificate to verify the server's certificate and
	// the client private key and associated certificate required for client authentication.
	KeyStores []DatabaseToolsKeyStoreGenericJdbc `mandatory:"false" json:"keyStores"`

	// The current state of the Database Tools connection.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Specifies whether this connection is supported by the Database Tools Runtime.
	RuntimeSupport RuntimeSupportEnum `mandatory:"true" json:"runtimeSupport"`
}

// GetId returns Id
func (m DatabaseToolsConnectionGenericJdbc) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m DatabaseToolsConnectionGenericJdbc) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m DatabaseToolsConnectionGenericJdbc) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetLifecycleState returns LifecycleState
func (m DatabaseToolsConnectionGenericJdbc) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m DatabaseToolsConnectionGenericJdbc) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m DatabaseToolsConnectionGenericJdbc) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DatabaseToolsConnectionGenericJdbc) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetDefinedTags returns DefinedTags
func (m DatabaseToolsConnectionGenericJdbc) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m DatabaseToolsConnectionGenericJdbc) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetSystemTags returns SystemTags
func (m DatabaseToolsConnectionGenericJdbc) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLocks returns Locks
func (m DatabaseToolsConnectionGenericJdbc) GetLocks() []ResourceLock {
	return m.Locks
}

// GetRuntimeSupport returns RuntimeSupport
func (m DatabaseToolsConnectionGenericJdbc) GetRuntimeSupport() RuntimeSupportEnum {
	return m.RuntimeSupport
}

func (m DatabaseToolsConnectionGenericJdbc) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsConnectionGenericJdbc) ValidateEnumValue() (bool, error) {
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

// MarshalJSON marshals to json representation
func (m DatabaseToolsConnectionGenericJdbc) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseToolsConnectionGenericJdbc DatabaseToolsConnectionGenericJdbc
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDatabaseToolsConnectionGenericJdbc
	}{
		"GENERIC_JDBC",
		(MarshalTypeDatabaseToolsConnectionGenericJdbc)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DatabaseToolsConnectionGenericJdbc) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		LifecycleDetails   *string                            `json:"lifecycleDetails"`
		DefinedTags        map[string]map[string]interface{}  `json:"definedTags"`
		FreeformTags       map[string]string                  `json:"freeformTags"`
		SystemTags         map[string]map[string]interface{}  `json:"systemTags"`
		Locks              []ResourceLock                     `json:"locks"`
		UserName           *string                            `json:"userName"`
		UserPassword       databasetoolsuserpassword          `json:"userPassword"`
		AdvancedProperties map[string]string                  `json:"advancedProperties"`
		KeyStores          []DatabaseToolsKeyStoreGenericJdbc `json:"keyStores"`
		Id                 *string                            `json:"id"`
		DisplayName        *string                            `json:"displayName"`
		CompartmentId      *string                            `json:"compartmentId"`
		LifecycleState     LifecycleStateEnum                 `json:"lifecycleState"`
		TimeCreated        *common.SDKTime                    `json:"timeCreated"`
		TimeUpdated        *common.SDKTime                    `json:"timeUpdated"`
		RuntimeSupport     RuntimeSupportEnum                 `json:"runtimeSupport"`
		Url                *string                            `json:"url"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.LifecycleDetails = model.LifecycleDetails

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.SystemTags = model.SystemTags

	m.Locks = make([]ResourceLock, len(model.Locks))
	copy(m.Locks, model.Locks)
	m.UserName = model.UserName

	nn, e = model.UserPassword.UnmarshalPolymorphicJSON(model.UserPassword.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.UserPassword = nn.(DatabaseToolsUserPassword)
	} else {
		m.UserPassword = nil
	}

	m.AdvancedProperties = model.AdvancedProperties

	m.KeyStores = make([]DatabaseToolsKeyStoreGenericJdbc, len(model.KeyStores))
	copy(m.KeyStores, model.KeyStores)
	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.RuntimeSupport = model.RuntimeSupport

	m.Url = model.Url

	return
}
