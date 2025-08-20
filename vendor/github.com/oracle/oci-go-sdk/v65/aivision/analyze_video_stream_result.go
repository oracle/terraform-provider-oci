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

// AnalyzeVideoStreamResult Video stream analysis results.
type AnalyzeVideoStreamResult struct {
	VideoStreamMetadata *VideoStreamMetadata `mandatory:"true" json:"videoStreamMetadata"`

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of streamJob.
	StreamJobId *string `mandatory:"true" json:"streamJobId"`

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of StreamSource.
	StreamSourceId *string `mandatory:"true" json:"streamSourceId"`

	// time stamp of frame in utc.
	Timestamp *string `mandatory:"true" json:"timestamp"`

	OntologyClasses *OntologyClass `mandatory:"false" json:"ontologyClasses"`

	// Base 64 encoded frame
	ImageData *string `mandatory:"false" json:"imageData"`

	// Tracked objects in a video stream.
	VideoStreamObjects []VideoStreamObject `mandatory:"false" json:"videoStreamObjects"`

	// List of Object Tracking model versions.
	ObjectTrackingModelVersions []ModelVersionDetails `mandatory:"false" json:"objectTrackingModelVersions"`

	// List of Object Detection model versions.
	ObjectDetectionModelVersions []ModelVersionDetails `mandatory:"false" json:"objectDetectionModelVersions"`

	// Array of possible errors.
	Errors []ProcessingError `mandatory:"false" json:"errors"`
}

func (m AnalyzeVideoStreamResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AnalyzeVideoStreamResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
