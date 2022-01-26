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

// MacsManagedExternalDatabaseInsightSummary Summary of a database insight resource.
type MacsManagedExternalDatabaseInsightSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database insight resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
	DatabaseId *string `mandatory:"true" json:"databaseId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The database name. The database name is unique within the tenancy.
	DatabaseName *string `mandatory:"false" json:"databaseName"`

	// The user-friendly name for the database. The name does not have to be unique.
	DatabaseDisplayName *string `mandatory:"false" json:"databaseDisplayName"`

	// Operations Insights internal representation of the database type.
	DatabaseType *string `mandatory:"false" json:"databaseType"`

	// The version of the database.
	DatabaseVersion *string `mandatory:"false" json:"databaseVersion"`

	// The hostnames for the database.
	DatabaseHostNames []string `mandatory:"false" json:"databaseHostNames"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Processor count. This is the OCPU count for Autonomous Database and CPU core count for other database types.
	ProcessorCount *int `mandatory:"false" json:"processorCount"`

	// The time the the database insight was first enabled. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the database insight was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// OCI database resource type
	DatabaseResourceType *string `mandatory:"false" json:"databaseResourceType"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent
	ManagementAgentId *string `mandatory:"false" json:"managementAgentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of External Database Connector
	ConnectorId *string `mandatory:"false" json:"connectorId"`

	// Indicates the status of a database insight in Operations Insights
	Status ResourceStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The current state of the database.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m MacsManagedExternalDatabaseInsightSummary) GetId() *string {
	return m.Id
}

//GetDatabaseId returns DatabaseId
func (m MacsManagedExternalDatabaseInsightSummary) GetDatabaseId() *string {
	return m.DatabaseId
}

//GetCompartmentId returns CompartmentId
func (m MacsManagedExternalDatabaseInsightSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDatabaseName returns DatabaseName
func (m MacsManagedExternalDatabaseInsightSummary) GetDatabaseName() *string {
	return m.DatabaseName
}

//GetDatabaseDisplayName returns DatabaseDisplayName
func (m MacsManagedExternalDatabaseInsightSummary) GetDatabaseDisplayName() *string {
	return m.DatabaseDisplayName
}

//GetDatabaseType returns DatabaseType
func (m MacsManagedExternalDatabaseInsightSummary) GetDatabaseType() *string {
	return m.DatabaseType
}

//GetDatabaseVersion returns DatabaseVersion
func (m MacsManagedExternalDatabaseInsightSummary) GetDatabaseVersion() *string {
	return m.DatabaseVersion
}

//GetDatabaseHostNames returns DatabaseHostNames
func (m MacsManagedExternalDatabaseInsightSummary) GetDatabaseHostNames() []string {
	return m.DatabaseHostNames
}

//GetFreeformTags returns FreeformTags
func (m MacsManagedExternalDatabaseInsightSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m MacsManagedExternalDatabaseInsightSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m MacsManagedExternalDatabaseInsightSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

//GetProcessorCount returns ProcessorCount
func (m MacsManagedExternalDatabaseInsightSummary) GetProcessorCount() *int {
	return m.ProcessorCount
}

//GetStatus returns Status
func (m MacsManagedExternalDatabaseInsightSummary) GetStatus() ResourceStatusEnum {
	return m.Status
}

//GetTimeCreated returns TimeCreated
func (m MacsManagedExternalDatabaseInsightSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m MacsManagedExternalDatabaseInsightSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m MacsManagedExternalDatabaseInsightSummary) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m MacsManagedExternalDatabaseInsightSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m MacsManagedExternalDatabaseInsightSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m MacsManagedExternalDatabaseInsightSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMacsManagedExternalDatabaseInsightSummary MacsManagedExternalDatabaseInsightSummary
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeMacsManagedExternalDatabaseInsightSummary
	}{
		"MACS_MANAGED_EXTERNAL_DATABASE",
		(MarshalTypeMacsManagedExternalDatabaseInsightSummary)(m),
	}

	return json.Marshal(&s)
}
