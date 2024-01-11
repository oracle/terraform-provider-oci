// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service API
//
// **Generative AI Service**
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable LLMs that cover a wide range of use cases for text generation. Use the playground to try out the models out-of-the-box or create and host your own fine-tuned custom models based on your own data on dedicated AI clusters.
//

package generativeai

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TextGenerationModelMetrics **TextGenerationModelMetrics**
// The text generation model metrics of the fine-tuning process.
type TextGenerationModelMetrics struct {

	// Fine-tuned model accuracy.
	FinalAccuracy *float64 `mandatory:"false" json:"finalAccuracy"`

	// Fine-tuned model loss.
	FinalLoss *float64 `mandatory:"false" json:"finalLoss"`
}

func (m TextGenerationModelMetrics) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TextGenerationModelMetrics) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TextGenerationModelMetrics) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTextGenerationModelMetrics TextGenerationModelMetrics
	s := struct {
		DiscriminatorParam string `json:"modelMetricsType"`
		MarshalTypeTextGenerationModelMetrics
	}{
		"TEXT_GENERATION_MODEL_METRICS",
		(MarshalTypeTextGenerationModelMetrics)(m),
	}

	return json.Marshal(&s)
}
