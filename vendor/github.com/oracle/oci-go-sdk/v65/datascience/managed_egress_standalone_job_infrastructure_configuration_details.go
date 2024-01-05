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

// ManagedEgressStandaloneJobInfrastructureConfigurationDetails The standalone job infrastructure configuration with network egress settings preconfigured.
type ManagedEgressStandaloneJobInfrastructureConfigurationDetails struct {

	// The shape used to launch the job run instances.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// The size of the block storage volume to attach to the instance running the job
	BlockStorageSizeInGBs *int `mandatory:"true" json:"blockStorageSizeInGBs"`

	JobShapeConfigDetails *JobShapeConfigDetails `mandatory:"false" json:"jobShapeConfigDetails"`
}

func (m ManagedEgressStandaloneJobInfrastructureConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedEgressStandaloneJobInfrastructureConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ManagedEgressStandaloneJobInfrastructureConfigurationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeManagedEgressStandaloneJobInfrastructureConfigurationDetails ManagedEgressStandaloneJobInfrastructureConfigurationDetails
	s := struct {
		DiscriminatorParam string `json:"jobInfrastructureType"`
		MarshalTypeManagedEgressStandaloneJobInfrastructureConfigurationDetails
	}{
		"ME_STANDALONE",
		(MarshalTypeManagedEgressStandaloneJobInfrastructureConfigurationDetails)(m),
	}

	return json.Marshal(&s)
}
