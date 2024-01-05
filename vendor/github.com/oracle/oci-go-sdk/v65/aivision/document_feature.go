// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vision API
//
// Using Vision, you can upload images to detect and classify objects in them. If you have lots of images, you can process them in batch using asynchronous API endpoints. Vision's features are thematically split between Document AI for document-centric images, and Image Analysis for object and scene-based images. Pretrained models and custom models are supported.
//

package aivision

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
	case "TABLE_DETECTION":
		mm := DocumentTableDetectionFeature{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "KEY_VALUE_DETECTION":
		mm := DocumentKeyValueDetectionFeature{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DOCUMENT_CLASSIFICATION":
		mm := DocumentClassificationFeature{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TEXT_DETECTION":
		mm := DocumentTextDetectionFeature{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LANGUAGE_CLASSIFICATION":
		mm := DocumentLanguageClassificationFeature{}
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
	DocumentFeatureFeatureTypeTextDetection          DocumentFeatureFeatureTypeEnum = "TEXT_DETECTION"
	DocumentFeatureFeatureTypeTableDetection         DocumentFeatureFeatureTypeEnum = "TABLE_DETECTION"
	DocumentFeatureFeatureTypeKeyValueDetection      DocumentFeatureFeatureTypeEnum = "KEY_VALUE_DETECTION"
	DocumentFeatureFeatureTypeDocumentClassification DocumentFeatureFeatureTypeEnum = "DOCUMENT_CLASSIFICATION"
)

var mappingDocumentFeatureFeatureTypeEnum = map[string]DocumentFeatureFeatureTypeEnum{
	"LANGUAGE_CLASSIFICATION": DocumentFeatureFeatureTypeLanguageClassification,
	"TEXT_DETECTION":          DocumentFeatureFeatureTypeTextDetection,
	"TABLE_DETECTION":         DocumentFeatureFeatureTypeTableDetection,
	"KEY_VALUE_DETECTION":     DocumentFeatureFeatureTypeKeyValueDetection,
	"DOCUMENT_CLASSIFICATION": DocumentFeatureFeatureTypeDocumentClassification,
}

var mappingDocumentFeatureFeatureTypeEnumLowerCase = map[string]DocumentFeatureFeatureTypeEnum{
	"language_classification": DocumentFeatureFeatureTypeLanguageClassification,
	"text_detection":          DocumentFeatureFeatureTypeTextDetection,
	"table_detection":         DocumentFeatureFeatureTypeTableDetection,
	"key_value_detection":     DocumentFeatureFeatureTypeKeyValueDetection,
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
		"TEXT_DETECTION",
		"TABLE_DETECTION",
		"KEY_VALUE_DETECTION",
		"DOCUMENT_CLASSIFICATION",
	}
}

// GetMappingDocumentFeatureFeatureTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDocumentFeatureFeatureTypeEnum(val string) (DocumentFeatureFeatureTypeEnum, bool) {
	enum, ok := mappingDocumentFeatureFeatureTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
