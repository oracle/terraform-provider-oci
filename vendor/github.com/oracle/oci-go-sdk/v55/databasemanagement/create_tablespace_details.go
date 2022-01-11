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
	"github.com/oracle/oci-go-sdk/v55/common"
	"strings"
)

// CreateTablespaceDetails The details required to create a tablespace.
type CreateTablespaceDetails struct {
	CredentialDetails TablespaceAdminCredentialDetails `mandatory:"true" json:"credentialDetails"`

	// The name of the tablespace. It must be unique within a database.
	Name *string `mandatory:"true" json:"name"`

	// The type of tablespace.
	Type CreateTablespaceDetailsTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Specifies whether the tablespace is a bigfile or smallfile tablespace.
	// A bigfile tablespace contains only one data file or temp file, which can contain up to approximately 4 billion (232) blocks.
	// A smallfile tablespace is a traditional Oracle tablespace, which can contain 1022 data files or temp files, each of which can contain up to approximately 4 million (222) blocks.
	IsBigfile *bool `mandatory:"false" json:"isBigfile"`

	// The list of data files or temp files created for the tablespace.
	DataFiles []string `mandatory:"false" json:"dataFiles"`

	// The number of data files or temp files created for the tablespace. This is for Oracle Managed Files only.
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

	// Block size for the tablespace.
	BlockSizeInKilobytes *int `mandatory:"false" json:"blockSizeInKilobytes"`

	// Indicates whether the tablespace is encrypted.
	IsEncrypted *bool `mandatory:"false" json:"isEncrypted"`

	// The name of the encryption algorithm to be used for tablespace encryption.
	EncryptionAlgorithm *string `mandatory:"false" json:"encryptionAlgorithm"`

	// The default compression of data for all tables created in the tablespace.
	DefaultCompress CreateTablespaceDetailsDefaultCompressEnum `mandatory:"false" json:"defaultCompress,omitempty"`

	// The status of the tablespace.
	Status CreateTablespaceDetailsStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Specifies how the extents of the tablespace should be managed.
	ExtentManagement CreateTablespaceDetailsExtentManagementEnum `mandatory:"false" json:"extentManagement,omitempty"`

	// The size of the extent when the tablespace is managed with uniform extents of a specific size.
	ExtentUniformSize *TablespaceStorageSize `mandatory:"false" json:"extentUniformSize"`

	// Specifies whether tablespace segment management should be automatic or manual.
	SegmentManagement CreateTablespaceDetailsSegmentManagementEnum `mandatory:"false" json:"segmentManagement,omitempty"`

	// Specifies whether the tablespace is the default tablespace.
	IsDefault *bool `mandatory:"false" json:"isDefault"`
}

func (m CreateTablespaceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateTablespaceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := mappingCreateTablespaceDetailsTypeEnum[string(m.Type)]; !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCreateTablespaceDetailsTypeEnumStringValues(), ",")))
	}
	if _, ok := mappingCreateTablespaceDetailsDefaultCompressEnum[string(m.DefaultCompress)]; !ok && m.DefaultCompress != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultCompress: %s. Supported values are: %s.", m.DefaultCompress, strings.Join(GetCreateTablespaceDetailsDefaultCompressEnumStringValues(), ",")))
	}
	if _, ok := mappingCreateTablespaceDetailsStatusEnum[string(m.Status)]; !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetCreateTablespaceDetailsStatusEnumStringValues(), ",")))
	}
	if _, ok := mappingCreateTablespaceDetailsExtentManagementEnum[string(m.ExtentManagement)]; !ok && m.ExtentManagement != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExtentManagement: %s. Supported values are: %s.", m.ExtentManagement, strings.Join(GetCreateTablespaceDetailsExtentManagementEnumStringValues(), ",")))
	}
	if _, ok := mappingCreateTablespaceDetailsSegmentManagementEnum[string(m.SegmentManagement)]; !ok && m.SegmentManagement != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SegmentManagement: %s. Supported values are: %s.", m.SegmentManagement, strings.Join(GetCreateTablespaceDetailsSegmentManagementEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateTablespaceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Type                 CreateTablespaceDetailsTypeEnum              `json:"type"`
		IsBigfile            *bool                                        `json:"isBigfile"`
		DataFiles            []string                                     `json:"dataFiles"`
		FileCount            *int                                         `json:"fileCount"`
		FileSize             *TablespaceStorageSize                       `json:"fileSize"`
		IsReusable           *bool                                        `json:"isReusable"`
		IsAutoExtensible     *bool                                        `json:"isAutoExtensible"`
		AutoExtendNextSize   *TablespaceStorageSize                       `json:"autoExtendNextSize"`
		AutoExtendMaxSize    *TablespaceStorageSize                       `json:"autoExtendMaxSize"`
		IsMaxSizeUnlimited   *bool                                        `json:"isMaxSizeUnlimited"`
		BlockSizeInKilobytes *int                                         `json:"blockSizeInKilobytes"`
		IsEncrypted          *bool                                        `json:"isEncrypted"`
		EncryptionAlgorithm  *string                                      `json:"encryptionAlgorithm"`
		DefaultCompress      CreateTablespaceDetailsDefaultCompressEnum   `json:"defaultCompress"`
		Status               CreateTablespaceDetailsStatusEnum            `json:"status"`
		ExtentManagement     CreateTablespaceDetailsExtentManagementEnum  `json:"extentManagement"`
		ExtentUniformSize    *TablespaceStorageSize                       `json:"extentUniformSize"`
		SegmentManagement    CreateTablespaceDetailsSegmentManagementEnum `json:"segmentManagement"`
		IsDefault            *bool                                        `json:"isDefault"`
		CredentialDetails    tablespaceadmincredentialdetails             `json:"credentialDetails"`
		Name                 *string                                      `json:"name"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Type = model.Type

	m.IsBigfile = model.IsBigfile

	m.DataFiles = make([]string, len(model.DataFiles))
	for i, n := range model.DataFiles {
		m.DataFiles[i] = n
	}

	m.FileCount = model.FileCount

	m.FileSize = model.FileSize

	m.IsReusable = model.IsReusable

	m.IsAutoExtensible = model.IsAutoExtensible

	m.AutoExtendNextSize = model.AutoExtendNextSize

	m.AutoExtendMaxSize = model.AutoExtendMaxSize

	m.IsMaxSizeUnlimited = model.IsMaxSizeUnlimited

	m.BlockSizeInKilobytes = model.BlockSizeInKilobytes

	m.IsEncrypted = model.IsEncrypted

	m.EncryptionAlgorithm = model.EncryptionAlgorithm

	m.DefaultCompress = model.DefaultCompress

	m.Status = model.Status

	m.ExtentManagement = model.ExtentManagement

	m.ExtentUniformSize = model.ExtentUniformSize

	m.SegmentManagement = model.SegmentManagement

	m.IsDefault = model.IsDefault

	nn, e = model.CredentialDetails.UnmarshalPolymorphicJSON(model.CredentialDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CredentialDetails = nn.(TablespaceAdminCredentialDetails)
	} else {
		m.CredentialDetails = nil
	}

	m.Name = model.Name

	return
}

// CreateTablespaceDetailsTypeEnum Enum with underlying type: string
type CreateTablespaceDetailsTypeEnum string

// Set of constants representing the allowable values for CreateTablespaceDetailsTypeEnum
const (
	CreateTablespaceDetailsTypePermanent CreateTablespaceDetailsTypeEnum = "PERMANENT"
	CreateTablespaceDetailsTypeTemporary CreateTablespaceDetailsTypeEnum = "TEMPORARY"
)

var mappingCreateTablespaceDetailsTypeEnum = map[string]CreateTablespaceDetailsTypeEnum{
	"PERMANENT": CreateTablespaceDetailsTypePermanent,
	"TEMPORARY": CreateTablespaceDetailsTypeTemporary,
}

// GetCreateTablespaceDetailsTypeEnumValues Enumerates the set of values for CreateTablespaceDetailsTypeEnum
func GetCreateTablespaceDetailsTypeEnumValues() []CreateTablespaceDetailsTypeEnum {
	values := make([]CreateTablespaceDetailsTypeEnum, 0)
	for _, v := range mappingCreateTablespaceDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateTablespaceDetailsTypeEnumStringValues Enumerates the set of values in String for CreateTablespaceDetailsTypeEnum
func GetCreateTablespaceDetailsTypeEnumStringValues() []string {
	return []string{
		"PERMANENT",
		"TEMPORARY",
	}
}

// CreateTablespaceDetailsDefaultCompressEnum Enum with underlying type: string
type CreateTablespaceDetailsDefaultCompressEnum string

// Set of constants representing the allowable values for CreateTablespaceDetailsDefaultCompressEnum
const (
	CreateTablespaceDetailsDefaultCompressNoCompress    CreateTablespaceDetailsDefaultCompressEnum = "NO_COMPRESS"
	CreateTablespaceDetailsDefaultCompressBasicCompress CreateTablespaceDetailsDefaultCompressEnum = "BASIC_COMPRESS"
)

var mappingCreateTablespaceDetailsDefaultCompressEnum = map[string]CreateTablespaceDetailsDefaultCompressEnum{
	"NO_COMPRESS":    CreateTablespaceDetailsDefaultCompressNoCompress,
	"BASIC_COMPRESS": CreateTablespaceDetailsDefaultCompressBasicCompress,
}

// GetCreateTablespaceDetailsDefaultCompressEnumValues Enumerates the set of values for CreateTablespaceDetailsDefaultCompressEnum
func GetCreateTablespaceDetailsDefaultCompressEnumValues() []CreateTablespaceDetailsDefaultCompressEnum {
	values := make([]CreateTablespaceDetailsDefaultCompressEnum, 0)
	for _, v := range mappingCreateTablespaceDetailsDefaultCompressEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateTablespaceDetailsDefaultCompressEnumStringValues Enumerates the set of values in String for CreateTablespaceDetailsDefaultCompressEnum
func GetCreateTablespaceDetailsDefaultCompressEnumStringValues() []string {
	return []string{
		"NO_COMPRESS",
		"BASIC_COMPRESS",
	}
}

// CreateTablespaceDetailsStatusEnum Enum with underlying type: string
type CreateTablespaceDetailsStatusEnum string

// Set of constants representing the allowable values for CreateTablespaceDetailsStatusEnum
const (
	CreateTablespaceDetailsStatusOnly  CreateTablespaceDetailsStatusEnum = "READ_ONLY"
	CreateTablespaceDetailsStatusWrite CreateTablespaceDetailsStatusEnum = "READ_WRITE"
)

var mappingCreateTablespaceDetailsStatusEnum = map[string]CreateTablespaceDetailsStatusEnum{
	"READ_ONLY":  CreateTablespaceDetailsStatusOnly,
	"READ_WRITE": CreateTablespaceDetailsStatusWrite,
}

// GetCreateTablespaceDetailsStatusEnumValues Enumerates the set of values for CreateTablespaceDetailsStatusEnum
func GetCreateTablespaceDetailsStatusEnumValues() []CreateTablespaceDetailsStatusEnum {
	values := make([]CreateTablespaceDetailsStatusEnum, 0)
	for _, v := range mappingCreateTablespaceDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateTablespaceDetailsStatusEnumStringValues Enumerates the set of values in String for CreateTablespaceDetailsStatusEnum
func GetCreateTablespaceDetailsStatusEnumStringValues() []string {
	return []string{
		"READ_ONLY",
		"READ_WRITE",
	}
}

// CreateTablespaceDetailsExtentManagementEnum Enum with underlying type: string
type CreateTablespaceDetailsExtentManagementEnum string

// Set of constants representing the allowable values for CreateTablespaceDetailsExtentManagementEnum
const (
	CreateTablespaceDetailsExtentManagementAutoallocate CreateTablespaceDetailsExtentManagementEnum = "AUTOALLOCATE"
	CreateTablespaceDetailsExtentManagementUniform      CreateTablespaceDetailsExtentManagementEnum = "UNIFORM"
)

var mappingCreateTablespaceDetailsExtentManagementEnum = map[string]CreateTablespaceDetailsExtentManagementEnum{
	"AUTOALLOCATE": CreateTablespaceDetailsExtentManagementAutoallocate,
	"UNIFORM":      CreateTablespaceDetailsExtentManagementUniform,
}

// GetCreateTablespaceDetailsExtentManagementEnumValues Enumerates the set of values for CreateTablespaceDetailsExtentManagementEnum
func GetCreateTablespaceDetailsExtentManagementEnumValues() []CreateTablespaceDetailsExtentManagementEnum {
	values := make([]CreateTablespaceDetailsExtentManagementEnum, 0)
	for _, v := range mappingCreateTablespaceDetailsExtentManagementEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateTablespaceDetailsExtentManagementEnumStringValues Enumerates the set of values in String for CreateTablespaceDetailsExtentManagementEnum
func GetCreateTablespaceDetailsExtentManagementEnumStringValues() []string {
	return []string{
		"AUTOALLOCATE",
		"UNIFORM",
	}
}

// CreateTablespaceDetailsSegmentManagementEnum Enum with underlying type: string
type CreateTablespaceDetailsSegmentManagementEnum string

// Set of constants representing the allowable values for CreateTablespaceDetailsSegmentManagementEnum
const (
	CreateTablespaceDetailsSegmentManagementAuto   CreateTablespaceDetailsSegmentManagementEnum = "AUTO"
	CreateTablespaceDetailsSegmentManagementManual CreateTablespaceDetailsSegmentManagementEnum = "MANUAL"
)

var mappingCreateTablespaceDetailsSegmentManagementEnum = map[string]CreateTablespaceDetailsSegmentManagementEnum{
	"AUTO":   CreateTablespaceDetailsSegmentManagementAuto,
	"MANUAL": CreateTablespaceDetailsSegmentManagementManual,
}

// GetCreateTablespaceDetailsSegmentManagementEnumValues Enumerates the set of values for CreateTablespaceDetailsSegmentManagementEnum
func GetCreateTablespaceDetailsSegmentManagementEnumValues() []CreateTablespaceDetailsSegmentManagementEnum {
	values := make([]CreateTablespaceDetailsSegmentManagementEnum, 0)
	for _, v := range mappingCreateTablespaceDetailsSegmentManagementEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateTablespaceDetailsSegmentManagementEnumStringValues Enumerates the set of values in String for CreateTablespaceDetailsSegmentManagementEnum
func GetCreateTablespaceDetailsSegmentManagementEnumStringValues() []string {
	return []string{
		"AUTO",
		"MANUAL",
	}
}
