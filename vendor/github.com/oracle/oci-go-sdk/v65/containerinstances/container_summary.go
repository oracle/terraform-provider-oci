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

// ContainerSummary Summary information about a container.
type ContainerSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the container.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The compartment OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The availability domain where the container instance that hosts this container runs.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The current state of the container.
	LifecycleState ContainerLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the the container was created in the format defined by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the container instance on which the container is running.
	ContainerInstanceId *string `mandatory:"true" json:"containerInstanceId"`

	// A URL identifying the image that the container runs in, such as docker.io/library/busybox:latest. If you do not provide a tag, the tag will default to latest.
	// If no registry is provided, will default the registry to public docker hub `docker.io/library`.
	// The registry used for container image must be reachable over the Container Instance's VNIC.
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

	// The fault domain where the container instance that hosts the container runs.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// A message that describes the current state of the container in more detail. Can be used to provide
	// actionable information.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The time the container was updated in the format defined by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	ResourceConfig *ContainerResourceConfig `mandatory:"false" json:"resourceConfig"`

	// Determines whether the container will have access to the container instance resource principal.
	// This method utilizes resource principal version 2.2. For information on how to use the exposed resource principal elements, see
	// https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdk_authentication_methods.htm#sdk_authentication_methods_resource_principal.
	IsResourcePrincipalDisabled *bool `mandatory:"false" json:"isResourcePrincipalDisabled"`

	SecurityContext SecurityContext `mandatory:"false" json:"securityContext"`
}

func (m ContainerSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContainerSummary) ValidateEnumValue() (bool, error) {
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
func (m *ContainerSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FreeformTags                map[string]string                 `json:"freeformTags"`
		DefinedTags                 map[string]map[string]interface{} `json:"definedTags"`
		SystemTags                  map[string]map[string]interface{} `json:"systemTags"`
		FaultDomain                 *string                           `json:"faultDomain"`
		LifecycleDetails            *string                           `json:"lifecycleDetails"`
		TimeUpdated                 *common.SDKTime                   `json:"timeUpdated"`
		ResourceConfig              *ContainerResourceConfig          `json:"resourceConfig"`
		IsResourcePrincipalDisabled *bool                             `json:"isResourcePrincipalDisabled"`
		SecurityContext             securitycontext                   `json:"securityContext"`
		Id                          *string                           `json:"id"`
		DisplayName                 *string                           `json:"displayName"`
		CompartmentId               *string                           `json:"compartmentId"`
		AvailabilityDomain          *string                           `json:"availabilityDomain"`
		LifecycleState              ContainerLifecycleStateEnum       `json:"lifecycleState"`
		TimeCreated                 *common.SDKTime                   `json:"timeCreated"`
		ContainerInstanceId         *string                           `json:"containerInstanceId"`
		ImageUrl                    *string                           `json:"imageUrl"`
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

	m.TimeUpdated = model.TimeUpdated

	m.ResourceConfig = model.ResourceConfig

	m.IsResourcePrincipalDisabled = model.IsResourcePrincipalDisabled

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
