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

// MySqlAssessment MySql Assessment resource
type MySqlAssessment struct {

	// The OCID of the resource being referenced.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the resource being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	SourceDatabaseConnection *SourceAssessmentConnection `mandatory:"true" json:"sourceDatabaseConnection"`

	TargetDatabaseConnection *TargetAssessmentConnection `mandatory:"true" json:"targetDatabaseConnection"`

	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A user-friendly description. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the resource being referenced.
	MigrationId *string `mandatory:"false" json:"migrationId"`

	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags. Example: {"Department": "Finance"}
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// A network speed in Megabits per second.
	NetworkSpeedMegabitPerSecond NetworkSpeedMegabitPerSecondEnum `mandatory:"true" json:"networkSpeedMegabitPerSecond"`

	// Time allowed for the application downtime.
	AcceptableDowntime AcceptableDowntimeEnum `mandatory:"true" json:"acceptableDowntime"`

	// The size of a source database.
	DatabaseDataSize DatabaseDataSizeEnum `mandatory:"true" json:"databaseDataSize"`

	// DDL expectation values.
	DdlExpectation DdlExpectationEnum `mandatory:"true" json:"ddlExpectation"`

	// The type of assessment creation.
	CreationType CreationTypeEnum `mandatory:"true" json:"creationType"`

	// The current state of the Assessment resource.
	LifecycleState AssessmentLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// The migration type of the migration to be performed.
	AssessmentMigrationType AssessmentMigrationTypesEnum `mandatory:"false" json:"assessmentMigrationType,omitempty"`
}

// GetId returns Id
func (m MySqlAssessment) GetId() *string {
	return m.Id
}

// GetDescription returns Description
func (m MySqlAssessment) GetDescription() *string {
	return m.Description
}

// GetDisplayName returns DisplayName
func (m MySqlAssessment) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m MySqlAssessment) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetNetworkSpeedMegabitPerSecond returns NetworkSpeedMegabitPerSecond
func (m MySqlAssessment) GetNetworkSpeedMegabitPerSecond() NetworkSpeedMegabitPerSecondEnum {
	return m.NetworkSpeedMegabitPerSecond
}

// GetAcceptableDowntime returns AcceptableDowntime
func (m MySqlAssessment) GetAcceptableDowntime() AcceptableDowntimeEnum {
	return m.AcceptableDowntime
}

// GetDatabaseDataSize returns DatabaseDataSize
func (m MySqlAssessment) GetDatabaseDataSize() DatabaseDataSizeEnum {
	return m.DatabaseDataSize
}

// GetDdlExpectation returns DdlExpectation
func (m MySqlAssessment) GetDdlExpectation() DdlExpectationEnum {
	return m.DdlExpectation
}

// GetCreationType returns CreationType
func (m MySqlAssessment) GetCreationType() CreationTypeEnum {
	return m.CreationType
}

// GetMigrationId returns MigrationId
func (m MySqlAssessment) GetMigrationId() *string {
	return m.MigrationId
}

// GetSourceDatabaseConnection returns SourceDatabaseConnection
func (m MySqlAssessment) GetSourceDatabaseConnection() *SourceAssessmentConnection {
	return m.SourceDatabaseConnection
}

// GetTargetDatabaseConnection returns TargetDatabaseConnection
func (m MySqlAssessment) GetTargetDatabaseConnection() *TargetAssessmentConnection {
	return m.TargetDatabaseConnection
}

// GetLifecycleState returns LifecycleState
func (m MySqlAssessment) GetLifecycleState() AssessmentLifecycleStatesEnum {
	return m.LifecycleState
}

// GetAssessmentMigrationType returns AssessmentMigrationType
func (m MySqlAssessment) GetAssessmentMigrationType() AssessmentMigrationTypesEnum {
	return m.AssessmentMigrationType
}

// GetTimeCreated returns TimeCreated
func (m MySqlAssessment) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m MySqlAssessment) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetFreeformTags returns FreeformTags
func (m MySqlAssessment) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m MySqlAssessment) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m MySqlAssessment) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m MySqlAssessment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MySqlAssessment) ValidateEnumValue() (bool, error) {
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

// MarshalJSON marshals to json representation
func (m MySqlAssessment) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMySqlAssessment MySqlAssessment
	s := struct {
		DiscriminatorParam string `json:"databaseCombination"`
		MarshalTypeMySqlAssessment
	}{
		"MYSQL",
		(MarshalTypeMySqlAssessment)(m),
	}

	return json.Marshal(&s)
}
