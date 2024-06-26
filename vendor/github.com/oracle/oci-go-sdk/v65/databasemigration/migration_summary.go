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

// MigrationSummary Migration resource
type MigrationSummary interface {

	// The OCID of the resource being referenced.
	GetId() *string

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	GetDisplayName() *string

	// The OCID of the resource being referenced.
	GetCompartmentId() *string

	// The type of the migration to be performed.
	// Example: ONLINE if no downtime is preferred for a migration. This method uses Oracle GoldenGate for replication.
	GetType() MigrationTypesEnum

	// The OCID of the resource being referenced.
	GetSourceDatabaseConnectionId() *string

	// The OCID of the resource being referenced.
	GetTargetDatabaseConnectionId() *string

	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	GetTimeCreated() *common.SDKTime

	// The current state of the Migration resource.
	GetLifecycleState() MigrationLifecycleStatesEnum

	// The OCID of the resource being referenced.
	GetExecutingJobId() *string

	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	GetTimeUpdated() *common.SDKTime

	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	GetTimeLastMigration() *common.SDKTime

	// Additional status related to the execution and current state of the Migration.
	GetLifecycleDetails() MigrationStatusEnum

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

type migrationsummary struct {
	JsonData                   []byte
	ExecutingJobId             *string                           `mandatory:"false" json:"executingJobId"`
	TimeUpdated                *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	TimeLastMigration          *common.SDKTime                   `mandatory:"false" json:"timeLastMigration"`
	LifecycleDetails           MigrationStatusEnum               `mandatory:"false" json:"lifecycleDetails,omitempty"`
	FreeformTags               map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags                map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags                 map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id                         *string                           `mandatory:"true" json:"id"`
	DisplayName                *string                           `mandatory:"true" json:"displayName"`
	CompartmentId              *string                           `mandatory:"true" json:"compartmentId"`
	Type                       MigrationTypesEnum                `mandatory:"true" json:"type"`
	SourceDatabaseConnectionId *string                           `mandatory:"true" json:"sourceDatabaseConnectionId"`
	TargetDatabaseConnectionId *string                           `mandatory:"true" json:"targetDatabaseConnectionId"`
	TimeCreated                *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	LifecycleState             MigrationLifecycleStatesEnum      `mandatory:"true" json:"lifecycleState"`
	DatabaseCombination        string                            `json:"databaseCombination"`
}

// UnmarshalJSON unmarshals json
func (m *migrationsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermigrationsummary migrationsummary
	s := struct {
		Model Unmarshalermigrationsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.Type = s.Model.Type
	m.SourceDatabaseConnectionId = s.Model.SourceDatabaseConnectionId
	m.TargetDatabaseConnectionId = s.Model.TargetDatabaseConnectionId
	m.TimeCreated = s.Model.TimeCreated
	m.LifecycleState = s.Model.LifecycleState
	m.ExecutingJobId = s.Model.ExecutingJobId
	m.TimeUpdated = s.Model.TimeUpdated
	m.TimeLastMigration = s.Model.TimeLastMigration
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.DatabaseCombination = s.Model.DatabaseCombination

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *migrationsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DatabaseCombination {
	case "ORACLE":
		mm := OracleMigrationSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL":
		mm := MySqlMigrationSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for MigrationSummary: %s.", m.DatabaseCombination)
		return *m, nil
	}
}

// GetExecutingJobId returns ExecutingJobId
func (m migrationsummary) GetExecutingJobId() *string {
	return m.ExecutingJobId
}

// GetTimeUpdated returns TimeUpdated
func (m migrationsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetTimeLastMigration returns TimeLastMigration
func (m migrationsummary) GetTimeLastMigration() *common.SDKTime {
	return m.TimeLastMigration
}

// GetLifecycleDetails returns LifecycleDetails
func (m migrationsummary) GetLifecycleDetails() MigrationStatusEnum {
	return m.LifecycleDetails
}

// GetFreeformTags returns FreeformTags
func (m migrationsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m migrationsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m migrationsummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m migrationsummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m migrationsummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m migrationsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetType returns Type
func (m migrationsummary) GetType() MigrationTypesEnum {
	return m.Type
}

// GetSourceDatabaseConnectionId returns SourceDatabaseConnectionId
func (m migrationsummary) GetSourceDatabaseConnectionId() *string {
	return m.SourceDatabaseConnectionId
}

// GetTargetDatabaseConnectionId returns TargetDatabaseConnectionId
func (m migrationsummary) GetTargetDatabaseConnectionId() *string {
	return m.TargetDatabaseConnectionId
}

// GetTimeCreated returns TimeCreated
func (m migrationsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetLifecycleState returns LifecycleState
func (m migrationsummary) GetLifecycleState() MigrationLifecycleStatesEnum {
	return m.LifecycleState
}

func (m migrationsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m migrationsummary) ValidateEnumValue() (bool, error) {
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
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
