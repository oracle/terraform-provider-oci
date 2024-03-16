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

// Job Job details which contain input document details on which prediction need to run, features (which and all language services ) need to run and where to store results
type Job struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job.
	Id *string `mandatory:"true" json:"id"`

	InputLocation InputLocation `mandatory:"true" json:"inputLocation"`

	// training model details
	// For this release only one model is allowed to be input here.
	// One of the three modelType, ModelId, endpointId should be given other wise error will be thrown from API
	ModelMetadataDetails []ModelMetadataDetails `mandatory:"true" json:"modelMetadataDetails"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	OutputLocation *ObjectPrefixOutputLocation `mandatory:"true" json:"outputLocation"`

	// A user-friendly display name for the job.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the job.
	Description *string `mandatory:"false" json:"description"`

	InputConfiguration *InputConfiguration `mandatory:"false" json:"inputConfiguration"`

	// The current state of the Job.
	LifecycleState JobLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// How much progress the operation has made, vs the total amount of work that must be performed.
	PercentComplete *int `mandatory:"false" json:"percentComplete"`

	// Total number of documents given as input for prediction. For CSV this signifies number of rows and for TXT this signifies number of files.
	TotalDocuments *int `mandatory:"false" json:"totalDocuments"`

	// Number of documents still to process. For CSV this signifies number of rows and for TXT this signifies number of files.
	PendingDocuments *int `mandatory:"false" json:"pendingDocuments"`

	// Number of documents processed for prediction. For CSV this signifies number of rows and for TXT this signifies number of files.
	CompletedDocuments *int `mandatory:"false" json:"completedDocuments"`

	// Number of documents failed for prediction. For CSV this signifies number of rows and for TXT this signifies number of files.
	FailedDocuments *int `mandatory:"false" json:"failedDocuments"`

	// warnings count
	WarningsCount *int `mandatory:"false" json:"warningsCount"`

	// Time to live duration in days for Job. Job will be available till max 90 days.
	TtlInDays *int `mandatory:"false" json:"ttlInDays"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the job.
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// Job accepted time.
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// Job started time.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Job finished time.
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`
}

func (m Job) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Job) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetJobLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Job) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName          *string                     `json:"displayName"`
		Description          *string                     `json:"description"`
		InputConfiguration   *InputConfiguration         `json:"inputConfiguration"`
		LifecycleState       JobLifecycleStateEnum       `json:"lifecycleState"`
		LifecycleDetails     *string                     `json:"lifecycleDetails"`
		PercentComplete      *int                        `json:"percentComplete"`
		TotalDocuments       *int                        `json:"totalDocuments"`
		PendingDocuments     *int                        `json:"pendingDocuments"`
		CompletedDocuments   *int                        `json:"completedDocuments"`
		FailedDocuments      *int                        `json:"failedDocuments"`
		WarningsCount        *int                        `json:"warningsCount"`
		TtlInDays            *int                        `json:"ttlInDays"`
		CreatedBy            *string                     `json:"createdBy"`
		TimeAccepted         *common.SDKTime             `json:"timeAccepted"`
		TimeStarted          *common.SDKTime             `json:"timeStarted"`
		TimeCompleted        *common.SDKTime             `json:"timeCompleted"`
		Id                   *string                     `json:"id"`
		InputLocation        inputlocation               `json:"inputLocation"`
		ModelMetadataDetails []ModelMetadataDetails      `json:"modelMetadataDetails"`
		CompartmentId        *string                     `json:"compartmentId"`
		OutputLocation       *ObjectPrefixOutputLocation `json:"outputLocation"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.InputConfiguration = model.InputConfiguration

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.PercentComplete = model.PercentComplete

	m.TotalDocuments = model.TotalDocuments

	m.PendingDocuments = model.PendingDocuments

	m.CompletedDocuments = model.CompletedDocuments

	m.FailedDocuments = model.FailedDocuments

	m.WarningsCount = model.WarningsCount

	m.TtlInDays = model.TtlInDays

	m.CreatedBy = model.CreatedBy

	m.TimeAccepted = model.TimeAccepted

	m.TimeStarted = model.TimeStarted

	m.TimeCompleted = model.TimeCompleted

	m.Id = model.Id

	nn, e = model.InputLocation.UnmarshalPolymorphicJSON(model.InputLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.InputLocation = nn.(InputLocation)
	} else {
		m.InputLocation = nil
	}

	m.ModelMetadataDetails = make([]ModelMetadataDetails, len(model.ModelMetadataDetails))
	copy(m.ModelMetadataDetails, model.ModelMetadataDetails)
	m.CompartmentId = model.CompartmentId

	m.OutputLocation = model.OutputLocation

	return
}

// JobLifecycleStateEnum Enum with underlying type: string
type JobLifecycleStateEnum string

// Set of constants representing the allowable values for JobLifecycleStateEnum
const (
	JobLifecycleStateAccepted   JobLifecycleStateEnum = "ACCEPTED"
	JobLifecycleStateInProgress JobLifecycleStateEnum = "IN_PROGRESS"
	JobLifecycleStateSucceeded  JobLifecycleStateEnum = "SUCCEEDED"
	JobLifecycleStateFailed     JobLifecycleStateEnum = "FAILED"
	JobLifecycleStateCanceling  JobLifecycleStateEnum = "CANCELING"
	JobLifecycleStateCanceled   JobLifecycleStateEnum = "CANCELED"
	JobLifecycleStateDeleting   JobLifecycleStateEnum = "DELETING"
	JobLifecycleStateDeleted    JobLifecycleStateEnum = "DELETED"
)

var mappingJobLifecycleStateEnum = map[string]JobLifecycleStateEnum{
	"ACCEPTED":    JobLifecycleStateAccepted,
	"IN_PROGRESS": JobLifecycleStateInProgress,
	"SUCCEEDED":   JobLifecycleStateSucceeded,
	"FAILED":      JobLifecycleStateFailed,
	"CANCELING":   JobLifecycleStateCanceling,
	"CANCELED":    JobLifecycleStateCanceled,
	"DELETING":    JobLifecycleStateDeleting,
	"DELETED":     JobLifecycleStateDeleted,
}

var mappingJobLifecycleStateEnumLowerCase = map[string]JobLifecycleStateEnum{
	"accepted":    JobLifecycleStateAccepted,
	"in_progress": JobLifecycleStateInProgress,
	"succeeded":   JobLifecycleStateSucceeded,
	"failed":      JobLifecycleStateFailed,
	"canceling":   JobLifecycleStateCanceling,
	"canceled":    JobLifecycleStateCanceled,
	"deleting":    JobLifecycleStateDeleting,
	"deleted":     JobLifecycleStateDeleted,
}

// GetJobLifecycleStateEnumValues Enumerates the set of values for JobLifecycleStateEnum
func GetJobLifecycleStateEnumValues() []JobLifecycleStateEnum {
	values := make([]JobLifecycleStateEnum, 0)
	for _, v := range mappingJobLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetJobLifecycleStateEnumStringValues Enumerates the set of values in String for JobLifecycleStateEnum
func GetJobLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"CANCELING",
		"CANCELED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingJobLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobLifecycleStateEnum(val string) (JobLifecycleStateEnum, bool) {
	enum, ok := mappingJobLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
