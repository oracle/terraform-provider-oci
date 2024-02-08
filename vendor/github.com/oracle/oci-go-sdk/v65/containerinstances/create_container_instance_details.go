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

// CreateContainerInstanceDetails Information to create a container instance.
type CreateContainerInstanceDetails struct {

	// The compartment OCID.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The availability domain where the container instance runs.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The shape of the container instance. The shape determines the resources available to the container instance.
	Shape *string `mandatory:"true" json:"shape"`

	ShapeConfig *CreateContainerInstanceShapeConfigDetails `mandatory:"true" json:"shapeConfig"`

	// The containers to create on this container instance.
	Containers []CreateContainerDetails `mandatory:"true" json:"containers"`

	// The networks available to containers on this container instance.
	Vnics []CreateContainerVnicDetails `mandatory:"true" json:"vnics"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. If you don't provide a name, a name is generated automatically.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The fault domain where the container instance runs.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// A volume is a directory with data that is accessible across multiple containers in a
	// container instance.
	// You can attach up to 32 volumes to single container instance.
	Volumes []CreateContainerVolumeDetails `mandatory:"false" json:"volumes"`

	DnsConfig *CreateContainerDnsConfigDetails `mandatory:"false" json:"dnsConfig"`

	// The amount of time that processes in a container have to gracefully end when the container must be stopped. For example, when you delete a container instance. After the timeout is reached, the processes are sent a signal to be deleted.
	GracefulShutdownTimeoutInSeconds *int64 `mandatory:"false" json:"gracefulShutdownTimeoutInSeconds"`

	// The image pulls secrets so you can access private registry to pull container images.
	ImagePullSecrets []CreateImagePullSecretDetails `mandatory:"false" json:"imagePullSecrets"`

	// Container restart policy
	ContainerRestartPolicy ContainerInstanceContainerRestartPolicyEnum `mandatory:"false" json:"containerRestartPolicy,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`.
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

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	m.AvailabilityDomain = model.AvailabilityDomain

	m.Shape = model.Shape

	m.ShapeConfig = model.ShapeConfig

	m.Containers = make([]CreateContainerDetails, len(model.Containers))
	copy(m.Containers, model.Containers)
	m.Vnics = make([]CreateContainerVnicDetails, len(model.Vnics))
	copy(m.Vnics, model.Vnics)
	return
}
