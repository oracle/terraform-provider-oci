// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// VisionService API
//
// A description of the VisionService API.
//

package aivision

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// AnalyzeImageResult Image analysis results.
type AnalyzeImageResult struct {

	// Detected objects.
	ImageObjects []ImageObject `mandatory:"false" json:"imageObjects"`

	// Image classification labels.
	Labels []Label `mandatory:"false" json:"labels"`

	// ontologyClasses of image labels.
	OntologyClasses []OntologyClass `mandatory:"false" json:"ontologyClasses"`

	ImageText *ImageText `mandatory:"false" json:"imageText"`

	// Image classification model version.
	ImageClassificationModelVersion *string `mandatory:"false" json:"imageClassificationModelVersion"`

	// Object detection model version.
	ObjectDetectionModelVersion *string `mandatory:"false" json:"objectDetectionModelVersion"`

	// Text detection model version.
	TextDetectionModelVersion *string `mandatory:"false" json:"textDetectionModelVersion"`

	// Errors encountered during image analysis.
	Errors []ProcessingError `mandatory:"false" json:"errors"`
}

func (m AnalyzeImageResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AnalyzeImageResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
