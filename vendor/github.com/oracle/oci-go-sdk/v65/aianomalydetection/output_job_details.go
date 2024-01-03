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

// OutputJobDetails Output details for detect anomaly job.
type OutputJobDetails interface {
}

type outputjobdetails struct {
	JsonData   []byte
	OutputType string `json:"outputType"`
}

// UnmarshalJSON unmarshals json
func (m *outputjobdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleroutputjobdetails outputjobdetails
	s := struct {
		Model Unmarshaleroutputjobdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.OutputType = s.Model.OutputType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *outputjobdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.OutputType {
	case "OBJECT_STORAGE":
		mm := ObjectStorageLocation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for OutputJobDetails: %s.", m.OutputType)
		return *m, nil
	}
}

func (m outputjobdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m outputjobdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OutputJobDetailsOutputTypeEnum Enum with underlying type: string
type OutputJobDetailsOutputTypeEnum string

// Set of constants representing the allowable values for OutputJobDetailsOutputTypeEnum
const (
	OutputJobDetailsOutputTypeObjectStorage OutputJobDetailsOutputTypeEnum = "OBJECT_STORAGE"
)

var mappingOutputJobDetailsOutputTypeEnum = map[string]OutputJobDetailsOutputTypeEnum{
	"OBJECT_STORAGE": OutputJobDetailsOutputTypeObjectStorage,
}

var mappingOutputJobDetailsOutputTypeEnumLowerCase = map[string]OutputJobDetailsOutputTypeEnum{
	"object_storage": OutputJobDetailsOutputTypeObjectStorage,
}

// GetOutputJobDetailsOutputTypeEnumValues Enumerates the set of values for OutputJobDetailsOutputTypeEnum
func GetOutputJobDetailsOutputTypeEnumValues() []OutputJobDetailsOutputTypeEnum {
	values := make([]OutputJobDetailsOutputTypeEnum, 0)
	for _, v := range mappingOutputJobDetailsOutputTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOutputJobDetailsOutputTypeEnumStringValues Enumerates the set of values in String for OutputJobDetailsOutputTypeEnum
func GetOutputJobDetailsOutputTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
	}
}

// GetMappingOutputJobDetailsOutputTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOutputJobDetailsOutputTypeEnum(val string) (OutputJobDetailsOutputTypeEnum, bool) {
	enum, ok := mappingOutputJobDetailsOutputTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
