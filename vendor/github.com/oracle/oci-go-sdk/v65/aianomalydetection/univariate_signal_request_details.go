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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UnivariateSignalRequestDetails This contains metadata for each signal of univariate model
type UnivariateSignalRequestDetails struct {

	// Name of the signal
	SignalName *string `mandatory:"true" json:"signalName"`

	// Names of associated categorical signal
	CategoricalSignalNames []string `mandatory:"false" json:"categoricalSignalNames"`

	// Window of data to look at for each signal
	WindowSize *int `mandatory:"false" json:"windowSize"`

	// Estimate of anomalies in the dataset
	ContaminationRatio *float32 `mandatory:"false" json:"contaminationRatio"`

	// Algorithm to be used for training
	AlgorithmHint *string `mandatory:"false" json:"algorithmHint"`

	// Algorithm Parameters to be used for training
	AlgorithmHintParameters *string `mandatory:"false" json:"algorithmHintParameters"`
}

func (m UnivariateSignalRequestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnivariateSignalRequestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
