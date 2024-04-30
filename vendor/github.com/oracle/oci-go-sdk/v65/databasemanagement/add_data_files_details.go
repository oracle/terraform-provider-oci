// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AddDataFilesDetails The details required to add data files or temp files to the tablespace.
// lease provide either credentialDetails or databaseCredential. It's recommended to provide databaseCredential
type AddDataFilesDetails struct {

	// Specifies whether the file is a data file or temp file.
	FileType AddDataFilesDetailsFileTypeEnum `mandatory:"true" json:"fileType"`

	CredentialDetails TablespaceAdminCredentialDetails `mandatory:"false" json:"credentialDetails"`

	DatabaseCredential DatabaseCredentialDetails `mandatory:"false" json:"databaseCredential"`

	// The list of data files or temp files added to the tablespace.
	DataFiles []string `mandatory:"false" json:"dataFiles"`

	// The number of data files or temp files to be added for the tablespace. This is for Oracle Managed Files only.
	FileCount *int `mandatory:"false" json:"fileCount"`

	// The size of each data file or temp file.
	FileSize *TablespaceStorageSize `mandatory:"false" json:"fileSize"`

	// Specifies whether Oracle can reuse the data file or temp file. Reuse is only allowed when the file name is provided.
	IsReusable *bool `mandatory:"false" json:"isReusable"`

	// Specifies whether the data file or temp file can be extended automatically.
	IsAutoExtensible *bool `mandatory:"false" json:"isAutoExtensible"`

	// The size of the next increment of disk space to be allocated automatically when more extents are required.
	AutoExtendNextSize *TablespaceStorageSize `mandatory:"false" json:"autoExtendNextSize"`

	// The maximum disk space allowed for automatic extension of the data files or temp files.
	AutoExtendMaxSize *TablespaceStorageSize `mandatory:"false" json:"autoExtendMaxSize"`

	// Specifies whether the disk space of the data file or temp file can be limited.
	IsMaxSizeUnlimited *bool `mandatory:"false" json:"isMaxSizeUnlimited"`
}

func (m AddDataFilesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddDataFilesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAddDataFilesDetailsFileTypeEnum(string(m.FileType)); !ok && m.FileType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FileType: %s. Supported values are: %s.", m.FileType, strings.Join(GetAddDataFilesDetailsFileTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *AddDataFilesDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CredentialDetails  tablespaceadmincredentialdetails `json:"credentialDetails"`
		DatabaseCredential databasecredentialdetails        `json:"databaseCredential"`
		DataFiles          []string                         `json:"dataFiles"`
		FileCount          *int                             `json:"fileCount"`
		FileSize           *TablespaceStorageSize           `json:"fileSize"`
		IsReusable         *bool                            `json:"isReusable"`
		IsAutoExtensible   *bool                            `json:"isAutoExtensible"`
		AutoExtendNextSize *TablespaceStorageSize           `json:"autoExtendNextSize"`
		AutoExtendMaxSize  *TablespaceStorageSize           `json:"autoExtendMaxSize"`
		IsMaxSizeUnlimited *bool                            `json:"isMaxSizeUnlimited"`
		FileType           AddDataFilesDetailsFileTypeEnum  `json:"fileType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.CredentialDetails.UnmarshalPolymorphicJSON(model.CredentialDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CredentialDetails = nn.(TablespaceAdminCredentialDetails)
	} else {
		m.CredentialDetails = nil
	}

	nn, e = model.DatabaseCredential.UnmarshalPolymorphicJSON(model.DatabaseCredential.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DatabaseCredential = nn.(DatabaseCredentialDetails)
	} else {
		m.DatabaseCredential = nil
	}

	m.DataFiles = make([]string, len(model.DataFiles))
	copy(m.DataFiles, model.DataFiles)
	m.FileCount = model.FileCount

	m.FileSize = model.FileSize

	m.IsReusable = model.IsReusable

	m.IsAutoExtensible = model.IsAutoExtensible

	m.AutoExtendNextSize = model.AutoExtendNextSize

	m.AutoExtendMaxSize = model.AutoExtendMaxSize

	m.IsMaxSizeUnlimited = model.IsMaxSizeUnlimited

	m.FileType = model.FileType

	return
}

// AddDataFilesDetailsFileTypeEnum Enum with underlying type: string
type AddDataFilesDetailsFileTypeEnum string

// Set of constants representing the allowable values for AddDataFilesDetailsFileTypeEnum
const (
	AddDataFilesDetailsFileTypeDatafile AddDataFilesDetailsFileTypeEnum = "DATAFILE"
	AddDataFilesDetailsFileTypeTempfile AddDataFilesDetailsFileTypeEnum = "TEMPFILE"
)

var mappingAddDataFilesDetailsFileTypeEnum = map[string]AddDataFilesDetailsFileTypeEnum{
	"DATAFILE": AddDataFilesDetailsFileTypeDatafile,
	"TEMPFILE": AddDataFilesDetailsFileTypeTempfile,
}

var mappingAddDataFilesDetailsFileTypeEnumLowerCase = map[string]AddDataFilesDetailsFileTypeEnum{
	"datafile": AddDataFilesDetailsFileTypeDatafile,
	"tempfile": AddDataFilesDetailsFileTypeTempfile,
}

// GetAddDataFilesDetailsFileTypeEnumValues Enumerates the set of values for AddDataFilesDetailsFileTypeEnum
func GetAddDataFilesDetailsFileTypeEnumValues() []AddDataFilesDetailsFileTypeEnum {
	values := make([]AddDataFilesDetailsFileTypeEnum, 0)
	for _, v := range mappingAddDataFilesDetailsFileTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAddDataFilesDetailsFileTypeEnumStringValues Enumerates the set of values in String for AddDataFilesDetailsFileTypeEnum
func GetAddDataFilesDetailsFileTypeEnumStringValues() []string {
	return []string{
		"DATAFILE",
		"TEMPFILE",
	}
}

// GetMappingAddDataFilesDetailsFileTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddDataFilesDetailsFileTypeEnum(val string) (AddDataFilesDetailsFileTypeEnum, bool) {
	enum, ok := mappingAddDataFilesDetailsFileTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
