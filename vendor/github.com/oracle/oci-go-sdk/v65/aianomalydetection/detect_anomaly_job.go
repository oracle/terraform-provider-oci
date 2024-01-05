// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Anomaly Detection API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DetectAnomalyJob Anomaly Job contains information for asynchronous detection of anomalies.
type DetectAnomalyJob struct {

	// Id of the job.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that starts the job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the trained model.
	ModelId *string `mandatory:"true" json:"modelId"`

	InputDetails InputJobDetails `mandatory:"true" json:"inputDetails"`

	OutputDetails OutputJobDetails `mandatory:"true" json:"outputDetails"`

	// Job accepted time
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The current state of the batch document job.
	LifecycleState DetectAnomalyJobLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Detect anomaly job display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detect anomaly job description.
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the project.
	ProjectId *string `mandatory:"false" json:"projectId"`

	// The value that customer can adjust to control the sensitivity of anomaly detection
	Sensitivity *float32 `mandatory:"false" json:"sensitivity"`

	// Flag to enable the service to return estimates for all data points rather than just the anomalous data points
	AreAllEstimatesRequired *bool `mandatory:"false" json:"areAllEstimatesRequired"`

	// Job started time
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Job finished time
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// The current state details of the batch document job.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{ "orcl-cloud": { "free-tier-retained": "true" } }`
	SystemTags map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DetectAnomalyJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DetectAnomalyJob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDetectAnomalyJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDetectAnomalyJobLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DetectAnomalyJob) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName             *string                            `json:"displayName"`
		Description             *string                            `json:"description"`
		ProjectId               *string                            `json:"projectId"`
		Sensitivity             *float32                           `json:"sensitivity"`
		AreAllEstimatesRequired *bool                              `json:"areAllEstimatesRequired"`
		TimeStarted             *common.SDKTime                    `json:"timeStarted"`
		TimeFinished            *common.SDKTime                    `json:"timeFinished"`
		LifecycleStateDetails   *string                            `json:"lifecycleStateDetails"`
		FreeformTags            map[string]string                  `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{}  `json:"definedTags"`
		SystemTags              map[string]interface{}             `json:"systemTags"`
		Id                      *string                            `json:"id"`
		CompartmentId           *string                            `json:"compartmentId"`
		ModelId                 *string                            `json:"modelId"`
		InputDetails            inputjobdetails                    `json:"inputDetails"`
		OutputDetails           outputjobdetails                   `json:"outputDetails"`
		TimeAccepted            *common.SDKTime                    `json:"timeAccepted"`
		LifecycleState          DetectAnomalyJobLifecycleStateEnum `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.ProjectId = model.ProjectId

	m.Sensitivity = model.Sensitivity

	m.AreAllEstimatesRequired = model.AreAllEstimatesRequired

	m.TimeStarted = model.TimeStarted

	m.TimeFinished = model.TimeFinished

	m.LifecycleStateDetails = model.LifecycleStateDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.ModelId = model.ModelId

	nn, e = model.InputDetails.UnmarshalPolymorphicJSON(model.InputDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.InputDetails = nn.(InputJobDetails)
	} else {
		m.InputDetails = nil
	}

	nn, e = model.OutputDetails.UnmarshalPolymorphicJSON(model.OutputDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.OutputDetails = nn.(OutputJobDetails)
	} else {
		m.OutputDetails = nil
	}

	m.TimeAccepted = model.TimeAccepted

	m.LifecycleState = model.LifecycleState

	return
}

// DetectAnomalyJobLifecycleStateEnum Enum with underlying type: string
type DetectAnomalyJobLifecycleStateEnum string

// Set of constants representing the allowable values for DetectAnomalyJobLifecycleStateEnum
const (
	DetectAnomalyJobLifecycleStateSucceeded          DetectAnomalyJobLifecycleStateEnum = "SUCCEEDED"
	DetectAnomalyJobLifecycleStatePartiallySucceeded DetectAnomalyJobLifecycleStateEnum = "PARTIALLY_SUCCEEDED"
	DetectAnomalyJobLifecycleStateFailed             DetectAnomalyJobLifecycleStateEnum = "FAILED"
	DetectAnomalyJobLifecycleStateAccepted           DetectAnomalyJobLifecycleStateEnum = "ACCEPTED"
	DetectAnomalyJobLifecycleStateCanceled           DetectAnomalyJobLifecycleStateEnum = "CANCELED"
	DetectAnomalyJobLifecycleStateInProgress         DetectAnomalyJobLifecycleStateEnum = "IN_PROGRESS"
)

var mappingDetectAnomalyJobLifecycleStateEnum = map[string]DetectAnomalyJobLifecycleStateEnum{
	"SUCCEEDED":           DetectAnomalyJobLifecycleStateSucceeded,
	"PARTIALLY_SUCCEEDED": DetectAnomalyJobLifecycleStatePartiallySucceeded,
	"FAILED":              DetectAnomalyJobLifecycleStateFailed,
	"ACCEPTED":            DetectAnomalyJobLifecycleStateAccepted,
	"CANCELED":            DetectAnomalyJobLifecycleStateCanceled,
	"IN_PROGRESS":         DetectAnomalyJobLifecycleStateInProgress,
}

var mappingDetectAnomalyJobLifecycleStateEnumLowerCase = map[string]DetectAnomalyJobLifecycleStateEnum{
	"succeeded":           DetectAnomalyJobLifecycleStateSucceeded,
	"partially_succeeded": DetectAnomalyJobLifecycleStatePartiallySucceeded,
	"failed":              DetectAnomalyJobLifecycleStateFailed,
	"accepted":            DetectAnomalyJobLifecycleStateAccepted,
	"canceled":            DetectAnomalyJobLifecycleStateCanceled,
	"in_progress":         DetectAnomalyJobLifecycleStateInProgress,
}

// GetDetectAnomalyJobLifecycleStateEnumValues Enumerates the set of values for DetectAnomalyJobLifecycleStateEnum
func GetDetectAnomalyJobLifecycleStateEnumValues() []DetectAnomalyJobLifecycleStateEnum {
	values := make([]DetectAnomalyJobLifecycleStateEnum, 0)
	for _, v := range mappingDetectAnomalyJobLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDetectAnomalyJobLifecycleStateEnumStringValues Enumerates the set of values in String for DetectAnomalyJobLifecycleStateEnum
func GetDetectAnomalyJobLifecycleStateEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"PARTIALLY_SUCCEEDED",
		"FAILED",
		"ACCEPTED",
		"CANCELED",
		"IN_PROGRESS",
	}
}

// GetMappingDetectAnomalyJobLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDetectAnomalyJobLifecycleStateEnum(val string) (DetectAnomalyJobLifecycleStateEnum, bool) {
	enum, ok := mappingDetectAnomalyJobLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
