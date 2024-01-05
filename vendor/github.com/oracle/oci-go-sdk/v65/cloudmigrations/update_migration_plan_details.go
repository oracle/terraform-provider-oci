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

// UpdateMigrationPlanDetails The information to be updated.
type UpdateMigrationPlanDetails struct {

	// Migration plan identifier
	DisplayName *string `mandatory:"false" json:"displayName"`

	// List of strategies for the resources to be migrated.
	Strategies []ResourceAssessmentStrategy `mandatory:"false" json:"strategies"`

	// List of target environments.
	TargetEnvironments []TargetEnvironment `mandatory:"false" json:"targetEnvironments"`

	// Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateMigrationPlanDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateMigrationPlanDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateMigrationPlanDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName        *string                           `json:"displayName"`
		Strategies         []resourceassessmentstrategy      `json:"strategies"`
		TargetEnvironments []targetenvironment               `json:"targetEnvironments"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

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
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
