// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools API
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

// CreateDatabaseToolsConnectionOracleDatabaseDetails Details of the new Database Tools connection for an Oracle Database.
type CreateDatabaseToolsConnectionOracleDatabaseDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools connection.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The connect descriptor or Easy Connect Naming method use to connect to the database.
	ConnectionString *string `mandatory:"true" json:"connectionString"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	RelatedResource *CreateDatabaseToolsRelatedResourceDetails `mandatory:"false" json:"relatedResource"`

	// The database user name.
	UserName *string `mandatory:"false" json:"userName"`

	UserPassword DatabaseToolsUserPasswordDetails `mandatory:"false" json:"userPassword"`

	// The advanced connection properties key-value pair (e.g., `oracle.net.ssl_server_dn_match`).
	AdvancedProperties map[string]string `mandatory:"false" json:"advancedProperties"`

	// Oracle wallet or Java Keystores containing trusted certificates for authenticating the server's public certificate and
	// the client private key and associated certificates required for client authentication.
	KeyStores []DatabaseToolsKeyStoreDetails `mandatory:"false" json:"keyStores"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools private endpoint used to access the database in the customer VCN.
	PrivateEndpointId *string `mandatory:"false" json:"privateEndpointId"`

	ProxyClient DatabaseToolsConnectionOracleDatabaseProxyClientDetails `mandatory:"false" json:"proxyClient"`

	// Specifies whether this connection is supported by the Database Tools Runtime.
	RuntimeSupport RuntimeSupportEnum `mandatory:"false" json:"runtimeSupport,omitempty"`

	// Specifies the identity used when accessing OCI resources at runtime. AUTHENTICATED_PRINCIPAL to use the caller’s identity (On-Behalf-Of token), or RESOURCE_PRINCIPAL to use the connection’s resource principal (RPST).
	RuntimeIdentity RuntimeIdentityEnum `mandatory:"false" json:"runtimeIdentity,omitempty"`

	// Specifies the authentication type used by the Database Tools service to authenticate with the database.
	AuthenticationType AuthenticationTypeEnum `mandatory:"false" json:"authenticationType,omitempty"`
}

// GetDisplayName returns DisplayName
func (m CreateDatabaseToolsConnectionOracleDatabaseDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m CreateDatabaseToolsConnectionOracleDatabaseDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDefinedTags returns DefinedTags
func (m CreateDatabaseToolsConnectionOracleDatabaseDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m CreateDatabaseToolsConnectionOracleDatabaseDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetLocks returns Locks
func (m CreateDatabaseToolsConnectionOracleDatabaseDetails) GetLocks() []ResourceLock {
	return m.Locks
}

// GetRuntimeSupport returns RuntimeSupport
func (m CreateDatabaseToolsConnectionOracleDatabaseDetails) GetRuntimeSupport() RuntimeSupportEnum {
	return m.RuntimeSupport
}

// GetRuntimeIdentity returns RuntimeIdentity
func (m CreateDatabaseToolsConnectionOracleDatabaseDetails) GetRuntimeIdentity() RuntimeIdentityEnum {
	return m.RuntimeIdentity
}

func (m CreateDatabaseToolsConnectionOracleDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDatabaseToolsConnectionOracleDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRuntimeSupportEnum(string(m.RuntimeSupport)); !ok && m.RuntimeSupport != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuntimeSupport: %s. Supported values are: %s.", m.RuntimeSupport, strings.Join(GetRuntimeSupportEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRuntimeIdentityEnum(string(m.RuntimeIdentity)); !ok && m.RuntimeIdentity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuntimeIdentity: %s. Supported values are: %s.", m.RuntimeIdentity, strings.Join(GetRuntimeIdentityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAuthenticationTypeEnum(string(m.AuthenticationType)); !ok && m.AuthenticationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticationType: %s. Supported values are: %s.", m.AuthenticationType, strings.Join(GetAuthenticationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDatabaseToolsConnectionOracleDatabaseDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDatabaseToolsConnectionOracleDatabaseDetails CreateDatabaseToolsConnectionOracleDatabaseDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateDatabaseToolsConnectionOracleDatabaseDetails
	}{
		"ORACLE_DATABASE",
		(MarshalTypeCreateDatabaseToolsConnectionOracleDatabaseDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateDatabaseToolsConnectionOracleDatabaseDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DefinedTags        map[string]map[string]interface{}                       `json:"definedTags"`
		FreeformTags       map[string]string                                       `json:"freeformTags"`
		Locks              []ResourceLock                                          `json:"locks"`
		RuntimeSupport     RuntimeSupportEnum                                      `json:"runtimeSupport"`
		RuntimeIdentity    RuntimeIdentityEnum                                     `json:"runtimeIdentity"`
		RelatedResource    *CreateDatabaseToolsRelatedResourceDetails              `json:"relatedResource"`
		UserName           *string                                                 `json:"userName"`
		UserPassword       databasetoolsuserpassworddetails                        `json:"userPassword"`
		AdvancedProperties map[string]string                                       `json:"advancedProperties"`
		KeyStores          []DatabaseToolsKeyStoreDetails                          `json:"keyStores"`
		PrivateEndpointId  *string                                                 `json:"privateEndpointId"`
		ProxyClient        databasetoolsconnectionoracledatabaseproxyclientdetails `json:"proxyClient"`
		AuthenticationType AuthenticationTypeEnum                                  `json:"authenticationType"`
		DisplayName        *string                                                 `json:"displayName"`
		CompartmentId      *string                                                 `json:"compartmentId"`
		ConnectionString   *string                                                 `json:"connectionString"`
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

	m.RuntimeIdentity = model.RuntimeIdentity

	m.RelatedResource = model.RelatedResource

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
	copy(m.KeyStores, model.KeyStores)
	m.PrivateEndpointId = model.PrivateEndpointId

	nn, e = model.ProxyClient.UnmarshalPolymorphicJSON(model.ProxyClient.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ProxyClient = nn.(DatabaseToolsConnectionOracleDatabaseProxyClientDetails)
	} else {
		m.ProxyClient = nil
	}

	m.AuthenticationType = model.AuthenticationType

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.ConnectionString = model.ConnectionString

	return
}
