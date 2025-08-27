// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// StreamOutputLocation Details about a where results will be Sent
type StreamOutputLocation interface {
}

type streamoutputlocation struct {
	JsonData           []byte
	OutputLocationType string `json:"outputLocationType"`
}

// UnmarshalJSON unmarshals json
func (m *streamoutputlocation) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerstreamoutputlocation streamoutputlocation
	s := struct {
		Model Unmarshalerstreamoutputlocation
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.OutputLocationType = s.Model.OutputLocationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *streamoutputlocation) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.OutputLocationType {
	case "OBJECT_STORAGE":
		mm := ObjectStorageOutputLocation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for StreamOutputLocation: %s.", m.OutputLocationType)
		return *m, nil
	}
}

func (m streamoutputlocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m streamoutputlocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StreamOutputLocationOutputLocationTypeEnum Enum with underlying type: string
type StreamOutputLocationOutputLocationTypeEnum string

// Set of constants representing the allowable values for StreamOutputLocationOutputLocationTypeEnum
const (
	StreamOutputLocationOutputLocationTypeObjectStorage StreamOutputLocationOutputLocationTypeEnum = "OBJECT_STORAGE"
)

var mappingStreamOutputLocationOutputLocationTypeEnum = map[string]StreamOutputLocationOutputLocationTypeEnum{
	"OBJECT_STORAGE": StreamOutputLocationOutputLocationTypeObjectStorage,
}

var mappingStreamOutputLocationOutputLocationTypeEnumLowerCase = map[string]StreamOutputLocationOutputLocationTypeEnum{
	"object_storage": StreamOutputLocationOutputLocationTypeObjectStorage,
}

// GetStreamOutputLocationOutputLocationTypeEnumValues Enumerates the set of values for StreamOutputLocationOutputLocationTypeEnum
func GetStreamOutputLocationOutputLocationTypeEnumValues() []StreamOutputLocationOutputLocationTypeEnum {
	values := make([]StreamOutputLocationOutputLocationTypeEnum, 0)
	for _, v := range mappingStreamOutputLocationOutputLocationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetStreamOutputLocationOutputLocationTypeEnumStringValues Enumerates the set of values in String for StreamOutputLocationOutputLocationTypeEnum
func GetStreamOutputLocationOutputLocationTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
	}
}

// GetMappingStreamOutputLocationOutputLocationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStreamOutputLocationOutputLocationTypeEnum(val string) (StreamOutputLocationOutputLocationTypeEnum, bool) {
	enum, ok := mappingStreamOutputLocationOutputLocationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
