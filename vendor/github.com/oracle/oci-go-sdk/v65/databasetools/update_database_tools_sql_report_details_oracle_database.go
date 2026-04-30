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

// UpdateDatabaseToolsSqlReportDetailsOracleDatabase Database Tools SQL report information to be updated for a report of type ORACLE_DATABASE.
type UpdateDatabaseToolsSqlReportDetailsOracleDatabase struct {

	// A user-friendly name. Does not have to be unique and can be updated. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// SQL query executed to generate the report.
	Source *string `mandatory:"false" json:"source"`

	// A description of the SQL report.
	Description *string `mandatory:"false" json:"description"`

	// Purpose of the Database Tools SQL report. Scenario or conditions describing when or why this report should be used. Provides selection criteria to AI agents to improve report selection accuracy.
	Purpose *string `mandatory:"false" json:"purpose"`

	// Instructions on how to use the SQL report. Step-by-step guidance for an AI agent on how to execute or fill in parameters for the report.
	Instructions *string `mandatory:"false" json:"instructions"`

	// Variables referenced in the Database Tools SQL Report source.
	Variables []DatabaseToolsSqlReportVariable `mandatory:"false" json:"variables"`

	// Descriptive information on columns referenced in the Database Tools SQL Report source.
	Columns []DatabaseToolsSqlReportColumn `mandatory:"false" json:"columns"`
}

// GetDisplayName returns DisplayName
func (m UpdateDatabaseToolsSqlReportDetailsOracleDatabase) GetDisplayName() *string {
	return m.DisplayName
}

// GetDefinedTags returns DefinedTags
func (m UpdateDatabaseToolsSqlReportDetailsOracleDatabase) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m UpdateDatabaseToolsSqlReportDetailsOracleDatabase) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetSource returns Source
func (m UpdateDatabaseToolsSqlReportDetailsOracleDatabase) GetSource() *string {
	return m.Source
}

// GetDescription returns Description
func (m UpdateDatabaseToolsSqlReportDetailsOracleDatabase) GetDescription() *string {
	return m.Description
}

// GetPurpose returns Purpose
func (m UpdateDatabaseToolsSqlReportDetailsOracleDatabase) GetPurpose() *string {
	return m.Purpose
}

// GetInstructions returns Instructions
func (m UpdateDatabaseToolsSqlReportDetailsOracleDatabase) GetInstructions() *string {
	return m.Instructions
}

// GetVariables returns Variables
func (m UpdateDatabaseToolsSqlReportDetailsOracleDatabase) GetVariables() []DatabaseToolsSqlReportVariable {
	return m.Variables
}

// GetColumns returns Columns
func (m UpdateDatabaseToolsSqlReportDetailsOracleDatabase) GetColumns() []DatabaseToolsSqlReportColumn {
	return m.Columns
}

func (m UpdateDatabaseToolsSqlReportDetailsOracleDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDatabaseToolsSqlReportDetailsOracleDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateDatabaseToolsSqlReportDetailsOracleDatabase) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateDatabaseToolsSqlReportDetailsOracleDatabase UpdateDatabaseToolsSqlReportDetailsOracleDatabase
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateDatabaseToolsSqlReportDetailsOracleDatabase
	}{
		"ORACLE_DATABASE",
		(MarshalTypeUpdateDatabaseToolsSqlReportDetailsOracleDatabase)(m),
	}

	return json.Marshal(&s)
}
