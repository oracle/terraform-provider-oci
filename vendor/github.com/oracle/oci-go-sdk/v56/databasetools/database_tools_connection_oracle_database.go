// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Database Tools APIs to manage Connections and Private Endpoints.
//

package databasetools

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DatabaseToolsConnectionOracleDatabase DatabaseToolsConnection of an Oracle Database.
type DatabaseToolsConnectionOracleDatabase struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DatabaseToolsConnection.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the containing Compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the DatabaseToolsConnection was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the DatabaseToolsConnection was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
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

	RelatedResource *DatabaseToolsRelatedResource `mandatory:"false" json:"relatedResource"`

	// Connect descriptor or Easy Connect Naming method to connect to the database.
	ConnectionString *string `mandatory:"false" json:"connectionString"`

	// Database user name.
	UserName *string `mandatory:"false" json:"userName"`

	UserPassword DatabaseToolsUserPassword `mandatory:"false" json:"userPassword"`

	// Advanced connection properties key-value pair (e.g., oracle.net.ssl_server_dn_match).
	AdvancedProperties map[string]string `mandatory:"false" json:"advancedProperties"`

	// Oracle wallet or Java Keystores containing trusted certificates for authenticating the server's public certificate and
	// the client private key and associated certificates required for client authentication.
	KeyStores []DatabaseToolsKeyStore `mandatory:"false" json:"keyStores"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DatabaseToolsPrivateEndpoint used to access the database in the Customer VCN.
	PrivateEndpointId *string `mandatory:"false" json:"privateEndpointId"`

	// The current state of the DatabaseToolsConnection.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

//GetId returns Id
func (m DatabaseToolsConnectionOracleDatabase) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m DatabaseToolsConnectionOracleDatabase) GetDisplayName() *string {
	return m.DisplayName
}

//GetCompartmentId returns CompartmentId
func (m DatabaseToolsConnectionOracleDatabase) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetLifecycleState returns LifecycleState
func (m DatabaseToolsConnectionOracleDatabase) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m DatabaseToolsConnectionOracleDatabase) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetTimeCreated returns TimeCreated
func (m DatabaseToolsConnectionOracleDatabase) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m DatabaseToolsConnectionOracleDatabase) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetDefinedTags returns DefinedTags
func (m DatabaseToolsConnectionOracleDatabase) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetFreeformTags returns FreeformTags
func (m DatabaseToolsConnectionOracleDatabase) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetSystemTags returns SystemTags
func (m DatabaseToolsConnectionOracleDatabase) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m DatabaseToolsConnectionOracleDatabase) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DatabaseToolsConnectionOracleDatabase) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseToolsConnectionOracleDatabase DatabaseToolsConnectionOracleDatabase
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDatabaseToolsConnectionOracleDatabase
	}{
		"ORACLE_DATABASE",
		(MarshalTypeDatabaseToolsConnectionOracleDatabase)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DatabaseToolsConnectionOracleDatabase) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		LifecycleDetails   *string                           `json:"lifecycleDetails"`
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		SystemTags         map[string]map[string]interface{} `json:"systemTags"`
		RelatedResource    *DatabaseToolsRelatedResource     `json:"relatedResource"`
		ConnectionString   *string                           `json:"connectionString"`
		UserName           *string                           `json:"userName"`
		UserPassword       databasetoolsuserpassword         `json:"userPassword"`
		AdvancedProperties map[string]string                 `json:"advancedProperties"`
		KeyStores          []DatabaseToolsKeyStore           `json:"keyStores"`
		PrivateEndpointId  *string                           `json:"privateEndpointId"`
		Id                 *string                           `json:"id"`
		DisplayName        *string                           `json:"displayName"`
		CompartmentId      *string                           `json:"compartmentId"`
		LifecycleState     LifecycleStateEnum                `json:"lifecycleState"`
		TimeCreated        *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated        *common.SDKTime                   `json:"timeUpdated"`
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

	m.RelatedResource = model.RelatedResource

	m.ConnectionString = model.ConnectionString

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

	m.KeyStores = make([]DatabaseToolsKeyStore, len(model.KeyStores))
	for i, n := range model.KeyStores {
		m.KeyStores[i] = n
	}

	m.PrivateEndpointId = model.PrivateEndpointId

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	return
}
