// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Artifacts and Container Images API
//
// Use the Artifacts and Container Images API to manage container images and non-container generic artifacts.
// - For container images such as Docker images, use the ContainerImage resource. Save the images in a ContainerRepository.
// - For non-container generic artifacts or blobs, use the GenericArtifact resource. Save the artifacts in an Repository.
// - To upload and download non-container generic artifacts, instead of the Artifacts and Container Images API, use the Generic Artifacts Content API.
// For more information, see the user guides for Container Registry (https://docs.oracle.com/iaas/Content/Registry/home.htm) and Artifact Registry (https://docs.oracle.com/iaas/Content/artifacts/home.htm).
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

	// The type of container repository. This determines the repository behavior and required settings.
	RepositoryType ContainerRepositoryRepositoryTypeEnum `mandatory:"true" json:"repositoryType"`

	// Whether the repository is top-level. A top-level repository reserves its name in namespace as a prefix and allows images to be hosted under child paths of that prefix.
	IsTopLevel *bool `mandatory:"true" json:"isTopLevel"`

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

	// The upstream registry URL for a pull-through cache repository.
	UpstreamUrl *string `mandatory:"false" json:"upstreamUrl"`

	// The upstream registry username for a pull-through cache repository.
	UpstreamUsername *string `mandatory:"false" json:"upstreamUsername"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OCI Vault secret that contains the upstream registry password as base64 encoded string for a pull-through cache repository.
	UpstreamSecretId *string `mandatory:"false" json:"upstreamSecretId"`

	// List of peer regions that define the multi-region group. Peer regions must be specified using canonical region names (for example, us-phoenix-1).
	// The peer set must match a supported multi-region endpoint group and MUST include the local OCIR region where the repository is created.
	MultiRegionPeers []string `mandatory:"false" json:"multiRegionPeers"`
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
	if _, ok := GetMappingContainerRepositoryRepositoryTypeEnum(string(m.RepositoryType)); !ok && m.RepositoryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RepositoryType: %s. Supported values are: %s.", m.RepositoryType, strings.Join(GetContainerRepositoryRepositoryTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
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

// ContainerRepositoryRepositoryTypeEnum Enum with underlying type: string
type ContainerRepositoryRepositoryTypeEnum string

// Set of constants representing the allowable values for ContainerRepositoryRepositoryTypeEnum
const (
	ContainerRepositoryRepositoryTypeStandard    ContainerRepositoryRepositoryTypeEnum = "STANDARD"
	ContainerRepositoryRepositoryTypePullThrough ContainerRepositoryRepositoryTypeEnum = "PULL_THROUGH"
	ContainerRepositoryRepositoryTypeMultiRegion ContainerRepositoryRepositoryTypeEnum = "MULTI_REGION"
)

var mappingContainerRepositoryRepositoryTypeEnum = map[string]ContainerRepositoryRepositoryTypeEnum{
	"STANDARD":     ContainerRepositoryRepositoryTypeStandard,
	"PULL_THROUGH": ContainerRepositoryRepositoryTypePullThrough,
	"MULTI_REGION": ContainerRepositoryRepositoryTypeMultiRegion,
}

var mappingContainerRepositoryRepositoryTypeEnumLowerCase = map[string]ContainerRepositoryRepositoryTypeEnum{
	"standard":     ContainerRepositoryRepositoryTypeStandard,
	"pull_through": ContainerRepositoryRepositoryTypePullThrough,
	"multi_region": ContainerRepositoryRepositoryTypeMultiRegion,
}

// GetContainerRepositoryRepositoryTypeEnumValues Enumerates the set of values for ContainerRepositoryRepositoryTypeEnum
func GetContainerRepositoryRepositoryTypeEnumValues() []ContainerRepositoryRepositoryTypeEnum {
	values := make([]ContainerRepositoryRepositoryTypeEnum, 0)
	for _, v := range mappingContainerRepositoryRepositoryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerRepositoryRepositoryTypeEnumStringValues Enumerates the set of values in String for ContainerRepositoryRepositoryTypeEnum
func GetContainerRepositoryRepositoryTypeEnumStringValues() []string {
	return []string{
		"STANDARD",
		"PULL_THROUGH",
		"MULTI_REGION",
	}
}

// GetMappingContainerRepositoryRepositoryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerRepositoryRepositoryTypeEnum(val string) (ContainerRepositoryRepositoryTypeEnum, bool) {
	enum, ok := mappingContainerRepositoryRepositoryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
