// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Migration Migration resource
type Migration struct {

	// The OCID of the resource
	Id *string `mandatory:"true" json:"id"`

	// Migration Display Name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// OCID of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Migration type.
	Type MigrationTypesEnum `mandatory:"true" json:"type"`

	// The OCID of the Source Database Connection.
	SourceDatabaseConnectionId *string `mandatory:"true" json:"sourceDatabaseConnectionId"`

	// The OCID of the Target Database Connection.
	TargetDatabaseConnectionId *string `mandatory:"true" json:"targetDatabaseConnectionId"`

	// The time the Migration was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the Migration resource.
	LifecycleState MigrationLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// Name of a migration phase. The Job will wait after executing this
	// phase until the Resume Job endpoint is called.
	WaitAfter OdmsJobPhasesEnum `mandatory:"false" json:"waitAfter,omitempty"`

	// The OCID of the registered on-premises ODMS Agent. Only valid for Offline Migrations.
	AgentId *string `mandatory:"false" json:"agentId"`

	// OCID of the Secret in the OCI vault containing the Migration credentials. Used to store GoldenGate administrator user credentials.
	CredentialsSecretId *string `mandatory:"false" json:"credentialsSecretId"`

	// The OCID of the Source Container Database Connection.
	SourceContainerDatabaseConnectionId *string `mandatory:"false" json:"sourceContainerDatabaseConnectionId"`

	// OCID of the current ODMS Job in execution for the Migration, if any.
	ExecutingJobId *string `mandatory:"false" json:"executingJobId"`

	DataTransferMediumDetails *DataTransferMediumDetails `mandatory:"false" json:"dataTransferMediumDetails"`

	DumpTransferDetails *DumpTransferDetails `mandatory:"false" json:"dumpTransferDetails"`

	DatapumpSettings *DataPumpSettings `mandatory:"false" json:"datapumpSettings"`

	AdvisorSettings *AdvisorSettings `mandatory:"false" json:"advisorSettings"`

	// Database objects to exclude from migration.
	// If 'includeObjects' are specified, only exclude object types can be specified with general wildcards (.*) for owner and objectName.
	ExcludeObjects []DatabaseObject `mandatory:"false" json:"excludeObjects"`

	// Database objects to include from migration.
	IncludeObjects []DatabaseObject `mandatory:"false" json:"includeObjects"`

	GoldenGateDetails *GoldenGateDetails `mandatory:"false" json:"goldenGateDetails"`

	VaultDetails *VaultDetails `mandatory:"false" json:"vaultDetails"`

	// The time of the last Migration details update. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The time of last Migration. An RFC3339 formatted datetime string.
	TimeLastMigration *common.SDKTime `mandatory:"false" json:"timeLastMigration"`

	// Additional status related to the execution and current state of the Migration.
	LifecycleDetails MigrationStatusEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Migration) String() string {
	return common.PointerString(m)
}
