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

// UpdateDatabaseToolsConnectionDetails The information to be updated.
type UpdateDatabaseToolsConnectionDetails interface {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	GetDisplayName() *string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Specifies the identity used by the Database Tools service to issue requests to other OCI services (e.g., Secrets in Vault).
	GetRuntimeIdentity() RuntimeIdentityEnum
}

type updatedatabasetoolsconnectiondetails struct {
	JsonData        []byte
	DisplayName     *string                           `mandatory:"false" json:"displayName"`
	DefinedTags     map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	FreeformTags    map[string]string                 `mandatory:"false" json:"freeformTags"`
	RuntimeIdentity RuntimeIdentityEnum               `mandatory:"false" json:"runtimeIdentity,omitempty"`
	Type            string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *updatedatabasetoolsconnectiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatedatabasetoolsconnectiondetails updatedatabasetoolsconnectiondetails
	s := struct {
		Model Unmarshalerupdatedatabasetoolsconnectiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.DefinedTags = s.Model.DefinedTags
	m.FreeformTags = s.Model.FreeformTags
	m.RuntimeIdentity = s.Model.RuntimeIdentity
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatedatabasetoolsconnectiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "GENERIC_JDBC":
		mm := UpdateDatabaseToolsConnectionGenericJdbcDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "POSTGRESQL":
		mm := UpdateDatabaseToolsConnectionPostgresqlDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL":
		mm := UpdateDatabaseToolsConnectionMySqlDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_DATABASE":
		mm := UpdateDatabaseToolsConnectionOracleDatabaseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateDatabaseToolsConnectionDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m updatedatabasetoolsconnectiondetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDefinedTags returns DefinedTags
func (m updatedatabasetoolsconnectiondetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m updatedatabasetoolsconnectiondetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetRuntimeIdentity returns RuntimeIdentity
func (m updatedatabasetoolsconnectiondetails) GetRuntimeIdentity() RuntimeIdentityEnum {
	return m.RuntimeIdentity
}

func (m updatedatabasetoolsconnectiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatedatabasetoolsconnectiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRuntimeIdentityEnum(string(m.RuntimeIdentity)); !ok && m.RuntimeIdentity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuntimeIdentity: %s. Supported values are: %s.", m.RuntimeIdentity, strings.Join(GetRuntimeIdentityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
