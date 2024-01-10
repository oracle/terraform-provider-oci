// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// AnalyzeDocumentResult The document analysis results.
type AnalyzeDocumentResult struct {
	DocumentMetadata *DocumentMetadata `mandatory:"true" json:"documentMetadata"`

	// The array of a Page.
	Pages []Page `mandatory:"true" json:"pages"`

	// An array of detected document types.
	DetectedDocumentTypes []DetectedDocumentType `mandatory:"false" json:"detectedDocumentTypes"`

	// An array of detected languages.
	DetectedLanguages []DetectedLanguage `mandatory:"false" json:"detectedLanguages"`

	// The document classification model version.
	DocumentClassificationModelVersion *string `mandatory:"false" json:"documentClassificationModelVersion"`

	// The document language classification model version.
	LanguageClassificationModelVersion *string `mandatory:"false" json:"languageClassificationModelVersion"`

	// The document text detection model version.
	TextDetectionModelVersion *string `mandatory:"false" json:"textDetectionModelVersion"`

	// The document keyValue detection model version.
	KeyValueDetectionModelVersion *string `mandatory:"false" json:"keyValueDetectionModelVersion"`

	// The document table detection model version.
	TableDetectionModelVersion *string `mandatory:"false" json:"tableDetectionModelVersion"`

	// The errors encountered during document analysis.
	Errors []ProcessingError `mandatory:"false" json:"errors"`

	// The searchable PDF file that was generated.
	SearchablePdf []byte `mandatory:"false" json:"searchablePdf"`
}

func (m AnalyzeDocumentResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AnalyzeDocumentResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
