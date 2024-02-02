// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Media Services API
//
// Media Services (includes Media Flow and Media Streams) is a fully managed service for processing media (video) source content. Use Media Flow and Media Streams to transcode and package digital video using configurable workflows and stream video outputs.
// Use the Media Services API to configure media workflows and run Media Flow jobs, create distribution channels, ingest assets, create Preview URLs and play assets. For more information, see Media Flow (https://docs.cloud.oracle.com/iaas/Content/dms-mediaflow/home.htm) and Media Streams (https://docs.cloud.oracle.com/iaas/Content/dms-mediastream/home.htm).
//

package mediaservices

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MediaWorkflowJob A MediaWorkflowJob represents a run of a MediaWorkflow for a specific set of parameters and configurations.
type MediaWorkflowJob struct {

	// The workflow to execute.
	MediaWorkflowId *string `mandatory:"true" json:"mediaWorkflowId"`

	// Unique identifier for this run of the workflow.
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Configurations to be applied to this run of the workflow.
	MediaWorkflowConfigurationIds []string `mandatory:"false" json:"mediaWorkflowConfigurationIds"`

	// Name of the Media Workflow Job. Does not have to be unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The current state of the MediaWorkflowJob.
	LifecycleState MediaWorkflowJobLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The lifecyle details.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Status of each task.
	TaskLifecycleState []MediaWorkflowTaskState `mandatory:"false" json:"taskLifecycleState"`

	// Parameters that override parameters specified in MediaWorkflowTaskDeclarations, the MediaWorkflow,
	// the MediaWorkflow's MediaWorkflowConfigurations and the MediaWorkflowConfigurations of this
	// MediaWorkflowJob. The parameters are given as JSON.  The top level and 2nd level elements must be
	// JSON objects (vs arrays, scalars, etc). The top level keys refer to a task's key and the 2nd level
	// keys refer to a parameter's name.
	Parameters map[string]interface{} `mandatory:"false" json:"parameters"`

	// Creation time of the job. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Updated time of the job. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A JSON representation of the job as it will be run by the system. All the task declarations, configurations
	// and parameters are merged. Parameter values are all fully resolved.
	Runnable map[string]interface{} `mandatory:"false" json:"runnable"`

	// A list of JobOutput for the workflowJob.
	Outputs []JobOutput `mandatory:"false" json:"outputs"`

	// Time when the job started to execute. An RFC3339 formatted datetime string.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Time when the job finished. An RFC3339 formatted datetime string.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m MediaWorkflowJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MediaWorkflowJob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMediaWorkflowJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMediaWorkflowJobLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MediaWorkflowJobLifecycleStateEnum Enum with underlying type: string
type MediaWorkflowJobLifecycleStateEnum string

// Set of constants representing the allowable values for MediaWorkflowJobLifecycleStateEnum
const (
	MediaWorkflowJobLifecycleStateAccepted   MediaWorkflowJobLifecycleStateEnum = "ACCEPTED"
	MediaWorkflowJobLifecycleStateInProgress MediaWorkflowJobLifecycleStateEnum = "IN_PROGRESS"
	MediaWorkflowJobLifecycleStateWaiting    MediaWorkflowJobLifecycleStateEnum = "WAITING"
	MediaWorkflowJobLifecycleStateRejected   MediaWorkflowJobLifecycleStateEnum = "REJECTED"
	MediaWorkflowJobLifecycleStateFailed     MediaWorkflowJobLifecycleStateEnum = "FAILED"
	MediaWorkflowJobLifecycleStateSucceeded  MediaWorkflowJobLifecycleStateEnum = "SUCCEEDED"
	MediaWorkflowJobLifecycleStateCanceling  MediaWorkflowJobLifecycleStateEnum = "CANCELING"
	MediaWorkflowJobLifecycleStateCanceled   MediaWorkflowJobLifecycleStateEnum = "CANCELED"
)

var mappingMediaWorkflowJobLifecycleStateEnum = map[string]MediaWorkflowJobLifecycleStateEnum{
	"ACCEPTED":    MediaWorkflowJobLifecycleStateAccepted,
	"IN_PROGRESS": MediaWorkflowJobLifecycleStateInProgress,
	"WAITING":     MediaWorkflowJobLifecycleStateWaiting,
	"REJECTED":    MediaWorkflowJobLifecycleStateRejected,
	"FAILED":      MediaWorkflowJobLifecycleStateFailed,
	"SUCCEEDED":   MediaWorkflowJobLifecycleStateSucceeded,
	"CANCELING":   MediaWorkflowJobLifecycleStateCanceling,
	"CANCELED":    MediaWorkflowJobLifecycleStateCanceled,
}

var mappingMediaWorkflowJobLifecycleStateEnumLowerCase = map[string]MediaWorkflowJobLifecycleStateEnum{
	"accepted":    MediaWorkflowJobLifecycleStateAccepted,
	"in_progress": MediaWorkflowJobLifecycleStateInProgress,
	"waiting":     MediaWorkflowJobLifecycleStateWaiting,
	"rejected":    MediaWorkflowJobLifecycleStateRejected,
	"failed":      MediaWorkflowJobLifecycleStateFailed,
	"succeeded":   MediaWorkflowJobLifecycleStateSucceeded,
	"canceling":   MediaWorkflowJobLifecycleStateCanceling,
	"canceled":    MediaWorkflowJobLifecycleStateCanceled,
}

// GetMediaWorkflowJobLifecycleStateEnumValues Enumerates the set of values for MediaWorkflowJobLifecycleStateEnum
func GetMediaWorkflowJobLifecycleStateEnumValues() []MediaWorkflowJobLifecycleStateEnum {
	values := make([]MediaWorkflowJobLifecycleStateEnum, 0)
	for _, v := range mappingMediaWorkflowJobLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMediaWorkflowJobLifecycleStateEnumStringValues Enumerates the set of values in String for MediaWorkflowJobLifecycleStateEnum
func GetMediaWorkflowJobLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"WAITING",
		"REJECTED",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingMediaWorkflowJobLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMediaWorkflowJobLifecycleStateEnum(val string) (MediaWorkflowJobLifecycleStateEnum, bool) {
	enum, ok := mappingMediaWorkflowJobLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
