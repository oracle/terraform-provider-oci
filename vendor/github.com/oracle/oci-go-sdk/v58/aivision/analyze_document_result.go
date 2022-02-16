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

// AnalyzeDocumentResult Document analysis results.
type AnalyzeDocumentResult struct {
	DocumentMetadata *DocumentMetadata `mandatory:"true" json:"documentMetadata"`

	// Array of Page.
	Pages []Page `mandatory:"true" json:"pages"`

	// An array of detected document types.
	DetectedDocumentTypes []DetectedDocumentType `mandatory:"false" json:"detectedDocumentTypes"`

	// An array of detected languages.
	DetectedLanguages []DetectedLanguage `mandatory:"false" json:"detectedLanguages"`

	// Document classification model version.
	DocumentClassificationModelVersion *string `mandatory:"false" json:"documentClassificationModelVersion"`

	// Document language classification model version.
	LanguageClassificationModelVersion *string `mandatory:"false" json:"languageClassificationModelVersion"`

	// Document text detection model version.
	TextDetectionModelVersion *string `mandatory:"false" json:"textDetectionModelVersion"`

	// Document keyValue detection model version.
	KeyValueDetectionModelVersion *string `mandatory:"false" json:"keyValueDetectionModelVersion"`

	// Document table detection model version.
	TableDetectionModelVersion *string `mandatory:"false" json:"tableDetectionModelVersion"`

	// Errors encountered during document analysis.
	Errors []ProcessingError `mandatory:"false" json:"errors"`

	// Generated searchable PDF file.
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
