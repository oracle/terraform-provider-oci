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

// ModelConfigurationDetails The model configuration details.
type ModelConfigurationDetails struct {

	// The OCID of the model you want to deploy.
	ModelId *string `mandatory:"true" json:"modelId"`

	InstanceConfiguration *InstanceConfiguration `mandatory:"true" json:"instanceConfiguration"`

	ScalingPolicy ScalingPolicy `mandatory:"false" json:"scalingPolicy"`

	// The minimum network bandwidth for the model deployment.
	BandwidthMbps *int `mandatory:"false" json:"bandwidthMbps"`

	// The maximum network bandwidth for the model deployment.
	MaximumBandwidthMbps *int `mandatory:"false" json:"maximumBandwidthMbps"`
}

func (m ModelConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModelConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ModelConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ScalingPolicy         scalingpolicy          `json:"scalingPolicy"`
		BandwidthMbps         *int                   `json:"bandwidthMbps"`
		MaximumBandwidthMbps  *int                   `json:"maximumBandwidthMbps"`
		ModelId               *string                `json:"modelId"`
		InstanceConfiguration *InstanceConfiguration `json:"instanceConfiguration"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
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

	m.ModelId = model.ModelId

	m.InstanceConfiguration = model.InstanceConfiguration

	return
}
