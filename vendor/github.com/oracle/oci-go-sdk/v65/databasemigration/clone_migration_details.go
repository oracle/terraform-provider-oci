// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CloneMigrationDetails Details that will override an existing Migration configuration that will be cloned.
type CloneMigrationDetails struct {

	// The OCID of the Source Database Connection.
	SourceDatabaseConnectionId *string `mandatory:"true" json:"sourceDatabaseConnectionId"`

	// The OCID of the Target Database Connection.
	TargetDatabaseConnectionId *string `mandatory:"true" json:"targetDatabaseConnectionId"`

	// Migration Display Name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// OCID of the compartment
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID of the registered on-premises ODMS Agent. Only valid for Offline Logical Migrations.
	AgentId *string `mandatory:"false" json:"agentId"`

	// The OCID of the Source Container Database Connection. Only used for Online migrations.
	// Only Connections of type Non-Autonomous can be used as source container databases.
	SourceContainerDatabaseConnectionId *string `mandatory:"false" json:"sourceContainerDatabaseConnectionId"`

	// Database objects to exclude from migration, cannot be specified alongside 'includeObjects'
	ExcludeObjects []DatabaseObject `mandatory:"false" json:"excludeObjects"`

	// Database objects to include from migration, cannot be specified alongside 'excludeObjects'
	IncludeObjects []DatabaseObject `mandatory:"false" json:"includeObjects"`

	VaultDetails *CreateVaultDetails `mandatory:"false" json:"vaultDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CloneMigrationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloneMigrationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
