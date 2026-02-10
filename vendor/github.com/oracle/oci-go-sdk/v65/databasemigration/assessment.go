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

// Assessment Assessment resource
type Assessment interface {

	// The OCID of the resource being referenced.
	GetId() *string

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	GetDisplayName() *string

	// The OCID of the resource being referenced.
	GetCompartmentId() *string

	// A network speed in Megabits per second.
	GetNetworkSpeedMegabitPerSecond() NetworkSpeedMegabitPerSecondEnum

	// Time allowed for the application downtime.
	GetAcceptableDowntime() AcceptableDowntimeEnum

	// The size of a source database.
	GetDatabaseDataSize() DatabaseDataSizeEnum

	// DDL expectation values.
	GetDdlExpectation() DdlExpectationEnum

	// The type of assessment creation.
	GetCreationType() CreationTypeEnum

	GetSourceDatabaseConnection() *SourceAssessmentConnection

	GetTargetDatabaseConnection() *TargetAssessmentConnection

	// The current state of the Assessment resource.
	GetLifecycleState() AssessmentLifecycleStatesEnum

	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	GetTimeCreated() *common.SDKTime

	// A user-friendly description. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	GetDescription() *string

	// The OCID of the resource being referenced.
	GetMigrationId() *string

	// The migration type of the migration to be performed.
	GetAssessmentMigrationType() AssessmentMigrationTypesEnum

	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	GetTimeUpdated() *common.SDKTime

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags. Example: {"Department": "Finance"}
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type assessment struct {
	JsonData                     []byte
	Description                  *string                           `mandatory:"false" json:"description"`
	MigrationId                  *string                           `mandatory:"false" json:"migrationId"`
	AssessmentMigrationType      AssessmentMigrationTypesEnum      `mandatory:"false" json:"assessmentMigrationType,omitempty"`
	TimeUpdated                  *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	FreeformTags                 map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags                  map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags                   map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id                           *string                           `mandatory:"true" json:"id"`
	DisplayName                  *string                           `mandatory:"true" json:"displayName"`
	CompartmentId                *string                           `mandatory:"true" json:"compartmentId"`
	NetworkSpeedMegabitPerSecond NetworkSpeedMegabitPerSecondEnum  `mandatory:"true" json:"networkSpeedMegabitPerSecond"`
	AcceptableDowntime           AcceptableDowntimeEnum            `mandatory:"true" json:"acceptableDowntime"`
	DatabaseDataSize             DatabaseDataSizeEnum              `mandatory:"true" json:"databaseDataSize"`
	DdlExpectation               DdlExpectationEnum                `mandatory:"true" json:"ddlExpectation"`
	CreationType                 CreationTypeEnum                  `mandatory:"true" json:"creationType"`
	SourceDatabaseConnection     *SourceAssessmentConnection       `mandatory:"true" json:"sourceDatabaseConnection"`
	TargetDatabaseConnection     *TargetAssessmentConnection       `mandatory:"true" json:"targetDatabaseConnection"`
	LifecycleState               AssessmentLifecycleStatesEnum     `mandatory:"true" json:"lifecycleState"`
	TimeCreated                  *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	DatabaseCombination          string                            `json:"databaseCombination"`
}

// UnmarshalJSON unmarshals json
func (m *assessment) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerassessment assessment
	s := struct {
		Model Unmarshalerassessment
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.NetworkSpeedMegabitPerSecond = s.Model.NetworkSpeedMegabitPerSecond
	m.AcceptableDowntime = s.Model.AcceptableDowntime
	m.DatabaseDataSize = s.Model.DatabaseDataSize
	m.DdlExpectation = s.Model.DdlExpectation
	m.CreationType = s.Model.CreationType
	m.SourceDatabaseConnection = s.Model.SourceDatabaseConnection
	m.TargetDatabaseConnection = s.Model.TargetDatabaseConnection
	m.LifecycleState = s.Model.LifecycleState
	m.TimeCreated = s.Model.TimeCreated
	m.Description = s.Model.Description
	m.MigrationId = s.Model.MigrationId
	m.AssessmentMigrationType = s.Model.AssessmentMigrationType
	m.TimeUpdated = s.Model.TimeUpdated
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.DatabaseCombination = s.Model.DatabaseCombination

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *assessment) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DatabaseCombination {
	case "MYSQL":
		mm := MySqlAssessment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE":
		mm := OracleAssessment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for Assessment: %s.", m.DatabaseCombination)
		return *m, nil
	}
}

// GetDescription returns Description
func (m assessment) GetDescription() *string {
	return m.Description
}

// GetMigrationId returns MigrationId
func (m assessment) GetMigrationId() *string {
	return m.MigrationId
}

// GetAssessmentMigrationType returns AssessmentMigrationType
func (m assessment) GetAssessmentMigrationType() AssessmentMigrationTypesEnum {
	return m.AssessmentMigrationType
}

// GetTimeUpdated returns TimeUpdated
func (m assessment) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetFreeformTags returns FreeformTags
func (m assessment) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m assessment) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m assessment) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m assessment) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m assessment) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m assessment) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetNetworkSpeedMegabitPerSecond returns NetworkSpeedMegabitPerSecond
func (m assessment) GetNetworkSpeedMegabitPerSecond() NetworkSpeedMegabitPerSecondEnum {
	return m.NetworkSpeedMegabitPerSecond
}

// GetAcceptableDowntime returns AcceptableDowntime
func (m assessment) GetAcceptableDowntime() AcceptableDowntimeEnum {
	return m.AcceptableDowntime
}

// GetDatabaseDataSize returns DatabaseDataSize
func (m assessment) GetDatabaseDataSize() DatabaseDataSizeEnum {
	return m.DatabaseDataSize
}

// GetDdlExpectation returns DdlExpectation
func (m assessment) GetDdlExpectation() DdlExpectationEnum {
	return m.DdlExpectation
}

// GetCreationType returns CreationType
func (m assessment) GetCreationType() CreationTypeEnum {
	return m.CreationType
}

// GetSourceDatabaseConnection returns SourceDatabaseConnection
func (m assessment) GetSourceDatabaseConnection() *SourceAssessmentConnection {
	return m.SourceDatabaseConnection
}

// GetTargetDatabaseConnection returns TargetDatabaseConnection
func (m assessment) GetTargetDatabaseConnection() *TargetAssessmentConnection {
	return m.TargetDatabaseConnection
}

// GetLifecycleState returns LifecycleState
func (m assessment) GetLifecycleState() AssessmentLifecycleStatesEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m assessment) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

func (m assessment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m assessment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNetworkSpeedMegabitPerSecondEnum(string(m.NetworkSpeedMegabitPerSecond)); !ok && m.NetworkSpeedMegabitPerSecond != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetworkSpeedMegabitPerSecond: %s. Supported values are: %s.", m.NetworkSpeedMegabitPerSecond, strings.Join(GetNetworkSpeedMegabitPerSecondEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAcceptableDowntimeEnum(string(m.AcceptableDowntime)); !ok && m.AcceptableDowntime != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AcceptableDowntime: %s. Supported values are: %s.", m.AcceptableDowntime, strings.Join(GetAcceptableDowntimeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseDataSizeEnum(string(m.DatabaseDataSize)); !ok && m.DatabaseDataSize != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseDataSize: %s. Supported values are: %s.", m.DatabaseDataSize, strings.Join(GetDatabaseDataSizeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDdlExpectationEnum(string(m.DdlExpectation)); !ok && m.DdlExpectation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DdlExpectation: %s. Supported values are: %s.", m.DdlExpectation, strings.Join(GetDdlExpectationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreationTypeEnum(string(m.CreationType)); !ok && m.CreationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CreationType: %s. Supported values are: %s.", m.CreationType, strings.Join(GetCreationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAssessmentLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAssessmentLifecycleStatesEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAssessmentMigrationTypesEnum(string(m.AssessmentMigrationType)); !ok && m.AssessmentMigrationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssessmentMigrationType: %s. Supported values are: %s.", m.AssessmentMigrationType, strings.Join(GetAssessmentMigrationTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
