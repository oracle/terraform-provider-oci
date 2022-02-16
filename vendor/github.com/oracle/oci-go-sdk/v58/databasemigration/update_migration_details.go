// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// UpdateMigrationDetails Update Migration resource parameters.
type UpdateMigrationDetails struct {

	// Migration type.
	Type MigrationTypesEnum `mandatory:"false" json:"type,omitempty"`

	// Migration Display Name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the registered ODMS Agent.
	AgentId *string `mandatory:"false" json:"agentId"`

	// The OCID of the Source Database Connection.
	SourceDatabaseConnectionId *string `mandatory:"false" json:"sourceDatabaseConnectionId"`

	// The OCID of the Source Container Database Connection. Only used for Online migrations.
	// Only Connections of type Non-Autonomous can be used as source container databases.
	// An empty value would remove the stored Connection ID.
	SourceContainerDatabaseConnectionId *string `mandatory:"false" json:"sourceContainerDatabaseConnectionId"`

	// The OCID of the Target Database Connection.
	TargetDatabaseConnectionId *string `mandatory:"false" json:"targetDatabaseConnectionId"`

	DataTransferMediumDetails *UpdateDataTransferMediumDetails `mandatory:"false" json:"dataTransferMediumDetails"`

	DumpTransferDetails *UpdateDumpTransferDetails `mandatory:"false" json:"dumpTransferDetails"`

	DatapumpSettings *UpdateDataPumpSettings `mandatory:"false" json:"datapumpSettings"`

	AdvisorSettings *UpdateAdvisorSettings `mandatory:"false" json:"advisorSettings"`

	// Database objects to exclude from migration, cannot be specified alongside 'includeObjects'.
	// If specified, the list will be replaced entirely. Empty list will remove stored excludeObjects details.
	ExcludeObjects []DatabaseObject `mandatory:"false" json:"excludeObjects"`

	// Database objects to include from migration, cannot be specified alongside 'excludeObjects'.
	// If specified, the list will be replaced entirely. Empty list will remove stored includeObjects details.
	IncludeObjects []DatabaseObject `mandatory:"false" json:"includeObjects"`

	GoldenGateDetails *UpdateGoldenGateDetails `mandatory:"false" json:"goldenGateDetails"`

	VaultDetails *UpdateVaultDetails `mandatory:"false" json:"vaultDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateMigrationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateMigrationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMigrationTypesEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMigrationTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
