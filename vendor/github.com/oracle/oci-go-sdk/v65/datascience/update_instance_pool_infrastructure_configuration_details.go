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

// UpdateInstancePoolInfrastructureConfigurationDetails Instance Pool based Infrastructure configuration details for update.
type UpdateInstancePoolInfrastructureConfigurationDetails struct {
	InstanceConfiguration *InstanceConfiguration `mandatory:"false" json:"instanceConfiguration"`

	ScalingPolicy ScalingPolicy `mandatory:"false" json:"scalingPolicy"`

	// The minimum network bandwidth for the model deployment.
	BandwidthMbps *int `mandatory:"false" json:"bandwidthMbps"`

	// The maximum network bandwidth for the model deployment.
	MaximumBandwidthMbps *int `mandatory:"false" json:"maximumBandwidthMbps"`
}

func (m UpdateInstancePoolInfrastructureConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateInstancePoolInfrastructureConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateInstancePoolInfrastructureConfigurationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateInstancePoolInfrastructureConfigurationDetails UpdateInstancePoolInfrastructureConfigurationDetails
	s := struct {
		DiscriminatorParam string `json:"infrastructureType"`
		MarshalTypeUpdateInstancePoolInfrastructureConfigurationDetails
	}{
		"INSTANCE_POOL",
		(MarshalTypeUpdateInstancePoolInfrastructureConfigurationDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateInstancePoolInfrastructureConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		InstanceConfiguration *InstanceConfiguration `json:"instanceConfiguration"`
		ScalingPolicy         scalingpolicy          `json:"scalingPolicy"`
		BandwidthMbps         *int                   `json:"bandwidthMbps"`
		MaximumBandwidthMbps  *int                   `json:"maximumBandwidthMbps"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.InstanceConfiguration = model.InstanceConfiguration

	nn, e = model.ScalingPolicy.UnmarshalPolymorphicJSON(model.ScalingPolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ScalingPolicy = nn.(ScalingPolicy)
	} else {
		m.ScalingPolicy = nil
	}

	m.BandwidthMbps = model.BandwidthMbps

	m.MaximumBandwidthMbps = model.MaximumBandwidthMbps

	return
}
