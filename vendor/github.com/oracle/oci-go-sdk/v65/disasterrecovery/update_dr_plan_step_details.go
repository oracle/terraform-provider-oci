// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateDrPlanStepDetails The details for updating a DR plan step.
type UpdateDrPlanStepDetails struct {

	// The unique id of the step.
	// Example: `sgid1.step..uniqueID`
	Id *string `mandatory:"false" json:"id"`

	// The display name of the step in a group.
	// Example: `My_STEP_3A - EBS Start - STAGE A`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The error mode for this step.
	ErrorMode DrPlanStepErrorModeEnum `mandatory:"false" json:"errorMode,omitempty"`

	// The timeout in seconds for executing this step.
	// Example: `600`
	Timeout *int `mandatory:"false" json:"timeout"`

	// A flag indicating whether this step should be enabled for execution.
	// Example: `true`
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	UserDefinedStep UpdateDrPlanUserDefinedStepDetails `mandatory:"false" json:"userDefinedStep"`
}

func (m UpdateDrPlanStepDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDrPlanStepDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDrPlanStepErrorModeEnum(string(m.ErrorMode)); !ok && m.ErrorMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ErrorMode: %s. Supported values are: %s.", m.ErrorMode, strings.Join(GetDrPlanStepErrorModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateDrPlanStepDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Id              *string                            `json:"id"`
		DisplayName     *string                            `json:"displayName"`
		ErrorMode       DrPlanStepErrorModeEnum            `json:"errorMode"`
		Timeout         *int                               `json:"timeout"`
		IsEnabled       *bool                              `json:"isEnabled"`
		UserDefinedStep updatedrplanuserdefinedstepdetails `json:"userDefinedStep"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.ErrorMode = model.ErrorMode

	m.Timeout = model.Timeout

	m.IsEnabled = model.IsEnabled

	nn, e = model.UserDefinedStep.UnmarshalPolymorphicJSON(model.UserDefinedStep.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.UserDefinedStep = nn.(UpdateDrPlanUserDefinedStepDetails)
	} else {
		m.UserDefinedStep = nil
	}

	return
}
