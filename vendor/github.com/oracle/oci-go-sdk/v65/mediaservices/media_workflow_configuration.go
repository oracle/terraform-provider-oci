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

// MediaWorkflowConfiguration Resusable set of values that can be referenced either in a MediaWorkflow or when running a MediaWorkflowJob.
type MediaWorkflowConfiguration struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Display name for the MediaWorkflowConfiguration. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Reuseable parameter values encoded as a JSON; the top and second level JSON elements are
	// objects. Each key of the top level object refer to a task key that is unqiue to the
	// workflow, each of the second level objects' keys refer to the name of a parameter that is
	// unique to the task. taskKey -> parameterName -> parameterValue
	Parameters map[string]interface{} `mandatory:"true" json:"parameters"`

	// The time when the the MediaWorkflowConfiguration was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time when the MediaWorkflowConfiguration was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the MediaWorkflowConfiguration.
	LifecycleState MediaWorkflowConfigurationLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecyleDetails *string `mandatory:"false" json:"lifecyleDetails"`

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

func (m MediaWorkflowConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MediaWorkflowConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMediaWorkflowConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMediaWorkflowConfigurationLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MediaWorkflowConfigurationLifecycleStateEnum Enum with underlying type: string
type MediaWorkflowConfigurationLifecycleStateEnum string

// Set of constants representing the allowable values for MediaWorkflowConfigurationLifecycleStateEnum
const (
	MediaWorkflowConfigurationLifecycleStateActive  MediaWorkflowConfigurationLifecycleStateEnum = "ACTIVE"
	MediaWorkflowConfigurationLifecycleStateDeleted MediaWorkflowConfigurationLifecycleStateEnum = "DELETED"
)

var mappingMediaWorkflowConfigurationLifecycleStateEnum = map[string]MediaWorkflowConfigurationLifecycleStateEnum{
	"ACTIVE":  MediaWorkflowConfigurationLifecycleStateActive,
	"DELETED": MediaWorkflowConfigurationLifecycleStateDeleted,
}

var mappingMediaWorkflowConfigurationLifecycleStateEnumLowerCase = map[string]MediaWorkflowConfigurationLifecycleStateEnum{
	"active":  MediaWorkflowConfigurationLifecycleStateActive,
	"deleted": MediaWorkflowConfigurationLifecycleStateDeleted,
}

// GetMediaWorkflowConfigurationLifecycleStateEnumValues Enumerates the set of values for MediaWorkflowConfigurationLifecycleStateEnum
func GetMediaWorkflowConfigurationLifecycleStateEnumValues() []MediaWorkflowConfigurationLifecycleStateEnum {
	values := make([]MediaWorkflowConfigurationLifecycleStateEnum, 0)
	for _, v := range mappingMediaWorkflowConfigurationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMediaWorkflowConfigurationLifecycleStateEnumStringValues Enumerates the set of values in String for MediaWorkflowConfigurationLifecycleStateEnum
func GetMediaWorkflowConfigurationLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingMediaWorkflowConfigurationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMediaWorkflowConfigurationLifecycleStateEnum(val string) (MediaWorkflowConfigurationLifecycleStateEnum, bool) {
	enum, ok := mappingMediaWorkflowConfigurationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
