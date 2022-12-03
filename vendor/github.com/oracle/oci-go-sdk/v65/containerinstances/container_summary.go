// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ContainerSummary A reduced set of details about a single Container returned by list APIs.
type ContainerSummary struct {

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

	// The time the Container was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	ResourceConfig *ContainerResourceConfig `mandatory:"false" json:"resourceConfig"`

	// Determines if the Container will have access to the Container Instance Resource Principal.
	// This method utilizes resource principal version 2.2. Please refer to
	// https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdk_authentication_methods.htm#sdk_authentication_methods_resource_principal
	// for detailed explanation of how to leverage the exposed resource principal elements.
	IsResourcePrincipalDisabled *bool `mandatory:"false" json:"isResourcePrincipalDisabled"`
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
