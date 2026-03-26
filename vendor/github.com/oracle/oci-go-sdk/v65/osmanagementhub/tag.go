// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Tag Metadata tag assigned to a resource.
type Tag interface {

	// The key of the tag.
	GetKey() *string

	// The value associated with the tag key.
	GetValue() *string
}

type tag struct {
	JsonData []byte
	Key      *string `mandatory:"false" json:"key"`
	Value    *string `mandatory:"false" json:"value"`
	Type     string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *tag) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertag tag
	s := struct {
		Model Unmarshalertag
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.Value = s.Model.Value
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *tag) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "FREEFORM":
		mm := FreeFormTag{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEFINED":
		mm := DefinedTag{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for Tag: %s.", m.Type)
		return *m, nil
	}
}

// GetKey returns Key
func (m tag) GetKey() *string {
	return m.Key
}

// GetValue returns Value
func (m tag) GetValue() *string {
	return m.Value
}

func (m tag) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m tag) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TagTypeEnum Enum with underlying type: string
type TagTypeEnum string

// Set of constants representing the allowable values for TagTypeEnum
const (
	TagTypeDefined  TagTypeEnum = "DEFINED"
	TagTypeFreeform TagTypeEnum = "FREEFORM"
)

var mappingTagTypeEnum = map[string]TagTypeEnum{
	"DEFINED":  TagTypeDefined,
	"FREEFORM": TagTypeFreeform,
}

var mappingTagTypeEnumLowerCase = map[string]TagTypeEnum{
	"defined":  TagTypeDefined,
	"freeform": TagTypeFreeform,
}

// GetTagTypeEnumValues Enumerates the set of values for TagTypeEnum
func GetTagTypeEnumValues() []TagTypeEnum {
	values := make([]TagTypeEnum, 0)
	for _, v := range mappingTagTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTagTypeEnumStringValues Enumerates the set of values in String for TagTypeEnum
func GetTagTypeEnumStringValues() []string {
	return []string{
		"DEFINED",
		"FREEFORM",
	}
}

// GetMappingTagTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTagTypeEnum(val string) (TagTypeEnum, bool) {
	enum, ok := mappingTagTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
