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

// TextClassificationModelEvaluationResult Possible TXTC model error analysis
type TextClassificationModelEvaluationResult struct {

	// For CSV format location is rowId(1 is header) and for JSONL location is jsonL line sequence(1 is metadata)
	Location *string `mandatory:"true" json:"location"`

	// List of true(actual) labels in test data for multi class or multi label TextClassification
	TrueLabels []string `mandatory:"true" json:"trueLabels"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// List of predicted labels by custom multi class or multi label TextClassification model
	PredictedLabels []string `mandatory:"false" json:"predictedLabels"`
}

// GetFreeformTags returns FreeformTags
func (m TextClassificationModelEvaluationResult) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m TextClassificationModelEvaluationResult) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m TextClassificationModelEvaluationResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TextClassificationModelEvaluationResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TextClassificationModelEvaluationResult) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTextClassificationModelEvaluationResult TextClassificationModelEvaluationResult
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeTextClassificationModelEvaluationResult
	}{
		"TEXT_CLASSIFICATION",
		(MarshalTypeTextClassificationModelEvaluationResult)(m),
	}

	return json.Marshal(&s)
}
