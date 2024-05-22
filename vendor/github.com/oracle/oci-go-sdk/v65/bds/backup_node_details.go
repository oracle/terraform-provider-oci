// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BackupNodeDetails The information about the nodes to backup.
type BackupNodeDetails struct {
	LevelTypeDetails LevelTypeDetails `mandatory:"true" json:"levelTypeDetails"`

	// Incremental backup type includes only the changes since the last backup. Full backup type includes all changes since the volume was created.
	BackupType NodeBackupBackupTypeEnum `mandatory:"false" json:"backupType,omitempty"`
}

func (m BackupNodeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackupNodeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingNodeBackupBackupTypeEnum(string(m.BackupType)); !ok && m.BackupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupType: %s. Supported values are: %s.", m.BackupType, strings.Join(GetNodeBackupBackupTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *BackupNodeDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		BackupType       NodeBackupBackupTypeEnum `json:"backupType"`
		LevelTypeDetails leveltypedetails         `json:"levelTypeDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.BackupType = model.BackupType

	nn, e = model.LevelTypeDetails.UnmarshalPolymorphicJSON(model.LevelTypeDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.LevelTypeDetails = nn.(LevelTypeDetails)
	} else {
		m.LevelTypeDetails = nil
	}

	return
}
