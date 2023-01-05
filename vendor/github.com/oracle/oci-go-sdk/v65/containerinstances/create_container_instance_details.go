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

// CreateContainerInstanceDetails The information about new ContainerInstance.
type CreateContainerInstanceDetails struct {

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Availability Domain where the ContainerInstance should be created.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The shape of the Container Instance. The shape determines the resources available to the Container Instance.
	Shape *string `mandatory:"true" json:"shape"`

	ShapeConfig *CreateContainerInstanceShapeConfigDetails `mandatory:"true" json:"shapeConfig"`

	// The Containers to create on this Instance.
	Containers []CreateContainerDetails `mandatory:"true" json:"containers"`

	// The networks to make available to containers on this Instance.
	Vnics []CreateContainerVnicDetails `mandatory:"true" json:"vnics"`

	// Human-readable name for the ContainerInstance. If none is provided,
	// OCI will select one for you.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Fault Domain where the ContainerInstance should run.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// A Volume represents a directory with data that is accessible across multiple containers in a
	// ContainerInstance.
	// Up to 32 volumes can be attached to single container instance.
	Volumes []CreateContainerVolumeDetails `mandatory:"false" json:"volumes"`

	DnsConfig *CreateContainerDnsConfigDetails `mandatory:"false" json:"dnsConfig"`

	// Duration in seconds processes within a Container have to gracefully terminate. This applies whenever a Container must be halted, such as when the Container Instance is deleted. Processes will first be sent a termination signal. After this timeout is reached, the processes will be sent a termination signal.
	GracefulShutdownTimeoutInSeconds *int64 `mandatory:"false" json:"gracefulShutdownTimeoutInSeconds"`

	// The image pull secrets for accessing private registry to pull images for containers
	ImagePullSecrets []CreateImagePullSecretDetails `mandatory:"false" json:"imagePullSecrets"`

	// Container restart policy
	ContainerRestartPolicy ContainerInstanceContainerRestartPolicyEnum `mandatory:"false" json:"containerRestartPolicy,omitempty"`

	// Customer's streaming OCID which is used for receiving a message whenever container health check status changes.
	StreamId *string `mandatory:"false" json:"streamId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateContainerInstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateContainerInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingContainerInstanceContainerRestartPolicyEnum(string(m.ContainerRestartPolicy)); !ok && m.ContainerRestartPolicy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ContainerRestartPolicy: %s. Supported values are: %s.", m.ContainerRestartPolicy, strings.Join(GetContainerInstanceContainerRestartPolicyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateContainerInstanceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                      *string                                     `json:"displayName"`
		FaultDomain                      *string                                     `json:"faultDomain"`
		Volumes                          []createcontainervolumedetails              `json:"volumes"`
		DnsConfig                        *CreateContainerDnsConfigDetails            `json:"dnsConfig"`
		GracefulShutdownTimeoutInSeconds *int64                                      `json:"gracefulShutdownTimeoutInSeconds"`
		ImagePullSecrets                 []createimagepullsecretdetails              `json:"imagePullSecrets"`
		ContainerRestartPolicy           ContainerInstanceContainerRestartPolicyEnum `json:"containerRestartPolicy"`
		StreamId                         *string                                     `json:"streamId"`
		FreeformTags                     map[string]string                           `json:"freeformTags"`
		DefinedTags                      map[string]map[string]interface{}           `json:"definedTags"`
		CompartmentId                    *string                                     `json:"compartmentId"`
		AvailabilityDomain               *string                                     `json:"availabilityDomain"`
		Shape                            *string                                     `json:"shape"`
		ShapeConfig                      *CreateContainerInstanceShapeConfigDetails  `json:"shapeConfig"`
		Containers                       []CreateContainerDetails                    `json:"containers"`
		Vnics                            []CreateContainerVnicDetails                `json:"vnics"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.FaultDomain = model.FaultDomain

	m.Volumes = make([]CreateContainerVolumeDetails, len(model.Volumes))
	for i, n := range model.Volumes {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Volumes[i] = nn.(CreateContainerVolumeDetails)
		} else {
			m.Volumes[i] = nil
		}
	}

	m.DnsConfig = model.DnsConfig

	m.GracefulShutdownTimeoutInSeconds = model.GracefulShutdownTimeoutInSeconds

	m.ImagePullSecrets = make([]CreateImagePullSecretDetails, len(model.ImagePullSecrets))
	for i, n := range model.ImagePullSecrets {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ImagePullSecrets[i] = nn.(CreateImagePullSecretDetails)
		} else {
			m.ImagePullSecrets[i] = nil
		}
	}

	m.ContainerRestartPolicy = model.ContainerRestartPolicy

	m.StreamId = model.StreamId

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	m.AvailabilityDomain = model.AvailabilityDomain

	m.Shape = model.Shape

	m.ShapeConfig = model.ShapeConfig

	m.Containers = make([]CreateContainerDetails, len(model.Containers))
	for i, n := range model.Containers {
		m.Containers[i] = n
	}

	m.Vnics = make([]CreateContainerVnicDetails, len(model.Vnics))
	for i, n := range model.Vnics {
		m.Vnics[i] = n
	}

	return
}
