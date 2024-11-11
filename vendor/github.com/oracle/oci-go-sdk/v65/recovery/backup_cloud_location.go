// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database Autonomous Recovery Service API
//
// Use Oracle Database Autonomous Recovery Service API to manage Protected Databases.
//

package recovery

import (
	"strings"
)

// BackupCloudLocationEnum Enum with underlying type: string
type BackupCloudLocationEnum string

// Set of constants representing the allowable values for BackupCloudLocationEnum
const (
	BackupCloudLocationAzure BackupCloudLocationEnum = "AZURE"
	BackupCloudLocationOci   BackupCloudLocationEnum = "OCI"
	BackupCloudLocationGcp   BackupCloudLocationEnum = "GCP"
	BackupCloudLocationAws   BackupCloudLocationEnum = "AWS"
)

var mappingBackupCloudLocationEnum = map[string]BackupCloudLocationEnum{
	"AZURE": BackupCloudLocationAzure,
	"OCI":   BackupCloudLocationOci,
	"GCP":   BackupCloudLocationGcp,
	"AWS":   BackupCloudLocationAws,
}

var mappingBackupCloudLocationEnumLowerCase = map[string]BackupCloudLocationEnum{
	"azure": BackupCloudLocationAzure,
	"oci":   BackupCloudLocationOci,
	"gcp":   BackupCloudLocationGcp,
	"aws":   BackupCloudLocationAws,
}

// GetBackupCloudLocationEnumValues Enumerates the set of values for BackupCloudLocationEnum
func GetBackupCloudLocationEnumValues() []BackupCloudLocationEnum {
	values := make([]BackupCloudLocationEnum, 0)
	for _, v := range mappingBackupCloudLocationEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupCloudLocationEnumStringValues Enumerates the set of values in String for BackupCloudLocationEnum
func GetBackupCloudLocationEnumStringValues() []string {
	return []string{
		"AZURE",
		"OCI",
		"GCP",
		"AWS",
	}
}

// GetMappingBackupCloudLocationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupCloudLocationEnum(val string) (BackupCloudLocationEnum, bool) {
	enum, ok := mappingBackupCloudLocationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
