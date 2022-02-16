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

// RemoveDataFileDetails The details required to remove a data file or temp file from the tablespace.
type RemoveDataFileDetails struct {
	CredentialDetails TablespaceAdminCredentialDetails `mandatory:"true" json:"credentialDetails"`

	// Specifies whether the file is a data file or temp file.
	FileType RemoveDataFileDetailsFileTypeEnum `mandatory:"true" json:"fileType"`

	// Name of the data file or temp file to be removed from the tablespace.
	DataFile *string `mandatory:"true" json:"dataFile"`
}

func (m RemoveDataFileDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RemoveDataFileDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRemoveDataFileDetailsFileTypeEnum(string(m.FileType)); !ok && m.FileType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FileType: %s. Supported values are: %s.", m.FileType, strings.Join(GetRemoveDataFileDetailsFileTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *RemoveDataFileDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CredentialDetails tablespaceadmincredentialdetails  `json:"credentialDetails"`
		FileType          RemoveDataFileDetailsFileTypeEnum `json:"fileType"`
		DataFile          *string                           `json:"dataFile"`
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

	m.FileType = model.FileType

	m.DataFile = model.DataFile

	return
}

// RemoveDataFileDetailsFileTypeEnum Enum with underlying type: string
type RemoveDataFileDetailsFileTypeEnum string

// Set of constants representing the allowable values for RemoveDataFileDetailsFileTypeEnum
const (
	RemoveDataFileDetailsFileTypeDatafile RemoveDataFileDetailsFileTypeEnum = "DATAFILE"
	RemoveDataFileDetailsFileTypeTempfile RemoveDataFileDetailsFileTypeEnum = "TEMPFILE"
)

var mappingRemoveDataFileDetailsFileTypeEnum = map[string]RemoveDataFileDetailsFileTypeEnum{
	"DATAFILE": RemoveDataFileDetailsFileTypeDatafile,
	"TEMPFILE": RemoveDataFileDetailsFileTypeTempfile,
}

// GetRemoveDataFileDetailsFileTypeEnumValues Enumerates the set of values for RemoveDataFileDetailsFileTypeEnum
func GetRemoveDataFileDetailsFileTypeEnumValues() []RemoveDataFileDetailsFileTypeEnum {
	values := make([]RemoveDataFileDetailsFileTypeEnum, 0)
	for _, v := range mappingRemoveDataFileDetailsFileTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRemoveDataFileDetailsFileTypeEnumStringValues Enumerates the set of values in String for RemoveDataFileDetailsFileTypeEnum
func GetRemoveDataFileDetailsFileTypeEnumStringValues() []string {
	return []string{
		"DATAFILE",
		"TEMPFILE",
	}
}

// GetMappingRemoveDataFileDetailsFileTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRemoveDataFileDetailsFileTypeEnum(val string) (RemoveDataFileDetailsFileTypeEnum, bool) {
	mappingRemoveDataFileDetailsFileTypeEnumIgnoreCase := make(map[string]RemoveDataFileDetailsFileTypeEnum)
	for k, v := range mappingRemoveDataFileDetailsFileTypeEnum {
		mappingRemoveDataFileDetailsFileTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingRemoveDataFileDetailsFileTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
