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

// DatabaseInsight Database insight resource.
type DatabaseInsight interface {

	// Database insight identifier
	GetId() *string

	// Compartment identifier of the database
	GetCompartmentId() *string

	// Indicates the status of a database insight in Operations Insights
	GetStatus() ResourceStatusEnum

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// The time the the database insight was first enabled. An RFC3339 formatted datetime string
	GetTimeCreated() *common.SDKTime

	// The current state of the database.
	GetLifecycleState() LifecycleStateEnum

	// Operations Insights internal representation of the database type.
	GetDatabaseType() *string

	// The version of the database.
	GetDatabaseVersion() *string

	// Processor count. This is the OCPU count for Autonomous Database and CPU core count for other database types.
	GetProcessorCount() *int

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}

	// The time the database insight was updated. An RFC3339 formatted datetime string
	GetTimeUpdated() *common.SDKTime

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string
}

type databaseinsight struct {
	JsonData         []byte
	Id               *string                           `mandatory:"true" json:"id"`
	CompartmentId    *string                           `mandatory:"true" json:"compartmentId"`
	Status           ResourceStatusEnum                `mandatory:"true" json:"status"`
	FreeformTags     map[string]string                 `mandatory:"true" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`
	TimeCreated      *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	LifecycleState   LifecycleStateEnum                `mandatory:"true" json:"lifecycleState"`
	DatabaseType     *string                           `mandatory:"false" json:"databaseType"`
	DatabaseVersion  *string                           `mandatory:"false" json:"databaseVersion"`
	ProcessorCount   *int                              `mandatory:"false" json:"processorCount"`
	SystemTags       map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	TimeUpdated      *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	LifecycleDetails *string                           `mandatory:"false" json:"lifecycleDetails"`
	EntitySource     string                            `json:"entitySource"`
}

// UnmarshalJSON unmarshals json
func (m *databaseinsight) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabaseinsight databaseinsight
	s := struct {
		Model Unmarshalerdatabaseinsight
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.Status = s.Model.Status
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.TimeCreated = s.Model.TimeCreated
	m.LifecycleState = s.Model.LifecycleState
	m.DatabaseType = s.Model.DatabaseType
	m.DatabaseVersion = s.Model.DatabaseVersion
	m.ProcessorCount = s.Model.ProcessorCount
	m.SystemTags = s.Model.SystemTags
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.EntitySource = s.Model.EntitySource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databaseinsight) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntitySource {
	case "EM_MANAGED_EXTERNAL_DATABASE":
		mm := EmManagedExternalDatabaseInsight{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MACS_MANAGED_EXTERNAL_DATABASE":
		mm := MacsManagedExternalDatabaseInsight{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AUTONOMOUS_DATABASE":
		mm := AutonomousDatabaseInsight{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m databaseinsight) GetId() *string {
	return m.Id
}

//GetCompartmentId returns CompartmentId
func (m databaseinsight) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetStatus returns Status
func (m databaseinsight) GetStatus() ResourceStatusEnum {
	return m.Status
}

//GetFreeformTags returns FreeformTags
func (m databaseinsight) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m databaseinsight) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetTimeCreated returns TimeCreated
func (m databaseinsight) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetLifecycleState returns LifecycleState
func (m databaseinsight) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

//GetDatabaseType returns DatabaseType
func (m databaseinsight) GetDatabaseType() *string {
	return m.DatabaseType
}

//GetDatabaseVersion returns DatabaseVersion
func (m databaseinsight) GetDatabaseVersion() *string {
	return m.DatabaseVersion
}

//GetProcessorCount returns ProcessorCount
func (m databaseinsight) GetProcessorCount() *int {
	return m.ProcessorCount
}

//GetSystemTags returns SystemTags
func (m databaseinsight) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

//GetTimeUpdated returns TimeUpdated
func (m databaseinsight) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleDetails returns LifecycleDetails
func (m databaseinsight) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m databaseinsight) String() string {
	return common.PointerString(m)
}
