// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MigrationPlanSummary Summary of the migration plan.
type MigrationPlanSummary struct {

	// The unique Oracle ID (OCID) that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the migration plan.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time when the migration plan was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the migration plan.
	LifecycleState MigrationPlanLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the associated migration.
	MigrationId *string `mandatory:"true" json:"migrationId"`

	// List of strategies for the resources to be migrated.
	Strategies []ResourceAssessmentStrategy `mandatory:"true" json:"strategies"`

	// Limits of the resources that are needed for migration. Example: {"BlockVolume": 2, "VCN": 1}
	CalculatedLimits map[string]int `mandatory:"true" json:"calculatedLimits"`

	// List of target environments.
	TargetEnvironments []TargetEnvironment `mandatory:"true" json:"targetEnvironments"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time when the migration plan was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	MigrationPlanStats *MigrationPlanStats `mandatory:"false" json:"migrationPlanStats"`

	// OCID of the referenced ORM job.
	ReferenceToRmsStack *string `mandatory:"false" json:"referenceToRmsStack"`

	// Source migraiton plan ID to be cloned.
	SourceMigrationPlanId *string `mandatory:"false" json:"sourceMigrationPlanId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m MigrationPlanSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MigrationPlanSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMigrationPlanLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMigrationPlanLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *MigrationPlanSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName           *string                           `json:"displayName"`
		TimeUpdated           *common.SDKTime                   `json:"timeUpdated"`
		LifecycleDetails      *string                           `json:"lifecycleDetails"`
		MigrationPlanStats    *MigrationPlanStats               `json:"migrationPlanStats"`
		ReferenceToRmsStack   *string                           `json:"referenceToRmsStack"`
		SourceMigrationPlanId *string                           `json:"sourceMigrationPlanId"`
		FreeformTags          map[string]string                 `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{} `json:"definedTags"`
		SystemTags            map[string]map[string]interface{} `json:"systemTags"`
		Id                    *string                           `json:"id"`
		CompartmentId         *string                           `json:"compartmentId"`
		TimeCreated           *common.SDKTime                   `json:"timeCreated"`
		LifecycleState        MigrationPlanLifecycleStateEnum   `json:"lifecycleState"`
		MigrationId           *string                           `json:"migrationId"`
		Strategies            []resourceassessmentstrategy      `json:"strategies"`
		CalculatedLimits      map[string]int                    `json:"calculatedLimits"`
		TargetEnvironments    []targetenvironment               `json:"targetEnvironments"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

	m.MigrationPlanStats = model.MigrationPlanStats

	m.ReferenceToRmsStack = model.ReferenceToRmsStack

	m.SourceMigrationPlanId = model.SourceMigrationPlanId

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.MigrationId = model.MigrationId

	m.Strategies = make([]ResourceAssessmentStrategy, len(model.Strategies))
	for i, n := range model.Strategies {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Strategies[i] = nn.(ResourceAssessmentStrategy)
		} else {
			m.Strategies[i] = nil
		}
	}
	m.CalculatedLimits = model.CalculatedLimits

	m.TargetEnvironments = make([]TargetEnvironment, len(model.TargetEnvironments))
	for i, n := range model.TargetEnvironments {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.TargetEnvironments[i] = nn.(TargetEnvironment)
		} else {
			m.TargetEnvironments[i] = nil
		}
	}
	return
}
