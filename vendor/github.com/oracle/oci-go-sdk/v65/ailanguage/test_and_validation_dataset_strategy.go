// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
// pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TestAndValidationDatasetStrategy This information will be used capture training, testing and validation dataset.
type TestAndValidationDatasetStrategy struct {
	TestingDataset DatasetDetails `mandatory:"true" json:"testingDataset"`

	ValidationDataset DatasetDetails `mandatory:"false" json:"validationDataset"`
}

func (m TestAndValidationDatasetStrategy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TestAndValidationDatasetStrategy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TestAndValidationDatasetStrategy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTestAndValidationDatasetStrategy TestAndValidationDatasetStrategy
	s := struct {
		DiscriminatorParam string `json:"strategyType"`
		MarshalTypeTestAndValidationDatasetStrategy
	}{
		"TEST_AND_VALIDATION_DATASET",
		(MarshalTypeTestAndValidationDatasetStrategy)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *TestAndValidationDatasetStrategy) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ValidationDataset datasetdetails `json:"validationDataset"`
		TestingDataset    datasetdetails `json:"testingDataset"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.ValidationDataset.UnmarshalPolymorphicJSON(model.ValidationDataset.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ValidationDataset = nn.(DatasetDetails)
	} else {
		m.ValidationDataset = nil
	}

	nn, e = model.TestingDataset.UnmarshalPolymorphicJSON(model.TestingDataset.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TestingDataset = nn.(DatasetDetails)
	} else {
		m.TestingDataset = nil
	}

	return
}
