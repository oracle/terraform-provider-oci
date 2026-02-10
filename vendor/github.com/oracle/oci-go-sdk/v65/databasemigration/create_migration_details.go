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

// CreateMigrationDetails Common Migration details.
type CreateMigrationDetails interface {

	// The OCID of the resource being referenced.
	GetCompartmentId() *string

	// The type of the migration to be performed.
	// Example: ONLINE if no downtime is preferred for a migration. This method uses Oracle GoldenGate for replication.
	GetType() MigrationTypesEnum

	// A user-friendly description. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	GetDescription() *string

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	GetDisplayName() *string

	// The OCID of the resource being referenced.
	GetSourceDatabaseConnectionId() *string

	// The OCID of the resource being referenced.
	GetTargetDatabaseConnectionId() *string

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags. Example: {"Department": "Finance"}
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// The OCID of the resource being referenced.
	GetAssessmentId() *string
}

type createmigrationdetails struct {
	JsonData                   []byte
	Description                *string                           `mandatory:"false" json:"description"`
	DisplayName                *string                           `mandatory:"false" json:"displayName"`
	SourceDatabaseConnectionId *string                           `mandatory:"false" json:"sourceDatabaseConnectionId"`
	TargetDatabaseConnectionId *string                           `mandatory:"false" json:"targetDatabaseConnectionId"`
	FreeformTags               map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags                map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	AssessmentId               *string                           `mandatory:"false" json:"assessmentId"`
	CompartmentId              *string                           `mandatory:"true" json:"compartmentId"`
	Type                       MigrationTypesEnum                `mandatory:"true" json:"type"`
	DatabaseCombination        string                            `json:"databaseCombination"`
}

// UnmarshalJSON unmarshals json
func (m *createmigrationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatemigrationdetails createmigrationdetails
	s := struct {
		Model Unmarshalercreatemigrationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.Type = s.Model.Type
	m.Description = s.Model.Description
	m.DisplayName = s.Model.DisplayName
	m.SourceDatabaseConnectionId = s.Model.SourceDatabaseConnectionId
	m.TargetDatabaseConnectionId = s.Model.TargetDatabaseConnectionId
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.AssessmentId = s.Model.AssessmentId
	m.DatabaseCombination = s.Model.DatabaseCombination

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createmigrationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DatabaseCombination {
	case "MYSQL":
		mm := CreateMySqlMigrationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE":
		mm := CreateOracleMigrationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateMigrationDetails: %s.", m.DatabaseCombination)
		return *m, nil
	}
}

// GetDescription returns Description
func (m createmigrationdetails) GetDescription() *string {
	return m.Description
}

// GetDisplayName returns DisplayName
func (m createmigrationdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetSourceDatabaseConnectionId returns SourceDatabaseConnectionId
func (m createmigrationdetails) GetSourceDatabaseConnectionId() *string {
	return m.SourceDatabaseConnectionId
}

// GetTargetDatabaseConnectionId returns TargetDatabaseConnectionId
func (m createmigrationdetails) GetTargetDatabaseConnectionId() *string {
	return m.TargetDatabaseConnectionId
}

// GetFreeformTags returns FreeformTags
func (m createmigrationdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createmigrationdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetAssessmentId returns AssessmentId
func (m createmigrationdetails) GetAssessmentId() *string {
	return m.AssessmentId
}

// GetCompartmentId returns CompartmentId
func (m createmigrationdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetType returns Type
func (m createmigrationdetails) GetType() MigrationTypesEnum {
	return m.Type
}

func (m createmigrationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createmigrationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMigrationTypesEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMigrationTypesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
