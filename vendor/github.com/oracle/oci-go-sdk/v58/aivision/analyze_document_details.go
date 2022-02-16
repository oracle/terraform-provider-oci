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

// AnalyzeDocumentDetails Details about how to analyze a document.
type AnalyzeDocumentDetails struct {

	// Types of document analysis requested.
	Features []DocumentFeature `mandatory:"true" json:"features"`

	Document DocumentDetails `mandatory:"true" json:"document"`

	// The OCID of the compartment that calls the API.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	OutputLocation *OutputLocation `mandatory:"false" json:"outputLocation"`

	// Language of the document, abbreviated according to ISO 639-2.
	Language DocumentLanguageEnum `mandatory:"false" json:"language,omitempty"`

	// The type of document.
	DocumentType DocumentTypeEnum `mandatory:"false" json:"documentType,omitempty"`
}

func (m AnalyzeDocumentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AnalyzeDocumentDetails) ValidateEnumValue() (bool, error) {
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
func (m *AnalyzeDocumentDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CompartmentId  *string              `json:"compartmentId"`
		OutputLocation *OutputLocation      `json:"outputLocation"`
		Language       DocumentLanguageEnum `json:"language"`
		DocumentType   DocumentTypeEnum     `json:"documentType"`
		Features       []documentfeature    `json:"features"`
		Document       documentdetails      `json:"document"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CompartmentId = model.CompartmentId

	m.OutputLocation = model.OutputLocation

	m.Language = model.Language

	m.DocumentType = model.DocumentType

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

	nn, e = model.Document.UnmarshalPolymorphicJSON(model.Document.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Document = nn.(DocumentDetails)
	} else {
		m.Document = nil
	}

	return
}
