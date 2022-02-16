// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// VisionService API
//
// A description of the VisionService API.
//

package aivision

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ImageDetails Details about an image to analyze.
type ImageDetails interface {
}

type imagedetails struct {
	JsonData []byte
	Source   string `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *imagedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerimagedetails imagedetails
	s := struct {
		Model Unmarshalerimagedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Source = s.Model.Source

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *imagedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "OBJECT_STORAGE":
		mm := ObjectStorageImageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INLINE":
		mm := InlineImageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m imagedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m imagedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ImageDetailsSourceEnum Enum with underlying type: string
type ImageDetailsSourceEnum string

// Set of constants representing the allowable values for ImageDetailsSourceEnum
const (
	ImageDetailsSourceInline        ImageDetailsSourceEnum = "INLINE"
	ImageDetailsSourceObjectStorage ImageDetailsSourceEnum = "OBJECT_STORAGE"
)

var mappingImageDetailsSourceEnum = map[string]ImageDetailsSourceEnum{
	"INLINE":         ImageDetailsSourceInline,
	"OBJECT_STORAGE": ImageDetailsSourceObjectStorage,
}

// GetImageDetailsSourceEnumValues Enumerates the set of values for ImageDetailsSourceEnum
func GetImageDetailsSourceEnumValues() []ImageDetailsSourceEnum {
	values := make([]ImageDetailsSourceEnum, 0)
	for _, v := range mappingImageDetailsSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetImageDetailsSourceEnumStringValues Enumerates the set of values in String for ImageDetailsSourceEnum
func GetImageDetailsSourceEnumStringValues() []string {
	return []string{
		"INLINE",
		"OBJECT_STORAGE",
	}
}

// GetMappingImageDetailsSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImageDetailsSourceEnum(val string) (ImageDetailsSourceEnum, bool) {
	mappingImageDetailsSourceEnumIgnoreCase := make(map[string]ImageDetailsSourceEnum)
	for k, v := range mappingImageDetailsSourceEnum {
		mappingImageDetailsSourceEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingImageDetailsSourceEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
