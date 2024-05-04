// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// MySqlMigration MySQL Migration resource
type MySqlMigration struct {

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

	// A user-friendly description. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the resource being referenced.
	ExecutingJobId *string `mandatory:"false" json:"executingJobId"`

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

	DataTransferMediumDetails MySqlDataTransferMediumDetails `mandatory:"false" json:"dataTransferMediumDetails"`

	InitialLoadSettings *MySqlInitialLoadSettings `mandatory:"false" json:"initialLoadSettings"`

	AdvisorSettings *MySqlAdvisorSettings `mandatory:"false" json:"advisorSettings"`

	HubDetails *GoldenGateHubDetails `mandatory:"false" json:"hubDetails"`

	GgsDetails *MySqlGgsDeploymentDetails `mandatory:"false" json:"ggsDetails"`

	// The type of the migration to be performed.
	// Example: ONLINE if no downtime is preferred for a migration. This method uses Oracle GoldenGate for replication.
	Type MigrationTypesEnum `mandatory:"true" json:"type"`

	// You can optionally pause a migration after a job phase.
	// This property allows you to optionally specify the phase after which you can pause the migration.
	WaitAfter OdmsJobPhasesEnum `mandatory:"false" json:"waitAfter,omitempty"`

	// The current state of the Migration resource.
	LifecycleState MigrationLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// Additional status related to the execution and current state of the Migration.
	LifecycleDetails MigrationStatusEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`
}

// GetId returns Id
func (m MySqlMigration) GetId() *string {
	return m.Id
}

// GetDescription returns Description
func (m MySqlMigration) GetDescription() *string {
	return m.Description
}

// GetDisplayName returns DisplayName
func (m MySqlMigration) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m MySqlMigration) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetType returns Type
func (m MySqlMigration) GetType() MigrationTypesEnum {
	return m.Type
}

// GetWaitAfter returns WaitAfter
func (m MySqlMigration) GetWaitAfter() OdmsJobPhasesEnum {
	return m.WaitAfter
}

// GetSourceDatabaseConnectionId returns SourceDatabaseConnectionId
func (m MySqlMigration) GetSourceDatabaseConnectionId() *string {
	return m.SourceDatabaseConnectionId
}

// GetTargetDatabaseConnectionId returns TargetDatabaseConnectionId
func (m MySqlMigration) GetTargetDatabaseConnectionId() *string {
	return m.TargetDatabaseConnectionId
}

// GetExecutingJobId returns ExecutingJobId
func (m MySqlMigration) GetExecutingJobId() *string {
	return m.ExecutingJobId
}

// GetTimeCreated returns TimeCreated
func (m MySqlMigration) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m MySqlMigration) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetTimeLastMigration returns TimeLastMigration
func (m MySqlMigration) GetTimeLastMigration() *common.SDKTime {
	return m.TimeLastMigration
}

// GetLifecycleState returns LifecycleState
func (m MySqlMigration) GetLifecycleState() MigrationLifecycleStatesEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m MySqlMigration) GetLifecycleDetails() MigrationStatusEnum {
	return m.LifecycleDetails
}

// GetFreeformTags returns FreeformTags
func (m MySqlMigration) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m MySqlMigration) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m MySqlMigration) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m MySqlMigration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MySqlMigration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMigrationTypesEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMigrationTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOdmsJobPhasesEnum(string(m.WaitAfter)); !ok && m.WaitAfter != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WaitAfter: %s. Supported values are: %s.", m.WaitAfter, strings.Join(GetOdmsJobPhasesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMigrationLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMigrationLifecycleStatesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMigrationStatusEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetMigrationStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MySqlMigration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMySqlMigration MySqlMigration
	s := struct {
		DiscriminatorParam string `json:"databaseCombination"`
		MarshalTypeMySqlMigration
	}{
		"MYSQL",
		(MarshalTypeMySqlMigration)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *MySqlMigration) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                *string                           `json:"description"`
		WaitAfter                  OdmsJobPhasesEnum                 `json:"waitAfter"`
		ExecutingJobId             *string                           `json:"executingJobId"`
		TimeUpdated                *common.SDKTime                   `json:"timeUpdated"`
		TimeLastMigration          *common.SDKTime                   `json:"timeLastMigration"`
		LifecycleDetails           MigrationStatusEnum               `json:"lifecycleDetails"`
		FreeformTags               map[string]string                 `json:"freeformTags"`
		DefinedTags                map[string]map[string]interface{} `json:"definedTags"`
		SystemTags                 map[string]map[string]interface{} `json:"systemTags"`
		DataTransferMediumDetails  mysqldatatransfermediumdetails    `json:"dataTransferMediumDetails"`
		InitialLoadSettings        *MySqlInitialLoadSettings         `json:"initialLoadSettings"`
		AdvisorSettings            *MySqlAdvisorSettings             `json:"advisorSettings"`
		HubDetails                 *GoldenGateHubDetails             `json:"hubDetails"`
		GgsDetails                 *MySqlGgsDeploymentDetails        `json:"ggsDetails"`
		Id                         *string                           `json:"id"`
		DisplayName                *string                           `json:"displayName"`
		CompartmentId              *string                           `json:"compartmentId"`
		Type                       MigrationTypesEnum                `json:"type"`
		SourceDatabaseConnectionId *string                           `json:"sourceDatabaseConnectionId"`
		TargetDatabaseConnectionId *string                           `json:"targetDatabaseConnectionId"`
		TimeCreated                *common.SDKTime                   `json:"timeCreated"`
		LifecycleState             MigrationLifecycleStatesEnum      `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.WaitAfter = model.WaitAfter

	m.ExecutingJobId = model.ExecutingJobId

	m.TimeUpdated = model.TimeUpdated

	m.TimeLastMigration = model.TimeLastMigration

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	nn, e = model.DataTransferMediumDetails.UnmarshalPolymorphicJSON(model.DataTransferMediumDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DataTransferMediumDetails = nn.(MySqlDataTransferMediumDetails)
	} else {
		m.DataTransferMediumDetails = nil
	}

	m.InitialLoadSettings = model.InitialLoadSettings

	m.AdvisorSettings = model.AdvisorSettings

	m.HubDetails = model.HubDetails

	m.GgsDetails = model.GgsDetails

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
