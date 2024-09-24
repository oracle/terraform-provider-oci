// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseInsightSummary Summary of a database insight resource.
type DatabaseInsightSummary interface {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database insight resource.
	GetId() *string

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
	GetDatabaseId() *string

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// The database name. The database name is unique within the tenancy.
	GetDatabaseName() *string

	// The user-friendly name for the database. The name does not have to be unique.
	GetDatabaseDisplayName() *string

	// Ops Insights internal representation of the database type.
	GetDatabaseType() *string

	// The version of the database.
	GetDatabaseVersion() *string

	// The hostnames for the database.
	GetDatabaseHostNames() []string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}

	// Processor count. This is the OCPU count for Autonomous Database and CPU core count for other database types.
	GetProcessorCount() *int

	// Indicates the status of a database insight in Operations Insights
	GetStatus() ResourceStatusEnum

	// The time the the database insight was first enabled. An RFC3339 formatted datetime string
	GetTimeCreated() *common.SDKTime

	// The time the database insight was updated. An RFC3339 formatted datetime string
	GetTimeUpdated() *common.SDKTime

	// The current state of the database.
	GetLifecycleState() LifecycleStateEnum

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string

	// A message describing the status of the database connection of this resource. For example, it can be used to provide actionable information about the permission and content validity of the database connection.
	GetDatabaseConnectionStatusDetails() *string
}

type databaseinsightsummary struct {
	JsonData                        []byte
	CompartmentId                   *string                           `mandatory:"false" json:"compartmentId"`
	DatabaseName                    *string                           `mandatory:"false" json:"databaseName"`
	DatabaseDisplayName             *string                           `mandatory:"false" json:"databaseDisplayName"`
	DatabaseType                    *string                           `mandatory:"false" json:"databaseType"`
	DatabaseVersion                 *string                           `mandatory:"false" json:"databaseVersion"`
	DatabaseHostNames               []string                          `mandatory:"false" json:"databaseHostNames"`
	FreeformTags                    map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags                     map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags                      map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	ProcessorCount                  *int                              `mandatory:"false" json:"processorCount"`
	Status                          ResourceStatusEnum                `mandatory:"false" json:"status,omitempty"`
	TimeCreated                     *common.SDKTime                   `mandatory:"false" json:"timeCreated"`
	TimeUpdated                     *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	LifecycleState                  LifecycleStateEnum                `mandatory:"false" json:"lifecycleState,omitempty"`
	LifecycleDetails                *string                           `mandatory:"false" json:"lifecycleDetails"`
	DatabaseConnectionStatusDetails *string                           `mandatory:"false" json:"databaseConnectionStatusDetails"`
	Id                              *string                           `mandatory:"true" json:"id"`
	DatabaseId                      *string                           `mandatory:"true" json:"databaseId"`
	EntitySource                    string                            `json:"entitySource"`
}

// UnmarshalJSON unmarshals json
func (m *databaseinsightsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabaseinsightsummary databaseinsightsummary
	s := struct {
		Model Unmarshalerdatabaseinsightsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DatabaseId = s.Model.DatabaseId
	m.CompartmentId = s.Model.CompartmentId
	m.DatabaseName = s.Model.DatabaseName
	m.DatabaseDisplayName = s.Model.DatabaseDisplayName
	m.DatabaseType = s.Model.DatabaseType
	m.DatabaseVersion = s.Model.DatabaseVersion
	m.DatabaseHostNames = s.Model.DatabaseHostNames
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.ProcessorCount = s.Model.ProcessorCount
	m.Status = s.Model.Status
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleState = s.Model.LifecycleState
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.DatabaseConnectionStatusDetails = s.Model.DatabaseConnectionStatusDetails
	m.EntitySource = s.Model.EntitySource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databaseinsightsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntitySource {
	case "MACS_MANAGED_EXTERNAL_DATABASE":
		mm := MacsManagedExternalDatabaseInsightSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AUTONOMOUS_DATABASE":
		mm := AutonomousDatabaseInsightSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MACS_MANAGED_CLOUD_DATABASE":
		mm := MacsManagedCloudDatabaseInsightSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MDS_MYSQL_DATABASE_SYSTEM":
		mm := MdsMySqlDatabaseInsightSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PE_COMANAGED_DATABASE":
		mm := PeComanagedDatabaseInsightSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EM_MANAGED_EXTERNAL_DATABASE":
		mm := EmManagedExternalDatabaseInsightSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DatabaseInsightSummary: %s.", m.EntitySource)
		return *m, nil
	}
}

// GetCompartmentId returns CompartmentId
func (m databaseinsightsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDatabaseName returns DatabaseName
func (m databaseinsightsummary) GetDatabaseName() *string {
	return m.DatabaseName
}

// GetDatabaseDisplayName returns DatabaseDisplayName
func (m databaseinsightsummary) GetDatabaseDisplayName() *string {
	return m.DatabaseDisplayName
}

// GetDatabaseType returns DatabaseType
func (m databaseinsightsummary) GetDatabaseType() *string {
	return m.DatabaseType
}

// GetDatabaseVersion returns DatabaseVersion
func (m databaseinsightsummary) GetDatabaseVersion() *string {
	return m.DatabaseVersion
}

// GetDatabaseHostNames returns DatabaseHostNames
func (m databaseinsightsummary) GetDatabaseHostNames() []string {
	return m.DatabaseHostNames
}

// GetFreeformTags returns FreeformTags
func (m databaseinsightsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m databaseinsightsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m databaseinsightsummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetProcessorCount returns ProcessorCount
func (m databaseinsightsummary) GetProcessorCount() *int {
	return m.ProcessorCount
}

// GetStatus returns Status
func (m databaseinsightsummary) GetStatus() ResourceStatusEnum {
	return m.Status
}

// GetTimeCreated returns TimeCreated
func (m databaseinsightsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m databaseinsightsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m databaseinsightsummary) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m databaseinsightsummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetDatabaseConnectionStatusDetails returns DatabaseConnectionStatusDetails
func (m databaseinsightsummary) GetDatabaseConnectionStatusDetails() *string {
	return m.DatabaseConnectionStatusDetails
}

// GetId returns Id
func (m databaseinsightsummary) GetId() *string {
	return m.Id
}

// GetDatabaseId returns DatabaseId
func (m databaseinsightsummary) GetDatabaseId() *string {
	return m.DatabaseId
}

func (m databaseinsightsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databaseinsightsummary) ValidateEnumValue() (bool, error) {
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
