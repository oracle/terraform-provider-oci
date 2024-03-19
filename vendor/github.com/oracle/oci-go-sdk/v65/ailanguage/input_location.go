// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
// pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InputLocation document location and other meta data about documents
// For TXT both ObjectStoragePrefixLocation and ObjectStorageFileNameLocation supported
// For CSV only ObjectStorageFileNameLocation is supported
type InputLocation interface {
}

type inputlocation struct {
	JsonData     []byte
	LocationType string `json:"locationType"`
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
	m.LocationType = s.Model.LocationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *inputlocation) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.LocationType {
	case "OBJECT_STORAGE_PREFIX":
		mm := ObjectStoragePrefixLocation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE_FILE_LIST":
		mm := ObjectStorageFileNameLocation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for InputLocation: %s.", m.LocationType)
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

// InputLocationLocationTypeEnum Enum with underlying type: string
type InputLocationLocationTypeEnum string

// Set of constants representing the allowable values for InputLocationLocationTypeEnum
const (
	InputLocationLocationTypePrefix   InputLocationLocationTypeEnum = "OBJECT_STORAGE_PREFIX"
	InputLocationLocationTypeFileList InputLocationLocationTypeEnum = "OBJECT_STORAGE_FILE_LIST"
)

var mappingInputLocationLocationTypeEnum = map[string]InputLocationLocationTypeEnum{
	"OBJECT_STORAGE_PREFIX":    InputLocationLocationTypePrefix,
	"OBJECT_STORAGE_FILE_LIST": InputLocationLocationTypeFileList,
}

var mappingInputLocationLocationTypeEnumLowerCase = map[string]InputLocationLocationTypeEnum{
	"object_storage_prefix":    InputLocationLocationTypePrefix,
	"object_storage_file_list": InputLocationLocationTypeFileList,
}

// GetInputLocationLocationTypeEnumValues Enumerates the set of values for InputLocationLocationTypeEnum
func GetInputLocationLocationTypeEnumValues() []InputLocationLocationTypeEnum {
	values := make([]InputLocationLocationTypeEnum, 0)
	for _, v := range mappingInputLocationLocationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInputLocationLocationTypeEnumStringValues Enumerates the set of values in String for InputLocationLocationTypeEnum
func GetInputLocationLocationTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE_PREFIX",
		"OBJECT_STORAGE_FILE_LIST",
	}
}

// GetMappingInputLocationLocationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInputLocationLocationTypeEnum(val string) (InputLocationLocationTypeEnum, bool) {
	enum, ok := mappingInputLocationLocationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
