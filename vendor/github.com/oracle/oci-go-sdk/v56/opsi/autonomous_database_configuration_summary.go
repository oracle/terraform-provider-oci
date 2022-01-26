// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// AutonomousDatabaseConfigurationSummary Configuration Summary of autonomous database.
type AutonomousDatabaseConfigurationSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database insight resource.
	DatabaseInsightId *string `mandatory:"true" json:"databaseInsightId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The database name. The database name is unique within the tenancy.
	DatabaseName *string `mandatory:"true" json:"databaseName"`

	// The user-friendly name for the database. The name does not have to be unique.
	DatabaseDisplayName *string `mandatory:"true" json:"databaseDisplayName"`

	// Operations Insights internal representation of the database type.
	DatabaseType *string `mandatory:"true" json:"databaseType"`

	// The version of the database.
	DatabaseVersion *string `mandatory:"true" json:"databaseVersion"`

	// Name of the CDB.Only applies to PDB.
	CdbName *string `mandatory:"true" json:"cdbName"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
	DatabaseId *string `mandatory:"true" json:"databaseId"`

	// Processor count. This is the OCPU count for Autonomous Database and CPU core count for other database types.
	ProcessorCount *int `mandatory:"false" json:"processorCount"`
}

//GetDatabaseInsightId returns DatabaseInsightId
func (m AutonomousDatabaseConfigurationSummary) GetDatabaseInsightId() *string {
	return m.DatabaseInsightId
}

//GetCompartmentId returns CompartmentId
func (m AutonomousDatabaseConfigurationSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDatabaseName returns DatabaseName
func (m AutonomousDatabaseConfigurationSummary) GetDatabaseName() *string {
	return m.DatabaseName
}

//GetDatabaseDisplayName returns DatabaseDisplayName
func (m AutonomousDatabaseConfigurationSummary) GetDatabaseDisplayName() *string {
	return m.DatabaseDisplayName
}

//GetDatabaseType returns DatabaseType
func (m AutonomousDatabaseConfigurationSummary) GetDatabaseType() *string {
	return m.DatabaseType
}

//GetDatabaseVersion returns DatabaseVersion
func (m AutonomousDatabaseConfigurationSummary) GetDatabaseVersion() *string {
	return m.DatabaseVersion
}

//GetCdbName returns CdbName
func (m AutonomousDatabaseConfigurationSummary) GetCdbName() *string {
	return m.CdbName
}

//GetDefinedTags returns DefinedTags
func (m AutonomousDatabaseConfigurationSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetFreeformTags returns FreeformTags
func (m AutonomousDatabaseConfigurationSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetProcessorCount returns ProcessorCount
func (m AutonomousDatabaseConfigurationSummary) GetProcessorCount() *int {
	return m.ProcessorCount
}

func (m AutonomousDatabaseConfigurationSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AutonomousDatabaseConfigurationSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAutonomousDatabaseConfigurationSummary AutonomousDatabaseConfigurationSummary
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeAutonomousDatabaseConfigurationSummary
	}{
		"AUTONOMOUS_DATABASE",
		(MarshalTypeAutonomousDatabaseConfigurationSummary)(m),
	}

	return json.Marshal(&s)
}
