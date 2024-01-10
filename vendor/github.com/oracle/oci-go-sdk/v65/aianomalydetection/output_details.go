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

// OutputDetails Detect anomaly job output details.
type OutputDetails interface {
}

type outputdetails struct {
	JsonData   []byte
	OutputType string `json:"outputType"`
}

// UnmarshalJSON unmarshals json
func (m *outputdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleroutputdetails outputdetails
	s := struct {
		Model Unmarshaleroutputdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.OutputType = s.Model.OutputType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *outputdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.OutputType {
	case "OBJECT_STORAGE":
		mm := ObjectStoreOutputDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for OutputDetails: %s.", m.OutputType)
		return *m, nil
	}
}

func (m outputdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m outputdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OutputDetailsOutputTypeEnum Enum with underlying type: string
type OutputDetailsOutputTypeEnum string

// Set of constants representing the allowable values for OutputDetailsOutputTypeEnum
const (
	OutputDetailsOutputTypeObjectStorage OutputDetailsOutputTypeEnum = "OBJECT_STORAGE"
)

var mappingOutputDetailsOutputTypeEnum = map[string]OutputDetailsOutputTypeEnum{
	"OBJECT_STORAGE": OutputDetailsOutputTypeObjectStorage,
}

var mappingOutputDetailsOutputTypeEnumLowerCase = map[string]OutputDetailsOutputTypeEnum{
	"object_storage": OutputDetailsOutputTypeObjectStorage,
}

// GetOutputDetailsOutputTypeEnumValues Enumerates the set of values for OutputDetailsOutputTypeEnum
func GetOutputDetailsOutputTypeEnumValues() []OutputDetailsOutputTypeEnum {
	values := make([]OutputDetailsOutputTypeEnum, 0)
	for _, v := range mappingOutputDetailsOutputTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOutputDetailsOutputTypeEnumStringValues Enumerates the set of values in String for OutputDetailsOutputTypeEnum
func GetOutputDetailsOutputTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
	}
}

// GetMappingOutputDetailsOutputTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOutputDetailsOutputTypeEnum(val string) (OutputDetailsOutputTypeEnum, bool) {
	enum, ok := mappingOutputDetailsOutputTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
