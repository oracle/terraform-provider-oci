// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Images API
//
// API covering the Registry (https://docs.cloud.oracle.com/iaas/Content/Registry/Concepts/registryoverview.htm) services.
// Use this API to manage resources such as container images and repositories.
//

package artifacts

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ContainerImage Container image metadata.
type ContainerImage struct {

	// The compartment OCID to which the container image belongs. Inferred from the container repository.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the user or principal that created the resource.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The container image digest.
	Digest *string `mandatory:"true" json:"digest"`

	// The repository name and the most recent version associated with the image.
	// If there are no versions associated with the image, then last known version and digest are used instead.
	// If the last known version is unavailable, then 'unknown' is used instead of the version.
	// Example: `ubuntu:latest` or `ubuntu:latest@sha256:45b23dee08af5e43a7fea6c4cf9c25ccf269ee113168c19722f87876677c5cb2`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the container image.
	// Example: `ocid1.containerimage.oc1..exampleuniqueID`
	Id *string `mandatory:"true" json:"id"`

	// Layers of which the image is composed, ordered by the layer digest.
	Layers []ContainerImageLayer `mandatory:"true" json:"layers"`

	// The total size of the container image layers in bytes.
	LayersSizeInBytes *int64 `mandatory:"true" json:"layersSizeInBytes"`

	// The current state of the container image.
	LifecycleState ContainerImageLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The size of the container image manifest in bytes.
	ManifestSizeInBytes *int `mandatory:"true" json:"manifestSizeInBytes"`

	// Total number of pulls.
	PullCount *int64 `mandatory:"true" json:"pullCount"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the container repository.
	RepositoryId *string `mandatory:"true" json:"repositoryId"`

	// The container repository name.
	RepositoryName *string `mandatory:"true" json:"repositoryName"`

	// An RFC 3339 timestamp indicating when the image was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The versions associated with this image.
	Versions []ContainerVersion `mandatory:"true" json:"versions"`

	// An RFC 3339 timestamp indicating when the image was last pulled.
	TimeLastPulled *common.SDKTime `mandatory:"false" json:"timeLastPulled"`

	// The most recent version associated with this image.
	Version *string `mandatory:"false" json:"version"`
}

func (m ContainerImage) String() string {
	return common.PointerString(m)
}

// ContainerImageLifecycleStateEnum Enum with underlying type: string
type ContainerImageLifecycleStateEnum string

// Set of constants representing the allowable values for ContainerImageLifecycleStateEnum
const (
	ContainerImageLifecycleStateAvailable ContainerImageLifecycleStateEnum = "AVAILABLE"
	ContainerImageLifecycleStateDeleted   ContainerImageLifecycleStateEnum = "DELETED"
	ContainerImageLifecycleStateDeleting  ContainerImageLifecycleStateEnum = "DELETING"
)

var mappingContainerImageLifecycleState = map[string]ContainerImageLifecycleStateEnum{
	"AVAILABLE": ContainerImageLifecycleStateAvailable,
	"DELETED":   ContainerImageLifecycleStateDeleted,
	"DELETING":  ContainerImageLifecycleStateDeleting,
}

// GetContainerImageLifecycleStateEnumValues Enumerates the set of values for ContainerImageLifecycleStateEnum
func GetContainerImageLifecycleStateEnumValues() []ContainerImageLifecycleStateEnum {
	values := make([]ContainerImageLifecycleStateEnum, 0)
	for _, v := range mappingContainerImageLifecycleState {
		values = append(values, v)
	}
	return values
}
