// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// ContainerInstanceSummary A reduced set of details about a single ContainerInstance returned by list APIs.
type ContainerInstanceSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Display name for the ContainerInstance. Can be renamed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Availability Domain where the ContainerInstance is running.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The current state of the ContainerInstance.
	LifecycleState ContainerInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the the ContainerInstance was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The shape of the Container Instance. The shape determines the resources available to the Container Instance.
	Shape *string `mandatory:"true" json:"shape"`

	ShapeConfig *ContainerInstanceShapeConfig `mandatory:"true" json:"shapeConfig"`

	// The number of containers on this Instance
	ContainerCount *int `mandatory:"true" json:"containerCount"`

	// Container Restart Policy
	ContainerRestartPolicy ContainerInstanceContainerRestartPolicyEnum `mandatory:"true" json:"containerRestartPolicy"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Fault Domain where the ContainerInstance is running.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// A message describing the current state in more detail. For example, can be used to provide
	// actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The time the ContainerInstance was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Duration in seconds processes within a Container have to gracefully terminate. This applies whenever a Container must be halted, such as when the Container Instance is deleted. Processes will first be sent a termination signal. After this timeout is reached, the processes will be sent a termination signal.
	GracefulShutdownTimeoutInSeconds *int64 `mandatory:"false" json:"gracefulShutdownTimeoutInSeconds"`

	// The number of volumes that attached to this Instance
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
