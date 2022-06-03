// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// InlineDetectAnomaliesRequest This is the specialised JSON format that we accept as Training data, with an additional
// field for 'requestType' which is a required field used deciding whether it is an inline
// request or contains embedded data.
type InlineDetectAnomaliesRequest struct {

	// The OCID of the trained model.
	ModelId *string `mandatory:"true" json:"modelId"`

	// List of signal names.
	SignalNames []string `mandatory:"true" json:"signalNames"`

	// Array containing data.
	Data []DataItem `mandatory:"true" json:"data"`

	// The value estimated during training would be used by default. Cutomer can choose to provide a custom value.
	Sensitivity *float32 `mandatory:"false" json:"sensitivity"`

	// State of the asset's behaviour calculated based on current and/or previous inference requests.
	InferenceState *string `mandatory:"false" json:"inferenceState"`
}

//GetModelId returns ModelId
func (m InlineDetectAnomaliesRequest) GetModelId() *string {
	return m.ModelId
}

//GetSensitivity returns Sensitivity
func (m InlineDetectAnomaliesRequest) GetSensitivity() *float32 {
	return m.Sensitivity
}

//GetInferenceState returns InferenceState
func (m InlineDetectAnomaliesRequest) GetInferenceState() *string {
	return m.InferenceState
}

func (m InlineDetectAnomaliesRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InlineDetectAnomaliesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m InlineDetectAnomaliesRequest) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInlineDetectAnomaliesRequest InlineDetectAnomaliesRequest
	s := struct {
		DiscriminatorParam string `json:"requestType"`
		MarshalTypeInlineDetectAnomaliesRequest
	}{
		"INLINE",
		(MarshalTypeInlineDetectAnomaliesRequest)(m),
	}

	return json.Marshal(&s)
}
