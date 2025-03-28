// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Artifacts and Container Images API
//
// API covering the Artifacts and Registry (https://docs.oracle.com/iaas/Content/Registry/Concepts/registryoverview.htm) services.
// Use this API to manage resources such as generic artifacts and container images.
//

package artifacts

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ContainerRepository Container repository metadata.
type ContainerRepository struct {

	// The OCID of the compartment in which the container repository exists.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The id of the user or principal that created the resource.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The container repository name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the container repository.
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

	// The tenancy namespace used in the container repository path.
	Namespace *string `mandatory:"true" json:"namespace"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The system tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	Readme *ContainerRepositoryReadme `mandatory:"false" json:"readme"`

	// An RFC 3339 timestamp indicating when an image was last pushed to the repository.
	TimeLastPushed *common.SDKTime `mandatory:"false" json:"timeLastPushed"`
}

func (m ContainerRepository) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContainerRepository) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingContainerRepositoryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetContainerRepositoryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ContainerRepositoryLifecycleStateEnum Enum with underlying type: string
type ContainerRepositoryLifecycleStateEnum string

// Set of constants representing the allowable values for ContainerRepositoryLifecycleStateEnum
const (
	ContainerRepositoryLifecycleStateAvailable ContainerRepositoryLifecycleStateEnum = "AVAILABLE"
	ContainerRepositoryLifecycleStateDeleting  ContainerRepositoryLifecycleStateEnum = "DELETING"
	ContainerRepositoryLifecycleStateDeleted   ContainerRepositoryLifecycleStateEnum = "DELETED"
)

var mappingContainerRepositoryLifecycleStateEnum = map[string]ContainerRepositoryLifecycleStateEnum{
	"AVAILABLE": ContainerRepositoryLifecycleStateAvailable,
	"DELETING":  ContainerRepositoryLifecycleStateDeleting,
	"DELETED":   ContainerRepositoryLifecycleStateDeleted,
}

var mappingContainerRepositoryLifecycleStateEnumLowerCase = map[string]ContainerRepositoryLifecycleStateEnum{
	"available": ContainerRepositoryLifecycleStateAvailable,
	"deleting":  ContainerRepositoryLifecycleStateDeleting,
	"deleted":   ContainerRepositoryLifecycleStateDeleted,
}

// GetContainerRepositoryLifecycleStateEnumValues Enumerates the set of values for ContainerRepositoryLifecycleStateEnum
func GetContainerRepositoryLifecycleStateEnumValues() []ContainerRepositoryLifecycleStateEnum {
	values := make([]ContainerRepositoryLifecycleStateEnum, 0)
	for _, v := range mappingContainerRepositoryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerRepositoryLifecycleStateEnumStringValues Enumerates the set of values in String for ContainerRepositoryLifecycleStateEnum
func GetContainerRepositoryLifecycleStateEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingContainerRepositoryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerRepositoryLifecycleStateEnum(val string) (ContainerRepositoryLifecycleStateEnum, bool) {
	enum, ok := mappingContainerRepositoryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
