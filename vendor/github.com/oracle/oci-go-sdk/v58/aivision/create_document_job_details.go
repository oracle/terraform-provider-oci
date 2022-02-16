// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// VisionService API
//
// A description of the VisionService API.
//

package aivision

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CreateDocumentJobDetails Details about the batch document analysis.
type CreateDocumentJobDetails struct {
	InputLocation InputLocation `mandatory:"true" json:"inputLocation"`

	// List of document analysis types requested.
	Features []DocumentFeature `mandatory:"true" json:"features"`

	OutputLocation *OutputLocation `mandatory:"true" json:"outputLocation"`

	// Compartment identifier from the requester.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Document job display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Language of the document, abbreviated according to ISO 639-2.
	Language DocumentLanguageEnum `mandatory:"false" json:"language,omitempty"`

	// The type of documents.
	DocumentType DocumentTypeEnum `mandatory:"false" json:"documentType,omitempty"`

	// Whether to generate a Zip file containing the results.
	IsZipOutputEnabled *bool `mandatory:"false" json:"isZipOutputEnabled"`
}

func (m CreateDocumentJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDocumentJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDocumentLanguageEnum(string(m.Language)); !ok && m.Language != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Language: %s. Supported values are: %s.", m.Language, strings.Join(GetDocumentLanguageEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDocumentTypeEnum(string(m.DocumentType)); !ok && m.DocumentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DocumentType: %s. Supported values are: %s.", m.DocumentType, strings.Join(GetDocumentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateDocumentJobDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CompartmentId      *string              `json:"compartmentId"`
		DisplayName        *string              `json:"displayName"`
		Language           DocumentLanguageEnum `json:"language"`
		DocumentType       DocumentTypeEnum     `json:"documentType"`
		IsZipOutputEnabled *bool                `json:"isZipOutputEnabled"`
		InputLocation      inputlocation        `json:"inputLocation"`
		Features           []documentfeature    `json:"features"`
		OutputLocation     *OutputLocation      `json:"outputLocation"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.Language = model.Language

	m.DocumentType = model.DocumentType

	m.IsZipOutputEnabled = model.IsZipOutputEnabled

	nn, e = model.InputLocation.UnmarshalPolymorphicJSON(model.InputLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.InputLocation = nn.(InputLocation)
	} else {
		m.InputLocation = nil
	}

	m.Features = make([]DocumentFeature, len(model.Features))
	for i, n := range model.Features {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Features[i] = nn.(DocumentFeature)
		} else {
			m.Features[i] = nil
		}
	}

	m.OutputLocation = model.OutputLocation

	return
}
