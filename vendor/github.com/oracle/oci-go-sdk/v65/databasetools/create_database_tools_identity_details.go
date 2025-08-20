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

// CreateDatabaseToolsIdentityDetails Details for the new Database Tools identity.
type CreateDatabaseToolsIdentityDetails interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools identity.
	GetCompartmentId() *string

	// A user-friendly name. Does not have to be unique and can be updated. Avoid entering confidential information.
	GetDisplayName() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Database Tools connection.
	GetDatabaseToolsConnectionId() *string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Locks associated with this resource.
	GetLocks() []ResourceLock
}

type createdatabasetoolsidentitydetails struct {
	JsonData                  []byte
	DefinedTags               map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	FreeformTags              map[string]string                 `mandatory:"false" json:"freeformTags"`
	Locks                     []ResourceLock                    `mandatory:"false" json:"locks"`
	CompartmentId             *string                           `mandatory:"true" json:"compartmentId"`
	DisplayName               *string                           `mandatory:"true" json:"displayName"`
	DatabaseToolsConnectionId *string                           `mandatory:"true" json:"databaseToolsConnectionId"`
	Type                      string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createdatabasetoolsidentitydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedatabasetoolsidentitydetails createdatabasetoolsidentitydetails
	s := struct {
		Model Unmarshalercreatedatabasetoolsidentitydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.DatabaseToolsConnectionId = s.Model.DatabaseToolsConnectionId
	m.DefinedTags = s.Model.DefinedTags
	m.FreeformTags = s.Model.FreeformTags
	m.Locks = s.Model.Locks
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdatabasetoolsidentitydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "ORACLE_DATABASE_RESOURCE_PRINCIPAL":
		mm := CreateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateDatabaseToolsIdentityDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetDefinedTags returns DefinedTags
func (m createdatabasetoolsidentitydetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m createdatabasetoolsidentitydetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetLocks returns Locks
func (m createdatabasetoolsidentitydetails) GetLocks() []ResourceLock {
	return m.Locks
}

// GetCompartmentId returns CompartmentId
func (m createdatabasetoolsidentitydetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m createdatabasetoolsidentitydetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDatabaseToolsConnectionId returns DatabaseToolsConnectionId
func (m createdatabasetoolsidentitydetails) GetDatabaseToolsConnectionId() *string {
	return m.DatabaseToolsConnectionId
}

func (m createdatabasetoolsidentitydetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdatabasetoolsidentitydetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
