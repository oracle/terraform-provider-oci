// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
// pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ModelSummary Summary of the language Model.
type ModelSummary struct {

	// Unique identifier model OCID of a model that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)  for the model's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ModelDetails ModelDetails `mandatory:"true" json:"modelDetails"`

	// The time the the Model was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The state of the model.
	LifecycleState ModelLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the model.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// A short description of the Model.
	Description *string `mandatory:"false" json:"description"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// For pre trained models this will identify model type version used for model creation
	// For custom this will identify model type version used for model creation and custom model on which training has to be done
	// <<service>>::<<service-name>>_<<model-type-version>>::<<custom model on which this training has to be done>>
	// ex: ai-lang::NER_V1::CUSTOM-V0
	Version *string `mandatory:"false" json:"version"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{ "orcl-cloud": { "free-tier-retained": "true" } }`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ModelSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModelSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingModelLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetModelLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ModelSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description      *string                           `json:"description"`
		LifecycleDetails *string                           `json:"lifecycleDetails"`
		Version          *string                           `json:"version"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		SystemTags       map[string]map[string]interface{} `json:"systemTags"`
		Id               *string                           `json:"id"`
		DisplayName      *string                           `json:"displayName"`
		CompartmentId    *string                           `json:"compartmentId"`
		ModelDetails     modeldetails                      `json:"modelDetails"`
		TimeCreated      *common.SDKTime                   `json:"timeCreated"`
		LifecycleState   ModelLifecycleStateEnum           `json:"lifecycleState"`
		ProjectId        *string                           `json:"projectId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.LifecycleDetails = model.LifecycleDetails

	m.Version = model.Version

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	nn, e = model.ModelDetails.UnmarshalPolymorphicJSON(model.ModelDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ModelDetails = nn.(ModelDetails)
	} else {
		m.ModelDetails = nil
	}

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.ProjectId = model.ProjectId

	return
}
