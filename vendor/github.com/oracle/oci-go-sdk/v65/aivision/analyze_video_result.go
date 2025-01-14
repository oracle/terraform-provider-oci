// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vision API
//
// Using Vision, you can upload images to detect and classify objects in them. If you have lots of images, you can process them in batch using asynchronous API endpoints. Vision's features are thematically split between Document AI for document-centric images, and Image Analysis for object and scene-based images. Pretrained models and custom models are supported.
//

package aivision

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AnalyzeVideoResult Video analysis results.
type AnalyzeVideoResult struct {
	VideoMetadata *VideoMetadata `mandatory:"true" json:"videoMetadata"`

	// Detected labels in a video.
	VideoLabels []VideoLabel `mandatory:"false" json:"videoLabels"`

	// Detected objects in a video.
	VideoObjects []VideoObject `mandatory:"false" json:"videoObjects"`

	// Tracked objects in a video.
	VideoTrackedObjects []VideoTrackedObject `mandatory:"false" json:"videoTrackedObjects"`

	// Detected text in a video.
	VideoText []VideoText `mandatory:"false" json:"videoText"`

	// Detected faces in a video.
	VideoFaces []VideoFace `mandatory:"false" json:"videoFaces"`

	// The ontologyClasses of video labels.
	OntologyClasses []OntologyClass `mandatory:"false" json:"ontologyClasses"`

	// Label Detection model version.
	LabelDetectionModelVersion *string `mandatory:"false" json:"labelDetectionModelVersion"`

	// Object Detection model version.
	ObjectDetectionModelVersion *string `mandatory:"false" json:"objectDetectionModelVersion"`

	// Object Tracking model version.
	ObjectTrackingModelVersion *string `mandatory:"false" json:"objectTrackingModelVersion"`

	// Text Detection model version.
	TextDetectionModelVersion *string `mandatory:"false" json:"textDetectionModelVersion"`

	// Face Detection model version.
	FaceDetectionModelVersion *string `mandatory:"false" json:"faceDetectionModelVersion"`

	// Array of possible errors.
	Errors []ProcessingError `mandatory:"false" json:"errors"`
}

func (m AnalyzeVideoResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AnalyzeVideoResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
