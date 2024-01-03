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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NamedEntityRecognitionModelMetrics Model level named entity recognition metrics
type NamedEntityRecognitionModelMetrics struct {

	// F1-score, is a measure of a model’s accuracy on a dataset
	MicroF1 *float32 `mandatory:"true" json:"microF1"`

	// Precision refers to the number of true positives divided by the total number of positive predictions (i.e., the number of true positives plus the number of false positives)
	MicroPrecision *float32 `mandatory:"true" json:"microPrecision"`

	// Measures the model's ability to predict actual positive classes. It is the ratio between the predicted true positives and what was actually tagged. The recall metric reveals how many of the predicted classes are correct.
	MicroRecall *float32 `mandatory:"true" json:"microRecall"`

	// F1-score, is a measure of a model’s accuracy on a dataset
	MacroF1 *float32 `mandatory:"true" json:"macroF1"`

	// Precision refers to the number of true positives divided by the total number of positive predictions (i.e., the number of true positives plus the number of false positives)
	MacroPrecision *float32 `mandatory:"true" json:"macroPrecision"`

	// Measures the model's ability to predict actual positive classes. It is the ratio between the predicted true positives and what was actually tagged. The recall metric reveals how many of the predicted classes are correct.
	MacroRecall *float32 `mandatory:"true" json:"macroRecall"`

	// F1-score, is a measure of a model’s accuracy on a dataset
	WeightedF1 *float32 `mandatory:"false" json:"weightedF1"`

	// Precision refers to the number of true positives divided by the total number of positive predictions (i.e., the number of true positives plus the number of false positives)
	WeightedPrecision *float32 `mandatory:"false" json:"weightedPrecision"`

	// Measures the model's ability to predict actual positive classes. It is the ratio between the predicted true positives and what was actually tagged. The recall metric reveals how many of the predicted classes are correct.
	WeightedRecall *float32 `mandatory:"false" json:"weightedRecall"`
}

func (m NamedEntityRecognitionModelMetrics) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NamedEntityRecognitionModelMetrics) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
