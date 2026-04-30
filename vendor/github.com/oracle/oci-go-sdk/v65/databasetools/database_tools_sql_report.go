// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools API
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseToolsSqlReport Allows generating a summary of SQL query execution details for database operations.
type DatabaseToolsSqlReport interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools SQL report.
	GetId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools SQL report.
	GetCompartmentId() *string

	// A user-friendly name. Does not have to be unique and can be updated. Avoid entering confidential information.
	GetDisplayName() *string

	// SQL query executed to generate the report.
	GetSource() *string

	// The current state of the Database Tools SQL report.
	GetLifecycleState() DatabaseToolsSqlReportLifecycleStateEnum

	// The time the Database Tools SQL report was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The time the Database Tools SQL report was updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	// A description of the SQL report.
	GetDescription() *string

	// Purpose of the Database Tools SQL report. Scenario or conditions describing when or why this report should be used. Provides selection criteria to AI agents to improve report selection accuracy.
	GetPurpose() *string

	// Instructions on how to use the SQL report. Step-by-step guidance for an AI agent on how to execute or fill in parameters for the report.
	GetInstructions() *string

	// Variables referenced in the Database Tools SQL Report source.
	GetVariables() []DatabaseToolsSqlReportVariable

	// Descriptive information on columns referenced in the Database Tools SQL Report source.
	GetColumns() []DatabaseToolsSqlReportColumn

	// A message describing the current state in more detail.
	GetLifecycleDetails() *string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}

	// Locks associated with this resource.
	GetLocks() []ResourceLock
}

type databasetoolssqlreport struct {
	JsonData         []byte
	Description      *string                                  `mandatory:"false" json:"description"`
	Purpose          *string                                  `mandatory:"false" json:"purpose"`
	Instructions     *string                                  `mandatory:"false" json:"instructions"`
	Variables        []DatabaseToolsSqlReportVariable         `mandatory:"false" json:"variables"`
	Columns          []DatabaseToolsSqlReportColumn           `mandatory:"false" json:"columns"`
	LifecycleDetails *string                                  `mandatory:"false" json:"lifecycleDetails"`
	DefinedTags      map[string]map[string]interface{}        `mandatory:"false" json:"definedTags"`
	FreeformTags     map[string]string                        `mandatory:"false" json:"freeformTags"`
	SystemTags       map[string]map[string]interface{}        `mandatory:"false" json:"systemTags"`
	Locks            []ResourceLock                           `mandatory:"false" json:"locks"`
	Id               *string                                  `mandatory:"true" json:"id"`
	CompartmentId    *string                                  `mandatory:"true" json:"compartmentId"`
	DisplayName      *string                                  `mandatory:"true" json:"displayName"`
	Source           *string                                  `mandatory:"true" json:"source"`
	LifecycleState   DatabaseToolsSqlReportLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	TimeCreated      *common.SDKTime                          `mandatory:"true" json:"timeCreated"`
	TimeUpdated      *common.SDKTime                          `mandatory:"true" json:"timeUpdated"`
	Type             string                                   `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolssqlreport) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolssqlreport databasetoolssqlreport
	s := struct {
		Model Unmarshalerdatabasetoolssqlreport
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.Source = s.Model.Source
	m.LifecycleState = s.Model.LifecycleState
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.Description = s.Model.Description
	m.Purpose = s.Model.Purpose
	m.Instructions = s.Model.Instructions
	m.Variables = s.Model.Variables
	m.Columns = s.Model.Columns
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.DefinedTags = s.Model.DefinedTags
	m.FreeformTags = s.Model.FreeformTags
	m.SystemTags = s.Model.SystemTags
	m.Locks = s.Model.Locks
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolssqlreport) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "ORACLE_DATABASE":
		mm := DatabaseToolsSqlReportOracleDatabase{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DatabaseToolsSqlReport: %s.", m.Type)
		return *m, nil
	}
}

// GetDescription returns Description
func (m databasetoolssqlreport) GetDescription() *string {
	return m.Description
}

// GetPurpose returns Purpose
func (m databasetoolssqlreport) GetPurpose() *string {
	return m.Purpose
}

// GetInstructions returns Instructions
func (m databasetoolssqlreport) GetInstructions() *string {
	return m.Instructions
}

// GetVariables returns Variables
func (m databasetoolssqlreport) GetVariables() []DatabaseToolsSqlReportVariable {
	return m.Variables
}

// GetColumns returns Columns
func (m databasetoolssqlreport) GetColumns() []DatabaseToolsSqlReportColumn {
	return m.Columns
}

// GetLifecycleDetails returns LifecycleDetails
func (m databasetoolssqlreport) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetDefinedTags returns DefinedTags
func (m databasetoolssqlreport) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m databasetoolssqlreport) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetSystemTags returns SystemTags
func (m databasetoolssqlreport) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLocks returns Locks
func (m databasetoolssqlreport) GetLocks() []ResourceLock {
	return m.Locks
}

// GetId returns Id
func (m databasetoolssqlreport) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m databasetoolssqlreport) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m databasetoolssqlreport) GetDisplayName() *string {
	return m.DisplayName
}

// GetSource returns Source
func (m databasetoolssqlreport) GetSource() *string {
	return m.Source
}

// GetLifecycleState returns LifecycleState
func (m databasetoolssqlreport) GetLifecycleState() DatabaseToolsSqlReportLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m databasetoolssqlreport) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m databasetoolssqlreport) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m databasetoolssqlreport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolssqlreport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseToolsSqlReportLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDatabaseToolsSqlReportLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
