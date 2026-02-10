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

// AssessmentSummary Assessment Summary resource
type AssessmentSummary interface {

	// The OCID of the resource being referenced.
	GetId() *string

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	GetDisplayName() *string

	// The OCID of the resource being referenced.
	GetCompartmentId() *string

	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	GetTimeCreated() *common.SDKTime

	// The current state of the Assessment resource.
	GetLifecycleState() AssessmentLifecycleStatesEnum

	// The OCID of the resource being referenced.
	GetMigrationId() *string

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

type assessmentsummary struct {
	JsonData            []byte
	MigrationId         *string                           `mandatory:"false" json:"migrationId"`
	TimeUpdated         *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	FreeformTags        map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags         map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags          map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id                  *string                           `mandatory:"true" json:"id"`
	DisplayName         *string                           `mandatory:"true" json:"displayName"`
	CompartmentId       *string                           `mandatory:"true" json:"compartmentId"`
	TimeCreated         *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	LifecycleState      AssessmentLifecycleStatesEnum     `mandatory:"true" json:"lifecycleState"`
	DatabaseCombination string                            `json:"databaseCombination"`
}

// UnmarshalJSON unmarshals json
func (m *assessmentsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerassessmentsummary assessmentsummary
	s := struct {
		Model Unmarshalerassessmentsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.TimeCreated = s.Model.TimeCreated
	m.LifecycleState = s.Model.LifecycleState
	m.MigrationId = s.Model.MigrationId
	m.TimeUpdated = s.Model.TimeUpdated
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.DatabaseCombination = s.Model.DatabaseCombination

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *assessmentsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DatabaseCombination {
	case "ORACLE":
		mm := OracleAssessmentSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL":
		mm := MySqlAssessmentSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for AssessmentSummary: %s.", m.DatabaseCombination)
		return *m, nil
	}
}

// GetMigrationId returns MigrationId
func (m assessmentsummary) GetMigrationId() *string {
	return m.MigrationId
}

// GetTimeUpdated returns TimeUpdated
func (m assessmentsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetFreeformTags returns FreeformTags
func (m assessmentsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m assessmentsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m assessmentsummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m assessmentsummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m assessmentsummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m assessmentsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetTimeCreated returns TimeCreated
func (m assessmentsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetLifecycleState returns LifecycleState
func (m assessmentsummary) GetLifecycleState() AssessmentLifecycleStatesEnum {
	return m.LifecycleState
}

func (m assessmentsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m assessmentsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAssessmentLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAssessmentLifecycleStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
