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

// CreateDataPumpSettings Optional settings for Data Pump Export and Import jobs
type CreateDataPumpSettings struct {

	// Data Pump job mode.
	// Refer to link text (https://docs.oracle.com/en/database/oracle/oracle-database/19/sutil/oracle-data-pump-export-utility.html#GUID-8E497131-6B9B-4CC8-AA50-35F480CAC2C4)
	JobMode DataPumpJobModeEnum `mandatory:"false" json:"jobMode,omitempty"`

	DataPumpParameters *CreateDataPumpParameters `mandatory:"false" json:"dataPumpParameters"`

	// Defines remapping to be applied to objects as they are processed.
	// Refer to DATA_REMAP (https://docs.oracle.com/en/database/oracle/oracle-database/19/arpls/DBMS_DATAPUMP.html#GUID-E75AAE6F-4EA6-4737-A752-6B62F5E9D460)
	MetadataRemaps []MetadataRemap `mandatory:"false" json:"metadataRemaps"`

	TablespaceDetails CreateTargetTypeTablespaceDetails `mandatory:"false" json:"tablespaceDetails"`

	ExportDirectoryObject *CreateDirectoryObject `mandatory:"false" json:"exportDirectoryObject"`

	ImportDirectoryObject *CreateDirectoryObject `mandatory:"false" json:"importDirectoryObject"`
}

func (m CreateDataPumpSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDataPumpSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDataPumpJobModeEnum(string(m.JobMode)); !ok && m.JobMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JobMode: %s. Supported values are: %s.", m.JobMode, strings.Join(GetDataPumpJobModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateDataPumpSettings) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		JobMode               DataPumpJobModeEnum               `json:"jobMode"`
		DataPumpParameters    *CreateDataPumpParameters         `json:"dataPumpParameters"`
		MetadataRemaps        []MetadataRemap                   `json:"metadataRemaps"`
		TablespaceDetails     createtargettypetablespacedetails `json:"tablespaceDetails"`
		ExportDirectoryObject *CreateDirectoryObject            `json:"exportDirectoryObject"`
		ImportDirectoryObject *CreateDirectoryObject            `json:"importDirectoryObject"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.JobMode = model.JobMode

	m.DataPumpParameters = model.DataPumpParameters

	m.MetadataRemaps = make([]MetadataRemap, len(model.MetadataRemaps))
	copy(m.MetadataRemaps, model.MetadataRemaps)
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

	return
}
