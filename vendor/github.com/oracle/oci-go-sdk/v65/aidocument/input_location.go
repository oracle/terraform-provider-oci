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

// InputLocation The location of the inputs.
type InputLocation interface {
}

type inputlocation struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *inputlocation) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerinputlocation inputlocation
	s := struct {
		Model Unmarshalerinputlocation
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *inputlocation) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "INLINE_DOCUMENT_CONTENT":
		mm := InlineDocumentContent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE_LOCATIONS":
		mm := ObjectStorageLocations{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for InputLocation: %s.", m.SourceType)
		return *m, nil
	}
}

func (m inputlocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m inputlocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InputLocationSourceTypeEnum Enum with underlying type: string
type InputLocationSourceTypeEnum string

// Set of constants representing the allowable values for InputLocationSourceTypeEnum
const (
	InputLocationSourceTypeObjectStorageLocations InputLocationSourceTypeEnum = "OBJECT_STORAGE_LOCATIONS"
	InputLocationSourceTypeInlineDocumentContent  InputLocationSourceTypeEnum = "INLINE_DOCUMENT_CONTENT"
)

var mappingInputLocationSourceTypeEnum = map[string]InputLocationSourceTypeEnum{
	"OBJECT_STORAGE_LOCATIONS": InputLocationSourceTypeObjectStorageLocations,
	"INLINE_DOCUMENT_CONTENT":  InputLocationSourceTypeInlineDocumentContent,
}

var mappingInputLocationSourceTypeEnumLowerCase = map[string]InputLocationSourceTypeEnum{
	"object_storage_locations": InputLocationSourceTypeObjectStorageLocations,
	"inline_document_content":  InputLocationSourceTypeInlineDocumentContent,
}

// GetInputLocationSourceTypeEnumValues Enumerates the set of values for InputLocationSourceTypeEnum
func GetInputLocationSourceTypeEnumValues() []InputLocationSourceTypeEnum {
	values := make([]InputLocationSourceTypeEnum, 0)
	for _, v := range mappingInputLocationSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInputLocationSourceTypeEnumStringValues Enumerates the set of values in String for InputLocationSourceTypeEnum
func GetInputLocationSourceTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE_LOCATIONS",
		"INLINE_DOCUMENT_CONTENT",
	}
}

// GetMappingInputLocationSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInputLocationSourceTypeEnum(val string) (InputLocationSourceTypeEnum, bool) {
	enum, ok := mappingInputLocationSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
