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

// UpdateArchiveSourceDetails Source details for updating an archive‑based function.
type UpdateArchiveSourceDetails interface {
}

type updatearchivesourcedetails struct {
	JsonData          []byte
	ArchiveSourceType string `json:"archiveSourceType"`
}

// UnmarshalJSON unmarshals json
func (m *updatearchivesourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatearchivesourcedetails updatearchivesourcedetails
	s := struct {
		Model Unmarshalerupdatearchivesourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ArchiveSourceType = s.Model.ArchiveSourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatearchivesourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ArchiveSourceType {
	case "DIRECT_ARCHIVE":
		mm := UpdateDirectArchiveSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE_ARCHIVE":
		mm := UpdateObjectStorageArchiveSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateArchiveSourceDetails: %s.", m.ArchiveSourceType)
		return *m, nil
	}
}

func (m updatearchivesourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatearchivesourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateArchiveSourceDetailsArchiveSourceTypeEnum Enum with underlying type: string
type UpdateArchiveSourceDetailsArchiveSourceTypeEnum string

// Set of constants representing the allowable values for UpdateArchiveSourceDetailsArchiveSourceTypeEnum
const (
	UpdateArchiveSourceDetailsArchiveSourceTypeObjectStorageArchive UpdateArchiveSourceDetailsArchiveSourceTypeEnum = "OBJECT_STORAGE_ARCHIVE"
	UpdateArchiveSourceDetailsArchiveSourceTypeDirectArchive        UpdateArchiveSourceDetailsArchiveSourceTypeEnum = "DIRECT_ARCHIVE"
)

var mappingUpdateArchiveSourceDetailsArchiveSourceTypeEnum = map[string]UpdateArchiveSourceDetailsArchiveSourceTypeEnum{
	"OBJECT_STORAGE_ARCHIVE": UpdateArchiveSourceDetailsArchiveSourceTypeObjectStorageArchive,
	"DIRECT_ARCHIVE":         UpdateArchiveSourceDetailsArchiveSourceTypeDirectArchive,
}

var mappingUpdateArchiveSourceDetailsArchiveSourceTypeEnumLowerCase = map[string]UpdateArchiveSourceDetailsArchiveSourceTypeEnum{
	"object_storage_archive": UpdateArchiveSourceDetailsArchiveSourceTypeObjectStorageArchive,
	"direct_archive":         UpdateArchiveSourceDetailsArchiveSourceTypeDirectArchive,
}

// GetUpdateArchiveSourceDetailsArchiveSourceTypeEnumValues Enumerates the set of values for UpdateArchiveSourceDetailsArchiveSourceTypeEnum
func GetUpdateArchiveSourceDetailsArchiveSourceTypeEnumValues() []UpdateArchiveSourceDetailsArchiveSourceTypeEnum {
	values := make([]UpdateArchiveSourceDetailsArchiveSourceTypeEnum, 0)
	for _, v := range mappingUpdateArchiveSourceDetailsArchiveSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateArchiveSourceDetailsArchiveSourceTypeEnumStringValues Enumerates the set of values in String for UpdateArchiveSourceDetailsArchiveSourceTypeEnum
func GetUpdateArchiveSourceDetailsArchiveSourceTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE_ARCHIVE",
		"DIRECT_ARCHIVE",
	}
}

// GetMappingUpdateArchiveSourceDetailsArchiveSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateArchiveSourceDetailsArchiveSourceTypeEnum(val string) (UpdateArchiveSourceDetailsArchiveSourceTypeEnum, bool) {
	enum, ok := mappingUpdateArchiveSourceDetailsArchiveSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
