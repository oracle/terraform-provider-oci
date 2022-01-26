// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud AI Services API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// InlineDetectAnomaliesRequest This is the specialised JSON format that we accept as Training data, with an additional
// field for 'requestType' which is a required field used deciding whether it is an inline
// request or contains embedded data.
type InlineDetectAnomaliesRequest struct {

	// The OCID of the trained model。
	ModelId *string `mandatory:"true" json:"modelId"`

	// List of signal names.
	SignalNames []string `mandatory:"true" json:"signalNames"`

	// Array containing data.
	Data []DataItem `mandatory:"true" json:"data"`
}

//GetModelId returns ModelId
func (m InlineDetectAnomaliesRequest) GetModelId() *string {
	return m.ModelId
}

func (m InlineDetectAnomaliesRequest) String() string {
	return common.PointerString(m)
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
