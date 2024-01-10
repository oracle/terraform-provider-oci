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

// ImageJob The job details for a batch image analysis.
type ImageJob struct {

	// The job id
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that starts the job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The list of requested document analysis types.
	Features []ImageFeature `mandatory:"true" json:"features"`

	// The job acceptance time.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	OutputLocation *OutputLocation `mandatory:"true" json:"outputLocation"`

	// The current state of the batch image job.
	LifecycleState ImageJobLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The image job display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	InputLocation InputLocation `mandatory:"false" json:"inputLocation"`

	// The job start time.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The job finish time.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// How much progress the operation has made, compared to the total amount of work to be performed.
	PercentComplete *float32 `mandatory:"false" json:"percentComplete"`

	// The detailed status of FAILED state.
	LifecycleDetails ImageJobLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// Whether or not to generate a ZIP file containing the results.
	IsZipOutputEnabled *bool `mandatory:"false" json:"isZipOutputEnabled"`
}

func (m ImageJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImageJob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingImageJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetImageJobLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingImageJobLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetImageJobLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ImageJob) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName        *string                      `json:"displayName"`
		InputLocation      inputlocation                `json:"inputLocation"`
		TimeStarted        *common.SDKTime              `json:"timeStarted"`
		TimeFinished       *common.SDKTime              `json:"timeFinished"`
		PercentComplete    *float32                     `json:"percentComplete"`
		LifecycleDetails   ImageJobLifecycleDetailsEnum `json:"lifecycleDetails"`
		IsZipOutputEnabled *bool                        `json:"isZipOutputEnabled"`
		Id                 *string                      `json:"id"`
		CompartmentId      *string                      `json:"compartmentId"`
		Features           []imagefeature               `json:"features"`
		TimeAccepted       *common.SDKTime              `json:"timeAccepted"`
		OutputLocation     *OutputLocation              `json:"outputLocation"`
		LifecycleState     ImageJobLifecycleStateEnum   `json:"lifecycleState"`
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

	m.IsZipOutputEnabled = model.IsZipOutputEnabled

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.Features = make([]ImageFeature, len(model.Features))
	for i, n := range model.Features {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Features[i] = nn.(ImageFeature)
		} else {
			m.Features[i] = nil
		}
	}
	m.TimeAccepted = model.TimeAccepted

	m.OutputLocation = model.OutputLocation

	m.LifecycleState = model.LifecycleState

	return
}

// ImageJobLifecycleStateEnum Enum with underlying type: string
type ImageJobLifecycleStateEnum string

// Set of constants representing the allowable values for ImageJobLifecycleStateEnum
const (
	ImageJobLifecycleStateSucceeded  ImageJobLifecycleStateEnum = "SUCCEEDED"
	ImageJobLifecycleStateFailed     ImageJobLifecycleStateEnum = "FAILED"
	ImageJobLifecycleStateAccepted   ImageJobLifecycleStateEnum = "ACCEPTED"
	ImageJobLifecycleStateCanceled   ImageJobLifecycleStateEnum = "CANCELED"
	ImageJobLifecycleStateInProgress ImageJobLifecycleStateEnum = "IN_PROGRESS"
	ImageJobLifecycleStateCanceling  ImageJobLifecycleStateEnum = "CANCELING"
)

var mappingImageJobLifecycleStateEnum = map[string]ImageJobLifecycleStateEnum{
	"SUCCEEDED":   ImageJobLifecycleStateSucceeded,
	"FAILED":      ImageJobLifecycleStateFailed,
	"ACCEPTED":    ImageJobLifecycleStateAccepted,
	"CANCELED":    ImageJobLifecycleStateCanceled,
	"IN_PROGRESS": ImageJobLifecycleStateInProgress,
	"CANCELING":   ImageJobLifecycleStateCanceling,
}

var mappingImageJobLifecycleStateEnumLowerCase = map[string]ImageJobLifecycleStateEnum{
	"succeeded":   ImageJobLifecycleStateSucceeded,
	"failed":      ImageJobLifecycleStateFailed,
	"accepted":    ImageJobLifecycleStateAccepted,
	"canceled":    ImageJobLifecycleStateCanceled,
	"in_progress": ImageJobLifecycleStateInProgress,
	"canceling":   ImageJobLifecycleStateCanceling,
}

// GetImageJobLifecycleStateEnumValues Enumerates the set of values for ImageJobLifecycleStateEnum
func GetImageJobLifecycleStateEnumValues() []ImageJobLifecycleStateEnum {
	values := make([]ImageJobLifecycleStateEnum, 0)
	for _, v := range mappingImageJobLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetImageJobLifecycleStateEnumStringValues Enumerates the set of values in String for ImageJobLifecycleStateEnum
func GetImageJobLifecycleStateEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
		"ACCEPTED",
		"CANCELED",
		"IN_PROGRESS",
		"CANCELING",
	}
}

// GetMappingImageJobLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImageJobLifecycleStateEnum(val string) (ImageJobLifecycleStateEnum, bool) {
	enum, ok := mappingImageJobLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ImageJobLifecycleDetailsEnum Enum with underlying type: string
type ImageJobLifecycleDetailsEnum string

// Set of constants representing the allowable values for ImageJobLifecycleDetailsEnum
const (
	ImageJobLifecycleDetailsPartiallySucceeded ImageJobLifecycleDetailsEnum = "PARTIALLY_SUCCEEDED"
	ImageJobLifecycleDetailsCompletelyFailed   ImageJobLifecycleDetailsEnum = "COMPLETELY_FAILED"
)

var mappingImageJobLifecycleDetailsEnum = map[string]ImageJobLifecycleDetailsEnum{
	"PARTIALLY_SUCCEEDED": ImageJobLifecycleDetailsPartiallySucceeded,
	"COMPLETELY_FAILED":   ImageJobLifecycleDetailsCompletelyFailed,
}

var mappingImageJobLifecycleDetailsEnumLowerCase = map[string]ImageJobLifecycleDetailsEnum{
	"partially_succeeded": ImageJobLifecycleDetailsPartiallySucceeded,
	"completely_failed":   ImageJobLifecycleDetailsCompletelyFailed,
}

// GetImageJobLifecycleDetailsEnumValues Enumerates the set of values for ImageJobLifecycleDetailsEnum
func GetImageJobLifecycleDetailsEnumValues() []ImageJobLifecycleDetailsEnum {
	values := make([]ImageJobLifecycleDetailsEnum, 0)
	for _, v := range mappingImageJobLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetImageJobLifecycleDetailsEnumStringValues Enumerates the set of values in String for ImageJobLifecycleDetailsEnum
func GetImageJobLifecycleDetailsEnumStringValues() []string {
	return []string{
		"PARTIALLY_SUCCEEDED",
		"COMPLETELY_FAILED",
	}
}

// GetMappingImageJobLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImageJobLifecycleDetailsEnum(val string) (ImageJobLifecycleDetailsEnum, bool) {
	enum, ok := mappingImageJobLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
