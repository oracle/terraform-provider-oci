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

// CreateContainerDetails Information to create a new Container within a ContainerInstance.
// The Container created by this call will contain both the tags specified
// in this object as well as any tags specified in the parent ContainerInstance object.
// The Container will be created with the same `compartmentId`, `availabilityDomain`,
// and `faultDomain` as the parent ContainerInstance object.
type CreateContainerDetails struct {

	// The container image information. Currently only support public docker registry. Can be either image name,
	// e.g `containerImage`, image name with version, e.g `containerImage:v1` or complete docker image Url e.g
	// `docker.io/library/containerImage:latest`.
	// If no registry is provided, will default the registry to public docker hub `docker.io/library`.
	// The registry used for container image must be reachable over the Container Instance's VNIC.
	ImageUrl *string `mandatory:"true" json:"imageUrl"`

	// Display name for the Container. There are no guarantees of uniqueness
	// for this name. If none is provided, it will be calculated automatically.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// This command will override the container's entrypoint process.
	// If not specified, the existing entrypoint process defined in the image will be used.
	Command []string `mandatory:"false" json:"command"`

	// A list of string arguments for a container's entrypoint process.
	// Many containers use an entrypoint process pointing to a shell,
	// for example /bin/bash. For such containers, this argument list
	// can also be used to specify the main command in the container process.
	// All arguments together must be 64KB or smaller.
	Arguments []string `mandatory:"false" json:"arguments"`

	// A list of additional capabilities for the container.
	AdditionalCapabilities []ContainerCapabilityEnum `mandatory:"false" json:"additionalCapabilities,omitempty"`

	// The working directory within the Container's filesystem for
	// the Container process. If none is set, the Container will run in the
	// working directory set by the container image.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`

	// A map of additional environment variables to set in the environment of the container's
	// entrypoint process. These variables are in addition to any variables already defined
	// in the container's image.
	// All environment variables together, name and values, must be 64KB or smaller.
	EnvironmentVariables map[string]string `mandatory:"false" json:"environmentVariables"`

	// List of the volume mounts.
	VolumeMounts []CreateVolumeMountDetails `mandatory:"false" json:"volumeMounts"`

	// Determines if the Container will have access to the Container Instance Resource Principal.
	// This method utilizes resource principal version 2.2. Please refer to
	// https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdk_authentication_methods.htm#sdk_authentication_methods_resource_principal
	// for detailed explanation of how to leverage the exposed resource principal elements.
	IsResourcePrincipalDisabled *bool `mandatory:"false" json:"isResourcePrincipalDisabled"`

	ResourceConfig *CreateContainerResourceConfigDetails `mandatory:"false" json:"resourceConfig"`

	// list of container health checks to check container status and take appropriate action if container status is failed.
	// There are three types of health checks that we currently support HTTP, TCP, and Command.
	HealthChecks []CreateContainerHealthCheckDetails `mandatory:"false" json:"healthChecks"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateContainerDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateContainerDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

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
func (m *CreateContainerDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                 *string                               `json:"displayName"`
		Command                     []string                              `json:"command"`
		Arguments                   []string                              `json:"arguments"`
		AdditionalCapabilities      []ContainerCapabilityEnum             `json:"additionalCapabilities"`
		WorkingDirectory            *string                               `json:"workingDirectory"`
		EnvironmentVariables        map[string]string                     `json:"environmentVariables"`
		VolumeMounts                []CreateVolumeMountDetails            `json:"volumeMounts"`
		IsResourcePrincipalDisabled *bool                                 `json:"isResourcePrincipalDisabled"`
		ResourceConfig              *CreateContainerResourceConfigDetails `json:"resourceConfig"`
		HealthChecks                []createcontainerhealthcheckdetails   `json:"healthChecks"`
		FreeformTags                map[string]string                     `json:"freeformTags"`
		DefinedTags                 map[string]map[string]interface{}     `json:"definedTags"`
		ImageUrl                    *string                               `json:"imageUrl"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

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

	m.VolumeMounts = make([]CreateVolumeMountDetails, len(model.VolumeMounts))
	for i, n := range model.VolumeMounts {
		m.VolumeMounts[i] = n
	}

	m.IsResourcePrincipalDisabled = model.IsResourcePrincipalDisabled

	m.ResourceConfig = model.ResourceConfig

	m.HealthChecks = make([]CreateContainerHealthCheckDetails, len(model.HealthChecks))
	for i, n := range model.HealthChecks {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.HealthChecks[i] = nn.(CreateContainerHealthCheckDetails)
		} else {
			m.HealthChecks[i] = nil
		}
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.ImageUrl = model.ImageUrl

	return
}
