// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vision API
//
// Using Vision, you can upload images to detect and classify objects in them. If you have lots of images, you can process them in batch using asynchronous API endpoints. Vision's features are thematically split between Document AI for document-centric images, and Image Analysis for object and scene-based images. Pretrained models and custom models are supported.
//

package aivision

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StreamSourceSummary Video stream analysis results.
type StreamSourceSummary struct {
	StreamSourceDetails StreamSourceDetails `mandatory:"true" json:"streamSourceDetails"`

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the streamSource.
	Id *string `mandatory:"true" json:"id"`

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// When the streamSource was created, as an RFC3339 datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the streamSource.
	LifecycleState StreamSourceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// StreamSource  display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// When the streamSource was created, as an RFC3339 datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only.
	// For example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// For example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m StreamSourceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StreamSourceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStreamSourceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetStreamSourceLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *StreamSourceSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName         *string                           `json:"displayName"`
		TimeUpdated         *common.SDKTime                   `json:"timeUpdated"`
		FreeformTags        map[string]string                 `json:"freeformTags"`
		DefinedTags         map[string]map[string]interface{} `json:"definedTags"`
		SystemTags          map[string]map[string]interface{} `json:"systemTags"`
		StreamSourceDetails streamsourcedetails               `json:"streamSourceDetails"`
		Id                  *string                           `json:"id"`
		CompartmentId       *string                           `json:"compartmentId"`
		TimeCreated         *common.SDKTime                   `json:"timeCreated"`
		LifecycleState      StreamSourceLifecycleStateEnum    `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.TimeUpdated = model.TimeUpdated

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	nn, e = model.StreamSourceDetails.UnmarshalPolymorphicJSON(model.StreamSourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.StreamSourceDetails = nn.(StreamSourceDetails)
	} else {
		m.StreamSourceDetails = nil
	}

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	return
}
