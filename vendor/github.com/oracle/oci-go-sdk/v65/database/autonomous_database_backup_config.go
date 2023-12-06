// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutonomousDatabaseBackupConfig Autonomous Database configuration details for storing manual backups (https://docs.oracle.com/en/cloud/paas/autonomous-database/adbsa/backup-restore.html#GUID-9035DFB8-4702-4CEB-8281-C2A303820809) in the Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) service.
type AutonomousDatabaseBackupConfig struct {

	// Name of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) bucket to use for storing manual backups.
	ManualBackupBucketName *string `mandatory:"false" json:"manualBackupBucketName"`

	// The manual backup destination type.
	ManualBackupType AutonomousDatabaseBackupConfigManualBackupTypeEnum `mandatory:"false" json:"manualBackupType,omitempty"`
}

func (m AutonomousDatabaseBackupConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDatabaseBackupConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAutonomousDatabaseBackupConfigManualBackupTypeEnum(string(m.ManualBackupType)); !ok && m.ManualBackupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManualBackupType: %s. Supported values are: %s.", m.ManualBackupType, strings.Join(GetAutonomousDatabaseBackupConfigManualBackupTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousDatabaseBackupConfigManualBackupTypeEnum Enum with underlying type: string
type AutonomousDatabaseBackupConfigManualBackupTypeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseBackupConfigManualBackupTypeEnum
const (
	AutonomousDatabaseBackupConfigManualBackupTypeNone        AutonomousDatabaseBackupConfigManualBackupTypeEnum = "NONE"
	AutonomousDatabaseBackupConfigManualBackupTypeObjectStore AutonomousDatabaseBackupConfigManualBackupTypeEnum = "OBJECT_STORE"
)

var mappingAutonomousDatabaseBackupConfigManualBackupTypeEnum = map[string]AutonomousDatabaseBackupConfigManualBackupTypeEnum{
	"NONE":         AutonomousDatabaseBackupConfigManualBackupTypeNone,
	"OBJECT_STORE": AutonomousDatabaseBackupConfigManualBackupTypeObjectStore,
}

var mappingAutonomousDatabaseBackupConfigManualBackupTypeEnumLowerCase = map[string]AutonomousDatabaseBackupConfigManualBackupTypeEnum{
	"none":         AutonomousDatabaseBackupConfigManualBackupTypeNone,
	"object_store": AutonomousDatabaseBackupConfigManualBackupTypeObjectStore,
}

// GetAutonomousDatabaseBackupConfigManualBackupTypeEnumValues Enumerates the set of values for AutonomousDatabaseBackupConfigManualBackupTypeEnum
func GetAutonomousDatabaseBackupConfigManualBackupTypeEnumValues() []AutonomousDatabaseBackupConfigManualBackupTypeEnum {
	values := make([]AutonomousDatabaseBackupConfigManualBackupTypeEnum, 0)
	for _, v := range mappingAutonomousDatabaseBackupConfigManualBackupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseBackupConfigManualBackupTypeEnumStringValues Enumerates the set of values in String for AutonomousDatabaseBackupConfigManualBackupTypeEnum
func GetAutonomousDatabaseBackupConfigManualBackupTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"OBJECT_STORE",
	}
}

// GetMappingAutonomousDatabaseBackupConfigManualBackupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseBackupConfigManualBackupTypeEnum(val string) (AutonomousDatabaseBackupConfigManualBackupTypeEnum, bool) {
	enum, ok := mappingAutonomousDatabaseBackupConfigManualBackupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
