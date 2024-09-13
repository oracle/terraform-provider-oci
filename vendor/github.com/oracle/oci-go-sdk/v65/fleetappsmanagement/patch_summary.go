// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchSummary Summary of the Patch.
type PatchSummary struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Should be unique within the tenancy, and cannot be changed after creation.
	// Avoid entering confidential information.
	Name *string `mandatory:"true" json:"name"`

	PatchType *PatchType `mandatory:"true" json:"patchType"`

	// Patch Severity.
	Severity PatchSeverityEnum `mandatory:"true" json:"severity"`

	// Date when the patch was released.
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased"`

	Product *PatchProduct `mandatory:"true" json:"product"`

	// OCID of the compartment to which the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the Patch.
	LifecycleState PatchLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Provide information on who defined the patch.
	// Example: For Custom Patches the value will be USER_DEFINED
	// For Oracle Defined Patches the value will be ORACLE_DEFINED
	Type PatchTypeEnum `mandatory:"false" json:"type,omitempty"`

	ArtifactDetails ArtifactDetails `mandatory:"false" json:"artifactDetails"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Associated region
	ResourceRegion *string `mandatory:"false" json:"resourceRegion"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m PatchSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetPatchSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPatchLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingPatchTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetPatchTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *PatchSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description      *string                           `json:"description"`
		Type             PatchTypeEnum                     `json:"type"`
		ArtifactDetails  artifactdetails                   `json:"artifactDetails"`
		LifecycleDetails *string                           `json:"lifecycleDetails"`
		ResourceRegion   *string                           `json:"resourceRegion"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		SystemTags       map[string]map[string]interface{} `json:"systemTags"`
		Id               *string                           `json:"id"`
		Name             *string                           `json:"name"`
		PatchType        *PatchType                        `json:"patchType"`
		Severity         PatchSeverityEnum                 `json:"severity"`
		TimeReleased     *common.SDKTime                   `json:"timeReleased"`
		Product          *PatchProduct                     `json:"product"`
		CompartmentId    *string                           `json:"compartmentId"`
		LifecycleState   PatchLifecycleStateEnum           `json:"lifecycleState"`
		TimeCreated      *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated      *common.SDKTime                   `json:"timeUpdated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.Type = model.Type

	nn, e = model.ArtifactDetails.UnmarshalPolymorphicJSON(model.ArtifactDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ArtifactDetails = nn.(ArtifactDetails)
	} else {
		m.ArtifactDetails = nil
	}

	m.LifecycleDetails = model.LifecycleDetails

	m.ResourceRegion = model.ResourceRegion

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.Name = model.Name

	m.PatchType = model.PatchType

	m.Severity = model.Severity

	m.TimeReleased = model.TimeReleased

	m.Product = model.Product

	m.CompartmentId = model.CompartmentId

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	return
}
