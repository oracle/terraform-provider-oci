// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ModelDeployWorkloadConfigurationDetails The model deployment workload configuration.
type ModelDeployWorkloadConfigurationDetails struct {

	// The container image run CMD (https://docs.docker.com/engine/reference/builder/#cmd) as a list of strings.
	// Use `CMD` as arguments to the `ENTRYPOINT` or the only command to run in the absence of an `ENTRYPOINT`.
	// The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes.
	Cmd *string `mandatory:"true" json:"cmd"`

	// The port on which the web server serving the inference is running.
	// The port can be anything between `1024` and `65535`. The following ports cannot be used `24224`, `8446`, `8447`.
	ServerPort *int `mandatory:"true" json:"serverPort"`

	// The port on which the container HEALTHCHECK (https://docs.docker.com/engine/reference/builder/#healthcheck) would listen.
	// The port can be anything between `1024` and `65535`. The following ports cannot be used `24224`, `8446`, `8447`.
	HealthCheckPort *int `mandatory:"true" json:"healthCheckPort"`

	// The additional configurations
	AdditionalConfigurations map[string]string `mandatory:"false" json:"additionalConfigurations"`
}

func (m ModelDeployWorkloadConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModelDeployWorkloadConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ModelDeployWorkloadConfigurationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeModelDeployWorkloadConfigurationDetails ModelDeployWorkloadConfigurationDetails
	s := struct {
		DiscriminatorParam string `json:"workloadType"`
		MarshalTypeModelDeployWorkloadConfigurationDetails
	}{
		"MODEL_DEPLOYMENT",
		(MarshalTypeModelDeployWorkloadConfigurationDetails)(m),
	}

	return json.Marshal(&s)
}
