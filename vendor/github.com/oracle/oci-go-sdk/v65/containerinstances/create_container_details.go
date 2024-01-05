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

// CreateContainerDetails Information to create a new container within a container instance.
// The container created by this call contains both the tags specified
// in this object and any tags specified in the parent container instance.
// The container is created in the same compartment, availability domain,
// and fault domain as its container instance.
type CreateContainerDetails struct {

	// A URL identifying the image that the container runs in, such as docker.io/library/busybox:latest. If you do not provide a tag, the tag will default to latest.
	// If no registry is provided, will default the registry to public docker hub `docker.io/library`.
	// The registry used for container image must be reachable over the Container Instance's VNIC.
	ImageUrl *string `mandatory:"true" json:"imageUrl"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	// If you don't provide a name, a name is generated automatically.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// An optional command that overrides the ENTRYPOINT process.
	// If you do not provide a value, the existing ENTRYPOINT process defined in the image is used.
	Command []string `mandatory:"false" json:"command"`

	// A list of string arguments for a container's ENTRYPOINT process.
	// Many containers use an ENTRYPOINT process pointing to a shell
	// (/bin/bash). For those containers, this argument list
	// specifies the main command in the container process.
	// The total size of all arguments combined must be 64 KB or smaller.
	Arguments []string `mandatory:"false" json:"arguments"`

	// The working directory within the container's filesystem for
	// the container process. If not specified, the default
	// working directory from the image is used.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`

	// A map of additional environment variables to set in the environment of the container's
	// ENTRYPOINT process. These variables are in addition to any variables already defined
	// in the container's image.
	// The total size of all environment variables combined, name and values, must be 64 KB or smaller.
	EnvironmentVariables map[string]string `mandatory:"false" json:"environmentVariables"`

	// List of the volume mounts.
	VolumeMounts []CreateVolumeMountDetails `mandatory:"false" json:"volumeMounts"`

	// Determines if the container will have access to the container instance resource principal.
	// This method utilizes resource principal version 2.2. For information on how to use the exposed resource principal elements, see
	// https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdk_authentication_methods.htm#sdk_authentication_methods_resource_principal.
	IsResourcePrincipalDisabled *bool `mandatory:"false" json:"isResourcePrincipalDisabled"`

	ResourceConfig *CreateContainerResourceConfigDetails `mandatory:"false" json:"resourceConfig"`

	// list of container health checks to check container status and take appropriate action if container status is failed.
	// There are three types of health checks that we currently support HTTP, TCP, and Command.
	HealthChecks []CreateContainerHealthCheckDetails `mandatory:"false" json:"healthChecks"`

	SecurityContext CreateSecurityContextDetails `mandatory:"false" json:"securityContext"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`.
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
		WorkingDirectory            *string                               `json:"workingDirectory"`
		EnvironmentVariables        map[string]string                     `json:"environmentVariables"`
		VolumeMounts                []CreateVolumeMountDetails            `json:"volumeMounts"`
		IsResourcePrincipalDisabled *bool                                 `json:"isResourcePrincipalDisabled"`
		ResourceConfig              *CreateContainerResourceConfigDetails `json:"resourceConfig"`
		HealthChecks                []createcontainerhealthcheckdetails   `json:"healthChecks"`
		SecurityContext             createsecuritycontextdetails          `json:"securityContext"`
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
	copy(m.Command, model.Command)
	m.Arguments = make([]string, len(model.Arguments))
	copy(m.Arguments, model.Arguments)
	m.WorkingDirectory = model.WorkingDirectory

	m.EnvironmentVariables = model.EnvironmentVariables

	m.VolumeMounts = make([]CreateVolumeMountDetails, len(model.VolumeMounts))
	copy(m.VolumeMounts, model.VolumeMounts)
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
	nn, e = model.SecurityContext.UnmarshalPolymorphicJSON(model.SecurityContext.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SecurityContext = nn.(CreateSecurityContextDetails)
	} else {
		m.SecurityContext = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.ImageUrl = model.ImageUrl

	return
}
