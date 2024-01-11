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

// FineTuneDetails **FineTuneDetails**
// Details about creating this fine-tune model,including training data, validation data, and hyperparameters.
type FineTuneDetails struct {
	TrainingDataset Dataset `mandatory:"true" json:"trainingDataset"`

	// The OCID of the Dedicated AI Cluster this fine-tuning runs on.
	DedicatedAiClusterId *string `mandatory:"true" json:"dedicatedAiClusterId"`

	TrainingConfig TrainingConfig `mandatory:"false" json:"trainingConfig"`
}

func (m FineTuneDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FineTuneDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *FineTuneDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TrainingConfig       trainingconfig `json:"trainingConfig"`
		TrainingDataset      dataset        `json:"trainingDataset"`
		DedicatedAiClusterId *string        `json:"dedicatedAiClusterId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.TrainingConfig.UnmarshalPolymorphicJSON(model.TrainingConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TrainingConfig = nn.(TrainingConfig)
	} else {
		m.TrainingConfig = nil
	}

	nn, e = model.TrainingDataset.UnmarshalPolymorphicJSON(model.TrainingDataset.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TrainingDataset = nn.(Dataset)
	} else {
		m.TrainingDataset = nil
	}

	m.DedicatedAiClusterId = model.DedicatedAiClusterId

	return
}
