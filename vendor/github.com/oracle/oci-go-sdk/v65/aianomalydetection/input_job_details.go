// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// InputJobDetails Input details for detect anomaly job.
type InputJobDetails interface {
}

type inputjobdetails struct {
	JsonData  []byte
	InputType string `json:"inputType"`
}

// UnmarshalJSON unmarshals json
func (m *inputjobdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerinputjobdetails inputjobdetails
	s := struct {
		Model Unmarshalerinputjobdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.InputType = s.Model.InputType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *inputjobdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.InputType {
	case "INLINE":
		mm := InlineInputJobDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_LIST":
		mm := ObjectListInputJobDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for InputJobDetails: %s.", m.InputType)
		return *m, nil
	}
}

func (m inputjobdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m inputjobdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InputJobDetailsInputTypeEnum Enum with underlying type: string
type InputJobDetailsInputTypeEnum string

// Set of constants representing the allowable values for InputJobDetailsInputTypeEnum
const (
	InputJobDetailsInputTypeInline     InputJobDetailsInputTypeEnum = "INLINE"
	InputJobDetailsInputTypeObjectList InputJobDetailsInputTypeEnum = "OBJECT_LIST"
)

var mappingInputJobDetailsInputTypeEnum = map[string]InputJobDetailsInputTypeEnum{
	"INLINE":      InputJobDetailsInputTypeInline,
	"OBJECT_LIST": InputJobDetailsInputTypeObjectList,
}

var mappingInputJobDetailsInputTypeEnumLowerCase = map[string]InputJobDetailsInputTypeEnum{
	"inline":      InputJobDetailsInputTypeInline,
	"object_list": InputJobDetailsInputTypeObjectList,
}

// GetInputJobDetailsInputTypeEnumValues Enumerates the set of values for InputJobDetailsInputTypeEnum
func GetInputJobDetailsInputTypeEnumValues() []InputJobDetailsInputTypeEnum {
	values := make([]InputJobDetailsInputTypeEnum, 0)
	for _, v := range mappingInputJobDetailsInputTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInputJobDetailsInputTypeEnumStringValues Enumerates the set of values in String for InputJobDetailsInputTypeEnum
func GetInputJobDetailsInputTypeEnumStringValues() []string {
	return []string{
		"INLINE",
		"OBJECT_LIST",
	}
}

// GetMappingInputJobDetailsInputTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInputJobDetailsInputTypeEnum(val string) (InputJobDetailsInputTypeEnum, bool) {
	enum, ok := mappingInputJobDetailsInputTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
