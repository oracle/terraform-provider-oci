// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.cloud.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogPipelineSummary Log Pipeline object summary.
type LogPipelineSummary struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly display name. This must be unique within the enclosing resource,
	// and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Inputs for the Log Pipeline
	Inputs []LogPipelineInput `mandatory:"true" json:"inputs"`

	// Routes for the Log Pipeline
	Routes []LogPipelineRoute `mandatory:"true" json:"routes"`

	// The pipeline state.
	LifecycleState LogPipelineLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Whether or not this resource is currently enabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// Description for this resource.
	Description *string `mandatory:"false" json:"description"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Time the resource was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time the resource was last modified.
	TimeLastModified *common.SDKTime `mandatory:"false" json:"timeLastModified"`

	// The OCID of the compartment that the resource belongs to.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
}

func (m LogPipelineSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogPipelineSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLogPipelineLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLogPipelineLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *LogPipelineSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description      *string                           `json:"description"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		SystemTags       map[string]map[string]interface{} `json:"systemTags"`
		TimeCreated      *common.SDKTime                   `json:"timeCreated"`
		TimeLastModified *common.SDKTime                   `json:"timeLastModified"`
		CompartmentId    *string                           `json:"compartmentId"`
		Id               *string                           `json:"id"`
		DisplayName      *string                           `json:"displayName"`
		Inputs           []logpipelineinput                `json:"inputs"`
		Routes           []LogPipelineRoute                `json:"routes"`
		LifecycleState   LogPipelineLifecycleStateEnum     `json:"lifecycleState"`
		IsEnabled        *bool                             `json:"isEnabled"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.SystemTags = model.SystemTags

	m.TimeCreated = model.TimeCreated

	m.TimeLastModified = model.TimeLastModified

	m.CompartmentId = model.CompartmentId

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.Inputs = make([]LogPipelineInput, len(model.Inputs))
	for i, n := range model.Inputs {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Inputs[i] = nn.(LogPipelineInput)
		} else {
			m.Inputs[i] = nil
		}
	}
	m.Routes = make([]LogPipelineRoute, len(model.Routes))
	copy(m.Routes, model.Routes)
	m.LifecycleState = model.LifecycleState

	m.IsEnabled = model.IsEnabled

	return
}
