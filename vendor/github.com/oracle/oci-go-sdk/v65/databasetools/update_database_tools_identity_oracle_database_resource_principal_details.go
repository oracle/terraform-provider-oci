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

// UpdateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails Database Tools identity information to be updated for the Oracle Database resource principal identity type.
type UpdateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails struct {

	// A user-friendly name. Does not have to be unique and can be updated. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

// GetDisplayName returns DisplayName
func (m UpdateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDefinedTags returns DefinedTags
func (m UpdateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m UpdateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

func (m UpdateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails UpdateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails
	}{
		"ORACLE_DATABASE_RESOURCE_PRINCIPAL",
		(MarshalTypeUpdateDatabaseToolsIdentityOracleDatabaseResourcePrincipalDetails)(m),
	}

	return json.Marshal(&s)
}
