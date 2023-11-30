// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateOracleInitialLoadSettings Optional settings for Data Pump Export and Import jobs
type CreateOracleInitialLoadSettings struct {
	DataPumpParameters *CreateDataPumpParameters `mandatory:"false" json:"dataPumpParameters"`

	TablespaceDetails CreateTargetTypeTablespaceDetails `mandatory:"false" json:"tablespaceDetails"`

	ExportDirectoryObject *CreateDirectoryObject `mandatory:"false" json:"exportDirectoryObject"`

	ImportDirectoryObject *CreateDirectoryObject `mandatory:"false" json:"importDirectoryObject"`

	// Data Pump job mode.
	JobMode CreateOracleInitialLoadSettingsJobModeEnum `mandatory:"false" json:"jobMode,omitempty"`

	// Defines remapping to be applied to objects as they are processed.
	MetadataRemaps []MetadataRemap `mandatory:"false" json:"metadataRemaps"`
}

func (m CreateOracleInitialLoadSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOracleInitialLoadSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateOracleInitialLoadSettingsJobModeEnum(string(m.JobMode)); !ok && m.JobMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JobMode: %s. Supported values are: %s.", m.JobMode, strings.Join(GetCreateOracleInitialLoadSettingsJobModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateOracleInitialLoadSettings) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DataPumpParameters    *CreateDataPumpParameters                  `json:"dataPumpParameters"`
		TablespaceDetails     createtargettypetablespacedetails          `json:"tablespaceDetails"`
		ExportDirectoryObject *CreateDirectoryObject                     `json:"exportDirectoryObject"`
		ImportDirectoryObject *CreateDirectoryObject                     `json:"importDirectoryObject"`
		JobMode               CreateOracleInitialLoadSettingsJobModeEnum `json:"jobMode"`
		MetadataRemaps        []MetadataRemap                            `json:"metadataRemaps"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DataPumpParameters = model.DataPumpParameters

	nn, e = model.TablespaceDetails.UnmarshalPolymorphicJSON(model.TablespaceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TablespaceDetails = nn.(CreateTargetTypeTablespaceDetails)
	} else {
		m.TablespaceDetails = nil
	}

	m.ExportDirectoryObject = model.ExportDirectoryObject

	m.ImportDirectoryObject = model.ImportDirectoryObject

	m.JobMode = model.JobMode

	m.MetadataRemaps = make([]MetadataRemap, len(model.MetadataRemaps))
	copy(m.MetadataRemaps, model.MetadataRemaps)
	return
}

// CreateOracleInitialLoadSettingsJobModeEnum Enum with underlying type: string
type CreateOracleInitialLoadSettingsJobModeEnum string

// Set of constants representing the allowable values for CreateOracleInitialLoadSettingsJobModeEnum
const (
	CreateOracleInitialLoadSettingsJobModeFull          CreateOracleInitialLoadSettingsJobModeEnum = "FULL"
	CreateOracleInitialLoadSettingsJobModeSchema        CreateOracleInitialLoadSettingsJobModeEnum = "SCHEMA"
	CreateOracleInitialLoadSettingsJobModeTable         CreateOracleInitialLoadSettingsJobModeEnum = "TABLE"
	CreateOracleInitialLoadSettingsJobModeTablespace    CreateOracleInitialLoadSettingsJobModeEnum = "TABLESPACE"
	CreateOracleInitialLoadSettingsJobModeTransportable CreateOracleInitialLoadSettingsJobModeEnum = "TRANSPORTABLE"
)

var mappingCreateOracleInitialLoadSettingsJobModeEnum = map[string]CreateOracleInitialLoadSettingsJobModeEnum{
	"FULL":          CreateOracleInitialLoadSettingsJobModeFull,
	"SCHEMA":        CreateOracleInitialLoadSettingsJobModeSchema,
	"TABLE":         CreateOracleInitialLoadSettingsJobModeTable,
	"TABLESPACE":    CreateOracleInitialLoadSettingsJobModeTablespace,
	"TRANSPORTABLE": CreateOracleInitialLoadSettingsJobModeTransportable,
}

var mappingCreateOracleInitialLoadSettingsJobModeEnumLowerCase = map[string]CreateOracleInitialLoadSettingsJobModeEnum{
	"full":          CreateOracleInitialLoadSettingsJobModeFull,
	"schema":        CreateOracleInitialLoadSettingsJobModeSchema,
	"table":         CreateOracleInitialLoadSettingsJobModeTable,
	"tablespace":    CreateOracleInitialLoadSettingsJobModeTablespace,
	"transportable": CreateOracleInitialLoadSettingsJobModeTransportable,
}

// GetCreateOracleInitialLoadSettingsJobModeEnumValues Enumerates the set of values for CreateOracleInitialLoadSettingsJobModeEnum
func GetCreateOracleInitialLoadSettingsJobModeEnumValues() []CreateOracleInitialLoadSettingsJobModeEnum {
	values := make([]CreateOracleInitialLoadSettingsJobModeEnum, 0)
	for _, v := range mappingCreateOracleInitialLoadSettingsJobModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateOracleInitialLoadSettingsJobModeEnumStringValues Enumerates the set of values in String for CreateOracleInitialLoadSettingsJobModeEnum
func GetCreateOracleInitialLoadSettingsJobModeEnumStringValues() []string {
	return []string{
		"FULL",
		"SCHEMA",
		"TABLE",
		"TABLESPACE",
		"TRANSPORTABLE",
	}
}

// GetMappingCreateOracleInitialLoadSettingsJobModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateOracleInitialLoadSettingsJobModeEnum(val string) (CreateOracleInitialLoadSettingsJobModeEnum, bool) {
	enum, ok := mappingCreateOracleInitialLoadSettingsJobModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
