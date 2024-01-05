// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling Management API
//
// Use Data Labeling Management API to create, list, edit & delete datasets.
//

package datalabelingservice

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TextFileTypeMetadata Metadata for files with text content.
type TextFileTypeMetadata interface {
}

type textfiletypemetadata struct {
	JsonData   []byte
	FormatType string `json:"formatType"`
}

// UnmarshalJSON unmarshals json
func (m *textfiletypemetadata) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertextfiletypemetadata textfiletypemetadata
	s := struct {
		Model Unmarshalertextfiletypemetadata
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.FormatType = s.Model.FormatType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *textfiletypemetadata) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.FormatType {
	case "DELIMITED":
		mm := DelimitedFileTypeMetadata{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for TextFileTypeMetadata: %s.", m.FormatType)
		return *m, nil
	}
}

func (m textfiletypemetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m textfiletypemetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TextFileTypeMetadataFormatTypeEnum Enum with underlying type: string
type TextFileTypeMetadataFormatTypeEnum string

// Set of constants representing the allowable values for TextFileTypeMetadataFormatTypeEnum
const (
	TextFileTypeMetadataFormatTypeDelimited TextFileTypeMetadataFormatTypeEnum = "DELIMITED"
)

var mappingTextFileTypeMetadataFormatTypeEnum = map[string]TextFileTypeMetadataFormatTypeEnum{
	"DELIMITED": TextFileTypeMetadataFormatTypeDelimited,
}

var mappingTextFileTypeMetadataFormatTypeEnumLowerCase = map[string]TextFileTypeMetadataFormatTypeEnum{
	"delimited": TextFileTypeMetadataFormatTypeDelimited,
}

// GetTextFileTypeMetadataFormatTypeEnumValues Enumerates the set of values for TextFileTypeMetadataFormatTypeEnum
func GetTextFileTypeMetadataFormatTypeEnumValues() []TextFileTypeMetadataFormatTypeEnum {
	values := make([]TextFileTypeMetadataFormatTypeEnum, 0)
	for _, v := range mappingTextFileTypeMetadataFormatTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTextFileTypeMetadataFormatTypeEnumStringValues Enumerates the set of values in String for TextFileTypeMetadataFormatTypeEnum
func GetTextFileTypeMetadataFormatTypeEnumStringValues() []string {
	return []string{
		"DELIMITED",
	}
}

// GetMappingTextFileTypeMetadataFormatTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTextFileTypeMetadataFormatTypeEnum(val string) (TextFileTypeMetadataFormatTypeEnum, bool) {
	enum, ok := mappingTextFileTypeMetadataFormatTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
