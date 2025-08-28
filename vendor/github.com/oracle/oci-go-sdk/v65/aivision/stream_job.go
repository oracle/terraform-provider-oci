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

// StreamJob Job details for a stream analysis.
type StreamJob struct {

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the streamJob.
	Id *string `mandatory:"true" json:"id"`

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the streamSource
	StreamSourceId *string `mandatory:"true" json:"streamSourceId"`

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// a list of document analysis features.
	Features []VideoStreamFeature `mandatory:"true" json:"features"`

	// The current state of the Stream job.
	LifecycleState StreamJobLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// When the streamJob was created, as an RFC3339 datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Stream job display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	StreamOutputLocation StreamOutputLocation `mandatory:"false" json:"streamOutputLocation"`

	// participant id of agent where results need to be sent
	AgentParticipantId *string `mandatory:"false" json:"agentParticipantId"`

	// Additional details about current state of streamJob
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// When the stream job was updated, as an RFC3339 datetime string.
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

func (m StreamJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StreamJob) ValidateEnumValue() (bool, error) {
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
func (m *StreamJob) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName          *string                           `json:"displayName"`
		StreamOutputLocation streamoutputlocation              `json:"streamOutputLocation"`
		AgentParticipantId   *string                           `json:"agentParticipantId"`
		LifecycleDetails     *string                           `json:"lifecycleDetails"`
		TimeUpdated          *common.SDKTime                   `json:"timeUpdated"`
		FreeformTags         map[string]string                 `json:"freeformTags"`
		DefinedTags          map[string]map[string]interface{} `json:"definedTags"`
		SystemTags           map[string]map[string]interface{} `json:"systemTags"`
		Id                   *string                           `json:"id"`
		StreamSourceId       *string                           `json:"streamSourceId"`
		CompartmentId        *string                           `json:"compartmentId"`
		Features             []videostreamfeature              `json:"features"`
		LifecycleState       StreamJobLifecycleStateEnum       `json:"lifecycleState"`
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

	m.AgentParticipantId = model.AgentParticipantId

	m.LifecycleDetails = model.LifecycleDetails

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

	m.TimeCreated = model.TimeCreated

	return
}

// StreamJobLifecycleStateEnum Enum with underlying type: string
type StreamJobLifecycleStateEnum string

// Set of constants representing the allowable values for StreamJobLifecycleStateEnum
const (
	StreamJobLifecycleStateCreating       StreamJobLifecycleStateEnum = "CREATING"
	StreamJobLifecycleStateUpdating       StreamJobLifecycleStateEnum = "UPDATING"
	StreamJobLifecycleStateActive         StreamJobLifecycleStateEnum = "ACTIVE"
	StreamJobLifecycleStateDeleting       StreamJobLifecycleStateEnum = "DELETING"
	StreamJobLifecycleStateDeleted        StreamJobLifecycleStateEnum = "DELETED"
	StreamJobLifecycleStateFailed         StreamJobLifecycleStateEnum = "FAILED"
	StreamJobLifecycleStateInactive       StreamJobLifecycleStateEnum = "INACTIVE"
	StreamJobLifecycleStateNeedsAttention StreamJobLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingStreamJobLifecycleStateEnum = map[string]StreamJobLifecycleStateEnum{
	"CREATING":        StreamJobLifecycleStateCreating,
	"UPDATING":        StreamJobLifecycleStateUpdating,
	"ACTIVE":          StreamJobLifecycleStateActive,
	"DELETING":        StreamJobLifecycleStateDeleting,
	"DELETED":         StreamJobLifecycleStateDeleted,
	"FAILED":          StreamJobLifecycleStateFailed,
	"INACTIVE":        StreamJobLifecycleStateInactive,
	"NEEDS_ATTENTION": StreamJobLifecycleStateNeedsAttention,
}

var mappingStreamJobLifecycleStateEnumLowerCase = map[string]StreamJobLifecycleStateEnum{
	"creating":        StreamJobLifecycleStateCreating,
	"updating":        StreamJobLifecycleStateUpdating,
	"active":          StreamJobLifecycleStateActive,
	"deleting":        StreamJobLifecycleStateDeleting,
	"deleted":         StreamJobLifecycleStateDeleted,
	"failed":          StreamJobLifecycleStateFailed,
	"inactive":        StreamJobLifecycleStateInactive,
	"needs_attention": StreamJobLifecycleStateNeedsAttention,
}

// GetStreamJobLifecycleStateEnumValues Enumerates the set of values for StreamJobLifecycleStateEnum
func GetStreamJobLifecycleStateEnumValues() []StreamJobLifecycleStateEnum {
	values := make([]StreamJobLifecycleStateEnum, 0)
	for _, v := range mappingStreamJobLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetStreamJobLifecycleStateEnumStringValues Enumerates the set of values in String for StreamJobLifecycleStateEnum
func GetStreamJobLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"INACTIVE",
		"NEEDS_ATTENTION",
	}
}

// GetMappingStreamJobLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStreamJobLifecycleStateEnum(val string) (StreamJobLifecycleStateEnum, bool) {
	enum, ok := mappingStreamJobLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
