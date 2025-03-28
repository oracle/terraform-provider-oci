// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MacsManagedCloudDatabaseInsight Database insight resource.
type MacsManagedCloudDatabaseInsight struct {

	// Database insight identifier
	Id *string `mandatory:"true" json:"id"`

	// Compartment identifier of the database
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The time the the database insight was first enabled. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
	DatabaseId *string `mandatory:"true" json:"databaseId"`

	// Name of database
	DatabaseName *string `mandatory:"true" json:"databaseName"`

	// OCI database resource type
	DatabaseResourceType *string `mandatory:"true" json:"databaseResourceType"`

	// Ops Insights internal representation of the database type.
	DatabaseType *string `mandatory:"false" json:"databaseType"`

	// The version of the database.
	DatabaseVersion *string `mandatory:"false" json:"databaseVersion"`

	// Processor count. This is the OCPU count for Autonomous Database and CPU core count for other database types.
	ProcessorCount *int `mandatory:"false" json:"processorCount"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The time the database insight was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// A message describing the status of the database connection of this resource. For example, it can be used to provide actionable information about the permission and content validity of the database connection.
	DatabaseConnectionStatusDetails *string `mandatory:"false" json:"databaseConnectionStatusDetails"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent
	ManagementAgentId *string `mandatory:"false" json:"managementAgentId"`

	ConnectionDetails *ConnectionDetails `mandatory:"false" json:"connectionDetails"`

	ConnectionCredentialDetails CredentialDetails `mandatory:"false" json:"connectionCredentialDetails"`

	// Display name of database
	DatabaseDisplayName *string `mandatory:"false" json:"databaseDisplayName"`

	// Additional details of a database in JSON format. For autonomous databases, this is the AutonomousDatabase object serialized as a JSON string as defined in https://docs.oracle.com/iaas/en-us/iaas/api/#/en/database/20160918/AutonomousDatabase/. For EM, pass in null or an empty string. Note that this string needs to be escaped when specified in the curl command.
	DbAdditionalDetails *interface{} `mandatory:"false" json:"dbAdditionalDetails"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM Cluster or DB System ID, depending on which configuration the resource belongs to.
	ParentId *string `mandatory:"false" json:"parentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Infrastructure.
	RootId *string `mandatory:"false" json:"rootId"`

	// Indicates the status of a database insight in Operations Insights
	Status ResourceStatusEnum `mandatory:"true" json:"status"`

	// The current state of the database.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m MacsManagedCloudDatabaseInsight) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m MacsManagedCloudDatabaseInsight) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetStatus returns Status
func (m MacsManagedCloudDatabaseInsight) GetStatus() ResourceStatusEnum {
	return m.Status
}

// GetDatabaseType returns DatabaseType
func (m MacsManagedCloudDatabaseInsight) GetDatabaseType() *string {
	return m.DatabaseType
}

// GetDatabaseVersion returns DatabaseVersion
func (m MacsManagedCloudDatabaseInsight) GetDatabaseVersion() *string {
	return m.DatabaseVersion
}

// GetProcessorCount returns ProcessorCount
func (m MacsManagedCloudDatabaseInsight) GetProcessorCount() *int {
	return m.ProcessorCount
}

// GetFreeformTags returns FreeformTags
func (m MacsManagedCloudDatabaseInsight) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m MacsManagedCloudDatabaseInsight) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m MacsManagedCloudDatabaseInsight) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetTimeCreated returns TimeCreated
func (m MacsManagedCloudDatabaseInsight) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m MacsManagedCloudDatabaseInsight) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m MacsManagedCloudDatabaseInsight) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m MacsManagedCloudDatabaseInsight) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetDatabaseConnectionStatusDetails returns DatabaseConnectionStatusDetails
func (m MacsManagedCloudDatabaseInsight) GetDatabaseConnectionStatusDetails() *string {
	return m.DatabaseConnectionStatusDetails
}

func (m MacsManagedCloudDatabaseInsight) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MacsManagedCloudDatabaseInsight) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingResourceStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetResourceStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MacsManagedCloudDatabaseInsight) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMacsManagedCloudDatabaseInsight MacsManagedCloudDatabaseInsight
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeMacsManagedCloudDatabaseInsight
	}{
		"MACS_MANAGED_CLOUD_DATABASE",
		(MarshalTypeMacsManagedCloudDatabaseInsight)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *MacsManagedCloudDatabaseInsight) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DatabaseType                    *string                           `json:"databaseType"`
		DatabaseVersion                 *string                           `json:"databaseVersion"`
		ProcessorCount                  *int                              `json:"processorCount"`
		SystemTags                      map[string]map[string]interface{} `json:"systemTags"`
		TimeUpdated                     *common.SDKTime                   `json:"timeUpdated"`
		LifecycleDetails                *string                           `json:"lifecycleDetails"`
		DatabaseConnectionStatusDetails *string                           `json:"databaseConnectionStatusDetails"`
		ManagementAgentId               *string                           `json:"managementAgentId"`
		ConnectionDetails               *ConnectionDetails                `json:"connectionDetails"`
		ConnectionCredentialDetails     credentialdetails                 `json:"connectionCredentialDetails"`
		DatabaseDisplayName             *string                           `json:"databaseDisplayName"`
		DbAdditionalDetails             *interface{}                      `json:"dbAdditionalDetails"`
		ParentId                        *string                           `json:"parentId"`
		RootId                          *string                           `json:"rootId"`
		Id                              *string                           `json:"id"`
		CompartmentId                   *string                           `json:"compartmentId"`
		Status                          ResourceStatusEnum                `json:"status"`
		FreeformTags                    map[string]string                 `json:"freeformTags"`
		DefinedTags                     map[string]map[string]interface{} `json:"definedTags"`
		TimeCreated                     *common.SDKTime                   `json:"timeCreated"`
		LifecycleState                  LifecycleStateEnum                `json:"lifecycleState"`
		DatabaseId                      *string                           `json:"databaseId"`
		DatabaseName                    *string                           `json:"databaseName"`
		DatabaseResourceType            *string                           `json:"databaseResourceType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DatabaseType = model.DatabaseType

	m.DatabaseVersion = model.DatabaseVersion

	m.ProcessorCount = model.ProcessorCount

	m.SystemTags = model.SystemTags

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

	m.DatabaseConnectionStatusDetails = model.DatabaseConnectionStatusDetails

	m.ManagementAgentId = model.ManagementAgentId

	m.ConnectionDetails = model.ConnectionDetails

	nn, e = model.ConnectionCredentialDetails.UnmarshalPolymorphicJSON(model.ConnectionCredentialDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ConnectionCredentialDetails = nn.(CredentialDetails)
	} else {
		m.ConnectionCredentialDetails = nil
	}

	m.DatabaseDisplayName = model.DatabaseDisplayName

	m.DbAdditionalDetails = model.DbAdditionalDetails

	m.ParentId = model.ParentId

	m.RootId = model.RootId

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.Status = model.Status

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.DatabaseId = model.DatabaseId

	m.DatabaseName = model.DatabaseName

	m.DatabaseResourceType = model.DatabaseResourceType

	return
}
