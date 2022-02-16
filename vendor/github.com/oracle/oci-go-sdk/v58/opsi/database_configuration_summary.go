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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// DatabaseConfigurationSummary Summary of a database configuration for a resource.
type DatabaseConfigurationSummary interface {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database insight resource.
	GetDatabaseInsightId() *string

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// The database name. The database name is unique within the tenancy.
	GetDatabaseName() *string

	// The user-friendly name for the database. The name does not have to be unique.
	GetDatabaseDisplayName() *string

	// Operations Insights internal representation of the database type.
	GetDatabaseType() *string

	// The version of the database.
	GetDatabaseVersion() *string

	// Name of the CDB.Only applies to PDB.
	GetCdbName() *string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Processor count. This is the OCPU count for Autonomous Database and CPU core count for other database types.
	GetProcessorCount() *int
}

type databaseconfigurationsummary struct {
	JsonData            []byte
	DatabaseInsightId   *string                           `mandatory:"true" json:"databaseInsightId"`
	CompartmentId       *string                           `mandatory:"true" json:"compartmentId"`
	DatabaseName        *string                           `mandatory:"true" json:"databaseName"`
	DatabaseDisplayName *string                           `mandatory:"true" json:"databaseDisplayName"`
	DatabaseType        *string                           `mandatory:"true" json:"databaseType"`
	DatabaseVersion     *string                           `mandatory:"true" json:"databaseVersion"`
	CdbName             *string                           `mandatory:"true" json:"cdbName"`
	DefinedTags         map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`
	FreeformTags        map[string]string                 `mandatory:"true" json:"freeformTags"`
	ProcessorCount      *int                              `mandatory:"false" json:"processorCount"`
	EntitySource        string                            `json:"entitySource"`
}

// UnmarshalJSON unmarshals json
func (m *databaseconfigurationsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabaseconfigurationsummary databaseconfigurationsummary
	s := struct {
		Model Unmarshalerdatabaseconfigurationsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DatabaseInsightId = s.Model.DatabaseInsightId
	m.CompartmentId = s.Model.CompartmentId
	m.DatabaseName = s.Model.DatabaseName
	m.DatabaseDisplayName = s.Model.DatabaseDisplayName
	m.DatabaseType = s.Model.DatabaseType
	m.DatabaseVersion = s.Model.DatabaseVersion
	m.CdbName = s.Model.CdbName
	m.DefinedTags = s.Model.DefinedTags
	m.FreeformTags = s.Model.FreeformTags
	m.ProcessorCount = s.Model.ProcessorCount
	m.EntitySource = s.Model.EntitySource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databaseconfigurationsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntitySource {
	case "EM_MANAGED_EXTERNAL_DATABASE":
		mm := EmManagedExternalDatabaseConfigurationSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AUTONOMOUS_DATABASE":
		mm := AutonomousDatabaseConfigurationSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MACS_MANAGED_EXTERNAL_DATABASE":
		mm := MacsManagedExternalDatabaseConfigurationSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDatabaseInsightId returns DatabaseInsightId
func (m databaseconfigurationsummary) GetDatabaseInsightId() *string {
	return m.DatabaseInsightId
}

//GetCompartmentId returns CompartmentId
func (m databaseconfigurationsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDatabaseName returns DatabaseName
func (m databaseconfigurationsummary) GetDatabaseName() *string {
	return m.DatabaseName
}

//GetDatabaseDisplayName returns DatabaseDisplayName
func (m databaseconfigurationsummary) GetDatabaseDisplayName() *string {
	return m.DatabaseDisplayName
}

//GetDatabaseType returns DatabaseType
func (m databaseconfigurationsummary) GetDatabaseType() *string {
	return m.DatabaseType
}

//GetDatabaseVersion returns DatabaseVersion
func (m databaseconfigurationsummary) GetDatabaseVersion() *string {
	return m.DatabaseVersion
}

//GetCdbName returns CdbName
func (m databaseconfigurationsummary) GetCdbName() *string {
	return m.CdbName
}

//GetDefinedTags returns DefinedTags
func (m databaseconfigurationsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetFreeformTags returns FreeformTags
func (m databaseconfigurationsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetProcessorCount returns ProcessorCount
func (m databaseconfigurationsummary) GetProcessorCount() *int {
	return m.ProcessorCount
}

func (m databaseconfigurationsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databaseconfigurationsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
