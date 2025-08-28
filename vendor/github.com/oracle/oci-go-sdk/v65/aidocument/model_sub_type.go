// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ModelSubType Sub type model based on the model type.
type ModelSubType interface {
}

type modelsubtype struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *modelsubtype) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermodelsubtype modelsubtype
	s := struct {
		Model Unmarshalermodelsubtype
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *modelsubtype) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "PRE_TRAINED_KEY_VALUE_EXTRACTION":
		mm := KvModelSubType{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION":
		mm := DocumentElementsSubType{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ModelSubType: %s.", m.ModelType)
		return *m, nil
	}
}

func (m modelsubtype) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m modelsubtype) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ModelSubTypeModelTypeEnum Enum with underlying type: string
type ModelSubTypeModelTypeEnum string

// Set of constants representing the allowable values for ModelSubTypeModelTypeEnum
const (
	ModelSubTypeModelTypeKeyValueExtraction         ModelSubTypeModelTypeEnum = "PRE_TRAINED_KEY_VALUE_EXTRACTION"
	ModelSubTypeModelTypeDocumentElementsExtraction ModelSubTypeModelTypeEnum = "PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION"
)

var mappingModelSubTypeModelTypeEnum = map[string]ModelSubTypeModelTypeEnum{
	"PRE_TRAINED_KEY_VALUE_EXTRACTION":         ModelSubTypeModelTypeKeyValueExtraction,
	"PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION": ModelSubTypeModelTypeDocumentElementsExtraction,
}

var mappingModelSubTypeModelTypeEnumLowerCase = map[string]ModelSubTypeModelTypeEnum{
	"pre_trained_key_value_extraction":         ModelSubTypeModelTypeKeyValueExtraction,
	"pre_trained_document_elements_extraction": ModelSubTypeModelTypeDocumentElementsExtraction,
}

// GetModelSubTypeModelTypeEnumValues Enumerates the set of values for ModelSubTypeModelTypeEnum
func GetModelSubTypeModelTypeEnumValues() []ModelSubTypeModelTypeEnum {
	values := make([]ModelSubTypeModelTypeEnum, 0)
	for _, v := range mappingModelSubTypeModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetModelSubTypeModelTypeEnumStringValues Enumerates the set of values in String for ModelSubTypeModelTypeEnum
func GetModelSubTypeModelTypeEnumStringValues() []string {
	return []string{
		"PRE_TRAINED_KEY_VALUE_EXTRACTION",
		"PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION",
	}
}

// GetMappingModelSubTypeModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelSubTypeModelTypeEnum(val string) (ModelSubTypeModelTypeEnum, bool) {
	enum, ok := mappingModelSubTypeModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
