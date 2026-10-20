// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ArchiveSourceDetails The details for the Archive source of the Function.
// This mode is used when the function code is provided as an archive, either from Object Storage or directly uploaded by the API caller.
// It is suitable for scenarios where the function code is packaged as a single archive file.
type ArchiveSourceDetails interface {
}

type archivesourcedetails struct {
	JsonData          []byte
	ArchiveSourceType string `json:"archiveSourceType"`
}

// UnmarshalJSON unmarshals json
func (m *archivesourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerarchivesourcedetails archivesourcedetails
	s := struct {
		Model Unmarshalerarchivesourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ArchiveSourceType = s.Model.ArchiveSourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *archivesourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ArchiveSourceType {
	case "DIRECT_ARCHIVE":
		mm := DirectArchiveSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE_ARCHIVE":
		mm := ObjectStorageArchiveSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ArchiveSourceDetails: %s.", m.ArchiveSourceType)
		return *m, nil
	}
}

func (m archivesourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m archivesourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ArchiveSourceDetailsArchiveSourceTypeEnum Enum with underlying type: string
type ArchiveSourceDetailsArchiveSourceTypeEnum string

// Set of constants representing the allowable values for ArchiveSourceDetailsArchiveSourceTypeEnum
const (
	ArchiveSourceDetailsArchiveSourceTypeObjectStorageArchive ArchiveSourceDetailsArchiveSourceTypeEnum = "OBJECT_STORAGE_ARCHIVE"
	ArchiveSourceDetailsArchiveSourceTypeDirectArchive        ArchiveSourceDetailsArchiveSourceTypeEnum = "DIRECT_ARCHIVE"
)

var mappingArchiveSourceDetailsArchiveSourceTypeEnum = map[string]ArchiveSourceDetailsArchiveSourceTypeEnum{
	"OBJECT_STORAGE_ARCHIVE": ArchiveSourceDetailsArchiveSourceTypeObjectStorageArchive,
	"DIRECT_ARCHIVE":         ArchiveSourceDetailsArchiveSourceTypeDirectArchive,
}

var mappingArchiveSourceDetailsArchiveSourceTypeEnumLowerCase = map[string]ArchiveSourceDetailsArchiveSourceTypeEnum{
	"object_storage_archive": ArchiveSourceDetailsArchiveSourceTypeObjectStorageArchive,
	"direct_archive":         ArchiveSourceDetailsArchiveSourceTypeDirectArchive,
}

// GetArchiveSourceDetailsArchiveSourceTypeEnumValues Enumerates the set of values for ArchiveSourceDetailsArchiveSourceTypeEnum
func GetArchiveSourceDetailsArchiveSourceTypeEnumValues() []ArchiveSourceDetailsArchiveSourceTypeEnum {
	values := make([]ArchiveSourceDetailsArchiveSourceTypeEnum, 0)
	for _, v := range mappingArchiveSourceDetailsArchiveSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetArchiveSourceDetailsArchiveSourceTypeEnumStringValues Enumerates the set of values in String for ArchiveSourceDetailsArchiveSourceTypeEnum
func GetArchiveSourceDetailsArchiveSourceTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE_ARCHIVE",
		"DIRECT_ARCHIVE",
	}
}

// GetMappingArchiveSourceDetailsArchiveSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingArchiveSourceDetailsArchiveSourceTypeEnum(val string) (ArchiveSourceDetailsArchiveSourceTypeEnum, bool) {
	enum, ok := mappingArchiveSourceDetailsArchiveSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
