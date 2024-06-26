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

// MySqlMigrationSummary MySQL Migration Summary
type MySqlMigrationSummary struct {

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

	// The type of the migration to be performed.
	// Example: ONLINE if no downtime is preferred for a migration. This method uses Oracle GoldenGate for replication.
	Type MigrationTypesEnum `mandatory:"true" json:"type"`

	// The current state of the Migration resource.
	LifecycleState MigrationLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// Additional status related to the execution and current state of the Migration.
	LifecycleDetails MigrationStatusEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`
}

// GetId returns Id
func (m MySqlMigrationSummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m MySqlMigrationSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m MySqlMigrationSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetType returns Type
func (m MySqlMigrationSummary) GetType() MigrationTypesEnum {
	return m.Type
}

// GetSourceDatabaseConnectionId returns SourceDatabaseConnectionId
func (m MySqlMigrationSummary) GetSourceDatabaseConnectionId() *string {
	return m.SourceDatabaseConnectionId
}

// GetTargetDatabaseConnectionId returns TargetDatabaseConnectionId
func (m MySqlMigrationSummary) GetTargetDatabaseConnectionId() *string {
	return m.TargetDatabaseConnectionId
}

// GetExecutingJobId returns ExecutingJobId
func (m MySqlMigrationSummary) GetExecutingJobId() *string {
	return m.ExecutingJobId
}

// GetTimeCreated returns TimeCreated
func (m MySqlMigrationSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m MySqlMigrationSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetTimeLastMigration returns TimeLastMigration
func (m MySqlMigrationSummary) GetTimeLastMigration() *common.SDKTime {
	return m.TimeLastMigration
}

// GetLifecycleState returns LifecycleState
func (m MySqlMigrationSummary) GetLifecycleState() MigrationLifecycleStatesEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m MySqlMigrationSummary) GetLifecycleDetails() MigrationStatusEnum {
	return m.LifecycleDetails
}

// GetFreeformTags returns FreeformTags
func (m MySqlMigrationSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m MySqlMigrationSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m MySqlMigrationSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m MySqlMigrationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MySqlMigrationSummary) ValidateEnumValue() (bool, error) {
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

// MarshalJSON marshals to json representation
func (m MySqlMigrationSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMySqlMigrationSummary MySqlMigrationSummary
	s := struct {
		DiscriminatorParam string `json:"databaseCombination"`
		MarshalTypeMySqlMigrationSummary
	}{
		"MYSQL",
		(MarshalTypeMySqlMigrationSummary)(m),
	}

	return json.Marshal(&s)
}
