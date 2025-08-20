// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FsuImage Exadata Fleet Update Image resource.
type FsuImage struct {

	// OCID identifier for the Exadata Fleet Update Image.
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Unique name of the Exadata Fleet Update Image.
	ImageName *string `mandatory:"true" json:"imageName"`

	ImageDetails ImageDetails `mandatory:"true" json:"imageDetails"`

	// The time the Exadata Fleet Update Image was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the Exadata Fleet Update Image.
	LifecycleState ImageLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// Exadata Fleet Update Image resource display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time the Exadata Fleet Update Image was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m FsuImage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FsuImage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingImageLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetImageLifecycleStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *FsuImage) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName      *string                           `json:"displayName"`
		TimeUpdated      *common.SDKTime                   `json:"timeUpdated"`
		LifecycleDetails *string                           `json:"lifecycleDetails"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		SystemTags       map[string]map[string]interface{} `json:"systemTags"`
		Id               *string                           `json:"id"`
		CompartmentId    *string                           `json:"compartmentId"`
		ImageName        *string                           `json:"imageName"`
		ImageDetails     imagedetails                      `json:"imageDetails"`
		TimeCreated      *common.SDKTime                   `json:"timeCreated"`
		LifecycleState   ImageLifecycleStatesEnum          `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.ImageName = model.ImageName

	nn, e = model.ImageDetails.UnmarshalPolymorphicJSON(model.ImageDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ImageDetails = nn.(ImageDetails)
	} else {
		m.ImageDetails = nil
	}

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	return
}
