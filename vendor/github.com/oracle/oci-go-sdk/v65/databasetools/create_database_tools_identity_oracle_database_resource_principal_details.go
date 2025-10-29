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

// CreateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails Details for the new Database Tools identity for the Oracle Database resource principal identity type.
type CreateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools identity.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique and can be updated. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Database Tools connection.
	DatabaseToolsConnectionId *string `mandatory:"true" json:"databaseToolsConnectionId"`

	// The name of the credential object created in the Oracle Database.
	CredentialKey *string `mandatory:"true" json:"credentialKey"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`
}

// GetCompartmentId returns CompartmentId
func (m CreateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m CreateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDatabaseToolsConnectionId returns DatabaseToolsConnectionId
func (m CreateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails) GetDatabaseToolsConnectionId() *string {
	return m.DatabaseToolsConnectionId
}

// GetDefinedTags returns DefinedTags
func (m CreateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m CreateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetLocks returns Locks
func (m CreateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails) GetLocks() []ResourceLock {
	return m.Locks
}

func (m CreateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails CreateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails
	}{
		"ORACLE_DATABASE_RESOURCE_PRINCIPAL",
		(MarshalTypeCreateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails)(m),
	}

	return json.Marshal(&s)
}
