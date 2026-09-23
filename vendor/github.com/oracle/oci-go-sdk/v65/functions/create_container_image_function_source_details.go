// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateContainerImageFunctionSourceDetails Source details for creating a container image based function.
type CreateContainerImageFunctionSourceDetails struct {

	// The qualified name of the Docker image to use in the function, including the image tag.
	// The image should be in the OCI Registry that is in the same region as the function itself.
	// Example: `phx.ocir.io/ten/functions/function:0.0.1`
	Image *string `mandatory:"true" json:"image"`

	// The image digest for the version of the image that will be pulled when invoking this function.
	// If no value is specified, the digest currently associated with the image in the OCI Registry will be used.
	// Example: `sha256:ca0eeb6fb05351dfc8759c20733c91def84cb8007aa89a5bf606bc8b315b9fc7`
	ImageDigest *string `mandatory:"false" json:"imageDigest"`
}

func (m CreateContainerImageFunctionSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateContainerImageFunctionSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateContainerImageFunctionSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateContainerImageFunctionSourceDetails CreateContainerImageFunctionSourceDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeCreateContainerImageFunctionSourceDetails
	}{
		"CONTAINER_IMAGE",
		(MarshalTypeCreateContainerImageFunctionSourceDetails)(m),
	}

	return json.Marshal(&s)
}
