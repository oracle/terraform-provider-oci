// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OcirContainerJobEnvironmentConfigurationDetails Environment configuration based on container image stored in OCI Container Registry.
type OcirContainerJobEnvironmentConfigurationDetails struct {

	// The full path to the Oracle Container Repository (OCIR) registry, image, and tag in a canonical format.
	// Acceptable format:
	// `<region>.ocir.io/<registry>/<image>:<tag>`
	// `<region>.ocir.io/<registry>/<image>:<tag>@digest`
	Image *string `mandatory:"true" json:"image"`

	// The container image run CMD (https://docs.docker.com/engine/reference/builder/#cmd) as a list of strings.
	// Use `CMD` as arguments to the `ENTRYPOINT` or the only command to run in the absence of an `ENTRYPOINT`.
	// The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes.
	Cmd []string `mandatory:"false" json:"cmd"`

	// The container image run ENTRYPOINT (https://docs.docker.com/engine/reference/builder/#entrypoint) as a list of strings.
	// Accept the `CMD` as extra arguments.
	// The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes.
	// More information on how `CMD` and `ENTRYPOINT` interact are here (https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact).
	Entrypoint []string `mandatory:"false" json:"entrypoint"`

	// The digest of the container image. For example,
	// `sha256:881303a6b2738834d795a32b4a98eb0e5e3d1cad590a712d1e04f9b2fa90a030`
	ImageDigest *string `mandatory:"false" json:"imageDigest"`

	// OCID of the container image signature
	ImageSignatureId *string `mandatory:"false" json:"imageSignatureId"`
}

func (m OcirContainerJobEnvironmentConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OcirContainerJobEnvironmentConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OcirContainerJobEnvironmentConfigurationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOcirContainerJobEnvironmentConfigurationDetails OcirContainerJobEnvironmentConfigurationDetails
	s := struct {
		DiscriminatorParam string `json:"jobEnvironmentType"`
		MarshalTypeOcirContainerJobEnvironmentConfigurationDetails
	}{
		"OCIR_CONTAINER",
		(MarshalTypeOcirContainerJobEnvironmentConfigurationDetails)(m),
	}

	return json.Marshal(&s)
}
