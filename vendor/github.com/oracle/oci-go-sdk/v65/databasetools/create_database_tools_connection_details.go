// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CreateDatabaseToolsConnectionDetails Details for the new Database Tools connection.
type CreateDatabaseToolsConnectionDetails interface {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	GetDisplayName() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools connection.
	GetCompartmentId() *string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Locks associated with this resource.
	GetLocks() []ResourceLock

	// Specifies whether this connection is supported by the Database Tools Runtime.
	GetRuntimeSupport() RuntimeSupportEnum
}

type createdatabasetoolsconnectiondetails struct {
	JsonData       []byte
	DefinedTags    map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	FreeformTags   map[string]string                 `mandatory:"false" json:"freeformTags"`
	Locks          []ResourceLock                    `mandatory:"false" json:"locks"`
	RuntimeSupport RuntimeSupportEnum                `mandatory:"false" json:"runtimeSupport,omitempty"`
	DisplayName    *string                           `mandatory:"true" json:"displayName"`
	CompartmentId  *string                           `mandatory:"true" json:"compartmentId"`
	Type           string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createdatabasetoolsconnectiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedatabasetoolsconnectiondetails createdatabasetoolsconnectiondetails
	s := struct {
		Model Unmarshalercreatedatabasetoolsconnectiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.DefinedTags = s.Model.DefinedTags
	m.FreeformTags = s.Model.FreeformTags
	m.Locks = s.Model.Locks
	m.RuntimeSupport = s.Model.RuntimeSupport
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdatabasetoolsconnectiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "GENERIC_JDBC":
		mm := CreateDatabaseToolsConnectionGenericJdbcDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "POSTGRESQL":
		mm := CreateDatabaseToolsConnectionPostgresqlDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL":
		mm := CreateDatabaseToolsConnectionMySqlDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_DATABASE":
		mm := CreateDatabaseToolsConnectionOracleDatabaseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateDatabaseToolsConnectionDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetDefinedTags returns DefinedTags
func (m createdatabasetoolsconnectiondetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m createdatabasetoolsconnectiondetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetLocks returns Locks
func (m createdatabasetoolsconnectiondetails) GetLocks() []ResourceLock {
	return m.Locks
}

// GetRuntimeSupport returns RuntimeSupport
func (m createdatabasetoolsconnectiondetails) GetRuntimeSupport() RuntimeSupportEnum {
	return m.RuntimeSupport
}

// GetDisplayName returns DisplayName
func (m createdatabasetoolsconnectiondetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m createdatabasetoolsconnectiondetails) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m createdatabasetoolsconnectiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdatabasetoolsconnectiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRuntimeSupportEnum(string(m.RuntimeSupport)); !ok && m.RuntimeSupport != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuntimeSupport: %s. Supported values are: %s.", m.RuntimeSupport, strings.Join(GetRuntimeSupportEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
