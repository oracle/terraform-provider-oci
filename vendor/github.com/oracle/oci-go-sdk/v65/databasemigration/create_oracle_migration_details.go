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

// CreateOracleMigrationDetails Create Migration resource parameters.
type CreateOracleMigrationDetails struct {

	// The OCID of the resource being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the resource being referenced.
	SourceDatabaseConnectionId *string `mandatory:"true" json:"sourceDatabaseConnectionId"`

	// The OCID of the resource being referenced.
	TargetDatabaseConnectionId *string `mandatory:"true" json:"targetDatabaseConnectionId"`

	// A user-friendly description. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags. Example: {"Department": "Finance"}
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	DataTransferMediumDetails CreateOracleDataTransferMediumDetails `mandatory:"false" json:"dataTransferMediumDetails"`

	InitialLoadSettings *CreateOracleInitialLoadSettings `mandatory:"false" json:"initialLoadSettings"`

	AdvisorSettings *CreateOracleAdvisorSettings `mandatory:"false" json:"advisorSettings"`

	HubDetails *CreateGoldenGateHubDetails `mandatory:"false" json:"hubDetails"`

	GgsDetails *CreateOracleGgsDeploymentDetails `mandatory:"false" json:"ggsDetails"`

	// List of Migration Parameter objects.
	AdvancedParameters []MigrationParameterDetails `mandatory:"false" json:"advancedParameters"`

	// The OCID of the resource being referenced.
	SourceContainerDatabaseConnectionId *string `mandatory:"false" json:"sourceContainerDatabaseConnectionId"`

	// Database objects to exclude from migration, cannot be specified alongside 'includeObjects'
	ExcludeObjects []OracleDatabaseObject `mandatory:"false" json:"excludeObjects"`

	// Database objects to include from migration, cannot be specified alongside 'excludeObjects'
	IncludeObjects []OracleDatabaseObject `mandatory:"false" json:"includeObjects"`

	// Specifies the database objects to be excluded from the migration in bulk.
	// The definition accepts input in a CSV format, newline separated for each entry.
	// More details can be found in the documentation.
	BulkIncludeExcludeData *string `mandatory:"false" json:"bulkIncludeExcludeData"`

	// The type of the migration to be performed.
	// Example: ONLINE if no downtime is preferred for a migration. This method uses Oracle GoldenGate for replication.
	Type MigrationTypesEnum `mandatory:"true" json:"type"`
}

// GetDescription returns Description
func (m CreateOracleMigrationDetails) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m CreateOracleMigrationDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetType returns Type
func (m CreateOracleMigrationDetails) GetType() MigrationTypesEnum {
	return m.Type
}

// GetDisplayName returns DisplayName
func (m CreateOracleMigrationDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetSourceDatabaseConnectionId returns SourceDatabaseConnectionId
func (m CreateOracleMigrationDetails) GetSourceDatabaseConnectionId() *string {
	return m.SourceDatabaseConnectionId
}

// GetTargetDatabaseConnectionId returns TargetDatabaseConnectionId
func (m CreateOracleMigrationDetails) GetTargetDatabaseConnectionId() *string {
	return m.TargetDatabaseConnectionId
}

// GetFreeformTags returns FreeformTags
func (m CreateOracleMigrationDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateOracleMigrationDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateOracleMigrationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOracleMigrationDetails) ValidateEnumValue() (bool, error) {
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
func (m CreateOracleMigrationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOracleMigrationDetails CreateOracleMigrationDetails
	s := struct {
		DiscriminatorParam string `json:"databaseCombination"`
		MarshalTypeCreateOracleMigrationDetails
	}{
		"ORACLE",
		(MarshalTypeCreateOracleMigrationDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateOracleMigrationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                         *string                               `json:"description"`
		DisplayName                         *string                               `json:"displayName"`
		FreeformTags                        map[string]string                     `json:"freeformTags"`
		DefinedTags                         map[string]map[string]interface{}     `json:"definedTags"`
		DataTransferMediumDetails           createoracledatatransfermediumdetails `json:"dataTransferMediumDetails"`
		InitialLoadSettings                 *CreateOracleInitialLoadSettings      `json:"initialLoadSettings"`
		AdvisorSettings                     *CreateOracleAdvisorSettings          `json:"advisorSettings"`
		HubDetails                          *CreateGoldenGateHubDetails           `json:"hubDetails"`
		GgsDetails                          *CreateOracleGgsDeploymentDetails     `json:"ggsDetails"`
		AdvancedParameters                  []MigrationParameterDetails           `json:"advancedParameters"`
		SourceContainerDatabaseConnectionId *string                               `json:"sourceContainerDatabaseConnectionId"`
		ExcludeObjects                      []OracleDatabaseObject                `json:"excludeObjects"`
		IncludeObjects                      []OracleDatabaseObject                `json:"includeObjects"`
		BulkIncludeExcludeData              *string                               `json:"bulkIncludeExcludeData"`
		CompartmentId                       *string                               `json:"compartmentId"`
		Type                                MigrationTypesEnum                    `json:"type"`
		SourceDatabaseConnectionId          *string                               `json:"sourceDatabaseConnectionId"`
		TargetDatabaseConnectionId          *string                               `json:"targetDatabaseConnectionId"`
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
		m.DataTransferMediumDetails = nn.(CreateOracleDataTransferMediumDetails)
	} else {
		m.DataTransferMediumDetails = nil
	}

	m.InitialLoadSettings = model.InitialLoadSettings

	m.AdvisorSettings = model.AdvisorSettings

	m.HubDetails = model.HubDetails

	m.GgsDetails = model.GgsDetails

	m.AdvancedParameters = make([]MigrationParameterDetails, len(model.AdvancedParameters))
	copy(m.AdvancedParameters, model.AdvancedParameters)
	m.SourceContainerDatabaseConnectionId = model.SourceContainerDatabaseConnectionId

	m.ExcludeObjects = make([]OracleDatabaseObject, len(model.ExcludeObjects))
	copy(m.ExcludeObjects, model.ExcludeObjects)
	m.IncludeObjects = make([]OracleDatabaseObject, len(model.IncludeObjects))
	copy(m.IncludeObjects, model.IncludeObjects)
	m.BulkIncludeExcludeData = model.BulkIncludeExcludeData

	m.CompartmentId = model.CompartmentId

	m.Type = model.Type

	m.SourceDatabaseConnectionId = model.SourceDatabaseConnectionId

	m.TargetDatabaseConnectionId = model.TargetDatabaseConnectionId

	return
}
