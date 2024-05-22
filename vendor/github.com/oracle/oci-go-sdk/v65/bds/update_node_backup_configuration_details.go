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

// UpdateNodeBackupConfigurationDetails The information about the NodeBackupConfiguration that is being updated.
type UpdateNodeBackupConfigurationDetails struct {
	LevelTypeDetails LevelTypeDetails `mandatory:"false" json:"levelTypeDetails"`

	// A user-friendly name. Only ASCII alphanumeric characters with no spaces allowed. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time zone of the execution schedule, in IANA time zone database name format
	Timezone *string `mandatory:"false" json:"timezone"`

	// Day/time recurrence (specified following RFC 5545) at which to trigger the backup process. Currently only DAILY, WEEKLY and MONTHLY frequency is supported. Days of the week are specified using BYDAY field. Time of the day is specified using BYHOUR. Other fields are not supported.
	Schedule *string `mandatory:"false" json:"schedule"`

	// Number of backup copies to retain.
	NumberOfBackupsToRetain *int `mandatory:"false" json:"numberOfBackupsToRetain"`

	// Incremental backup type includes only the changes since the last backup. Full backup type includes all changes since the volume was created.
	BackupType NodeBackupBackupTypeEnum `mandatory:"false" json:"backupType,omitempty"`
}

func (m UpdateNodeBackupConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateNodeBackupConfigurationDetails) ValidateEnumValue() (bool, error) {
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
func (m *UpdateNodeBackupConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		LevelTypeDetails        leveltypedetails         `json:"levelTypeDetails"`
		DisplayName             *string                  `json:"displayName"`
		Timezone                *string                  `json:"timezone"`
		Schedule                *string                  `json:"schedule"`
		NumberOfBackupsToRetain *int                     `json:"numberOfBackupsToRetain"`
		BackupType              NodeBackupBackupTypeEnum `json:"backupType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.LevelTypeDetails.UnmarshalPolymorphicJSON(model.LevelTypeDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.LevelTypeDetails = nn.(LevelTypeDetails)
	} else {
		m.LevelTypeDetails = nil
	}

	m.DisplayName = model.DisplayName

	m.Timezone = model.Timezone

	m.Schedule = model.Schedule

	m.NumberOfBackupsToRetain = model.NumberOfBackupsToRetain

	m.BackupType = model.BackupType

	return
}
