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

// DocumentDetails Details about a document to analyze.
type DocumentDetails interface {
}

type documentdetails struct {
	JsonData []byte
	Source   string `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *documentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdocumentdetails documentdetails
	s := struct {
		Model Unmarshalerdocumentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Source = s.Model.Source

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *documentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "OBJECT_STORAGE":
		mm := ObjectStorageDocumentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INLINE":
		mm := InlineDocumentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m documentdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m documentdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DocumentDetailsSourceEnum Enum with underlying type: string
type DocumentDetailsSourceEnum string

// Set of constants representing the allowable values for DocumentDetailsSourceEnum
const (
	DocumentDetailsSourceInline        DocumentDetailsSourceEnum = "INLINE"
	DocumentDetailsSourceObjectStorage DocumentDetailsSourceEnum = "OBJECT_STORAGE"
)

var mappingDocumentDetailsSourceEnum = map[string]DocumentDetailsSourceEnum{
	"INLINE":         DocumentDetailsSourceInline,
	"OBJECT_STORAGE": DocumentDetailsSourceObjectStorage,
}

// GetDocumentDetailsSourceEnumValues Enumerates the set of values for DocumentDetailsSourceEnum
func GetDocumentDetailsSourceEnumValues() []DocumentDetailsSourceEnum {
	values := make([]DocumentDetailsSourceEnum, 0)
	for _, v := range mappingDocumentDetailsSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetDocumentDetailsSourceEnumStringValues Enumerates the set of values in String for DocumentDetailsSourceEnum
func GetDocumentDetailsSourceEnumStringValues() []string {
	return []string{
		"INLINE",
		"OBJECT_STORAGE",
	}
}

// GetMappingDocumentDetailsSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDocumentDetailsSourceEnum(val string) (DocumentDetailsSourceEnum, bool) {
	mappingDocumentDetailsSourceEnumIgnoreCase := make(map[string]DocumentDetailsSourceEnum)
	for k, v := range mappingDocumentDetailsSourceEnum {
		mappingDocumentDetailsSourceEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDocumentDetailsSourceEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
