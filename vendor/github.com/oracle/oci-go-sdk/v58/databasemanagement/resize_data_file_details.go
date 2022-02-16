// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ResizeDataFileDetails The details required to resize a data file or temp file within the tablespace.
type ResizeDataFileDetails struct {
	CredentialDetails TablespaceAdminCredentialDetails `mandatory:"true" json:"credentialDetails"`

	// Specifies whether the file is a data file or temp file.
	FileType ResizeDataFileDetailsFileTypeEnum `mandatory:"true" json:"fileType"`

	// Name of the data file or temp file to be resized.
	DataFile *string `mandatory:"true" json:"dataFile"`

	// The new size of the data file or temp file.
	FileSize *TablespaceStorageSize `mandatory:"false" json:"fileSize"`

	// Specifies whether the data file or temp file can be extended automatically.
	IsAutoExtensible *bool `mandatory:"false" json:"isAutoExtensible"`

	// The size of the next increment of disk space to be allocated automatically when more extents are required.
	AutoExtendNextSize *TablespaceStorageSize `mandatory:"false" json:"autoExtendNextSize"`

	// The maximum disk space allowed for automatic extension of the data files or temp files.
	AutoExtendMaxSize *TablespaceStorageSize `mandatory:"false" json:"autoExtendMaxSize"`

	// Specifies whether the disk space of the data file or temp file can be limited.
	IsMaxSizeUnlimited *bool `mandatory:"false" json:"isMaxSizeUnlimited"`
}

func (m ResizeDataFileDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResizeDataFileDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResizeDataFileDetailsFileTypeEnum(string(m.FileType)); !ok && m.FileType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FileType: %s. Supported values are: %s.", m.FileType, strings.Join(GetResizeDataFileDetailsFileTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ResizeDataFileDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FileSize           *TablespaceStorageSize            `json:"fileSize"`
		IsAutoExtensible   *bool                             `json:"isAutoExtensible"`
		AutoExtendNextSize *TablespaceStorageSize            `json:"autoExtendNextSize"`
		AutoExtendMaxSize  *TablespaceStorageSize            `json:"autoExtendMaxSize"`
		IsMaxSizeUnlimited *bool                             `json:"isMaxSizeUnlimited"`
		CredentialDetails  tablespaceadmincredentialdetails  `json:"credentialDetails"`
		FileType           ResizeDataFileDetailsFileTypeEnum `json:"fileType"`
		DataFile           *string                           `json:"dataFile"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.FileSize = model.FileSize

	m.IsAutoExtensible = model.IsAutoExtensible

	m.AutoExtendNextSize = model.AutoExtendNextSize

	m.AutoExtendMaxSize = model.AutoExtendMaxSize

	m.IsMaxSizeUnlimited = model.IsMaxSizeUnlimited

	nn, e = model.CredentialDetails.UnmarshalPolymorphicJSON(model.CredentialDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CredentialDetails = nn.(TablespaceAdminCredentialDetails)
	} else {
		m.CredentialDetails = nil
	}

	m.FileType = model.FileType

	m.DataFile = model.DataFile

	return
}

// ResizeDataFileDetailsFileTypeEnum Enum with underlying type: string
type ResizeDataFileDetailsFileTypeEnum string

// Set of constants representing the allowable values for ResizeDataFileDetailsFileTypeEnum
const (
	ResizeDataFileDetailsFileTypeDatafile ResizeDataFileDetailsFileTypeEnum = "DATAFILE"
	ResizeDataFileDetailsFileTypeTempfile ResizeDataFileDetailsFileTypeEnum = "TEMPFILE"
)

var mappingResizeDataFileDetailsFileTypeEnum = map[string]ResizeDataFileDetailsFileTypeEnum{
	"DATAFILE": ResizeDataFileDetailsFileTypeDatafile,
	"TEMPFILE": ResizeDataFileDetailsFileTypeTempfile,
}

// GetResizeDataFileDetailsFileTypeEnumValues Enumerates the set of values for ResizeDataFileDetailsFileTypeEnum
func GetResizeDataFileDetailsFileTypeEnumValues() []ResizeDataFileDetailsFileTypeEnum {
	values := make([]ResizeDataFileDetailsFileTypeEnum, 0)
	for _, v := range mappingResizeDataFileDetailsFileTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResizeDataFileDetailsFileTypeEnumStringValues Enumerates the set of values in String for ResizeDataFileDetailsFileTypeEnum
func GetResizeDataFileDetailsFileTypeEnumStringValues() []string {
	return []string{
		"DATAFILE",
		"TEMPFILE",
	}
}

// GetMappingResizeDataFileDetailsFileTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResizeDataFileDetailsFileTypeEnum(val string) (ResizeDataFileDetailsFileTypeEnum, bool) {
	mappingResizeDataFileDetailsFileTypeEnumIgnoreCase := make(map[string]ResizeDataFileDetailsFileTypeEnum)
	for k, v := range mappingResizeDataFileDetailsFileTypeEnum {
		mappingResizeDataFileDetailsFileTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingResizeDataFileDetailsFileTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
