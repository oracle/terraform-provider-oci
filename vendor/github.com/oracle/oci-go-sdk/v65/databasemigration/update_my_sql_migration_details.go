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

// UpdateMySqlMigrationDetails Update Migration parameters.
type UpdateMySqlMigrationDetails struct {

	// A user-friendly description. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the resource being referenced.
	SourceDatabaseConnectionId *string `mandatory:"false" json:"sourceDatabaseConnectionId"`

	// The OCID of the resource being referenced.
	TargetDatabaseConnectionId *string `mandatory:"false" json:"targetDatabaseConnectionId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags. Example: {"Department": "Finance"}
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	DataTransferMediumDetails UpdateMySqlDataTransferMediumDetails `mandatory:"false" json:"dataTransferMediumDetails"`

	InitialLoadSettings *UpdateMySqlInitialLoadSettings `mandatory:"false" json:"initialLoadSettings"`

	AdvisorSettings *UpdateMySqlAdvisorSettings `mandatory:"false" json:"advisorSettings"`

	HubDetails *UpdateGoldenGateHubDetails `mandatory:"false" json:"hubDetails"`

	GgsDetails *UpdateMySqlGgsDeploymentDetails `mandatory:"false" json:"ggsDetails"`

	// The type of the migration to be performed.
	// Example: ONLINE if no downtime is preferred for a migration. This method uses Oracle GoldenGate for replication.
	Type MigrationTypesEnum `mandatory:"false" json:"type,omitempty"`
}

// GetDescription returns Description
func (m UpdateMySqlMigrationDetails) GetDescription() *string {
	return m.Description
}

// GetType returns Type
func (m UpdateMySqlMigrationDetails) GetType() MigrationTypesEnum {
	return m.Type
}

// GetDisplayName returns DisplayName
func (m UpdateMySqlMigrationDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetSourceDatabaseConnectionId returns SourceDatabaseConnectionId
func (m UpdateMySqlMigrationDetails) GetSourceDatabaseConnectionId() *string {
	return m.SourceDatabaseConnectionId
}

// GetTargetDatabaseConnectionId returns TargetDatabaseConnectionId
func (m UpdateMySqlMigrationDetails) GetTargetDatabaseConnectionId() *string {
	return m.TargetDatabaseConnectionId
}

// GetFreeformTags returns FreeformTags
func (m UpdateMySqlMigrationDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateMySqlMigrationDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateMySqlMigrationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateMySqlMigrationDetails) ValidateEnumValue() (bool, error) {
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
func (m UpdateMySqlMigrationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateMySqlMigrationDetails UpdateMySqlMigrationDetails
	s := struct {
		DiscriminatorParam string `json:"databaseCombination"`
		MarshalTypeUpdateMySqlMigrationDetails
	}{
		"MYSQL",
		(MarshalTypeUpdateMySqlMigrationDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateMySqlMigrationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                *string                              `json:"description"`
		Type                       MigrationTypesEnum                   `json:"type"`
		DisplayName                *string                              `json:"displayName"`
		SourceDatabaseConnectionId *string                              `json:"sourceDatabaseConnectionId"`
		TargetDatabaseConnectionId *string                              `json:"targetDatabaseConnectionId"`
		FreeformTags               map[string]string                    `json:"freeformTags"`
		DefinedTags                map[string]map[string]interface{}    `json:"definedTags"`
		DataTransferMediumDetails  updatemysqldatatransfermediumdetails `json:"dataTransferMediumDetails"`
		InitialLoadSettings        *UpdateMySqlInitialLoadSettings      `json:"initialLoadSettings"`
		AdvisorSettings            *UpdateMySqlAdvisorSettings          `json:"advisorSettings"`
		HubDetails                 *UpdateGoldenGateHubDetails          `json:"hubDetails"`
		GgsDetails                 *UpdateMySqlGgsDeploymentDetails     `json:"ggsDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.Type = model.Type

	m.DisplayName = model.DisplayName

	m.SourceDatabaseConnectionId = model.SourceDatabaseConnectionId

	m.TargetDatabaseConnectionId = model.TargetDatabaseConnectionId

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	nn, e = model.DataTransferMediumDetails.UnmarshalPolymorphicJSON(model.DataTransferMediumDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DataTransferMediumDetails = nn.(UpdateMySqlDataTransferMediumDetails)
	} else {
		m.DataTransferMediumDetails = nil
	}

	m.InitialLoadSettings = model.InitialLoadSettings

	m.AdvisorSettings = model.AdvisorSettings

	m.HubDetails = model.HubDetails

	m.GgsDetails = model.GgsDetails

	return
}
