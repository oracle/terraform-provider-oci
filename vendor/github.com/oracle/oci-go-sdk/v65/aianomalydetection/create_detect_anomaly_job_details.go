// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Anomaly Detection API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDetectAnomalyJobDetails Base class for the DetectAnomalies async call. It contains the identifier that is
// used for deciding what type of request this is.
type CreateDetectAnomalyJobDetails struct {

	// The OCID of the compartment that starts the job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the trained model.
	ModelId *string `mandatory:"true" json:"modelId"`

	InputDetails InputDetails `mandatory:"true" json:"inputDetails"`

	OutputDetails OutputDetails `mandatory:"true" json:"outputDetails"`

	// A short description of the detect anomaly job.
	Description *string `mandatory:"false" json:"description"`

	// Detect anomaly job display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The value that customer can adjust to control the sensitivity of anomaly detection
	Sensitivity *float32 `mandatory:"false" json:"sensitivity"`

	// Flag to enable the service to return estimates for all data points rather than just the anomalous data points.
	AreAllEstimatesRequired *bool `mandatory:"false" json:"areAllEstimatesRequired"`
}

func (m CreateDetectAnomalyJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDetectAnomalyJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateDetectAnomalyJobDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description             *string       `json:"description"`
		DisplayName             *string       `json:"displayName"`
		Sensitivity             *float32      `json:"sensitivity"`
		AreAllEstimatesRequired *bool         `json:"areAllEstimatesRequired"`
		CompartmentId           *string       `json:"compartmentId"`
		ModelId                 *string       `json:"modelId"`
		InputDetails            inputdetails  `json:"inputDetails"`
		OutputDetails           outputdetails `json:"outputDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.Sensitivity = model.Sensitivity

	m.AreAllEstimatesRequired = model.AreAllEstimatesRequired

	m.CompartmentId = model.CompartmentId

	m.ModelId = model.ModelId

	nn, e = model.InputDetails.UnmarshalPolymorphicJSON(model.InputDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.InputDetails = nn.(InputDetails)
	} else {
		m.InputDetails = nil
	}

	nn, e = model.OutputDetails.UnmarshalPolymorphicJSON(model.OutputDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.OutputDetails = nn.(OutputDetails)
	} else {
		m.OutputDetails = nil
	}

	return
}
