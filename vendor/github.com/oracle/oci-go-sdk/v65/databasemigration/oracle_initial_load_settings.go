// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// OracleInitialLoadSettings Optional settings for Data Pump Export and Import jobs
type OracleInitialLoadSettings struct {

	// Oracle job mode.
	JobMode OracleInitialLoadSettingsJobModeEnum `mandatory:"true" json:"jobMode"`

	DataPumpParameters *DataPumpParameters `mandatory:"false" json:"dataPumpParameters"`

	TablespaceDetails TargetTypeTablespaceDetails `mandatory:"false" json:"tablespaceDetails"`

	ExportDirectoryObject *DirectoryObject `mandatory:"false" json:"exportDirectoryObject"`

	ImportDirectoryObject *DirectoryObject `mandatory:"false" json:"importDirectoryObject"`

	// Defines remapping to be applied to objects as they are processed.
	MetadataRemaps []MetadataRemap `mandatory:"false" json:"metadataRemaps"`
}

func (m OracleInitialLoadSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OracleInitialLoadSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOracleInitialLoadSettingsJobModeEnum(string(m.JobMode)); !ok && m.JobMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JobMode: %s. Supported values are: %s.", m.JobMode, strings.Join(GetOracleInitialLoadSettingsJobModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *OracleInitialLoadSettings) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DataPumpParameters    *DataPumpParameters                  `json:"dataPumpParameters"`
		TablespaceDetails     targettypetablespacedetails          `json:"tablespaceDetails"`
		ExportDirectoryObject *DirectoryObject                     `json:"exportDirectoryObject"`
		ImportDirectoryObject *DirectoryObject                     `json:"importDirectoryObject"`
		MetadataRemaps        []MetadataRemap                      `json:"metadataRemaps"`
		JobMode               OracleInitialLoadSettingsJobModeEnum `json:"jobMode"`
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
		m.TablespaceDetails = nn.(TargetTypeTablespaceDetails)
	} else {
		m.TablespaceDetails = nil
	}

	m.ExportDirectoryObject = model.ExportDirectoryObject

	m.ImportDirectoryObject = model.ImportDirectoryObject

	m.MetadataRemaps = make([]MetadataRemap, len(model.MetadataRemaps))
	copy(m.MetadataRemaps, model.MetadataRemaps)
	m.JobMode = model.JobMode

	return
}

// OracleInitialLoadSettingsJobModeEnum Enum with underlying type: string
type OracleInitialLoadSettingsJobModeEnum string

// Set of constants representing the allowable values for OracleInitialLoadSettingsJobModeEnum
const (
	OracleInitialLoadSettingsJobModeFull          OracleInitialLoadSettingsJobModeEnum = "FULL"
	OracleInitialLoadSettingsJobModeSchema        OracleInitialLoadSettingsJobModeEnum = "SCHEMA"
	OracleInitialLoadSettingsJobModeTable         OracleInitialLoadSettingsJobModeEnum = "TABLE"
	OracleInitialLoadSettingsJobModeTablespace    OracleInitialLoadSettingsJobModeEnum = "TABLESPACE"
	OracleInitialLoadSettingsJobModeTransportable OracleInitialLoadSettingsJobModeEnum = "TRANSPORTABLE"
)

var mappingOracleInitialLoadSettingsJobModeEnum = map[string]OracleInitialLoadSettingsJobModeEnum{
	"FULL":          OracleInitialLoadSettingsJobModeFull,
	"SCHEMA":        OracleInitialLoadSettingsJobModeSchema,
	"TABLE":         OracleInitialLoadSettingsJobModeTable,
	"TABLESPACE":    OracleInitialLoadSettingsJobModeTablespace,
	"TRANSPORTABLE": OracleInitialLoadSettingsJobModeTransportable,
}

var mappingOracleInitialLoadSettingsJobModeEnumLowerCase = map[string]OracleInitialLoadSettingsJobModeEnum{
	"full":          OracleInitialLoadSettingsJobModeFull,
	"schema":        OracleInitialLoadSettingsJobModeSchema,
	"table":         OracleInitialLoadSettingsJobModeTable,
	"tablespace":    OracleInitialLoadSettingsJobModeTablespace,
	"transportable": OracleInitialLoadSettingsJobModeTransportable,
}

// GetOracleInitialLoadSettingsJobModeEnumValues Enumerates the set of values for OracleInitialLoadSettingsJobModeEnum
func GetOracleInitialLoadSettingsJobModeEnumValues() []OracleInitialLoadSettingsJobModeEnum {
	values := make([]OracleInitialLoadSettingsJobModeEnum, 0)
	for _, v := range mappingOracleInitialLoadSettingsJobModeEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleInitialLoadSettingsJobModeEnumStringValues Enumerates the set of values in String for OracleInitialLoadSettingsJobModeEnum
func GetOracleInitialLoadSettingsJobModeEnumStringValues() []string {
	return []string{
		"FULL",
		"SCHEMA",
		"TABLE",
		"TABLESPACE",
		"TRANSPORTABLE",
	}
}

// GetMappingOracleInitialLoadSettingsJobModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleInitialLoadSettingsJobModeEnum(val string) (OracleInitialLoadSettingsJobModeEnum, bool) {
	enum, ok := mappingOracleInitialLoadSettingsJobModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
