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

// CreateMySqlMigrationDetails Create Migration resource parameters.
type CreateMySqlMigrationDetails struct {

	// OCI resource ID.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OCI resource ID.
	SourceDatabaseConnectionId *string `mandatory:"true" json:"sourceDatabaseConnectionId"`

	// OCI resource ID.
	TargetDatabaseConnectionId *string `mandatory:"true" json:"targetDatabaseConnectionId"`

	// An object's Description.
	Description *string `mandatory:"false" json:"description"`

	// An object's Display Name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	DataTransferMediumDetails CreateMySqlDataTransferMediumDetails `mandatory:"false" json:"dataTransferMediumDetails"`

	InitialLoadSettings *CreateMySqlInitialLoadSettings `mandatory:"false" json:"initialLoadSettings"`

	AdvisorSettings *CreateMySqlAdvisorSettings `mandatory:"false" json:"advisorSettings"`

	// Database objects to exclude from migration, cannot be specified alongside 'includeObjects'
	ExcludeObjects []MySqlDatabaseObject `mandatory:"false" json:"excludeObjects"`

	// Database objects to include from migration, cannot be specified alongside 'excludeObjects'
	IncludeObjects []MySqlDatabaseObject `mandatory:"false" json:"includeObjects"`

	// Database objects to exclude/include from migration in CSV format. The excludeObjects and includeObjects
	// fields will be ignored if this field is not null.
	BulkIncludeExcludeData *string `mandatory:"false" json:"bulkIncludeExcludeData"`

	HubDetails *CreateGoldenGateHubDetails `mandatory:"false" json:"hubDetails"`

	GgsDetails *CreateMySqlGgsDeploymentDetails `mandatory:"false" json:"ggsDetails"`

	// Migration type (ONLINE/OFFLINE).
	Type MigrationTypesEnum `mandatory:"true" json:"type"`
}

// GetDescription returns Description
func (m CreateMySqlMigrationDetails) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m CreateMySqlMigrationDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetType returns Type
func (m CreateMySqlMigrationDetails) GetType() MigrationTypesEnum {
	return m.Type
}

// GetDisplayName returns DisplayName
func (m CreateMySqlMigrationDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetSourceDatabaseConnectionId returns SourceDatabaseConnectionId
func (m CreateMySqlMigrationDetails) GetSourceDatabaseConnectionId() *string {
	return m.SourceDatabaseConnectionId
}

// GetTargetDatabaseConnectionId returns TargetDatabaseConnectionId
func (m CreateMySqlMigrationDetails) GetTargetDatabaseConnectionId() *string {
	return m.TargetDatabaseConnectionId
}

// GetFreeformTags returns FreeformTags
func (m CreateMySqlMigrationDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateMySqlMigrationDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateMySqlMigrationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMySqlMigrationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMigrationTypesEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMigrationTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateMySqlMigrationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateMySqlMigrationDetails CreateMySqlMigrationDetails
	s := struct {
		DiscriminatorParam string `json:"databaseCombination"`
		MarshalTypeCreateMySqlMigrationDetails
	}{
		"MYSQL",
		(MarshalTypeCreateMySqlMigrationDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateMySqlMigrationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                *string                              `json:"description"`
		DisplayName                *string                              `json:"displayName"`
		FreeformTags               map[string]string                    `json:"freeformTags"`
		DefinedTags                map[string]map[string]interface{}    `json:"definedTags"`
		DataTransferMediumDetails  createmysqldatatransfermediumdetails `json:"dataTransferMediumDetails"`
		InitialLoadSettings        *CreateMySqlInitialLoadSettings      `json:"initialLoadSettings"`
		AdvisorSettings            *CreateMySqlAdvisorSettings          `json:"advisorSettings"`
		ExcludeObjects             []MySqlDatabaseObject                `json:"excludeObjects"`
		IncludeObjects             []MySqlDatabaseObject                `json:"includeObjects"`
		BulkIncludeExcludeData     *string                              `json:"bulkIncludeExcludeData"`
		HubDetails                 *CreateGoldenGateHubDetails          `json:"hubDetails"`
		GgsDetails                 *CreateMySqlGgsDeploymentDetails     `json:"ggsDetails"`
		CompartmentId              *string                              `json:"compartmentId"`
		Type                       MigrationTypesEnum                   `json:"type"`
		SourceDatabaseConnectionId *string                              `json:"sourceDatabaseConnectionId"`
		TargetDatabaseConnectionId *string                              `json:"targetDatabaseConnectionId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	nn, e = model.DataTransferMediumDetails.UnmarshalPolymorphicJSON(model.DataTransferMediumDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DataTransferMediumDetails = nn.(CreateMySqlDataTransferMediumDetails)
	} else {
		m.DataTransferMediumDetails = nil
	}

	m.InitialLoadSettings = model.InitialLoadSettings

	m.AdvisorSettings = model.AdvisorSettings

	m.ExcludeObjects = make([]MySqlDatabaseObject, len(model.ExcludeObjects))
	copy(m.ExcludeObjects, model.ExcludeObjects)
	m.IncludeObjects = make([]MySqlDatabaseObject, len(model.IncludeObjects))
	copy(m.IncludeObjects, model.IncludeObjects)
	m.BulkIncludeExcludeData = model.BulkIncludeExcludeData

	m.HubDetails = model.HubDetails

	m.GgsDetails = model.GgsDetails

	m.CompartmentId = model.CompartmentId

	m.Type = model.Type

	m.SourceDatabaseConnectionId = model.SourceDatabaseConnectionId

	m.TargetDatabaseConnectionId = model.TargetDatabaseConnectionId

	return
}
