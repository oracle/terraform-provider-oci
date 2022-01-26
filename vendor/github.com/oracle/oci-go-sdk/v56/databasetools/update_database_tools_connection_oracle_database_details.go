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

// UpdateDatabaseToolsConnectionOracleDatabaseDetails The information about new DatabaseToolsConnection for an Oracle Database.
type UpdateDatabaseToolsConnectionOracleDatabaseDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	RelatedResource *UpdateDatabaseToolsRelatedResourceDetails `mandatory:"false" json:"relatedResource"`

	// Connect descriptor or Easy Connect Naming method to connect to the database.
	ConnectionString *string `mandatory:"false" json:"connectionString"`

	// Database user name.
	UserName *string `mandatory:"false" json:"userName"`

	UserPassword DatabaseToolsUserPasswordDetails `mandatory:"false" json:"userPassword"`

	// Advanced connection properties key-value pair (e.g., oracle.net.ssl_server_dn_match).
	AdvancedProperties map[string]string `mandatory:"false" json:"advancedProperties"`

	// Oracle wallet or Java Keystores containing trusted certificates for authenticating the server's public certificate and
	// the client private key and associated certificates required for client authentication.
	KeyStores []DatabaseToolsKeyStoreDetails `mandatory:"false" json:"keyStores"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DatabaseToolsPrivateEndpoint used to access the database in the Customer VCN.
	PrivateEndpointId *string `mandatory:"false" json:"privateEndpointId"`
}

//GetDisplayName returns DisplayName
func (m UpdateDatabaseToolsConnectionOracleDatabaseDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDefinedTags returns DefinedTags
func (m UpdateDatabaseToolsConnectionOracleDatabaseDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetFreeformTags returns FreeformTags
func (m UpdateDatabaseToolsConnectionOracleDatabaseDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

func (m UpdateDatabaseToolsConnectionOracleDatabaseDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateDatabaseToolsConnectionOracleDatabaseDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateDatabaseToolsConnectionOracleDatabaseDetails UpdateDatabaseToolsConnectionOracleDatabaseDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateDatabaseToolsConnectionOracleDatabaseDetails
	}{
		"ORACLE_DATABASE",
		(MarshalTypeUpdateDatabaseToolsConnectionOracleDatabaseDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateDatabaseToolsConnectionOracleDatabaseDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName        *string                                    `json:"displayName"`
		DefinedTags        map[string]map[string]interface{}          `json:"definedTags"`
		FreeformTags       map[string]string                          `json:"freeformTags"`
		RelatedResource    *UpdateDatabaseToolsRelatedResourceDetails `json:"relatedResource"`
		ConnectionString   *string                                    `json:"connectionString"`
		UserName           *string                                    `json:"userName"`
		UserPassword       databasetoolsuserpassworddetails           `json:"userPassword"`
		AdvancedProperties map[string]string                          `json:"advancedProperties"`
		KeyStores          []DatabaseToolsKeyStoreDetails             `json:"keyStores"`
		PrivateEndpointId  *string                                    `json:"privateEndpointId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.RelatedResource = model.RelatedResource

	m.ConnectionString = model.ConnectionString

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

	m.AdvancedProperties = model.AdvancedProperties

	m.KeyStores = make([]DatabaseToolsKeyStoreDetails, len(model.KeyStores))
	for i, n := range model.KeyStores {
		m.KeyStores[i] = n
	}

	m.PrivateEndpointId = model.PrivateEndpointId

	return
}
