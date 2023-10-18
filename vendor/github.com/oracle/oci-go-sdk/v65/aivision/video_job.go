// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// VideoJob Job details for a video analysis.
type VideoJob struct {

	// Id of the job.
	Id *string `mandatory:"true" json:"id"`

	// The ocid of the compartment that starts the job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// a list of document analysis features.
	Features []VideoFeature `mandatory:"true" json:"features"`

	// Job accepted time.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	OutputLocation *OutputLocation `mandatory:"true" json:"outputLocation"`

	// The current state of the batch document job.
	LifecycleState VideoJobLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Video job display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	InputLocation InputLocation `mandatory:"false" json:"inputLocation"`

	// Job started time.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Job finished time.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// How much progress the operation has made, vs the total amount of work that must be performed.
	PercentComplete *float32 `mandatory:"false" json:"percentComplete"`

	// Detailed status of FAILED state.
	LifecycleDetails VideoJobLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`

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

func (m VideoJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VideoJob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVideoJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVideoJobLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingVideoJobLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetVideoJobLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *VideoJob) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName      *string                           `json:"displayName"`
		InputLocation    inputlocation                     `json:"inputLocation"`
		TimeStarted      *common.SDKTime                   `json:"timeStarted"`
		TimeFinished     *common.SDKTime                   `json:"timeFinished"`
		PercentComplete  *float32                          `json:"percentComplete"`
		LifecycleDetails VideoJobLifecycleDetailsEnum      `json:"lifecycleDetails"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		SystemTags       map[string]map[string]interface{} `json:"systemTags"`
		Id               *string                           `json:"id"`
		CompartmentId    *string                           `json:"compartmentId"`
		Features         []videofeature                    `json:"features"`
		TimeAccepted     *common.SDKTime                   `json:"timeAccepted"`
		OutputLocation   *OutputLocation                   `json:"outputLocation"`
		LifecycleState   VideoJobLifecycleStateEnum        `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	nn, e = model.InputLocation.UnmarshalPolymorphicJSON(model.InputLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.InputLocation = nn.(InputLocation)
	} else {
		m.InputLocation = nil
	}

	m.TimeStarted = model.TimeStarted

	m.TimeFinished = model.TimeFinished

	m.PercentComplete = model.PercentComplete

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

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
	m.TimeAccepted = model.TimeAccepted

	m.OutputLocation = model.OutputLocation

	m.LifecycleState = model.LifecycleState

	return
}

// VideoJobLifecycleStateEnum Enum with underlying type: string
type VideoJobLifecycleStateEnum string

// Set of constants representing the allowable values for VideoJobLifecycleStateEnum
const (
	VideoJobLifecycleStateSucceeded  VideoJobLifecycleStateEnum = "SUCCEEDED"
	VideoJobLifecycleStateFailed     VideoJobLifecycleStateEnum = "FAILED"
	VideoJobLifecycleStateAccepted   VideoJobLifecycleStateEnum = "ACCEPTED"
	VideoJobLifecycleStateCanceled   VideoJobLifecycleStateEnum = "CANCELED"
	VideoJobLifecycleStateInProgress VideoJobLifecycleStateEnum = "IN_PROGRESS"
	VideoJobLifecycleStateCanceling  VideoJobLifecycleStateEnum = "CANCELING"
)

var mappingVideoJobLifecycleStateEnum = map[string]VideoJobLifecycleStateEnum{
	"SUCCEEDED":   VideoJobLifecycleStateSucceeded,
	"FAILED":      VideoJobLifecycleStateFailed,
	"ACCEPTED":    VideoJobLifecycleStateAccepted,
	"CANCELED":    VideoJobLifecycleStateCanceled,
	"IN_PROGRESS": VideoJobLifecycleStateInProgress,
	"CANCELING":   VideoJobLifecycleStateCanceling,
}

var mappingVideoJobLifecycleStateEnumLowerCase = map[string]VideoJobLifecycleStateEnum{
	"succeeded":   VideoJobLifecycleStateSucceeded,
	"failed":      VideoJobLifecycleStateFailed,
	"accepted":    VideoJobLifecycleStateAccepted,
	"canceled":    VideoJobLifecycleStateCanceled,
	"in_progress": VideoJobLifecycleStateInProgress,
	"canceling":   VideoJobLifecycleStateCanceling,
}

// GetVideoJobLifecycleStateEnumValues Enumerates the set of values for VideoJobLifecycleStateEnum
func GetVideoJobLifecycleStateEnumValues() []VideoJobLifecycleStateEnum {
	values := make([]VideoJobLifecycleStateEnum, 0)
	for _, v := range mappingVideoJobLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetVideoJobLifecycleStateEnumStringValues Enumerates the set of values in String for VideoJobLifecycleStateEnum
func GetVideoJobLifecycleStateEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
		"ACCEPTED",
		"CANCELED",
		"IN_PROGRESS",
		"CANCELING",
	}
}

// GetMappingVideoJobLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVideoJobLifecycleStateEnum(val string) (VideoJobLifecycleStateEnum, bool) {
	enum, ok := mappingVideoJobLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VideoJobLifecycleDetailsEnum Enum with underlying type: string
type VideoJobLifecycleDetailsEnum string

// Set of constants representing the allowable values for VideoJobLifecycleDetailsEnum
const (
	VideoJobLifecycleDetailsPartiallySucceeded VideoJobLifecycleDetailsEnum = "PARTIALLY_SUCCEEDED"
	VideoJobLifecycleDetailsCompletelyFailed   VideoJobLifecycleDetailsEnum = "COMPLETELY_FAILED"
)

var mappingVideoJobLifecycleDetailsEnum = map[string]VideoJobLifecycleDetailsEnum{
	"PARTIALLY_SUCCEEDED": VideoJobLifecycleDetailsPartiallySucceeded,
	"COMPLETELY_FAILED":   VideoJobLifecycleDetailsCompletelyFailed,
}

var mappingVideoJobLifecycleDetailsEnumLowerCase = map[string]VideoJobLifecycleDetailsEnum{
	"partially_succeeded": VideoJobLifecycleDetailsPartiallySucceeded,
	"completely_failed":   VideoJobLifecycleDetailsCompletelyFailed,
}

// GetVideoJobLifecycleDetailsEnumValues Enumerates the set of values for VideoJobLifecycleDetailsEnum
func GetVideoJobLifecycleDetailsEnumValues() []VideoJobLifecycleDetailsEnum {
	values := make([]VideoJobLifecycleDetailsEnum, 0)
	for _, v := range mappingVideoJobLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetVideoJobLifecycleDetailsEnumStringValues Enumerates the set of values in String for VideoJobLifecycleDetailsEnum
func GetVideoJobLifecycleDetailsEnumStringValues() []string {
	return []string{
		"PARTIALLY_SUCCEEDED",
		"COMPLETELY_FAILED",
	}
}

// GetMappingVideoJobLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVideoJobLifecycleDetailsEnum(val string) (VideoJobLifecycleDetailsEnum, bool) {
	enum, ok := mappingVideoJobLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
