// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// StreamRun Job details for a stream analysis.
type StreamRun struct {

	// Id of the job.
	StreamId *string `mandatory:"true" json:"streamId"`

	// id of the run
	RunId *string `mandatory:"true" json:"runId"`

	// a list of document analysis features.
	Features []VideoFeature `mandatory:"true" json:"features"`

	OutputLocation *OutputLocation `mandatory:"true" json:"outputLocation"`

	// The current state of the Stream job.
	LifecycleState StreamRunLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The ocid of the compartment that starts the job.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Stream job display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Job accepted time.
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// Job started time.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Job finished time.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

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

func (m StreamRun) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StreamRun) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStreamRunLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetStreamRunLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *StreamRun) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CompartmentId  *string                           `json:"compartmentId"`
		DisplayName    *string                           `json:"displayName"`
		TimeAccepted   *common.SDKTime                   `json:"timeAccepted"`
		TimeStarted    *common.SDKTime                   `json:"timeStarted"`
		TimeFinished   *common.SDKTime                   `json:"timeFinished"`
		FreeformTags   map[string]string                 `json:"freeformTags"`
		DefinedTags    map[string]map[string]interface{} `json:"definedTags"`
		SystemTags     map[string]map[string]interface{} `json:"systemTags"`
		StreamId       *string                           `json:"streamId"`
		RunId          *string                           `json:"runId"`
		Features       []videofeature                    `json:"features"`
		OutputLocation *OutputLocation                   `json:"outputLocation"`
		LifecycleState StreamRunLifecycleStateEnum       `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.TimeAccepted = model.TimeAccepted

	m.TimeStarted = model.TimeStarted

	m.TimeFinished = model.TimeFinished

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.StreamId = model.StreamId

	m.RunId = model.RunId

	m.Features = make([]VideoFeature, len(model.Features))
	for i, n := range model.Features {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Features[i] = nn.(VideoFeature)
		} else {
			m.Features[i] = nil
		}
	}
	m.OutputLocation = model.OutputLocation

	m.LifecycleState = model.LifecycleState

	return
}

// StreamRunLifecycleStateEnum Enum with underlying type: string
type StreamRunLifecycleStateEnum string

// Set of constants representing the allowable values for StreamRunLifecycleStateEnum
const (
	StreamRunLifecycleStateAccepted   StreamRunLifecycleStateEnum = "ACCEPTED"
	StreamRunLifecycleStateInProgress StreamRunLifecycleStateEnum = "IN_PROGRESS"
	StreamRunLifecycleStateRunning    StreamRunLifecycleStateEnum = "RUNNING"
	StreamRunLifecycleStateStopping   StreamRunLifecycleStateEnum = "STOPPING"
	StreamRunLifecycleStateStopped    StreamRunLifecycleStateEnum = "STOPPED"
	StreamRunLifecycleStateFailed     StreamRunLifecycleStateEnum = "FAILED"
)

var mappingStreamRunLifecycleStateEnum = map[string]StreamRunLifecycleStateEnum{
	"ACCEPTED":    StreamRunLifecycleStateAccepted,
	"IN_PROGRESS": StreamRunLifecycleStateInProgress,
	"RUNNING":     StreamRunLifecycleStateRunning,
	"STOPPING":    StreamRunLifecycleStateStopping,
	"STOPPED":     StreamRunLifecycleStateStopped,
	"FAILED":      StreamRunLifecycleStateFailed,
}

var mappingStreamRunLifecycleStateEnumLowerCase = map[string]StreamRunLifecycleStateEnum{
	"accepted":    StreamRunLifecycleStateAccepted,
	"in_progress": StreamRunLifecycleStateInProgress,
	"running":     StreamRunLifecycleStateRunning,
	"stopping":    StreamRunLifecycleStateStopping,
	"stopped":     StreamRunLifecycleStateStopped,
	"failed":      StreamRunLifecycleStateFailed,
}

// GetStreamRunLifecycleStateEnumValues Enumerates the set of values for StreamRunLifecycleStateEnum
func GetStreamRunLifecycleStateEnumValues() []StreamRunLifecycleStateEnum {
	values := make([]StreamRunLifecycleStateEnum, 0)
	for _, v := range mappingStreamRunLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetStreamRunLifecycleStateEnumStringValues Enumerates the set of values in String for StreamRunLifecycleStateEnum
func GetStreamRunLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"RUNNING",
		"STOPPING",
		"STOPPED",
		"FAILED",
	}
}

// GetMappingStreamRunLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStreamRunLifecycleStateEnum(val string) (StreamRunLifecycleStateEnum, bool) {
	enum, ok := mappingStreamRunLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
