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

// CreateContainerRepositoryDetails Create container repository details.
type CreateContainerRepositoryDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to create the resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The container repository name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of container repository. This determines the repository behavior and required settings.
	RepositoryType ContainerRepositoryRepositoryTypeEnum `mandatory:"false" json:"repositoryType,omitempty"`

	// Whether the repository is top-level. A top-level repository reserves its name in namespace as a prefix and allows images to be hosted under child paths of that prefix.
	IsTopLevel *bool `mandatory:"false" json:"isTopLevel"`

	// The upstream registry URL for a pull-through cache repository. Required when repositoryType is PULL_THROUGH.
	// The value must include the registry endpoint and repository path, and must not include a tag or digest.
	// If isTopLevel is false, the value must be an exact upstream image full path.
	// Example: `iad.ocir.io/<upstreamNamespace>/library/ubuntu`
	// If isTopLevel is true, the value must be an upstream repository path prefix, allowing OCIR to pull multiple repositories under that prefix.
	// Example: `iad.ocir.io/<upstreamNamespace>/library` or `iad.ocir.io/<upstreamNamespace>`
	UpstreamUrl *string `mandatory:"false" json:"upstreamUrl"`

	// The upstream registry username for a pull-through cache repository. Required when repositoryType is PULL_THROUGH.
	UpstreamUsername *string `mandatory:"false" json:"upstreamUsername"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OCI Vault secret that contains the upstream registry password as base64 encoded string for a pull-through cache repository. Required when repositoryType is PULL_THROUGH.
	// Example: `ocid1.vaultsecret.oc1..exampleuniqueID`
	UpstreamSecretId *string `mandatory:"false" json:"upstreamSecretId"`

	// List of peer regions that define the multi-region group. Peer regions must be specified using canonical region names (for example, `us-phoenix-1`).
	// The peer set must match a supported multi-region endpoint group and MUST include the local OCIR region where the repository is created.
	// Required when repositoryType is MULTI_REGION.
	MultiRegionPeers []string `mandatory:"false" json:"multiRegionPeers"`

	// Whether the repository is immutable. Images cannot be overwritten in an immutable repository.
	IsImmutable *bool `mandatory:"false" json:"isImmutable"`

	// Whether the repository is public. A public repository allows unauthenticated access.
	IsPublic *bool `mandatory:"false" json:"isPublic"`

	Readme *ContainerRepositoryReadme `mandatory:"false" json:"readme"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateContainerRepositoryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateContainerRepositoryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingContainerRepositoryRepositoryTypeEnum(string(m.RepositoryType)); !ok && m.RepositoryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RepositoryType: %s. Supported values are: %s.", m.RepositoryType, strings.Join(GetContainerRepositoryRepositoryTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
