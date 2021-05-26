// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// AutonomousDatabaseBackupConfig Autonomous Database configuration details for storing manual backups (https://docs.cloud.oracle.com/Content/Database/Tasks/adbbackingup.htm) in the Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) service.
type AutonomousDatabaseBackupConfig struct {

	// Name of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) bucket to use for storing manual backups.
	ManualBackupBucketName *string `mandatory:"false" json:"manualBackupBucketName"`

	// The manual backup destination type.
	ManualBackupType AutonomousDatabaseBackupConfigManualBackupTypeEnum `mandatory:"false" json:"manualBackupType,omitempty"`
}

func (m AutonomousDatabaseBackupConfig) String() string {
	return common.PointerString(m)
}

// AutonomousDatabaseBackupConfigManualBackupTypeEnum Enum with underlying type: string
type AutonomousDatabaseBackupConfigManualBackupTypeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseBackupConfigManualBackupTypeEnum
const (
	AutonomousDatabaseBackupConfigManualBackupTypeNone        AutonomousDatabaseBackupConfigManualBackupTypeEnum = "NONE"
	AutonomousDatabaseBackupConfigManualBackupTypeObjectStore AutonomousDatabaseBackupConfigManualBackupTypeEnum = "OBJECT_STORE"
)

var mappingAutonomousDatabaseBackupConfigManualBackupType = map[string]AutonomousDatabaseBackupConfigManualBackupTypeEnum{
	"NONE":         AutonomousDatabaseBackupConfigManualBackupTypeNone,
	"OBJECT_STORE": AutonomousDatabaseBackupConfigManualBackupTypeObjectStore,
}

// GetAutonomousDatabaseBackupConfigManualBackupTypeEnumValues Enumerates the set of values for AutonomousDatabaseBackupConfigManualBackupTypeEnum
func GetAutonomousDatabaseBackupConfigManualBackupTypeEnumValues() []AutonomousDatabaseBackupConfigManualBackupTypeEnum {
	values := make([]AutonomousDatabaseBackupConfigManualBackupTypeEnum, 0)
	for _, v := range mappingAutonomousDatabaseBackupConfigManualBackupType {
		values = append(values, v)
	}
	return values
}
