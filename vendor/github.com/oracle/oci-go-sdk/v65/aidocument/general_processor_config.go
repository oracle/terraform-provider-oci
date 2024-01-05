// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Document Understanding API
//
// Document AI helps customers perform various analysis on their documents. If a customer has lots of documents, they can process them in batch using asynchronous API endpoints.
//

package aidocument

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GeneralProcessorConfig The configuration of a general processor.
type GeneralProcessorConfig struct {

	// The types of document analysis requested.
	Features []DocumentFeature `mandatory:"true" json:"features"`

	// Whether or not to generate a ZIP file containing the results.
	IsZipOutputEnabled *bool `mandatory:"false" json:"isZipOutputEnabled"`

	// The document language, abbreviated according to the BCP 47 Language-Tag syntax.
	Language *string `mandatory:"false" json:"language"`

	// The document type.
	DocumentType DocumentTypeEnum `mandatory:"false" json:"documentType,omitempty"`
}

func (m GeneralProcessorConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GeneralProcessorConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDocumentTypeEnum(string(m.DocumentType)); !ok && m.DocumentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DocumentType: %s. Supported values are: %s.", m.DocumentType, strings.Join(GetDocumentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GeneralProcessorConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGeneralProcessorConfig GeneralProcessorConfig
	s := struct {
		DiscriminatorParam string `json:"processorType"`
		MarshalTypeGeneralProcessorConfig
	}{
		"GENERAL",
		(MarshalTypeGeneralProcessorConfig)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *GeneralProcessorConfig) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DocumentType       DocumentTypeEnum  `json:"documentType"`
		IsZipOutputEnabled *bool             `json:"isZipOutputEnabled"`
		Language           *string           `json:"language"`
		Features           []documentfeature `json:"features"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DocumentType = model.DocumentType

	m.IsZipOutputEnabled = model.IsZipOutputEnabled

	m.Language = model.Language

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
	return
}
