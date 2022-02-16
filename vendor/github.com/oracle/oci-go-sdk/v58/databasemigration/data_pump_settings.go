// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// DataPumpSettings Optional settings for Data Pump Export and Import jobs
type DataPumpSettings struct {

	// Data Pump job mode.
	// Refer to Data Pump Export Modes  (https://docs.oracle.com/en/database/oracle/oracle-database/19/sutil/oracle-data-pump-export-utility.html#GUID-8E497131-6B9B-4CC8-AA50-35F480CAC2C4)
	JobMode DataPumpJobModeEnum `mandatory:"false" json:"jobMode,omitempty"`

	DataPumpParameters *DataPumpParameters `mandatory:"false" json:"dataPumpParameters"`

	// Defines remapping to be applied to objects as they are processed.
	// Refer to METADATA_REMAP Procedure  (https://docs.oracle.com/en/database/oracle/oracle-database/19/arpls/DBMS_DATAPUMP.html#GUID-0FC32790-91E6-4781-87A3-229DE024CB3D)
	MetadataRemaps []MetadataRemap `mandatory:"false" json:"metadataRemaps"`

	ExportDirectoryObject *DirectoryObject `mandatory:"false" json:"exportDirectoryObject"`

	ImportDirectoryObject *DirectoryObject `mandatory:"false" json:"importDirectoryObject"`
}

func (m DataPumpSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataPumpSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDataPumpJobModeEnum(string(m.JobMode)); !ok && m.JobMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JobMode: %s. Supported values are: %s.", m.JobMode, strings.Join(GetDataPumpJobModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
