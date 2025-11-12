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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RemoveContainerVersionDetails Remove version details.
type RemoveContainerVersionDetails struct {

	// The version to remove.
	Version *string `mandatory:"true" json:"version"`
}

func (m RemoveContainerVersionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RemoveContainerVersionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
