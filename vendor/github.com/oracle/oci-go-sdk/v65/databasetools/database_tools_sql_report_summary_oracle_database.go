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

// DatabaseToolsSqlReportSummaryOracleDatabase Summary of the Database Tools SQL report for an Oracle Database.
type DatabaseToolsSqlReportSummaryOracleDatabase struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools SQL Report.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools SQL Report.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique and can be updated. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time the Database Tools SQL Report was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the Database Tools SQL Report was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// A description of the SQL report.
	Description *string `mandatory:"false" json:"description"`

	// Purpose of the Database Tools SQL report. Scenario or conditions describing when or why this report should be used. Provides selection criteria to AI agents to improve report selection accuracy.
	Purpose *string `mandatory:"false" json:"purpose"`

	// Instructions on how to use the SQL report. Step-by-step guidance for an AI agent on how to execute or fill in parameters for the report.
	Instructions *string `mandatory:"false" json:"instructions"`

	// A message describing the current state of the Database Tools SQL Report.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// The current state of the Database Tools SQL Report.
	LifecycleState DatabaseToolsSqlReportLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m DatabaseToolsSqlReportSummaryOracleDatabase) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m DatabaseToolsSqlReportSummaryOracleDatabase) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m DatabaseToolsSqlReportSummaryOracleDatabase) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m DatabaseToolsSqlReportSummaryOracleDatabase) GetDescription() *string {
	return m.Description
}

// GetPurpose returns Purpose
func (m DatabaseToolsSqlReportSummaryOracleDatabase) GetPurpose() *string {
	return m.Purpose
}

// GetInstructions returns Instructions
func (m DatabaseToolsSqlReportSummaryOracleDatabase) GetInstructions() *string {
	return m.Instructions
}

// GetLifecycleState returns LifecycleState
func (m DatabaseToolsSqlReportSummaryOracleDatabase) GetLifecycleState() DatabaseToolsSqlReportLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m DatabaseToolsSqlReportSummaryOracleDatabase) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m DatabaseToolsSqlReportSummaryOracleDatabase) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DatabaseToolsSqlReportSummaryOracleDatabase) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetDefinedTags returns DefinedTags
func (m DatabaseToolsSqlReportSummaryOracleDatabase) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m DatabaseToolsSqlReportSummaryOracleDatabase) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetSystemTags returns SystemTags
func (m DatabaseToolsSqlReportSummaryOracleDatabase) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLocks returns Locks
func (m DatabaseToolsSqlReportSummaryOracleDatabase) GetLocks() []ResourceLock {
	return m.Locks
}

func (m DatabaseToolsSqlReportSummaryOracleDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsSqlReportSummaryOracleDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDatabaseToolsSqlReportLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDatabaseToolsSqlReportLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DatabaseToolsSqlReportSummaryOracleDatabase) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseToolsSqlReportSummaryOracleDatabase DatabaseToolsSqlReportSummaryOracleDatabase
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDatabaseToolsSqlReportSummaryOracleDatabase
	}{
		"ORACLE_DATABASE",
		(MarshalTypeDatabaseToolsSqlReportSummaryOracleDatabase)(m),
	}

	return json.Marshal(&s)
}
