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

// Page One page document analysis result.
type Page struct {

	// Document page number.
	PageNumber *int `mandatory:"true" json:"pageNumber"`

	Dimensions *Dimensions `mandatory:"false" json:"dimensions"`

	// An array of detected document types.
	DetectedDocumentTypes []DetectedDocumentType `mandatory:"false" json:"detectedDocumentTypes"`

	// An array of detected languages.
	DetectedLanguages []DetectedLanguage `mandatory:"false" json:"detectedLanguages"`

	// Words detected on the page.
	Words []Word `mandatory:"false" json:"words"`

	// Text lines detected on the page.
	Lines []Line `mandatory:"false" json:"lines"`

	// Tables detected on the page.
	Tables []Table `mandatory:"false" json:"tables"`

	// Form fields detected on the page.
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
