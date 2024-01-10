// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Document Understanding API
//
// Document AI helps customers perform various analysis on their documents. If a customer has lots of documents, they can process them in batch using asynchronous API endpoints.
//

package aidocument

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ProcessorJob Details of a processor job.
type ProcessorJob struct {

	// The id of the processor job.
	Id *string `mandatory:"true" json:"id"`

	// The compartment identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ProcessorConfig ProcessorConfig `mandatory:"true" json:"processorConfig"`

	// The job acceptance time.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	OutputLocation *OutputLocation `mandatory:"true" json:"outputLocation"`

	// The current state of the processor job.
	LifecycleState ProcessorJobLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The display name of the processor job.
	DisplayName *string `mandatory:"false" json:"displayName"`

	InputLocation InputLocation `mandatory:"false" json:"inputLocation"`

	// The job start time.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The job finish time.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// How much progress the operation has made, compared to the total amount of work to be performed.
	PercentComplete *float32 `mandatory:"false" json:"percentComplete"`

	// The detailed status of FAILED state.
	LifecycleDetails ProcessorJobLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`
}

func (m ProcessorJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProcessorJob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingProcessorJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetProcessorJobLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingProcessorJobLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetProcessorJobLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ProcessorJob) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName      *string                          `json:"displayName"`
		InputLocation    inputlocation                    `json:"inputLocation"`
		TimeStarted      *common.SDKTime                  `json:"timeStarted"`
		TimeFinished     *common.SDKTime                  `json:"timeFinished"`
		PercentComplete  *float32                         `json:"percentComplete"`
		LifecycleDetails ProcessorJobLifecycleDetailsEnum `json:"lifecycleDetails"`
		Id               *string                          `json:"id"`
		CompartmentId    *string                          `json:"compartmentId"`
		ProcessorConfig  processorconfig                  `json:"processorConfig"`
		TimeAccepted     *common.SDKTime                  `json:"timeAccepted"`
		OutputLocation   *OutputLocation                  `json:"outputLocation"`
		LifecycleState   ProcessorJobLifecycleStateEnum   `json:"lifecycleState"`
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

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	nn, e = model.ProcessorConfig.UnmarshalPolymorphicJSON(model.ProcessorConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ProcessorConfig = nn.(ProcessorConfig)
	} else {
		m.ProcessorConfig = nil
	}

	m.TimeAccepted = model.TimeAccepted

	m.OutputLocation = model.OutputLocation

	m.LifecycleState = model.LifecycleState

	return
}

// ProcessorJobLifecycleStateEnum Enum with underlying type: string
type ProcessorJobLifecycleStateEnum string

// Set of constants representing the allowable values for ProcessorJobLifecycleStateEnum
const (
	ProcessorJobLifecycleStateSucceeded  ProcessorJobLifecycleStateEnum = "SUCCEEDED"
	ProcessorJobLifecycleStateFailed     ProcessorJobLifecycleStateEnum = "FAILED"
	ProcessorJobLifecycleStateAccepted   ProcessorJobLifecycleStateEnum = "ACCEPTED"
	ProcessorJobLifecycleStateCanceled   ProcessorJobLifecycleStateEnum = "CANCELED"
	ProcessorJobLifecycleStateInProgress ProcessorJobLifecycleStateEnum = "IN_PROGRESS"
	ProcessorJobLifecycleStateCanceling  ProcessorJobLifecycleStateEnum = "CANCELING"
)

var mappingProcessorJobLifecycleStateEnum = map[string]ProcessorJobLifecycleStateEnum{
	"SUCCEEDED":   ProcessorJobLifecycleStateSucceeded,
	"FAILED":      ProcessorJobLifecycleStateFailed,
	"ACCEPTED":    ProcessorJobLifecycleStateAccepted,
	"CANCELED":    ProcessorJobLifecycleStateCanceled,
	"IN_PROGRESS": ProcessorJobLifecycleStateInProgress,
	"CANCELING":   ProcessorJobLifecycleStateCanceling,
}

var mappingProcessorJobLifecycleStateEnumLowerCase = map[string]ProcessorJobLifecycleStateEnum{
	"succeeded":   ProcessorJobLifecycleStateSucceeded,
	"failed":      ProcessorJobLifecycleStateFailed,
	"accepted":    ProcessorJobLifecycleStateAccepted,
	"canceled":    ProcessorJobLifecycleStateCanceled,
	"in_progress": ProcessorJobLifecycleStateInProgress,
	"canceling":   ProcessorJobLifecycleStateCanceling,
}

// GetProcessorJobLifecycleStateEnumValues Enumerates the set of values for ProcessorJobLifecycleStateEnum
func GetProcessorJobLifecycleStateEnumValues() []ProcessorJobLifecycleStateEnum {
	values := make([]ProcessorJobLifecycleStateEnum, 0)
	for _, v := range mappingProcessorJobLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetProcessorJobLifecycleStateEnumStringValues Enumerates the set of values in String for ProcessorJobLifecycleStateEnum
func GetProcessorJobLifecycleStateEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
		"ACCEPTED",
		"CANCELED",
		"IN_PROGRESS",
		"CANCELING",
	}
}

// GetMappingProcessorJobLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProcessorJobLifecycleStateEnum(val string) (ProcessorJobLifecycleStateEnum, bool) {
	enum, ok := mappingProcessorJobLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ProcessorJobLifecycleDetailsEnum Enum with underlying type: string
type ProcessorJobLifecycleDetailsEnum string

// Set of constants representing the allowable values for ProcessorJobLifecycleDetailsEnum
const (
	ProcessorJobLifecycleDetailsPartiallySucceeded ProcessorJobLifecycleDetailsEnum = "PARTIALLY_SUCCEEDED"
	ProcessorJobLifecycleDetailsCompletelyFailed   ProcessorJobLifecycleDetailsEnum = "COMPLETELY_FAILED"
)

var mappingProcessorJobLifecycleDetailsEnum = map[string]ProcessorJobLifecycleDetailsEnum{
	"PARTIALLY_SUCCEEDED": ProcessorJobLifecycleDetailsPartiallySucceeded,
	"COMPLETELY_FAILED":   ProcessorJobLifecycleDetailsCompletelyFailed,
}

var mappingProcessorJobLifecycleDetailsEnumLowerCase = map[string]ProcessorJobLifecycleDetailsEnum{
	"partially_succeeded": ProcessorJobLifecycleDetailsPartiallySucceeded,
	"completely_failed":   ProcessorJobLifecycleDetailsCompletelyFailed,
}

// GetProcessorJobLifecycleDetailsEnumValues Enumerates the set of values for ProcessorJobLifecycleDetailsEnum
func GetProcessorJobLifecycleDetailsEnumValues() []ProcessorJobLifecycleDetailsEnum {
	values := make([]ProcessorJobLifecycleDetailsEnum, 0)
	for _, v := range mappingProcessorJobLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetProcessorJobLifecycleDetailsEnumStringValues Enumerates the set of values in String for ProcessorJobLifecycleDetailsEnum
func GetProcessorJobLifecycleDetailsEnumStringValues() []string {
	return []string{
		"PARTIALLY_SUCCEEDED",
		"COMPLETELY_FAILED",
	}
}

// GetMappingProcessorJobLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProcessorJobLifecycleDetailsEnum(val string) (ProcessorJobLifecycleDetailsEnum, bool) {
	enum, ok := mappingProcessorJobLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
