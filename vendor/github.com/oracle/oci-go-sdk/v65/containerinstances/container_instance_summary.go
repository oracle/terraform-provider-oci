// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ContainerInstanceSummary A set of details about a single container instance returned by list APIs.
type ContainerInstanceSummary struct {

	// OCID that cannot be changed.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment to create the container instance in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The availability domain where the container instance runs.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The current state of the container instance.
	LifecycleState ContainerInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the container instance was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The shape of the container instance. The shape determines the resources available to the container instance.
	Shape *string `mandatory:"true" json:"shape"`

	ShapeConfig *ContainerInstanceShapeConfig `mandatory:"true" json:"shapeConfig"`

	// The number of containers in the container instance.
	ContainerCount *int `mandatory:"true" json:"containerCount"`

	// Container Restart Policy
	ContainerRestartPolicy ContainerInstanceContainerRestartPolicyEnum `mandatory:"true" json:"containerRestartPolicy"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`.
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`.
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The fault domain where the container instance runs.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// A message that describes the current state of the container instance in more detail. Can be used to provide
	// actionable information.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The time the container instance was updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The amount of time that processes in a container have to gracefully end when the container must be stopped. For example, when you delete a container instance. After the timeout is reached, the processes are sent a signal to be deleted.
	GracefulShutdownTimeoutInSeconds *int64 `mandatory:"false" json:"gracefulShutdownTimeoutInSeconds"`

	// The number of volumes that are attached to the container instance.
	VolumeCount *int `mandatory:"false" json:"volumeCount"`
}

func (m ContainerInstanceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContainerInstanceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingContainerInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetContainerInstanceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingContainerInstanceContainerRestartPolicyEnum(string(m.ContainerRestartPolicy)); !ok && m.ContainerRestartPolicy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ContainerRestartPolicy: %s. Supported values are: %s.", m.ContainerRestartPolicy, strings.Join(GetContainerInstanceContainerRestartPolicyEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
