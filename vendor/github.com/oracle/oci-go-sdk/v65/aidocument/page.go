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

// Page One page document analysis result.
type Page struct {

	// The document page number.
	PageNumber *int `mandatory:"true" json:"pageNumber"`

	Dimensions *Dimensions `mandatory:"false" json:"dimensions"`

	// An array of detected document types.
	DetectedDocumentTypes []DetectedDocumentType `mandatory:"false" json:"detectedDocumentTypes"`

	// An array of detected languages.
	DetectedLanguages []DetectedLanguage `mandatory:"false" json:"detectedLanguages"`

	// The words detected on the page.
	Words []Word `mandatory:"false" json:"words"`

	// The lines of text detected on the page.
	Lines []Line `mandatory:"false" json:"lines"`

	// The tables detected on the page.
	Tables []Table `mandatory:"false" json:"tables"`

	// The form fields detected on the page.
	DocumentFields []DocumentField `mandatory:"false" json:"documentFields"`
}

func (m Page) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Page) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
