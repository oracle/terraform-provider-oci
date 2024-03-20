// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MySqlCloneMigrationDetails MySQL Clone Migration Summary
type MySqlCloneMigrationDetails struct {

	// OCI resource ID.
	SourceDatabaseConnectionId *string `mandatory:"true" json:"sourceDatabaseConnectionId"`

	// OCI resource ID.
	TargetDatabaseConnectionId *string `mandatory:"true" json:"targetDatabaseConnectionId"`

	// An object's Display Name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// OCI resource ID.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

// GetDisplayName returns DisplayName
func (m MySqlCloneMigrationDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m MySqlCloneMigrationDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetSourceDatabaseConnectionId returns SourceDatabaseConnectionId
func (m MySqlCloneMigrationDetails) GetSourceDatabaseConnectionId() *string {
	return m.SourceDatabaseConnectionId
}

// GetTargetDatabaseConnectionId returns TargetDatabaseConnectionId
func (m MySqlCloneMigrationDetails) GetTargetDatabaseConnectionId() *string {
	return m.TargetDatabaseConnectionId
}

// GetFreeformTags returns FreeformTags
func (m MySqlCloneMigrationDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m MySqlCloneMigrationDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m MySqlCloneMigrationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MySqlCloneMigrationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MySqlCloneMigrationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMySqlCloneMigrationDetails MySqlCloneMigrationDetails
	s := struct {
		DiscriminatorParam string `json:"databaseCombination"`
		MarshalTypeMySqlCloneMigrationDetails
	}{
		"MYSQL",
		(MarshalTypeMySqlCloneMigrationDetails)(m),
	}

	return json.Marshal(&s)
}
