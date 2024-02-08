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

// DocumentFeature The type of document analysis.
type DocumentFeature interface {
}

type documentfeature struct {
	JsonData    []byte
	FeatureType string `json:"featureType"`
}

// UnmarshalJSON unmarshals json
func (m *documentfeature) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdocumentfeature documentfeature
	s := struct {
		Model Unmarshalerdocumentfeature
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.FeatureType = s.Model.FeatureType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *documentfeature) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.FeatureType {
	case "DOCUMENT_CLASSIFICATION":
		mm := DocumentClassificationFeature{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "KEY_VALUE_EXTRACTION":
		mm := DocumentKeyValueExtractionFeature{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LANGUAGE_CLASSIFICATION":
		mm := DocumentLanguageClassificationFeature{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TEXT_EXTRACTION":
		mm := DocumentTextExtractionFeature{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TABLE_EXTRACTION":
		mm := DocumentTableExtractionFeature{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DocumentFeature: %s.", m.FeatureType)
		return *m, nil
	}
}

func (m documentfeature) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m documentfeature) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DocumentFeatureFeatureTypeEnum Enum with underlying type: string
type DocumentFeatureFeatureTypeEnum string

// Set of constants representing the allowable values for DocumentFeatureFeatureTypeEnum
const (
	DocumentFeatureFeatureTypeLanguageClassification DocumentFeatureFeatureTypeEnum = "LANGUAGE_CLASSIFICATION"
	DocumentFeatureFeatureTypeTextExtraction         DocumentFeatureFeatureTypeEnum = "TEXT_EXTRACTION"
	DocumentFeatureFeatureTypeTableExtraction        DocumentFeatureFeatureTypeEnum = "TABLE_EXTRACTION"
	DocumentFeatureFeatureTypeKeyValueExtraction     DocumentFeatureFeatureTypeEnum = "KEY_VALUE_EXTRACTION"
	DocumentFeatureFeatureTypeDocumentClassification DocumentFeatureFeatureTypeEnum = "DOCUMENT_CLASSIFICATION"
)

var mappingDocumentFeatureFeatureTypeEnum = map[string]DocumentFeatureFeatureTypeEnum{
	"LANGUAGE_CLASSIFICATION": DocumentFeatureFeatureTypeLanguageClassification,
	"TEXT_EXTRACTION":         DocumentFeatureFeatureTypeTextExtraction,
	"TABLE_EXTRACTION":        DocumentFeatureFeatureTypeTableExtraction,
	"KEY_VALUE_EXTRACTION":    DocumentFeatureFeatureTypeKeyValueExtraction,
	"DOCUMENT_CLASSIFICATION": DocumentFeatureFeatureTypeDocumentClassification,
}

var mappingDocumentFeatureFeatureTypeEnumLowerCase = map[string]DocumentFeatureFeatureTypeEnum{
	"language_classification": DocumentFeatureFeatureTypeLanguageClassification,
	"text_extraction":         DocumentFeatureFeatureTypeTextExtraction,
	"table_extraction":        DocumentFeatureFeatureTypeTableExtraction,
	"key_value_extraction":    DocumentFeatureFeatureTypeKeyValueExtraction,
	"document_classification": DocumentFeatureFeatureTypeDocumentClassification,
}

// GetDocumentFeatureFeatureTypeEnumValues Enumerates the set of values for DocumentFeatureFeatureTypeEnum
func GetDocumentFeatureFeatureTypeEnumValues() []DocumentFeatureFeatureTypeEnum {
	values := make([]DocumentFeatureFeatureTypeEnum, 0)
	for _, v := range mappingDocumentFeatureFeatureTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDocumentFeatureFeatureTypeEnumStringValues Enumerates the set of values in String for DocumentFeatureFeatureTypeEnum
func GetDocumentFeatureFeatureTypeEnumStringValues() []string {
	return []string{
		"LANGUAGE_CLASSIFICATION",
		"TEXT_EXTRACTION",
		"TABLE_EXTRACTION",
		"KEY_VALUE_EXTRACTION",
		"DOCUMENT_CLASSIFICATION",
	}
}

// GetMappingDocumentFeatureFeatureTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDocumentFeatureFeatureTypeEnum(val string) (DocumentFeatureFeatureTypeEnum, bool) {
	enum, ok := mappingDocumentFeatureFeatureTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
