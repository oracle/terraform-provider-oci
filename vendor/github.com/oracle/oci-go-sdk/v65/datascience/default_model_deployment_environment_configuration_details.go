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

// DefaultModelDeploymentEnvironmentConfigurationDetails The environment configuration details object for managed container
type DefaultModelDeploymentEnvironmentConfigurationDetails struct {

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

func (m DefaultModelDeploymentEnvironmentConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DefaultModelDeploymentEnvironmentConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DefaultModelDeploymentEnvironmentConfigurationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDefaultModelDeploymentEnvironmentConfigurationDetails DefaultModelDeploymentEnvironmentConfigurationDetails
	s := struct {
		DiscriminatorParam string `json:"environmentConfigurationType"`
		MarshalTypeDefaultModelDeploymentEnvironmentConfigurationDetails
	}{
		"DEFAULT",
		(MarshalTypeDefaultModelDeploymentEnvironmentConfigurationDetails)(m),
	}

	return json.Marshal(&s)
}
