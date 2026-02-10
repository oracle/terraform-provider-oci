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

// MySqlAssessmentSummary MySql Assessment Summary
type MySqlAssessmentSummary struct {

	// The OCID of the resource being referenced.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the resource being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

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

	// The current state of the Assessment resource.
	LifecycleState AssessmentLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m MySqlAssessmentSummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m MySqlAssessmentSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetMigrationId returns MigrationId
func (m MySqlAssessmentSummary) GetMigrationId() *string {
	return m.MigrationId
}

// GetCompartmentId returns CompartmentId
func (m MySqlAssessmentSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetTimeCreated returns TimeCreated
func (m MySqlAssessmentSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m MySqlAssessmentSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m MySqlAssessmentSummary) GetLifecycleState() AssessmentLifecycleStatesEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m MySqlAssessmentSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m MySqlAssessmentSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m MySqlAssessmentSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m MySqlAssessmentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MySqlAssessmentSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAssessmentLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAssessmentLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MySqlAssessmentSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMySqlAssessmentSummary MySqlAssessmentSummary
	s := struct {
		DiscriminatorParam string `json:"databaseCombination"`
		MarshalTypeMySqlAssessmentSummary
	}{
		"MYSQL",
		(MarshalTypeMySqlAssessmentSummary)(m),
	}

	return json.Marshal(&s)
}
