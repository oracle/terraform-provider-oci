// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// DrPlanStep Details of a step in a DR plan.
type DrPlanStep struct {

	// The unique id of the step. Must not be modified by the user.
	// Example: `sgid1.step..uniqueID`
	Id *string `mandatory:"true" json:"id"`

	// The unique id of the group to which this step belongs. Must not be modified by user.
	// Example: `sgid1.group..uniqueID`
	GroupId *string `mandatory:"true" json:"groupId"`

	// The plan step type.
	Type DrPlanStepTypeEnum `mandatory:"true" json:"type"`

	// The display name of the group.
	// Example: `DATABASE_SWITCHOVER`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The error mode for this step.
	ErrorMode DrPlanStepErrorModeEnum `mandatory:"true" json:"errorMode"`

	// The timeout in seconds for executing this step.
	// Example: `600`
	Timeout *int `mandatory:"true" json:"timeout"`

	// A flag indicating whether this step should be enabled for execution.
	// Example: `true`
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The OCID of the member associated with this step.
	// Example: `ocid1.database.oc1..uniqueID`
	MemberId *string `mandatory:"false" json:"memberId"`

	UserDefinedStep DrPlanUserDefinedStep `mandatory:"false" json:"userDefinedStep"`
}

func (m DrPlanStep) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrPlanStep) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDrPlanStepTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDrPlanStepTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDrPlanStepErrorModeEnum(string(m.ErrorMode)); !ok && m.ErrorMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ErrorMode: %s. Supported values are: %s.", m.ErrorMode, strings.Join(GetDrPlanStepErrorModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DrPlanStep) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		MemberId        *string                 `json:"memberId"`
		UserDefinedStep drplanuserdefinedstep   `json:"userDefinedStep"`
		Id              *string                 `json:"id"`
		GroupId         *string                 `json:"groupId"`
		Type            DrPlanStepTypeEnum      `json:"type"`
		DisplayName     *string                 `json:"displayName"`
		ErrorMode       DrPlanStepErrorModeEnum `json:"errorMode"`
		Timeout         *int                    `json:"timeout"`
		IsEnabled       *bool                   `json:"isEnabled"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.MemberId = model.MemberId

	nn, e = model.UserDefinedStep.UnmarshalPolymorphicJSON(model.UserDefinedStep.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.UserDefinedStep = nn.(DrPlanUserDefinedStep)
	} else {
		m.UserDefinedStep = nil
	}

	m.Id = model.Id

	m.GroupId = model.GroupId

	m.Type = model.Type

	m.DisplayName = model.DisplayName

	m.ErrorMode = model.ErrorMode

	m.Timeout = model.Timeout

	m.IsEnabled = model.IsEnabled

	return
}
