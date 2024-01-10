// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateTablespaceDetails The details required to update a tablespace.
type UpdateTablespaceDetails struct {
	CredentialDetails TablespaceAdminCredentialDetails `mandatory:"true" json:"credentialDetails"`

	// The name of the tablespace. It must be unique within a database.
	Name *string `mandatory:"false" json:"name"`

	// The type of tablespace.
	Type UpdateTablespaceDetailsTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The size of each data file or temp file.
	FileSize *TablespaceStorageSize `mandatory:"false" json:"fileSize"`

	// The status of the tablespace.
	Status UpdateTablespaceDetailsStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Specifies whether the data file or temp file can be extended automatically.
	IsAutoExtensible *bool `mandatory:"false" json:"isAutoExtensible"`

	// The size of the next increment of disk space to be allocated automatically when more extents are required.
	AutoExtendNextSize *TablespaceStorageSize `mandatory:"false" json:"autoExtendNextSize"`

	// The maximum disk space allowed for automatic extension of the data files or temp files.
	AutoExtendMaxSize *TablespaceStorageSize `mandatory:"false" json:"autoExtendMaxSize"`

	// Specifies whether the disk space of the data file or temp file can be limited.
	IsMaxSizeUnlimited *bool `mandatory:"false" json:"isMaxSizeUnlimited"`

	// Specifies whether the tablespace is the default tablespace.
	IsDefault *bool `mandatory:"false" json:"isDefault"`
}

func (m UpdateTablespaceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateTablespaceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateTablespaceDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetUpdateTablespaceDetailsTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateTablespaceDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetUpdateTablespaceDetailsStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateTablespaceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Name               *string                           `json:"name"`
		Type               UpdateTablespaceDetailsTypeEnum   `json:"type"`
		FileSize           *TablespaceStorageSize            `json:"fileSize"`
		Status             UpdateTablespaceDetailsStatusEnum `json:"status"`
		IsAutoExtensible   *bool                             `json:"isAutoExtensible"`
		AutoExtendNextSize *TablespaceStorageSize            `json:"autoExtendNextSize"`
		AutoExtendMaxSize  *TablespaceStorageSize            `json:"autoExtendMaxSize"`
		IsMaxSizeUnlimited *bool                             `json:"isMaxSizeUnlimited"`
		IsDefault          *bool                             `json:"isDefault"`
		CredentialDetails  tablespaceadmincredentialdetails  `json:"credentialDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Name = model.Name

	m.Type = model.Type

	m.FileSize = model.FileSize

	m.Status = model.Status

	m.IsAutoExtensible = model.IsAutoExtensible

	m.AutoExtendNextSize = model.AutoExtendNextSize

	m.AutoExtendMaxSize = model.AutoExtendMaxSize

	m.IsMaxSizeUnlimited = model.IsMaxSizeUnlimited

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

	return
}

// UpdateTablespaceDetailsTypeEnum Enum with underlying type: string
type UpdateTablespaceDetailsTypeEnum string

// Set of constants representing the allowable values for UpdateTablespaceDetailsTypeEnum
const (
	UpdateTablespaceDetailsTypePermanent UpdateTablespaceDetailsTypeEnum = "PERMANENT"
	UpdateTablespaceDetailsTypeTemporary UpdateTablespaceDetailsTypeEnum = "TEMPORARY"
)

var mappingUpdateTablespaceDetailsTypeEnum = map[string]UpdateTablespaceDetailsTypeEnum{
	"PERMANENT": UpdateTablespaceDetailsTypePermanent,
	"TEMPORARY": UpdateTablespaceDetailsTypeTemporary,
}

var mappingUpdateTablespaceDetailsTypeEnumLowerCase = map[string]UpdateTablespaceDetailsTypeEnum{
	"permanent": UpdateTablespaceDetailsTypePermanent,
	"temporary": UpdateTablespaceDetailsTypeTemporary,
}

// GetUpdateTablespaceDetailsTypeEnumValues Enumerates the set of values for UpdateTablespaceDetailsTypeEnum
func GetUpdateTablespaceDetailsTypeEnumValues() []UpdateTablespaceDetailsTypeEnum {
	values := make([]UpdateTablespaceDetailsTypeEnum, 0)
	for _, v := range mappingUpdateTablespaceDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateTablespaceDetailsTypeEnumStringValues Enumerates the set of values in String for UpdateTablespaceDetailsTypeEnum
func GetUpdateTablespaceDetailsTypeEnumStringValues() []string {
	return []string{
		"PERMANENT",
		"TEMPORARY",
	}
}

// GetMappingUpdateTablespaceDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateTablespaceDetailsTypeEnum(val string) (UpdateTablespaceDetailsTypeEnum, bool) {
	enum, ok := mappingUpdateTablespaceDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateTablespaceDetailsStatusEnum Enum with underlying type: string
type UpdateTablespaceDetailsStatusEnum string

// Set of constants representing the allowable values for UpdateTablespaceDetailsStatusEnum
const (
	UpdateTablespaceDetailsStatusOnly  UpdateTablespaceDetailsStatusEnum = "READ_ONLY"
	UpdateTablespaceDetailsStatusWrite UpdateTablespaceDetailsStatusEnum = "READ_WRITE"
)

var mappingUpdateTablespaceDetailsStatusEnum = map[string]UpdateTablespaceDetailsStatusEnum{
	"READ_ONLY":  UpdateTablespaceDetailsStatusOnly,
	"READ_WRITE": UpdateTablespaceDetailsStatusWrite,
}

var mappingUpdateTablespaceDetailsStatusEnumLowerCase = map[string]UpdateTablespaceDetailsStatusEnum{
	"read_only":  UpdateTablespaceDetailsStatusOnly,
	"read_write": UpdateTablespaceDetailsStatusWrite,
}

// GetUpdateTablespaceDetailsStatusEnumValues Enumerates the set of values for UpdateTablespaceDetailsStatusEnum
func GetUpdateTablespaceDetailsStatusEnumValues() []UpdateTablespaceDetailsStatusEnum {
	values := make([]UpdateTablespaceDetailsStatusEnum, 0)
	for _, v := range mappingUpdateTablespaceDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateTablespaceDetailsStatusEnumStringValues Enumerates the set of values in String for UpdateTablespaceDetailsStatusEnum
func GetUpdateTablespaceDetailsStatusEnumStringValues() []string {
	return []string{
		"READ_ONLY",
		"READ_WRITE",
	}
}

// GetMappingUpdateTablespaceDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateTablespaceDetailsStatusEnum(val string) (UpdateTablespaceDetailsStatusEnum, bool) {
	enum, ok := mappingUpdateTablespaceDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
