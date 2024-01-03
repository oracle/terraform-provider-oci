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

// OcirModelDeploymentEnvironmentConfigurationDetails The environment configuration details object for OCI Registry
type OcirModelDeploymentEnvironmentConfigurationDetails struct {

	// The full path to the Oracle Container Repository (OCIR) registry, image, and tag in a canonical format.
	// Acceptable format:
	// `<region>.ocir.io/<registry>/<image>:<tag>`
	// `<region>.ocir.io/<registry>/<image>:<tag>@digest`
	Image *string `mandatory:"true" json:"image"`

	// The digest of the container image. For example,
	// `sha256:881303a6b2738834d795a32b4a98eb0e5e3d1cad590a712d1e04f9b2fa90a030`
	ImageDigest *string `mandatory:"false" json:"imageDigest"`

	// The container image run CMD (https://docs.docker.com/engine/reference/builder/#cmd) as a list of strings.
	// Use `CMD` as arguments to the `ENTRYPOINT` or the only command to run in the absence of an `ENTRYPOINT`.
	// The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes.
	Cmd []string `mandatory:"false" json:"cmd"`

	// The container image run ENTRYPOINT (https://docs.docker.com/engine/reference/builder/#entrypoint) as a list of strings.
	// Accept the `CMD` as extra arguments.
	// The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes.
	// More information on how `CMD` and `ENTRYPOINT` interact are here (https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact).
	Entrypoint []string `mandatory:"false" json:"entrypoint"`

	// The port on which the web server serving the inference is running.
	// The port can be anything between `1024` and `65535`. The following ports cannot be used `24224`, `8446`, `8447`.
	ServerPort *int `mandatory:"false" json:"serverPort"`

	// The port on which the container HEALTHCHECK (https://docs.docker.com/engine/reference/builder/#healthcheck) would listen.
	// The port can be anything between `1024` and `65535`. The following ports cannot be used `24224`, `8446`, `8447`.
	HealthCheckPort *int `mandatory:"false" json:"healthCheckPort"`

	// Environment variables to set for the web server container.
	// The size of envVars must be less than 2048 bytes.
	// Key should be under 32 characters.
	// Key should contain only letters, digits and underscore (_)
	// Key should start with a letter.
	// Key should have at least 2 characters.
	// Key should not end with underscore eg. `TEST_`
	// Key if added cannot be empty. Value can be empty.
	// No specific size limits on individual Values. But overall environment variables is limited to 2048 bytes.
	// Key can't be reserved Model Deployment environment variables.
	EnvironmentVariables map[string]string `mandatory:"false" json:"environmentVariables"`
}

func (m OcirModelDeploymentEnvironmentConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OcirModelDeploymentEnvironmentConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OcirModelDeploymentEnvironmentConfigurationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOcirModelDeploymentEnvironmentConfigurationDetails OcirModelDeploymentEnvironmentConfigurationDetails
	s := struct {
		DiscriminatorParam string `json:"environmentConfigurationType"`
		MarshalTypeOcirModelDeploymentEnvironmentConfigurationDetails
	}{
		"OCIR_CONTAINER",
		(MarshalTypeOcirModelDeploymentEnvironmentConfigurationDetails)(m),
	}

	return json.Marshal(&s)
}
