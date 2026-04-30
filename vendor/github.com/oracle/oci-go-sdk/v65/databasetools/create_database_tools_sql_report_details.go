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

// CreateDatabaseToolsSqlReportDetails Details for the new Database Tools SQL report.
type CreateDatabaseToolsSqlReportDetails interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools SQL report.
	GetCompartmentId() *string

	// A user-friendly name. Does not have to be unique and can be updated. Avoid entering confidential information.
	GetDisplayName() *string

	// SQL query executed to generate the report.
	GetSource() *string

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

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Locks associated with this resource.
	GetLocks() []ResourceLock
}

type createdatabasetoolssqlreportdetails struct {
	JsonData      []byte
	Description   *string                           `mandatory:"false" json:"description"`
	Purpose       *string                           `mandatory:"false" json:"purpose"`
	Instructions  *string                           `mandatory:"false" json:"instructions"`
	Variables     []DatabaseToolsSqlReportVariable  `mandatory:"false" json:"variables"`
	Columns       []DatabaseToolsSqlReportColumn    `mandatory:"false" json:"columns"`
	DefinedTags   map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	FreeformTags  map[string]string                 `mandatory:"false" json:"freeformTags"`
	Locks         []ResourceLock                    `mandatory:"false" json:"locks"`
	CompartmentId *string                           `mandatory:"true" json:"compartmentId"`
	DisplayName   *string                           `mandatory:"true" json:"displayName"`
	Source        *string                           `mandatory:"true" json:"source"`
	Type          string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createdatabasetoolssqlreportdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedatabasetoolssqlreportdetails createdatabasetoolssqlreportdetails
	s := struct {
		Model Unmarshalercreatedatabasetoolssqlreportdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.Source = s.Model.Source
	m.Description = s.Model.Description
	m.Purpose = s.Model.Purpose
	m.Instructions = s.Model.Instructions
	m.Variables = s.Model.Variables
	m.Columns = s.Model.Columns
	m.DefinedTags = s.Model.DefinedTags
	m.FreeformTags = s.Model.FreeformTags
	m.Locks = s.Model.Locks
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdatabasetoolssqlreportdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "ORACLE_DATABASE":
		mm := CreateDatabaseToolsSqlReportOracleDatabaseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateDatabaseToolsSqlReportDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetDescription returns Description
func (m createdatabasetoolssqlreportdetails) GetDescription() *string {
	return m.Description
}

// GetPurpose returns Purpose
func (m createdatabasetoolssqlreportdetails) GetPurpose() *string {
	return m.Purpose
}

// GetInstructions returns Instructions
func (m createdatabasetoolssqlreportdetails) GetInstructions() *string {
	return m.Instructions
}

// GetVariables returns Variables
func (m createdatabasetoolssqlreportdetails) GetVariables() []DatabaseToolsSqlReportVariable {
	return m.Variables
}

// GetColumns returns Columns
func (m createdatabasetoolssqlreportdetails) GetColumns() []DatabaseToolsSqlReportColumn {
	return m.Columns
}

// GetDefinedTags returns DefinedTags
func (m createdatabasetoolssqlreportdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m createdatabasetoolssqlreportdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetLocks returns Locks
func (m createdatabasetoolssqlreportdetails) GetLocks() []ResourceLock {
	return m.Locks
}

// GetCompartmentId returns CompartmentId
func (m createdatabasetoolssqlreportdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m createdatabasetoolssqlreportdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetSource returns Source
func (m createdatabasetoolssqlreportdetails) GetSource() *string {
	return m.Source
}

func (m createdatabasetoolssqlreportdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdatabasetoolssqlreportdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
