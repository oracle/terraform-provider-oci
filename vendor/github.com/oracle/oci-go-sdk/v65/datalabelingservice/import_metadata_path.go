// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling Management API
//
// Use Data Labeling Management API to create, list, edit & delete datasets.
//

package datalabelingservice

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ImportMetadataPath Object storage path for the metadata file
type ImportMetadataPath interface {
}

type importmetadatapath struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *importmetadatapath) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerimportmetadatapath importmetadatapath
	s := struct {
		Model Unmarshalerimportmetadatapath
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *importmetadatapath) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "OBJECT_STORAGE":
		mm := ObjectStorageImportMetadataPath{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ImportMetadataPath: %s.", m.SourceType)
		return *m, nil
	}
}

func (m importmetadatapath) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m importmetadatapath) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ImportMetadataPathSourceTypeEnum Enum with underlying type: string
type ImportMetadataPathSourceTypeEnum string

// Set of constants representing the allowable values for ImportMetadataPathSourceTypeEnum
const (
	ImportMetadataPathSourceTypeObjectStorage ImportMetadataPathSourceTypeEnum = "OBJECT_STORAGE"
)

var mappingImportMetadataPathSourceTypeEnum = map[string]ImportMetadataPathSourceTypeEnum{
	"OBJECT_STORAGE": ImportMetadataPathSourceTypeObjectStorage,
}

var mappingImportMetadataPathSourceTypeEnumLowerCase = map[string]ImportMetadataPathSourceTypeEnum{
	"object_storage": ImportMetadataPathSourceTypeObjectStorage,
}

// GetImportMetadataPathSourceTypeEnumValues Enumerates the set of values for ImportMetadataPathSourceTypeEnum
func GetImportMetadataPathSourceTypeEnumValues() []ImportMetadataPathSourceTypeEnum {
	values := make([]ImportMetadataPathSourceTypeEnum, 0)
	for _, v := range mappingImportMetadataPathSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetImportMetadataPathSourceTypeEnumStringValues Enumerates the set of values in String for ImportMetadataPathSourceTypeEnum
func GetImportMetadataPathSourceTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
	}
}

// GetMappingImportMetadataPathSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImportMetadataPathSourceTypeEnum(val string) (ImportMetadataPathSourceTypeEnum, bool) {
	enum, ok := mappingImportMetadataPathSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
