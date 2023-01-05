// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// ContainerInstance A ContainerInstance for hosting Containers.
// If this ContainerInstance is DELETED, the record will remain visible for a short period
// of time before being permanently removed.
type ContainerInstance struct {

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

	// The Containers on this Instance
	Containers []ContainerInstanceContainer `mandatory:"true" json:"containers"`

	// The number of containers on this Instance
	ContainerCount *int `mandatory:"true" json:"containerCount"`

	// The time the the ContainerInstance was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The shape of the Container Instance. The shape determines the resources available to the Container Instance.
	Shape *string `mandatory:"true" json:"shape"`

	ShapeConfig *ContainerInstanceShapeConfig `mandatory:"true" json:"shapeConfig"`

	// The virtual networks available to containers running on this Container Instance.
	Vnics []ContainerVnic `mandatory:"true" json:"vnics"`

	// The container restart policy is applied for all containers in container instance.
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

	// A Volume represents a directory with data that is accessible across multiple containers in a
	// ContainerInstance.
	Volumes []ContainerVolume `mandatory:"false" json:"volumes"`

	// The number of volumes that attached to this Instance
	VolumeCount *int `mandatory:"false" json:"volumeCount"`

	// The time the ContainerInstance was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	DnsConfig *ContainerDnsConfig `mandatory:"false" json:"dnsConfig"`

	// Duration in seconds processes within a Container have to gracefully terminate. This applies whenever a Container must be halted, such as when the Container Instance is deleted. Processes will first be sent a termination signal. After this timeout is reached, the processes will be sent a termination signal.
	GracefulShutdownTimeoutInSeconds *int64 `mandatory:"false" json:"gracefulShutdownTimeoutInSeconds"`

	// The image pull secrets for accessing private registry to pull images for containers
	ImagePullSecrets []ImagePullSecret `mandatory:"false" json:"imagePullSecrets"`

	// Customer's streaming OCID which is used for receiving a message whenever container health check status changes.
	StreamId *string `mandatory:"false" json:"streamId"`
}

func (m ContainerInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContainerInstance) ValidateEnumValue() (bool, error) {
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

// UnmarshalJSON unmarshals from json
func (m *ContainerInstance) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FreeformTags                     map[string]string                           `json:"freeformTags"`
		DefinedTags                      map[string]map[string]interface{}           `json:"definedTags"`
		SystemTags                       map[string]map[string]interface{}           `json:"systemTags"`
		FaultDomain                      *string                                     `json:"faultDomain"`
		LifecycleDetails                 *string                                     `json:"lifecycleDetails"`
		Volumes                          []containervolume                           `json:"volumes"`
		VolumeCount                      *int                                        `json:"volumeCount"`
		TimeUpdated                      *common.SDKTime                             `json:"timeUpdated"`
		DnsConfig                        *ContainerDnsConfig                         `json:"dnsConfig"`
		GracefulShutdownTimeoutInSeconds *int64                                      `json:"gracefulShutdownTimeoutInSeconds"`
		ImagePullSecrets                 []imagepullsecret                           `json:"imagePullSecrets"`
		StreamId                         *string                                     `json:"streamId"`
		Id                               *string                                     `json:"id"`
		DisplayName                      *string                                     `json:"displayName"`
		CompartmentId                    *string                                     `json:"compartmentId"`
		AvailabilityDomain               *string                                     `json:"availabilityDomain"`
		LifecycleState                   ContainerInstanceLifecycleStateEnum         `json:"lifecycleState"`
		Containers                       []ContainerInstanceContainer                `json:"containers"`
		ContainerCount                   *int                                        `json:"containerCount"`
		TimeCreated                      *common.SDKTime                             `json:"timeCreated"`
		Shape                            *string                                     `json:"shape"`
		ShapeConfig                      *ContainerInstanceShapeConfig               `json:"shapeConfig"`
		Vnics                            []ContainerVnic                             `json:"vnics"`
		ContainerRestartPolicy           ContainerInstanceContainerRestartPolicyEnum `json:"containerRestartPolicy"`
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

	m.Volumes = make([]ContainerVolume, len(model.Volumes))
	for i, n := range model.Volumes {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Volumes[i] = nn.(ContainerVolume)
		} else {
			m.Volumes[i] = nil
		}
	}

	m.VolumeCount = model.VolumeCount

	m.TimeUpdated = model.TimeUpdated

	m.DnsConfig = model.DnsConfig

	m.GracefulShutdownTimeoutInSeconds = model.GracefulShutdownTimeoutInSeconds

	m.ImagePullSecrets = make([]ImagePullSecret, len(model.ImagePullSecrets))
	for i, n := range model.ImagePullSecrets {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ImagePullSecrets[i] = nn.(ImagePullSecret)
		} else {
			m.ImagePullSecrets[i] = nil
		}
	}

	m.StreamId = model.StreamId

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.AvailabilityDomain = model.AvailabilityDomain

	m.LifecycleState = model.LifecycleState

	m.Containers = make([]ContainerInstanceContainer, len(model.Containers))
	for i, n := range model.Containers {
		m.Containers[i] = n
	}

	m.ContainerCount = model.ContainerCount

	m.TimeCreated = model.TimeCreated

	m.Shape = model.Shape

	m.ShapeConfig = model.ShapeConfig

	m.Vnics = make([]ContainerVnic, len(model.Vnics))
	for i, n := range model.Vnics {
		m.Vnics[i] = n
	}

	m.ContainerRestartPolicy = model.ContainerRestartPolicy

	return
}

// ContainerInstanceLifecycleStateEnum Enum with underlying type: string
type ContainerInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for ContainerInstanceLifecycleStateEnum
const (
	ContainerInstanceLifecycleStateCreating ContainerInstanceLifecycleStateEnum = "CREATING"
	ContainerInstanceLifecycleStateUpdating ContainerInstanceLifecycleStateEnum = "UPDATING"
	ContainerInstanceLifecycleStateActive   ContainerInstanceLifecycleStateEnum = "ACTIVE"
	ContainerInstanceLifecycleStateInactive ContainerInstanceLifecycleStateEnum = "INACTIVE"
	ContainerInstanceLifecycleStateDeleting ContainerInstanceLifecycleStateEnum = "DELETING"
	ContainerInstanceLifecycleStateDeleted  ContainerInstanceLifecycleStateEnum = "DELETED"
	ContainerInstanceLifecycleStateFailed   ContainerInstanceLifecycleStateEnum = "FAILED"
)

var mappingContainerInstanceLifecycleStateEnum = map[string]ContainerInstanceLifecycleStateEnum{
	"CREATING": ContainerInstanceLifecycleStateCreating,
	"UPDATING": ContainerInstanceLifecycleStateUpdating,
	"ACTIVE":   ContainerInstanceLifecycleStateActive,
	"INACTIVE": ContainerInstanceLifecycleStateInactive,
	"DELETING": ContainerInstanceLifecycleStateDeleting,
	"DELETED":  ContainerInstanceLifecycleStateDeleted,
	"FAILED":   ContainerInstanceLifecycleStateFailed,
}

var mappingContainerInstanceLifecycleStateEnumLowerCase = map[string]ContainerInstanceLifecycleStateEnum{
	"creating": ContainerInstanceLifecycleStateCreating,
	"updating": ContainerInstanceLifecycleStateUpdating,
	"active":   ContainerInstanceLifecycleStateActive,
	"inactive": ContainerInstanceLifecycleStateInactive,
	"deleting": ContainerInstanceLifecycleStateDeleting,
	"deleted":  ContainerInstanceLifecycleStateDeleted,
	"failed":   ContainerInstanceLifecycleStateFailed,
}

// GetContainerInstanceLifecycleStateEnumValues Enumerates the set of values for ContainerInstanceLifecycleStateEnum
func GetContainerInstanceLifecycleStateEnumValues() []ContainerInstanceLifecycleStateEnum {
	values := make([]ContainerInstanceLifecycleStateEnum, 0)
	for _, v := range mappingContainerInstanceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerInstanceLifecycleStateEnumStringValues Enumerates the set of values in String for ContainerInstanceLifecycleStateEnum
func GetContainerInstanceLifecycleStateEnumStringValues() []string {
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

// GetMappingContainerInstanceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerInstanceLifecycleStateEnum(val string) (ContainerInstanceLifecycleStateEnum, bool) {
	enum, ok := mappingContainerInstanceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ContainerInstanceContainerRestartPolicyEnum Enum with underlying type: string
type ContainerInstanceContainerRestartPolicyEnum string

// Set of constants representing the allowable values for ContainerInstanceContainerRestartPolicyEnum
const (
	ContainerInstanceContainerRestartPolicyAlways    ContainerInstanceContainerRestartPolicyEnum = "ALWAYS"
	ContainerInstanceContainerRestartPolicyNever     ContainerInstanceContainerRestartPolicyEnum = "NEVER"
	ContainerInstanceContainerRestartPolicyOnFailure ContainerInstanceContainerRestartPolicyEnum = "ON_FAILURE"
)

var mappingContainerInstanceContainerRestartPolicyEnum = map[string]ContainerInstanceContainerRestartPolicyEnum{
	"ALWAYS":     ContainerInstanceContainerRestartPolicyAlways,
	"NEVER":      ContainerInstanceContainerRestartPolicyNever,
	"ON_FAILURE": ContainerInstanceContainerRestartPolicyOnFailure,
}

var mappingContainerInstanceContainerRestartPolicyEnumLowerCase = map[string]ContainerInstanceContainerRestartPolicyEnum{
	"always":     ContainerInstanceContainerRestartPolicyAlways,
	"never":      ContainerInstanceContainerRestartPolicyNever,
	"on_failure": ContainerInstanceContainerRestartPolicyOnFailure,
}

// GetContainerInstanceContainerRestartPolicyEnumValues Enumerates the set of values for ContainerInstanceContainerRestartPolicyEnum
func GetContainerInstanceContainerRestartPolicyEnumValues() []ContainerInstanceContainerRestartPolicyEnum {
	values := make([]ContainerInstanceContainerRestartPolicyEnum, 0)
	for _, v := range mappingContainerInstanceContainerRestartPolicyEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerInstanceContainerRestartPolicyEnumStringValues Enumerates the set of values in String for ContainerInstanceContainerRestartPolicyEnum
func GetContainerInstanceContainerRestartPolicyEnumStringValues() []string {
	return []string{
		"ALWAYS",
		"NEVER",
		"ON_FAILURE",
	}
}

// GetMappingContainerInstanceContainerRestartPolicyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerInstanceContainerRestartPolicyEnum(val string) (ContainerInstanceContainerRestartPolicyEnum, bool) {
	enum, ok := mappingContainerInstanceContainerRestartPolicyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
