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

// StreamJobSummary Job details for a stream analysis.
type StreamJobSummary struct {

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the streamJob.
	Id *string `mandatory:"true" json:"id"`

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the streamSource
	StreamSourceId *string `mandatory:"true" json:"streamSourceId"`

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// List of document analysis features.
	Features []VideoStreamFeature `mandatory:"true" json:"features"`

	// The current state of the Stream job.
	LifecycleState StreamJobLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Additional Details of the state of streamJob
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// When the streamJob was created, as an RFC3339 datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Stream job display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	StreamOutputLocation StreamOutputLocation `mandatory:"false" json:"streamOutputLocation"`

	// When the streamJob was updated, as an RFC3339 datetime string.
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

func (m StreamJobSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StreamJobSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStreamJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetStreamJobLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *StreamJobSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName          *string                           `json:"displayName"`
		StreamOutputLocation streamoutputlocation              `json:"streamOutputLocation"`
		TimeUpdated          *common.SDKTime                   `json:"timeUpdated"`
		FreeformTags         map[string]string                 `json:"freeformTags"`
		DefinedTags          map[string]map[string]interface{} `json:"definedTags"`
		SystemTags           map[string]map[string]interface{} `json:"systemTags"`
		Id                   *string                           `json:"id"`
		StreamSourceId       *string                           `json:"streamSourceId"`
		CompartmentId        *string                           `json:"compartmentId"`
		Features             []videostreamfeature              `json:"features"`
		LifecycleState       StreamJobLifecycleStateEnum       `json:"lifecycleState"`
		LifecycleDetails     *string                           `json:"lifecycleDetails"`
		TimeCreated          *common.SDKTime                   `json:"timeCreated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	nn, e = model.StreamOutputLocation.UnmarshalPolymorphicJSON(model.StreamOutputLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.StreamOutputLocation = nn.(StreamOutputLocation)
	} else {
		m.StreamOutputLocation = nil
	}

	m.TimeUpdated = model.TimeUpdated

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.StreamSourceId = model.StreamSourceId

	m.CompartmentId = model.CompartmentId

	m.Features = make([]VideoStreamFeature, len(model.Features))
	for i, n := range model.Features {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Features[i] = nn.(VideoStreamFeature)
		} else {
			m.Features[i] = nil
		}
	}
	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.TimeCreated = model.TimeCreated

	return
}
