// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Anomaly Detection API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InputDetails Detect anomaly asynchronous job details.
type InputDetails interface {
}

type inputdetails struct {
	JsonData  []byte
	InputType string `json:"inputType"`
}

// UnmarshalJSON unmarshals json
func (m *inputdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerinputdetails inputdetails
	s := struct {
		Model Unmarshalerinputdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.InputType = s.Model.InputType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *inputdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.InputType {
	case "BASE64_ENCODED":
		mm := EmbeddedInputDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_LIST":
		mm := ObjectListInputDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INLINE":
		mm := InlineInputDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for InputDetails: %s.", m.InputType)
		return *m, nil
	}
}

func (m inputdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m inputdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InputDetailsInputTypeEnum Enum with underlying type: string
type InputDetailsInputTypeEnum string

// Set of constants representing the allowable values for InputDetailsInputTypeEnum
const (
	InputDetailsInputTypeInline        InputDetailsInputTypeEnum = "INLINE"
	InputDetailsInputTypeBase64Encoded InputDetailsInputTypeEnum = "BASE64_ENCODED"
	InputDetailsInputTypeObjectList    InputDetailsInputTypeEnum = "OBJECT_LIST"
)

var mappingInputDetailsInputTypeEnum = map[string]InputDetailsInputTypeEnum{
	"INLINE":         InputDetailsInputTypeInline,
	"BASE64_ENCODED": InputDetailsInputTypeBase64Encoded,
	"OBJECT_LIST":    InputDetailsInputTypeObjectList,
}

var mappingInputDetailsInputTypeEnumLowerCase = map[string]InputDetailsInputTypeEnum{
	"inline":         InputDetailsInputTypeInline,
	"base64_encoded": InputDetailsInputTypeBase64Encoded,
	"object_list":    InputDetailsInputTypeObjectList,
}

// GetInputDetailsInputTypeEnumValues Enumerates the set of values for InputDetailsInputTypeEnum
func GetInputDetailsInputTypeEnumValues() []InputDetailsInputTypeEnum {
	values := make([]InputDetailsInputTypeEnum, 0)
	for _, v := range mappingInputDetailsInputTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInputDetailsInputTypeEnumStringValues Enumerates the set of values in String for InputDetailsInputTypeEnum
func GetInputDetailsInputTypeEnumStringValues() []string {
	return []string{
		"INLINE",
		"BASE64_ENCODED",
		"OBJECT_LIST",
	}
}

// GetMappingInputDetailsInputTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInputDetailsInputTypeEnum(val string) (InputDetailsInputTypeEnum, bool) {
	enum, ok := mappingInputDetailsInputTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
