// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Container A single container on a Container Instance.
// If this Container is DELETED, the record will remain visible for a short period
// of time before being permanently removed.
type Container struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Display name for the Container. Can be renamed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Availability Domain where the Container's Instance is running.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The current state of the Container.
	LifecycleState ContainerLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the the Container was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The identifier of the Container Instance on which this container is running.
	ContainerInstanceId *string `mandatory:"true" json:"containerInstanceId"`

	// The container image information. Currently only support public docker registry. Can be either image name,
	// e.g `containerImage`, image name with version, e.g `containerImage:v1` or complete docker image Url e.g
	// `docker.io/library/containerImage:latest`.
	// If no registry is provided, will default the registry to public docker hub `docker.io/library`.
	// The registry used for container image must be reachable over the Container Instance's VNIC.
	ImageUrl *string `mandatory:"true" json:"imageUrl"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Fault Domain where the Container's Instance is running.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// A message describing the current state in more detail. For example, can be used to provide
	// actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The exit code of the container process if it has stopped executing.
	ExitCode *int `mandatory:"false" json:"exitCode"`

	// Time at which the container last terminated. An RFC3339 formatted datetime string
	TimeTerminated *common.SDKTime `mandatory:"false" json:"timeTerminated"`

	// The time the Container was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// This command will override the container's entrypoint process.
	// If not specified, the existing entrypoint process defined in the image will be used.
	Command []string `mandatory:"false" json:"command"`

	// A list of string arguments for a Container's entrypoint process.
	// Many containers use an entrypoint process pointing to a shell,
	// for example /bin/bash. For such containers, this argument list
	// can also be used to specify the main command in the container process.
	Arguments []string `mandatory:"false" json:"arguments"`

	// A list of additional configurable container capabilities
	AdditionalCapabilities []ContainerCapabilityEnum `mandatory:"false" json:"additionalCapabilities,omitempty"`

	// The working directory within the Container's filesystem for
	// the Container process. If this is not present, the default
	// working directory from the image will be used.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`

	// A map of additional environment variables to set in the environment of the container's
	// entrypoint process. These variables are in addition to any variables already defined
	// in the container's image.
	EnvironmentVariables map[string]string `mandatory:"false" json:"environmentVariables"`

	// List of the volume mounts.
	VolumeMounts []VolumeMount `mandatory:"false" json:"volumeMounts"`

	// List of container health checks
	HealthChecks []ContainerHealthCheck `mandatory:"false" json:"healthChecks"`

	// Determines if the Container will have access to the Container Instance Resource Principal.
	// This method utilizes resource principal version 2.2. Please refer to
	// https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdk_authentication_methods.htm#sdk_authentication_methods_resource_principal
	// for detailed explanation of how to leverage the exposed resource principal elements.
	IsResourcePrincipalDisabled *bool `mandatory:"false" json:"isResourcePrincipalDisabled"`

	ResourceConfig *ContainerResourceConfig `mandatory:"false" json:"resourceConfig"`

	// The number of container restart attempts. A restart may be attempted after a health check failure or a container exit, based on the restart policy.
	ContainerRestartAttemptCount *int `mandatory:"false" json:"containerRestartAttemptCount"`
}

func (m Container) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Container) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingContainerLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetContainerLifecycleStateEnumStringValues(), ",")))
	}

	for _, val := range m.AdditionalCapabilities {
		if _, ok := GetMappingContainerCapabilityEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdditionalCapabilities: %s. Supported values are: %s.", val, strings.Join(GetContainerCapabilityEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Container) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FreeformTags                 map[string]string                 `json:"freeformTags"`
		DefinedTags                  map[string]map[string]interface{} `json:"definedTags"`
		SystemTags                   map[string]map[string]interface{} `json:"systemTags"`
		FaultDomain                  *string                           `json:"faultDomain"`
		LifecycleDetails             *string                           `json:"lifecycleDetails"`
		ExitCode                     *int                              `json:"exitCode"`
		TimeTerminated               *common.SDKTime                   `json:"timeTerminated"`
		TimeUpdated                  *common.SDKTime                   `json:"timeUpdated"`
		Command                      []string                          `json:"command"`
		Arguments                    []string                          `json:"arguments"`
		AdditionalCapabilities       []ContainerCapabilityEnum         `json:"additionalCapabilities"`
		WorkingDirectory             *string                           `json:"workingDirectory"`
		EnvironmentVariables         map[string]string                 `json:"environmentVariables"`
		VolumeMounts                 []VolumeMount                     `json:"volumeMounts"`
		HealthChecks                 []containerhealthcheck            `json:"healthChecks"`
		IsResourcePrincipalDisabled  *bool                             `json:"isResourcePrincipalDisabled"`
		ResourceConfig               *ContainerResourceConfig          `json:"resourceConfig"`
		ContainerRestartAttemptCount *int                              `json:"containerRestartAttemptCount"`
		Id                           *string                           `json:"id"`
		DisplayName                  *string                           `json:"displayName"`
		CompartmentId                *string                           `json:"compartmentId"`
		AvailabilityDomain           *string                           `json:"availabilityDomain"`
		LifecycleState               ContainerLifecycleStateEnum       `json:"lifecycleState"`
		TimeCreated                  *common.SDKTime                   `json:"timeCreated"`
		ContainerInstanceId          *string                           `json:"containerInstanceId"`
		ImageUrl                     *string                           `json:"imageUrl"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.FaultDomain = model.FaultDomain

	m.LifecycleDetails = model.LifecycleDetails

	m.ExitCode = model.ExitCode

	m.TimeTerminated = model.TimeTerminated

	m.TimeUpdated = model.TimeUpdated

	m.Command = make([]string, len(model.Command))
	for i, n := range model.Command {
		m.Command[i] = n
	}

	m.Arguments = make([]string, len(model.Arguments))
	for i, n := range model.Arguments {
		m.Arguments[i] = n
	}

	m.AdditionalCapabilities = make([]ContainerCapabilityEnum, len(model.AdditionalCapabilities))
	for i, n := range model.AdditionalCapabilities {
		m.AdditionalCapabilities[i] = n
	}

	m.WorkingDirectory = model.WorkingDirectory

	m.EnvironmentVariables = model.EnvironmentVariables

	m.VolumeMounts = make([]VolumeMount, len(model.VolumeMounts))
	for i, n := range model.VolumeMounts {
		m.VolumeMounts[i] = n
	}

	m.HealthChecks = make([]ContainerHealthCheck, len(model.HealthChecks))
	for i, n := range model.HealthChecks {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.HealthChecks[i] = nn.(ContainerHealthCheck)
		} else {
			m.HealthChecks[i] = nil
		}
	}

	m.IsResourcePrincipalDisabled = model.IsResourcePrincipalDisabled

	m.ResourceConfig = model.ResourceConfig

	m.ContainerRestartAttemptCount = model.ContainerRestartAttemptCount

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.AvailabilityDomain = model.AvailabilityDomain

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.ContainerInstanceId = model.ContainerInstanceId

	m.ImageUrl = model.ImageUrl

	return
}

// ContainerLifecycleStateEnum Enum with underlying type: string
type ContainerLifecycleStateEnum string

// Set of constants representing the allowable values for ContainerLifecycleStateEnum
const (
	ContainerLifecycleStateCreating ContainerLifecycleStateEnum = "CREATING"
	ContainerLifecycleStateUpdating ContainerLifecycleStateEnum = "UPDATING"
	ContainerLifecycleStateActive   ContainerLifecycleStateEnum = "ACTIVE"
	ContainerLifecycleStateInactive ContainerLifecycleStateEnum = "INACTIVE"
	ContainerLifecycleStateDeleting ContainerLifecycleStateEnum = "DELETING"
	ContainerLifecycleStateDeleted  ContainerLifecycleStateEnum = "DELETED"
	ContainerLifecycleStateFailed   ContainerLifecycleStateEnum = "FAILED"
)

var mappingContainerLifecycleStateEnum = map[string]ContainerLifecycleStateEnum{
	"CREATING": ContainerLifecycleStateCreating,
	"UPDATING": ContainerLifecycleStateUpdating,
	"ACTIVE":   ContainerLifecycleStateActive,
	"INACTIVE": ContainerLifecycleStateInactive,
	"DELETING": ContainerLifecycleStateDeleting,
	"DELETED":  ContainerLifecycleStateDeleted,
	"FAILED":   ContainerLifecycleStateFailed,
}

var mappingContainerLifecycleStateEnumLowerCase = map[string]ContainerLifecycleStateEnum{
	"creating": ContainerLifecycleStateCreating,
	"updating": ContainerLifecycleStateUpdating,
	"active":   ContainerLifecycleStateActive,
	"inactive": ContainerLifecycleStateInactive,
	"deleting": ContainerLifecycleStateDeleting,
	"deleted":  ContainerLifecycleStateDeleted,
	"failed":   ContainerLifecycleStateFailed,
}

// GetContainerLifecycleStateEnumValues Enumerates the set of values for ContainerLifecycleStateEnum
func GetContainerLifecycleStateEnumValues() []ContainerLifecycleStateEnum {
	values := make([]ContainerLifecycleStateEnum, 0)
	for _, v := range mappingContainerLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerLifecycleStateEnumStringValues Enumerates the set of values in String for ContainerLifecycleStateEnum
func GetContainerLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingContainerLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerLifecycleStateEnum(val string) (ContainerLifecycleStateEnum, bool) {
	enum, ok := mappingContainerLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
