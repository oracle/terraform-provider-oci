// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateGenericRepositoryDetails Details for updating an artifact repository.
type UpdateGenericRepositoryDetails struct {

	// The repository name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The repository description.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

// GetDisplayName returns DisplayName
func (m UpdateGenericRepositoryDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m UpdateGenericRepositoryDetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m UpdateGenericRepositoryDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateGenericRepositoryDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateGenericRepositoryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateGenericRepositoryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateGenericRepositoryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateGenericRepositoryDetails UpdateGenericRepositoryDetails
	s := struct {
		DiscriminatorParam string `json:"repositoryType"`
		MarshalTypeUpdateGenericRepositoryDetails
	}{
		"GENERIC",
		(MarshalTypeUpdateGenericRepositoryDetails)(m),
	}

	return json.Marshal(&s)
}
