// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// StandaloneJobInfrastructureConfigurationDetails The standalone job infrastructure configuration.
type StandaloneJobInfrastructureConfigurationDetails struct {

	// The shape used to launch the job run instances.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// The subnet to create a secondary vnic in to attach to the instance running the job
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The size of the block storage volume to attach to the instance running the job
	BlockStorageSizeInGBs *int `mandatory:"true" json:"blockStorageSizeInGBs"`
}

func (m StandaloneJobInfrastructureConfigurationDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m StandaloneJobInfrastructureConfigurationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStandaloneJobInfrastructureConfigurationDetails StandaloneJobInfrastructureConfigurationDetails
	s := struct {
		DiscriminatorParam string `json:"jobInfrastructureType"`
		MarshalTypeStandaloneJobInfrastructureConfigurationDetails
	}{
		"STANDALONE",
		(MarshalTypeStandaloneJobInfrastructureConfigurationDetails)(m),
	}

	return json.Marshal(&s)
}
