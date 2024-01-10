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

// MediaWorkflow Configurable workflows that define the series of tasks that will be used to process video files.
type MediaWorkflow struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Name of the Media Workflow. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The processing to be done in this workflow. Each key of the MediaWorkflowTasks in this array is unique
	// within the array.  The order of the items is preserved from the order of the tasks array in
	// CreateMediaWorkflowDetails or UpdateMediaWorkflowDetails.
	Tasks []MediaWorkflowTask `mandatory:"false" json:"tasks"`

	// Configurations to be applied to all the runs of this workflow. Parameters in these configurations are
	// overridden by parameters in the MediaWorkflowConfigurations of the MediaWorkflowJob and the
	// parameters of the MediaWorkflowJob. If the same parameter appears in multiple configurations, the values that
	// appear in the configuration at the highest index will be used.
	MediaWorkflowConfigurationIds []string `mandatory:"false" json:"mediaWorkflowConfigurationIds"`

	// JSON object representing named parameters and their default values that can be referenced throughout this workflow.
	// The values declared here can be overridden by the MediaWorkflowConfigurations or parameters supplied when creating
	// MediaWorkflowJobs from this MediaWorkflow.
	Parameters map[string]interface{} `mandatory:"false" json:"parameters"`

	// The time when the MediaWorkflow was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time when the MediaWorkflow was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the MediaWorkflow.
	LifecycleState MediaWorkflowLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecyleDetails *string `mandatory:"false" json:"lifecyleDetails"`

	// The version of the MediaWorkflow.
	Version *int64 `mandatory:"false" json:"version"`

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

func (m MediaWorkflow) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MediaWorkflow) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMediaWorkflowLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMediaWorkflowLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MediaWorkflowLifecycleStateEnum Enum with underlying type: string
type MediaWorkflowLifecycleStateEnum string

// Set of constants representing the allowable values for MediaWorkflowLifecycleStateEnum
const (
	MediaWorkflowLifecycleStateActive         MediaWorkflowLifecycleStateEnum = "ACTIVE"
	MediaWorkflowLifecycleStateNeedsAttention MediaWorkflowLifecycleStateEnum = "NEEDS_ATTENTION"
	MediaWorkflowLifecycleStateDeleted        MediaWorkflowLifecycleStateEnum = "DELETED"
)

var mappingMediaWorkflowLifecycleStateEnum = map[string]MediaWorkflowLifecycleStateEnum{
	"ACTIVE":          MediaWorkflowLifecycleStateActive,
	"NEEDS_ATTENTION": MediaWorkflowLifecycleStateNeedsAttention,
	"DELETED":         MediaWorkflowLifecycleStateDeleted,
}

var mappingMediaWorkflowLifecycleStateEnumLowerCase = map[string]MediaWorkflowLifecycleStateEnum{
	"active":          MediaWorkflowLifecycleStateActive,
	"needs_attention": MediaWorkflowLifecycleStateNeedsAttention,
	"deleted":         MediaWorkflowLifecycleStateDeleted,
}

// GetMediaWorkflowLifecycleStateEnumValues Enumerates the set of values for MediaWorkflowLifecycleStateEnum
func GetMediaWorkflowLifecycleStateEnumValues() []MediaWorkflowLifecycleStateEnum {
	values := make([]MediaWorkflowLifecycleStateEnum, 0)
	for _, v := range mappingMediaWorkflowLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMediaWorkflowLifecycleStateEnumStringValues Enumerates the set of values in String for MediaWorkflowLifecycleStateEnum
func GetMediaWorkflowLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"NEEDS_ATTENTION",
		"DELETED",
	}
}

// GetMappingMediaWorkflowLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMediaWorkflowLifecycleStateEnum(val string) (MediaWorkflowLifecycleStateEnum, bool) {
	enum, ok := mappingMediaWorkflowLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
