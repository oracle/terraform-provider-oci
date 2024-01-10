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
	case "OBJECT_LIST_INLINE_INPUT_LOCATION":
		mm := ObjectListInlineInputLocation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for InputLocation: %s.", m.SourceType)
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
	InputLocationSourceTypeObjectListInlineInputLocation InputLocationSourceTypeEnum = "OBJECT_LIST_INLINE_INPUT_LOCATION"
)

var mappingInputLocationSourceTypeEnum = map[string]InputLocationSourceTypeEnum{
	"OBJECT_LIST_INLINE_INPUT_LOCATION": InputLocationSourceTypeObjectListInlineInputLocation,
}

var mappingInputLocationSourceTypeEnumLowerCase = map[string]InputLocationSourceTypeEnum{
	"object_list_inline_input_location": InputLocationSourceTypeObjectListInlineInputLocation,
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
		"OBJECT_LIST_INLINE_INPUT_LOCATION",
	}
}

// GetMappingInputLocationSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInputLocationSourceTypeEnum(val string) (InputLocationSourceTypeEnum, bool) {
	enum, ok := mappingInputLocationSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
