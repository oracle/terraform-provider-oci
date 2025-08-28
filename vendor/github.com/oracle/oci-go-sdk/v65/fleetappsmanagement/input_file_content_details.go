// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InputFileContentDetails Content Source details.
type InputFileContentDetails interface {
}

type inputfilecontentdetails struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *inputfilecontentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerinputfilecontentdetails inputfilecontentdetails
	s := struct {
		Model Unmarshalerinputfilecontentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *inputfilecontentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "OBJECT_STORAGE_BUCKET":
		mm := InputFileObjectStorageBucketContentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for InputFileContentDetails: %s.", m.SourceType)
		return *m, nil
	}
}

func (m inputfilecontentdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m inputfilecontentdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InputFileContentDetailsSourceTypeEnum Enum with underlying type: string
type InputFileContentDetailsSourceTypeEnum string

// Set of constants representing the allowable values for InputFileContentDetailsSourceTypeEnum
const (
	InputFileContentDetailsSourceTypeObjectStorageBucket InputFileContentDetailsSourceTypeEnum = "OBJECT_STORAGE_BUCKET"
)

var mappingInputFileContentDetailsSourceTypeEnum = map[string]InputFileContentDetailsSourceTypeEnum{
	"OBJECT_STORAGE_BUCKET": InputFileContentDetailsSourceTypeObjectStorageBucket,
}

var mappingInputFileContentDetailsSourceTypeEnumLowerCase = map[string]InputFileContentDetailsSourceTypeEnum{
	"object_storage_bucket": InputFileContentDetailsSourceTypeObjectStorageBucket,
}

// GetInputFileContentDetailsSourceTypeEnumValues Enumerates the set of values for InputFileContentDetailsSourceTypeEnum
func GetInputFileContentDetailsSourceTypeEnumValues() []InputFileContentDetailsSourceTypeEnum {
	values := make([]InputFileContentDetailsSourceTypeEnum, 0)
	for _, v := range mappingInputFileContentDetailsSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInputFileContentDetailsSourceTypeEnumStringValues Enumerates the set of values in String for InputFileContentDetailsSourceTypeEnum
func GetInputFileContentDetailsSourceTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE_BUCKET",
	}
}

// GetMappingInputFileContentDetailsSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInputFileContentDetailsSourceTypeEnum(val string) (InputFileContentDetailsSourceTypeEnum, bool) {
	enum, ok := mappingInputFileContentDetailsSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
