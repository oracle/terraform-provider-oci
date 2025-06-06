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

// CreateDatabaseToolsConnectionGenericJdbcDetails Details of the new Database Tools connection for a Generic JDBC database system.
type CreateDatabaseToolsConnectionGenericJdbcDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools connection.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The JDBC URL used to connect to the Generic JDBC database system.
	Url *string `mandatory:"true" json:"url"`

	// The user name.
	UserName *string `mandatory:"true" json:"userName"`

	UserPassword DatabaseToolsUserPasswordDetails `mandatory:"true" json:"userPassword"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// The advanced connection properties key-value pair.
	AdvancedProperties map[string]string `mandatory:"false" json:"advancedProperties"`

	// The CA certificate to verify the server's certificate and
	// the client private key and associated certificate required for client authentication.
	KeyStores []DatabaseToolsKeyStoreGenericJdbcDetails `mandatory:"false" json:"keyStores"`

	// Specifies whether this connection is supported by the Database Tools Runtime.
	RuntimeSupport RuntimeSupportEnum `mandatory:"false" json:"runtimeSupport,omitempty"`
}

// GetDisplayName returns DisplayName
func (m CreateDatabaseToolsConnectionGenericJdbcDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m CreateDatabaseToolsConnectionGenericJdbcDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDefinedTags returns DefinedTags
func (m CreateDatabaseToolsConnectionGenericJdbcDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m CreateDatabaseToolsConnectionGenericJdbcDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetLocks returns Locks
func (m CreateDatabaseToolsConnectionGenericJdbcDetails) GetLocks() []ResourceLock {
	return m.Locks
}

// GetRuntimeSupport returns RuntimeSupport
func (m CreateDatabaseToolsConnectionGenericJdbcDetails) GetRuntimeSupport() RuntimeSupportEnum {
	return m.RuntimeSupport
}

func (m CreateDatabaseToolsConnectionGenericJdbcDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDatabaseToolsConnectionGenericJdbcDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRuntimeSupportEnum(string(m.RuntimeSupport)); !ok && m.RuntimeSupport != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuntimeSupport: %s. Supported values are: %s.", m.RuntimeSupport, strings.Join(GetRuntimeSupportEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDatabaseToolsConnectionGenericJdbcDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDatabaseToolsConnectionGenericJdbcDetails CreateDatabaseToolsConnectionGenericJdbcDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateDatabaseToolsConnectionGenericJdbcDetails
	}{
		"GENERIC_JDBC",
		(MarshalTypeCreateDatabaseToolsConnectionGenericJdbcDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateDatabaseToolsConnectionGenericJdbcDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DefinedTags        map[string]map[string]interface{}         `json:"definedTags"`
		FreeformTags       map[string]string                         `json:"freeformTags"`
		Locks              []ResourceLock                            `json:"locks"`
		RuntimeSupport     RuntimeSupportEnum                        `json:"runtimeSupport"`
		AdvancedProperties map[string]string                         `json:"advancedProperties"`
		KeyStores          []DatabaseToolsKeyStoreGenericJdbcDetails `json:"keyStores"`
		DisplayName        *string                                   `json:"displayName"`
		CompartmentId      *string                                   `json:"compartmentId"`
		Url                *string                                   `json:"url"`
		UserName           *string                                   `json:"userName"`
		UserPassword       databasetoolsuserpassworddetails          `json:"userPassword"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.Locks = make([]ResourceLock, len(model.Locks))
	copy(m.Locks, model.Locks)
	m.RuntimeSupport = model.RuntimeSupport

	m.AdvancedProperties = model.AdvancedProperties

	m.KeyStores = make([]DatabaseToolsKeyStoreGenericJdbcDetails, len(model.KeyStores))
	copy(m.KeyStores, model.KeyStores)
	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.Url = model.Url

	m.UserName = model.UserName

	nn, e = model.UserPassword.UnmarshalPolymorphicJSON(model.UserPassword.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.UserPassword = nn.(DatabaseToolsUserPasswordDetails)
	} else {
		m.UserPassword = nil
	}

	return
}
