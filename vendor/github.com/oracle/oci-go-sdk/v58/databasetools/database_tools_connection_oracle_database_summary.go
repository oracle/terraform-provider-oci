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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// DatabaseToolsConnectionOracleDatabaseSummary DatabaseToolsConnectionSummary of an Oracle Database.
type DatabaseToolsConnectionOracleDatabaseSummary struct {

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

	UserPassword DatabaseToolsUserPasswordSummary `mandatory:"false" json:"userPassword"`

	// Advanced connection properties key-value pair (e.g., oracle.net.ssl_server_dn_match).
	AdvancedProperties map[string]string `mandatory:"false" json:"advancedProperties"`

	// Oracle wallet or Java Keystores containing trusted certificates for authenticating the server's public certificate and
	// the client private key and associated certificates required for client authentication.
	KeyStores []DatabaseToolsKeyStoreSummary `mandatory:"false" json:"keyStores"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DatabaseToolsPrivateEndpoint used to access the database in the Customer VCN.
	PrivateEndpointId *string `mandatory:"false" json:"privateEndpointId"`

	// The current state of the DatabaseToolsConnection.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

//GetId returns Id
func (m DatabaseToolsConnectionOracleDatabaseSummary) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m DatabaseToolsConnectionOracleDatabaseSummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetCompartmentId returns CompartmentId
func (m DatabaseToolsConnectionOracleDatabaseSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetLifecycleState returns LifecycleState
func (m DatabaseToolsConnectionOracleDatabaseSummary) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m DatabaseToolsConnectionOracleDatabaseSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetTimeCreated returns TimeCreated
func (m DatabaseToolsConnectionOracleDatabaseSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m DatabaseToolsConnectionOracleDatabaseSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetDefinedTags returns DefinedTags
func (m DatabaseToolsConnectionOracleDatabaseSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetFreeformTags returns FreeformTags
func (m DatabaseToolsConnectionOracleDatabaseSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetSystemTags returns SystemTags
func (m DatabaseToolsConnectionOracleDatabaseSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m DatabaseToolsConnectionOracleDatabaseSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsConnectionOracleDatabaseSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DatabaseToolsConnectionOracleDatabaseSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseToolsConnectionOracleDatabaseSummary DatabaseToolsConnectionOracleDatabaseSummary
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDatabaseToolsConnectionOracleDatabaseSummary
	}{
		"ORACLE_DATABASE",
		(MarshalTypeDatabaseToolsConnectionOracleDatabaseSummary)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DatabaseToolsConnectionOracleDatabaseSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		LifecycleDetails   *string                           `json:"lifecycleDetails"`
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		SystemTags         map[string]map[string]interface{} `json:"systemTags"`
		RelatedResource    *DatabaseToolsRelatedResource     `json:"relatedResource"`
		ConnectionString   *string                           `json:"connectionString"`
		UserName           *string                           `json:"userName"`
		UserPassword       databasetoolsuserpasswordsummary  `json:"userPassword"`
		AdvancedProperties map[string]string                 `json:"advancedProperties"`
		KeyStores          []DatabaseToolsKeyStoreSummary    `json:"keyStores"`
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
		m.UserPassword = nn.(DatabaseToolsUserPasswordSummary)
	} else {
		m.UserPassword = nil
	}

	m.AdvancedProperties = model.AdvancedProperties

	m.KeyStores = make([]DatabaseToolsKeyStoreSummary, len(model.KeyStores))
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
