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

// UpdatePatchDetails The patch information to be updated.
type UpdatePatchDetails struct {

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	PatchType *PatchType `mandatory:"false" json:"patchType"`

	// Patch Severity.
	Severity PatchSeverityEnum `mandatory:"false" json:"severity,omitempty"`

	// Date when the patch was released.
	TimeReleased *common.SDKTime `mandatory:"false" json:"timeReleased"`

	ArtifactDetails ArtifactDetails `mandatory:"false" json:"artifactDetails"`

	Product *PatchProduct `mandatory:"false" json:"product"`

	// Dependent Patches.
	DependentPatches []DependentPatchDetails `mandatory:"false" json:"dependentPatches"`

	// OCID of the compartment to which the resource belongs to.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdatePatchDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdatePatchDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPatchSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetPatchSeverityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdatePatchDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description      *string                           `json:"description"`
		PatchType        *PatchType                        `json:"patchType"`
		Severity         PatchSeverityEnum                 `json:"severity"`
		TimeReleased     *common.SDKTime                   `json:"timeReleased"`
		ArtifactDetails  artifactdetails                   `json:"artifactDetails"`
		Product          *PatchProduct                     `json:"product"`
		DependentPatches []DependentPatchDetails           `json:"dependentPatches"`
		CompartmentId    *string                           `json:"compartmentId"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.PatchType = model.PatchType

	m.Severity = model.Severity

	m.TimeReleased = model.TimeReleased

	nn, e = model.ArtifactDetails.UnmarshalPolymorphicJSON(model.ArtifactDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ArtifactDetails = nn.(ArtifactDetails)
	} else {
		m.ArtifactDetails = nil
	}

	m.Product = model.Product

	m.DependentPatches = make([]DependentPatchDetails, len(model.DependentPatches))
	copy(m.DependentPatches, model.DependentPatches)
	m.CompartmentId = model.CompartmentId

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
