// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OracleMigrationSummary Oracle Migration Summary.
// Deprecated: The parent-level Oracle migration configuration properties on this model are deprecated.
// Use `migrationSettings` instead.
type OracleMigrationSummary struct {

	// The OCID of the resource being referenced.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the resource being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the resource being referenced.
	SourceDatabaseConnectionId *string `mandatory:"true" json:"sourceDatabaseConnectionId"`

	// The OCID of the resource being referenced.
	TargetDatabaseConnectionId *string `mandatory:"true" json:"targetDatabaseConnectionId"`

	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID of the resource being referenced.
	ExecutingJobId *string `mandatory:"false" json:"executingJobId"`

	// The OCID of the resource being referenced.
	AssessmentId *string `mandatory:"false" json:"assessmentId"`

	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	TimeLastMigration *common.SDKTime `mandatory:"false" json:"timeLastMigration"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags. Example: {"Department": "Finance"}
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	MigrationSettings OracleMigrationSettingsSummary `mandatory:"false" json:"migrationSettings"`

	// The OCID of the resource being referenced.
	SourceContainerDatabaseConnectionId *string `mandatory:"false" json:"sourceContainerDatabaseConnectionId"`

	// The OCID of the resource being referenced.
	SourceStandbyDatabaseConnectionId *string `mandatory:"false" json:"sourceStandbyDatabaseConnectionId"`

	// The type of the migration to be performed.
	// Example: ONLINE if no downtime is preferred for a migration. This method uses Oracle GoldenGate for replication.
	Type MigrationTypesEnum `mandatory:"true" json:"type"`

	// The current state of the Migration resource.
	LifecycleState MigrationLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// Additional status related to the execution and current state of the Migration.
	LifecycleDetails MigrationStatusEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`
}

// GetId returns Id
func (m OracleMigrationSummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m OracleMigrationSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m OracleMigrationSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetType returns Type
func (m OracleMigrationSummary) GetType() MigrationTypesEnum {
	return m.Type
}

// GetSourceDatabaseConnectionId returns SourceDatabaseConnectionId
func (m OracleMigrationSummary) GetSourceDatabaseConnectionId() *string {
	return m.SourceDatabaseConnectionId
}

// GetTargetDatabaseConnectionId returns TargetDatabaseConnectionId
func (m OracleMigrationSummary) GetTargetDatabaseConnectionId() *string {
	return m.TargetDatabaseConnectionId
}

// GetExecutingJobId returns ExecutingJobId
func (m OracleMigrationSummary) GetExecutingJobId() *string {
	return m.ExecutingJobId
}

// GetAssessmentId returns AssessmentId
func (m OracleMigrationSummary) GetAssessmentId() *string {
	return m.AssessmentId
}

// GetTimeCreated returns TimeCreated
func (m OracleMigrationSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m OracleMigrationSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetTimeLastMigration returns TimeLastMigration
func (m OracleMigrationSummary) GetTimeLastMigration() *common.SDKTime {
	return m.TimeLastMigration
}

// GetLifecycleState returns LifecycleState
func (m OracleMigrationSummary) GetLifecycleState() MigrationLifecycleStatesEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m OracleMigrationSummary) GetLifecycleDetails() MigrationStatusEnum {
	return m.LifecycleDetails
}

// GetFreeformTags returns FreeformTags
func (m OracleMigrationSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m OracleMigrationSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m OracleMigrationSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m OracleMigrationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OracleMigrationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMigrationTypesEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMigrationTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMigrationLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMigrationLifecycleStatesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMigrationStatusEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetMigrationStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OracleMigrationSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOracleMigrationSummary OracleMigrationSummary
	s := struct {
		DiscriminatorParam string `json:"databaseCombination"`
		MarshalTypeOracleMigrationSummary
	}{
		"ORACLE",
		(MarshalTypeOracleMigrationSummary)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *OracleMigrationSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ExecutingJobId                      *string                           `json:"executingJobId"`
		AssessmentId                        *string                           `json:"assessmentId"`
		TimeUpdated                         *common.SDKTime                   `json:"timeUpdated"`
		TimeLastMigration                   *common.SDKTime                   `json:"timeLastMigration"`
		LifecycleDetails                    MigrationStatusEnum               `json:"lifecycleDetails"`
		FreeformTags                        map[string]string                 `json:"freeformTags"`
		DefinedTags                         map[string]map[string]interface{} `json:"definedTags"`
		SystemTags                          map[string]map[string]interface{} `json:"systemTags"`
		MigrationSettings                   oraclemigrationsettingssummary    `json:"migrationSettings"`
		SourceContainerDatabaseConnectionId *string                           `json:"sourceContainerDatabaseConnectionId"`
		SourceStandbyDatabaseConnectionId   *string                           `json:"sourceStandbyDatabaseConnectionId"`
		Id                                  *string                           `json:"id"`
		DisplayName                         *string                           `json:"displayName"`
		CompartmentId                       *string                           `json:"compartmentId"`
		Type                                MigrationTypesEnum                `json:"type"`
		SourceDatabaseConnectionId          *string                           `json:"sourceDatabaseConnectionId"`
		TargetDatabaseConnectionId          *string                           `json:"targetDatabaseConnectionId"`
		TimeCreated                         *common.SDKTime                   `json:"timeCreated"`
		LifecycleState                      MigrationLifecycleStatesEnum      `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ExecutingJobId = model.ExecutingJobId

	m.AssessmentId = model.AssessmentId

	m.TimeUpdated = model.TimeUpdated

	m.TimeLastMigration = model.TimeLastMigration

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	nn, e = model.MigrationSettings.UnmarshalPolymorphicJSON(model.MigrationSettings.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.MigrationSettings = nn.(OracleMigrationSettingsSummary)
	} else {
		m.MigrationSettings = nil
	}

	m.SourceContainerDatabaseConnectionId = model.SourceContainerDatabaseConnectionId

	m.SourceStandbyDatabaseConnectionId = model.SourceStandbyDatabaseConnectionId

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.Type = model.Type

	m.SourceDatabaseConnectionId = model.SourceDatabaseConnectionId

	m.TargetDatabaseConnectionId = model.TargetDatabaseConnectionId

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	return
}
