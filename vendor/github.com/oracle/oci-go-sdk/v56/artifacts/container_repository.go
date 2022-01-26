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

// ContainerRepository Container repository metadata.
type ContainerRepository struct {

	// The OCID of the compartment in which the container repository exists.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The id of the user or principal that created the resource.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The container repository name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the container repository.
	// Example: `ocid1.containerrepo.oc1..exampleuniqueID`
	Id *string `mandatory:"true" json:"id"`

	// Total number of images.
	ImageCount *int `mandatory:"true" json:"imageCount"`

	// Whether the repository is immutable. Images cannot be overwritten in an immutable repository.
	IsImmutable *bool `mandatory:"true" json:"isImmutable"`

	// Whether the repository is public. A public repository allows unauthenticated access.
	IsPublic *bool `mandatory:"true" json:"isPublic"`

	// Total number of layers.
	LayerCount *int `mandatory:"true" json:"layerCount"`

	// Total storage in bytes consumed by layers.
	LayersSizeInBytes *int64 `mandatory:"true" json:"layersSizeInBytes"`

	// The current state of the container repository.
	LifecycleState ContainerRepositoryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// An RFC 3339 timestamp indicating when the repository was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Total storage size in GBs that will be charged.
	BillableSizeInGBs *int64 `mandatory:"true" json:"billableSizeInGBs"`

	Readme *ContainerRepositoryReadme `mandatory:"false" json:"readme"`

	// An RFC 3339 timestamp indicating when an image was last pushed to the repository.
	TimeLastPushed *common.SDKTime `mandatory:"false" json:"timeLastPushed"`
}

func (m ContainerRepository) String() string {
	return common.PointerString(m)
}

// ContainerRepositoryLifecycleStateEnum Enum with underlying type: string
type ContainerRepositoryLifecycleStateEnum string

// Set of constants representing the allowable values for ContainerRepositoryLifecycleStateEnum
const (
	ContainerRepositoryLifecycleStateAvailable ContainerRepositoryLifecycleStateEnum = "AVAILABLE"
	ContainerRepositoryLifecycleStateDeleting  ContainerRepositoryLifecycleStateEnum = "DELETING"
	ContainerRepositoryLifecycleStateDeleted   ContainerRepositoryLifecycleStateEnum = "DELETED"
)

var mappingContainerRepositoryLifecycleState = map[string]ContainerRepositoryLifecycleStateEnum{
	"AVAILABLE": ContainerRepositoryLifecycleStateAvailable,
	"DELETING":  ContainerRepositoryLifecycleStateDeleting,
	"DELETED":   ContainerRepositoryLifecycleStateDeleted,
}

// GetContainerRepositoryLifecycleStateEnumValues Enumerates the set of values for ContainerRepositoryLifecycleStateEnum
func GetContainerRepositoryLifecycleStateEnumValues() []ContainerRepositoryLifecycleStateEnum {
	values := make([]ContainerRepositoryLifecycleStateEnum, 0)
	for _, v := range mappingContainerRepositoryLifecycleState {
		values = append(values, v)
	}
	return values
}
