// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Document Understanding API
//
// Document AI helps customers perform various analysis on their documents. If a customer has lots of documents, they can process them in batch using asynchronous API endpoints.
//

package aidocument

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

	// The document text extraction model version.
	TextExtractionModelVersion *string `mandatory:"false" json:"textExtractionModelVersion"`

	// The document keyValue extraction model version.
	KeyValueExtractionModelVersion *string `mandatory:"false" json:"keyValueExtractionModelVersion"`

	// The document table extraction model version.
	TableExtractionModelVersion *string `mandatory:"false" json:"tableExtractionModelVersion"`

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
