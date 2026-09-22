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

// CreateArchiveSourceDetails The details required to create an Archive-based function source.
// This mode is used when the function code is provided as an archive, either from Object Storage or directly uploaded by the API caller.
// It is suitable for scenarios where the function code is packaged as a single archive file.
type CreateArchiveSourceDetails interface {
}

type createarchivesourcedetails struct {
	JsonData          []byte
	ArchiveSourceType string `json:"archiveSourceType"`
}

// UnmarshalJSON unmarshals json
func (m *createarchivesourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatearchivesourcedetails createarchivesourcedetails
	s := struct {
		Model Unmarshalercreatearchivesourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ArchiveSourceType = s.Model.ArchiveSourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createarchivesourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ArchiveSourceType {
	case "DIRECT_ARCHIVE":
		mm := CreateDirectArchiveSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE_ARCHIVE":
		mm := CreateObjectStorageArchiveSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateArchiveSourceDetails: %s.", m.ArchiveSourceType)
		return *m, nil
	}
}

func (m createarchivesourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createarchivesourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateArchiveSourceDetailsArchiveSourceTypeEnum Enum with underlying type: string
type CreateArchiveSourceDetailsArchiveSourceTypeEnum string

// Set of constants representing the allowable values for CreateArchiveSourceDetailsArchiveSourceTypeEnum
const (
	CreateArchiveSourceDetailsArchiveSourceTypeObjectStorageArchive CreateArchiveSourceDetailsArchiveSourceTypeEnum = "OBJECT_STORAGE_ARCHIVE"
	CreateArchiveSourceDetailsArchiveSourceTypeDirectArchive        CreateArchiveSourceDetailsArchiveSourceTypeEnum = "DIRECT_ARCHIVE"
)

var mappingCreateArchiveSourceDetailsArchiveSourceTypeEnum = map[string]CreateArchiveSourceDetailsArchiveSourceTypeEnum{
	"OBJECT_STORAGE_ARCHIVE": CreateArchiveSourceDetailsArchiveSourceTypeObjectStorageArchive,
	"DIRECT_ARCHIVE":         CreateArchiveSourceDetailsArchiveSourceTypeDirectArchive,
}

var mappingCreateArchiveSourceDetailsArchiveSourceTypeEnumLowerCase = map[string]CreateArchiveSourceDetailsArchiveSourceTypeEnum{
	"object_storage_archive": CreateArchiveSourceDetailsArchiveSourceTypeObjectStorageArchive,
	"direct_archive":         CreateArchiveSourceDetailsArchiveSourceTypeDirectArchive,
}

// GetCreateArchiveSourceDetailsArchiveSourceTypeEnumValues Enumerates the set of values for CreateArchiveSourceDetailsArchiveSourceTypeEnum
func GetCreateArchiveSourceDetailsArchiveSourceTypeEnumValues() []CreateArchiveSourceDetailsArchiveSourceTypeEnum {
	values := make([]CreateArchiveSourceDetailsArchiveSourceTypeEnum, 0)
	for _, v := range mappingCreateArchiveSourceDetailsArchiveSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateArchiveSourceDetailsArchiveSourceTypeEnumStringValues Enumerates the set of values in String for CreateArchiveSourceDetailsArchiveSourceTypeEnum
func GetCreateArchiveSourceDetailsArchiveSourceTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE_ARCHIVE",
		"DIRECT_ARCHIVE",
	}
}

// GetMappingCreateArchiveSourceDetailsArchiveSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateArchiveSourceDetailsArchiveSourceTypeEnum(val string) (CreateArchiveSourceDetailsArchiveSourceTypeEnum, bool) {
	enum, ok := mappingCreateArchiveSourceDetailsArchiveSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
