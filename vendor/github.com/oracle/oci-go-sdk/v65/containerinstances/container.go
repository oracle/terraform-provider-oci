// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// Container A single container on a container instance.
// If you delete a container, the record remains visible for a short period
// of time before being permanently removed.
type Container struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the container.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment that contains the container.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The availability domain where the container instance that hosts the container runs.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The current state of the container.
	LifecycleState ContainerLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the container was created, in the format defined by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the container instance that the container is running on.
	ContainerInstanceId *string `mandatory:"true" json:"containerInstanceId"`

	// The container image information. Currently only supports public Docker registry.
	// You can provide either the image name (containerImage), image name with version (containerImagev1), or complete Docker image URL
	// `docker.io/library/containerImage:latest`.
	// If you do not provide a registry, the registry defaults to public Docker hub `docker.io/library`.
	// The registry used for the container image must be reachable over the VNIC of the container instance.
	ImageUrl *string `mandatory:"true" json:"imageUrl"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`.
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`.
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The fault domain of the container instance that hosts the container runs.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// A message that describes the current state of the container in more detail. Can be used to provide
	// actionable information.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The exit code of the container process when it stopped running.
	ExitCode *int `mandatory:"false" json:"exitCode"`

	// The time when the container last deleted (terminated), in the format defined by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeTerminated *common.SDKTime `mandatory:"false" json:"timeTerminated"`

	// The time the container was updated, in the format defined by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// This command overrides ENTRYPOINT process of the container.
	// If you do not specify this command, the existing ENTRYPOINT process defined in the image is the default.
	Command []string `mandatory:"false" json:"command"`

	// A list of string arguments for the ENTRYPOINT process of the container.
	// Many containers use an ENTRYPOINT process pointing to a shell
	// `/bin/bash`. For those containers, you can use the argument list to specify the main command in the container process.
	Arguments []string `mandatory:"false" json:"arguments"`

	// The working directory within the container's filesystem for
	// the container process. If not specified, the default
	// working directory from the image is used.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`

	// A map of additional environment variables to set in the environment of the
	// ENTRYPOINT process of the container. These variables are in addition to any variables already defined
	// in the container's image.
	EnvironmentVariables map[string]string `mandatory:"false" json:"environmentVariables"`

	// List of the volume mounts.
	VolumeMounts []VolumeMount `mandatory:"false" json:"volumeMounts"`

	// List of container health checks
	HealthChecks []ContainerHealthCheck `mandatory:"false" json:"healthChecks"`

	// Determines if the container will have access to the container instance resource principal.
	// This method utilizes resource principal version 2.2. For more information on how to use the exposed resource principal elements, see
	// https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdk_authentication_methods.htm#sdk_authentication_methods_resource_principal.
	IsResourcePrincipalDisabled *bool `mandatory:"false" json:"isResourcePrincipalDisabled"`

	ResourceConfig *ContainerResourceConfig `mandatory:"false" json:"resourceConfig"`

	// The number of container restart attempts. Depending on the restart policy, a restart might be attempted after a health check failure or a container exit.
	ContainerRestartAttemptCount *int `mandatory:"false" json:"containerRestartAttemptCount"`

	SecurityContext SecurityContext `mandatory:"false" json:"securityContext"`
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
		WorkingDirectory             *string                           `json:"workingDirectory"`
		EnvironmentVariables         map[string]string                 `json:"environmentVariables"`
		VolumeMounts                 []VolumeMount                     `json:"volumeMounts"`
		HealthChecks                 []containerhealthcheck            `json:"healthChecks"`
		IsResourcePrincipalDisabled  *bool                             `json:"isResourcePrincipalDisabled"`
		ResourceConfig               *ContainerResourceConfig          `json:"resourceConfig"`
		ContainerRestartAttemptCount *int                              `json:"containerRestartAttemptCount"`
		SecurityContext              securitycontext                   `json:"securityContext"`
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
	copy(m.Command, model.Command)
	m.Arguments = make([]string, len(model.Arguments))
	copy(m.Arguments, model.Arguments)
	m.WorkingDirectory = model.WorkingDirectory

	m.EnvironmentVariables = model.EnvironmentVariables

	m.VolumeMounts = make([]VolumeMount, len(model.VolumeMounts))
	copy(m.VolumeMounts, model.VolumeMounts)
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

	nn, e = model.SecurityContext.UnmarshalPolymorphicJSON(model.SecurityContext.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SecurityContext = nn.(SecurityContext)
	} else {
		m.SecurityContext = nil
	}

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
